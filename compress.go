package fasthttp

import (
	"io"
	"io/fs"
	"sync"

	"github.com/klauspost/compress/flate"
	"github.com/klauspost/compress/gzip"
	"github.com/klauspost/compress/zlib"
	"github.com/valyala/fasthttp/stackless"
)

const (
	CompressNoCompression      = flate.NoCompression
	CompressBestSpeed          = flate.BestSpeed
	CompressBestCompression    = flate.BestCompression
	CompressDefaultCompression = 6
	CompressHuffmanOnly        = -2
)

func acquireGzipReader(r io.Reader) (*gzip.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forcetypeassert

func releaseGzipReader(zr *gzip.Reader) { _ = "STUB: not implemented"; return }

var gzipReaderPool sync.Pool

func acquireFlateReader(r io.Reader) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

//nolint:forcetypeassert

func releaseFlateReader(zr io.ReadCloser) { _ = "STUB: not implemented"; return }

func resetFlateReader(zr io.ReadCloser, r io.Reader) error { _ = "STUB: not implemented"; return nil }

var flateReaderPool sync.Pool

func acquireStacklessGzipWriter(w io.Writer, level int) stackless.Writer {
	_ = "STUB: not implemented"
	return *new(stackless.Writer)
}

//nolint:forcetypeassert

func releaseStacklessGzipWriter(sw stackless.Writer, level int) { _ = "STUB: not implemented"; return }

func acquireRealGzipWriter(w io.Writer, level int) *gzip.Writer {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

func releaseRealGzipWriter(zw *gzip.Writer, level int) { _ = "STUB: not implemented"; return }

var (
	stacklessGzipWriterPoolMap = newCompressWriterPoolMap()
	realGzipWriterPoolMap      = newCompressWriterPoolMap()
)

func AppendGzipBytesLevel(dst, src []byte, level int) []byte { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func WriteGzipLevel(w io.Writer, p []byte, level int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var (
	stacklessWriteGzipOnce sync.Once
	stacklessWriteGzipFunc func(ctx any) bool
)

func stacklessWriteGzip(ctx any) { _ = "STUB: not implemented"; return }

func nonblockingWriteGzip(ctxv any) { _ = "STUB: not implemented"; return }

//nolint:forcetypeassert

//nolint:errcheck // no way to handle this error anyway

func WriteGzip(w io.Writer, p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func AppendGzipBytes(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func WriteGunzip(w io.Writer, p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func writeGunzip(w io.Writer, p []byte, maxBodySize int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func AppendGunzipBytes(dst, src []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func AppendDeflateBytesLevel(dst, src []byte, level int) []byte {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func WriteDeflateLevel(w io.Writer, p []byte, level int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var (
	stacklessWriteDeflateOnce sync.Once
	stacklessWriteDeflateFunc func(ctx any) bool
)

func stacklessWriteDeflate(ctx any) { _ = "STUB: not implemented"; return }

func nonblockingWriteDeflate(ctxv any) { _ = "STUB: not implemented"; return }

//nolint:forcetypeassert

//nolint:errcheck // no way to handle this error anyway

type compressCtx struct {
	w     io.Writer
	p     []byte
	level int
}

func WriteDeflate(w io.Writer, p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func AppendDeflateBytes(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func WriteInflate(w io.Writer, p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func writeInflate(w io.Writer, p []byte, maxBodySize int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func AppendInflateBytes(dst, src []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type byteSliceWriter struct {
	b []byte
}

func (w *byteSliceWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *byteSliceWriter) WriteString(s string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type byteSliceReader struct {
	b []byte
}

func (r *byteSliceReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *byteSliceReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func acquireStacklessDeflateWriter(w io.Writer, level int) stackless.Writer {
	_ = "STUB: not implemented"
	return *new(stackless.Writer)
}

//nolint:forcetypeassert

func releaseStacklessDeflateWriter(sw stackless.Writer, level int) {
	_ = "STUB: not implemented"
	return
}

func acquireRealDeflateWriter(w io.Writer, level int) *zlib.Writer {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

func releaseRealDeflateWriter(zw *zlib.Writer, level int) { _ = "STUB: not implemented"; return }

var (
	stacklessDeflateWriterPoolMap = newCompressWriterPoolMap()
	realDeflateWriterPoolMap      = newCompressWriterPoolMap()
)

func newCompressWriterPoolMap() []*sync.Pool { _ = "STUB: not implemented"; return nil }

func isFileCompressible(f fs.File, minCompressRatio float64) bool {
	_ = "STUB: not implemented"
	return false
}

//nolint:errcheck

func normalizeCompressLevel(level int) int { _ = "STUB: not implemented"; return 0 }
