package fasthttp

import (
	"bufio"
	"context"
	"crypto/tls"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var errNoCertOrKeyProvided = errors.New("cert or key has not provided")

var ErrAlreadyServing = errors.New("fasthttp: server is already serving connections")

func ServeConn(c net.Conn, handler RequestHandler) error { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

var serverPool sync.Pool

func Serve(ln net.Listener, handler RequestHandler) error { _ = "STUB: not implemented"; return nil }

func ServeTLS(ln net.Listener, certFile, keyFile string, handler RequestHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func ServeTLSEmbed(ln net.Listener, certData, keyData []byte, handler RequestHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func ListenAndServe(addr string, handler RequestHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func ListenAndServeUNIX(addr string, mode os.FileMode, handler RequestHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func ListenAndServeTLS(addr, certFile, keyFile string, handler RequestHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func ListenAndServeTLSEmbed(addr string, certData, keyData []byte, handler RequestHandler) error {
	_ = "STUB: not implemented"
	return nil
}

type RequestHandler func(ctx *RequestCtx)

type ServeHandler func(c net.Conn) error

type Server struct {
	noCopy noCopy

	perIPConnCounter perIPConnCounter

	ctxPool        sync.Pool
	readerPool     sync.Pool
	writerPool     sync.Pool
	hijackConnPool sync.Pool

	Logger Logger

	Handler RequestHandler

	ErrorHandler func(ctx *RequestCtx, err error)

	HeaderReceived func(header *RequestHeader) RequestConfig

	ContinueHandler func(header *RequestHeader) bool

	ExpectHandler func(ctx *RequestCtx) int

	ConnState func(net.Conn, ConnState)

	TLSConfig *tls.Config

	FormValueFunc FormValueFunc

	nextProtos map[string]ServeHandler

	concurrencyCh chan struct{}

	idleConns map[net.Conn]*atomic.Int64
	done      chan struct{}

	Name string

	ln []net.Listener

	Concurrency int

	ReadBufferSize int

	WriteBufferSize int

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	IdleTimeout time.Duration

	MaxConnsPerIP int

	MaxRequestsPerConn int

	MaxKeepaliveDuration time.Duration

	MaxIdleWorkerDuration time.Duration

	TCPKeepalivePeriod time.Duration

	MaxRequestBodySize int

	SleepWhenConcurrencyLimitsExceeded time.Duration

	idleConnsMu sync.Mutex

	mu sync.Mutex

	concurrency atomic.Uint32
	open        atomic.Int32
	stop        atomic.Int32

	rejectedRequestsCount atomic.Uint32

	DisableKeepalive bool

	TCPKeepalive bool

	ReduceMemoryUsage bool

	GetOnly bool

	DisablePreParseMultipartForm bool

	LogAllErrors bool

	SecureErrorLogMessage bool

	DisableHeaderNamesNormalizing bool

	NoDefaultServerHeader bool

	NoDefaultDate bool

	NoDefaultContentType bool

	KeepHijackedConns bool

	CloseOnShutdown bool

	StreamRequestBody bool
}

func TimeoutHandler(h RequestHandler, timeout time.Duration, msg string) RequestHandler {
	_ = "STUB: not implemented"
	return *new(RequestHandler)
}

func TimeoutWithCodeHandler(h RequestHandler, timeout time.Duration, msg string, statusCode int) RequestHandler {
	_ = "STUB: not implemented"
	return *new(RequestHandler)
}

type RequestConfig struct {
	ReadTimeout time.Duration

	WriteTimeout time.Duration

	MaxRequestBodySize int
}

func CompressHandler(h RequestHandler) RequestHandler {
	_ = "STUB: not implemented"
	return *new(RequestHandler)
}

func CompressHandlerLevel(h RequestHandler, level int) RequestHandler {
	_ = "STUB: not implemented"
	return *new(RequestHandler)
}

func CompressHandlerBrotliLevel(h RequestHandler, brotliLevel, otherLevel int) RequestHandler {
	_ = "STUB: not implemented"
	return *new(RequestHandler)
}

type RequestCtx struct {
	noCopy noCopy

	Response Response

	connTime time.Time

	time time.Time

	logger     ctxLogger
	remoteAddr net.Addr

	c net.Conn
	s *Server

	timeoutResponse *Response
	timeoutCh       chan struct{}
	timeoutTimer    *time.Timer

	hijackHandler HijackHandler
	formValueFunc FormValueFunc
	fbr           firstByteReader

	Request Request

	connID           uint64
	connRequestNum   uint64
	hijackNoResponse bool
}

func (ctx *RequestCtx) EarlyHints() error { _ = "STUB: not implemented"; return nil }

type HijackHandler func(c net.Conn)

func (ctx *RequestCtx) Hijack(handler HijackHandler) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) HijackSetNoResponse(noResponse bool) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) Hijacked() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) SetUserValue(key, value any) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) SetUserValueBytes(key []byte, value any) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) UserValue(key any) any { _ = "STUB: not implemented"; return *new(any) }

func (ctx *RequestCtx) UserValueBytes(key []byte) any { _ = "STUB: not implemented"; return *new(any) }

func (ctx *RequestCtx) VisitUserValues(visitor func([]byte, any)) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestCtx) VisitUserValuesAll(visitor func(any, any)) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestCtx) ResetUserValues() { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) RemoveUserValue(key any) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) RemoveUserValueBytes(key []byte) { _ = "STUB: not implemented"; return }

type tlsConn interface {
	Handshake() error
	ConnectionState() tls.ConnectionState
}

func (ctx *RequestCtx) IsTLS() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) TLSConnectionState() *tls.ConnectionState {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestCtx) Conn() net.Conn { _ = "STUB: not implemented"; return *new(net.Conn) }

func (ctx *RequestCtx) reset() { _ = "STUB: not implemented"; return }

type firstByteReader struct {
	c        net.Conn
	ch       byte
	byteRead bool
}

func (r *firstByteReader) reset() { _ = "STUB: not implemented"; return }

func (r *firstByteReader) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type Logger interface {
	Printf(format string, args ...any)
}

var ctxLoggerLock sync.Mutex

type ctxLogger struct {
	ctx    *RequestCtx
	logger Logger
}

func (cl *ctxLogger) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

var zeroTCPAddr = &net.TCPAddr{
	IP: net.IPv4zero,
}

func (ctx *RequestCtx) String() string { _ = "STUB: not implemented"; return "" }

func (ctx *RequestCtx) ID() uint64 { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestCtx) ConnID() uint64 { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestCtx) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (ctx *RequestCtx) ConnTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (ctx *RequestCtx) ConnRequestNum() uint64 { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestCtx) SetConnectionClose() { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) SetStatusCode(statusCode int) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) SetContentType(contentType string) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) SetContentTypeBytes(contentType []byte) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) RequestURI() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) URI() *URI { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) Referer() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) UserAgent() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) Path() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) Host() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) QueryArgs() *Args { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) PostArgs() *Args { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) MultipartForm() (*multipart.Form, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctx *RequestCtx) MultipartFormWithLimit(maxBodySize int) (*multipart.Form, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctx *RequestCtx) FormFile(key string) (*multipart.FileHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var ErrMissingFile = errors.New("fasthttp: there is no uploaded file associated with the given key")

func SaveMultipartFile(fh *multipart.FileHeader, path string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestCtx) FormValue(key string) []byte { _ = "STUB: not implemented"; return nil }

type FormValueFunc func(*RequestCtx, string) []byte

var (
	defaultFormValue = func(ctx *RequestCtx, key string) []byte {
		v := ctx.QueryArgs().Peek(key)
		if len(v) > 0 {
			return v
		}
		v = ctx.PostArgs().Peek(key)
		if len(v) > 0 {
			return v
		}
		mf, err := ctx.MultipartForm()
		if err == nil && mf.Value != nil {
			vv := mf.Value[key]
			if len(vv) > 0 {
				return []byte(vv[0])
			}
		}
		return nil
	}

	//nolint:staticcheck // backwards compatibility
	NetHttpFormValueFunc = func(ctx *RequestCtx, key string) []byte {
		v := ctx.PostArgs().Peek(key)
		if len(v) > 0 {
			return v
		}
		mf, err := ctx.MultipartForm()
		if err == nil && mf.Value != nil {
			vv := mf.Value[key]
			if len(vv) > 0 {
				return []byte(vv[0])
			}
		}
		v = ctx.QueryArgs().Peek(key)
		if len(v) > 0 {
			return v
		}
		return nil
	}
)

func (ctx *RequestCtx) IsGet() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) IsPost() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) IsPut() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) IsDelete() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) IsConnect() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) IsOptions() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) IsTrace() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) IsPatch() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) Method() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) IsHead() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (ctx *RequestCtx) SetRemoteAddr(remoteAddr net.Addr) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (ctx *RequestCtx) RemoteIP() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func (ctx *RequestCtx) LocalIP() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func addrToIP(addr net.Addr) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func (ctx *RequestCtx) Error(msg string, statusCode int) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) Success(contentType string, body []byte) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) SuccessString(contentType, body string) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) Redirect(uri string, statusCode int) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) RedirectBytes(uri []byte, statusCode int) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) redirect(uri []byte, statusCode int) { _ = "STUB: not implemented"; return }

func getRedirectStatusCode(statusCode int) int { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestCtx) SetBody(body []byte) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) SetBodyString(body string) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) ResetBody() { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) SendFile(path string) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) SendFileLiteral(path string) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) SendFileBytes(path []byte) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) IfModifiedSince(lastModified time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (ctx *RequestCtx) NotModified() { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) NotFound() { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (ctx *RequestCtx) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (ctx *RequestCtx) PostBody() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) SetBodyStream(bodyStream io.Reader, bodySize int) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestCtx) SetBodyStreamWriter(sw StreamWriter) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) IsBodyStream() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestCtx) Logger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (ctx *RequestCtx) TimeoutError(msg string) { _ = "STUB: not implemented"; return }

func (ctx *RequestCtx) TimeoutErrorWithCode(msg string, statusCode int) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestCtx) TimeoutErrorWithResponse(resp *Response) { _ = "STUB: not implemented"; return }

func (s *Server) NextProto(key string, nph ServeHandler) { _ = "STUB: not implemented"; return }

func (s *Server) getNextProto(c net.Conn) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *Server) ListenAndServe(addr string) error { _ = "STUB: not implemented"; return nil }

func (s *Server) ListenAndServeUNIX(addr string, mode os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) ListenAndServeTLS(addr, certFile, keyFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) ListenAndServeTLSEmbed(addr string, certData, keyData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) ServeTLS(ln net.Listener, certFile, keyFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) ServeTLSEmbed(ln net.Listener, certData, keyData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) AppendCert(certFile, keyFile string) error { _ = "STUB: not implemented"; return nil }

func loadX509KeyPair(certFile, keyFile string) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

func (s *Server) AppendCertEmbed(certData, keyData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func x509KeyPair(certData, keyData []byte) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

func (s *Server) appendCertLocked(cert *tls.Certificate) { _ = "STUB: not implemented"; return }

func (s *Server) configTLS() { _ = "STUB: not implemented"; return }

const DefaultConcurrency = 256 * 1024

func (s *Server) Serve(ln net.Listener) error { _ = "STUB: not implemented"; return nil }

func (s *Server) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (s *Server) ShutdownWithContext(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type keepAliveConn interface {
	SetKeepAlive(keepalive bool) error
	SetKeepAlivePeriod(d time.Duration) error
	io.Closer
}

func acceptConn(s *Server, ln net.Listener, lastPerIPErrorTime *time.Time) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func wrapPerIPConn(s *Server, c net.Conn) net.Conn {
	_ = "STUB: not implemented"
	return *new(net.Conn)
}

var defaultLogger = Logger(log.New(os.Stderr, "", log.LstdFlags))

func (s *Server) logger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

var (
	ErrPerIPConnLimit = errors.New("fasthttp: too many connections per ip")

	ErrConcurrencyLimit = errors.New("fasthttp: cannot serve the connection because server.concurrency " +
		"concurrent connections are served")
)

func (s *Server) ServeConn(c net.Conn) error { _ = "STUB: not implemented"; return nil }

func (s *Server) tryAcquireConcurrency() bool { _ = "STUB: not implemented"; return false }

func (s *Server) releaseConcurrency() { _ = "STUB: not implemented"; return }

var errHijacked = errors.New("connection has been hijacked")

func (s *Server) GetCurrentConcurrency() uint32 { _ = "STUB: not implemented"; return 0 }

func (s *Server) GetOpenConnectionsCount() int32 { _ = "STUB: not implemented"; return 0 }

func (s *Server) GetRejectedConnectionsCount() uint32 { _ = "STUB: not implemented"; return 0 }

func (s *Server) getConcurrency() int { _ = "STUB: not implemented"; return 0 }

var globalConnID uint64

func nextConnID() uint64 { _ = "STUB: not implemented"; return 0 }

const DefaultMaxRequestBodySize = 4 * 1024 * 1024

func (s *Server) idleTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (s *Server) serveConnCleanup(countConcurrency bool) { _ = "STUB: not implemented"; return }

func (s *Server) serveConn(c net.Conn) error { _ = "STUB: not implemented"; return nil }

func (s *Server) serveConnCounted(c net.Conn, countConcurrency bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

func (s *Server) setState(nc net.Conn, state ConnState) { _ = "STUB: not implemented"; return }

func hijackConnHandler(ctx *RequestCtx, r io.Reader, c net.Conn, s *Server, h HijackHandler) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) acquireHijackConn(r io.Reader, c net.Conn) *hijackConn {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

func (s *Server) releaseHijackConn(hjc *hijackConn) { _ = "STUB: not implemented"; return }

type hijackConn struct {
	net.Conn

	r io.Reader
	s *Server
}

func (c *hijackConn) UnsafeConn() net.Conn { _ = "STUB: not implemented"; return *new(net.Conn) }

func (c *hijackConn) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *hijackConn) Close() error { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) LastTimeoutErrorResponse() *Response { _ = "STUB: not implemented"; return nil }

func writeResponse(ctx *RequestCtx, w *bufio.Writer) error { _ = "STUB: not implemented"; return nil }

const (
	defaultReadBufferSize  = 4096
	defaultWriteBufferSize = 4096
)

func acquireByteReader(ctxP **RequestCtx) (*bufio.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:wastedassign // Make GC happy, so it could garbage collect ctx while we wait for the

func acquireReader(ctx *RequestCtx) *bufio.Reader { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func releaseReader(s *Server, r *bufio.Reader) { _ = "STUB: not implemented"; return }

func acquireWriter(ctx *RequestCtx) *bufio.Writer { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func releaseWriter(s *Server, w *bufio.Writer) { _ = "STUB: not implemented"; return }

func (s *Server) acquireCtx(c net.Conn) (ctx *RequestCtx) { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func (ctx *RequestCtx) Init2(conn net.Conn, logger Logger, reduceMemoryUsage bool) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestCtx) Init(req *Request, remoteAddr net.Addr, logger Logger) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestCtx) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (ctx *RequestCtx) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) Err() error { _ = "STUB: not implemented"; return nil }

func (ctx *RequestCtx) Value(key any) any { _ = "STUB: not implemented"; return *new(any) }

var fakeServer = &Server{
	done: make(chan struct{}),

	concurrencyCh: make(chan struct{}, DefaultConcurrency),
}

type fakeAddrer struct {
	net.Conn

	laddr net.Addr
	raddr net.Addr
}

func (fa *fakeAddrer) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (fa *fakeAddrer) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (fa *fakeAddrer) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (fa *fakeAddrer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (fa *fakeAddrer) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Server) releaseCtx(ctx *RequestCtx) { _ = "STUB: not implemented"; return }

func (s *Server) getServerName() string { _ = "STUB: not implemented"; return "" }

func (s *Server) writeFastError(w io.Writer, statusCode int, msg string) {
	_ = "STUB: not implemented"
	return
}

//nolint:errcheck

func defaultErrorHandler(ctx *RequestCtx, err error) { _ = "STUB: not implemented"; return }

func (s *Server) writeErrorResponse(bw *bufio.Writer, ctx *RequestCtx, serverName string, err error) *bufio.Writer {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

var idleConnTimePool sync.Pool

func (s *Server) closeIdleConns() { _ = "STUB: not implemented"; return }

func (s *Server) closeListenersLocked() error { _ = "STUB: not implemented"; return nil }

type ConnState int

const (
	StateNew ConnState = iota

	StateActive

	StateIdle

	StateHijacked

	StateClosed
)

var stateName = []string{
	StateNew:      "new",
	StateActive:   "active",
	StateIdle:     "idle",
	StateHijacked: "hijacked",
	StateClosed:   "closed",
}

func (c ConnState) String() string { _ = "STUB: not implemented"; return "" }
