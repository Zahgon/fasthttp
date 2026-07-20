package fasthttp

import (
	"crypto/tls"
	"net"
	"sync"
)

type perIPConnCounter struct {
	perIPConnPool    sync.Pool
	perIPTLSConnPool sync.Pool
	m                map[uint32]int
	lock             sync.Mutex
}

func (cc *perIPConnCounter) Register(ip uint32) int { _ = "STUB: not implemented"; return 0 }

func (cc *perIPConnCounter) Unregister(ip uint32) { _ = "STUB: not implemented"; return }

type perIPConn struct {
	net.Conn

	perIPConnCounter *perIPConnCounter

	ip   uint32
	lock sync.Mutex
}

type perIPTLSConn struct {
	*tls.Conn

	perIPConnCounter *perIPConnCounter

	ip   uint32
	lock sync.Mutex
}

func acquirePerIPConn(conn net.Conn, ip uint32, counter *perIPConnCounter) net.Conn {
	_ = "STUB: not implemented"
	return *new(net.Conn)
}

//nolint:forcetypeassert

//nolint:forcetypeassert

func (c *perIPConn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *perIPTLSConn) Close() error { _ = "STUB: not implemented"; return nil }

func getUint32IP(c net.Conn) uint32 { _ = "STUB: not implemented"; return 0 }

func getConnIP4(c net.Conn) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }
