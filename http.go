package fasthttp

import (
	"bufio"
	"errors"
	"io"
	"mime/multipart"
	"net"
	"sync"
	"time"

	"github.com/valyala/bytebufferpool"
)

var (
	requestBodyPoolSizeLimit  int64 = -1
	responseBodyPoolSizeLimit int64 = -1
)

func SetBodySizePoolLimit(reqBodyLimit, respBodyLimit int) { _ = "STUB: not implemented"; return }

type Request struct {
	noCopy noCopy

	bodyStream io.Reader
	w          requestBodyWriter
	body       *bytebufferpool.ByteBuffer

	multipartForm         *multipart.Form
	multipartFormBoundary string

	postArgs   Args
	userValues userData

	bodyRaw []byte

	uri URI

	Header RequestHeader

	timeout time.Duration

	secureErrorLogMessage bool

	parsedURI      bool
	parsedPostArgs bool
	uriParseErr    error

	keepBodyBuffer bool

	forceResponseBodyBuffering bool

	isTLS bool

	UseHostHeader bool

	DisableRedirectPathNormalizing bool
}

type Response struct {
	noCopy noCopy

	bodyStream io.Reader

	raddr net.Addr

	laddr net.Addr
	w     responseBodyWriter
	body  *bytebufferpool.ByteBuffer

	bodyRaw []byte

	Header ResponseHeader

	ImmediateHeaderFlush bool

	StreamBody bool

	SkipBody bool

	keepBodyBuffer        bool
	preserveBodyBuffer    bool
	secureErrorLogMessage bool
}

func (req *Request) SetHost(host string) { _ = "STUB: not implemented"; return }

func (req *Request) SetHostBytes(host []byte) { _ = "STUB: not implemented"; return }

func (req *Request) Host() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) SetRequestURI(requestURI string) { _ = "STUB: not implemented"; return }

func (req *Request) SetRequestURIBytes(requestURI []byte) { _ = "STUB: not implemented"; return }

func (req *Request) RequestURI() []byte { _ = "STUB: not implemented"; return nil }

func (resp *Response) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (resp *Response) SetStatusCode(statusCode int) { _ = "STUB: not implemented"; return }

func (resp *Response) ConnectionClose() bool { _ = "STUB: not implemented"; return false }

func (resp *Response) SetConnectionClose() { _ = "STUB: not implemented"; return }

func (req *Request) ConnectionClose() bool { _ = "STUB: not implemented"; return false }

func (req *Request) SetConnectionClose() { _ = "STUB: not implemented"; return }

func (req *Request) GetTimeOut() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (resp *Response) SendFile(path string) error { _ = "STUB: not implemented"; return nil }

func (req *Request) SetBodyStream(bodyStream io.Reader, bodySize int) {
	_ = "STUB: not implemented"
	return
}

func (resp *Response) SetBodyStream(bodyStream io.Reader, bodySize int) {
	_ = "STUB: not implemented"
	return
}

func (req *Request) IsBodyStream() bool { _ = "STUB: not implemented"; return false }

func (resp *Response) IsBodyStream() bool { _ = "STUB: not implemented"; return false }

func (req *Request) SetBodyStreamWriter(sw StreamWriter) { _ = "STUB: not implemented"; return }

func (resp *Response) SetBodyStreamWriter(sw StreamWriter) { _ = "STUB: not implemented"; return }

func (resp *Response) BodyWriter() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (req *Request) BodyStream() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (req *Request) CloseBodyStream() error { _ = "STUB: not implemented"; return nil }

func (resp *Response) BodyStream() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (resp *Response) CloseBodyStream() error { _ = "STUB: not implemented"; return nil }

type ReadCloserWithError interface {
	io.Reader
	CloseWithError(err error) error
}

type closeReader struct {
	io.Reader

	closeFunc func(err error) error
}

func newCloseReaderWithError(r io.Reader, closeFunc func(err error) error) ReadCloserWithError {
	_ = "STUB: not implemented"
	return *new(ReadCloserWithError)
}

func (c *closeReader) CloseWithError(err error) error { _ = "STUB: not implemented"; return nil }

func (req *Request) BodyWriter() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

type responseBodyWriter struct {
	r *Response
}

func (w *responseBodyWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *responseBodyWriter) WriteString(s string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type requestBodyWriter struct {
	r *Request
}

func (w *requestBodyWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *requestBodyWriter) WriteString(s string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (resp *Response) ParseNetConn(conn net.Conn) { _ = "STUB: not implemented"; return }

func (resp *Response) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (resp *Response) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (resp *Response) Body() []byte { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (resp *Response) bodyBytes() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) bodyBytes() []byte { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (resp *Response) bodyBuffer() *bytebufferpool.ByteBuffer {
	_ = "STUB: not implemented"
	return nil
}

func (req *Request) bodyBuffer() *bytebufferpool.ByteBuffer { _ = "STUB: not implemented"; return nil }

var (
	responseBodyPool bytebufferpool.Pool
	requestBodyPool  bytebufferpool.Pool
)

func (req *Request) BodyGunzip() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (req *Request) BodyGunzipWithLimit(maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resp *Response) BodyGunzip() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (resp *Response) BodyGunzipWithLimit(maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func gunzipData(p []byte, maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (req *Request) BodyUnbrotli() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (req *Request) BodyUnbrotliWithLimit(maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resp *Response) BodyUnbrotli() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (resp *Response) BodyUnbrotliWithLimit(maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unBrotliData(p []byte, maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (req *Request) BodyInflate() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (req *Request) BodyInflateWithLimit(maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resp *Response) BodyInflate() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (resp *Response) BodyInflateWithLimit(maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctx *RequestCtx) RequestBodyStream() io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (req *Request) BodyUnzstd() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (req *Request) BodyUnzstdWithLimit(maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resp *Response) BodyUnzstd() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (resp *Response) BodyUnzstdWithLimit(maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unzstdData(p []byte, maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inflateData(p []byte, maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var ErrContentEncodingUnsupported = errors.New("fasthttp: unsupported content-encoding")

func (req *Request) BodyUncompressed() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (req *Request) BodyUncompressedWithLimit(maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resp *Response) BodyUncompressed() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resp *Response) BodyUncompressedWithLimit(maxBodySize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (req *Request) BodyWriteTo(w io.Writer) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (resp *Response) BodyWriteTo(w io.Writer) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (resp *Response) AppendBody(p []byte) { _ = "STUB: not implemented"; return }

//nolint:errcheck
//nolint:errcheck

func (resp *Response) AppendBodyString(s string) { _ = "STUB: not implemented"; return }

//nolint:errcheck
//nolint:errcheck

func (resp *Response) SetBody(body []byte) { _ = "STUB: not implemented"; return }

//nolint:errcheck

//nolint:errcheck

func (resp *Response) SetBodyString(body string) { _ = "STUB: not implemented"; return }

//nolint:errcheck

//nolint:errcheck

func (resp *Response) ResetBody() { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (resp *Response) SetBodyRaw(body []byte) { _ = "STUB: not implemented"; return }

func (req *Request) SetBodyRaw(body []byte) { _ = "STUB: not implemented"; return }

func (resp *Response) ReleaseBody(size int) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (req *Request) ReleaseBody(size int) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (resp *Response) SwapBody(body []byte) []byte { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (req *Request) SwapBody(body []byte) []byte { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (req *Request) Body() []byte { _ = "STUB: not implemented"; return nil }

func (req *Request) AppendBody(p []byte) { _ = "STUB: not implemented"; return }

//nolint:errcheck
//nolint:errcheck

func (req *Request) AppendBodyString(s string) { _ = "STUB: not implemented"; return }

//nolint:errcheck
//nolint:errcheck

func (req *Request) SetBody(body []byte) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (req *Request) SetBodyString(body string) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (req *Request) ResetBody() { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (req *Request) CopyTo(dst *Request) { _ = "STUB: not implemented"; return }

func (req *Request) copyToSkipBody(dst *Request) { _ = "STUB: not implemented"; return }

func (resp *Response) CopyTo(dst *Response) { _ = "STUB: not implemented"; return }

func (resp *Response) copyToSkipBody(dst *Response) { _ = "STUB: not implemented"; return }

func swapRequestBody(a, b *Request) { _ = "STUB: not implemented"; return }

func swapResponseBody(a, b *Response) { _ = "STUB: not implemented"; return }

func (req *Request) URI() *URI {
	_ = "STUB: not implemented"
	//nolint:errcheck
	return nil
}

func (req *Request) SetURI(newURI *URI) { _ = "STUB: not implemented"; return }

func (req *Request) parseURI() error { _ = "STUB: not implemented"; return nil }

func (req *Request) PostArgs() *Args { _ = "STUB: not implemented"; return nil }

func (req *Request) parsePostArgs() { _ = "STUB: not implemented"; return }

var ErrNoMultipartForm = errors.New("fasthttp: request content-type has bad boundary or is not multipart/form-data")

func (req *Request) MultipartForm() (*multipart.Form, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (req *Request) MultipartFormWithLimit(maxBodySize int) (*multipart.Form, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalMultipartForm(f *multipart.Form, boundary string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WriteMultipartForm(w io.Writer, f *multipart.Form, boundary string) error {
	_ = "STUB: not implemented"
	return nil
}

func readMultipartForm(r io.Reader, boundary string, size, maxInMemoryFileSize int) (*multipart.Form, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (req *Request) Reset() { _ = "STUB: not implemented"; return }

func (req *Request) resetSkipHeader() { _ = "STUB: not implemented"; return }

func (req *Request) RemoveMultipartFormFiles() { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (resp *Response) Reset() { _ = "STUB: not implemented"; return }

func (resp *Response) resetSkipHeader() { _ = "STUB: not implemented"; return }

func (req *Request) Read(r *bufio.Reader) error { _ = "STUB: not implemented"; return nil }

const defaultMaxInMemoryFileSize = 16 * 1024 * 1024

var ErrGetOnly = errors.New("fasthttp: non-get request received")

func (req *Request) ReadLimitBody(r *bufio.Reader, maxBodySize int) error {
	_ = "STUB: not implemented"
	return nil
}

func (req *Request) readLimitBody(r *bufio.Reader, maxBodySize int, getOnly, preParseMultipartForm bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (req *Request) readBodyStream(r *bufio.Reader, maxBodySize int, getOnly, preParseMultipartForm bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (req *Request) MayContinue() bool { _ = "STUB: not implemented"; return false }

func (req *Request) ContinueReadBody(r *bufio.Reader, maxBodySize int, preParseMultipartForm ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (req *Request) ReadBody(r *bufio.Reader, contentLength, maxBodySize int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (req *Request) ContinueReadBodyStream(r *bufio.Reader, maxBodySize int, preParseMultipartForm ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (resp *Response) Read(r *bufio.Reader) error { _ = "STUB: not implemented"; return nil }

const maxInterimResponses = 100

var errTooManyInterimResponses = errors.New("fasthttp: too many 1xx informational responses received")

func (resp *Response) ReadLimitBody(r *bufio.Reader, maxBodySize int) error {
	_ = "STUB: not implemented"
	return nil
}

func (resp *Response) ReadBody(r *bufio.Reader, maxBodySize int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (resp *Response) mustSkipBody() bool { _ = "STUB: not implemented"; return false }

var errRequestHostRequired = errors.New("missing required host header in request")

func (req *Request) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (resp *Response) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func writeBufio(hw httpWriter, w io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type statsWriter struct {
	w            io.Writer
	bytesWritten int64
}

func (w *statsWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *statsWriter) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func acquireStatsWriter(w io.Writer) *statsWriter { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func releaseStatsWriter(sw *statsWriter) { _ = "STUB: not implemented"; return }

var statsWriterPool sync.Pool

func acquireBufioWriter(w io.Writer) *bufio.Writer { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func releaseBufioWriter(bw *bufio.Writer) { _ = "STUB: not implemented"; return }

var bufioWriterPool sync.Pool

func (req *Request) onlyMultipartForm() bool { _ = "STUB: not implemented"; return false }

func (req *Request) Write(w *bufio.Writer) error { _ = "STUB: not implemented"; return nil }

func (resp *Response) WriteGzip(w *bufio.Writer) error { _ = "STUB: not implemented"; return nil }

func (resp *Response) WriteGzipLevel(w *bufio.Writer, level int) error {
	_ = "STUB: not implemented"
	return nil
}

func (resp *Response) WriteDeflate(w *bufio.Writer) error { _ = "STUB: not implemented"; return nil }

func (resp *Response) WriteDeflateLevel(w *bufio.Writer, level int) error {
	_ = "STUB: not implemented"
	return nil
}

func (resp *Response) brotliBody(level int) { _ = "STUB: not implemented"; return }

func (resp *Response) gzipBody(level int) { _ = "STUB: not implemented"; return }

func (resp *Response) deflateBody(level int) { _ = "STUB: not implemented"; return }

func (resp *Response) zstdBody(level int) { _ = "STUB: not implemented"; return }

type compressedBodyStream struct {
	io.ReadCloser

	bodyStream io.Reader
	level      int
	compress   compressBodyStream

	done chan struct{}

	closeReadOnce sync.Once
	closeReadErr  error

	originalLock   sync.Mutex
	originalClosed bool
	closeErr       error
}

func (s *compressedBodyStream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *compressedBodyStream) write(sw *bufio.Writer) { _ = "STUB: not implemented"; return }

func (s *compressedBodyStream) closeOriginal(wErr error) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *compressedBodyStream) closeOriginalForDiscard() error {
	_ = "STUB: not implemented"
	return nil
}

type compressBodyStream func(sw *bufio.Writer, bodyStream io.Reader, level int) error

func newCompressedBodyStream(bodyStream io.Reader, level int, compress compressBodyStream) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func compressBrotliBodyStream(sw *bufio.Writer, bodyStream io.Reader, level int) error {
	_ = "STUB: not implemented"
	return nil
}

func compressGzipBodyStream(sw *bufio.Writer, bodyStream io.Reader, level int) error {
	_ = "STUB: not implemented"
	return nil
}

func compressDeflateBodyStream(sw *bufio.Writer, bodyStream io.Reader, level int) error {
	_ = "STUB: not implemented"
	return nil
}

func compressZstdBodyStream(sw *bufio.Writer, bodyStream io.Reader, level int) error {
	_ = "STUB: not implemented"
	return nil
}

func closeBodyStreamReader(bodyStream io.Reader, wErr error) error {
	_ = "STUB: not implemented"
	return nil
}

const minCompressLen = 200

type writeFlusher interface {
	io.Writer
	Flush() error
}

type flushWriter struct {
	wf writeFlusher
	bw *bufio.Writer
}

func (w *flushWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *flushWriter) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (resp *Response) Write(w *bufio.Writer) error { _ = "STUB: not implemented"; return nil }

func (req *Request) writeBodyStream(w *bufio.Writer) error { _ = "STUB: not implemented"; return nil }

type ErrBodyStreamWritePanic struct {
	error
}

func (resp *Response) writeBodyStream(w *bufio.Writer, sendBody bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (req *Request) closeBodyStream() error { _ = "STUB: not implemented"; return nil }

func (resp *Response) closeBodyStream(wErr error) error { _ = "STUB: not implemented"; return nil }

func (req *Request) String() string { _ = "STUB: not implemented"; return "" }

func (resp *Response) String() string { _ = "STUB: not implemented"; return "" }

func (req *Request) SetUserValue(key, value any) { _ = "STUB: not implemented"; return }

func (req *Request) SetUserValueBytes(key []byte, value any) { _ = "STUB: not implemented"; return }

func (req *Request) UserValue(key any) any { _ = "STUB: not implemented"; return *new(any) }

func (req *Request) UserValueBytes(key []byte) any { _ = "STUB: not implemented"; return *new(any) }

func (req *Request) VisitUserValues(visitor func([]byte, any)) { _ = "STUB: not implemented"; return }

func (req *Request) VisitUserValuesAll(visitor func(any, any)) { _ = "STUB: not implemented"; return }

func (req *Request) ResetUserValues() { _ = "STUB: not implemented"; return }

func (req *Request) RemoveUserValue(key any) { _ = "STUB: not implemented"; return }

func (req *Request) RemoveUserValueBytes(key []byte) { _ = "STUB: not implemented"; return }

func getHTTPString(hw httpWriter) string { _ = "STUB: not implemented"; return "" }

type httpWriter interface {
	Write(w *bufio.Writer) error
}

type BodyWriterTo interface {
	io.WriterTo
	SupportsBodyWriteTo() bool
}

type chunkedBodyWriter struct {
	w   *bufio.Writer
	err error
}

func (cw *chunkedBodyWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func writeBodyChunked(w *bufio.Writer, r io.Reader) error { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func limitedReaderSize(r io.Reader) int64 { _ = "STUB: not implemented"; return 0 }

func writeBodyFixedSize(w *bufio.Writer, r io.Reader, size int64) error {
	_ = "STUB: not implemented"
	return nil
}

func copyBodyStream(w io.Writer, r io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:forcetypeassert

//nolint:dupword
func copyZeroAlloc(w io.Writer, r io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:forcetypeassert

func copyBuffer(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var copyBufPool = sync.Pool{
	New: func() any {
		return make([]byte, 4096)
	},
}

func writeChunk(w *bufio.Writer, b []byte) error { _ = "STUB: not implemented"; return nil }

var ErrBodyTooLarge = errors.New("fasthttp: body size exceeds the given limit")

func copyZeroAllocWithLimit(w io.Writer, r io.Reader, maxBodySize int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func readBody(r *bufio.Reader, contentLength, maxBodySize int, dst []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var errChunkedStream = errors.New("chunked stream")

func readBodyWithStreaming(r *bufio.Reader, contentLength, maxBodySize int, dst []byte) (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readBodyIdentity(r *bufio.Reader, maxBodySize int, dst []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendBodyFixedSize(r *bufio.Reader, dst []byte, n int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ErrBrokenChunk struct {
	error
}

func readBodyChunked(r *bufio.Reader, maxBodySize int, dst []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseChunkSize(r *bufio.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func readCrLf(r *bufio.Reader) error { _ = "STUB: not implemented"; return nil }

func (req *Request) SetTimeout(t time.Duration) { _ = "STUB: not implemented"; return }
