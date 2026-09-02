//go:build linux || darwin || dragonfly || freebsd || netbsd || openbsd || rumprun || (zos && s390x)

package tcplisten

import (
	"net"

	"golang.org/x/sys/unix"
)

type Config struct {
	ReusePort bool

	DeferAccept bool

	FastOpen bool

	Backlog int
}

func (cfg *Config) NewListener(network, addr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (cfg *Config) fdSetup(fd int, sa unix.Sockaddr, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

func getSockaddr(network, addr string) (sa unix.Sockaddr, soType int, err error) {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr), 0, nil
}

func safeIntToUint32(i int) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func safeIntToUintptr(i int) (uintptr, error) { _ = "STUB: not implemented"; return 0, nil }
