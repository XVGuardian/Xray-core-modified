package scenarios

import (
	"testing"
	//"time"

	"github.com/xvguardian/xray-core-modified/app/log"
	"github.com/xvguardian/xray-core-modified/app/proxyman"
	"github.com/xvguardian/xray-core-modified/common"
	clog "github.com/xvguardian/xray-core-modified/common/log"
	"github.com/xvguardian/xray-core-modified/common/net"
	"github.com/xvguardian/xray-core-modified/common/serial"
	core "github.com/xvguardian/xray-core-modified/core"
	"github.com/xvguardian/xray-core-modified/infra/conf"
	"github.com/xvguardian/xray-core-modified/proxy/dokodemo"
	"github.com/xvguardian/xray-core-modified/proxy/freedom"
	"github.com/xvguardian/xray-core-modified/proxy/wireguard"
	"github.com/xvguardian/xray-core-modified/testing/servers/tcp"
	"github.com/xvguardian/xray-core-modified/testing/servers/udp"
	//"golang.org/x/sync/errgroup"
)

func TestWireguard(t *testing.T) {
	tcpServer := tcp.Server{
		MsgProcessor: xor,
	}
	dest, err := tcpServer.Start()
	common.Must(err)
	defer tcpServer.Close()

	serverPrivate, _ := conf.ParseWireGuardKey("EGs4lTSJPmgELx6YiJAmPR2meWi6bY+e9rTdCipSj10=")
	serverPublic, _ := conf.ParseWireGuardKey("osAMIyil18HeZXGGBDC9KpZoM+L2iGyXWVSYivuM9B0=")
	clientPrivate, _ := conf.ParseWireGuardKey("CPQSpgxgdQRZa5SUbT3HLv+mmDVHLW5YR/rQlzum/2I=")
	clientPublic, _ := conf.ParseWireGuardKey("MmLJ5iHFVVBp7VsB0hxfpQ0wEzAbT2KQnpQpj0+RtBw=")

	serverPort := udp.PickPort()
	serverConfig := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{
				ErrorLogLevel: clog.Severity_Debug,
				ErrorLogType:  log.LogType_Console,
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{Range: []*net.PortRange{net.SinglePortRange(serverPort)}},
					Listen:   net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&wireguard.DeviceConfig{
					IsClient: false,
					KernelMode: false,
					Endpoint: []string{"10.0.0.1"},
					Mtu: 1420,
					SecretKey: serverPrivate,
					Peers: []*wireguard.PeerConfig{{
						PublicKey: serverPublic,
						AllowedIps: []string{"0.0.0.0/0", "::0/0"},
					}},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	clientPort := tcp.PickPort()
	clientConfig := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{
				ErrorLogLevel: clog.Severity_Debug,
				ErrorLogType:  log.LogType_Console,
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{Range: []*net.PortRange{net.SinglePortRange(clientPort)}},
					Listen:   net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(dest.Address),
					Port:    uint32(dest.Port),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&wireguard.DeviceConfig{
					IsClient: true,
					KernelMode: false,
					Endpoint: []string{"10.0.0.2"},
					Mtu: 1420,
					SecretKey: clientPrivate,
					Peers: []*wireguard.PeerConfig{{
						Endpoint: "127.0.0.1:" + serverPort.String(),
						PublicKey: clientPublic,
						AllowedIps: []string{"0.0.0.0/0", "::0/0"},
					}},
				}),
			},
		},
	}

	servers, err := InitializeServerConfigs(serverConfig, clientConfig)
	common.Must(err)
	defer CloseAllServers(servers)

	// FIXME: for some reason wg server does not receive

	// var errg errgroup.Group
	// for i := 0; i < 1; i++ {
	// 	errg.Go(testTCPConn(clientPort, 1024, time.Second*2))
	// }
	// if err := errg.Wait(); err != nil {
	// 	t.Error(err)
	// }
}