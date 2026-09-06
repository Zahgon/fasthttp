package fasthttp

type headerScanner struct {
	initialized bool

	b []byte
	r int

	blockEnd int

	key   []byte
	value []byte

	keyHasSpace bool

	err error
}

func (s *headerScanner) next() bool { _ = "STUB: not implemented"; return false }

func (s *headerScanner) readLine() []byte { _ = "STUB: not implemented"; return nil }

func (s *headerScanner) readContinuedLineSlice() ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *headerScanner) skipSpace() bool { _ = "STUB: not implemented"; return false }

func isASCIILetter(b byte) bool { _ = "STUB: not implemented"; return false }

func trim(s []byte) []byte { _ = "STUB: not implemented"; return nil }

func trimTrailingSpace(s []byte) []byte { _ = "STUB: not implemented"; return nil }
