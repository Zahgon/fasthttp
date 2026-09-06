package fasthttp

import (
	"errors"
)

var (
	errInvalidIPv6Host    = errors.New("invalid ipv6 host")
	errInvalidIPv6Zone    = errors.New("invalid ipv6 zone")
	errInvalidIPv6Address = errors.New("invalid ipv6 address")
)

func validateIPv6Literal(host []byte) error { _ = "STUB: not implemented"; return nil }

func parseIPv6Hextets(s []byte, allowTrailingColon bool) (groups int, seenDouble, ok bool) {
	_ = "STUB: not implemented"
	return 0, false, false
}

func validIPv4(s []byte) bool { _ = "STUB: not implemented"; return false }
