package fasthttp

import (
	"sync"
	"time"
)

func initTimer(t *time.Timer, timeout time.Duration) *time.Timer {
	_ = "STUB: not implemented"
	return nil
}

func stopTimer(t *time.Timer) { _ = "STUB: not implemented"; return }

func AcquireTimer(timeout time.Duration) *time.Timer { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func ReleaseTimer(t *time.Timer) { _ = "STUB: not implemented"; return }

var timerPool sync.Pool
