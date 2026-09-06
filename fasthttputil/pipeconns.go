package fasthttputil

import (
	"errors"
	"net"
	"sync"
	"time"
)

func NewPipeConns() *PipeConns { _ = "STUB: not implemented"; return nil }

type PipeConns struct {
	stopCh     chan struct{}
	c1         pipeConn
	c2         pipeConn
	stopChLock sync.Mutex
}

func (pc *PipeConns) SetAddresses(localAddr1, remoteAddr1, localAddr2, remoteAddr2 net.Addr) {
	_ = "STUB: not implemented"
	return
}

func (pc *PipeConns) Conn1() net.Conn { _ = "STUB: not implemented"; return *new(net.Conn) }

func (pc *PipeConns) Conn2() net.Conn { _ = "STUB: not implemented"; return *new(net.Conn) }

func (pc *PipeConns) Close() error { _ = "STUB: not implemented"; return nil }

type pipeConn struct {
	localAddr  net.Addr
	remoteAddr net.Addr
	b          *byteBuffer

	rCh chan *byteBuffer
	wCh chan *byteBuffer
	pc  *PipeConns

	readDeadlineTimer  *time.Timer
	writeDeadlineTimer *time.Timer

	readDeadlineCh  <-chan time.Time
	writeDeadlineCh <-chan time.Time

	bb []byte

	addrLock sync.RWMutex

	readDeadlineChLock sync.Mutex
}

func (c *pipeConn) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *pipeConn) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *pipeConn) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *pipeConn) read(p []byte, mayBlock bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *pipeConn) readNextByteBuffer(mayBlock bool) error { _ = "STUB: not implemented"; return nil }

var errWouldBlock = errors.New("would block")

var ErrConnectionClosed = errors.New("fasthttputil: connection closed")

type timeoutError struct{}

func (e *timeoutError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *timeoutError) Timeout() bool { _ = "STUB: not implemented"; return false }

var ErrTimeout = &timeoutError{}

func (c *pipeConn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *pipeConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *pipeConn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *pipeConn) SetDeadline(deadline time.Time) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck
//nolint:errcheck

func (c *pipeConn) SetReadDeadline(deadline time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *pipeConn) SetWriteDeadline(deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func updateTimer(t *time.Timer, deadline time.Time) <-chan time.Time {
	_ = "STUB: not implemented"
	return nil
}

var closedDeadlineCh = func() <-chan time.Time {
	ch := make(chan time.Time)
	close(ch)
	return ch
}()

type pipeAddr int

func (pipeAddr) Network() string { _ = "STUB: not implemented"; return "" }

func (pipeAddr) String() string { _ = "STUB: not implemented"; return "" }

type byteBuffer struct {
	b []byte
}

func acquireByteBuffer() *byteBuffer { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func releaseByteBuffer(b *byteBuffer) { _ = "STUB: not implemented"; return }

var byteBufferPool = &sync.Pool{
	New: func() any {
		return &byteBuffer{
			b: make([]byte, 1024),
		}
	},
}
