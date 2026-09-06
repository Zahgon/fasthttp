//go:generate go run bytesconv_table_gen.go

package fasthttp

import (
	"bufio"
	"errors"
	"math"
	"net"
	"strconv"
	"sync"
	"time"
)

func AppendHTMLEscape(dst []byte, s string) []byte { _ = "STUB: not implemented"; return nil }

func AppendHTMLEscapeBytes(dst, s []byte) []byte { _ = "STUB: not implemented"; return nil }

func AppendIPv4(dst []byte, ip net.IP) []byte { _ = "STUB: not implemented"; return nil }

var errEmptyIPStr = errors.New("empty ip address string")

var httpDateGMT = time.FixedZone("GMT", 0)

func ParseIPv4(dst net.IP, ipStr []byte) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func AppendHTTPDate(dst []byte, date time.Time) []byte { _ = "STUB: not implemented"; return nil }

func ParseHTTPDate(date []byte) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func parseRFC1123DateGMT(b []byte) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func isWeekday3(a, b, c byte) bool { _ = "STUB: not implemented"; return false }

func parse2Digits(a, b byte) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func parse4Digits(a, b, c, d byte) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func parseMonth3(a, b, c byte) (time.Month, bool) {
	_ = "STUB: not implemented"
	return *new(time.Month), false
}

func AppendUint(dst []byte, n int) []byte { _ = "STUB: not implemented"; return nil }

func ParseUint(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

var (
	errEmptyInt               = errors.New("empty integer")
	errIPv4PartTooLarge       = errors.New("ip part cannot exceed 255")
	errUnexpectedFirstChar    = errors.New("unexpected first char found: expecting 0-9")
	errUnexpectedTrailingChar = errors.New("unexpected trailing char found: expecting 0-9")
	errTooLongInt             = errors.New("too long int")
)

const (
	maxIntDiv10 = math.MaxInt / 10

	maxSafeIntDigits = 9 * (strconv.IntSize / 32)
)

func parseUintBuf(b []byte) (int, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func parseIPv4Octet(b []byte) (byte, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func ParseUfloat(buf []byte) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

var (
	errEmptyHexNum    = errors.New("empty hex number")
	errTooLargeHexNum = errors.New("too large hex number")
)

func readHexInt(r *bufio.Reader) (int, error) { _ = "STUB: not implemented"; return 0, nil }

var hexIntBufPool sync.Pool

func writeHexInt(w *bufio.Writer, n int) error { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

const (
	upperhex = "0123456789ABCDEF"
	lowerhex = "0123456789abcdef"
)

func lowercaseBytes(b []byte) { _ = "STUB: not implemented"; return }

func AppendUnquotedArg(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func AppendQuotedArg(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func appendQuotedPath(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }
