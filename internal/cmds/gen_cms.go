// Code generated DO NOT EDIT

package cmds

type CmsIncrby Incomplete

func (b Builder) CmsIncrby() (c CmsIncrby) { _ = "STUB: not implemented"; return *new(CmsIncrby) }

func (c CmsIncrby) Key(key string) CmsIncrbyKey {
	_ = "STUB: not implemented"
	return *new(CmsIncrbyKey)
}

type CmsIncrbyItemsIncrement Incomplete

func (c CmsIncrbyItemsIncrement) Item(item string) CmsIncrbyItemsItem {
	_ = "STUB: not implemented"
	return *new(CmsIncrbyItemsItem)
}

func (c CmsIncrbyItemsIncrement) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type CmsIncrbyItemsItem Incomplete

func (c CmsIncrbyItemsItem) Increment(increment int64) CmsIncrbyItemsIncrement {
	_ = "STUB: not implemented"
	return *new(CmsIncrbyItemsIncrement)
}

type CmsIncrbyKey Incomplete

func (c CmsIncrbyKey) Item(item string) CmsIncrbyItemsItem {
	_ = "STUB: not implemented"
	return *new(CmsIncrbyItemsItem)
}

type CmsInfo Incomplete

func (b Builder) CmsInfo() (c CmsInfo) { _ = "STUB: not implemented"; return *new(CmsInfo) }

func (c CmsInfo) Key(key string) CmsInfoKey { _ = "STUB: not implemented"; return *new(CmsInfoKey) }

type CmsInfoKey Incomplete

func (c CmsInfoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c CmsInfoKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type CmsInitbydim Incomplete

func (b Builder) CmsInitbydim() (c CmsInitbydim) {
	_ = "STUB: not implemented"
	return *new(CmsInitbydim)
}

func (c CmsInitbydim) Key(key string) CmsInitbydimKey {
	_ = "STUB: not implemented"
	return *new(CmsInitbydimKey)
}

type CmsInitbydimDepth Incomplete

func (c CmsInitbydimDepth) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CmsInitbydimKey Incomplete

func (c CmsInitbydimKey) Width(width int64) CmsInitbydimWidth {
	_ = "STUB: not implemented"
	return *new(CmsInitbydimWidth)
}

type CmsInitbydimWidth Incomplete

func (c CmsInitbydimWidth) Depth(depth int64) CmsInitbydimDepth {
	_ = "STUB: not implemented"
	return *new(CmsInitbydimDepth)
}

type CmsInitbyprob Incomplete

func (b Builder) CmsInitbyprob() (c CmsInitbyprob) {
	_ = "STUB: not implemented"
	return *new(CmsInitbyprob)
}

func (c CmsInitbyprob) Key(key string) CmsInitbyprobKey {
	_ = "STUB: not implemented"
	return *new(CmsInitbyprobKey)
}

type CmsInitbyprobError Incomplete

func (c CmsInitbyprobError) Probability(probability float64) CmsInitbyprobProbability {
	_ = "STUB: not implemented"
	return *new(CmsInitbyprobProbability)
}

type CmsInitbyprobKey Incomplete

func (c CmsInitbyprobKey) Error(error float64) CmsInitbyprobError {
	_ = "STUB: not implemented"
	return *new(CmsInitbyprobError)
}

type CmsInitbyprobProbability Incomplete

func (c CmsInitbyprobProbability) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type CmsMerge Incomplete

func (b Builder) CmsMerge() (c CmsMerge) { _ = "STUB: not implemented"; return *new(CmsMerge) }

func (c CmsMerge) Destination(destination string) CmsMergeDestination {
	_ = "STUB: not implemented"
	return *new(CmsMergeDestination)
}

type CmsMergeDestination Incomplete

func (c CmsMergeDestination) Numkeys(numkeys int64) CmsMergeNumkeys {
	_ = "STUB: not implemented"
	return *new(CmsMergeNumkeys)
}

type CmsMergeNumkeys Incomplete

func (c CmsMergeNumkeys) Source(source ...string) CmsMergeSource {
	_ = "STUB: not implemented"
	return *new(CmsMergeSource)
}

type CmsMergeSource Incomplete

func (c CmsMergeSource) Source(source ...string) CmsMergeSource {
	_ = "STUB: not implemented"
	return *new(CmsMergeSource)
}

func (c CmsMergeSource) Weights() CmsMergeWeightWeights {
	_ = "STUB: not implemented"
	return *new(CmsMergeWeightWeights)
}

func (c CmsMergeSource) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CmsMergeWeightWeight Incomplete

func (c CmsMergeWeightWeight) Weight(weight ...float64) CmsMergeWeightWeight {
	_ = "STUB: not implemented"
	return *new(CmsMergeWeightWeight)
}

func (c CmsMergeWeightWeight) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CmsMergeWeightWeights Incomplete

func (c CmsMergeWeightWeights) Weight(weight ...float64) CmsMergeWeightWeight {
	_ = "STUB: not implemented"
	return *new(CmsMergeWeightWeight)
}

type CmsQuery Incomplete

func (b Builder) CmsQuery() (c CmsQuery) { _ = "STUB: not implemented"; return *new(CmsQuery) }

func (c CmsQuery) Key(key string) CmsQueryKey { _ = "STUB: not implemented"; return *new(CmsQueryKey) }

type CmsQueryItem Incomplete

func (c CmsQueryItem) Item(item ...string) CmsQueryItem {
	_ = "STUB: not implemented"
	return *new(CmsQueryItem)
}

func (c CmsQueryItem) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c CmsQueryItem) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type CmsQueryKey Incomplete

func (c CmsQueryKey) Item(item ...string) CmsQueryItem {
	_ = "STUB: not implemented"
	return *new(CmsQueryItem)
}
