package fasthttpproxy

import (
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
	"golang.org/x/net/http/httpproxy"
	"golang.org/x/net/proxy"
)

var (
	colonTLSPort = ":443"
	tmpURL       = &url.URL{Scheme: httpsScheme, Host: "example.com"}
)

func dialFuncOrError(dialFunc fasthttp.DialFunc, err error) fasthttp.DialFunc {
	_ = "STUB: not implemented"
	return *new(fasthttp.DialFunc)
}

type Dialer struct {
	fasthttp.TCPDialer

	httpproxy.Config

	DialDualStack bool

	Timeout time.Duration

	ConnectTimeout time.Duration
}

func (d *Dialer) GetDialFunc(useEnv bool) (fasthttp.DialFunc, error) {
	_ = "STUB: not implemented"
	return *new(fasthttp.DialFunc), nil
}

func (d *Dialer) Dial(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d *Dialer) connectTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type httpProxyDialer interface {
	connectTimeout() time.Duration
}

type DialerFunc func(network, addr string) (net.Conn, error)

func (d DialerFunc) Dial(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func httpProxyDial(dialer proxy.Dialer, network, addr, proxyAddr, auth string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

type proxyInfo struct {
	auth string
	addr string
}

func addrAndAuth(pu *url.URL, authCache *sync.Map) (proxyAddr, auth string) {
	_ = "STUB: not implemented"
	return "", ""
}

//nolint:forcetypeassert
