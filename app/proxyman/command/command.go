package command

import (
	"context"

	"github.com/xvguardian/xray-core-modified/common"
	"github.com/xvguardian/xray-core-modified/common/errors"
	"github.com/xvguardian/xray-core-modified/core"
	"github.com/xvguardian/xray-core-modified/features/inbound"
	"github.com/xvguardian/xray-core-modified/features/outbound"
	"github.com/xvguardian/xray-core-modified/proxy"
	grpc "google.golang.org/grpc"
)

// InboundOperation is the interface for operations that applies to inbound handlers.
type InboundOperation interface {
	// ApplyInbound applies this operation to the given inbound handler.
	ApplyInbound(context.Context, inbound.Handler) error
}

// OutboundOperation is the interface for operations that applies to outbound handlers.
type OutboundOperation interface {
	// ApplyOutbound applies this operation to the given outbound handler.
	ApplyOutbound(context.Context, outbound.Handler) error
}

func getInbound(handler inbound.Handler) (proxy.Inbound, error) {
	gi, ok := handler.(proxy.GetInbound)
	if !ok {
		return nil, errors.New("can't get inbound proxy from handler.")
	}
	return gi.GetInbound(), nil
}

// ApplyInbound implements InboundOperation.
func (op *AddUserOperation) ApplyInbound(ctx context.Context, handler inbound.Handler) error {
	p, err := getInbound(handler)
	if err != nil {
		return err
	}
	um, ok := p.(proxy.UserManager)
	if !ok {
		return errors.New("proxy is not a UserManager")
	}
	mUser, err := op.User.ToMemoryUser()
	if err != nil {
		return errors.New("failed to parse user").Base(err)
	}
	return um.AddUser(ctx, mUser)
}

// ApplyInbound implements InboundOperation.
func (op *RemoveUserOperation) ApplyInbound(ctx context.Context, handler inbound.Handler) error {
	p, err := getInbound(handler)
	if err != nil {
		return err
	}
	um, ok := p.(proxy.UserManager)
	if !ok {
		return errors.New("proxy is not a UserManager")
	}
	return um.RemoveUser(ctx, op.Email)
}

type handlerServer struct {
	s   *core.Instance
	ihm inbound.Manager
	ohm outbound.Manager
}

func (s *handlerServer) AddInbound(ctx context.Context, request *AddInboundRequest) (*AddInboundResponse, error) {
	if err := core.AddInboundHandler(s.s, request.Inbound); err != nil {
		return nil, err
	}

	return &AddInboundResponse{}, nil
}

func (s *handlerServer) RemoveInbound(ctx context.Context, request *RemoveInboundRequest) (*RemoveInboundResponse, error) {
	return &RemoveInboundResponse{}, s.ihm.RemoveHandler(ctx, request.Tag)
}

func (s *handlerServer) AlterInbound(ctx context.Context, request *AlterInboundRequest) (*AlterInboundResponse, error) {
	rawOperation, err := request.Operation.GetInstance()
	if err != nil {
		return nil, errors.New("unknown operation").Base(err)
	}
	operation, ok := rawOperation.(InboundOperation)
	if !ok {
		return nil, errors.New("not an inbound operation")
	}

	handler, err := s.ihm.GetHandler(ctx, request.Tag)
	if err != nil {
		return nil, errors.New("failed to get handler: ", request.Tag).Base(err)
	}

	return &AlterInboundResponse{}, operation.ApplyInbound(ctx, handler)
}

func (s *handlerServer) AddOutbound(ctx context.Context, request *AddOutboundRequest) (*AddOutboundResponse, error) {
	if err := core.AddOutboundHandler(s.s, request.Outbound); err != nil {
		return nil, err
	}
	return &AddOutboundResponse{}, nil
}

func (s *handlerServer) RemoveOutbound(ctx context.Context, request *RemoveOutboundRequest) (*RemoveOutboundResponse, error) {
	return &RemoveOutboundResponse{}, s.ohm.RemoveHandler(ctx, request.Tag)
}

func (s *handlerServer) AlterOutbound(ctx context.Context, request *AlterOutboundRequest) (*AlterOutboundResponse, error) {
	rawOperation, err := request.Operation.GetInstance()
	if err != nil {
		return nil, errors.New("unknown operation").Base(err)
	}
	operation, ok := rawOperation.(OutboundOperation)
	if !ok {
		return nil, errors.New("not an outbound operation")
	}

	handler := s.ohm.GetHandler(request.Tag)
	return &AlterOutboundResponse{}, operation.ApplyOutbound(ctx, handler)
}

func (s *handlerServer) mustEmbedUnimplementedHandlerServiceServer() {}

type service struct {
	v *core.Instance
}

func (s *service) Register(server *grpc.Server) {
	hs := &handlerServer{
		s: s.v,
	}
	common.Must(s.v.RequireFeatures(func(im inbound.Manager, om outbound.Manager) {
		hs.ihm = im
		hs.ohm = om
	}))
	RegisterHandlerServiceServer(server, hs)

	// For compatibility purposes
	vCoreDesc := HandlerService_ServiceDesc
	vCoreDesc.ServiceName = "v2ray.core.app.proxyman.command.HandlerService"
	server.RegisterService(&vCoreDesc, hs)
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, cfg interface{}) (interface{}, error) {
		s := core.MustFromContext(ctx)
		return &service{v: s}, nil
	}))
}
