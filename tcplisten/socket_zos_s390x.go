//go:build zos && s390x

package tcplisten

func newSocketCloexec(domain, typ, proto int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:errcheck

//nolint:errcheck
