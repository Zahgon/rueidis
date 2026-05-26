// Code generated DO NOT EDIT

package cmds

type TsAdd Incomplete

func (b Builder) TsAdd() (c TsAdd) { _ = "STUB: not implemented"; return *new(TsAdd) }

func (c TsAdd) Key(key string) TsAddKey { _ = "STUB: not implemented"; return *new(TsAddKey) }

type TsAddChunkSize Incomplete

func (c TsAddChunkSize) OnDuplicateBlock() TsAddOnDuplicateBlock {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateBlock)
}

func (c TsAddChunkSize) OnDuplicateFirst() TsAddOnDuplicateFirst {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateFirst)
}

func (c TsAddChunkSize) OnDuplicateLast() TsAddOnDuplicateLast {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateLast)
}

func (c TsAddChunkSize) OnDuplicateMin() TsAddOnDuplicateMin {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateMin)
}

func (c TsAddChunkSize) OnDuplicateMax() TsAddOnDuplicateMax {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateMax)
}

func (c TsAddChunkSize) OnDuplicateSum() TsAddOnDuplicateSum {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateSum)
}

func (c TsAddChunkSize) Labels() TsAddLabels { _ = "STUB: not implemented"; return *new(TsAddLabels) }

func (c TsAddChunkSize) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAddEncodingCompressed Incomplete

func (c TsAddEncodingCompressed) ChunkSize(size int64) TsAddChunkSize {
	_ = "STUB: not implemented"
	return *new(TsAddChunkSize)
}

func (c TsAddEncodingCompressed) OnDuplicateBlock() TsAddOnDuplicateBlock {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateBlock)
}

func (c TsAddEncodingCompressed) OnDuplicateFirst() TsAddOnDuplicateFirst {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateFirst)
}

func (c TsAddEncodingCompressed) OnDuplicateLast() TsAddOnDuplicateLast {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateLast)
}

func (c TsAddEncodingCompressed) OnDuplicateMin() TsAddOnDuplicateMin {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateMin)
}

func (c TsAddEncodingCompressed) OnDuplicateMax() TsAddOnDuplicateMax {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateMax)
}

func (c TsAddEncodingCompressed) OnDuplicateSum() TsAddOnDuplicateSum {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateSum)
}

func (c TsAddEncodingCompressed) Labels() TsAddLabels {
	_ = "STUB: not implemented"
	return *new(TsAddLabels)
}

func (c TsAddEncodingCompressed) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsAddEncodingUncompressed Incomplete

func (c TsAddEncodingUncompressed) ChunkSize(size int64) TsAddChunkSize {
	_ = "STUB: not implemented"
	return *new(TsAddChunkSize)
}

func (c TsAddEncodingUncompressed) OnDuplicateBlock() TsAddOnDuplicateBlock {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateBlock)
}

func (c TsAddEncodingUncompressed) OnDuplicateFirst() TsAddOnDuplicateFirst {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateFirst)
}

func (c TsAddEncodingUncompressed) OnDuplicateLast() TsAddOnDuplicateLast {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateLast)
}

func (c TsAddEncodingUncompressed) OnDuplicateMin() TsAddOnDuplicateMin {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateMin)
}

func (c TsAddEncodingUncompressed) OnDuplicateMax() TsAddOnDuplicateMax {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateMax)
}

func (c TsAddEncodingUncompressed) OnDuplicateSum() TsAddOnDuplicateSum {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateSum)
}

func (c TsAddEncodingUncompressed) Labels() TsAddLabels {
	_ = "STUB: not implemented"
	return *new(TsAddLabels)
}

func (c TsAddEncodingUncompressed) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsAddKey Incomplete

func (c TsAddKey) Timestamp(timestamp string) TsAddTimestamp {
	_ = "STUB: not implemented"
	return *new(TsAddTimestamp)
}

type TsAddLabels Incomplete

func (c TsAddLabels) Labels(label string, value string) TsAddLabels {
	_ = "STUB: not implemented"
	return *new(TsAddLabels)
}

func (c TsAddLabels) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAddOnDuplicateBlock Incomplete

func (c TsAddOnDuplicateBlock) Labels() TsAddLabels {
	_ = "STUB: not implemented"
	return *new(TsAddLabels)
}

func (c TsAddOnDuplicateBlock) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAddOnDuplicateFirst Incomplete

func (c TsAddOnDuplicateFirst) Labels() TsAddLabels {
	_ = "STUB: not implemented"
	return *new(TsAddLabels)
}

func (c TsAddOnDuplicateFirst) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAddOnDuplicateLast Incomplete

func (c TsAddOnDuplicateLast) Labels() TsAddLabels {
	_ = "STUB: not implemented"
	return *new(TsAddLabels)
}

func (c TsAddOnDuplicateLast) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAddOnDuplicateMax Incomplete

func (c TsAddOnDuplicateMax) Labels() TsAddLabels {
	_ = "STUB: not implemented"
	return *new(TsAddLabels)
}

func (c TsAddOnDuplicateMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAddOnDuplicateMin Incomplete

func (c TsAddOnDuplicateMin) Labels() TsAddLabels {
	_ = "STUB: not implemented"
	return *new(TsAddLabels)
}

func (c TsAddOnDuplicateMin) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAddOnDuplicateSum Incomplete

func (c TsAddOnDuplicateSum) Labels() TsAddLabels {
	_ = "STUB: not implemented"
	return *new(TsAddLabels)
}

func (c TsAddOnDuplicateSum) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAddRetention Incomplete

func (c TsAddRetention) EncodingUncompressed() TsAddEncodingUncompressed {
	_ = "STUB: not implemented"
	return *new(TsAddEncodingUncompressed)
}

func (c TsAddRetention) EncodingCompressed() TsAddEncodingCompressed {
	_ = "STUB: not implemented"
	return *new(TsAddEncodingCompressed)
}

func (c TsAddRetention) ChunkSize(size int64) TsAddChunkSize {
	_ = "STUB: not implemented"
	return *new(TsAddChunkSize)
}

func (c TsAddRetention) OnDuplicateBlock() TsAddOnDuplicateBlock {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateBlock)
}

func (c TsAddRetention) OnDuplicateFirst() TsAddOnDuplicateFirst {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateFirst)
}

func (c TsAddRetention) OnDuplicateLast() TsAddOnDuplicateLast {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateLast)
}

func (c TsAddRetention) OnDuplicateMin() TsAddOnDuplicateMin {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateMin)
}

func (c TsAddRetention) OnDuplicateMax() TsAddOnDuplicateMax {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateMax)
}

func (c TsAddRetention) OnDuplicateSum() TsAddOnDuplicateSum {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateSum)
}

func (c TsAddRetention) Labels() TsAddLabels { _ = "STUB: not implemented"; return *new(TsAddLabels) }

func (c TsAddRetention) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAddTimestamp Incomplete

func (c TsAddTimestamp) Value(value float64) TsAddValue {
	_ = "STUB: not implemented"
	return *new(TsAddValue)
}

type TsAddValue Incomplete

func (c TsAddValue) Retention(retentionperiod int64) TsAddRetention {
	_ = "STUB: not implemented"
	return *new(TsAddRetention)
}

func (c TsAddValue) EncodingUncompressed() TsAddEncodingUncompressed {
	_ = "STUB: not implemented"
	return *new(TsAddEncodingUncompressed)
}

func (c TsAddValue) EncodingCompressed() TsAddEncodingCompressed {
	_ = "STUB: not implemented"
	return *new(TsAddEncodingCompressed)
}

func (c TsAddValue) ChunkSize(size int64) TsAddChunkSize {
	_ = "STUB: not implemented"
	return *new(TsAddChunkSize)
}

func (c TsAddValue) OnDuplicateBlock() TsAddOnDuplicateBlock {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateBlock)
}

func (c TsAddValue) OnDuplicateFirst() TsAddOnDuplicateFirst {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateFirst)
}

func (c TsAddValue) OnDuplicateLast() TsAddOnDuplicateLast {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateLast)
}

func (c TsAddValue) OnDuplicateMin() TsAddOnDuplicateMin {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateMin)
}

func (c TsAddValue) OnDuplicateMax() TsAddOnDuplicateMax {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateMax)
}

func (c TsAddValue) OnDuplicateSum() TsAddOnDuplicateSum {
	_ = "STUB: not implemented"
	return *new(TsAddOnDuplicateSum)
}

func (c TsAddValue) Labels() TsAddLabels { _ = "STUB: not implemented"; return *new(TsAddLabels) }

func (c TsAddValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAlter Incomplete

func (b Builder) TsAlter() (c TsAlter) { _ = "STUB: not implemented"; return *new(TsAlter) }

func (c TsAlter) Key(key string) TsAlterKey { _ = "STUB: not implemented"; return *new(TsAlterKey) }

type TsAlterChunkSize Incomplete

func (c TsAlterChunkSize) DuplicatePolicyBlock() TsAlterDuplicatePolicyBlock {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyBlock)
}

func (c TsAlterChunkSize) DuplicatePolicyFirst() TsAlterDuplicatePolicyFirst {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyFirst)
}

func (c TsAlterChunkSize) DuplicatePolicyLast() TsAlterDuplicatePolicyLast {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyLast)
}

func (c TsAlterChunkSize) DuplicatePolicyMin() TsAlterDuplicatePolicyMin {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyMin)
}

func (c TsAlterChunkSize) DuplicatePolicyMax() TsAlterDuplicatePolicyMax {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyMax)
}

func (c TsAlterChunkSize) DuplicatePolicySum() TsAlterDuplicatePolicySum {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicySum)
}

func (c TsAlterChunkSize) Labels() TsAlterLabels {
	_ = "STUB: not implemented"
	return *new(TsAlterLabels)
}

func (c TsAlterChunkSize) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAlterDuplicatePolicyBlock Incomplete

func (c TsAlterDuplicatePolicyBlock) Labels() TsAlterLabels {
	_ = "STUB: not implemented"
	return *new(TsAlterLabels)
}

func (c TsAlterDuplicatePolicyBlock) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsAlterDuplicatePolicyFirst Incomplete

func (c TsAlterDuplicatePolicyFirst) Labels() TsAlterLabels {
	_ = "STUB: not implemented"
	return *new(TsAlterLabels)
}

func (c TsAlterDuplicatePolicyFirst) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsAlterDuplicatePolicyLast Incomplete

func (c TsAlterDuplicatePolicyLast) Labels() TsAlterLabels {
	_ = "STUB: not implemented"
	return *new(TsAlterLabels)
}

func (c TsAlterDuplicatePolicyLast) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsAlterDuplicatePolicyMax Incomplete

func (c TsAlterDuplicatePolicyMax) Labels() TsAlterLabels {
	_ = "STUB: not implemented"
	return *new(TsAlterLabels)
}

func (c TsAlterDuplicatePolicyMax) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsAlterDuplicatePolicyMin Incomplete

func (c TsAlterDuplicatePolicyMin) Labels() TsAlterLabels {
	_ = "STUB: not implemented"
	return *new(TsAlterLabels)
}

func (c TsAlterDuplicatePolicyMin) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsAlterDuplicatePolicySum Incomplete

func (c TsAlterDuplicatePolicySum) Labels() TsAlterLabels {
	_ = "STUB: not implemented"
	return *new(TsAlterLabels)
}

func (c TsAlterDuplicatePolicySum) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsAlterKey Incomplete

func (c TsAlterKey) Retention(retentionperiod int64) TsAlterRetention {
	_ = "STUB: not implemented"
	return *new(TsAlterRetention)
}

func (c TsAlterKey) ChunkSize(size int64) TsAlterChunkSize {
	_ = "STUB: not implemented"
	return *new(TsAlterChunkSize)
}

func (c TsAlterKey) DuplicatePolicyBlock() TsAlterDuplicatePolicyBlock {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyBlock)
}

func (c TsAlterKey) DuplicatePolicyFirst() TsAlterDuplicatePolicyFirst {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyFirst)
}

func (c TsAlterKey) DuplicatePolicyLast() TsAlterDuplicatePolicyLast {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyLast)
}

func (c TsAlterKey) DuplicatePolicyMin() TsAlterDuplicatePolicyMin {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyMin)
}

func (c TsAlterKey) DuplicatePolicyMax() TsAlterDuplicatePolicyMax {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyMax)
}

func (c TsAlterKey) DuplicatePolicySum() TsAlterDuplicatePolicySum {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicySum)
}

func (c TsAlterKey) Labels() TsAlterLabels { _ = "STUB: not implemented"; return *new(TsAlterLabels) }

func (c TsAlterKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAlterLabels Incomplete

func (c TsAlterLabels) Labels(label string, value string) TsAlterLabels {
	_ = "STUB: not implemented"
	return *new(TsAlterLabels)
}

func (c TsAlterLabels) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsAlterRetention Incomplete

func (c TsAlterRetention) ChunkSize(size int64) TsAlterChunkSize {
	_ = "STUB: not implemented"
	return *new(TsAlterChunkSize)
}

func (c TsAlterRetention) DuplicatePolicyBlock() TsAlterDuplicatePolicyBlock {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyBlock)
}

func (c TsAlterRetention) DuplicatePolicyFirst() TsAlterDuplicatePolicyFirst {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyFirst)
}

func (c TsAlterRetention) DuplicatePolicyLast() TsAlterDuplicatePolicyLast {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyLast)
}

func (c TsAlterRetention) DuplicatePolicyMin() TsAlterDuplicatePolicyMin {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyMin)
}

func (c TsAlterRetention) DuplicatePolicyMax() TsAlterDuplicatePolicyMax {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicyMax)
}

func (c TsAlterRetention) DuplicatePolicySum() TsAlterDuplicatePolicySum {
	_ = "STUB: not implemented"
	return *new(TsAlterDuplicatePolicySum)
}

func (c TsAlterRetention) Labels() TsAlterLabels {
	_ = "STUB: not implemented"
	return *new(TsAlterLabels)
}

func (c TsAlterRetention) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsCreate Incomplete

func (b Builder) TsCreate() (c TsCreate) { _ = "STUB: not implemented"; return *new(TsCreate) }

func (c TsCreate) Key(key string) TsCreateKey { _ = "STUB: not implemented"; return *new(TsCreateKey) }

type TsCreateChunkSize Incomplete

func (c TsCreateChunkSize) DuplicatePolicyBlock() TsCreateDuplicatePolicyBlock {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyBlock)
}

func (c TsCreateChunkSize) DuplicatePolicyFirst() TsCreateDuplicatePolicyFirst {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyFirst)
}

func (c TsCreateChunkSize) DuplicatePolicyLast() TsCreateDuplicatePolicyLast {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyLast)
}

func (c TsCreateChunkSize) DuplicatePolicyMin() TsCreateDuplicatePolicyMin {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyMin)
}

func (c TsCreateChunkSize) DuplicatePolicyMax() TsCreateDuplicatePolicyMax {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyMax)
}

func (c TsCreateChunkSize) DuplicatePolicySum() TsCreateDuplicatePolicySum {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicySum)
}

func (c TsCreateChunkSize) Labels() TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateChunkSize) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsCreateDuplicatePolicyBlock Incomplete

func (c TsCreateDuplicatePolicyBlock) Labels() TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateDuplicatePolicyBlock) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsCreateDuplicatePolicyFirst Incomplete

func (c TsCreateDuplicatePolicyFirst) Labels() TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateDuplicatePolicyFirst) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsCreateDuplicatePolicyLast Incomplete

func (c TsCreateDuplicatePolicyLast) Labels() TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateDuplicatePolicyLast) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsCreateDuplicatePolicyMax Incomplete

func (c TsCreateDuplicatePolicyMax) Labels() TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateDuplicatePolicyMax) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsCreateDuplicatePolicyMin Incomplete

func (c TsCreateDuplicatePolicyMin) Labels() TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateDuplicatePolicyMin) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsCreateDuplicatePolicySum Incomplete

func (c TsCreateDuplicatePolicySum) Labels() TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateDuplicatePolicySum) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsCreateEncodingCompressed Incomplete

func (c TsCreateEncodingCompressed) ChunkSize(size int64) TsCreateChunkSize {
	_ = "STUB: not implemented"
	return *new(TsCreateChunkSize)
}

func (c TsCreateEncodingCompressed) DuplicatePolicyBlock() TsCreateDuplicatePolicyBlock {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyBlock)
}

func (c TsCreateEncodingCompressed) DuplicatePolicyFirst() TsCreateDuplicatePolicyFirst {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyFirst)
}

func (c TsCreateEncodingCompressed) DuplicatePolicyLast() TsCreateDuplicatePolicyLast {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyLast)
}

func (c TsCreateEncodingCompressed) DuplicatePolicyMin() TsCreateDuplicatePolicyMin {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyMin)
}

func (c TsCreateEncodingCompressed) DuplicatePolicyMax() TsCreateDuplicatePolicyMax {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyMax)
}

func (c TsCreateEncodingCompressed) DuplicatePolicySum() TsCreateDuplicatePolicySum {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicySum)
}

func (c TsCreateEncodingCompressed) Labels() TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateEncodingCompressed) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsCreateEncodingUncompressed Incomplete

func (c TsCreateEncodingUncompressed) ChunkSize(size int64) TsCreateChunkSize {
	_ = "STUB: not implemented"
	return *new(TsCreateChunkSize)
}

func (c TsCreateEncodingUncompressed) DuplicatePolicyBlock() TsCreateDuplicatePolicyBlock {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyBlock)
}

func (c TsCreateEncodingUncompressed) DuplicatePolicyFirst() TsCreateDuplicatePolicyFirst {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyFirst)
}

func (c TsCreateEncodingUncompressed) DuplicatePolicyLast() TsCreateDuplicatePolicyLast {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyLast)
}

func (c TsCreateEncodingUncompressed) DuplicatePolicyMin() TsCreateDuplicatePolicyMin {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyMin)
}

func (c TsCreateEncodingUncompressed) DuplicatePolicyMax() TsCreateDuplicatePolicyMax {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyMax)
}

func (c TsCreateEncodingUncompressed) DuplicatePolicySum() TsCreateDuplicatePolicySum {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicySum)
}

func (c TsCreateEncodingUncompressed) Labels() TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateEncodingUncompressed) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsCreateKey Incomplete

func (c TsCreateKey) Retention(retentionperiod int64) TsCreateRetention {
	_ = "STUB: not implemented"
	return *new(TsCreateRetention)
}

func (c TsCreateKey) EncodingUncompressed() TsCreateEncodingUncompressed {
	_ = "STUB: not implemented"
	return *new(TsCreateEncodingUncompressed)
}

func (c TsCreateKey) EncodingCompressed() TsCreateEncodingCompressed {
	_ = "STUB: not implemented"
	return *new(TsCreateEncodingCompressed)
}

func (c TsCreateKey) ChunkSize(size int64) TsCreateChunkSize {
	_ = "STUB: not implemented"
	return *new(TsCreateChunkSize)
}

func (c TsCreateKey) DuplicatePolicyBlock() TsCreateDuplicatePolicyBlock {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyBlock)
}

func (c TsCreateKey) DuplicatePolicyFirst() TsCreateDuplicatePolicyFirst {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyFirst)
}

func (c TsCreateKey) DuplicatePolicyLast() TsCreateDuplicatePolicyLast {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyLast)
}

func (c TsCreateKey) DuplicatePolicyMin() TsCreateDuplicatePolicyMin {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyMin)
}

func (c TsCreateKey) DuplicatePolicyMax() TsCreateDuplicatePolicyMax {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyMax)
}

func (c TsCreateKey) DuplicatePolicySum() TsCreateDuplicatePolicySum {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicySum)
}

func (c TsCreateKey) Labels() TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsCreateLabels Incomplete

func (c TsCreateLabels) Labels(label string, value string) TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateLabels) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsCreateRetention Incomplete

func (c TsCreateRetention) EncodingUncompressed() TsCreateEncodingUncompressed {
	_ = "STUB: not implemented"
	return *new(TsCreateEncodingUncompressed)
}

func (c TsCreateRetention) EncodingCompressed() TsCreateEncodingCompressed {
	_ = "STUB: not implemented"
	return *new(TsCreateEncodingCompressed)
}

func (c TsCreateRetention) ChunkSize(size int64) TsCreateChunkSize {
	_ = "STUB: not implemented"
	return *new(TsCreateChunkSize)
}

func (c TsCreateRetention) DuplicatePolicyBlock() TsCreateDuplicatePolicyBlock {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyBlock)
}

func (c TsCreateRetention) DuplicatePolicyFirst() TsCreateDuplicatePolicyFirst {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyFirst)
}

func (c TsCreateRetention) DuplicatePolicyLast() TsCreateDuplicatePolicyLast {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyLast)
}

func (c TsCreateRetention) DuplicatePolicyMin() TsCreateDuplicatePolicyMin {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyMin)
}

func (c TsCreateRetention) DuplicatePolicyMax() TsCreateDuplicatePolicyMax {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicyMax)
}

func (c TsCreateRetention) DuplicatePolicySum() TsCreateDuplicatePolicySum {
	_ = "STUB: not implemented"
	return *new(TsCreateDuplicatePolicySum)
}

func (c TsCreateRetention) Labels() TsCreateLabels {
	_ = "STUB: not implemented"
	return *new(TsCreateLabels)
}

func (c TsCreateRetention) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsCreaterule Incomplete

func (b Builder) TsCreaterule() (c TsCreaterule) {
	_ = "STUB: not implemented"
	return *new(TsCreaterule)
}

func (c TsCreaterule) Sourcekey(sourcekey string) TsCreateruleSourcekey {
	_ = "STUB: not implemented"
	return *new(TsCreateruleSourcekey)
}

type TsCreateruleAggregationAvg Incomplete

func (c TsCreateruleAggregationAvg) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationCount Incomplete

func (c TsCreateruleAggregationCount) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationFirst Incomplete

func (c TsCreateruleAggregationFirst) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationLast Incomplete

func (c TsCreateruleAggregationLast) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationMax Incomplete

func (c TsCreateruleAggregationMax) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationMin Incomplete

func (c TsCreateruleAggregationMin) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationRange Incomplete

func (c TsCreateruleAggregationRange) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationStdP Incomplete

func (c TsCreateruleAggregationStdP) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationStdS Incomplete

func (c TsCreateruleAggregationStdS) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationSum Incomplete

func (c TsCreateruleAggregationSum) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationTwa Incomplete

func (c TsCreateruleAggregationTwa) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationVarP Incomplete

func (c TsCreateruleAggregationVarP) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAggregationVarS Incomplete

func (c TsCreateruleAggregationVarS) Bucketduration(bucketduration int64) TsCreateruleBucketduration {
	_ = "STUB: not implemented"
	return *new(TsCreateruleBucketduration)
}

type TsCreateruleAligntimestamp Incomplete

func (c TsCreateruleAligntimestamp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsCreateruleBucketduration Incomplete

func (c TsCreateruleBucketduration) Aligntimestamp(aligntimestamp int64) TsCreateruleAligntimestamp {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAligntimestamp)
}

func (c TsCreateruleBucketduration) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsCreateruleDestkey Incomplete

func (c TsCreateruleDestkey) AggregationAvg() TsCreateruleAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationAvg)
}

func (c TsCreateruleDestkey) AggregationSum() TsCreateruleAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationSum)
}

func (c TsCreateruleDestkey) AggregationMin() TsCreateruleAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationMin)
}

func (c TsCreateruleDestkey) AggregationMax() TsCreateruleAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationMax)
}

func (c TsCreateruleDestkey) AggregationRange() TsCreateruleAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationRange)
}

func (c TsCreateruleDestkey) AggregationCount() TsCreateruleAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationCount)
}

func (c TsCreateruleDestkey) AggregationFirst() TsCreateruleAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationFirst)
}

func (c TsCreateruleDestkey) AggregationLast() TsCreateruleAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationLast)
}

func (c TsCreateruleDestkey) AggregationStdP() TsCreateruleAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationStdP)
}

func (c TsCreateruleDestkey) AggregationStdS() TsCreateruleAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationStdS)
}

func (c TsCreateruleDestkey) AggregationVarP() TsCreateruleAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationVarP)
}

func (c TsCreateruleDestkey) AggregationVarS() TsCreateruleAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationVarS)
}

func (c TsCreateruleDestkey) AggregationTwa() TsCreateruleAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsCreateruleAggregationTwa)
}

type TsCreateruleSourcekey Incomplete

func (c TsCreateruleSourcekey) Destkey(destkey string) TsCreateruleDestkey {
	_ = "STUB: not implemented"
	return *new(TsCreateruleDestkey)
}

type TsDecrby Incomplete

func (b Builder) TsDecrby() (c TsDecrby) { _ = "STUB: not implemented"; return *new(TsDecrby) }

func (c TsDecrby) Key(key string) TsDecrbyKey { _ = "STUB: not implemented"; return *new(TsDecrbyKey) }

type TsDecrbyChunkSize Incomplete

func (c TsDecrbyChunkSize) Labels() TsDecrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsDecrbyLabels)
}

func (c TsDecrbyChunkSize) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsDecrbyKey Incomplete

func (c TsDecrbyKey) Value(value float64) TsDecrbyValue {
	_ = "STUB: not implemented"
	return *new(TsDecrbyValue)
}

type TsDecrbyLabels Incomplete

func (c TsDecrbyLabels) Labels(label string, value string) TsDecrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsDecrbyLabels)
}

func (c TsDecrbyLabels) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsDecrbyRetention Incomplete

func (c TsDecrbyRetention) Uncompressed() TsDecrbyUncompressed {
	_ = "STUB: not implemented"
	return *new(TsDecrbyUncompressed)
}

func (c TsDecrbyRetention) ChunkSize(size int64) TsDecrbyChunkSize {
	_ = "STUB: not implemented"
	return *new(TsDecrbyChunkSize)
}

func (c TsDecrbyRetention) Labels() TsDecrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsDecrbyLabels)
}

func (c TsDecrbyRetention) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsDecrbyTimestamp Incomplete

func (c TsDecrbyTimestamp) Retention(retentionperiod int64) TsDecrbyRetention {
	_ = "STUB: not implemented"
	return *new(TsDecrbyRetention)
}

func (c TsDecrbyTimestamp) Uncompressed() TsDecrbyUncompressed {
	_ = "STUB: not implemented"
	return *new(TsDecrbyUncompressed)
}

func (c TsDecrbyTimestamp) ChunkSize(size int64) TsDecrbyChunkSize {
	_ = "STUB: not implemented"
	return *new(TsDecrbyChunkSize)
}

func (c TsDecrbyTimestamp) Labels() TsDecrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsDecrbyLabels)
}

func (c TsDecrbyTimestamp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsDecrbyUncompressed Incomplete

func (c TsDecrbyUncompressed) ChunkSize(size int64) TsDecrbyChunkSize {
	_ = "STUB: not implemented"
	return *new(TsDecrbyChunkSize)
}

func (c TsDecrbyUncompressed) Labels() TsDecrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsDecrbyLabels)
}

func (c TsDecrbyUncompressed) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsDecrbyValue Incomplete

func (c TsDecrbyValue) Timestamp(timestamp string) TsDecrbyTimestamp {
	_ = "STUB: not implemented"
	return *new(TsDecrbyTimestamp)
}

func (c TsDecrbyValue) Retention(retentionperiod int64) TsDecrbyRetention {
	_ = "STUB: not implemented"
	return *new(TsDecrbyRetention)
}

func (c TsDecrbyValue) Uncompressed() TsDecrbyUncompressed {
	_ = "STUB: not implemented"
	return *new(TsDecrbyUncompressed)
}

func (c TsDecrbyValue) ChunkSize(size int64) TsDecrbyChunkSize {
	_ = "STUB: not implemented"
	return *new(TsDecrbyChunkSize)
}

func (c TsDecrbyValue) Labels() TsDecrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsDecrbyLabels)
}

func (c TsDecrbyValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsDel Incomplete

func (b Builder) TsDel() (c TsDel) { _ = "STUB: not implemented"; return *new(TsDel) }

func (c TsDel) Key(key string) TsDelKey { _ = "STUB: not implemented"; return *new(TsDelKey) }

type TsDelFromTimestamp Incomplete

func (c TsDelFromTimestamp) ToTimestamp(toTimestamp int64) TsDelToTimestamp {
	_ = "STUB: not implemented"
	return *new(TsDelToTimestamp)
}

type TsDelKey Incomplete

func (c TsDelKey) FromTimestamp(fromTimestamp int64) TsDelFromTimestamp {
	_ = "STUB: not implemented"
	return *new(TsDelFromTimestamp)
}

type TsDelToTimestamp Incomplete

func (c TsDelToTimestamp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsDeleterule Incomplete

func (b Builder) TsDeleterule() (c TsDeleterule) {
	_ = "STUB: not implemented"
	return *new(TsDeleterule)
}

func (c TsDeleterule) Sourcekey(sourcekey string) TsDeleteruleSourcekey {
	_ = "STUB: not implemented"
	return *new(TsDeleteruleSourcekey)
}

type TsDeleteruleDestkey Incomplete

func (c TsDeleteruleDestkey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsDeleteruleSourcekey Incomplete

func (c TsDeleteruleSourcekey) Destkey(destkey string) TsDeleteruleDestkey {
	_ = "STUB: not implemented"
	return *new(TsDeleteruleDestkey)
}

type TsGet Incomplete

func (b Builder) TsGet() (c TsGet) { _ = "STUB: not implemented"; return *new(TsGet) }

func (c TsGet) Key(key string) TsGetKey { _ = "STUB: not implemented"; return *new(TsGetKey) }

type TsGetKey Incomplete

func (c TsGetKey) Latest() TsGetLatest { _ = "STUB: not implemented"; return *new(TsGetLatest) }

func (c TsGetKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsGetLatest Incomplete

func (c TsGetLatest) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsIncrby Incomplete

func (b Builder) TsIncrby() (c TsIncrby) { _ = "STUB: not implemented"; return *new(TsIncrby) }

func (c TsIncrby) Key(key string) TsIncrbyKey { _ = "STUB: not implemented"; return *new(TsIncrbyKey) }

type TsIncrbyChunkSize Incomplete

func (c TsIncrbyChunkSize) Labels() TsIncrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsIncrbyLabels)
}

func (c TsIncrbyChunkSize) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsIncrbyKey Incomplete

func (c TsIncrbyKey) Value(value float64) TsIncrbyValue {
	_ = "STUB: not implemented"
	return *new(TsIncrbyValue)
}

type TsIncrbyLabels Incomplete

func (c TsIncrbyLabels) Labels(label string, value string) TsIncrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsIncrbyLabels)
}

func (c TsIncrbyLabels) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsIncrbyRetention Incomplete

func (c TsIncrbyRetention) Uncompressed() TsIncrbyUncompressed {
	_ = "STUB: not implemented"
	return *new(TsIncrbyUncompressed)
}

func (c TsIncrbyRetention) ChunkSize(size int64) TsIncrbyChunkSize {
	_ = "STUB: not implemented"
	return *new(TsIncrbyChunkSize)
}

func (c TsIncrbyRetention) Labels() TsIncrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsIncrbyLabels)
}

func (c TsIncrbyRetention) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsIncrbyTimestamp Incomplete

func (c TsIncrbyTimestamp) Retention(retentionperiod int64) TsIncrbyRetention {
	_ = "STUB: not implemented"
	return *new(TsIncrbyRetention)
}

func (c TsIncrbyTimestamp) Uncompressed() TsIncrbyUncompressed {
	_ = "STUB: not implemented"
	return *new(TsIncrbyUncompressed)
}

func (c TsIncrbyTimestamp) ChunkSize(size int64) TsIncrbyChunkSize {
	_ = "STUB: not implemented"
	return *new(TsIncrbyChunkSize)
}

func (c TsIncrbyTimestamp) Labels() TsIncrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsIncrbyLabels)
}

func (c TsIncrbyTimestamp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsIncrbyUncompressed Incomplete

func (c TsIncrbyUncompressed) ChunkSize(size int64) TsIncrbyChunkSize {
	_ = "STUB: not implemented"
	return *new(TsIncrbyChunkSize)
}

func (c TsIncrbyUncompressed) Labels() TsIncrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsIncrbyLabels)
}

func (c TsIncrbyUncompressed) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsIncrbyValue Incomplete

func (c TsIncrbyValue) Timestamp(timestamp string) TsIncrbyTimestamp {
	_ = "STUB: not implemented"
	return *new(TsIncrbyTimestamp)
}

func (c TsIncrbyValue) Retention(retentionperiod int64) TsIncrbyRetention {
	_ = "STUB: not implemented"
	return *new(TsIncrbyRetention)
}

func (c TsIncrbyValue) Uncompressed() TsIncrbyUncompressed {
	_ = "STUB: not implemented"
	return *new(TsIncrbyUncompressed)
}

func (c TsIncrbyValue) ChunkSize(size int64) TsIncrbyChunkSize {
	_ = "STUB: not implemented"
	return *new(TsIncrbyChunkSize)
}

func (c TsIncrbyValue) Labels() TsIncrbyLabels {
	_ = "STUB: not implemented"
	return *new(TsIncrbyLabels)
}

func (c TsIncrbyValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsInfo Incomplete

func (b Builder) TsInfo() (c TsInfo) { _ = "STUB: not implemented"; return *new(TsInfo) }

func (c TsInfo) Key(key string) TsInfoKey { _ = "STUB: not implemented"; return *new(TsInfoKey) }

type TsInfoDebug Incomplete

func (c TsInfoDebug) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsInfoKey Incomplete

func (c TsInfoKey) Debug(debug string) TsInfoDebug {
	_ = "STUB: not implemented"
	return *new(TsInfoDebug)
}

func (c TsInfoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsMadd Incomplete

func (b Builder) TsMadd() (c TsMadd) { _ = "STUB: not implemented"; return *new(TsMadd) }

func (c TsMadd) KeyTimestampValue() TsMaddKeyTimestampValue {
	_ = "STUB: not implemented"
	return *new(TsMaddKeyTimestampValue)
}

type TsMaddKeyTimestampValue Incomplete

func (c TsMaddKeyTimestampValue) KeyTimestampValue(key string, timestamp int64, value float64) TsMaddKeyTimestampValue {
	_ = "STUB: not implemented"
	return *new(TsMaddKeyTimestampValue)
}

func (c TsMaddKeyTimestampValue) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsMget Incomplete

func (b Builder) TsMget() (c TsMget) { _ = "STUB: not implemented"; return *new(TsMget) }

func (c TsMget) Latest() TsMgetLatest { _ = "STUB: not implemented"; return *new(TsMgetLatest) }

func (c TsMget) Withlabels() TsMgetWithlabels {
	_ = "STUB: not implemented"
	return *new(TsMgetWithlabels)
}

func (c TsMget) SelectedLabels(labels []string) TsMgetSelectedLabels {
	_ = "STUB: not implemented"
	return *new(TsMgetSelectedLabels)
}

func (c TsMget) Filter(filter ...string) TsMgetFilter {
	_ = "STUB: not implemented"
	return *new(TsMgetFilter)
}

type TsMgetFilter Incomplete

func (c TsMgetFilter) Filter(filter ...string) TsMgetFilter {
	_ = "STUB: not implemented"
	return *new(TsMgetFilter)
}

func (c TsMgetFilter) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsMgetLatest Incomplete

func (c TsMgetLatest) Withlabels() TsMgetWithlabels {
	_ = "STUB: not implemented"
	return *new(TsMgetWithlabels)
}

func (c TsMgetLatest) SelectedLabels(labels []string) TsMgetSelectedLabels {
	_ = "STUB: not implemented"
	return *new(TsMgetSelectedLabels)
}

func (c TsMgetLatest) Filter(filter ...string) TsMgetFilter {
	_ = "STUB: not implemented"
	return *new(TsMgetFilter)
}

type TsMgetSelectedLabels Incomplete

func (c TsMgetSelectedLabels) Filter(filter ...string) TsMgetFilter {
	_ = "STUB: not implemented"
	return *new(TsMgetFilter)
}

type TsMgetWithlabels Incomplete

func (c TsMgetWithlabels) Filter(filter ...string) TsMgetFilter {
	_ = "STUB: not implemented"
	return *new(TsMgetFilter)
}

type TsMrange Incomplete

func (b Builder) TsMrange() (c TsMrange) { _ = "STUB: not implemented"; return *new(TsMrange) }

func (c TsMrange) Fromtimestamp(fromtimestamp string) TsMrangeFromtimestamp {
	_ = "STUB: not implemented"
	return *new(TsMrangeFromtimestamp)
}

type TsMrangeAggregationAggregationAvg Incomplete

func (c TsMrangeAggregationAggregationAvg) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationCount Incomplete

func (c TsMrangeAggregationAggregationCount) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationFirst Incomplete

func (c TsMrangeAggregationAggregationFirst) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationLast Incomplete

func (c TsMrangeAggregationAggregationLast) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationMax Incomplete

func (c TsMrangeAggregationAggregationMax) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationMin Incomplete

func (c TsMrangeAggregationAggregationMin) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationRange Incomplete

func (c TsMrangeAggregationAggregationRange) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationStdP Incomplete

func (c TsMrangeAggregationAggregationStdP) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationStdS Incomplete

func (c TsMrangeAggregationAggregationStdS) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationSum Incomplete

func (c TsMrangeAggregationAggregationSum) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationTwa Incomplete

func (c TsMrangeAggregationAggregationTwa) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationVarP Incomplete

func (c TsMrangeAggregationAggregationVarP) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationAggregationVarS Incomplete

func (c TsMrangeAggregationAggregationVarS) Bucketduration(bucketduration int64) TsMrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBucketduration)
}

type TsMrangeAggregationBucketduration Incomplete

func (c TsMrangeAggregationBucketduration) Buckettimestamp(buckettimestamp string) TsMrangeAggregationBuckettimestamp {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationBuckettimestamp)
}

func (c TsMrangeAggregationBucketduration) Empty() TsMrangeAggregationEmpty {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationEmpty)
}

func (c TsMrangeAggregationBucketduration) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

type TsMrangeAggregationBuckettimestamp Incomplete

func (c TsMrangeAggregationBuckettimestamp) Empty() TsMrangeAggregationEmpty {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationEmpty)
}

func (c TsMrangeAggregationBuckettimestamp) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

type TsMrangeAggregationEmpty Incomplete

func (c TsMrangeAggregationEmpty) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

type TsMrangeAlign Incomplete

func (c TsMrangeAlign) AggregationAvg() TsMrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationAvg)
}

func (c TsMrangeAlign) AggregationSum() TsMrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationSum)
}

func (c TsMrangeAlign) AggregationMin() TsMrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMin)
}

func (c TsMrangeAlign) AggregationMax() TsMrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMax)
}

func (c TsMrangeAlign) AggregationRange() TsMrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationRange)
}

func (c TsMrangeAlign) AggregationCount() TsMrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationCount)
}

func (c TsMrangeAlign) AggregationFirst() TsMrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationFirst)
}

func (c TsMrangeAlign) AggregationLast() TsMrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationLast)
}

func (c TsMrangeAlign) AggregationStdP() TsMrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdP)
}

func (c TsMrangeAlign) AggregationStdS() TsMrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdS)
}

func (c TsMrangeAlign) AggregationVarP() TsMrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarP)
}

func (c TsMrangeAlign) AggregationVarS() TsMrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarS)
}

func (c TsMrangeAlign) AggregationTwa() TsMrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationTwa)
}

func (c TsMrangeAlign) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

type TsMrangeCount Incomplete

func (c TsMrangeCount) Align(value string) TsMrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrangeAlign)
}

func (c TsMrangeCount) AggregationAvg() TsMrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationAvg)
}

func (c TsMrangeCount) AggregationSum() TsMrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationSum)
}

func (c TsMrangeCount) AggregationMin() TsMrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMin)
}

func (c TsMrangeCount) AggregationMax() TsMrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMax)
}

func (c TsMrangeCount) AggregationRange() TsMrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationRange)
}

func (c TsMrangeCount) AggregationCount() TsMrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationCount)
}

func (c TsMrangeCount) AggregationFirst() TsMrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationFirst)
}

func (c TsMrangeCount) AggregationLast() TsMrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationLast)
}

func (c TsMrangeCount) AggregationStdP() TsMrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdP)
}

func (c TsMrangeCount) AggregationStdS() TsMrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdS)
}

func (c TsMrangeCount) AggregationVarP() TsMrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarP)
}

func (c TsMrangeCount) AggregationVarS() TsMrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarS)
}

func (c TsMrangeCount) AggregationTwa() TsMrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationTwa)
}

func (c TsMrangeCount) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

type TsMrangeFilter Incomplete

func (c TsMrangeFilter) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

func (c TsMrangeFilter) Groupby(label string, reduce string, reducer string) TsMrangeGroupby {
	_ = "STUB: not implemented"
	return *new(TsMrangeGroupby)
}

func (c TsMrangeFilter) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsMrangeFilterByTs Incomplete

func (c TsMrangeFilterByTs) FilterByTs(timestamp ...int64) TsMrangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilterByTs)
}

func (c TsMrangeFilterByTs) FilterByValue(min float64, max float64) TsMrangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilterByValue)
}

func (c TsMrangeFilterByTs) Withlabels() TsMrangeWithlabels {
	_ = "STUB: not implemented"
	return *new(TsMrangeWithlabels)
}

func (c TsMrangeFilterByTs) SelectedLabels(labels []string) TsMrangeSelectedLabels {
	_ = "STUB: not implemented"
	return *new(TsMrangeSelectedLabels)
}

func (c TsMrangeFilterByTs) Count(count int64) TsMrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeCount)
}

func (c TsMrangeFilterByTs) Align(value string) TsMrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrangeAlign)
}

func (c TsMrangeFilterByTs) AggregationAvg() TsMrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationAvg)
}

func (c TsMrangeFilterByTs) AggregationSum() TsMrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationSum)
}

func (c TsMrangeFilterByTs) AggregationMin() TsMrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMin)
}

func (c TsMrangeFilterByTs) AggregationMax() TsMrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMax)
}

func (c TsMrangeFilterByTs) AggregationRange() TsMrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationRange)
}

func (c TsMrangeFilterByTs) AggregationCount() TsMrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationCount)
}

func (c TsMrangeFilterByTs) AggregationFirst() TsMrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationFirst)
}

func (c TsMrangeFilterByTs) AggregationLast() TsMrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationLast)
}

func (c TsMrangeFilterByTs) AggregationStdP() TsMrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdP)
}

func (c TsMrangeFilterByTs) AggregationStdS() TsMrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdS)
}

func (c TsMrangeFilterByTs) AggregationVarP() TsMrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarP)
}

func (c TsMrangeFilterByTs) AggregationVarS() TsMrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarS)
}

func (c TsMrangeFilterByTs) AggregationTwa() TsMrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationTwa)
}

func (c TsMrangeFilterByTs) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

type TsMrangeFilterByValue Incomplete

func (c TsMrangeFilterByValue) Withlabels() TsMrangeWithlabels {
	_ = "STUB: not implemented"
	return *new(TsMrangeWithlabels)
}

func (c TsMrangeFilterByValue) SelectedLabels(labels []string) TsMrangeSelectedLabels {
	_ = "STUB: not implemented"
	return *new(TsMrangeSelectedLabels)
}

func (c TsMrangeFilterByValue) Count(count int64) TsMrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeCount)
}

func (c TsMrangeFilterByValue) Align(value string) TsMrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrangeAlign)
}

func (c TsMrangeFilterByValue) AggregationAvg() TsMrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationAvg)
}

func (c TsMrangeFilterByValue) AggregationSum() TsMrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationSum)
}

func (c TsMrangeFilterByValue) AggregationMin() TsMrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMin)
}

func (c TsMrangeFilterByValue) AggregationMax() TsMrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMax)
}

func (c TsMrangeFilterByValue) AggregationRange() TsMrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationRange)
}

func (c TsMrangeFilterByValue) AggregationCount() TsMrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationCount)
}

func (c TsMrangeFilterByValue) AggregationFirst() TsMrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationFirst)
}

func (c TsMrangeFilterByValue) AggregationLast() TsMrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationLast)
}

func (c TsMrangeFilterByValue) AggregationStdP() TsMrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdP)
}

func (c TsMrangeFilterByValue) AggregationStdS() TsMrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdS)
}

func (c TsMrangeFilterByValue) AggregationVarP() TsMrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarP)
}

func (c TsMrangeFilterByValue) AggregationVarS() TsMrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarS)
}

func (c TsMrangeFilterByValue) AggregationTwa() TsMrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationTwa)
}

func (c TsMrangeFilterByValue) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

type TsMrangeFromtimestamp Incomplete

func (c TsMrangeFromtimestamp) Totimestamp(totimestamp string) TsMrangeTotimestamp {
	_ = "STUB: not implemented"
	return *new(TsMrangeTotimestamp)
}

type TsMrangeGroupby Incomplete

func (c TsMrangeGroupby) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsMrangeLatest Incomplete

func (c TsMrangeLatest) FilterByTs(timestamp ...int64) TsMrangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilterByTs)
}

func (c TsMrangeLatest) FilterByValue(min float64, max float64) TsMrangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilterByValue)
}

func (c TsMrangeLatest) Withlabels() TsMrangeWithlabels {
	_ = "STUB: not implemented"
	return *new(TsMrangeWithlabels)
}

func (c TsMrangeLatest) SelectedLabels(labels []string) TsMrangeSelectedLabels {
	_ = "STUB: not implemented"
	return *new(TsMrangeSelectedLabels)
}

func (c TsMrangeLatest) Count(count int64) TsMrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeCount)
}

func (c TsMrangeLatest) Align(value string) TsMrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrangeAlign)
}

func (c TsMrangeLatest) AggregationAvg() TsMrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationAvg)
}

func (c TsMrangeLatest) AggregationSum() TsMrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationSum)
}

func (c TsMrangeLatest) AggregationMin() TsMrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMin)
}

func (c TsMrangeLatest) AggregationMax() TsMrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMax)
}

func (c TsMrangeLatest) AggregationRange() TsMrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationRange)
}

func (c TsMrangeLatest) AggregationCount() TsMrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationCount)
}

func (c TsMrangeLatest) AggregationFirst() TsMrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationFirst)
}

func (c TsMrangeLatest) AggregationLast() TsMrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationLast)
}

func (c TsMrangeLatest) AggregationStdP() TsMrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdP)
}

func (c TsMrangeLatest) AggregationStdS() TsMrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdS)
}

func (c TsMrangeLatest) AggregationVarP() TsMrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarP)
}

func (c TsMrangeLatest) AggregationVarS() TsMrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarS)
}

func (c TsMrangeLatest) AggregationTwa() TsMrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationTwa)
}

func (c TsMrangeLatest) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

type TsMrangeSelectedLabels Incomplete

func (c TsMrangeSelectedLabels) Count(count int64) TsMrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeCount)
}

func (c TsMrangeSelectedLabels) Align(value string) TsMrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrangeAlign)
}

func (c TsMrangeSelectedLabels) AggregationAvg() TsMrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationAvg)
}

func (c TsMrangeSelectedLabels) AggregationSum() TsMrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationSum)
}

func (c TsMrangeSelectedLabels) AggregationMin() TsMrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMin)
}

func (c TsMrangeSelectedLabels) AggregationMax() TsMrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMax)
}

func (c TsMrangeSelectedLabels) AggregationRange() TsMrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationRange)
}

func (c TsMrangeSelectedLabels) AggregationCount() TsMrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationCount)
}

func (c TsMrangeSelectedLabels) AggregationFirst() TsMrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationFirst)
}

func (c TsMrangeSelectedLabels) AggregationLast() TsMrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationLast)
}

func (c TsMrangeSelectedLabels) AggregationStdP() TsMrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdP)
}

func (c TsMrangeSelectedLabels) AggregationStdS() TsMrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdS)
}

func (c TsMrangeSelectedLabels) AggregationVarP() TsMrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarP)
}

func (c TsMrangeSelectedLabels) AggregationVarS() TsMrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarS)
}

func (c TsMrangeSelectedLabels) AggregationTwa() TsMrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationTwa)
}

func (c TsMrangeSelectedLabels) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

type TsMrangeTotimestamp Incomplete

func (c TsMrangeTotimestamp) Latest() TsMrangeLatest {
	_ = "STUB: not implemented"
	return *new(TsMrangeLatest)
}

func (c TsMrangeTotimestamp) FilterByTs(timestamp ...int64) TsMrangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilterByTs)
}

func (c TsMrangeTotimestamp) FilterByValue(min float64, max float64) TsMrangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilterByValue)
}

func (c TsMrangeTotimestamp) Withlabels() TsMrangeWithlabels {
	_ = "STUB: not implemented"
	return *new(TsMrangeWithlabels)
}

func (c TsMrangeTotimestamp) SelectedLabels(labels []string) TsMrangeSelectedLabels {
	_ = "STUB: not implemented"
	return *new(TsMrangeSelectedLabels)
}

func (c TsMrangeTotimestamp) Count(count int64) TsMrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeCount)
}

func (c TsMrangeTotimestamp) Align(value string) TsMrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrangeAlign)
}

func (c TsMrangeTotimestamp) AggregationAvg() TsMrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationAvg)
}

func (c TsMrangeTotimestamp) AggregationSum() TsMrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationSum)
}

func (c TsMrangeTotimestamp) AggregationMin() TsMrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMin)
}

func (c TsMrangeTotimestamp) AggregationMax() TsMrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMax)
}

func (c TsMrangeTotimestamp) AggregationRange() TsMrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationRange)
}

func (c TsMrangeTotimestamp) AggregationCount() TsMrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationCount)
}

func (c TsMrangeTotimestamp) AggregationFirst() TsMrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationFirst)
}

func (c TsMrangeTotimestamp) AggregationLast() TsMrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationLast)
}

func (c TsMrangeTotimestamp) AggregationStdP() TsMrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdP)
}

func (c TsMrangeTotimestamp) AggregationStdS() TsMrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdS)
}

func (c TsMrangeTotimestamp) AggregationVarP() TsMrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarP)
}

func (c TsMrangeTotimestamp) AggregationVarS() TsMrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarS)
}

func (c TsMrangeTotimestamp) AggregationTwa() TsMrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationTwa)
}

func (c TsMrangeTotimestamp) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

type TsMrangeWithlabels Incomplete

func (c TsMrangeWithlabels) Count(count int64) TsMrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeCount)
}

func (c TsMrangeWithlabels) Align(value string) TsMrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrangeAlign)
}

func (c TsMrangeWithlabels) AggregationAvg() TsMrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationAvg)
}

func (c TsMrangeWithlabels) AggregationSum() TsMrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationSum)
}

func (c TsMrangeWithlabels) AggregationMin() TsMrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMin)
}

func (c TsMrangeWithlabels) AggregationMax() TsMrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationMax)
}

func (c TsMrangeWithlabels) AggregationRange() TsMrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationRange)
}

func (c TsMrangeWithlabels) AggregationCount() TsMrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationCount)
}

func (c TsMrangeWithlabels) AggregationFirst() TsMrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationFirst)
}

func (c TsMrangeWithlabels) AggregationLast() TsMrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationLast)
}

func (c TsMrangeWithlabels) AggregationStdP() TsMrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdP)
}

func (c TsMrangeWithlabels) AggregationStdS() TsMrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationStdS)
}

func (c TsMrangeWithlabels) AggregationVarP() TsMrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarP)
}

func (c TsMrangeWithlabels) AggregationVarS() TsMrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationVarS)
}

func (c TsMrangeWithlabels) AggregationTwa() TsMrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrangeAggregationAggregationTwa)
}

func (c TsMrangeWithlabels) Filter(filter ...string) TsMrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrangeFilter)
}

type TsMrevrange Incomplete

func (b Builder) TsMrevrange() (c TsMrevrange) { _ = "STUB: not implemented"; return *new(TsMrevrange) }

func (c TsMrevrange) Fromtimestamp(fromtimestamp string) TsMrevrangeFromtimestamp {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFromtimestamp)
}

type TsMrevrangeAggregationAggregationAvg Incomplete

func (c TsMrevrangeAggregationAggregationAvg) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationCount Incomplete

func (c TsMrevrangeAggregationAggregationCount) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationFirst Incomplete

func (c TsMrevrangeAggregationAggregationFirst) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationLast Incomplete

func (c TsMrevrangeAggregationAggregationLast) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationMax Incomplete

func (c TsMrevrangeAggregationAggregationMax) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationMin Incomplete

func (c TsMrevrangeAggregationAggregationMin) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationRange Incomplete

func (c TsMrevrangeAggregationAggregationRange) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationStdP Incomplete

func (c TsMrevrangeAggregationAggregationStdP) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationStdS Incomplete

func (c TsMrevrangeAggregationAggregationStdS) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationSum Incomplete

func (c TsMrevrangeAggregationAggregationSum) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationTwa Incomplete

func (c TsMrevrangeAggregationAggregationTwa) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationVarP Incomplete

func (c TsMrevrangeAggregationAggregationVarP) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationAggregationVarS Incomplete

func (c TsMrevrangeAggregationAggregationVarS) Bucketduration(bucketduration int64) TsMrevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBucketduration)
}

type TsMrevrangeAggregationBucketduration Incomplete

func (c TsMrevrangeAggregationBucketduration) Buckettimestamp(buckettimestamp string) TsMrevrangeAggregationBuckettimestamp {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationBuckettimestamp)
}

func (c TsMrevrangeAggregationBucketduration) Empty() TsMrevrangeAggregationEmpty {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationEmpty)
}

func (c TsMrevrangeAggregationBucketduration) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

type TsMrevrangeAggregationBuckettimestamp Incomplete

func (c TsMrevrangeAggregationBuckettimestamp) Empty() TsMrevrangeAggregationEmpty {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationEmpty)
}

func (c TsMrevrangeAggregationBuckettimestamp) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

type TsMrevrangeAggregationEmpty Incomplete

func (c TsMrevrangeAggregationEmpty) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

type TsMrevrangeAlign Incomplete

func (c TsMrevrangeAlign) AggregationAvg() TsMrevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationAvg)
}

func (c TsMrevrangeAlign) AggregationSum() TsMrevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationSum)
}

func (c TsMrevrangeAlign) AggregationMin() TsMrevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMin)
}

func (c TsMrevrangeAlign) AggregationMax() TsMrevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMax)
}

func (c TsMrevrangeAlign) AggregationRange() TsMrevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationRange)
}

func (c TsMrevrangeAlign) AggregationCount() TsMrevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationCount)
}

func (c TsMrevrangeAlign) AggregationFirst() TsMrevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationFirst)
}

func (c TsMrevrangeAlign) AggregationLast() TsMrevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationLast)
}

func (c TsMrevrangeAlign) AggregationStdP() TsMrevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdP)
}

func (c TsMrevrangeAlign) AggregationStdS() TsMrevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdS)
}

func (c TsMrevrangeAlign) AggregationVarP() TsMrevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarP)
}

func (c TsMrevrangeAlign) AggregationVarS() TsMrevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarS)
}

func (c TsMrevrangeAlign) AggregationTwa() TsMrevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationTwa)
}

func (c TsMrevrangeAlign) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

type TsMrevrangeCount Incomplete

func (c TsMrevrangeCount) Align(value string) TsMrevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAlign)
}

func (c TsMrevrangeCount) AggregationAvg() TsMrevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationAvg)
}

func (c TsMrevrangeCount) AggregationSum() TsMrevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationSum)
}

func (c TsMrevrangeCount) AggregationMin() TsMrevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMin)
}

func (c TsMrevrangeCount) AggregationMax() TsMrevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMax)
}

func (c TsMrevrangeCount) AggregationRange() TsMrevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationRange)
}

func (c TsMrevrangeCount) AggregationCount() TsMrevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationCount)
}

func (c TsMrevrangeCount) AggregationFirst() TsMrevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationFirst)
}

func (c TsMrevrangeCount) AggregationLast() TsMrevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationLast)
}

func (c TsMrevrangeCount) AggregationStdP() TsMrevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdP)
}

func (c TsMrevrangeCount) AggregationStdS() TsMrevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdS)
}

func (c TsMrevrangeCount) AggregationVarP() TsMrevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarP)
}

func (c TsMrevrangeCount) AggregationVarS() TsMrevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarS)
}

func (c TsMrevrangeCount) AggregationTwa() TsMrevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationTwa)
}

func (c TsMrevrangeCount) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

type TsMrevrangeFilter Incomplete

func (c TsMrevrangeFilter) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

func (c TsMrevrangeFilter) Groupby(label string, reduce string, reducer string) TsMrevrangeGroupby {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeGroupby)
}

func (c TsMrevrangeFilter) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsMrevrangeFilterByTs Incomplete

func (c TsMrevrangeFilterByTs) FilterByTs(timestamp ...int64) TsMrevrangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilterByTs)
}

func (c TsMrevrangeFilterByTs) FilterByValue(min float64, max float64) TsMrevrangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilterByValue)
}

func (c TsMrevrangeFilterByTs) Withlabels() TsMrevrangeWithlabels {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeWithlabels)
}

func (c TsMrevrangeFilterByTs) SelectedLabels(labels []string) TsMrevrangeSelectedLabels {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeSelectedLabels)
}

func (c TsMrevrangeFilterByTs) Count(count int64) TsMrevrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeCount)
}

func (c TsMrevrangeFilterByTs) Align(value string) TsMrevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAlign)
}

func (c TsMrevrangeFilterByTs) AggregationAvg() TsMrevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationAvg)
}

func (c TsMrevrangeFilterByTs) AggregationSum() TsMrevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationSum)
}

func (c TsMrevrangeFilterByTs) AggregationMin() TsMrevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMin)
}

func (c TsMrevrangeFilterByTs) AggregationMax() TsMrevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMax)
}

func (c TsMrevrangeFilterByTs) AggregationRange() TsMrevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationRange)
}

func (c TsMrevrangeFilterByTs) AggregationCount() TsMrevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationCount)
}

func (c TsMrevrangeFilterByTs) AggregationFirst() TsMrevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationFirst)
}

func (c TsMrevrangeFilterByTs) AggregationLast() TsMrevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationLast)
}

func (c TsMrevrangeFilterByTs) AggregationStdP() TsMrevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdP)
}

func (c TsMrevrangeFilterByTs) AggregationStdS() TsMrevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdS)
}

func (c TsMrevrangeFilterByTs) AggregationVarP() TsMrevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarP)
}

func (c TsMrevrangeFilterByTs) AggregationVarS() TsMrevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarS)
}

func (c TsMrevrangeFilterByTs) AggregationTwa() TsMrevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationTwa)
}

func (c TsMrevrangeFilterByTs) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

type TsMrevrangeFilterByValue Incomplete

func (c TsMrevrangeFilterByValue) Withlabels() TsMrevrangeWithlabels {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeWithlabels)
}

func (c TsMrevrangeFilterByValue) SelectedLabels(labels []string) TsMrevrangeSelectedLabels {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeSelectedLabels)
}

func (c TsMrevrangeFilterByValue) Count(count int64) TsMrevrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeCount)
}

func (c TsMrevrangeFilterByValue) Align(value string) TsMrevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAlign)
}

func (c TsMrevrangeFilterByValue) AggregationAvg() TsMrevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationAvg)
}

func (c TsMrevrangeFilterByValue) AggregationSum() TsMrevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationSum)
}

func (c TsMrevrangeFilterByValue) AggregationMin() TsMrevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMin)
}

func (c TsMrevrangeFilterByValue) AggregationMax() TsMrevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMax)
}

func (c TsMrevrangeFilterByValue) AggregationRange() TsMrevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationRange)
}

func (c TsMrevrangeFilterByValue) AggregationCount() TsMrevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationCount)
}

func (c TsMrevrangeFilterByValue) AggregationFirst() TsMrevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationFirst)
}

func (c TsMrevrangeFilterByValue) AggregationLast() TsMrevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationLast)
}

func (c TsMrevrangeFilterByValue) AggregationStdP() TsMrevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdP)
}

func (c TsMrevrangeFilterByValue) AggregationStdS() TsMrevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdS)
}

func (c TsMrevrangeFilterByValue) AggregationVarP() TsMrevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarP)
}

func (c TsMrevrangeFilterByValue) AggregationVarS() TsMrevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarS)
}

func (c TsMrevrangeFilterByValue) AggregationTwa() TsMrevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationTwa)
}

func (c TsMrevrangeFilterByValue) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

type TsMrevrangeFromtimestamp Incomplete

func (c TsMrevrangeFromtimestamp) Totimestamp(totimestamp string) TsMrevrangeTotimestamp {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeTotimestamp)
}

type TsMrevrangeGroupby Incomplete

func (c TsMrevrangeGroupby) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsMrevrangeLatest Incomplete

func (c TsMrevrangeLatest) FilterByTs(timestamp ...int64) TsMrevrangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilterByTs)
}

func (c TsMrevrangeLatest) FilterByValue(min float64, max float64) TsMrevrangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilterByValue)
}

func (c TsMrevrangeLatest) Withlabels() TsMrevrangeWithlabels {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeWithlabels)
}

func (c TsMrevrangeLatest) SelectedLabels(labels []string) TsMrevrangeSelectedLabels {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeSelectedLabels)
}

func (c TsMrevrangeLatest) Count(count int64) TsMrevrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeCount)
}

func (c TsMrevrangeLatest) Align(value string) TsMrevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAlign)
}

func (c TsMrevrangeLatest) AggregationAvg() TsMrevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationAvg)
}

func (c TsMrevrangeLatest) AggregationSum() TsMrevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationSum)
}

func (c TsMrevrangeLatest) AggregationMin() TsMrevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMin)
}

func (c TsMrevrangeLatest) AggregationMax() TsMrevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMax)
}

func (c TsMrevrangeLatest) AggregationRange() TsMrevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationRange)
}

func (c TsMrevrangeLatest) AggregationCount() TsMrevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationCount)
}

func (c TsMrevrangeLatest) AggregationFirst() TsMrevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationFirst)
}

func (c TsMrevrangeLatest) AggregationLast() TsMrevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationLast)
}

func (c TsMrevrangeLatest) AggregationStdP() TsMrevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdP)
}

func (c TsMrevrangeLatest) AggregationStdS() TsMrevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdS)
}

func (c TsMrevrangeLatest) AggregationVarP() TsMrevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarP)
}

func (c TsMrevrangeLatest) AggregationVarS() TsMrevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarS)
}

func (c TsMrevrangeLatest) AggregationTwa() TsMrevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationTwa)
}

func (c TsMrevrangeLatest) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

type TsMrevrangeSelectedLabels Incomplete

func (c TsMrevrangeSelectedLabels) Count(count int64) TsMrevrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeCount)
}

func (c TsMrevrangeSelectedLabels) Align(value string) TsMrevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAlign)
}

func (c TsMrevrangeSelectedLabels) AggregationAvg() TsMrevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationAvg)
}

func (c TsMrevrangeSelectedLabels) AggregationSum() TsMrevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationSum)
}

func (c TsMrevrangeSelectedLabels) AggregationMin() TsMrevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMin)
}

func (c TsMrevrangeSelectedLabels) AggregationMax() TsMrevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMax)
}

func (c TsMrevrangeSelectedLabels) AggregationRange() TsMrevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationRange)
}

func (c TsMrevrangeSelectedLabels) AggregationCount() TsMrevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationCount)
}

func (c TsMrevrangeSelectedLabels) AggregationFirst() TsMrevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationFirst)
}

func (c TsMrevrangeSelectedLabels) AggregationLast() TsMrevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationLast)
}

func (c TsMrevrangeSelectedLabels) AggregationStdP() TsMrevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdP)
}

func (c TsMrevrangeSelectedLabels) AggregationStdS() TsMrevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdS)
}

func (c TsMrevrangeSelectedLabels) AggregationVarP() TsMrevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarP)
}

func (c TsMrevrangeSelectedLabels) AggregationVarS() TsMrevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarS)
}

func (c TsMrevrangeSelectedLabels) AggregationTwa() TsMrevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationTwa)
}

func (c TsMrevrangeSelectedLabels) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

type TsMrevrangeTotimestamp Incomplete

func (c TsMrevrangeTotimestamp) Latest() TsMrevrangeLatest {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeLatest)
}

func (c TsMrevrangeTotimestamp) FilterByTs(timestamp ...int64) TsMrevrangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilterByTs)
}

func (c TsMrevrangeTotimestamp) FilterByValue(min float64, max float64) TsMrevrangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilterByValue)
}

func (c TsMrevrangeTotimestamp) Withlabels() TsMrevrangeWithlabels {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeWithlabels)
}

func (c TsMrevrangeTotimestamp) SelectedLabels(labels []string) TsMrevrangeSelectedLabels {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeSelectedLabels)
}

func (c TsMrevrangeTotimestamp) Count(count int64) TsMrevrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeCount)
}

func (c TsMrevrangeTotimestamp) Align(value string) TsMrevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAlign)
}

func (c TsMrevrangeTotimestamp) AggregationAvg() TsMrevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationAvg)
}

func (c TsMrevrangeTotimestamp) AggregationSum() TsMrevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationSum)
}

func (c TsMrevrangeTotimestamp) AggregationMin() TsMrevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMin)
}

func (c TsMrevrangeTotimestamp) AggregationMax() TsMrevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMax)
}

func (c TsMrevrangeTotimestamp) AggregationRange() TsMrevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationRange)
}

func (c TsMrevrangeTotimestamp) AggregationCount() TsMrevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationCount)
}

func (c TsMrevrangeTotimestamp) AggregationFirst() TsMrevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationFirst)
}

func (c TsMrevrangeTotimestamp) AggregationLast() TsMrevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationLast)
}

func (c TsMrevrangeTotimestamp) AggregationStdP() TsMrevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdP)
}

func (c TsMrevrangeTotimestamp) AggregationStdS() TsMrevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdS)
}

func (c TsMrevrangeTotimestamp) AggregationVarP() TsMrevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarP)
}

func (c TsMrevrangeTotimestamp) AggregationVarS() TsMrevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarS)
}

func (c TsMrevrangeTotimestamp) AggregationTwa() TsMrevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationTwa)
}

func (c TsMrevrangeTotimestamp) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

type TsMrevrangeWithlabels Incomplete

func (c TsMrevrangeWithlabels) Count(count int64) TsMrevrangeCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeCount)
}

func (c TsMrevrangeWithlabels) Align(value string) TsMrevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAlign)
}

func (c TsMrevrangeWithlabels) AggregationAvg() TsMrevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationAvg)
}

func (c TsMrevrangeWithlabels) AggregationSum() TsMrevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationSum)
}

func (c TsMrevrangeWithlabels) AggregationMin() TsMrevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMin)
}

func (c TsMrevrangeWithlabels) AggregationMax() TsMrevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationMax)
}

func (c TsMrevrangeWithlabels) AggregationRange() TsMrevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationRange)
}

func (c TsMrevrangeWithlabels) AggregationCount() TsMrevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationCount)
}

func (c TsMrevrangeWithlabels) AggregationFirst() TsMrevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationFirst)
}

func (c TsMrevrangeWithlabels) AggregationLast() TsMrevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationLast)
}

func (c TsMrevrangeWithlabels) AggregationStdP() TsMrevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdP)
}

func (c TsMrevrangeWithlabels) AggregationStdS() TsMrevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationStdS)
}

func (c TsMrevrangeWithlabels) AggregationVarP() TsMrevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarP)
}

func (c TsMrevrangeWithlabels) AggregationVarS() TsMrevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationVarS)
}

func (c TsMrevrangeWithlabels) AggregationTwa() TsMrevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeAggregationAggregationTwa)
}

func (c TsMrevrangeWithlabels) Filter(filter ...string) TsMrevrangeFilter {
	_ = "STUB: not implemented"
	return *new(TsMrevrangeFilter)
}

type TsQueryindex Incomplete

func (b Builder) TsQueryindex() (c TsQueryindex) {
	_ = "STUB: not implemented"
	return *new(TsQueryindex)
}

func (c TsQueryindex) Filter(filter ...string) TsQueryindexFilter {
	_ = "STUB: not implemented"
	return *new(TsQueryindexFilter)
}

type TsQueryindexFilter Incomplete

func (c TsQueryindexFilter) Filter(filter ...string) TsQueryindexFilter {
	_ = "STUB: not implemented"
	return *new(TsQueryindexFilter)
}

func (c TsQueryindexFilter) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsRange Incomplete

func (b Builder) TsRange() (c TsRange) { _ = "STUB: not implemented"; return *new(TsRange) }

func (c TsRange) Key(key string) TsRangeKey { _ = "STUB: not implemented"; return *new(TsRangeKey) }

type TsRangeAggregationAggregationAvg Incomplete

func (c TsRangeAggregationAggregationAvg) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationCount Incomplete

func (c TsRangeAggregationAggregationCount) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationFirst Incomplete

func (c TsRangeAggregationAggregationFirst) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationLast Incomplete

func (c TsRangeAggregationAggregationLast) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationMax Incomplete

func (c TsRangeAggregationAggregationMax) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationMin Incomplete

func (c TsRangeAggregationAggregationMin) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationRange Incomplete

func (c TsRangeAggregationAggregationRange) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationStdP Incomplete

func (c TsRangeAggregationAggregationStdP) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationStdS Incomplete

func (c TsRangeAggregationAggregationStdS) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationSum Incomplete

func (c TsRangeAggregationAggregationSum) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationTwa Incomplete

func (c TsRangeAggregationAggregationTwa) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationVarP Incomplete

func (c TsRangeAggregationAggregationVarP) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationAggregationVarS Incomplete

func (c TsRangeAggregationAggregationVarS) Bucketduration(bucketduration int64) TsRangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBucketduration)
}

type TsRangeAggregationBucketduration Incomplete

func (c TsRangeAggregationBucketduration) Buckettimestamp(buckettimestamp string) TsRangeAggregationBuckettimestamp {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationBuckettimestamp)
}

func (c TsRangeAggregationBucketduration) Empty() TsRangeAggregationEmpty {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationEmpty)
}

func (c TsRangeAggregationBucketduration) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsRangeAggregationBuckettimestamp Incomplete

func (c TsRangeAggregationBuckettimestamp) Empty() TsRangeAggregationEmpty {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationEmpty)
}

func (c TsRangeAggregationBuckettimestamp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsRangeAggregationEmpty Incomplete

func (c TsRangeAggregationEmpty) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsRangeAlign Incomplete

func (c TsRangeAlign) AggregationAvg() TsRangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationAvg)
}

func (c TsRangeAlign) AggregationSum() TsRangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationSum)
}

func (c TsRangeAlign) AggregationMin() TsRangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMin)
}

func (c TsRangeAlign) AggregationMax() TsRangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMax)
}

func (c TsRangeAlign) AggregationRange() TsRangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationRange)
}

func (c TsRangeAlign) AggregationCount() TsRangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationCount)
}

func (c TsRangeAlign) AggregationFirst() TsRangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationFirst)
}

func (c TsRangeAlign) AggregationLast() TsRangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationLast)
}

func (c TsRangeAlign) AggregationStdP() TsRangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdP)
}

func (c TsRangeAlign) AggregationStdS() TsRangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdS)
}

func (c TsRangeAlign) AggregationVarP() TsRangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarP)
}

func (c TsRangeAlign) AggregationVarS() TsRangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarS)
}

func (c TsRangeAlign) AggregationTwa() TsRangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationTwa)
}

func (c TsRangeAlign) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsRangeCount Incomplete

func (c TsRangeCount) Align(value string) TsRangeAlign {
	_ = "STUB: not implemented"
	return *new(TsRangeAlign)
}

func (c TsRangeCount) AggregationAvg() TsRangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationAvg)
}

func (c TsRangeCount) AggregationSum() TsRangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationSum)
}

func (c TsRangeCount) AggregationMin() TsRangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMin)
}

func (c TsRangeCount) AggregationMax() TsRangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMax)
}

func (c TsRangeCount) AggregationRange() TsRangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationRange)
}

func (c TsRangeCount) AggregationCount() TsRangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationCount)
}

func (c TsRangeCount) AggregationFirst() TsRangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationFirst)
}

func (c TsRangeCount) AggregationLast() TsRangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationLast)
}

func (c TsRangeCount) AggregationStdP() TsRangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdP)
}

func (c TsRangeCount) AggregationStdS() TsRangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdS)
}

func (c TsRangeCount) AggregationVarP() TsRangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarP)
}

func (c TsRangeCount) AggregationVarS() TsRangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarS)
}

func (c TsRangeCount) AggregationTwa() TsRangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationTwa)
}

func (c TsRangeCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsRangeFilterByTs Incomplete

func (c TsRangeFilterByTs) FilterByTs(timestamp ...int64) TsRangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsRangeFilterByTs)
}

func (c TsRangeFilterByTs) FilterByValue(min float64, max float64) TsRangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsRangeFilterByValue)
}

func (c TsRangeFilterByTs) Count(count int64) TsRangeCount {
	_ = "STUB: not implemented"
	return *new(TsRangeCount)
}

func (c TsRangeFilterByTs) Align(value string) TsRangeAlign {
	_ = "STUB: not implemented"
	return *new(TsRangeAlign)
}

func (c TsRangeFilterByTs) AggregationAvg() TsRangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationAvg)
}

func (c TsRangeFilterByTs) AggregationSum() TsRangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationSum)
}

func (c TsRangeFilterByTs) AggregationMin() TsRangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMin)
}

func (c TsRangeFilterByTs) AggregationMax() TsRangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMax)
}

func (c TsRangeFilterByTs) AggregationRange() TsRangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationRange)
}

func (c TsRangeFilterByTs) AggregationCount() TsRangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationCount)
}

func (c TsRangeFilterByTs) AggregationFirst() TsRangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationFirst)
}

func (c TsRangeFilterByTs) AggregationLast() TsRangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationLast)
}

func (c TsRangeFilterByTs) AggregationStdP() TsRangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdP)
}

func (c TsRangeFilterByTs) AggregationStdS() TsRangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdS)
}

func (c TsRangeFilterByTs) AggregationVarP() TsRangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarP)
}

func (c TsRangeFilterByTs) AggregationVarS() TsRangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarS)
}

func (c TsRangeFilterByTs) AggregationTwa() TsRangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationTwa)
}

func (c TsRangeFilterByTs) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsRangeFilterByValue Incomplete

func (c TsRangeFilterByValue) Count(count int64) TsRangeCount {
	_ = "STUB: not implemented"
	return *new(TsRangeCount)
}

func (c TsRangeFilterByValue) Align(value string) TsRangeAlign {
	_ = "STUB: not implemented"
	return *new(TsRangeAlign)
}

func (c TsRangeFilterByValue) AggregationAvg() TsRangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationAvg)
}

func (c TsRangeFilterByValue) AggregationSum() TsRangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationSum)
}

func (c TsRangeFilterByValue) AggregationMin() TsRangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMin)
}

func (c TsRangeFilterByValue) AggregationMax() TsRangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMax)
}

func (c TsRangeFilterByValue) AggregationRange() TsRangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationRange)
}

func (c TsRangeFilterByValue) AggregationCount() TsRangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationCount)
}

func (c TsRangeFilterByValue) AggregationFirst() TsRangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationFirst)
}

func (c TsRangeFilterByValue) AggregationLast() TsRangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationLast)
}

func (c TsRangeFilterByValue) AggregationStdP() TsRangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdP)
}

func (c TsRangeFilterByValue) AggregationStdS() TsRangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdS)
}

func (c TsRangeFilterByValue) AggregationVarP() TsRangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarP)
}

func (c TsRangeFilterByValue) AggregationVarS() TsRangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarS)
}

func (c TsRangeFilterByValue) AggregationTwa() TsRangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationTwa)
}

func (c TsRangeFilterByValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsRangeFromtimestamp Incomplete

func (c TsRangeFromtimestamp) Totimestamp(totimestamp string) TsRangeTotimestamp {
	_ = "STUB: not implemented"
	return *new(TsRangeTotimestamp)
}

type TsRangeKey Incomplete

func (c TsRangeKey) Fromtimestamp(fromtimestamp string) TsRangeFromtimestamp {
	_ = "STUB: not implemented"
	return *new(TsRangeFromtimestamp)
}

type TsRangeLatest Incomplete

func (c TsRangeLatest) FilterByTs(timestamp ...int64) TsRangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsRangeFilterByTs)
}

func (c TsRangeLatest) FilterByValue(min float64, max float64) TsRangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsRangeFilterByValue)
}

func (c TsRangeLatest) Count(count int64) TsRangeCount {
	_ = "STUB: not implemented"
	return *new(TsRangeCount)
}

func (c TsRangeLatest) Align(value string) TsRangeAlign {
	_ = "STUB: not implemented"
	return *new(TsRangeAlign)
}

func (c TsRangeLatest) AggregationAvg() TsRangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationAvg)
}

func (c TsRangeLatest) AggregationSum() TsRangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationSum)
}

func (c TsRangeLatest) AggregationMin() TsRangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMin)
}

func (c TsRangeLatest) AggregationMax() TsRangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMax)
}

func (c TsRangeLatest) AggregationRange() TsRangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationRange)
}

func (c TsRangeLatest) AggregationCount() TsRangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationCount)
}

func (c TsRangeLatest) AggregationFirst() TsRangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationFirst)
}

func (c TsRangeLatest) AggregationLast() TsRangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationLast)
}

func (c TsRangeLatest) AggregationStdP() TsRangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdP)
}

func (c TsRangeLatest) AggregationStdS() TsRangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdS)
}

func (c TsRangeLatest) AggregationVarP() TsRangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarP)
}

func (c TsRangeLatest) AggregationVarS() TsRangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarS)
}

func (c TsRangeLatest) AggregationTwa() TsRangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationTwa)
}

func (c TsRangeLatest) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsRangeTotimestamp Incomplete

func (c TsRangeTotimestamp) Latest() TsRangeLatest {
	_ = "STUB: not implemented"
	return *new(TsRangeLatest)
}

func (c TsRangeTotimestamp) FilterByTs(timestamp ...int64) TsRangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsRangeFilterByTs)
}

func (c TsRangeTotimestamp) FilterByValue(min float64, max float64) TsRangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsRangeFilterByValue)
}

func (c TsRangeTotimestamp) Count(count int64) TsRangeCount {
	_ = "STUB: not implemented"
	return *new(TsRangeCount)
}

func (c TsRangeTotimestamp) Align(value string) TsRangeAlign {
	_ = "STUB: not implemented"
	return *new(TsRangeAlign)
}

func (c TsRangeTotimestamp) AggregationAvg() TsRangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationAvg)
}

func (c TsRangeTotimestamp) AggregationSum() TsRangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationSum)
}

func (c TsRangeTotimestamp) AggregationMin() TsRangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMin)
}

func (c TsRangeTotimestamp) AggregationMax() TsRangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationMax)
}

func (c TsRangeTotimestamp) AggregationRange() TsRangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationRange)
}

func (c TsRangeTotimestamp) AggregationCount() TsRangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationCount)
}

func (c TsRangeTotimestamp) AggregationFirst() TsRangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationFirst)
}

func (c TsRangeTotimestamp) AggregationLast() TsRangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationLast)
}

func (c TsRangeTotimestamp) AggregationStdP() TsRangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdP)
}

func (c TsRangeTotimestamp) AggregationStdS() TsRangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationStdS)
}

func (c TsRangeTotimestamp) AggregationVarP() TsRangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarP)
}

func (c TsRangeTotimestamp) AggregationVarS() TsRangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationVarS)
}

func (c TsRangeTotimestamp) AggregationTwa() TsRangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRangeAggregationAggregationTwa)
}

func (c TsRangeTotimestamp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsRevrange Incomplete

func (b Builder) TsRevrange() (c TsRevrange) { _ = "STUB: not implemented"; return *new(TsRevrange) }

func (c TsRevrange) Key(key string) TsRevrangeKey {
	_ = "STUB: not implemented"
	return *new(TsRevrangeKey)
}

type TsRevrangeAggregationAggregationAvg Incomplete

func (c TsRevrangeAggregationAggregationAvg) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationCount Incomplete

func (c TsRevrangeAggregationAggregationCount) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationFirst Incomplete

func (c TsRevrangeAggregationAggregationFirst) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationLast Incomplete

func (c TsRevrangeAggregationAggregationLast) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationMax Incomplete

func (c TsRevrangeAggregationAggregationMax) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationMin Incomplete

func (c TsRevrangeAggregationAggregationMin) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationRange Incomplete

func (c TsRevrangeAggregationAggregationRange) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationStdP Incomplete

func (c TsRevrangeAggregationAggregationStdP) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationStdS Incomplete

func (c TsRevrangeAggregationAggregationStdS) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationSum Incomplete

func (c TsRevrangeAggregationAggregationSum) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationTwa Incomplete

func (c TsRevrangeAggregationAggregationTwa) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationVarP Incomplete

func (c TsRevrangeAggregationAggregationVarP) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationAggregationVarS Incomplete

func (c TsRevrangeAggregationAggregationVarS) Bucketduration(bucketduration int64) TsRevrangeAggregationBucketduration {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBucketduration)
}

type TsRevrangeAggregationBucketduration Incomplete

func (c TsRevrangeAggregationBucketduration) Buckettimestamp(buckettimestamp string) TsRevrangeAggregationBuckettimestamp {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationBuckettimestamp)
}

func (c TsRevrangeAggregationBucketduration) Empty() TsRevrangeAggregationEmpty {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationEmpty)
}

func (c TsRevrangeAggregationBucketduration) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsRevrangeAggregationBuckettimestamp Incomplete

func (c TsRevrangeAggregationBuckettimestamp) Empty() TsRevrangeAggregationEmpty {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationEmpty)
}

func (c TsRevrangeAggregationBuckettimestamp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsRevrangeAggregationEmpty Incomplete

func (c TsRevrangeAggregationEmpty) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsRevrangeAlign Incomplete

func (c TsRevrangeAlign) AggregationAvg() TsRevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationAvg)
}

func (c TsRevrangeAlign) AggregationSum() TsRevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationSum)
}

func (c TsRevrangeAlign) AggregationMin() TsRevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMin)
}

func (c TsRevrangeAlign) AggregationMax() TsRevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMax)
}

func (c TsRevrangeAlign) AggregationRange() TsRevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationRange)
}

func (c TsRevrangeAlign) AggregationCount() TsRevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationCount)
}

func (c TsRevrangeAlign) AggregationFirst() TsRevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationFirst)
}

func (c TsRevrangeAlign) AggregationLast() TsRevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationLast)
}

func (c TsRevrangeAlign) AggregationStdP() TsRevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdP)
}

func (c TsRevrangeAlign) AggregationStdS() TsRevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdS)
}

func (c TsRevrangeAlign) AggregationVarP() TsRevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarP)
}

func (c TsRevrangeAlign) AggregationVarS() TsRevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarS)
}

func (c TsRevrangeAlign) AggregationTwa() TsRevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationTwa)
}

func (c TsRevrangeAlign) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsRevrangeCount Incomplete

func (c TsRevrangeCount) Align(value string) TsRevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAlign)
}

func (c TsRevrangeCount) AggregationAvg() TsRevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationAvg)
}

func (c TsRevrangeCount) AggregationSum() TsRevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationSum)
}

func (c TsRevrangeCount) AggregationMin() TsRevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMin)
}

func (c TsRevrangeCount) AggregationMax() TsRevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMax)
}

func (c TsRevrangeCount) AggregationRange() TsRevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationRange)
}

func (c TsRevrangeCount) AggregationCount() TsRevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationCount)
}

func (c TsRevrangeCount) AggregationFirst() TsRevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationFirst)
}

func (c TsRevrangeCount) AggregationLast() TsRevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationLast)
}

func (c TsRevrangeCount) AggregationStdP() TsRevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdP)
}

func (c TsRevrangeCount) AggregationStdS() TsRevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdS)
}

func (c TsRevrangeCount) AggregationVarP() TsRevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarP)
}

func (c TsRevrangeCount) AggregationVarS() TsRevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarS)
}

func (c TsRevrangeCount) AggregationTwa() TsRevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationTwa)
}

func (c TsRevrangeCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsRevrangeFilterByTs Incomplete

func (c TsRevrangeFilterByTs) FilterByTs(timestamp ...int64) TsRevrangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsRevrangeFilterByTs)
}

func (c TsRevrangeFilterByTs) FilterByValue(min float64, max float64) TsRevrangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsRevrangeFilterByValue)
}

func (c TsRevrangeFilterByTs) Count(count int64) TsRevrangeCount {
	_ = "STUB: not implemented"
	return *new(TsRevrangeCount)
}

func (c TsRevrangeFilterByTs) Align(value string) TsRevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAlign)
}

func (c TsRevrangeFilterByTs) AggregationAvg() TsRevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationAvg)
}

func (c TsRevrangeFilterByTs) AggregationSum() TsRevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationSum)
}

func (c TsRevrangeFilterByTs) AggregationMin() TsRevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMin)
}

func (c TsRevrangeFilterByTs) AggregationMax() TsRevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMax)
}

func (c TsRevrangeFilterByTs) AggregationRange() TsRevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationRange)
}

func (c TsRevrangeFilterByTs) AggregationCount() TsRevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationCount)
}

func (c TsRevrangeFilterByTs) AggregationFirst() TsRevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationFirst)
}

func (c TsRevrangeFilterByTs) AggregationLast() TsRevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationLast)
}

func (c TsRevrangeFilterByTs) AggregationStdP() TsRevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdP)
}

func (c TsRevrangeFilterByTs) AggregationStdS() TsRevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdS)
}

func (c TsRevrangeFilterByTs) AggregationVarP() TsRevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarP)
}

func (c TsRevrangeFilterByTs) AggregationVarS() TsRevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarS)
}

func (c TsRevrangeFilterByTs) AggregationTwa() TsRevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationTwa)
}

func (c TsRevrangeFilterByTs) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsRevrangeFilterByValue Incomplete

func (c TsRevrangeFilterByValue) Count(count int64) TsRevrangeCount {
	_ = "STUB: not implemented"
	return *new(TsRevrangeCount)
}

func (c TsRevrangeFilterByValue) Align(value string) TsRevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAlign)
}

func (c TsRevrangeFilterByValue) AggregationAvg() TsRevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationAvg)
}

func (c TsRevrangeFilterByValue) AggregationSum() TsRevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationSum)
}

func (c TsRevrangeFilterByValue) AggregationMin() TsRevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMin)
}

func (c TsRevrangeFilterByValue) AggregationMax() TsRevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMax)
}

func (c TsRevrangeFilterByValue) AggregationRange() TsRevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationRange)
}

func (c TsRevrangeFilterByValue) AggregationCount() TsRevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationCount)
}

func (c TsRevrangeFilterByValue) AggregationFirst() TsRevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationFirst)
}

func (c TsRevrangeFilterByValue) AggregationLast() TsRevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationLast)
}

func (c TsRevrangeFilterByValue) AggregationStdP() TsRevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdP)
}

func (c TsRevrangeFilterByValue) AggregationStdS() TsRevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdS)
}

func (c TsRevrangeFilterByValue) AggregationVarP() TsRevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarP)
}

func (c TsRevrangeFilterByValue) AggregationVarS() TsRevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarS)
}

func (c TsRevrangeFilterByValue) AggregationTwa() TsRevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationTwa)
}

func (c TsRevrangeFilterByValue) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TsRevrangeFromtimestamp Incomplete

func (c TsRevrangeFromtimestamp) Totimestamp(totimestamp string) TsRevrangeTotimestamp {
	_ = "STUB: not implemented"
	return *new(TsRevrangeTotimestamp)
}

type TsRevrangeKey Incomplete

func (c TsRevrangeKey) Fromtimestamp(fromtimestamp string) TsRevrangeFromtimestamp {
	_ = "STUB: not implemented"
	return *new(TsRevrangeFromtimestamp)
}

type TsRevrangeLatest Incomplete

func (c TsRevrangeLatest) FilterByTs(timestamp ...int64) TsRevrangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsRevrangeFilterByTs)
}

func (c TsRevrangeLatest) FilterByValue(min float64, max float64) TsRevrangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsRevrangeFilterByValue)
}

func (c TsRevrangeLatest) Count(count int64) TsRevrangeCount {
	_ = "STUB: not implemented"
	return *new(TsRevrangeCount)
}

func (c TsRevrangeLatest) Align(value string) TsRevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAlign)
}

func (c TsRevrangeLatest) AggregationAvg() TsRevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationAvg)
}

func (c TsRevrangeLatest) AggregationSum() TsRevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationSum)
}

func (c TsRevrangeLatest) AggregationMin() TsRevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMin)
}

func (c TsRevrangeLatest) AggregationMax() TsRevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMax)
}

func (c TsRevrangeLatest) AggregationRange() TsRevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationRange)
}

func (c TsRevrangeLatest) AggregationCount() TsRevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationCount)
}

func (c TsRevrangeLatest) AggregationFirst() TsRevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationFirst)
}

func (c TsRevrangeLatest) AggregationLast() TsRevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationLast)
}

func (c TsRevrangeLatest) AggregationStdP() TsRevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdP)
}

func (c TsRevrangeLatest) AggregationStdS() TsRevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdS)
}

func (c TsRevrangeLatest) AggregationVarP() TsRevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarP)
}

func (c TsRevrangeLatest) AggregationVarS() TsRevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarS)
}

func (c TsRevrangeLatest) AggregationTwa() TsRevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationTwa)
}

func (c TsRevrangeLatest) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TsRevrangeTotimestamp Incomplete

func (c TsRevrangeTotimestamp) Latest() TsRevrangeLatest {
	_ = "STUB: not implemented"
	return *new(TsRevrangeLatest)
}

func (c TsRevrangeTotimestamp) FilterByTs(timestamp ...int64) TsRevrangeFilterByTs {
	_ = "STUB: not implemented"
	return *new(TsRevrangeFilterByTs)
}

func (c TsRevrangeTotimestamp) FilterByValue(min float64, max float64) TsRevrangeFilterByValue {
	_ = "STUB: not implemented"
	return *new(TsRevrangeFilterByValue)
}

func (c TsRevrangeTotimestamp) Count(count int64) TsRevrangeCount {
	_ = "STUB: not implemented"
	return *new(TsRevrangeCount)
}

func (c TsRevrangeTotimestamp) Align(value string) TsRevrangeAlign {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAlign)
}

func (c TsRevrangeTotimestamp) AggregationAvg() TsRevrangeAggregationAggregationAvg {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationAvg)
}

func (c TsRevrangeTotimestamp) AggregationSum() TsRevrangeAggregationAggregationSum {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationSum)
}

func (c TsRevrangeTotimestamp) AggregationMin() TsRevrangeAggregationAggregationMin {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMin)
}

func (c TsRevrangeTotimestamp) AggregationMax() TsRevrangeAggregationAggregationMax {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationMax)
}

func (c TsRevrangeTotimestamp) AggregationRange() TsRevrangeAggregationAggregationRange {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationRange)
}

func (c TsRevrangeTotimestamp) AggregationCount() TsRevrangeAggregationAggregationCount {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationCount)
}

func (c TsRevrangeTotimestamp) AggregationFirst() TsRevrangeAggregationAggregationFirst {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationFirst)
}

func (c TsRevrangeTotimestamp) AggregationLast() TsRevrangeAggregationAggregationLast {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationLast)
}

func (c TsRevrangeTotimestamp) AggregationStdP() TsRevrangeAggregationAggregationStdP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdP)
}

func (c TsRevrangeTotimestamp) AggregationStdS() TsRevrangeAggregationAggregationStdS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationStdS)
}

func (c TsRevrangeTotimestamp) AggregationVarP() TsRevrangeAggregationAggregationVarP {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarP)
}

func (c TsRevrangeTotimestamp) AggregationVarS() TsRevrangeAggregationAggregationVarS {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationVarS)
}

func (c TsRevrangeTotimestamp) AggregationTwa() TsRevrangeAggregationAggregationTwa {
	_ = "STUB: not implemented"
	return *new(TsRevrangeAggregationAggregationTwa)
}

func (c TsRevrangeTotimestamp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }
