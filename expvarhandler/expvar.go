package expvarhandler

import (
	"expvar"
	"regexp"

	"github.com/valyala/fasthttp"
)

var (
	expvarHandlerCalls = expvar.NewInt("expvarHandlerCalls")
	expvarRegexpErrors = expvar.NewInt("expvarRegexpErrors")

	defaultRE = regexp.MustCompile(".")
)

func ExpvarHandler(ctx *fasthttp.RequestCtx) { _ = "STUB: not implemented"; return }

func getExpvarRegexp(ctx *fasthttp.RequestCtx) (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
