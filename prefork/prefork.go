package prefork

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	preforkChildEnvVariable = "FASTHTTP_PREFORK_CHILD"
	preforkChildEnvValue    = "1"
	defaultNetwork          = "tcp4"

	inheritedListenerFD = 3

	masterPollInterval = 500 * time.Millisecond

	defaultShutdownGracePeriod = 5 * time.Second
)

var (
	defaultLogger = Logger(log.New(os.Stderr, "", log.LstdFlags))

	tcpListenerFile = (*net.TCPListener).File

	ErrOverRecovery = errors.New("prefork: exceeding the value of recoverthreshold")

	ErrOnlyReuseportOnWindows = errors.New("prefork: windows only supports reuseport = true")

	ErrCommandProducerNilCmd = errors.New("prefork: commandproducer returned nil command")

	ErrCommandProducerNotStarted = errors.New("prefork: commandproducer must return a started command")
)

type Logger interface {
	Printf(format string, args ...any)
}

var _ Logger = fasthttp.Logger(nil)

type Prefork struct {
	Logger Logger

	ln net.Listener

	ServeFunc         func(ln net.Listener) error
	ServeTLSFunc      func(ln net.Listener, certFile, keyFile string) error
	ServeTLSEmbedFunc func(ln net.Listener, certData, keyData []byte) error

	Network string

	files []*os.File

	RecoverThreshold int

	RecoverInterval time.Duration

	ShutdownGracePeriod time.Duration

	Reuseport bool

	OnMasterDeath func()

	OnChildSpawn func(pid int) error

	OnMasterReady func(childPIDs []int) error

	OnChildRecover func(oldPID, newPID int)

	CommandProducer func(files []*os.File) (*exec.Cmd, error)
}

func IsChild() bool { _ = "STUB: not implemented"; return false }

func New(s *fasthttp.Server) *Prefork { _ = "STUB: not implemented"; return nil }

func defaultRecoverThreshold() int { _ = "STUB: not implemented"; return 0 }

func (p *Prefork) logger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (p *Prefork) watchMaster(masterPID int) { _ = "STUB: not implemented"; return }

func (p *Prefork) listen(addr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (p *Prefork) listenAsChild(addr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (p *Prefork) setTCPListenerFiles(addr string) error { _ = "STUB: not implemented"; return nil }

func childEnv() []string { _ = "STUB: not implemented"; return nil }

func (p *Prefork) doCommand() (*exec.Cmd, error) { _ = "STUB: not implemented"; return nil, nil }

type childExit struct {
	err error
	pid int
}

func (p *Prefork) shutdownChildren(
	childProcs map[int]*exec.Cmd,
	wg *sync.WaitGroup,
	cancel context.CancelFunc,
	grace time.Duration,
) {
	_ = "STUB: not implemented"
	return
}

func (p *Prefork) killChild(pid int, proc *exec.Cmd) { _ = "STUB: not implemented"; return }

func (p *Prefork) prefork(addr string) (err error) {
	_ = "STUB: not implemented" //nolint:gocyclo
	return nil
}

func (p *Prefork) ListenAndServe(addr string) error { _ = "STUB: not implemented"; return nil }

func (p *Prefork) ListenAndServeTLS(addr, certKey, certFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Prefork) ListenAndServeTLSEmbed(addr string, certData, keyData []byte) error {
	_ = "STUB: not implemented"
	return nil
}
