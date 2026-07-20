package fasthttp

type noCopy struct{}

func (*noCopy) Lock()   { _ = "STUB: not implemented"; return }
func (*noCopy) Unlock() { _ = "STUB: not implemented"; return }
