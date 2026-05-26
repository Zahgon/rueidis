package rueidis

type bitmap struct {
	exts []uint64  // 24 bytes
	bits [3]uint64 // 24 bytes
}

func (b *bitmap) Init(n int) { _ = "STUB: not implemented"; return }

func (b *bitmap) Set(i int) { _ = "STUB: not implemented"; return }

func (b *bitmap) Get(i int) bool { _ = "STUB: not implemented"; return false }

func (b *bitmap) Len() int { _ = "STUB: not implemented"; return 0 }
