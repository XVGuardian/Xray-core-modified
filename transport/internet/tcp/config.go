package tcp

import (
	"github.com/xvguardian/xray-core-modified/common"
	"github.com/xvguardian/xray-core-modified/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
