package dns_test

import (
	"context"
	"testing"
	"time"

	. "github.com/xvguardian/xray-core-modified/app/dns"
	"github.com/xvguardian/xray-core-modified/common"
	"github.com/xvguardian/xray-core-modified/common/net"
	"github.com/xvguardian/xray-core-modified/features/dns"
)

func TestLocalNameServer(t *testing.T) {
	s := NewLocalNameServer()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	ips, err := s.QueryIP(ctx, "google.com", net.IP{}, dns.IPOption{
		IPv4Enable: true,
		IPv6Enable: true,
		FakeEnable: false,
	}, false)
	cancel()
	common.Must(err)
	if len(ips) == 0 {
		t.Error("expect some ips, but got 0")
	}
}
