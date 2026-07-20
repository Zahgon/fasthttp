//go:build !js && !wasm && (linux || darwin || dragonfly || freebsd || netbsd || openbsd || rumprun || (zos && s390x))

package tcplisten

func newSocketCloexecOld(domain, typ, proto int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
