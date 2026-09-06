//go:build !amd64 && !arm64 && !ppc64 && !ppc64le && !riscv64 && !s390x

package fasthttp

func roundUpForSliceCap(n int) int { _ = "STUB: not implemented"; return 0 }
