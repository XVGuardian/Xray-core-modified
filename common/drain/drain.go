package drain

import "io"

//go:generate go run github.com/xvguardian/xray-core-modified/common/errors/errorgen

type Drainer interface {
	AcknowledgeReceive(size int)
	Drain(reader io.Reader) error
}
