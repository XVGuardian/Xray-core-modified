package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/xvguardian/xray-core-modified/app/dispatcher"
	_ "github.com/xvguardian/xray-core-modified/app/proxyman/inbound"
	_ "github.com/xvguardian/xray-core-modified/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/xvguardian/xray-core-modified/app/commander"
	_ "github.com/xvguardian/xray-core-modified/app/log/command"
	_ "github.com/xvguardian/xray-core-modified/app/proxyman/command"
	_ "github.com/xvguardian/xray-core-modified/app/stats/command"

	// Developer preview services
	_ "github.com/xvguardian/xray-core-modified/app/observatory/command"

	// Other optional features.
	_ "github.com/xvguardian/xray-core-modified/app/dns"
	_ "github.com/xvguardian/xray-core-modified/app/dns/fakedns"
	_ "github.com/xvguardian/xray-core-modified/app/log"
	_ "github.com/xvguardian/xray-core-modified/app/metrics"
	_ "github.com/xvguardian/xray-core-modified/app/policy"
	_ "github.com/xvguardian/xray-core-modified/app/reverse"
	_ "github.com/xvguardian/xray-core-modified/app/router"
	_ "github.com/xvguardian/xray-core-modified/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/xvguardian/xray-core-modified/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/xvguardian/xray-core-modified/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/xvguardian/xray-core-modified/proxy/blackhole"
	_ "github.com/xvguardian/xray-core-modified/proxy/dns"
	_ "github.com/xvguardian/xray-core-modified/proxy/dokodemo"
	_ "github.com/xvguardian/xray-core-modified/proxy/freedom"
	_ "github.com/xvguardian/xray-core-modified/proxy/http"
	_ "github.com/xvguardian/xray-core-modified/proxy/loopback"
	_ "github.com/xvguardian/xray-core-modified/proxy/shadowsocks"
	_ "github.com/xvguardian/xray-core-modified/proxy/socks"
	_ "github.com/xvguardian/xray-core-modified/proxy/trojan"
	_ "github.com/xvguardian/xray-core-modified/proxy/vless/inbound"
	_ "github.com/xvguardian/xray-core-modified/proxy/vless/outbound"
	_ "github.com/xvguardian/xray-core-modified/proxy/vmess/inbound"
	_ "github.com/xvguardian/xray-core-modified/proxy/vmess/outbound"
	_ "github.com/xvguardian/xray-core-modified/proxy/wireguard"

	// Transports
	_ "github.com/xvguardian/xray-core-modified/transport/internet/domainsocket"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/grpc"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/http"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/httpupgrade"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/kcp"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/quic"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/reality"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/splithttp"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/tcp"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/tls"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/udp"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/websocket"

	// Transport headers
	_ "github.com/xvguardian/xray-core-modified/transport/internet/headers/http"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/headers/noop"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/headers/srtp"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/headers/tls"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/headers/utp"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/headers/wechat"
	_ "github.com/xvguardian/xray-core-modified/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/xvguardian/xray-core-modified/main/json"
	_ "github.com/xvguardian/xray-core-modified/main/toml"
	_ "github.com/xvguardian/xray-core-modified/main/yaml"

	// Load config from file or http(s)
	_ "github.com/xvguardian/xray-core-modified/main/confloader/external"

	// Commands
	_ "github.com/xvguardian/xray-core-modified/main/commands/all"
)
