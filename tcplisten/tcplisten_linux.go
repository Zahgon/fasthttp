//go:build linux

package tcplisten

const (
	soReusePort = 0x0F
	tcpFastOpen = 0x17
)

func enableDeferAccept(fd int) error { _ = "STUB: not implemented"; return nil }

func enableFastOpen(fd int) error { _ = "STUB: not implemented"; return nil }

const fastOpenQlen = 16 * 1024

func soMaxConn() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func kernelVersion() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func maxAckBacklog(n int) int { _ = "STUB: not implemented"; return 0 }

const soMaxConnFilePath = "/proc/sys/net/core/somaxconn"
