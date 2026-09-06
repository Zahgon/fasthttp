package fasthttp

type userDataKV struct {
	key   any
	value any
}

type userData []userDataKV

func (d *userData) Set(key, value any) { _ = "STUB: not implemented"; return }

func (d *userData) SetBytes(key []byte, value any) { _ = "STUB: not implemented"; return }

func (d *userData) Get(key any) any { _ = "STUB: not implemented"; return *new(any) }

func (d *userData) GetBytes(key []byte) any { _ = "STUB: not implemented"; return *new(any) }

func (d *userData) Reset() { _ = "STUB: not implemented"; return }

func (d *userData) Remove(key any) { _ = "STUB: not implemented"; return }

func (d *userData) RemoveBytes(key []byte) { _ = "STUB: not implemented"; return }
