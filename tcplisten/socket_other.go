//go:build !js && !wasm && (linux || dragonfly || freebsd || netbsd || openbsd || rumprun)

package tcplisten

func newSocketCloexec(domain, typ, proto int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
