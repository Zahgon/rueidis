// Code generated DO NOT EDIT

package cmds

type BfAdd Incomplete

func (b Builder) BfAdd() (c BfAdd) { _ = "STUB: not implemented"; return *new(BfAdd) }

func (c BfAdd) Key(key string) BfAddKey { _ = "STUB: not implemented"; return *new(BfAddKey) }

type BfAddItem Incomplete

func (c BfAddItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BfAddKey Incomplete

func (c BfAddKey) Item(item string) BfAddItem { _ = "STUB: not implemented"; return *new(BfAddItem) }

type BfCard Incomplete

func (b Builder) BfCard() (c BfCard) { _ = "STUB: not implemented"; return *new(BfCard) }

func (c BfCard) Key(key string) BfCardKey { _ = "STUB: not implemented"; return *new(BfCardKey) }

type BfCardKey Incomplete

func (c BfCardKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BfExists Incomplete

func (b Builder) BfExists() (c BfExists) { _ = "STUB: not implemented"; return *new(BfExists) }

func (c BfExists) Key(key string) BfExistsKey { _ = "STUB: not implemented"; return *new(BfExistsKey) }

type BfExistsItem Incomplete

func (c BfExistsItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c BfExistsItem) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type BfExistsKey Incomplete

func (c BfExistsKey) Item(item string) BfExistsItem {
	_ = "STUB: not implemented"
	return *new(BfExistsItem)
}

type BfInfo Incomplete

func (b Builder) BfInfo() (c BfInfo) { _ = "STUB: not implemented"; return *new(BfInfo) }

func (c BfInfo) Key(key string) BfInfoKey { _ = "STUB: not implemented"; return *new(BfInfoKey) }

type BfInfoKey Incomplete

func (c BfInfoKey) Capacity() BfInfoSingleValueCapacity {
	_ = "STUB: not implemented"
	return *new(BfInfoSingleValueCapacity)
}

func (c BfInfoKey) Size() BfInfoSingleValueSize {
	_ = "STUB: not implemented"
	return *new(BfInfoSingleValueSize)
}

func (c BfInfoKey) Filters() BfInfoSingleValueFilters {
	_ = "STUB: not implemented"
	return *new(BfInfoSingleValueFilters)
}

func (c BfInfoKey) Items() BfInfoSingleValueItems {
	_ = "STUB: not implemented"
	return *new(BfInfoSingleValueItems)
}

func (c BfInfoKey) Expansion() BfInfoSingleValueExpansion {
	_ = "STUB: not implemented"
	return *new(BfInfoSingleValueExpansion)
}

func (c BfInfoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c BfInfoKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type BfInfoSingleValueCapacity Incomplete

func (c BfInfoSingleValueCapacity) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c BfInfoSingleValueCapacity) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type BfInfoSingleValueExpansion Incomplete

func (c BfInfoSingleValueExpansion) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c BfInfoSingleValueExpansion) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type BfInfoSingleValueFilters Incomplete

func (c BfInfoSingleValueFilters) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c BfInfoSingleValueFilters) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type BfInfoSingleValueItems Incomplete

func (c BfInfoSingleValueItems) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c BfInfoSingleValueItems) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type BfInfoSingleValueSize Incomplete

func (c BfInfoSingleValueSize) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c BfInfoSingleValueSize) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type BfInsert Incomplete

func (b Builder) BfInsert() (c BfInsert) { _ = "STUB: not implemented"; return *new(BfInsert) }

func (c BfInsert) Key(key string) BfInsertKey { _ = "STUB: not implemented"; return *new(BfInsertKey) }

type BfInsertCapacity Incomplete

func (c BfInsertCapacity) Error(error float64) BfInsertError {
	_ = "STUB: not implemented"
	return *new(BfInsertError)
}

func (c BfInsertCapacity) Expansion(expansion int64) BfInsertExpansion {
	_ = "STUB: not implemented"
	return *new(BfInsertExpansion)
}

func (c BfInsertCapacity) Nocreate() BfInsertNocreate {
	_ = "STUB: not implemented"
	return *new(BfInsertNocreate)
}

func (c BfInsertCapacity) Nonscaling() BfInsertNonscaling {
	_ = "STUB: not implemented"
	return *new(BfInsertNonscaling)
}

func (c BfInsertCapacity) Items() BfInsertItems {
	_ = "STUB: not implemented"
	return *new(BfInsertItems)
}

type BfInsertError Incomplete

func (c BfInsertError) Expansion(expansion int64) BfInsertExpansion {
	_ = "STUB: not implemented"
	return *new(BfInsertExpansion)
}

func (c BfInsertError) Nocreate() BfInsertNocreate {
	_ = "STUB: not implemented"
	return *new(BfInsertNocreate)
}

func (c BfInsertError) Nonscaling() BfInsertNonscaling {
	_ = "STUB: not implemented"
	return *new(BfInsertNonscaling)
}

func (c BfInsertError) Items() BfInsertItems { _ = "STUB: not implemented"; return *new(BfInsertItems) }

type BfInsertExpansion Incomplete

func (c BfInsertExpansion) Nocreate() BfInsertNocreate {
	_ = "STUB: not implemented"
	return *new(BfInsertNocreate)
}

func (c BfInsertExpansion) Nonscaling() BfInsertNonscaling {
	_ = "STUB: not implemented"
	return *new(BfInsertNonscaling)
}

func (c BfInsertExpansion) Items() BfInsertItems {
	_ = "STUB: not implemented"
	return *new(BfInsertItems)
}

type BfInsertItem Incomplete

func (c BfInsertItem) Item(item ...string) BfInsertItem {
	_ = "STUB: not implemented"
	return *new(BfInsertItem)
}

func (c BfInsertItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BfInsertItems Incomplete

func (c BfInsertItems) Item(item ...string) BfInsertItem {
	_ = "STUB: not implemented"
	return *new(BfInsertItem)
}

type BfInsertKey Incomplete

func (c BfInsertKey) Capacity(capacity int64) BfInsertCapacity {
	_ = "STUB: not implemented"
	return *new(BfInsertCapacity)
}

func (c BfInsertKey) Error(error float64) BfInsertError {
	_ = "STUB: not implemented"
	return *new(BfInsertError)
}

func (c BfInsertKey) Expansion(expansion int64) BfInsertExpansion {
	_ = "STUB: not implemented"
	return *new(BfInsertExpansion)
}

func (c BfInsertKey) Nocreate() BfInsertNocreate {
	_ = "STUB: not implemented"
	return *new(BfInsertNocreate)
}

func (c BfInsertKey) Nonscaling() BfInsertNonscaling {
	_ = "STUB: not implemented"
	return *new(BfInsertNonscaling)
}

func (c BfInsertKey) Items() BfInsertItems { _ = "STUB: not implemented"; return *new(BfInsertItems) }

type BfInsertNocreate Incomplete

func (c BfInsertNocreate) Nonscaling() BfInsertNonscaling {
	_ = "STUB: not implemented"
	return *new(BfInsertNonscaling)
}

func (c BfInsertNocreate) Items() BfInsertItems {
	_ = "STUB: not implemented"
	return *new(BfInsertItems)
}

type BfInsertNonscaling Incomplete

func (c BfInsertNonscaling) Items() BfInsertItems {
	_ = "STUB: not implemented"
	return *new(BfInsertItems)
}

type BfLoadchunk Incomplete

func (b Builder) BfLoadchunk() (c BfLoadchunk) { _ = "STUB: not implemented"; return *new(BfLoadchunk) }

func (c BfLoadchunk) Key(key string) BfLoadchunkKey {
	_ = "STUB: not implemented"
	return *new(BfLoadchunkKey)
}

type BfLoadchunkData Incomplete

func (c BfLoadchunkData) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BfLoadchunkIterator Incomplete

func (c BfLoadchunkIterator) Data(data string) BfLoadchunkData {
	_ = "STUB: not implemented"
	return *new(BfLoadchunkData)
}

type BfLoadchunkKey Incomplete

func (c BfLoadchunkKey) Iterator(iterator int64) BfLoadchunkIterator {
	_ = "STUB: not implemented"
	return *new(BfLoadchunkIterator)
}

type BfMadd Incomplete

func (b Builder) BfMadd() (c BfMadd) { _ = "STUB: not implemented"; return *new(BfMadd) }

func (c BfMadd) Key(key string) BfMaddKey { _ = "STUB: not implemented"; return *new(BfMaddKey) }

type BfMaddItem Incomplete

func (c BfMaddItem) Item(item ...string) BfMaddItem {
	_ = "STUB: not implemented"
	return *new(BfMaddItem)
}

func (c BfMaddItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BfMaddKey Incomplete

func (c BfMaddKey) Item(item ...string) BfMaddItem {
	_ = "STUB: not implemented"
	return *new(BfMaddItem)
}

type BfMexists Incomplete

func (b Builder) BfMexists() (c BfMexists) { _ = "STUB: not implemented"; return *new(BfMexists) }

func (c BfMexists) Key(key string) BfMexistsKey {
	_ = "STUB: not implemented"
	return *new(BfMexistsKey)
}

type BfMexistsItem Incomplete

func (c BfMexistsItem) Item(item ...string) BfMexistsItem {
	_ = "STUB: not implemented"
	return *new(BfMexistsItem)
}

func (c BfMexistsItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BfMexistsKey Incomplete

func (c BfMexistsKey) Item(item ...string) BfMexistsItem {
	_ = "STUB: not implemented"
	return *new(BfMexistsItem)
}

type BfReserve Incomplete

func (b Builder) BfReserve() (c BfReserve) { _ = "STUB: not implemented"; return *new(BfReserve) }

func (c BfReserve) Key(key string) BfReserveKey {
	_ = "STUB: not implemented"
	return *new(BfReserveKey)
}

type BfReserveCapacity Incomplete

func (c BfReserveCapacity) Expansion(expansion int64) BfReserveExpansion {
	_ = "STUB: not implemented"
	return *new(BfReserveExpansion)
}

func (c BfReserveCapacity) Nonscaling() BfReserveNonscaling {
	_ = "STUB: not implemented"
	return *new(BfReserveNonscaling)
}

func (c BfReserveCapacity) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BfReserveErrorRate Incomplete

func (c BfReserveErrorRate) Capacity(capacity int64) BfReserveCapacity {
	_ = "STUB: not implemented"
	return *new(BfReserveCapacity)
}

type BfReserveExpansion Incomplete

func (c BfReserveExpansion) Nonscaling() BfReserveNonscaling {
	_ = "STUB: not implemented"
	return *new(BfReserveNonscaling)
}

func (c BfReserveExpansion) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BfReserveKey Incomplete

func (c BfReserveKey) ErrorRate(errorRate float64) BfReserveErrorRate {
	_ = "STUB: not implemented"
	return *new(BfReserveErrorRate)
}

type BfReserveNonscaling Incomplete

func (c BfReserveNonscaling) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BfScandump Incomplete

func (b Builder) BfScandump() (c BfScandump) { _ = "STUB: not implemented"; return *new(BfScandump) }

func (c BfScandump) Key(key string) BfScandumpKey {
	_ = "STUB: not implemented"
	return *new(BfScandumpKey)
}

type BfScandumpIterator Incomplete

func (c BfScandumpIterator) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BfScandumpKey Incomplete

func (c BfScandumpKey) Iterator(iterator int64) BfScandumpIterator {
	_ = "STUB: not implemented"
	return *new(BfScandumpIterator)
}
