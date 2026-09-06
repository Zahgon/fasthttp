package reuseport

type ErrNoReusePort struct {
	err error
}

func (e *ErrNoReusePort) Error() string { _ = "STUB: not implemented"; return "" }
