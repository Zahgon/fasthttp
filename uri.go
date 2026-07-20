package fasthttp

import (
	"errors"
	"io"
	"sync"
)

func AcquireURI() *URI { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

func ReleaseURI(u *URI) { _ = "STUB: not implemented"; return }

var uriPool = &sync.Pool{
	New: func() any {
		return &URI{}
	},
}

type URI struct {
	noCopy noCopy

	queryArgs Args

	pathOriginal []byte
	scheme       []byte
	path         []byte
	queryString  []byte
	hash         []byte
	host         []byte

	fullURI    []byte
	requestURI []byte

	username        []byte
	password        []byte
	parsedQueryArgs bool

	DisablePathNormalizing bool
}

func (u *URI) CopyTo(dst *URI) { _ = "STUB: not implemented"; return }

func (u *URI) Hash() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetHash(hash string) { _ = "STUB: not implemented"; return }

func (u *URI) SetHashBytes(hash []byte) { _ = "STUB: not implemented"; return }

func (u *URI) Username() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetUsername(username string) { _ = "STUB: not implemented"; return }

func (u *URI) SetUsernameBytes(username []byte) { _ = "STUB: not implemented"; return }

func (u *URI) Password() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetPassword(password string) { _ = "STUB: not implemented"; return }

func (u *URI) SetPasswordBytes(password []byte) { _ = "STUB: not implemented"; return }

func (u *URI) QueryString() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetQueryString(queryString string) { _ = "STUB: not implemented"; return }

func (u *URI) SetQueryStringBytes(queryString []byte) { _ = "STUB: not implemented"; return }

func (u *URI) Path() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetPath(path string) { _ = "STUB: not implemented"; return }

func (u *URI) SetPathBytes(path []byte) { _ = "STUB: not implemented"; return }

func (u *URI) PathOriginal() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) Scheme() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetScheme(scheme string) { _ = "STUB: not implemented"; return }

func (u *URI) SetSchemeBytes(scheme []byte) { _ = "STUB: not implemented"; return }

func (u *URI) isHTTPS() bool { _ = "STUB: not implemented"; return false }

func (u *URI) isHTTP() bool { _ = "STUB: not implemented"; return false }

func (u *URI) Reset() { _ = "STUB: not implemented"; return }

func (u *URI) Host() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) SetHost(host string) { _ = "STUB: not implemented"; return }

func (u *URI) SetHostBytes(host []byte) { _ = "STUB: not implemented"; return }

var ErrorInvalidURI = errors.New("fasthttp: invalid uri")

func (u *URI) Parse(host, uri []byte) error { _ = "STUB: not implemented"; return nil }

func (u *URI) parse(host, uri []byte, isTLS bool) error { _ = "STUB: not implemented"; return nil }

func validUserinfo(userinfo []byte) bool { _ = "STUB: not implemented"; return false }

func isValidScheme(scheme []byte) bool { _ = "STUB: not implemented"; return false }

func parseHost(host []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type encoding int

const (
	encodeHost encoding = 1 + iota
	encodeZone
)

type EscapeError string

func (e EscapeError) Error() string { _ = "STUB: not implemented"; return "" }

type InvalidHostError string

func (e InvalidHostError) Error() string { _ = "STUB: not implemented"; return "" }

func unescape(s []byte, mode encoding) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func shouldEscape(c byte, mode encoding) bool { _ = "STUB: not implemented"; return false }

func ishex(c byte) bool { _ = "STUB: not implemented"; return false }

func unhex(c byte) byte { _ = "STUB: not implemented"; return 0 }

func validOptionalPort(port []byte) bool { _ = "STUB: not implemented"; return false }

func normalizePath(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) RequestURI() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) LastPathSegment() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) Update(newURI string) { _ = "STUB: not implemented"; return }

func (u *URI) UpdateBytes(newURI []byte) { _ = "STUB: not implemented"; return }

func (u *URI) updateBytes(newURI, buf []byte) []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) FullURI() []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) AppendBytes(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) appendSchemeHost(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (u *URI) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (u *URI) String() string { _ = "STUB: not implemented"; return "" }

func splitHostURI(host, uri []byte) ([]byte, []byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (u *URI) QueryArgs() *Args { _ = "STUB: not implemented"; return nil }

func (u *URI) parseQueryArgs() { _ = "STUB: not implemented"; return }

func stringContainsCTLByte(s []byte) bool { _ = "STUB: not implemented"; return false }
