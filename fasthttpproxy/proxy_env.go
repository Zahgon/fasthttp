package fasthttpproxy

import (
	"time"

	"github.com/valyala/fasthttp"
)

const (
	httpsScheme = "https"
	httpScheme  = "http"
)

func FasthttpProxyHTTPDialer() fasthttp.DialFunc {
	_ = "STUB: not implemented"
	return *new(fasthttp.DialFunc)
}

func FasthttpProxyHTTPDialerTimeout(timeout time.Duration) fasthttp.DialFunc {
	_ = "STUB: not implemented"
	return *new(fasthttp.DialFunc)
}
