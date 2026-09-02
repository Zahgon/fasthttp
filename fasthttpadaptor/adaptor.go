package fasthttpadaptor

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/valyala/fasthttp"
)

func NewFastHTTPHandlerFunc(h http.HandlerFunc) fasthttp.RequestHandler {
	_ = "STUB: not implemented"
	return *new(fasthttp.RequestHandler)
}

func NewFastHTTPHandler(h http.Handler) fasthttp.RequestHandler {
	_ = "STUB: not implemented"
	return *new(fasthttp.RequestHandler)
}

//nolint:forcetypeassert

var bufferPool = sync.Pool{
	New: func() any {
		b := make([]byte, 32*1024)
		return &b
	},
}

const (
	modeDone = iota + 1
	modeFlushed
	modeHijacked
	modePanicked
)

type writer struct {
	ctx        *fasthttp.RequestCtx
	h          http.Header
	statusCode atomic.Int64

	mu           sync.Mutex
	responseBody []byte
	bufPool      *[]byte

	pr *io.PipeReader
	pw *io.PipeWriter

	hijacked atomic.Bool

	modeCh chan int

	streamReady chan struct{}

	flushOnce sync.Once
	closeOnce sync.Once
}

func acquireWriter(ctx *fasthttp.RequestCtx) *writer { _ = "STUB: not implemented"; return nil }

func releaseWriter(w *writer) { _ = "STUB: not implemented"; return }

func (w *writer) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (w *writer) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func (w *writer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:forcetypeassert

func (w *writer) Flush() { _ = "STUB: not implemented"; return }

type wrappedConn struct {
	net.Conn

	wg   sync.WaitGroup
	once sync.Once
}

func (c *wrappedConn) Close() (err error) { _ = "STUB: not implemented"; return nil }

func (w *writer) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

func (w *writer) Close() error { _ = "STUB: not implemented"; return nil }

func (w *writer) status() int { _ = "STUB: not implemented"; return 0 }

func (w *writer) consumePreflush() []byte { _ = "STUB: not implemented"; return nil }
