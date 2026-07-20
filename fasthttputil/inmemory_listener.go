package fasthttputil

import (
	"errors"
	"net"
	"sync"
)

var ErrInmemoryListenerClosed = errors.New("fasthttputil: inmemorylistener is already closed: use of closed network connection")

type InmemoryListener struct {
	listenerAddr net.Addr
	conns        chan acceptConn
	done         chan struct{}
	addrLock     sync.RWMutex
	lock         sync.Mutex
	closed       bool
}

type acceptConn struct {
	conn     net.Conn
	accepted chan struct{}
}

func NewInmemoryListener() *InmemoryListener { _ = "STUB: not implemented"; return nil }

func (ln *InmemoryListener) SetLocalAddr(localAddr net.Addr) { _ = "STUB: not implemented"; return }

func (ln *InmemoryListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (ln *InmemoryListener) Close() error { _ = "STUB: not implemented"; return nil }

func (ln *InmemoryListener) closePendingConns() { _ = "STUB: not implemented"; return }

type inmemoryAddr int

func (inmemoryAddr) Network() string { _ = "STUB: not implemented"; return "" }

func (inmemoryAddr) String() string { _ = "STUB: not implemented"; return "" }

func (ln *InmemoryListener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (ln *InmemoryListener) Dial() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (ln *InmemoryListener) DialWithLocalAddr(local net.Addr) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
