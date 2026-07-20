package fasthttp

import (
	"net"
	"runtime"
	"sync"
	"time"
)

type workerPool struct {
	workerChanPool sync.Pool

	Logger Logger

	WorkerFunc ServeHandler

	stopCh chan struct{}

	connState func(net.Conn, ConnState)

	ready []*workerChan

	MaxWorkersCount int

	MaxIdleWorkerDuration time.Duration

	workersCount int

	lock sync.Mutex

	LogAllErrors bool
	mustStop     bool
}

type workerChan struct {
	lastUseTime time.Time
	ch          chan net.Conn
}

func (wp *workerPool) Start() { _ = "STUB: not implemented"; return }

func (wp *workerPool) Stop() { _ = "STUB: not implemented"; return }

func (wp *workerPool) getMaxIdleWorkerDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (wp *workerPool) clean(scratch *[]*workerChan) { _ = "STUB: not implemented"; return }

func (wp *workerPool) Serve(c net.Conn) bool { _ = "STUB: not implemented"; return false }

var workerChanCap = func() int {

	if runtime.GOMAXPROCS(0) == 1 {
		return 0
	}

	return 1
}()

func (wp *workerPool) getCh() *workerChan { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func (wp *workerPool) release(ch *workerChan) bool { _ = "STUB: not implemented"; return false }

func (wp *workerPool) workerFunc(ch *workerChan) { _ = "STUB: not implemented"; return }
