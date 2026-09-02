package fasthttp

import (
	"errors"
	"io"
	"iter"
	"sync"
)

const (
	argsNoValue  = true
	argsHasValue = false
)

func AcquireArgs() *Args { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func ReleaseArgs(a *Args) { _ = "STUB: not implemented"; return }

var argsPool = &sync.Pool{
	New: func() any {
		return &Args{}
	},
}

type Args struct {
	noCopy noCopy

	args []argsKV
	buf  []byte
}

type argsKV struct {
	key     []byte
	value   []byte
	noValue bool
}

func (a *Args) Reset() { _ = "STUB: not implemented"; return }

func (a *Args) CopyTo(dst *Args) { _ = "STUB: not implemented"; return }

func (a *Args) All() iter.Seq2[[]byte, []byte] { _ = "STUB: not implemented"; return nil }

func (a *Args) VisitAll(f func(key, value []byte)) { _ = "STUB: not implemented"; return }

func (a *Args) Len() int { _ = "STUB: not implemented"; return 0 }

func (a *Args) Parse(s string) { _ = "STUB: not implemented"; return }

func (a *Args) ParseBytes(b []byte) { _ = "STUB: not implemented"; return }

func (a *Args) String() string { _ = "STUB: not implemented"; return "" }

func (a *Args) QueryString() []byte { _ = "STUB: not implemented"; return nil }

func (a *Args) Sort(f func(x, y []byte) int) { _ = "STUB: not implemented"; return }

func (a *Args) SortKeys(f func(x, y []byte) int) { _ = "STUB: not implemented"; return }

func (a *Args) AppendBytes(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (a *Args) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (a *Args) Del(key string) { _ = "STUB: not implemented"; return }

func (a *Args) DelBytes(key []byte) { _ = "STUB: not implemented"; return }

func (a *Args) Add(key, value string) { _ = "STUB: not implemented"; return }

func (a *Args) AddBytesK(key []byte, value string) { _ = "STUB: not implemented"; return }

func (a *Args) AddBytesV(key string, value []byte) { _ = "STUB: not implemented"; return }

func (a *Args) AddBytesKV(key, value []byte) { _ = "STUB: not implemented"; return }

func (a *Args) AddNoValue(key string) { _ = "STUB: not implemented"; return }

func (a *Args) AddBytesKNoValue(key []byte) { _ = "STUB: not implemented"; return }

func (a *Args) Set(key, value string) { _ = "STUB: not implemented"; return }

func (a *Args) SetBytesK(key []byte, value string) { _ = "STUB: not implemented"; return }

func (a *Args) SetBytesV(key string, value []byte) { _ = "STUB: not implemented"; return }

func (a *Args) SetBytesKV(key, value []byte) { _ = "STUB: not implemented"; return }

func (a *Args) SetNoValue(key string) { _ = "STUB: not implemented"; return }

func (a *Args) SetBytesKNoValue(key []byte) { _ = "STUB: not implemented"; return }

func (a *Args) Peek(key string) []byte { _ = "STUB: not implemented"; return nil }

func (a *Args) PeekBytes(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (a *Args) PeekMulti(key string) [][]byte { _ = "STUB: not implemented"; return nil }

func (a *Args) PeekMultiBytes(key []byte) [][]byte { _ = "STUB: not implemented"; return nil }

func (a *Args) Has(key string) bool { _ = "STUB: not implemented"; return false }

func (a *Args) HasBytes(key []byte) bool { _ = "STUB: not implemented"; return false }

var ErrNoArgValue = errors.New("fasthttp: no args value for the given key")

func (a *Args) GetUint(key string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (a *Args) SetUint(key string, value int) { _ = "STUB: not implemented"; return }

func (a *Args) SetUintBytes(key []byte, value int) { _ = "STUB: not implemented"; return }

func (a *Args) GetUintOrZero(key string) int { _ = "STUB: not implemented"; return 0 }

func (a *Args) GetUfloat(key string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (a *Args) GetUfloatOrZero(key string) float64 { _ = "STUB: not implemented"; return 0 }

func (a *Args) GetBool(key string) bool { _ = "STUB: not implemented"; return false }

func copyArgs(dst, src []argsKV) []argsKV { _ = "STUB: not implemented"; return nil }

func delAllArgsStable(args []argsKV, key string) []argsKV { _ = "STUB: not implemented"; return nil }

func delAllArgs(args []argsKV, key string) []argsKV { _ = "STUB: not implemented"; return nil }

func setArgBytes(h []argsKV, key, value []byte, noValue bool) []argsKV {
	_ = "STUB: not implemented"
	return nil
}

func setArg(h []argsKV, key, value string, noValue bool) []argsKV {
	_ = "STUB: not implemented"
	return nil
}

func appendArgBytes(h []argsKV, key, value []byte, noValue bool) []argsKV {
	_ = "STUB: not implemented"
	return nil
}

func appendArg(args []argsKV, key, value string, noValue bool) []argsKV {
	_ = "STUB: not implemented"
	return nil
}

func allocArg(h []argsKV) ([]argsKV, *argsKV) { _ = "STUB: not implemented"; return nil, nil }

func releaseArg(h []argsKV) []argsKV { _ = "STUB: not implemented"; return nil }

func hasArg(h []argsKV, key string) bool { _ = "STUB: not implemented"; return false }

func peekArgBytes(h []argsKV, k []byte) []byte { _ = "STUB: not implemented"; return nil }

func peekArgStr(h []argsKV, k string) []byte { _ = "STUB: not implemented"; return nil }

type argsScanner struct {
	b []byte
}

func (s *argsScanner) next(kv *argsKV) bool { _ = "STUB: not implemented"; return false }

func decodeArgAppend(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func decodeArgAppendNoPlus(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func peekAllArgBytesToDst(dst [][]byte, h []argsKV, k []byte) [][]byte {
	_ = "STUB: not implemented"
	return nil
}
