package fasthttp

import (
	"bufio"
	"bytes"
	"sync"

	"github.com/valyala/bytebufferpool"
)

type bodyStreamHeader interface {
	ContentLength() int
	ReadTrailer(r *bufio.Reader) error
}

type requestStream struct {
	header          bodyStreamHeader
	prefetchedBytes bytes.Reader
	reader          *bufio.Reader
	contentLength   int
	totalBytesRead  int
	chunkLeft       int
	strictEOF       bool
}

func (rs *requestStream) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func acquireRequestStream(b *bytebufferpool.ByteBuffer, r *bufio.Reader, h bodyStreamHeader) *requestStream {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

func acquireResponseStream(b *bytebufferpool.ByteBuffer, r *bufio.Reader, h bodyStreamHeader) *requestStream {
	_ = "STUB: not implemented"
	return nil
}

func releaseRequestStream(rs *requestStream) { _ = "STUB: not implemented"; return }

var requestStreamPool = sync.Pool{
	New: func() any {
		return &requestStream{}
	},
}
