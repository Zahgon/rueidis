// Code generated DO NOT EDIT

package cmds

type TopkAdd Incomplete

func (b Builder) TopkAdd() (c TopkAdd) { _ = "STUB: not implemented"; return *new(TopkAdd) }

func (c TopkAdd) Key(key string) TopkAddKey { _ = "STUB: not implemented"; return *new(TopkAddKey) }

type TopkAddItems Incomplete

func (c TopkAddItems) Items(items ...string) TopkAddItems {
	_ = "STUB: not implemented"
	return *new(TopkAddItems)
}

func (c TopkAddItems) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TopkAddKey Incomplete

func (c TopkAddKey) Items(items ...string) TopkAddItems {
	_ = "STUB: not implemented"
	return *new(TopkAddItems)
}

type TopkCount Incomplete

func (b Builder) TopkCount() (c TopkCount) { _ = "STUB: not implemented"; return *new(TopkCount) }

func (c TopkCount) Key(key string) TopkCountKey {
	_ = "STUB: not implemented"
	return *new(TopkCountKey)
}

type TopkCountItem Incomplete

func (c TopkCountItem) Item(item ...string) TopkCountItem {
	_ = "STUB: not implemented"
	return *new(TopkCountItem)
}

func (c TopkCountItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TopkCountKey Incomplete

func (c TopkCountKey) Item(item ...string) TopkCountItem {
	_ = "STUB: not implemented"
	return *new(TopkCountItem)
}

type TopkIncrby Incomplete

func (b Builder) TopkIncrby() (c TopkIncrby) { _ = "STUB: not implemented"; return *new(TopkIncrby) }

func (c TopkIncrby) Key(key string) TopkIncrbyKey {
	_ = "STUB: not implemented"
	return *new(TopkIncrbyKey)
}

type TopkIncrbyItemsIncrement Incomplete

func (c TopkIncrbyItemsIncrement) Item(item string) TopkIncrbyItemsItem {
	_ = "STUB: not implemented"
	return *new(TopkIncrbyItemsItem)
}

func (c TopkIncrbyItemsIncrement) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TopkIncrbyItemsItem Incomplete

func (c TopkIncrbyItemsItem) Increment(increment int64) TopkIncrbyItemsIncrement {
	_ = "STUB: not implemented"
	return *new(TopkIncrbyItemsIncrement)
}

type TopkIncrbyKey Incomplete

func (c TopkIncrbyKey) Item(item string) TopkIncrbyItemsItem {
	_ = "STUB: not implemented"
	return *new(TopkIncrbyItemsItem)
}

type TopkInfo Incomplete

func (b Builder) TopkInfo() (c TopkInfo) { _ = "STUB: not implemented"; return *new(TopkInfo) }

func (c TopkInfo) Key(key string) TopkInfoKey { _ = "STUB: not implemented"; return *new(TopkInfoKey) }

type TopkInfoKey Incomplete

func (c TopkInfoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c TopkInfoKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type TopkList Incomplete

func (b Builder) TopkList() (c TopkList) { _ = "STUB: not implemented"; return *new(TopkList) }

func (c TopkList) Key(key string) TopkListKey { _ = "STUB: not implemented"; return *new(TopkListKey) }

type TopkListKey Incomplete

func (c TopkListKey) Withcount() TopkListWithcount {
	_ = "STUB: not implemented"
	return *new(TopkListWithcount)
}

func (c TopkListKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c TopkListKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type TopkListWithcount Incomplete

func (c TopkListWithcount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c TopkListWithcount) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type TopkQuery Incomplete

func (b Builder) TopkQuery() (c TopkQuery) { _ = "STUB: not implemented"; return *new(TopkQuery) }

func (c TopkQuery) Key(key string) TopkQueryKey {
	_ = "STUB: not implemented"
	return *new(TopkQueryKey)
}

type TopkQueryItem Incomplete

func (c TopkQueryItem) Item(item ...string) TopkQueryItem {
	_ = "STUB: not implemented"
	return *new(TopkQueryItem)
}

func (c TopkQueryItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c TopkQueryItem) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type TopkQueryKey Incomplete

func (c TopkQueryKey) Item(item ...string) TopkQueryItem {
	_ = "STUB: not implemented"
	return *new(TopkQueryItem)
}

type TopkReserve Incomplete

func (b Builder) TopkReserve() (c TopkReserve) { _ = "STUB: not implemented"; return *new(TopkReserve) }

func (c TopkReserve) Key(key string) TopkReserveKey {
	_ = "STUB: not implemented"
	return *new(TopkReserveKey)
}

type TopkReserveKey Incomplete

func (c TopkReserveKey) Topk(topk int64) TopkReserveTopk {
	_ = "STUB: not implemented"
	return *new(TopkReserveTopk)
}

type TopkReserveParamsDecay Incomplete

func (c TopkReserveParamsDecay) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TopkReserveParamsDepth Incomplete

func (c TopkReserveParamsDepth) Decay(decay float64) TopkReserveParamsDecay {
	_ = "STUB: not implemented"
	return *new(TopkReserveParamsDecay)
}

type TopkReserveParamsWidth Incomplete

func (c TopkReserveParamsWidth) Depth(depth int64) TopkReserveParamsDepth {
	_ = "STUB: not implemented"
	return *new(TopkReserveParamsDepth)
}

type TopkReserveTopk Incomplete

func (c TopkReserveTopk) Width(width int64) TopkReserveParamsWidth {
	_ = "STUB: not implemented"
	return *new(TopkReserveParamsWidth)
}

func (c TopkReserveTopk) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }
