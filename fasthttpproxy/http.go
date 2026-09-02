package fasthttpproxy

import (
	"time"

	"github.com/valyala/fasthttp"
)

func FasthttpHTTPDialer(proxy string) fasthttp.DialFunc {
	_ = "STUB: not implemented"
	return *new(fasthttp.DialFunc)
}

func FasthttpHTTPDialerTimeout(proxy string, timeout time.Duration) fasthttp.DialFunc {
	_ = "STUB: not implemented"
	return *new(fasthttp.DialFunc)
}

func FasthttpHTTPDialerDualStack(proxy string) fasthttp.DialFunc {
	_ = "STUB: not implemented"
	return *new(fasthttp.DialFunc)
}

func FasthttpHTTPDialerDualStackTimeout(proxy string, timeout time.Duration) fasthttp.DialFunc {
	_ = "STUB: not implemented"
	return *new(fasthttp.DialFunc)
}
