// Code generated DO NOT EDIT

package cmds

type CfAdd Incomplete

func (b Builder) CfAdd() (c CfAdd) { _ = "STUB: not implemented"; return *new(CfAdd) }

func (c CfAdd) Key(key string) CfAddKey { _ = "STUB: not implemented"; return *new(CfAddKey) }

type CfAddItem Incomplete

func (c CfAddItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CfAddKey Incomplete

func (c CfAddKey) Item(item string) CfAddItem { _ = "STUB: not implemented"; return *new(CfAddItem) }

type CfAddnx Incomplete

func (b Builder) CfAddnx() (c CfAddnx) { _ = "STUB: not implemented"; return *new(CfAddnx) }

func (c CfAddnx) Key(key string) CfAddnxKey { _ = "STUB: not implemented"; return *new(CfAddnxKey) }

type CfAddnxItem Incomplete

func (c CfAddnxItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CfAddnxKey Incomplete

func (c CfAddnxKey) Item(item string) CfAddnxItem {
	_ = "STUB: not implemented"
	return *new(CfAddnxItem)
}

type CfCount Incomplete

func (b Builder) CfCount() (c CfCount) { _ = "STUB: not implemented"; return *new(CfCount) }

func (c CfCount) Key(key string) CfCountKey { _ = "STUB: not implemented"; return *new(CfCountKey) }

type CfCountItem Incomplete

func (c CfCountItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c CfCountItem) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type CfCountKey Incomplete

func (c CfCountKey) Item(item string) CfCountItem {
	_ = "STUB: not implemented"
	return *new(CfCountItem)
}

type CfDel Incomplete

func (b Builder) CfDel() (c CfDel) { _ = "STUB: not implemented"; return *new(CfDel) }

func (c CfDel) Key(key string) CfDelKey { _ = "STUB: not implemented"; return *new(CfDelKey) }

type CfDelItem Incomplete

func (c CfDelItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CfDelKey Incomplete

func (c CfDelKey) Item(item string) CfDelItem { _ = "STUB: not implemented"; return *new(CfDelItem) }

type CfExists Incomplete

func (b Builder) CfExists() (c CfExists) { _ = "STUB: not implemented"; return *new(CfExists) }

func (c CfExists) Key(key string) CfExistsKey { _ = "STUB: not implemented"; return *new(CfExistsKey) }

type CfExistsItem Incomplete

func (c CfExistsItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c CfExistsItem) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type CfExistsKey Incomplete

func (c CfExistsKey) Item(item string) CfExistsItem {
	_ = "STUB: not implemented"
	return *new(CfExistsItem)
}

type CfInfo Incomplete

func (b Builder) CfInfo() (c CfInfo) { _ = "STUB: not implemented"; return *new(CfInfo) }

func (c CfInfo) Key(key string) CfInfoKey { _ = "STUB: not implemented"; return *new(CfInfoKey) }

type CfInfoKey Incomplete

func (c CfInfoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c CfInfoKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type CfInsert Incomplete

func (b Builder) CfInsert() (c CfInsert) { _ = "STUB: not implemented"; return *new(CfInsert) }

func (c CfInsert) Key(key string) CfInsertKey { _ = "STUB: not implemented"; return *new(CfInsertKey) }

type CfInsertCapacity Incomplete

func (c CfInsertCapacity) Nocreate() CfInsertNocreate {
	_ = "STUB: not implemented"
	return *new(CfInsertNocreate)
}

func (c CfInsertCapacity) Items() CfInsertItems {
	_ = "STUB: not implemented"
	return *new(CfInsertItems)
}

type CfInsertItem Incomplete

func (c CfInsertItem) Item(item ...string) CfInsertItem {
	_ = "STUB: not implemented"
	return *new(CfInsertItem)
}

func (c CfInsertItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CfInsertItems Incomplete

func (c CfInsertItems) Item(item ...string) CfInsertItem {
	_ = "STUB: not implemented"
	return *new(CfInsertItem)
}

type CfInsertKey Incomplete

func (c CfInsertKey) Capacity(capacity int64) CfInsertCapacity {
	_ = "STUB: not implemented"
	return *new(CfInsertCapacity)
}

func (c CfInsertKey) Nocreate() CfInsertNocreate {
	_ = "STUB: not implemented"
	return *new(CfInsertNocreate)
}

func (c CfInsertKey) Items() CfInsertItems { _ = "STUB: not implemented"; return *new(CfInsertItems) }

type CfInsertNocreate Incomplete

func (c CfInsertNocreate) Items() CfInsertItems {
	_ = "STUB: not implemented"
	return *new(CfInsertItems)
}

type CfInsertnx Incomplete

func (b Builder) CfInsertnx() (c CfInsertnx) { _ = "STUB: not implemented"; return *new(CfInsertnx) }

func (c CfInsertnx) Key(key string) CfInsertnxKey {
	_ = "STUB: not implemented"
	return *new(CfInsertnxKey)
}

type CfInsertnxCapacity Incomplete

func (c CfInsertnxCapacity) Nocreate() CfInsertnxNocreate {
	_ = "STUB: not implemented"
	return *new(CfInsertnxNocreate)
}

func (c CfInsertnxCapacity) Items() CfInsertnxItems {
	_ = "STUB: not implemented"
	return *new(CfInsertnxItems)
}

type CfInsertnxItem Incomplete

func (c CfInsertnxItem) Item(item ...string) CfInsertnxItem {
	_ = "STUB: not implemented"
	return *new(CfInsertnxItem)
}

func (c CfInsertnxItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CfInsertnxItems Incomplete

func (c CfInsertnxItems) Item(item ...string) CfInsertnxItem {
	_ = "STUB: not implemented"
	return *new(CfInsertnxItem)
}

type CfInsertnxKey Incomplete

func (c CfInsertnxKey) Capacity(capacity int64) CfInsertnxCapacity {
	_ = "STUB: not implemented"
	return *new(CfInsertnxCapacity)
}

func (c CfInsertnxKey) Nocreate() CfInsertnxNocreate {
	_ = "STUB: not implemented"
	return *new(CfInsertnxNocreate)
}

func (c CfInsertnxKey) Items() CfInsertnxItems {
	_ = "STUB: not implemented"
	return *new(CfInsertnxItems)
}

type CfInsertnxNocreate Incomplete

func (c CfInsertnxNocreate) Items() CfInsertnxItems {
	_ = "STUB: not implemented"
	return *new(CfInsertnxItems)
}

type CfLoadchunk Incomplete

func (b Builder) CfLoadchunk() (c CfLoadchunk) { _ = "STUB: not implemented"; return *new(CfLoadchunk) }

func (c CfLoadchunk) Key(key string) CfLoadchunkKey {
	_ = "STUB: not implemented"
	return *new(CfLoadchunkKey)
}

type CfLoadchunkData Incomplete

func (c CfLoadchunkData) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CfLoadchunkIterator Incomplete

func (c CfLoadchunkIterator) Data(data string) CfLoadchunkData {
	_ = "STUB: not implemented"
	return *new(CfLoadchunkData)
}

type CfLoadchunkKey Incomplete

func (c CfLoadchunkKey) Iterator(iterator int64) CfLoadchunkIterator {
	_ = "STUB: not implemented"
	return *new(CfLoadchunkIterator)
}

type CfMexists Incomplete

func (b Builder) CfMexists() (c CfMexists) { _ = "STUB: not implemented"; return *new(CfMexists) }

func (c CfMexists) Key(key string) CfMexistsKey {
	_ = "STUB: not implemented"
	return *new(CfMexistsKey)
}

type CfMexistsItem Incomplete

func (c CfMexistsItem) Item(item ...string) CfMexistsItem {
	_ = "STUB: not implemented"
	return *new(CfMexistsItem)
}

func (c CfMexistsItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CfMexistsKey Incomplete

func (c CfMexistsKey) Item(item ...string) CfMexistsItem {
	_ = "STUB: not implemented"
	return *new(CfMexistsItem)
}

type CfReserve Incomplete

func (b Builder) CfReserve() (c CfReserve) { _ = "STUB: not implemented"; return *new(CfReserve) }

func (c CfReserve) Key(key string) CfReserveKey {
	_ = "STUB: not implemented"
	return *new(CfReserveKey)
}

type CfReserveBucketsize Incomplete

func (c CfReserveBucketsize) Maxiterations(maxiterations int64) CfReserveMaxiterations {
	_ = "STUB: not implemented"
	return *new(CfReserveMaxiterations)
}

func (c CfReserveBucketsize) Expansion(expansion int64) CfReserveExpansion {
	_ = "STUB: not implemented"
	return *new(CfReserveExpansion)
}

func (c CfReserveBucketsize) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CfReserveCapacity Incomplete

func (c CfReserveCapacity) Bucketsize(bucketsize int64) CfReserveBucketsize {
	_ = "STUB: not implemented"
	return *new(CfReserveBucketsize)
}

func (c CfReserveCapacity) Maxiterations(maxiterations int64) CfReserveMaxiterations {
	_ = "STUB: not implemented"
	return *new(CfReserveMaxiterations)
}

func (c CfReserveCapacity) Expansion(expansion int64) CfReserveExpansion {
	_ = "STUB: not implemented"
	return *new(CfReserveExpansion)
}

func (c CfReserveCapacity) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CfReserveExpansion Incomplete

func (c CfReserveExpansion) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CfReserveKey Incomplete

func (c CfReserveKey) Capacity(capacity int64) CfReserveCapacity {
	_ = "STUB: not implemented"
	return *new(CfReserveCapacity)
}

type CfReserveMaxiterations Incomplete

func (c CfReserveMaxiterations) Expansion(expansion int64) CfReserveExpansion {
	_ = "STUB: not implemented"
	return *new(CfReserveExpansion)
}

func (c CfReserveMaxiterations) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type CfScandump Incomplete

func (b Builder) CfScandump() (c CfScandump) { _ = "STUB: not implemented"; return *new(CfScandump) }

func (c CfScandump) Key(key string) CfScandumpKey {
	_ = "STUB: not implemented"
	return *new(CfScandumpKey)
}

type CfScandumpIterator Incomplete

func (c CfScandumpIterator) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CfScandumpKey Incomplete

func (c CfScandumpKey) Iterator(iterator int64) CfScandumpIterator {
	_ = "STUB: not implemented"
	return *new(CfScandumpIterator)
}
