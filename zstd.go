package fasthttp

import (
	"io"
	"sync"

	"github.com/klauspost/compress/zstd"
	"github.com/valyala/fasthttp/stackless"
)

const (
	CompressZstdSpeedNotSet = iota
	CompressZstdBestSpeed
	CompressZstdDefault
	CompressZstdSpeedBetter
	CompressZstdBestCompression
)

var (
	zstdDecoderPool            sync.Pool
	realZstdWriterPoolMap      = newCompressWriterPoolMap()
	stacklessZstdWriterPoolMap = newCompressWriterPoolMap()
)

func acquireZstdReader(r io.Reader) (*zstd.Decoder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forcetypeassert

func releaseZstdReader(zr *zstd.Decoder) { _ = "STUB: not implemented"; return }

func acquireStacklessZstdWriter(w io.Writer, compressLevel int) stackless.Writer {
	_ = "STUB: not implemented"
	return *new(stackless.Writer)
}

//nolint:forcetypeassert

func releaseStacklessZstdWriter(zf stackless.Writer, level int) { _ = "STUB: not implemented"; return }

func acquireRealZstdWriter(w io.Writer, level int) *zstd.Encoder {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

func releaseRealZstdWriter(zw *zstd.Encoder, level int) { _ = "STUB: not implemented"; return }

func AppendZstdBytesLevel(dst, src []byte, level int) []byte { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func WriteZstdLevel(w io.Writer, p []byte, level int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var (
	stacklessWriteZstdOnce sync.Once
	stacklessWriteZstdFunc func(ctx any) bool
)

func stacklessWriteZstd(ctx any) { _ = "STUB: not implemented"; return }

func nonblockingWriteZstd(ctxv any) { _ = "STUB: not implemented"; return }

//nolint:forcetypeassert

//nolint:errcheck

func AppendZstdBytes(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func WriteUnzstd(w io.Writer, p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func writeUnzstd(w io.Writer, p []byte, maxBodySize int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func estimateUnzstdSize(p []byte) int { _ = "STUB: not implemented"; return 0 }

func AppendUnzstdBytes(dst, src []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func normalizeZstdCompressLevel(level int) int { _ = "STUB: not implemented"; return 0 }
