// Code generated DO NOT EDIT

package cmds

type TdigestAdd Incomplete

func (b Builder) TdigestAdd() (c TdigestAdd) { _ = "STUB: not implemented"; return *new(TdigestAdd) }

func (c TdigestAdd) Key(key string) TdigestAddKey {
	_ = "STUB: not implemented"
	return *new(TdigestAddKey)
}

type TdigestAddKey Incomplete

func (c TdigestAddKey) Value(value float64) TdigestAddValuesValue {
	_ = "STUB: not implemented"
	return *new(TdigestAddValuesValue)
}

type TdigestAddValuesValue Incomplete

func (c TdigestAddValuesValue) Value(value float64) TdigestAddValuesValue {
	_ = "STUB: not implemented"
	return *new(TdigestAddValuesValue)
}

func (c TdigestAddValuesValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestByrank Incomplete

func (b Builder) TdigestByrank() (c TdigestByrank) {
	_ = "STUB: not implemented"
	return *new(TdigestByrank)
}

func (c TdigestByrank) Key(key string) TdigestByrankKey {
	_ = "STUB: not implemented"
	return *new(TdigestByrankKey)
}

type TdigestByrankKey Incomplete

func (c TdigestByrankKey) Rank(rank ...float64) TdigestByrankRank {
	_ = "STUB: not implemented"
	return *new(TdigestByrankRank)
}

type TdigestByrankRank Incomplete

func (c TdigestByrankRank) Rank(rank ...float64) TdigestByrankRank {
	_ = "STUB: not implemented"
	return *new(TdigestByrankRank)
}

func (c TdigestByrankRank) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestByrevrank Incomplete

func (b Builder) TdigestByrevrank() (c TdigestByrevrank) {
	_ = "STUB: not implemented"
	return *new(TdigestByrevrank)
}

func (c TdigestByrevrank) Key(key string) TdigestByrevrankKey {
	_ = "STUB: not implemented"
	return *new(TdigestByrevrankKey)
}

type TdigestByrevrankKey Incomplete

func (c TdigestByrevrankKey) ReverseRank(reverseRank ...float64) TdigestByrevrankReverseRank {
	_ = "STUB: not implemented"
	return *new(TdigestByrevrankReverseRank)
}

type TdigestByrevrankReverseRank Incomplete

func (c TdigestByrevrankReverseRank) ReverseRank(reverseRank ...float64) TdigestByrevrankReverseRank {
	_ = "STUB: not implemented"
	return *new(TdigestByrevrankReverseRank)
}

func (c TdigestByrevrankReverseRank) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TdigestCdf Incomplete

func (b Builder) TdigestCdf() (c TdigestCdf) { _ = "STUB: not implemented"; return *new(TdigestCdf) }

func (c TdigestCdf) Key(key string) TdigestCdfKey {
	_ = "STUB: not implemented"
	return *new(TdigestCdfKey)
}

type TdigestCdfKey Incomplete

func (c TdigestCdfKey) Value(value ...float64) TdigestCdfValue {
	_ = "STUB: not implemented"
	return *new(TdigestCdfValue)
}

type TdigestCdfValue Incomplete

func (c TdigestCdfValue) Value(value ...float64) TdigestCdfValue {
	_ = "STUB: not implemented"
	return *new(TdigestCdfValue)
}

func (c TdigestCdfValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestCreate Incomplete

func (b Builder) TdigestCreate() (c TdigestCreate) {
	_ = "STUB: not implemented"
	return *new(TdigestCreate)
}

func (c TdigestCreate) Key(key string) TdigestCreateKey {
	_ = "STUB: not implemented"
	return *new(TdigestCreateKey)
}

type TdigestCreateCompression Incomplete

func (c TdigestCreateCompression) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TdigestCreateKey Incomplete

func (c TdigestCreateKey) Compression(compression int64) TdigestCreateCompression {
	_ = "STUB: not implemented"
	return *new(TdigestCreateCompression)
}

func (c TdigestCreateKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestInfo Incomplete

func (b Builder) TdigestInfo() (c TdigestInfo) { _ = "STUB: not implemented"; return *new(TdigestInfo) }

func (c TdigestInfo) Key(key string) TdigestInfoKey {
	_ = "STUB: not implemented"
	return *new(TdigestInfoKey)
}

type TdigestInfoKey Incomplete

func (c TdigestInfoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestMax Incomplete

func (b Builder) TdigestMax() (c TdigestMax) { _ = "STUB: not implemented"; return *new(TdigestMax) }

func (c TdigestMax) Key(key string) TdigestMaxKey {
	_ = "STUB: not implemented"
	return *new(TdigestMaxKey)
}

type TdigestMaxKey Incomplete

func (c TdigestMaxKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestMerge Incomplete

func (b Builder) TdigestMerge() (c TdigestMerge) {
	_ = "STUB: not implemented"
	return *new(TdigestMerge)
}

func (c TdigestMerge) DestinationKey(destinationKey string) TdigestMergeDestinationKey {
	_ = "STUB: not implemented"
	return *new(TdigestMergeDestinationKey)
}

type TdigestMergeConfigCompression Incomplete

func (c TdigestMergeConfigCompression) Override() TdigestMergeOverride {
	_ = "STUB: not implemented"
	return *new(TdigestMergeOverride)
}

func (c TdigestMergeConfigCompression) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TdigestMergeDestinationKey Incomplete

func (c TdigestMergeDestinationKey) Numkeys(numkeys int64) TdigestMergeNumkeys {
	_ = "STUB: not implemented"
	return *new(TdigestMergeNumkeys)
}

type TdigestMergeNumkeys Incomplete

func (c TdigestMergeNumkeys) SourceKey(sourceKey ...string) TdigestMergeSourceKey {
	_ = "STUB: not implemented"
	return *new(TdigestMergeSourceKey)
}

type TdigestMergeOverride Incomplete

func (c TdigestMergeOverride) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestMergeSourceKey Incomplete

func (c TdigestMergeSourceKey) SourceKey(sourceKey ...string) TdigestMergeSourceKey {
	_ = "STUB: not implemented"
	return *new(TdigestMergeSourceKey)
}

func (c TdigestMergeSourceKey) Compression(compression int64) TdigestMergeConfigCompression {
	_ = "STUB: not implemented"
	return *new(TdigestMergeConfigCompression)
}

func (c TdigestMergeSourceKey) Override() TdigestMergeOverride {
	_ = "STUB: not implemented"
	return *new(TdigestMergeOverride)
}

func (c TdigestMergeSourceKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestMin Incomplete

func (b Builder) TdigestMin() (c TdigestMin) { _ = "STUB: not implemented"; return *new(TdigestMin) }

func (c TdigestMin) Key(key string) TdigestMinKey {
	_ = "STUB: not implemented"
	return *new(TdigestMinKey)
}

type TdigestMinKey Incomplete

func (c TdigestMinKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestQuantile Incomplete

func (b Builder) TdigestQuantile() (c TdigestQuantile) {
	_ = "STUB: not implemented"
	return *new(TdigestQuantile)
}

func (c TdigestQuantile) Key(key string) TdigestQuantileKey {
	_ = "STUB: not implemented"
	return *new(TdigestQuantileKey)
}

type TdigestQuantileKey Incomplete

func (c TdigestQuantileKey) Quantile(quantile ...float64) TdigestQuantileQuantile {
	_ = "STUB: not implemented"
	return *new(TdigestQuantileQuantile)
}

type TdigestQuantileQuantile Incomplete

func (c TdigestQuantileQuantile) Quantile(quantile ...float64) TdigestQuantileQuantile {
	_ = "STUB: not implemented"
	return *new(TdigestQuantileQuantile)
}

func (c TdigestQuantileQuantile) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TdigestRank Incomplete

func (b Builder) TdigestRank() (c TdigestRank) { _ = "STUB: not implemented"; return *new(TdigestRank) }

func (c TdigestRank) Key(key string) TdigestRankKey {
	_ = "STUB: not implemented"
	return *new(TdigestRankKey)
}

type TdigestRankKey Incomplete

func (c TdigestRankKey) Value(value ...float64) TdigestRankValue {
	_ = "STUB: not implemented"
	return *new(TdigestRankValue)
}

type TdigestRankValue Incomplete

func (c TdigestRankValue) Value(value ...float64) TdigestRankValue {
	_ = "STUB: not implemented"
	return *new(TdigestRankValue)
}

func (c TdigestRankValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestReset Incomplete

func (b Builder) TdigestReset() (c TdigestReset) {
	_ = "STUB: not implemented"
	return *new(TdigestReset)
}

func (c TdigestReset) Key(key string) TdigestResetKey {
	_ = "STUB: not implemented"
	return *new(TdigestResetKey)
}

type TdigestResetKey Incomplete

func (c TdigestResetKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestRevrank Incomplete

func (b Builder) TdigestRevrank() (c TdigestRevrank) {
	_ = "STUB: not implemented"
	return *new(TdigestRevrank)
}

func (c TdigestRevrank) Key(key string) TdigestRevrankKey {
	_ = "STUB: not implemented"
	return *new(TdigestRevrankKey)
}

type TdigestRevrankKey Incomplete

func (c TdigestRevrankKey) Value(value ...float64) TdigestRevrankValue {
	_ = "STUB: not implemented"
	return *new(TdigestRevrankValue)
}

type TdigestRevrankValue Incomplete

func (c TdigestRevrankValue) Value(value ...float64) TdigestRevrankValue {
	_ = "STUB: not implemented"
	return *new(TdigestRevrankValue)
}

func (c TdigestRevrankValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TdigestTrimmedMean Incomplete

func (b Builder) TdigestTrimmedMean() (c TdigestTrimmedMean) {
	_ = "STUB: not implemented"
	return *new(TdigestTrimmedMean)
}

func (c TdigestTrimmedMean) Key(key string) TdigestTrimmedMeanKey {
	_ = "STUB: not implemented"
	return *new(TdigestTrimmedMeanKey)
}

type TdigestTrimmedMeanHighCutQuantile Incomplete

func (c TdigestTrimmedMeanHighCutQuantile) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TdigestTrimmedMeanKey Incomplete

func (c TdigestTrimmedMeanKey) LowCutQuantile(lowCutQuantile float64) TdigestTrimmedMeanLowCutQuantile {
	_ = "STUB: not implemented"
	return *new(TdigestTrimmedMeanLowCutQuantile)
}

type TdigestTrimmedMeanLowCutQuantile Incomplete

func (c TdigestTrimmedMeanLowCutQuantile) HighCutQuantile(highCutQuantile float64) TdigestTrimmedMeanHighCutQuantile {
	_ = "STUB: not implemented"
	return *new(TdigestTrimmedMeanHighCutQuantile)
}
