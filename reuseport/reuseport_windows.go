package reuseport

import (
	"net"
	"syscall"

	"golang.org/x/sys/windows"
)

var listenConfig = net.ListenConfig{
	Control: func(network, address string, c syscall.RawConn) (err error) {
		return c.Control(func(fd uintptr) {
			err = windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1)
		})
	},
}

func Listen(network, addr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}
