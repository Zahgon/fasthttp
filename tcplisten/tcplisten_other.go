//go:build darwin || dragonfly || freebsd || netbsd || openbsd || rumprun || (zos && s390x)

package tcplisten

import "golang.org/x/sys/unix"

const soReusePort = unix.SO_REUSEPORT

func enableDeferAccept(fd int) error { _ = "STUB: not implemented"; return nil }

func enableFastOpen(fd int) error { _ = "STUB: not implemented"; return nil }

func soMaxConn() (int, error) { _ = "STUB: not implemented"; return 0, nil }
