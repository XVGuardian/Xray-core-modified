package command_test

import (
	"context"
	"testing"

	"github.com/xvguardian/xray-core-modified/app/dispatcher"
	"github.com/xvguardian/xray-core-modified/app/log"
	. "github.com/xvguardian/xray-core-modified/app/log/command"
	"github.com/xvguardian/xray-core-modified/app/proxyman"
	_ "github.com/xvguardian/xray-core-modified/app/proxyman/inbound"
	_ "github.com/xvguardian/xray-core-modified/app/proxyman/outbound"
	"github.com/xvguardian/xray-core-modified/common"
	"github.com/xvguardian/xray-core-modified/common/serial"
	"github.com/xvguardian/xray-core-modified/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
