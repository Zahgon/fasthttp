package fasthttp

import (
	"errors"
	"io"
	"sync"
	"time"
)

var zeroTime time.Time

var (
	CookieExpireDelete = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	CookieExpireUnlimited = zeroTime
)

type CookieSameSite int

const (
	CookieSameSiteDisabled CookieSameSite = iota

	CookieSameSiteDefaultMode

	CookieSameSiteLaxMode

	CookieSameSiteStrictMode

	CookieSameSiteNoneMode
)

func AcquireCookie() *Cookie { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func ReleaseCookie(c *Cookie) { _ = "STUB: not implemented"; return }

var cookiePool = &sync.Pool{
	New: func() any {
		return &Cookie{}
	},
}

type Cookie struct {
	noCopy noCopy

	expire time.Time

	key    []byte
	value  []byte
	domain []byte
	path   []byte

	bufK []byte
	bufV []byte

	maxAge int

	sameSite    CookieSameSite
	httpOnly    bool
	secure      bool
	partitioned bool
}

func (c *Cookie) CopyTo(src *Cookie) { _ = "STUB: not implemented"; return }

func (c *Cookie) HTTPOnly() bool { _ = "STUB: not implemented"; return false }

func (c *Cookie) SetHTTPOnly(httpOnly bool) { _ = "STUB: not implemented"; return }

func (c *Cookie) Secure() bool { _ = "STUB: not implemented"; return false }

func (c *Cookie) SetSecure(secure bool) { _ = "STUB: not implemented"; return }

func (c *Cookie) SameSite() CookieSameSite { _ = "STUB: not implemented"; return *new(CookieSameSite) }

func (c *Cookie) SetSameSite(mode CookieSameSite) { _ = "STUB: not implemented"; return }

func (c *Cookie) Partitioned() bool { _ = "STUB: not implemented"; return false }

func (c *Cookie) SetPartitioned(partitioned bool) { _ = "STUB: not implemented"; return }

func (c *Cookie) Path() []byte { _ = "STUB: not implemented"; return nil }

func (c *Cookie) SetPath(path string) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetPathBytes(path []byte) { _ = "STUB: not implemented"; return }

func (c *Cookie) Domain() []byte { _ = "STUB: not implemented"; return nil }

func (c *Cookie) SetDomain(domain string) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetDomainBytes(domain []byte) { _ = "STUB: not implemented"; return }

func (c *Cookie) MaxAge() int { _ = "STUB: not implemented"; return 0 }

func (c *Cookie) SetMaxAge(seconds int) { _ = "STUB: not implemented"; return }

func (c *Cookie) Expire() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *Cookie) SetExpire(expire time.Time) { _ = "STUB: not implemented"; return }

func (c *Cookie) Value() []byte { _ = "STUB: not implemented"; return nil }

func (c *Cookie) SetValue(value string) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetValueBytes(value []byte) { _ = "STUB: not implemented"; return }

func (c *Cookie) Key() []byte { _ = "STUB: not implemented"; return nil }

func (c *Cookie) SetKey(key string) { _ = "STUB: not implemented"; return }

func (c *Cookie) SetKeyBytes(key []byte) { _ = "STUB: not implemented"; return }

func (c *Cookie) Reset() { _ = "STUB: not implemented"; return }

func (c *Cookie) AppendBytes(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (c *Cookie) Cookie() []byte { _ = "STUB: not implemented"; return nil }

func (c *Cookie) String() string { _ = "STUB: not implemented"; return "" }

func (c *Cookie) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

var (
	ErrNoCookies          = errors.New("fasthttp: no cookies found")
	ErrInvalidCookieValue = errors.New("fasthttp: invalid cookie value")
)

func (c *Cookie) Parse(src string) error { _ = "STUB: not implemented"; return nil }

func (c *Cookie) ParseBytes(src []byte) error { _ = "STUB: not implemented"; return nil }

func removeSemicolons(raw []byte) []byte { _ = "STUB: not implemented"; return nil }

func appendCookiePart(dst, key, value []byte) []byte { _ = "STUB: not implemented"; return nil }

func getCookieKey(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func appendRequestCookieBytes(dst []byte, cookies []argsKV) []byte {
	_ = "STUB: not implemented"
	return nil
}

func appendResponseCookieBytes(dst []byte, cookies []argsKV) []byte {
	_ = "STUB: not implemented"
	return nil
}

func parseRequestCookies(cookies []argsKV, src []byte) []argsKV {
	_ = "STUB: not implemented"
	return nil
}

type cookieScanner struct {
	b []byte
}

func (s *cookieScanner) nextRaw(key, val *[]byte) bool { _ = "STUB: not implemented"; return false }

func (s *cookieScanner) next(key, val *[]byte) bool { _ = "STUB: not implemented"; return false }

func decodeCookieArg(dst, src []byte, skipQuotes bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

func validCookieValue(value []byte) bool { _ = "STUB: not implemented"; return false }

func validCookiePathValue(value []byte) bool { _ = "STUB: not implemented"; return false }

func trimCookieArgNoCopy(src []byte, skipQuotes bool) []byte { _ = "STUB: not implemented"; return nil }

func parseCookieExpires(src []byte) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func caseInsensitiveCompare(a, b []byte) bool { _ = "STUB: not implemented"; return false }
