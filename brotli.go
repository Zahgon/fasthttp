package fasthttp

import (
	"io"
	"sync"

	"github.com/andybalholm/brotli"
	"github.com/valyala/fasthttp/stackless"
)

const (
	CompressBrotliNoCompression   = 0
	CompressBrotliBestSpeed       = brotli.BestSpeed
	CompressBrotliBestCompression = brotli.BestCompression

	CompressBrotliDefaultCompression = 4
)

func acquireBrotliReader(r io.Reader) (*brotli.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forcetypeassert

func releaseBrotliReader(zr *brotli.Reader) { _ = "STUB: not implemented"; return }

var brotliReaderPool sync.Pool

func acquireStacklessBrotliWriter(w io.Writer, level int) stackless.Writer {
	_ = "STUB: not implemented"
	return *new(stackless.Writer)
}

//nolint:forcetypeassert

func releaseStacklessBrotliWriter(sw stackless.Writer, level int) {
	_ = "STUB: not implemented"
	return
}

func acquireRealBrotliWriter(w io.Writer, level int) *brotli.Writer {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

func releaseRealBrotliWriter(zw *brotli.Writer, level int) { _ = "STUB: not implemented"; return }

var (
	stacklessBrotliWriterPoolMap = newCompressWriterPoolMap()
	realBrotliWriterPoolMap      = newCompressWriterPoolMap()
)

func AppendBrotliBytesLevel(dst, src []byte, level int) []byte {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func WriteBrotliLevel(w io.Writer, p []byte, level int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var (
	stacklessWriteBrotliOnce sync.Once
	stacklessWriteBrotliFunc func(ctx any) bool
)

func stacklessWriteBrotli(ctx any) { _ = "STUB: not implemented"; return }

func nonblockingWriteBrotli(ctxv any) { _ = "STUB: not implemented"; return }

//nolint:forcetypeassert

//nolint:errcheck // no way to handle this error anyway

func WriteBrotli(w io.Writer, p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func AppendBrotliBytes(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func WriteUnbrotli(w io.Writer, p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func writeUnbrotli(w io.Writer, p []byte, maxBodySize int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func AppendUnbrotliBytes(dst, src []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeBrotliCompressLevel(level int) int { _ = "STUB: not implemented"; return 0 }
