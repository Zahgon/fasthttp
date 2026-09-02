//go:build !windows && !aix && !solaris

package reuseport

import (
	"net"

	"github.com/valyala/fasthttp/tcplisten"
)

func Listen(network, addr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

var cfg = &tcplisten.Config{
	ReusePort:   true,
	DeferAccept: true,
	FastOpen:    true,
}
