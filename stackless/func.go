package stackless

import (
	"sync"
)

func NewFunc(f func(ctx any)) func(ctx any) bool { _ = "STUB: not implemented"; return nil }

func funcWorker(funcWorkCh <-chan *funcWork, f func(ctx any)) { _ = "STUB: not implemented"; return }

func getFuncWork() *funcWork { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func putFuncWork(fw *funcWork) { _ = "STUB: not implemented"; return }

var funcWorkPool sync.Pool

type funcWork struct {
	ctx  any
	done chan struct{}
}
