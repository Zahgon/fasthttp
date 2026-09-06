package tcplisten

import (
	"net"
)

type Config struct {
	ReusePort   bool
	DeferAccept bool
	FastOpen    bool
	Backlog     int
}

func (cfg *Config) NewListener(network, addr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}
