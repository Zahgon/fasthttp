package fasthttp

import (
	"bufio"
	"errors"
	"io"
	"iter"
	"sync"
	"sync/atomic"
	"time"
)

const (
	rChar = byte('\r')
	nChar = byte('\n')
)

type header struct {
	h       []argsKV
	cookies []argsKV

	bufK               []byte
	bufV               []byte
	contentLengthBytes []byte
	contentType        []byte
	protocol           []byte

	mulHeader [][]byte
	trailer   [][]byte

	contentLength int

	disableNormalizing    bool
	secureErrorLogMessage bool
	noHTTP11              bool
	connectionClose       bool
	noDefaultContentType  bool
}

type ResponseHeader struct {
	header

	noCopy noCopy

	statusMessage   []byte
	contentEncoding []byte
	server          []byte

	statusCode int

	noDefaultDate bool
}

type RequestHeader struct {
	header

	noCopy noCopy

	method     []byte
	requestURI []byte
	host       []byte
	userAgent  []byte

	rawHeaders []byte

	disableSpecialHeader bool
	cookiesCollected     bool
}

func (h *ResponseHeader) SetContentRange(startPos, endPos, contentLength int) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) SetByteRange(startPos, endPos int) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (h *ResponseHeader) SetStatusCode(statusCode int) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) StatusMessage() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) SetStatusMessage(statusMessage []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetProtocol(protocol []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetLastModified(t time.Time) { _ = "STUB: not implemented"; return }

func (h *header) ConnectionClose() bool { _ = "STUB: not implemented"; return false }

func (h *header) SetConnectionClose() { _ = "STUB: not implemented"; return }

func (h *header) ResetConnectionClose() { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) ConnectionUpgrade() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) ConnectionUpgrade() bool { _ = "STUB: not implemented"; return false }

func (h *ResponseHeader) PeekCookie(key string) []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) ContentLength() int { _ = "STUB: not implemented"; return 0 }

func (h *RequestHeader) ContentLength() int { _ = "STUB: not implemented"; return 0 }

func (h *ResponseHeader) SetContentLength(contentLength int) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) mustSkipContentLength() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) SetContentLength(contentLength int) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) isCompressibleContentType() bool { _ = "STUB: not implemented"; return false }

func (h *ResponseHeader) ContentType() []byte { _ = "STUB: not implemented"; return nil }

func (h *header) SetContentType(contentType string) { _ = "STUB: not implemented"; return }

func (h *header) SetContentTypeBytes(contentType []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) ContentEncoding() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) SetContentEncoding(contentEncoding string) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) SetContentEncodingBytes(contentEncoding []byte) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) addVaryBytes(value []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) Server() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) SetServer(server string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetServerBytes(server []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) ContentType() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) ContentEncoding() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) SetContentEncoding(contentEncoding string) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) SetContentEncodingBytes(contentEncoding []byte) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) SetMultipartFormBoundary(boundary string) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) SetMultipartFormBoundaryBytes(boundary []byte) {
	_ = "STUB: not implemented"
	return
}

func (h *header) SetTrailer(trailer string) error { _ = "STUB: not implemented"; return nil }

func (h *header) SetTrailerBytes(trailer []byte) error { _ = "STUB: not implemented"; return nil }

func (h *header) AddTrailer(trailer string) error { _ = "STUB: not implemented"; return nil }

var (
	ErrBadTrailer                    = errors.New("fasthttp: contain forbidden trailer")
	ErrReadingResponseHeaders        = errors.New("fasthttp: error when reading response headers")
	ErrReadingResponseTrailer        = errors.New("fasthttp: error when reading response trailer")
	ErrResponseFirstLineMissingSpace = errors.New("fasthttp: cannot find whitespace in the first line of response")
	ErrUnexpectedStatusCodeChar      = errors.New("fasthttp: unexpected char at the end of status code")
	ErrMissingRequestMethod          = errors.New("fasthttp: cannot find http request method")
	ErrUnsupportedRequestMethod      = errors.New("fasthttp: unsupported http request method")
	ErrExtraWhitespaceInRequestLine  = errors.New("fasthttp: extra whitespace in request line")
	ErrEmptyRequestURI               = errors.New("fasthttp: requesturi cannot be empty")
	ErrDuplicateContentLength        = errors.New("fasthttp: duplicate content-length header")
	ErrUnsupportedTransferEncoding   = errors.New("fasthttp: unsupported transfer-encoding")
	ErrNonNumericChars               = errors.New("fasthttp: non-numeric chars found")
	ErrNeedMore                      = errors.New("fasthttp: need more data: cannot find trailing lf")
	ErrSmallReadBuffer               = errors.New("fasthttp: small read buffer. increase readbuffersize")
)

func (h *header) AddTrailerBytes(trailer []byte) (err error) { _ = "STUB: not implemented"; return nil }

func isValidTrailerKey(key []byte) bool { _ = "STUB: not implemented"; return false }

func validHeaderFieldByte(c byte) bool { _ = "STUB: not implemented"; return false }

func validHeaderValueByte(c byte) bool { _ = "STUB: not implemented"; return false }

func isValidHeaderKey(a []byte) (valid, innerSpace bool) {
	_ = "STUB: not implemented"
	return false, false
}

func VisitHeaderParams(b []byte, f func(key, value []byte) bool) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) MultipartFormBoundary() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) Host() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) SetHost(host string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetHostBytes(host []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) UserAgent() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) SetUserAgent(userAgent string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetUserAgentBytes(userAgent []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) Referer() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) SetReferer(referer string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetRefererBytes(referer []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) Method() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) SetMethod(method string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetMethodBytes(method []byte) { _ = "STUB: not implemented"; return }

func (h *header) Protocol() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) SetProtocol(protocol string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetProtocolBytes(protocol []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) RequestURI() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) SetRequestURI(requestURI string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetRequestURIBytes(requestURI []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) IsGet() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsPost() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsPut() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsHead() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsDelete() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsConnect() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsOptions() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsTrace() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsPatch() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsQuery() bool { _ = "STUB: not implemented"; return false }

func (h *header) IsHTTP11() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) HasAcceptEncoding(acceptEncoding string) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *RequestHeader) HasAcceptEncodingBytes(acceptEncoding []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *ResponseHeader) Len() int { _ = "STUB: not implemented"; return 0 }

func (h *RequestHeader) Len() int { _ = "STUB: not implemented"; return 0 }

func (h *RequestHeader) DisableSpecialHeader() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) EnableSpecialHeader() bool { _ = "STUB: not implemented"; return false }

func (h *header) DisableNormalizing() bool { _ = "STUB: not implemented"; return false }

func (h *header) EnableNormalizing() bool { _ = "STUB: not implemented"; return false }

func (h *header) SetNoDefaultContentType(noDefaultContentType bool) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) Reset() { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) resetSkipNormalize() { _ = "STUB: not implemented"; return }

func (h *RequestHeader) Reset() { _ = "STUB: not implemented"; return }

func (h *RequestHeader) resetSkipNormalize() { _ = "STUB: not implemented"; return }

func (h *header) copyTo(dst *header) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) CopyTo(dst *ResponseHeader) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) CopyTo(dst *RequestHeader) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) All() iter.Seq2[[]byte, []byte] { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) VisitAll(f func(key, value []byte)) { _ = "STUB: not implemented"; return }

func (h *header) Trailers() iter.Seq[[]byte] { _ = "STUB: not implemented"; return nil }

func (h *header) VisitAllTrailer(f func(value []byte)) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) Cookies() iter.Seq2[[]byte, []byte] { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) VisitAllCookie(f func(key, value []byte)) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) Cookies() iter.Seq2[[]byte, []byte] { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) VisitAllCookie(f func(key, value []byte)) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) All() iter.Seq2[[]byte, []byte] { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) VisitAll(f func(key, value []byte)) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) AllInOrder() iter.Seq2[[]byte, []byte] {
	_ = "STUB: not implemented"
	return nil
}

func (h *RequestHeader) VisitAllInOrder(f func(key, value []byte)) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) Del(key string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) DelBytes(key []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) del(key []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) Del(key string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) DelBytes(key []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) del(key []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) setSpecialHeader(key, value []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *header) setNonSpecial(key, value []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) setSpecialHeader(key, value []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *ResponseHeader) Add(key, value string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) AddBytesK(key []byte, value string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) AddBytesV(key string, value []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) AddBytesKV(key, value []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) Set(key, value string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetBytesK(key []byte, value string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetBytesV(key string, value []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetBytesKV(key, value []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetCanonical(key, value []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetCookie(cookie *Cookie) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetCookie(key, value string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetCookieBytesK(key []byte, value string) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) SetCookieBytesKV(key, value []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) DelClientCookie(key string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) DelClientCookieBytes(key []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) DelCookie(key string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) DelCookieBytes(key []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) DelCookie(key string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) DelCookieBytes(key []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) DelAllCookies() { _ = "STUB: not implemented"; return }

func (h *RequestHeader) DelAllCookies() { _ = "STUB: not implemented"; return }

func (h *RequestHeader) Add(key, value string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) AddBytesK(key []byte, value string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) AddBytesV(key string, value []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) AddBytesKV(key, value []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) Set(key, value string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetBytesK(key []byte, value string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetBytesV(key string, value []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetBytesKV(key, value []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetCanonical(key, value []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) Peek(key string) []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) PeekBytes(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) PeekCanonical(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) Peek(key string) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) PeekBytes(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) PeekCanonical(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) peek(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) peek(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) PeekAll(key string) [][]byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) peekAll(key []byte) [][]byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) PeekAll(key string) [][]byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) peekAll(key []byte) [][]byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) PeekKeys() [][]byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) PeekKeys() [][]byte { _ = "STUB: not implemented"; return nil }

func (h *header) PeekTrailerKeys() [][]byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) Cookie(key string) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) CookieBytes(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) Cookie(cookie *Cookie) bool { _ = "STUB: not implemented"; return false }

//nolint:errcheck

func (h *ResponseHeader) Read(r *bufio.Reader) error { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) tryRead(r *bufio.Reader, n int) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *header) ReadTrailer(r *bufio.Reader) error { _ = "STUB: not implemented"; return nil }

func (h *header) tryReadTrailer(r *bufio.Reader, n int) error {
	_ = "STUB: not implemented"
	return nil
}

func headerError(typ string, err, errParse error, b []byte, secureErrorLogMessage bool) error {
	_ = "STUB: not implemented"
	return nil
}

func headerErrorMsg(typ string, err error, b []byte, secureErrorLogMessage bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *RequestHeader) Read(r *bufio.Reader) error { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) readLoop(r *bufio.Reader, waitForMore bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *RequestHeader) tryRead(r *bufio.Reader, n int) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *RequestHeader) validate() error { _ = "STUB: not implemented"; return nil }

func bufferSnippet(b []byte) string { _ = "STUB: not implemented"; return "" }

func isOnlyCRLF(b []byte) bool { _ = "STUB: not implemented"; return false }

func updateServerDate() { _ = "STUB: not implemented"; return }

var (
	serverDate     atomic.Pointer[[]byte]
	serverDateOnce sync.Once
)

func refreshServerDate() { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) Write(w *bufio.Writer) error { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) WriteTo(w io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *ResponseHeader) Header() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) writeTrailer(w *bufio.Writer) error { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) TrailerHeader() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) String() string { _ = "STUB: not implemented"; return "" }

func (h *ResponseHeader) appendStatusLine(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) AppendBytes(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) Write(w *bufio.Writer) error { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) WriteTo(w io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *RequestHeader) Header() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) writeTrailer(w *bufio.Writer) error { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) TrailerHeader() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) RawHeaders() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) String() string { _ = "STUB: not implemented"; return "" }

func (h *RequestHeader) AppendBytes(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func appendHeaderLine(dst, key, value []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) parse(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (h *RequestHeader) ignoreBody() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) parse(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func parseTrailer(src []byte, dest []argsKV, disableNormalizing bool) ([]argsKV, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func isBadTrailer(key []byte) bool { _ = "STUB: not implemented"; return false }

func isHTTPVersion(proto []byte) bool { _ = "STUB: not implemented"; return false }

func (h *ResponseHeader) parseFirstLine(buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func isValidMethod(method []byte) bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) parseFirstLine(buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func validateRequestURI(method, requestURI []byte) error { _ = "STUB: not implemented"; return nil }

func readRawHeaders(dst, buf []byte) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (h *ResponseHeader) parseHeaders(buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *RequestHeader) parseHeaders(buf []byte, blockEnd int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *RequestHeader) collectCookies() { _ = "STUB: not implemented"; return }

func parseContentLength(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type headerValueScanner struct {
	b     []byte
	value []byte
}

func (s *headerValueScanner) next() bool { _ = "STUB: not implemented"; return false }

func stripSpace(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func hasHeaderValue(s, value []byte) bool { _ = "STUB: not implemented"; return false }

func nextLine(b []byte) ([]byte, []byte, error) { _ = "STUB: not implemented"; return nil, nil, nil }

func initHeaderKV(bufK, bufV []byte, key, value string, disableNormalizing bool) ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initHeaderValueString(bufV []byte, value string) []byte { _ = "STUB: not implemented"; return nil }

func initHeaderValueBytes(bufV, value []byte) []byte { _ = "STUB: not implemented"; return nil }

func getHeaderKeyBytes(bufK []byte, key string, disableNormalizing bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

func normalizeHeaderKey(b []byte, disableNormalizing bool) { _ = "STUB: not implemented"; return }

func normalizeHeaderKeyValidated(b []byte, disableNormalizing bool) {
	_ = "STUB: not implemented"
	return
}

func removeNewLines(raw []byte) []byte { _ = "STUB: not implemented"; return nil }

func AppendNormalizedHeaderKey(dst []byte, key string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func AppendNormalizedHeaderKeyBytes(dst, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func appendTrailerBytes(dst []byte, trailer [][]byte, sep []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func copyTrailer(dst, src [][]byte) [][]byte { _ = "STUB: not implemented"; return nil }

type ErrNothingRead struct {
	error
}

type ErrSmallBuffer struct {
	error
}

func mustPeekBuffered(r *bufio.Reader) []byte { _ = "STUB: not implemented"; return nil }

func mustDiscard(r *bufio.Reader, n int) { _ = "STUB: not implemented"; return }
