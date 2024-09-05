package all

import (
	"github.com/xvguardian/xray-core-modified/main/commands/all/api"
	"github.com/xvguardian/xray-core-modified/main/commands/all/convert"
	"github.com/xvguardian/xray-core-modified/main/commands/all/tls"
	"github.com/xvguardian/xray-core-modified/main/commands/base"
)

// go:generate go run github.com/xtls/xray-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		convert.CmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
