// Code generated DO NOT EDIT

package cmds

type Pfadd Incomplete

func (b Builder) Pfadd() (c Pfadd) { _ = "STUB: not implemented"; return *new(Pfadd) }

func (c Pfadd) Key(key string) PfaddKey { _ = "STUB: not implemented"; return *new(PfaddKey) }

type PfaddElement Incomplete

func (c PfaddElement) Element(element ...string) PfaddElement {
	_ = "STUB: not implemented"
	return *new(PfaddElement)
}

func (c PfaddElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PfaddKey Incomplete

func (c PfaddKey) Element(element ...string) PfaddElement {
	_ = "STUB: not implemented"
	return *new(PfaddElement)
}

func (c PfaddKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Pfcount Incomplete

func (b Builder) Pfcount() (c Pfcount) { _ = "STUB: not implemented"; return *new(Pfcount) }

func (c Pfcount) Key(key ...string) PfcountKey { _ = "STUB: not implemented"; return *new(PfcountKey) }

type PfcountKey Incomplete

func (c PfcountKey) Key(key ...string) PfcountKey {
	_ = "STUB: not implemented"
	return *new(PfcountKey)
}

func (c PfcountKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Pfmerge Incomplete

func (b Builder) Pfmerge() (c Pfmerge) { _ = "STUB: not implemented"; return *new(Pfmerge) }

func (c Pfmerge) Destkey(destkey string) PfmergeDestkey {
	_ = "STUB: not implemented"
	return *new(PfmergeDestkey)
}

type PfmergeDestkey Incomplete

func (c PfmergeDestkey) Sourcekey(sourcekey ...string) PfmergeSourcekey {
	_ = "STUB: not implemented"
	return *new(PfmergeSourcekey)
}

func (c PfmergeDestkey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PfmergeSourcekey Incomplete

func (c PfmergeSourcekey) Sourcekey(sourcekey ...string) PfmergeSourcekey {
	_ = "STUB: not implemented"
	return *new(PfmergeSourcekey)
}

func (c PfmergeSourcekey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }
