package fasthttpproxy

import (
	"github.com/valyala/fasthttp"
)

func FasthttpSocksDialer(proxyAddr string) fasthttp.DialFunc {
	_ = "STUB: not implemented"
	return *new(fasthttp.DialFunc)
}

func FasthttpSocksDialerDualStack(proxyAddr string) fasthttp.DialFunc {
	_ = "STUB: not implemented"
	return *new(fasthttp.DialFunc)
}
