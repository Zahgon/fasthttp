//go:build !windows

package fasthttp

func hasWindowsReservedPathColon(_ []byte, _ bool) bool { _ = "STUB: not implemented"; return false }
