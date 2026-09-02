package fasthttp

import (
	"bufio"
	"io"
	"sync"
)

type StreamWriter func(w *bufio.Writer)

func NewStreamReader(sw StreamWriter) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

//nolint:forcetypeassert

var streamWriterBufPool sync.Pool
