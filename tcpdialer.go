package fasthttp

import (
	"context"
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

func Dial(addr string) (net.Conn, error) { _ = "STUB: not implemented"; return *new(net.Conn), nil }

func DialTimeout(addr string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func DialDualStack(addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func DialDualStackTimeout(addr string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

var defaultDialer = &TCPDialer{Concurrency: 1000}

type Resolver interface {
	LookupIPAddr(context.Context, string) (names []net.IPAddr, err error)
}

type TCPDialer struct {
	Resolver Resolver

	LocalAddr *net.TCPAddr

	concurrencyCh chan struct{}

	tcpAddrsMap    sync.Map
	cleanerRunning atomic.Bool

	Concurrency int

	DNSCacheDuration time.Duration

	once sync.Once

	DisableDNSResolution bool
}

func (d *TCPDialer) Dial(addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d *TCPDialer) DialTimeout(addr string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d *TCPDialer) DialDualStack(addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d *TCPDialer) DialDualStackTimeout(addr string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d *TCPDialer) FlushDNSCache() { _ = "STUB: not implemented"; return }

func FlushDNSCache() { _ = "STUB: not implemented"; return }

func (d *TCPDialer) dial(addr string, dualStack bool, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d *TCPDialer) tryDial(
	network string, addr string, deadline time.Time, concurrencyCh chan struct{},
) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

var ErrDialTimeout = errors.New("fasthttp: dialing to the given tcp address timed out")

type ErrDialWithUpstream struct {
	wrapErr  error
	Upstream string
}

func (e *ErrDialWithUpstream) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ErrDialWithUpstream) Unwrap() error { _ = "STUB: not implemented"; return nil }

func wrapDialWithUpstream(err error, upstream string) error { _ = "STUB: not implemented"; return nil }

const DefaultDialTimeout = 3 * time.Second

type tcpAddrEntry struct {
	resolveTime time.Time
	addrs       []net.TCPAddr
	addrsIdx    uint32

	pending int32
}

const DefaultDNSCacheDuration = time.Minute

func (d *TCPDialer) cleanExpiredDNSEntries() bool { _ = "STUB: not implemented"; return false }

func (d *TCPDialer) startTCPAddrsClean() { _ = "STUB: not implemented"; return }

func (d *TCPDialer) hasTCPAddrsEntries() bool { _ = "STUB: not implemented"; return false }

var tcpAddrsCleanInterval = int64(time.Second)

func (d *TCPDialer) tcpAddrsClean() { _ = "STUB: not implemented"; return }

func (d *TCPDialer) getTCPAddrs(addr string, dualStack bool, deadline time.Time) ([]net.TCPAddr, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func resolveTCPAddrs(addr string, dualStack bool, resolver Resolver, deadline time.Time) ([]net.TCPAddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var errNoDNSEntries = errors.New("couldn't find dns entries for the given domain: try using dual-stack dialing")
