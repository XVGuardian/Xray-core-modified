package udp

import (
	"github.com/xvguardian/xray-core-modified/common/buf"
	"github.com/xvguardian/xray-core-modified/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
