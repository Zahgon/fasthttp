package pprofhandler

import (
	"net/http/pprof"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var (
	cmdline = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Cmdline)
	profile = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile)
	symbol  = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Symbol)
	trace   = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Trace)
	index   = fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index)
)

var (
	cmdlinePath = []byte("/debug/pprof/cmdline")
	profilePath = []byte("/debug/pprof/profile")
	symbolPath  = []byte("/debug/pprof/symbol")
	tracePath   = []byte("/debug/pprof/trace")
)

func matchPprofPath(path, endpoint []byte) bool { _ = "STUB: not implemented"; return false }

func PprofHandler(ctx *fasthttp.RequestCtx) { _ = "STUB: not implemented"; return }
