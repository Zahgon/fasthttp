package fasthttp

import (
	"bufio"
	"crypto/tls"
	"errors"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

func Do(req *Request, resp *Response) error { _ = "STUB: not implemented"; return nil }

func DoTimeout(req *Request, resp *Response, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func DoDeadline(req *Request, resp *Response, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func DoRedirects(req *Request, resp *Response, maxRedirectsCount int) error {
	_ = "STUB: not implemented"
	return nil
}

func Get(dst []byte, url string) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func GetTimeout(dst []byte, url string, timeout time.Duration) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func GetDeadline(dst []byte, url string, deadline time.Time) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func Post(dst []byte, url string, postArgs *Args) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

var defaultClient Client

type Client struct {
	noCopy noCopy

	readerPool sync.Pool
	writerPool sync.Pool

	Transport RoundTripper

	DialTimeout DialFuncWithTimeout

	Dial DialFunc

	TLSConfig *tls.Config

	RetryIf RetryIfFunc

	RetryIfErr RetryIfErrFunc

	RetryIfErrUpstream RetryIfErrUpstreamFunc

	ConfigureClient func(hc *HostClient) error

	m  map[string]*HostClient
	ms map[string]*HostClient

	Name string

	MaxConnsPerHost int

	MaxIdleConnDuration time.Duration

	MaxConnDuration time.Duration

	MaxIdemponentCallAttempts int

	ReadBufferSize int

	WriteBufferSize int

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	MaxResponseBodySize int

	MaxConnWaitTimeout time.Duration

	ConnPoolStrategy ConnPoolStrategyType

	mLock sync.RWMutex
	mOnce sync.Once

	NoDefaultUserAgentHeader bool

	DialDualStack bool

	DisableHeaderNamesNormalizing bool

	DisablePathNormalizing bool

	StreamResponseBody bool
}

func (c *Client) Get(dst []byte, url string) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *Client) GetTimeout(dst []byte, url string, timeout time.Duration) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *Client) GetDeadline(dst []byte, url string, deadline time.Time) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *Client) Post(dst []byte, url string, postArgs *Args) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *Client) DoTimeout(req *Request, resp *Response, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DoDeadline(req *Request, resp *Response, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DoRedirects(req *Request, resp *Response, maxRedirectsCount int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Do(req *Request, resp *Response) error { _ = "STUB: not implemented"; return nil }

func (c *Client) hostClient(host []byte, isTLS bool) (*HostClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) CloseIdleConnections() { _ = "STUB: not implemented"; return }

func (c *Client) ConnsCount() int { _ = "STUB: not implemented"; return 0 }

func (c *Client) IdleConnsCount() int { _ = "STUB: not implemented"; return 0 }

func (c *Client) mCleaner(m map[string]*HostClient) { _ = "STUB: not implemented"; return }

const DefaultMaxConnsPerHost = 512

const DefaultMaxIdleConnDuration = 10 * time.Second

const DefaultMaxIdemponentCallAttempts = 5

type DialFunc func(addr string) (net.Conn, error)

type DialFuncWithTimeout func(addr string, timeout time.Duration) (net.Conn, error)

type RetryIfFunc func(request *Request) bool

type RetryIfErrFunc func(request *Request, attempts int, err error) (resetTimeout bool, retry bool)

type RetryIfErrUpstreamFunc func(request *Request, attempts int, err error, upstream string) (resetTimeout bool, retry bool)

type RoundTripper interface {
	RoundTrip(hc *HostClient, req *Request, resp *Response) (retry bool, err error)
}

type ConnPoolStrategyType int

const (
	FIFO ConnPoolStrategyType = iota
	LIFO
)

type HostClient struct {
	noCopy noCopy

	readerPool sync.Pool
	writerPool sync.Pool

	Transport RoundTripper

	DialTimeout DialFuncWithTimeout

	Dial DialFunc

	TLSConfig *tls.Config

	RetryIf RetryIfFunc

	RetryIfErr RetryIfErrFunc

	RetryIfErrUpstream RetryIfErrUpstreamFunc

	connsWait *wantConnQueue

	tlsConfigMap map[string]*tls.Config

	clientReaderPool *sync.Pool
	clientWriterPool *sync.Pool

	Addr string

	Name string

	conns []*clientConn
	addrs []string

	MaxConns int

	MaxConnDuration time.Duration

	MaxIdleConnDuration time.Duration

	MaxIdemponentCallAttempts int

	ReadBufferSize int

	WriteBufferSize int

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	MaxResponseBodySize int

	MaxConnWaitTimeout time.Duration

	ConnPoolStrategy ConnPoolStrategyType

	connsCount int

	connsLock sync.Mutex

	addrsLock        sync.Mutex
	tlsConfigMapLock sync.Mutex

	addrIdx     uint32
	lastUseTime uint32

	pendingRequests int32

	pendingClientRequests int32

	NoDefaultUserAgentHeader bool

	DialDualStack bool

	IsTLS bool

	DisableHeaderNamesNormalizing bool

	DisablePathNormalizing bool

	SecureErrorLogMessage bool

	StreamResponseBody bool

	connsCleanerRun bool
}

type clientConn struct {
	c net.Conn

	createdTime time.Time
	lastUseTime time.Time
}

func (cc *clientConn) Conn() net.Conn { _ = "STUB: not implemented"; return *new(net.Conn) }

func (cc *clientConn) CreatedTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (cc *clientConn) LastUseTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

var startTimeUnix = time.Now().Unix()

func (c *HostClient) LastUseTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *HostClient) Get(dst []byte, url string) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *HostClient) GetTimeout(dst []byte, url string, timeout time.Duration) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *HostClient) GetDeadline(dst []byte, url string, deadline time.Time) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *HostClient) Post(dst []byte, url string, postArgs *Args) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

type clientDoer interface {
	Do(req *Request, resp *Response) error
}

func clientGetURL(dst []byte, url string, c clientDoer) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func clientGetURLTimeout(dst []byte, url string, timeout time.Duration, c clientDoer) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

type clientURLResponse struct {
	err        error
	body       []byte
	statusCode int
}

func clientGetURLDeadline(dst []byte, url string, deadline time.Time, c clientDoer) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

//nolint:forcetypeassert

var clientURLResponseChPool sync.Pool

func clientPostURL(dst []byte, url string, postArgs *Args, c clientDoer) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

var (
	ErrMissingLocation = errors.New("fasthttp: missing location header for http redirect")

	ErrTooManyRedirects = errors.New("fasthttp: too many redirects detected when doing the request")

	ErrRedirectBodyStream = errors.New("fasthttp: cannot follow a body-preserving redirect for a request with a body stream")

	ErrHostClientRedirectToDifferentScheme = errors.New("fasthttp: hostclient can't follow redirects to a different protocol," +
		" please use client instead")
)

const defaultMaxRedirectsCount = 16

const maxResponseBodyDrainSize = 8 * 1024

func closeOrDrainResponseBody(resp *Response, maxDrainSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func responseBodyDrainSize(maxBodySize int) int { _ = "STUB: not implemented"; return 0 }

func doRequestFollowRedirectsBuffer(req *Request, dst []byte, url string, c clientDoer) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func doRequestFollowRedirects(
	req *Request, resp *Response, url string, maxRedirectsCount int, c clientDoer,
) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func getRedirectURL(baseURL string, location []byte, disablePathNormalizing bool, dst *URI) string {
	_ = "STUB: not implemented"
	return ""
}

func stripSensitiveHeadersOnRedirect(req *Request, initialHost []byte, redirectURI *URI) {
	_ = "STUB: not implemented"
	return
}

func shouldStripSensitiveHeadersOnRedirect(initialHost, redirectHostPort []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func hostnameFromURLString(url string) []byte { _ = "STUB: not implemented"; return nil }

func hostnameFromHostPortBytes(hostPort []byte) []byte { _ = "STUB: not implemented"; return nil }

func isDomainOrSubdomainBytes(sub, parent []byte) bool { _ = "STUB: not implemented"; return false }

func splitHostPortBytes(hostPort []byte) ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func StatusCodeIsRedirect(statusCode int) bool { _ = "STUB: not implemented"; return false }

var (
	requestPool  sync.Pool
	responsePool sync.Pool
)

func AcquireRequest() *Request { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func ReleaseRequest(req *Request) { _ = "STUB: not implemented"; return }

func AcquireResponse() *Response { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func ReleaseResponse(resp *Response) { _ = "STUB: not implemented"; return }

func (c *HostClient) DoTimeout(req *Request, resp *Response, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *HostClient) DoDeadline(req *Request, resp *Response, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *HostClient) DoRedirects(req *Request, resp *Response, maxRedirectsCount int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *HostClient) Do(req *Request, resp *Response) error { _ = "STUB: not implemented"; return nil }

func (c *HostClient) PendingRequests() int { _ = "STUB: not implemented"; return 0 }

func isIdempotent(req *Request) bool { _ = "STUB: not implemented"; return false }

func (c *HostClient) do(req *Request, resp *Response) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *HostClient) doNonNilReqResp(req *Request, resp *Response) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *HostClient) transport() RoundTripper { _ = "STUB: not implemented"; return *new(RoundTripper) }

var (
	ErrNoFreeConns = errors.New("fasthttp: no free connections available to host")

	ErrConnectionClosed = errors.New("fasthttp: the server closed connection before returning the first response byte. " +
		"make sure the server returns 'connection: close' response header before closing the connection")

	ErrConnPoolStrategyNotImpl = errors.New("fasthttp: connection pool strategy is not implement")
)

type timeoutError struct{}

func (e *timeoutError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *timeoutError) Timeout() bool { _ = "STUB: not implemented"; return false }

var ErrTimeout = &timeoutError{}

func (c *HostClient) SetMaxConns(newMaxConns int) { _ = "STUB: not implemented"; return }

func (c *HostClient) AcquireConn(reqTimeout time.Duration, connectionClose bool) (cc *clientConn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:dupword

func (c *HostClient) queueForIdle(w *wantConn) { _ = "STUB: not implemented"; return }

func (c *HostClient) dialConnFor(w *wantConn) { _ = "STUB: not implemented"; return }

func (c *HostClient) CloseIdleConnections() { _ = "STUB: not implemented"; return }

func (c *HostClient) connsCleaner() { _ = "STUB: not implemented"; return }

func (c *HostClient) CloseConn(cc *clientConn) { _ = "STUB: not implemented"; return }

func (c *HostClient) decConnsCount() { _ = "STUB: not implemented"; return }

func (c *HostClient) ConnsCount() int { _ = "STUB: not implemented"; return 0 }

func (c *HostClient) IdleConnsCount() int { _ = "STUB: not implemented"; return 0 }

func acquireClientConn(conn net.Conn) *clientConn { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func releaseClientConn(cc *clientConn) { _ = "STUB: not implemented"; return }

var clientConnPool sync.Pool

func (c *HostClient) ReleaseConn(cc *clientConn) { _ = "STUB: not implemented"; return }

func (c *HostClient) AcquireWriter(conn net.Conn) *bufio.Writer {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

func (c *HostClient) ReleaseWriter(bw *bufio.Writer) { _ = "STUB: not implemented"; return }

func (c *HostClient) AcquireReader(conn net.Conn) *bufio.Reader {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

func (c *HostClient) ReleaseReader(br *bufio.Reader) { _ = "STUB: not implemented"; return }

func newClientTLSConfig(c *tls.Config, addr string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tlsServerName(addr string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *HostClient) nextAddr() string { _ = "STUB: not implemented"; return "" }

func (c *HostClient) dialHostHard(dialTimeout time.Duration) (conn net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (c *HostClient) cachedTLSConfig(addr string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var ErrTLSHandshakeTimeout = errors.New("fasthttp: tls handshake timed out")

func tlsClientHandshake(rawConn net.Conn, tlsConfig *tls.Config, deadline time.Time) (_ net.Conn, retErr error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func dialAddr(
	addr string, dial DialFunc, dialWithTimeout DialFuncWithTimeout, dialDualStack, isTLS bool,
	tlsConfig *tls.Config, dialTimeout, writeTimeout time.Duration,
) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func callDialFunc(
	addr string, dial DialFunc, dialWithTimeout DialFuncWithTimeout, dialDualStack, isTLS bool, timeout time.Duration,
) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func AddMissingPort(addr string, isTLS bool) string { _ = "STUB: not implemented"; return "" }

type wantConn struct {
	err   error
	ready chan struct{}
	conn  *clientConn
	mu    sync.Mutex
}

func (w *wantConn) waiting() bool { _ = "STUB: not implemented"; return false }

func (w *wantConn) tryDeliver(conn *clientConn, err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *wantConn) cancel(c *HostClient, err error) { _ = "STUB: not implemented"; return }

type wantConnQueue struct {
	head    []*wantConn
	tail    []*wantConn
	headPos int
}

func (q *wantConnQueue) len() int { _ = "STUB: not implemented"; return 0 }

func (q *wantConnQueue) pushBack(w *wantConn) { _ = "STUB: not implemented"; return }

func (q *wantConnQueue) popFront() *wantConn { _ = "STUB: not implemented"; return nil }

func (q *wantConnQueue) peekFront() *wantConn { _ = "STUB: not implemented"; return nil }

func (q *wantConnQueue) clearFront() (cleaned bool) { _ = "STUB: not implemented"; return false }

type PipelineClient struct {
	noCopy noCopy

	Logger Logger

	Dial DialFunc

	TLSConfig *tls.Config

	Addr string

	Name string

	connClients []*pipelineConnClient

	MaxConns int

	MaxPendingRequests int

	MaxBatchDelay time.Duration

	MaxIdleConnDuration time.Duration

	ReadBufferSize int

	WriteBufferSize int

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	connClientsLock sync.Mutex

	NoDefaultUserAgentHeader bool

	DialDualStack bool

	DisableHeaderNamesNormalizing bool

	DisablePathNormalizing bool

	IsTLS bool
}

type pipelineConnClient struct {
	noCopy noCopy

	workPool sync.Pool

	Logger Logger

	Dial      DialFunc
	TLSConfig *tls.Config

	tlsConfig *tls.Config
	chs       *pipelineConnChannels

	Addr                string
	Name                string
	MaxPendingRequests  int
	MaxBatchDelay       time.Duration
	MaxIdleConnDuration time.Duration
	ReadBufferSize      int
	WriteBufferSize     int
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration

	chLock sync.Mutex

	tlsConfigLock                 sync.Mutex
	NoDefaultUserAgentHeader      bool
	DialDualStack                 bool
	DisableHeaderNamesNormalizing bool
	DisablePathNormalizing        bool
	IsTLS                         bool
}

type pipelineConnChannels struct {
	chW   chan *pipelineWork
	chR   chan *pipelineWork
	users int
}

type pipelineWork struct {
	respCopy Response
	deadline time.Time
	err      error
	req      *Request
	resp     *Response
	t        *time.Timer
	done     chan struct{}
	reqCopy  Request
}

func (c *PipelineClient) DoTimeout(req *Request, resp *Response, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PipelineClient) DoDeadline(req *Request, resp *Response, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *pipelineConnClient) DoDeadline(req *Request, resp *Response, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *pipelineConnClient) acquirePipelineWork(timeout time.Duration) (w *pipelineWork) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

func (c *pipelineConnClient) releasePipelineWork(w *pipelineWork) {
	_ = "STUB: not implemented"
	return
}

func (c *PipelineClient) Do(req *Request, resp *Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *pipelineConnClient) Do(req *Request, resp *Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PipelineClient) getConnClient() *pipelineConnClient { _ = "STUB: not implemented"; return nil }

func (c *PipelineClient) getConnClientUnlocked() *pipelineConnClient {
	_ = "STUB: not implemented"
	return nil
}

func (c *PipelineClient) newConnClient() *pipelineConnClient { _ = "STUB: not implemented"; return nil }

var ErrPipelineOverflow = errors.New("fasthttp: pipelined requests' queue has been overflowed. " +
	"increase maxconns and/or maxpendingrequests")

const DefaultMaxPendingRequests = 1024

func (c *pipelineConnClient) acquirePipelineConnChannels() *pipelineConnChannels {
	_ = "STUB: not implemented"
	return nil
}

func (c *pipelineConnClient) releasePipelineConnChannels(chs *pipelineConnChannels) {
	_ = "STUB: not implemented"
	return
}

func (c *pipelineConnClient) pipelineWorker(chs *pipelineConnChannels) {
	_ = "STUB: not implemented"
	return
}

func (c *pipelineConnClient) tryRetirePipelineConnChannels(chs *pipelineConnChannels) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *pipelineConnClient) worker(chs *pipelineConnChannels) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *pipelineConnClient) cachedTLSConfig() (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *pipelineConnClient) ensureTLSConfig() error { _ = "STUB: not implemented"; return nil }

func (c *pipelineConnClient) writer(conn net.Conn, stopCh <-chan struct{}, chs *pipelineConnChannels) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *pipelineConnClient) canStopPipelineConn(chs *pipelineConnChannels) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *pipelineConnClient) reader(conn net.Conn, stopCh <-chan struct{}, chs *pipelineConnChannels) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *pipelineConnClient) logger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (c *PipelineClient) PendingRequests() int { _ = "STUB: not implemented"; return 0 }

func (c *pipelineConnClient) PendingRequests() int { _ = "STUB: not implemented"; return 0 }

var errPipelineConnStopped = errors.New("pipeline connection has been stopped")

var DefaultTransport RoundTripper = &transport{}

type transport struct{}

type clientStreamBody struct {
	reader    io.Reader
	interrupt func()
	release   func(bool)
	closed    atomic.Bool
	fullyRead bool
	readLock  sync.Mutex
	closeOnce sync.Once
}

func (s *clientStreamBody) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *clientStreamBody) CloseWithError(err error) error { _ = "STUB: not implemented"; return nil }

func (t *transport) RoundTrip(hc *HostClient, req *Request, resp *Response) (retry bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

//nolint:forcetypeassert
