// Code generated DO NOT EDIT

package cmds

type FtAggregate Incomplete

func (b Builder) FtAggregate() (c FtAggregate) { _ = "STUB: not implemented"; return *new(FtAggregate) }

func (c FtAggregate) Index(index string) FtAggregateIndex {
	_ = "STUB: not implemented"
	return *new(FtAggregateIndex)
}

type FtAggregateAddscores Incomplete

func (c FtAggregateAddscores) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateAddscores) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAggregateCursorCount Incomplete

func (c FtAggregateCursorCount) Maxidle(idleTime int64) FtAggregateCursorMaxidle {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorMaxidle)
}

func (c FtAggregateCursorCount) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateCursorCount) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateCursorCount) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateCursorCount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateCursorMaxidle Incomplete

func (c FtAggregateCursorMaxidle) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateCursorMaxidle) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateCursorMaxidle) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateCursorMaxidle) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateCursorWithcursor Incomplete

func (c FtAggregateCursorWithcursor) Count(readSize int64) FtAggregateCursorCount {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorCount)
}

func (c FtAggregateCursorWithcursor) Maxidle(idleTime int64) FtAggregateCursorMaxidle {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorMaxidle)
}

func (c FtAggregateCursorWithcursor) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateCursorWithcursor) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateCursorWithcursor) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateCursorWithcursor) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateDialect Incomplete

func (c FtAggregateDialect) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAggregateIndex Incomplete

func (c FtAggregateIndex) Query(query string) FtAggregateQuery {
	_ = "STUB: not implemented"
	return *new(FtAggregateQuery)
}

type FtAggregateOpApplyApply Incomplete

func (c FtAggregateOpApplyApply) As(name string) FtAggregateOpApplyAs {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyAs)
}

type FtAggregateOpApplyAs Incomplete

func (c FtAggregateOpApplyAs) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpApplyAs) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpApplyAs) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpApplyAs) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpApplyAs) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpApplyAs) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpApplyAs) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpApplyAs) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpApplyAs) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpApplyAs) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpApplyAs) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpApplyAs) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAggregateOpFilter Incomplete

func (c FtAggregateOpFilter) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpFilter) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpFilter) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpFilter) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpFilter) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpFilter) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpFilter) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpFilter) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpFilter) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpFilter) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpFilter) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpFilter) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAggregateOpGroupbyGroupby Incomplete

func (c FtAggregateOpGroupbyGroupby) Property(property ...string) FtAggregateOpGroupbyProperty {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyProperty)
}

func (c FtAggregateOpGroupbyGroupby) Reduce(function string) FtAggregateOpGroupbyReduceReduce {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceReduce)
}

func (c FtAggregateOpGroupbyGroupby) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpGroupbyGroupby) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpGroupbyGroupby) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpGroupbyGroupby) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpGroupbyGroupby) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpGroupbyGroupby) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpGroupbyGroupby) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpGroupbyGroupby) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpGroupbyGroupby) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpGroupbyGroupby) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpGroupbyGroupby) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpGroupbyGroupby) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpGroupbyProperty Incomplete

func (c FtAggregateOpGroupbyProperty) Property(property ...string) FtAggregateOpGroupbyProperty {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyProperty)
}

func (c FtAggregateOpGroupbyProperty) Reduce(function string) FtAggregateOpGroupbyReduceReduce {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceReduce)
}

func (c FtAggregateOpGroupbyProperty) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpGroupbyProperty) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpGroupbyProperty) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpGroupbyProperty) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpGroupbyProperty) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpGroupbyProperty) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpGroupbyProperty) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpGroupbyProperty) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpGroupbyProperty) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpGroupbyProperty) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpGroupbyProperty) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpGroupbyProperty) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpGroupbyReduceArg Incomplete

func (c FtAggregateOpGroupbyReduceArg) Arg(arg ...string) FtAggregateOpGroupbyReduceArg {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceArg)
}

func (c FtAggregateOpGroupbyReduceArg) As(name string) FtAggregateOpGroupbyReduceAs {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceAs)
}

func (c FtAggregateOpGroupbyReduceArg) By(by string) FtAggregateOpGroupbyReduceBy {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceBy)
}

func (c FtAggregateOpGroupbyReduceArg) Asc() FtAggregateOpGroupbyReduceOrderAsc {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceOrderAsc)
}

func (c FtAggregateOpGroupbyReduceArg) Desc() FtAggregateOpGroupbyReduceOrderDesc {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceOrderDesc)
}

func (c FtAggregateOpGroupbyReduceArg) Reduce(function string) FtAggregateOpGroupbyReduceReduce {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceReduce)
}

func (c FtAggregateOpGroupbyReduceArg) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpGroupbyReduceArg) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpGroupbyReduceArg) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpGroupbyReduceArg) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpGroupbyReduceArg) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpGroupbyReduceArg) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpGroupbyReduceArg) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpGroupbyReduceArg) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpGroupbyReduceArg) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpGroupbyReduceArg) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpGroupbyReduceArg) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpGroupbyReduceArg) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpGroupbyReduceAs Incomplete

func (c FtAggregateOpGroupbyReduceAs) By(by string) FtAggregateOpGroupbyReduceBy {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceBy)
}

func (c FtAggregateOpGroupbyReduceAs) Asc() FtAggregateOpGroupbyReduceOrderAsc {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceOrderAsc)
}

func (c FtAggregateOpGroupbyReduceAs) Desc() FtAggregateOpGroupbyReduceOrderDesc {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceOrderDesc)
}

func (c FtAggregateOpGroupbyReduceAs) Reduce(function string) FtAggregateOpGroupbyReduceReduce {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceReduce)
}

func (c FtAggregateOpGroupbyReduceAs) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpGroupbyReduceAs) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpGroupbyReduceAs) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpGroupbyReduceAs) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpGroupbyReduceAs) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpGroupbyReduceAs) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpGroupbyReduceAs) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpGroupbyReduceAs) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpGroupbyReduceAs) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpGroupbyReduceAs) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpGroupbyReduceAs) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpGroupbyReduceAs) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpGroupbyReduceBy Incomplete

func (c FtAggregateOpGroupbyReduceBy) Asc() FtAggregateOpGroupbyReduceOrderAsc {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceOrderAsc)
}

func (c FtAggregateOpGroupbyReduceBy) Desc() FtAggregateOpGroupbyReduceOrderDesc {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceOrderDesc)
}

func (c FtAggregateOpGroupbyReduceBy) Reduce(function string) FtAggregateOpGroupbyReduceReduce {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceReduce)
}

func (c FtAggregateOpGroupbyReduceBy) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpGroupbyReduceBy) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpGroupbyReduceBy) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpGroupbyReduceBy) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpGroupbyReduceBy) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpGroupbyReduceBy) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpGroupbyReduceBy) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpGroupbyReduceBy) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpGroupbyReduceBy) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpGroupbyReduceBy) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpGroupbyReduceBy) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpGroupbyReduceBy) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpGroupbyReduceNargs Incomplete

func (c FtAggregateOpGroupbyReduceNargs) Arg(arg ...string) FtAggregateOpGroupbyReduceArg {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceArg)
}

func (c FtAggregateOpGroupbyReduceNargs) As(name string) FtAggregateOpGroupbyReduceAs {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceAs)
}

func (c FtAggregateOpGroupbyReduceNargs) By(by string) FtAggregateOpGroupbyReduceBy {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceBy)
}

func (c FtAggregateOpGroupbyReduceNargs) Asc() FtAggregateOpGroupbyReduceOrderAsc {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceOrderAsc)
}

func (c FtAggregateOpGroupbyReduceNargs) Desc() FtAggregateOpGroupbyReduceOrderDesc {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceOrderDesc)
}

func (c FtAggregateOpGroupbyReduceNargs) Reduce(function string) FtAggregateOpGroupbyReduceReduce {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceReduce)
}

func (c FtAggregateOpGroupbyReduceNargs) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpGroupbyReduceNargs) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpGroupbyReduceNargs) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpGroupbyReduceNargs) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpGroupbyReduceNargs) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpGroupbyReduceNargs) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpGroupbyReduceNargs) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpGroupbyReduceNargs) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpGroupbyReduceNargs) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpGroupbyReduceNargs) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpGroupbyReduceNargs) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpGroupbyReduceNargs) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpGroupbyReduceOrderAsc Incomplete

func (c FtAggregateOpGroupbyReduceOrderAsc) Reduce(function string) FtAggregateOpGroupbyReduceReduce {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceReduce)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpGroupbyReduceOrderAsc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpGroupbyReduceOrderDesc Incomplete

func (c FtAggregateOpGroupbyReduceOrderDesc) Reduce(function string) FtAggregateOpGroupbyReduceReduce {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceReduce)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpGroupbyReduceOrderDesc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpGroupbyReduceReduce Incomplete

func (c FtAggregateOpGroupbyReduceReduce) Nargs(nargs int64) FtAggregateOpGroupbyReduceNargs {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyReduceNargs)
}

type FtAggregateOpLimitLimit Incomplete

func (c FtAggregateOpLimitLimit) OffsetNum(offset int64, num int64) FtAggregateOpLimitOffsetNum {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitOffsetNum)
}

type FtAggregateOpLimitOffsetNum Incomplete

func (c FtAggregateOpLimitOffsetNum) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpLimitOffsetNum) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpLimitOffsetNum) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpLimitOffsetNum) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpLimitOffsetNum) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpLimitOffsetNum) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpLimitOffsetNum) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpLimitOffsetNum) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpLimitOffsetNum) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpLimitOffsetNum) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpLimitOffsetNum) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpLimitOffsetNum) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpLoadField Incomplete

func (c FtAggregateOpLoadField) Field(field ...string) FtAggregateOpLoadField {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadField)
}

func (c FtAggregateOpLoadField) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpLoadField) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpLoadField) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpLoadField) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpLoadField) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpLoadField) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpLoadField) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpLoadField) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpLoadField) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpLoadField) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpLoadField) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpLoadField) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpLoadLoad Incomplete

func (c FtAggregateOpLoadLoad) Field(field ...string) FtAggregateOpLoadField {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadField)
}

type FtAggregateOpLoadallLoadAll Incomplete

func (c FtAggregateOpLoadallLoadAll) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpLoadallLoadAll) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpLoadallLoadAll) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpLoadallLoadAll) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpLoadallLoadAll) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpLoadallLoadAll) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpLoadallLoadAll) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpLoadallLoadAll) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpLoadallLoadAll) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpLoadallLoadAll) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpLoadallLoadAll) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpLoadallLoadAll) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpSortbyFieldsOrderAsc Incomplete

func (c FtAggregateOpSortbyFieldsOrderAsc) Property(property string) FtAggregateOpSortbyFieldsProperty {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyFieldsProperty)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Max(num int64) FtAggregateOpSortbyMax {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyMax)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Withcount() FtAggregateOpSortbyWithcount {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyWithcount)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpSortbyFieldsOrderAsc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpSortbyFieldsOrderDesc Incomplete

func (c FtAggregateOpSortbyFieldsOrderDesc) Property(property string) FtAggregateOpSortbyFieldsProperty {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyFieldsProperty)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Max(num int64) FtAggregateOpSortbyMax {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyMax)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Withcount() FtAggregateOpSortbyWithcount {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyWithcount)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpSortbyFieldsOrderDesc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpSortbyFieldsProperty Incomplete

func (c FtAggregateOpSortbyFieldsProperty) Asc() FtAggregateOpSortbyFieldsOrderAsc {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyFieldsOrderAsc)
}

func (c FtAggregateOpSortbyFieldsProperty) Desc() FtAggregateOpSortbyFieldsOrderDesc {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyFieldsOrderDesc)
}

func (c FtAggregateOpSortbyFieldsProperty) Property(property string) FtAggregateOpSortbyFieldsProperty {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyFieldsProperty)
}

func (c FtAggregateOpSortbyFieldsProperty) Max(num int64) FtAggregateOpSortbyMax {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyMax)
}

func (c FtAggregateOpSortbyFieldsProperty) Withcount() FtAggregateOpSortbyWithcount {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyWithcount)
}

func (c FtAggregateOpSortbyFieldsProperty) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpSortbyFieldsProperty) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpSortbyFieldsProperty) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpSortbyFieldsProperty) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpSortbyFieldsProperty) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpSortbyFieldsProperty) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpSortbyFieldsProperty) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpSortbyFieldsProperty) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpSortbyFieldsProperty) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpSortbyFieldsProperty) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpSortbyFieldsProperty) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpSortbyFieldsProperty) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpSortbyMax Incomplete

func (c FtAggregateOpSortbyMax) Withcount() FtAggregateOpSortbyWithcount {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyWithcount)
}

func (c FtAggregateOpSortbyMax) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpSortbyMax) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpSortbyMax) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpSortbyMax) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpSortbyMax) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpSortbyMax) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpSortbyMax) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpSortbyMax) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpSortbyMax) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpSortbyMax) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpSortbyMax) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpSortbyMax) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpSortbySortby Incomplete

func (c FtAggregateOpSortbySortby) Property(property string) FtAggregateOpSortbyFieldsProperty {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyFieldsProperty)
}

func (c FtAggregateOpSortbySortby) Max(num int64) FtAggregateOpSortbyMax {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyMax)
}

func (c FtAggregateOpSortbySortby) Withcount() FtAggregateOpSortbyWithcount {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbyWithcount)
}

func (c FtAggregateOpSortbySortby) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpSortbySortby) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpSortbySortby) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpSortbySortby) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpSortbySortby) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpSortbySortby) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpSortbySortby) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpSortbySortby) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpSortbySortby) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpSortbySortby) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpSortbySortby) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpSortbySortby) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateOpSortbyWithcount Incomplete

func (c FtAggregateOpSortbyWithcount) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateOpSortbyWithcount) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateOpSortbyWithcount) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateOpSortbyWithcount) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateOpSortbyWithcount) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateOpSortbyWithcount) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateOpSortbyWithcount) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateOpSortbyWithcount) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateOpSortbyWithcount) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateOpSortbyWithcount) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateOpSortbyWithcount) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateOpSortbyWithcount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateParamsNameValue Incomplete

func (c FtAggregateParamsNameValue) NameValue(name string, value string) FtAggregateParamsNameValue {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsNameValue)
}

func (c FtAggregateParamsNameValue) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateParamsNameValue) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateParamsNameValue) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtAggregateParamsNargs Incomplete

func (c FtAggregateParamsNargs) NameValue() FtAggregateParamsNameValue {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsNameValue)
}

type FtAggregateParamsParams Incomplete

func (c FtAggregateParamsParams) Nargs(nargs int64) FtAggregateParamsNargs {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsNargs)
}

type FtAggregateQuery Incomplete

func (c FtAggregateQuery) Verbatim() FtAggregateVerbatim {
	_ = "STUB: not implemented"
	return *new(FtAggregateVerbatim)
}

func (c FtAggregateQuery) Scorer(scorer string) FtAggregateScorer {
	_ = "STUB: not implemented"
	return *new(FtAggregateScorer)
}

func (c FtAggregateQuery) Timeout(timeout int64) FtAggregateTimeout {
	_ = "STUB: not implemented"
	return *new(FtAggregateTimeout)
}

func (c FtAggregateQuery) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateQuery) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateQuery) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateQuery) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateQuery) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateQuery) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateQuery) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateQuery) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateQuery) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateQuery) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateQuery) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateQuery) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAggregateScorer Incomplete

func (c FtAggregateScorer) Timeout(timeout int64) FtAggregateTimeout {
	_ = "STUB: not implemented"
	return *new(FtAggregateTimeout)
}

func (c FtAggregateScorer) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateScorer) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateScorer) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateScorer) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateScorer) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateScorer) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateScorer) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateScorer) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateScorer) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateScorer) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateScorer) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateScorer) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAggregateTimeout Incomplete

func (c FtAggregateTimeout) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateTimeout) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateTimeout) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateTimeout) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateTimeout) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateTimeout) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateTimeout) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateTimeout) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateTimeout) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateTimeout) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateTimeout) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAggregateVerbatim Incomplete

func (c FtAggregateVerbatim) Scorer(scorer string) FtAggregateScorer {
	_ = "STUB: not implemented"
	return *new(FtAggregateScorer)
}

func (c FtAggregateVerbatim) Timeout(timeout int64) FtAggregateTimeout {
	_ = "STUB: not implemented"
	return *new(FtAggregateTimeout)
}

func (c FtAggregateVerbatim) LoadAll() FtAggregateOpLoadallLoadAll {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadallLoadAll)
}

func (c FtAggregateVerbatim) Load(count int64) FtAggregateOpLoadLoad {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLoadLoad)
}

func (c FtAggregateVerbatim) Apply(expression string) FtAggregateOpApplyApply {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpApplyApply)
}

func (c FtAggregateVerbatim) Groupby(nargs int64) FtAggregateOpGroupbyGroupby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpGroupbyGroupby)
}

func (c FtAggregateVerbatim) Sortby(nargs int64) FtAggregateOpSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpSortbySortby)
}

func (c FtAggregateVerbatim) Limit() FtAggregateOpLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpLimitLimit)
}

func (c FtAggregateVerbatim) Filter(filter string) FtAggregateOpFilter {
	_ = "STUB: not implemented"
	return *new(FtAggregateOpFilter)
}

func (c FtAggregateVerbatim) Withcursor() FtAggregateCursorWithcursor {
	_ = "STUB: not implemented"
	return *new(FtAggregateCursorWithcursor)
}

func (c FtAggregateVerbatim) Params() FtAggregateParamsParams {
	_ = "STUB: not implemented"
	return *new(FtAggregateParamsParams)
}

func (c FtAggregateVerbatim) Addscores() FtAggregateAddscores {
	_ = "STUB: not implemented"
	return *new(FtAggregateAddscores)
}

func (c FtAggregateVerbatim) Dialect(dialect int64) FtAggregateDialect {
	_ = "STUB: not implemented"
	return *new(FtAggregateDialect)
}

func (c FtAggregateVerbatim) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAliasadd Incomplete

func (b Builder) FtAliasadd() (c FtAliasadd) { _ = "STUB: not implemented"; return *new(FtAliasadd) }

func (c FtAliasadd) Alias(alias string) FtAliasaddAlias {
	_ = "STUB: not implemented"
	return *new(FtAliasaddAlias)
}

type FtAliasaddAlias Incomplete

func (c FtAliasaddAlias) Index(index string) FtAliasaddIndex {
	_ = "STUB: not implemented"
	return *new(FtAliasaddIndex)
}

type FtAliasaddIndex Incomplete

func (c FtAliasaddIndex) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAliasdel Incomplete

func (b Builder) FtAliasdel() (c FtAliasdel) { _ = "STUB: not implemented"; return *new(FtAliasdel) }

func (c FtAliasdel) Alias(alias string) FtAliasdelAlias {
	_ = "STUB: not implemented"
	return *new(FtAliasdelAlias)
}

type FtAliasdelAlias Incomplete

func (c FtAliasdelAlias) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAliasupdate Incomplete

func (b Builder) FtAliasupdate() (c FtAliasupdate) {
	_ = "STUB: not implemented"
	return *new(FtAliasupdate)
}

func (c FtAliasupdate) Alias(alias string) FtAliasupdateAlias {
	_ = "STUB: not implemented"
	return *new(FtAliasupdateAlias)
}

type FtAliasupdateAlias Incomplete

func (c FtAliasupdateAlias) Index(index string) FtAliasupdateIndex {
	_ = "STUB: not implemented"
	return *new(FtAliasupdateIndex)
}

type FtAliasupdateIndex Incomplete

func (c FtAliasupdateIndex) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAlter Incomplete

func (b Builder) FtAlter() (c FtAlter) { _ = "STUB: not implemented"; return *new(FtAlter) }

func (c FtAlter) Index(index string) FtAlterIndex {
	_ = "STUB: not implemented"
	return *new(FtAlterIndex)
}

type FtAlterAdd Incomplete

func (c FtAlterAdd) Field(field string) FtAlterField {
	_ = "STUB: not implemented"
	return *new(FtAlterField)
}

type FtAlterField Incomplete

func (c FtAlterField) Options(options ...string) FtAlterOptions {
	_ = "STUB: not implemented"
	return *new(FtAlterOptions)
}

type FtAlterIndex Incomplete

func (c FtAlterIndex) Skipinitialscan() FtAlterSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtAlterSkipinitialscan)
}

func (c FtAlterIndex) Schema() FtAlterSchema { _ = "STUB: not implemented"; return *new(FtAlterSchema) }

type FtAlterOptions Incomplete

func (c FtAlterOptions) Options(options ...string) FtAlterOptions {
	_ = "STUB: not implemented"
	return *new(FtAlterOptions)
}

func (c FtAlterOptions) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtAlterSchema Incomplete

func (c FtAlterSchema) Add() FtAlterAdd { _ = "STUB: not implemented"; return *new(FtAlterAdd) }

type FtAlterSkipinitialscan Incomplete

func (c FtAlterSkipinitialscan) Schema() FtAlterSchema {
	_ = "STUB: not implemented"
	return *new(FtAlterSchema)
}

type FtConfigGet Incomplete

func (b Builder) FtConfigGet() (c FtConfigGet) { _ = "STUB: not implemented"; return *new(FtConfigGet) }

func (c FtConfigGet) Option(option string) FtConfigGetOption {
	_ = "STUB: not implemented"
	return *new(FtConfigGetOption)
}

type FtConfigGetOption Incomplete

func (c FtConfigGetOption) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtConfigHelp Incomplete

func (b Builder) FtConfigHelp() (c FtConfigHelp) {
	_ = "STUB: not implemented"
	return *new(FtConfigHelp)
}

func (c FtConfigHelp) Option(option string) FtConfigHelpOption {
	_ = "STUB: not implemented"
	return *new(FtConfigHelpOption)
}

type FtConfigHelpOption Incomplete

func (c FtConfigHelpOption) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtConfigSet Incomplete

func (b Builder) FtConfigSet() (c FtConfigSet) { _ = "STUB: not implemented"; return *new(FtConfigSet) }

func (c FtConfigSet) Option(option string) FtConfigSetOption {
	_ = "STUB: not implemented"
	return *new(FtConfigSetOption)
}

type FtConfigSetOption Incomplete

func (c FtConfigSetOption) Value(value string) FtConfigSetValue {
	_ = "STUB: not implemented"
	return *new(FtConfigSetValue)
}

type FtConfigSetValue Incomplete

func (c FtConfigSetValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtCreate Incomplete

func (b Builder) FtCreate() (c FtCreate) { _ = "STUB: not implemented"; return *new(FtCreate) }

func (c FtCreate) Index(index string) FtCreateIndex {
	_ = "STUB: not implemented"
	return *new(FtCreateIndex)
}

type FtCreateFieldAs Incomplete

func (c FtCreateFieldAs) Text() FtCreateFieldFieldTypeText {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeText)
}

func (c FtCreateFieldAs) Tag() FtCreateFieldFieldTypeTag {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeTag)
}

func (c FtCreateFieldAs) Numeric() FtCreateFieldFieldTypeNumeric {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeNumeric)
}

func (c FtCreateFieldAs) Geo() FtCreateFieldFieldTypeGeo {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeGeo)
}

func (c FtCreateFieldAs) Vector(algo string, nargs int64, args ...string) FtCreateFieldFieldTypeVector {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeVector)
}

func (c FtCreateFieldAs) Geoshape() FtCreateFieldFieldTypeGeoshape {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeGeoshape)
}

type FtCreateFieldFieldName Incomplete

func (c FtCreateFieldFieldName) As(alias string) FtCreateFieldAs {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldAs)
}

func (c FtCreateFieldFieldName) Text() FtCreateFieldFieldTypeText {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeText)
}

func (c FtCreateFieldFieldName) Tag() FtCreateFieldFieldTypeTag {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeTag)
}

func (c FtCreateFieldFieldName) Numeric() FtCreateFieldFieldTypeNumeric {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeNumeric)
}

func (c FtCreateFieldFieldName) Geo() FtCreateFieldFieldTypeGeo {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeGeo)
}

func (c FtCreateFieldFieldName) Vector(algo string, nargs int64, args ...string) FtCreateFieldFieldTypeVector {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeVector)
}

func (c FtCreateFieldFieldName) Geoshape() FtCreateFieldFieldTypeGeoshape {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldTypeGeoshape)
}

type FtCreateFieldFieldTypeGeo Incomplete

func (c FtCreateFieldFieldTypeGeo) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldFieldTypeGeo) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldFieldTypeGeo) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldFieldTypeGeo) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldFieldTypeGeo) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldFieldTypeGeo) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldFieldTypeGeo) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldFieldTypeGeo) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldFieldTypeGeo) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldFieldTypeGeo) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldFieldTypeGeo) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldFieldTypeGeo) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldFieldTypeGeoshape Incomplete

func (c FtCreateFieldFieldTypeGeoshape) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldFieldTypeGeoshape) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldFieldTypeGeoshape) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldFieldTypeGeoshape) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldFieldTypeGeoshape) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldFieldTypeGeoshape) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldFieldTypeGeoshape) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldFieldTypeGeoshape) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldFieldTypeGeoshape) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldFieldTypeGeoshape) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldFieldTypeGeoshape) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldFieldTypeGeoshape) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldFieldTypeNumeric Incomplete

func (c FtCreateFieldFieldTypeNumeric) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldFieldTypeNumeric) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldFieldTypeNumeric) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldFieldTypeNumeric) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldFieldTypeNumeric) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldFieldTypeNumeric) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldFieldTypeNumeric) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldFieldTypeNumeric) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldFieldTypeNumeric) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldFieldTypeNumeric) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldFieldTypeNumeric) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldFieldTypeNumeric) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldFieldTypeTag Incomplete

func (c FtCreateFieldFieldTypeTag) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldFieldTypeTag) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldFieldTypeTag) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldFieldTypeTag) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldFieldTypeTag) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldFieldTypeTag) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldFieldTypeTag) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldFieldTypeTag) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldFieldTypeTag) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldFieldTypeTag) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldFieldTypeTag) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldFieldTypeTag) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldFieldTypeText Incomplete

func (c FtCreateFieldFieldTypeText) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldFieldTypeText) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldFieldTypeText) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldFieldTypeText) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldFieldTypeText) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldFieldTypeText) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldFieldTypeText) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldFieldTypeText) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldFieldTypeText) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldFieldTypeText) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldFieldTypeText) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldFieldTypeText) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldFieldTypeVector Incomplete

func (c FtCreateFieldFieldTypeVector) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldFieldTypeVector) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldFieldTypeVector) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldFieldTypeVector) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldFieldTypeVector) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldFieldTypeVector) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldFieldTypeVector) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldFieldTypeVector) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldFieldTypeVector) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldFieldTypeVector) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldFieldTypeVector) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldFieldTypeVector) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldOptionCasesensitive Incomplete

func (c FtCreateFieldOptionCasesensitive) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldOptionCasesensitive) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldOptionCasesensitive) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldOptionCasesensitive) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldOptionCasesensitive) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldOptionCasesensitive) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldOptionCasesensitive) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldOptionCasesensitive) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldOptionCasesensitive) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldOptionCasesensitive) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldOptionCasesensitive) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldOptionCasesensitive) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldOptionIndexempty Incomplete

func (c FtCreateFieldOptionIndexempty) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldOptionIndexempty) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldOptionIndexempty) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldOptionIndexempty) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldOptionIndexempty) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldOptionIndexempty) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldOptionIndexempty) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldOptionIndexempty) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldOptionIndexempty) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldOptionIndexempty) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldOptionIndexempty) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldOptionIndexempty) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldOptionIndexmissing Incomplete

func (c FtCreateFieldOptionIndexmissing) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldOptionIndexmissing) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldOptionIndexmissing) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldOptionIndexmissing) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldOptionIndexmissing) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldOptionIndexmissing) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldOptionIndexmissing) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldOptionIndexmissing) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldOptionIndexmissing) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldOptionIndexmissing) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldOptionIndexmissing) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldOptionIndexmissing) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldOptionNoindex Incomplete

func (c FtCreateFieldOptionNoindex) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldOptionNoindex) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldOptionNoindex) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldOptionNoindex) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldOptionNoindex) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldOptionNoindex) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldOptionNoindex) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldOptionNoindex) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldOptionNoindex) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldOptionNoindex) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldOptionNoindex) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldOptionNoindex) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldOptionNostem Incomplete

func (c FtCreateFieldOptionNostem) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldOptionNostem) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldOptionNostem) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldOptionNostem) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldOptionNostem) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldOptionNostem) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldOptionNostem) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldOptionNostem) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldOptionNostem) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldOptionNostem) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldOptionNostem) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldOptionNostem) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldOptionPhonetic Incomplete

func (c FtCreateFieldOptionPhonetic) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldOptionPhonetic) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldOptionPhonetic) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldOptionPhonetic) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldOptionPhonetic) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldOptionPhonetic) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldOptionPhonetic) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldOptionPhonetic) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldOptionPhonetic) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldOptionPhonetic) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldOptionPhonetic) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldOptionPhonetic) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldOptionSeparator Incomplete

func (c FtCreateFieldOptionSeparator) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldOptionSeparator) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldOptionSeparator) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldOptionSeparator) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldOptionSeparator) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldOptionSeparator) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldOptionSeparator) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldOptionSeparator) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldOptionSeparator) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldOptionSeparator) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldOptionSeparator) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldOptionSeparator) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldOptionSortableSortable Incomplete

func (c FtCreateFieldOptionSortableSortable) Unf() FtCreateFieldOptionSortableUnf {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableUnf)
}

func (c FtCreateFieldOptionSortableSortable) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldOptionSortableSortable) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldOptionSortableSortable) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldOptionSortableSortable) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldOptionSortableSortable) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldOptionSortableSortable) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldOptionSortableSortable) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldOptionSortableSortable) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldOptionSortableSortable) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldOptionSortableSortable) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldOptionSortableSortable) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldOptionSortableSortable) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldOptionSortableUnf Incomplete

func (c FtCreateFieldOptionSortableUnf) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldOptionSortableUnf) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldOptionSortableUnf) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldOptionSortableUnf) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldOptionSortableUnf) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldOptionSortableUnf) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldOptionSortableUnf) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldOptionSortableUnf) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldOptionSortableUnf) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldOptionSortableUnf) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldOptionSortableUnf) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldOptionSortableUnf) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldOptionWeight Incomplete

func (c FtCreateFieldOptionWeight) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldOptionWeight) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldOptionWeight) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldOptionWeight) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldOptionWeight) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldOptionWeight) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldOptionWeight) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldOptionWeight) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldOptionWeight) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldOptionWeight) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldOptionWeight) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldOptionWeight) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFieldOptionWithsuffixtrie Incomplete

func (c FtCreateFieldOptionWithsuffixtrie) Indexempty() FtCreateFieldOptionIndexempty {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexempty)
}

func (c FtCreateFieldOptionWithsuffixtrie) Indexmissing() FtCreateFieldOptionIndexmissing {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionIndexmissing)
}

func (c FtCreateFieldOptionWithsuffixtrie) Sortable() FtCreateFieldOptionSortableSortable {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSortableSortable)
}

func (c FtCreateFieldOptionWithsuffixtrie) Noindex() FtCreateFieldOptionNoindex {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNoindex)
}

func (c FtCreateFieldOptionWithsuffixtrie) Nostem() FtCreateFieldOptionNostem {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionNostem)
}

func (c FtCreateFieldOptionWithsuffixtrie) Phonetic(phonetic string) FtCreateFieldOptionPhonetic {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionPhonetic)
}

func (c FtCreateFieldOptionWithsuffixtrie) Weight(weight float64) FtCreateFieldOptionWeight {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWeight)
}

func (c FtCreateFieldOptionWithsuffixtrie) Separator(separator string) FtCreateFieldOptionSeparator {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionSeparator)
}

func (c FtCreateFieldOptionWithsuffixtrie) Casesensitive() FtCreateFieldOptionCasesensitive {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionCasesensitive)
}

func (c FtCreateFieldOptionWithsuffixtrie) Withsuffixtrie() FtCreateFieldOptionWithsuffixtrie {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldOptionWithsuffixtrie)
}

func (c FtCreateFieldOptionWithsuffixtrie) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

func (c FtCreateFieldOptionWithsuffixtrie) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtCreateFilter Incomplete

func (c FtCreateFilter) Language(defaultLang string) FtCreateLanguage {
	_ = "STUB: not implemented"
	return *new(FtCreateLanguage)
}

func (c FtCreateFilter) LanguageField(langAttribute string) FtCreateLanguageField {
	_ = "STUB: not implemented"
	return *new(FtCreateLanguageField)
}

func (c FtCreateFilter) Score(defaultScore float64) FtCreateScore {
	_ = "STUB: not implemented"
	return *new(FtCreateScore)
}

func (c FtCreateFilter) ScoreField(scoreAttribute string) FtCreateScoreField {
	_ = "STUB: not implemented"
	return *new(FtCreateScoreField)
}

func (c FtCreateFilter) PayloadField(payloadAttribute string) FtCreatePayloadField {
	_ = "STUB: not implemented"
	return *new(FtCreatePayloadField)
}

func (c FtCreateFilter) Maxtextfields() FtCreateMaxtextfields {
	_ = "STUB: not implemented"
	return *new(FtCreateMaxtextfields)
}

func (c FtCreateFilter) Temporary(seconds float64) FtCreateTemporary {
	_ = "STUB: not implemented"
	return *new(FtCreateTemporary)
}

func (c FtCreateFilter) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreateFilter) Nohl() FtCreateNohl { _ = "STUB: not implemented"; return *new(FtCreateNohl) }

func (c FtCreateFilter) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateFilter) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateFilter) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateFilter) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateFilter) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateIndex Incomplete

func (c FtCreateIndex) OnHash() FtCreateOnHash {
	_ = "STUB: not implemented"
	return *new(FtCreateOnHash)
}

func (c FtCreateIndex) OnJson() FtCreateOnJson {
	_ = "STUB: not implemented"
	return *new(FtCreateOnJson)
}

func (c FtCreateIndex) Prefix(count int64) FtCreatePrefixCount {
	_ = "STUB: not implemented"
	return *new(FtCreatePrefixCount)
}

func (c FtCreateIndex) Filter(filter string) FtCreateFilter {
	_ = "STUB: not implemented"
	return *new(FtCreateFilter)
}

func (c FtCreateIndex) Language(defaultLang string) FtCreateLanguage {
	_ = "STUB: not implemented"
	return *new(FtCreateLanguage)
}

func (c FtCreateIndex) LanguageField(langAttribute string) FtCreateLanguageField {
	_ = "STUB: not implemented"
	return *new(FtCreateLanguageField)
}

func (c FtCreateIndex) Score(defaultScore float64) FtCreateScore {
	_ = "STUB: not implemented"
	return *new(FtCreateScore)
}

func (c FtCreateIndex) ScoreField(scoreAttribute string) FtCreateScoreField {
	_ = "STUB: not implemented"
	return *new(FtCreateScoreField)
}

func (c FtCreateIndex) PayloadField(payloadAttribute string) FtCreatePayloadField {
	_ = "STUB: not implemented"
	return *new(FtCreatePayloadField)
}

func (c FtCreateIndex) Maxtextfields() FtCreateMaxtextfields {
	_ = "STUB: not implemented"
	return *new(FtCreateMaxtextfields)
}

func (c FtCreateIndex) Temporary(seconds float64) FtCreateTemporary {
	_ = "STUB: not implemented"
	return *new(FtCreateTemporary)
}

func (c FtCreateIndex) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreateIndex) Nohl() FtCreateNohl { _ = "STUB: not implemented"; return *new(FtCreateNohl) }

func (c FtCreateIndex) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateIndex) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateIndex) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateIndex) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateIndex) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateLanguage Incomplete

func (c FtCreateLanguage) LanguageField(langAttribute string) FtCreateLanguageField {
	_ = "STUB: not implemented"
	return *new(FtCreateLanguageField)
}

func (c FtCreateLanguage) Score(defaultScore float64) FtCreateScore {
	_ = "STUB: not implemented"
	return *new(FtCreateScore)
}

func (c FtCreateLanguage) ScoreField(scoreAttribute string) FtCreateScoreField {
	_ = "STUB: not implemented"
	return *new(FtCreateScoreField)
}

func (c FtCreateLanguage) PayloadField(payloadAttribute string) FtCreatePayloadField {
	_ = "STUB: not implemented"
	return *new(FtCreatePayloadField)
}

func (c FtCreateLanguage) Maxtextfields() FtCreateMaxtextfields {
	_ = "STUB: not implemented"
	return *new(FtCreateMaxtextfields)
}

func (c FtCreateLanguage) Temporary(seconds float64) FtCreateTemporary {
	_ = "STUB: not implemented"
	return *new(FtCreateTemporary)
}

func (c FtCreateLanguage) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreateLanguage) Nohl() FtCreateNohl { _ = "STUB: not implemented"; return *new(FtCreateNohl) }

func (c FtCreateLanguage) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateLanguage) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateLanguage) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateLanguage) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateLanguage) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateLanguageField Incomplete

func (c FtCreateLanguageField) Score(defaultScore float64) FtCreateScore {
	_ = "STUB: not implemented"
	return *new(FtCreateScore)
}

func (c FtCreateLanguageField) ScoreField(scoreAttribute string) FtCreateScoreField {
	_ = "STUB: not implemented"
	return *new(FtCreateScoreField)
}

func (c FtCreateLanguageField) PayloadField(payloadAttribute string) FtCreatePayloadField {
	_ = "STUB: not implemented"
	return *new(FtCreatePayloadField)
}

func (c FtCreateLanguageField) Maxtextfields() FtCreateMaxtextfields {
	_ = "STUB: not implemented"
	return *new(FtCreateMaxtextfields)
}

func (c FtCreateLanguageField) Temporary(seconds float64) FtCreateTemporary {
	_ = "STUB: not implemented"
	return *new(FtCreateTemporary)
}

func (c FtCreateLanguageField) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreateLanguageField) Nohl() FtCreateNohl {
	_ = "STUB: not implemented"
	return *new(FtCreateNohl)
}

func (c FtCreateLanguageField) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateLanguageField) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateLanguageField) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateLanguageField) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateLanguageField) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateMaxtextfields Incomplete

func (c FtCreateMaxtextfields) Temporary(seconds float64) FtCreateTemporary {
	_ = "STUB: not implemented"
	return *new(FtCreateTemporary)
}

func (c FtCreateMaxtextfields) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreateMaxtextfields) Nohl() FtCreateNohl {
	_ = "STUB: not implemented"
	return *new(FtCreateNohl)
}

func (c FtCreateMaxtextfields) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateMaxtextfields) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateMaxtextfields) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateMaxtextfields) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateMaxtextfields) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateNofields Incomplete

func (c FtCreateNofields) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateNofields) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateNofields) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateNofields) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateNofreqs Incomplete

func (c FtCreateNofreqs) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateNofreqs) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateNofreqs) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateNohl Incomplete

func (c FtCreateNohl) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateNohl) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateNohl) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateNohl) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateNohl) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateNooffsets Incomplete

func (c FtCreateNooffsets) Nohl() FtCreateNohl {
	_ = "STUB: not implemented"
	return *new(FtCreateNohl)
}

func (c FtCreateNooffsets) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateNooffsets) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateNooffsets) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateNooffsets) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateNooffsets) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateOnHash Incomplete

func (c FtCreateOnHash) Prefix(count int64) FtCreatePrefixCount {
	_ = "STUB: not implemented"
	return *new(FtCreatePrefixCount)
}

func (c FtCreateOnHash) Filter(filter string) FtCreateFilter {
	_ = "STUB: not implemented"
	return *new(FtCreateFilter)
}

func (c FtCreateOnHash) Language(defaultLang string) FtCreateLanguage {
	_ = "STUB: not implemented"
	return *new(FtCreateLanguage)
}

func (c FtCreateOnHash) LanguageField(langAttribute string) FtCreateLanguageField {
	_ = "STUB: not implemented"
	return *new(FtCreateLanguageField)
}

func (c FtCreateOnHash) Score(defaultScore float64) FtCreateScore {
	_ = "STUB: not implemented"
	return *new(FtCreateScore)
}

func (c FtCreateOnHash) ScoreField(scoreAttribute string) FtCreateScoreField {
	_ = "STUB: not implemented"
	return *new(FtCreateScoreField)
}

func (c FtCreateOnHash) PayloadField(payloadAttribute string) FtCreatePayloadField {
	_ = "STUB: not implemented"
	return *new(FtCreatePayloadField)
}

func (c FtCreateOnHash) Maxtextfields() FtCreateMaxtextfields {
	_ = "STUB: not implemented"
	return *new(FtCreateMaxtextfields)
}

func (c FtCreateOnHash) Temporary(seconds float64) FtCreateTemporary {
	_ = "STUB: not implemented"
	return *new(FtCreateTemporary)
}

func (c FtCreateOnHash) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreateOnHash) Nohl() FtCreateNohl { _ = "STUB: not implemented"; return *new(FtCreateNohl) }

func (c FtCreateOnHash) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateOnHash) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateOnHash) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateOnHash) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateOnHash) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateOnJson Incomplete

func (c FtCreateOnJson) Prefix(count int64) FtCreatePrefixCount {
	_ = "STUB: not implemented"
	return *new(FtCreatePrefixCount)
}

func (c FtCreateOnJson) Filter(filter string) FtCreateFilter {
	_ = "STUB: not implemented"
	return *new(FtCreateFilter)
}

func (c FtCreateOnJson) Language(defaultLang string) FtCreateLanguage {
	_ = "STUB: not implemented"
	return *new(FtCreateLanguage)
}

func (c FtCreateOnJson) LanguageField(langAttribute string) FtCreateLanguageField {
	_ = "STUB: not implemented"
	return *new(FtCreateLanguageField)
}

func (c FtCreateOnJson) Score(defaultScore float64) FtCreateScore {
	_ = "STUB: not implemented"
	return *new(FtCreateScore)
}

func (c FtCreateOnJson) ScoreField(scoreAttribute string) FtCreateScoreField {
	_ = "STUB: not implemented"
	return *new(FtCreateScoreField)
}

func (c FtCreateOnJson) PayloadField(payloadAttribute string) FtCreatePayloadField {
	_ = "STUB: not implemented"
	return *new(FtCreatePayloadField)
}

func (c FtCreateOnJson) Maxtextfields() FtCreateMaxtextfields {
	_ = "STUB: not implemented"
	return *new(FtCreateMaxtextfields)
}

func (c FtCreateOnJson) Temporary(seconds float64) FtCreateTemporary {
	_ = "STUB: not implemented"
	return *new(FtCreateTemporary)
}

func (c FtCreateOnJson) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreateOnJson) Nohl() FtCreateNohl { _ = "STUB: not implemented"; return *new(FtCreateNohl) }

func (c FtCreateOnJson) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateOnJson) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateOnJson) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateOnJson) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateOnJson) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreatePayloadField Incomplete

func (c FtCreatePayloadField) Maxtextfields() FtCreateMaxtextfields {
	_ = "STUB: not implemented"
	return *new(FtCreateMaxtextfields)
}

func (c FtCreatePayloadField) Temporary(seconds float64) FtCreateTemporary {
	_ = "STUB: not implemented"
	return *new(FtCreateTemporary)
}

func (c FtCreatePayloadField) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreatePayloadField) Nohl() FtCreateNohl {
	_ = "STUB: not implemented"
	return *new(FtCreateNohl)
}

func (c FtCreatePayloadField) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreatePayloadField) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreatePayloadField) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreatePayloadField) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreatePayloadField) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreatePrefixCount Incomplete

func (c FtCreatePrefixCount) Prefix(prefix ...string) FtCreatePrefixPrefix {
	_ = "STUB: not implemented"
	return *new(FtCreatePrefixPrefix)
}

type FtCreatePrefixPrefix Incomplete

func (c FtCreatePrefixPrefix) Prefix(prefix ...string) FtCreatePrefixPrefix {
	_ = "STUB: not implemented"
	return *new(FtCreatePrefixPrefix)
}

func (c FtCreatePrefixPrefix) Filter(filter string) FtCreateFilter {
	_ = "STUB: not implemented"
	return *new(FtCreateFilter)
}

func (c FtCreatePrefixPrefix) Language(defaultLang string) FtCreateLanguage {
	_ = "STUB: not implemented"
	return *new(FtCreateLanguage)
}

func (c FtCreatePrefixPrefix) LanguageField(langAttribute string) FtCreateLanguageField {
	_ = "STUB: not implemented"
	return *new(FtCreateLanguageField)
}

func (c FtCreatePrefixPrefix) Score(defaultScore float64) FtCreateScore {
	_ = "STUB: not implemented"
	return *new(FtCreateScore)
}

func (c FtCreatePrefixPrefix) ScoreField(scoreAttribute string) FtCreateScoreField {
	_ = "STUB: not implemented"
	return *new(FtCreateScoreField)
}

func (c FtCreatePrefixPrefix) PayloadField(payloadAttribute string) FtCreatePayloadField {
	_ = "STUB: not implemented"
	return *new(FtCreatePayloadField)
}

func (c FtCreatePrefixPrefix) Maxtextfields() FtCreateMaxtextfields {
	_ = "STUB: not implemented"
	return *new(FtCreateMaxtextfields)
}

func (c FtCreatePrefixPrefix) Temporary(seconds float64) FtCreateTemporary {
	_ = "STUB: not implemented"
	return *new(FtCreateTemporary)
}

func (c FtCreatePrefixPrefix) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreatePrefixPrefix) Nohl() FtCreateNohl {
	_ = "STUB: not implemented"
	return *new(FtCreateNohl)
}

func (c FtCreatePrefixPrefix) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreatePrefixPrefix) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreatePrefixPrefix) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreatePrefixPrefix) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreatePrefixPrefix) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateSchema Incomplete

func (c FtCreateSchema) FieldName(fieldName string) FtCreateFieldFieldName {
	_ = "STUB: not implemented"
	return *new(FtCreateFieldFieldName)
}

type FtCreateScore Incomplete

func (c FtCreateScore) ScoreField(scoreAttribute string) FtCreateScoreField {
	_ = "STUB: not implemented"
	return *new(FtCreateScoreField)
}

func (c FtCreateScore) PayloadField(payloadAttribute string) FtCreatePayloadField {
	_ = "STUB: not implemented"
	return *new(FtCreatePayloadField)
}

func (c FtCreateScore) Maxtextfields() FtCreateMaxtextfields {
	_ = "STUB: not implemented"
	return *new(FtCreateMaxtextfields)
}

func (c FtCreateScore) Temporary(seconds float64) FtCreateTemporary {
	_ = "STUB: not implemented"
	return *new(FtCreateTemporary)
}

func (c FtCreateScore) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreateScore) Nohl() FtCreateNohl { _ = "STUB: not implemented"; return *new(FtCreateNohl) }

func (c FtCreateScore) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateScore) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateScore) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateScore) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateScore) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateScoreField Incomplete

func (c FtCreateScoreField) PayloadField(payloadAttribute string) FtCreatePayloadField {
	_ = "STUB: not implemented"
	return *new(FtCreatePayloadField)
}

func (c FtCreateScoreField) Maxtextfields() FtCreateMaxtextfields {
	_ = "STUB: not implemented"
	return *new(FtCreateMaxtextfields)
}

func (c FtCreateScoreField) Temporary(seconds float64) FtCreateTemporary {
	_ = "STUB: not implemented"
	return *new(FtCreateTemporary)
}

func (c FtCreateScoreField) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreateScoreField) Nohl() FtCreateNohl {
	_ = "STUB: not implemented"
	return *new(FtCreateNohl)
}

func (c FtCreateScoreField) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateScoreField) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateScoreField) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateScoreField) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateScoreField) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateSkipinitialscan Incomplete

func (c FtCreateSkipinitialscan) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateStopwordsStopword Incomplete

func (c FtCreateStopwordsStopword) Stopword(stopword ...string) FtCreateStopwordsStopword {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopword)
}

func (c FtCreateStopwordsStopword) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateStopwordsStopword) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateStopwordsStopwords Incomplete

func (c FtCreateStopwordsStopwords) Stopword(stopword ...string) FtCreateStopwordsStopword {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopword)
}

func (c FtCreateStopwordsStopwords) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateStopwordsStopwords) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCreateTemporary Incomplete

func (c FtCreateTemporary) Nooffsets() FtCreateNooffsets {
	_ = "STUB: not implemented"
	return *new(FtCreateNooffsets)
}

func (c FtCreateTemporary) Nohl() FtCreateNohl {
	_ = "STUB: not implemented"
	return *new(FtCreateNohl)
}

func (c FtCreateTemporary) Nofields() FtCreateNofields {
	_ = "STUB: not implemented"
	return *new(FtCreateNofields)
}

func (c FtCreateTemporary) Nofreqs() FtCreateNofreqs {
	_ = "STUB: not implemented"
	return *new(FtCreateNofreqs)
}

func (c FtCreateTemporary) Stopwords(count int64) FtCreateStopwordsStopwords {
	_ = "STUB: not implemented"
	return *new(FtCreateStopwordsStopwords)
}

func (c FtCreateTemporary) Skipinitialscan() FtCreateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtCreateSkipinitialscan)
}

func (c FtCreateTemporary) Schema() FtCreateSchema {
	_ = "STUB: not implemented"
	return *new(FtCreateSchema)
}

type FtCursorDel Incomplete

func (b Builder) FtCursorDel() (c FtCursorDel) { _ = "STUB: not implemented"; return *new(FtCursorDel) }

func (c FtCursorDel) Index(index string) FtCursorDelIndex {
	_ = "STUB: not implemented"
	return *new(FtCursorDelIndex)
}

type FtCursorDelCursorId Incomplete

func (c FtCursorDelCursorId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtCursorDelIndex Incomplete

func (c FtCursorDelIndex) CursorId(cursorId int64) FtCursorDelCursorId {
	_ = "STUB: not implemented"
	return *new(FtCursorDelCursorId)
}

type FtCursorRead Incomplete

func (b Builder) FtCursorRead() (c FtCursorRead) {
	_ = "STUB: not implemented"
	return *new(FtCursorRead)
}

func (c FtCursorRead) Index(index string) FtCursorReadIndex {
	_ = "STUB: not implemented"
	return *new(FtCursorReadIndex)
}

type FtCursorReadCount Incomplete

func (c FtCursorReadCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtCursorReadCursorId Incomplete

func (c FtCursorReadCursorId) Count(readSize int64) FtCursorReadCount {
	_ = "STUB: not implemented"
	return *new(FtCursorReadCount)
}

func (c FtCursorReadCursorId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtCursorReadIndex Incomplete

func (c FtCursorReadIndex) CursorId(cursorId int64) FtCursorReadCursorId {
	_ = "STUB: not implemented"
	return *new(FtCursorReadCursorId)
}

type FtDictadd Incomplete

func (b Builder) FtDictadd() (c FtDictadd) { _ = "STUB: not implemented"; return *new(FtDictadd) }

func (c FtDictadd) Dict(dict string) FtDictaddDict {
	_ = "STUB: not implemented"
	return *new(FtDictaddDict)
}

type FtDictaddDict Incomplete

func (c FtDictaddDict) Term(term ...string) FtDictaddTerm {
	_ = "STUB: not implemented"
	return *new(FtDictaddTerm)
}

type FtDictaddTerm Incomplete

func (c FtDictaddTerm) Term(term ...string) FtDictaddTerm {
	_ = "STUB: not implemented"
	return *new(FtDictaddTerm)
}

func (c FtDictaddTerm) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtDictdel Incomplete

func (b Builder) FtDictdel() (c FtDictdel) { _ = "STUB: not implemented"; return *new(FtDictdel) }

func (c FtDictdel) Dict(dict string) FtDictdelDict {
	_ = "STUB: not implemented"
	return *new(FtDictdelDict)
}

type FtDictdelDict Incomplete

func (c FtDictdelDict) Term(term ...string) FtDictdelTerm {
	_ = "STUB: not implemented"
	return *new(FtDictdelTerm)
}

type FtDictdelTerm Incomplete

func (c FtDictdelTerm) Term(term ...string) FtDictdelTerm {
	_ = "STUB: not implemented"
	return *new(FtDictdelTerm)
}

func (c FtDictdelTerm) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtDictdump Incomplete

func (b Builder) FtDictdump() (c FtDictdump) { _ = "STUB: not implemented"; return *new(FtDictdump) }

func (c FtDictdump) Dict(dict string) FtDictdumpDict {
	_ = "STUB: not implemented"
	return *new(FtDictdumpDict)
}

type FtDictdumpDict Incomplete

func (c FtDictdumpDict) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtDropindex Incomplete

func (b Builder) FtDropindex() (c FtDropindex) { _ = "STUB: not implemented"; return *new(FtDropindex) }

func (c FtDropindex) Index(index string) FtDropindexIndex {
	_ = "STUB: not implemented"
	return *new(FtDropindexIndex)
}

type FtDropindexDeleteDocsDd Incomplete

func (c FtDropindexDeleteDocsDd) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtDropindexIndex Incomplete

func (c FtDropindexIndex) Dd() FtDropindexDeleteDocsDd {
	_ = "STUB: not implemented"
	return *new(FtDropindexDeleteDocsDd)
}

func (c FtDropindexIndex) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtExplain Incomplete

func (b Builder) FtExplain() (c FtExplain) { _ = "STUB: not implemented"; return *new(FtExplain) }

func (c FtExplain) Index(index string) FtExplainIndex {
	_ = "STUB: not implemented"
	return *new(FtExplainIndex)
}

type FtExplainDialect Incomplete

func (c FtExplainDialect) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtExplainIndex Incomplete

func (c FtExplainIndex) Query(query string) FtExplainQuery {
	_ = "STUB: not implemented"
	return *new(FtExplainQuery)
}

type FtExplainQuery Incomplete

func (c FtExplainQuery) Dialect(dialect int64) FtExplainDialect {
	_ = "STUB: not implemented"
	return *new(FtExplainDialect)
}

func (c FtExplainQuery) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtExplaincli Incomplete

func (b Builder) FtExplaincli() (c FtExplaincli) {
	_ = "STUB: not implemented"
	return *new(FtExplaincli)
}

func (c FtExplaincli) Index(index string) FtExplaincliIndex {
	_ = "STUB: not implemented"
	return *new(FtExplaincliIndex)
}

type FtExplaincliDialect Incomplete

func (c FtExplaincliDialect) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtExplaincliIndex Incomplete

func (c FtExplaincliIndex) Query(query string) FtExplaincliQuery {
	_ = "STUB: not implemented"
	return *new(FtExplaincliQuery)
}

type FtExplaincliQuery Incomplete

func (c FtExplaincliQuery) Dialect(dialect int64) FtExplaincliDialect {
	_ = "STUB: not implemented"
	return *new(FtExplaincliDialect)
}

func (c FtExplaincliQuery) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtInfo Incomplete

func (b Builder) FtInfo() (c FtInfo) { _ = "STUB: not implemented"; return *new(FtInfo) }

func (c FtInfo) Index(index string) FtInfoIndex {
	_ = "STUB: not implemented"
	return *new(FtInfoIndex)
}

type FtInfoIndex Incomplete

func (c FtInfoIndex) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtList Incomplete

func (b Builder) FtList() (c FtList) { _ = "STUB: not implemented"; return *new(FtList) }

func (c FtList) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtProfile Incomplete

func (b Builder) FtProfile() (c FtProfile) { _ = "STUB: not implemented"; return *new(FtProfile) }

func (c FtProfile) Index(index string) FtProfileIndex {
	_ = "STUB: not implemented"
	return *new(FtProfileIndex)
}

type FtProfileIndex Incomplete

func (c FtProfileIndex) Search() FtProfileQuerytypeSearch {
	_ = "STUB: not implemented"
	return *new(FtProfileQuerytypeSearch)
}

func (c FtProfileIndex) Aggregate() FtProfileQuerytypeAggregate {
	_ = "STUB: not implemented"
	return *new(FtProfileQuerytypeAggregate)
}

type FtProfileLimited Incomplete

func (c FtProfileLimited) Query(query string) FtProfileQuery {
	_ = "STUB: not implemented"
	return *new(FtProfileQuery)
}

type FtProfileQuery Incomplete

func (c FtProfileQuery) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtProfileQuerytypeAggregate Incomplete

func (c FtProfileQuerytypeAggregate) Limited() FtProfileLimited {
	_ = "STUB: not implemented"
	return *new(FtProfileLimited)
}

func (c FtProfileQuerytypeAggregate) Query(query string) FtProfileQuery {
	_ = "STUB: not implemented"
	return *new(FtProfileQuery)
}

type FtProfileQuerytypeSearch Incomplete

func (c FtProfileQuerytypeSearch) Limited() FtProfileLimited {
	_ = "STUB: not implemented"
	return *new(FtProfileLimited)
}

func (c FtProfileQuerytypeSearch) Query(query string) FtProfileQuery {
	_ = "STUB: not implemented"
	return *new(FtProfileQuery)
}

type FtSearch Incomplete

func (b Builder) FtSearch() (c FtSearch) { _ = "STUB: not implemented"; return *new(FtSearch) }

func (c FtSearch) Index(index string) FtSearchIndex {
	_ = "STUB: not implemented"
	return *new(FtSearchIndex)
}

type FtSearchDialect Incomplete

func (c FtSearchDialect) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchExpander Incomplete

func (c FtSearchExpander) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchExpander) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchExpander) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchExpander) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchExpander) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchExpander) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchExpander) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchExpander) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchExplainscore Incomplete

func (c FtSearchExplainscore) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchExplainscore) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchExplainscore) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchExplainscore) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchExplainscore) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchExplainscore) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchFilterFilter Incomplete

func (c FtSearchFilterFilter) Min(min float64) FtSearchFilterMin {
	_ = "STUB: not implemented"
	return *new(FtSearchFilterMin)
}

type FtSearchFilterMax Incomplete

func (c FtSearchFilterMax) Filter(numericField string) FtSearchFilterFilter {
	_ = "STUB: not implemented"
	return *new(FtSearchFilterFilter)
}

func (c FtSearchFilterMax) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchFilterMax) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchFilterMax) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchFilterMax) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchFilterMax) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchFilterMax) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchFilterMax) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchFilterMax) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchFilterMax) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchFilterMax) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchFilterMax) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchFilterMax) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchFilterMax) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchFilterMax) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchFilterMax) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchFilterMax) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchFilterMax) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchFilterMax) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchFilterMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchFilterMin Incomplete

func (c FtSearchFilterMin) Max(max float64) FtSearchFilterMax {
	_ = "STUB: not implemented"
	return *new(FtSearchFilterMax)
}

type FtSearchGeoFilterGeofilter Incomplete

func (c FtSearchGeoFilterGeofilter) Lon(lon float64) FtSearchGeoFilterLon {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterLon)
}

type FtSearchGeoFilterLat Incomplete

func (c FtSearchGeoFilterLat) Radius(radius float64) FtSearchGeoFilterRadius {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterRadius)
}

type FtSearchGeoFilterLon Incomplete

func (c FtSearchGeoFilterLon) Lat(lat float64) FtSearchGeoFilterLat {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterLat)
}

type FtSearchGeoFilterRadius Incomplete

func (c FtSearchGeoFilterRadius) M() FtSearchGeoFilterRadiusTypeM {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterRadiusTypeM)
}

func (c FtSearchGeoFilterRadius) Km() FtSearchGeoFilterRadiusTypeKm {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterRadiusTypeKm)
}

func (c FtSearchGeoFilterRadius) Mi() FtSearchGeoFilterRadiusTypeMi {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterRadiusTypeMi)
}

func (c FtSearchGeoFilterRadius) Ft() FtSearchGeoFilterRadiusTypeFt {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterRadiusTypeFt)
}

type FtSearchGeoFilterRadiusTypeFt Incomplete

func (c FtSearchGeoFilterRadiusTypeFt) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchGeoFilterRadiusTypeFt) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchGeoFilterRadiusTypeFt) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchGeoFilterRadiusTypeFt) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchGeoFilterRadiusTypeFt) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchGeoFilterRadiusTypeFt) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchGeoFilterRadiusTypeFt) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchGeoFilterRadiusTypeFt) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchGeoFilterRadiusTypeFt) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchGeoFilterRadiusTypeFt) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchGeoFilterRadiusTypeFt) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchGeoFilterRadiusTypeFt) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchGeoFilterRadiusTypeFt) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchGeoFilterRadiusTypeFt) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchGeoFilterRadiusTypeFt) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchGeoFilterRadiusTypeFt) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchGeoFilterRadiusTypeFt) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchGeoFilterRadiusTypeFt) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchGeoFilterRadiusTypeFt) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchGeoFilterRadiusTypeKm Incomplete

func (c FtSearchGeoFilterRadiusTypeKm) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchGeoFilterRadiusTypeKm) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchGeoFilterRadiusTypeKm) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchGeoFilterRadiusTypeKm) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchGeoFilterRadiusTypeKm) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchGeoFilterRadiusTypeKm) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchGeoFilterRadiusTypeKm) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchGeoFilterRadiusTypeKm) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchGeoFilterRadiusTypeKm) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchGeoFilterRadiusTypeKm) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchGeoFilterRadiusTypeKm) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchGeoFilterRadiusTypeKm) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchGeoFilterRadiusTypeKm) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchGeoFilterRadiusTypeKm) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchGeoFilterRadiusTypeKm) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchGeoFilterRadiusTypeKm) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchGeoFilterRadiusTypeKm) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchGeoFilterRadiusTypeKm) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchGeoFilterRadiusTypeKm) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchGeoFilterRadiusTypeM Incomplete

func (c FtSearchGeoFilterRadiusTypeM) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchGeoFilterRadiusTypeM) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchGeoFilterRadiusTypeM) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchGeoFilterRadiusTypeM) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchGeoFilterRadiusTypeM) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchGeoFilterRadiusTypeM) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchGeoFilterRadiusTypeM) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchGeoFilterRadiusTypeM) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchGeoFilterRadiusTypeM) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchGeoFilterRadiusTypeM) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchGeoFilterRadiusTypeM) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchGeoFilterRadiusTypeM) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchGeoFilterRadiusTypeM) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchGeoFilterRadiusTypeM) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchGeoFilterRadiusTypeM) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchGeoFilterRadiusTypeM) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchGeoFilterRadiusTypeM) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchGeoFilterRadiusTypeM) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchGeoFilterRadiusTypeM) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchGeoFilterRadiusTypeMi Incomplete

func (c FtSearchGeoFilterRadiusTypeMi) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchGeoFilterRadiusTypeMi) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchGeoFilterRadiusTypeMi) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchGeoFilterRadiusTypeMi) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchGeoFilterRadiusTypeMi) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchGeoFilterRadiusTypeMi) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchGeoFilterRadiusTypeMi) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchGeoFilterRadiusTypeMi) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchGeoFilterRadiusTypeMi) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchGeoFilterRadiusTypeMi) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchGeoFilterRadiusTypeMi) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchGeoFilterRadiusTypeMi) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchGeoFilterRadiusTypeMi) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchGeoFilterRadiusTypeMi) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchGeoFilterRadiusTypeMi) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchGeoFilterRadiusTypeMi) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchGeoFilterRadiusTypeMi) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchGeoFilterRadiusTypeMi) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchGeoFilterRadiusTypeMi) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchHighlightFieldsField Incomplete

func (c FtSearchHighlightFieldsField) Field(field ...string) FtSearchHighlightFieldsField {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightFieldsField)
}

func (c FtSearchHighlightFieldsField) Tags() FtSearchHighlightTagsTags {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightTagsTags)
}

func (c FtSearchHighlightFieldsField) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchHighlightFieldsField) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchHighlightFieldsField) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchHighlightFieldsField) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchHighlightFieldsField) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchHighlightFieldsField) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchHighlightFieldsField) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchHighlightFieldsField) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchHighlightFieldsField) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchHighlightFieldsField) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchHighlightFieldsField) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchHighlightFieldsField) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchHighlightFieldsField) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchHighlightFieldsFields Incomplete

func (c FtSearchHighlightFieldsFields) Field(field ...string) FtSearchHighlightFieldsField {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightFieldsField)
}

type FtSearchHighlightHighlight Incomplete

func (c FtSearchHighlightHighlight) Fields(count string) FtSearchHighlightFieldsFields {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightFieldsFields)
}

func (c FtSearchHighlightHighlight) Tags() FtSearchHighlightTagsTags {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightTagsTags)
}

func (c FtSearchHighlightHighlight) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchHighlightHighlight) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchHighlightHighlight) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchHighlightHighlight) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchHighlightHighlight) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchHighlightHighlight) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchHighlightHighlight) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchHighlightHighlight) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchHighlightHighlight) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchHighlightHighlight) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchHighlightHighlight) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchHighlightHighlight) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchHighlightHighlight) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchHighlightTagsOpenClose Incomplete

func (c FtSearchHighlightTagsOpenClose) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchHighlightTagsOpenClose) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchHighlightTagsOpenClose) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchHighlightTagsOpenClose) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchHighlightTagsOpenClose) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchHighlightTagsOpenClose) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchHighlightTagsOpenClose) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchHighlightTagsOpenClose) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchHighlightTagsOpenClose) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchHighlightTagsOpenClose) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchHighlightTagsOpenClose) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchHighlightTagsOpenClose) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchHighlightTagsOpenClose) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchHighlightTagsTags Incomplete

func (c FtSearchHighlightTagsTags) OpenClose(open string, close string) FtSearchHighlightTagsOpenClose {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightTagsOpenClose)
}

type FtSearchInFieldsField Incomplete

func (c FtSearchInFieldsField) Field(field ...string) FtSearchInFieldsField {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsField)
}

func (c FtSearchInFieldsField) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchInFieldsField) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchInFieldsField) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchInFieldsField) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchInFieldsField) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchInFieldsField) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchInFieldsField) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchInFieldsField) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchInFieldsField) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchInFieldsField) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchInFieldsField) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchInFieldsField) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchInFieldsField) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchInFieldsField) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchInFieldsField) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchInFieldsField) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchInFieldsInfields Incomplete

func (c FtSearchInFieldsInfields) Field(field ...string) FtSearchInFieldsField {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsField)
}

type FtSearchInKeysInkeys Incomplete

func (c FtSearchInKeysInkeys) Key(key ...string) FtSearchInKeysKey {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysKey)
}

type FtSearchInKeysKey Incomplete

func (c FtSearchInKeysKey) Key(key ...string) FtSearchInKeysKey {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysKey)
}

func (c FtSearchInKeysKey) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchInKeysKey) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchInKeysKey) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchInKeysKey) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchInKeysKey) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchInKeysKey) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchInKeysKey) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchInKeysKey) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchInKeysKey) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchInKeysKey) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchInKeysKey) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchInKeysKey) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchInKeysKey) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchInKeysKey) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchInKeysKey) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchInKeysKey) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchInKeysKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchIndex Incomplete

func (c FtSearchIndex) Query(query string) FtSearchQuery {
	_ = "STUB: not implemented"
	return *new(FtSearchQuery)
}

type FtSearchLanguage Incomplete

func (c FtSearchLanguage) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchLanguage) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchLanguage) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchLanguage) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchLanguage) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchLanguage) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchLanguage) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchLanguage) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchLanguage) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchLimitLimit Incomplete

func (c FtSearchLimitLimit) OffsetNum(offset int64, num int64) FtSearchLimitOffsetNum {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitOffsetNum)
}

type FtSearchLimitOffsetNum Incomplete

func (c FtSearchLimitOffsetNum) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchLimitOffsetNum) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchLimitOffsetNum) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchNocontent Incomplete

func (c FtSearchNocontent) Verbatim() FtSearchVerbatim {
	_ = "STUB: not implemented"
	return *new(FtSearchVerbatim)
}

func (c FtSearchNocontent) Nostopwords() FtSearchNostopwords {
	_ = "STUB: not implemented"
	return *new(FtSearchNostopwords)
}

func (c FtSearchNocontent) Withscores() FtSearchWithscores {
	_ = "STUB: not implemented"
	return *new(FtSearchWithscores)
}

func (c FtSearchNocontent) Withpayloads() FtSearchWithpayloads {
	_ = "STUB: not implemented"
	return *new(FtSearchWithpayloads)
}

func (c FtSearchNocontent) Withsortkeys() FtSearchWithsortkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchWithsortkeys)
}

func (c FtSearchNocontent) Filter(numericField string) FtSearchFilterFilter {
	_ = "STUB: not implemented"
	return *new(FtSearchFilterFilter)
}

func (c FtSearchNocontent) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchNocontent) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchNocontent) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchNocontent) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchNocontent) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchNocontent) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchNocontent) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchNocontent) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchNocontent) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchNocontent) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchNocontent) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchNocontent) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchNocontent) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchNocontent) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchNocontent) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchNocontent) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchNocontent) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchNocontent) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchNocontent) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchNostopwords Incomplete

func (c FtSearchNostopwords) Withscores() FtSearchWithscores {
	_ = "STUB: not implemented"
	return *new(FtSearchWithscores)
}

func (c FtSearchNostopwords) Withpayloads() FtSearchWithpayloads {
	_ = "STUB: not implemented"
	return *new(FtSearchWithpayloads)
}

func (c FtSearchNostopwords) Withsortkeys() FtSearchWithsortkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchWithsortkeys)
}

func (c FtSearchNostopwords) Filter(numericField string) FtSearchFilterFilter {
	_ = "STUB: not implemented"
	return *new(FtSearchFilterFilter)
}

func (c FtSearchNostopwords) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchNostopwords) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchNostopwords) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchNostopwords) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchNostopwords) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchNostopwords) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchNostopwords) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchNostopwords) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchNostopwords) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchNostopwords) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchNostopwords) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchNostopwords) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchNostopwords) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchNostopwords) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchNostopwords) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchNostopwords) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchNostopwords) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchNostopwords) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchNostopwords) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchParamsNameValue Incomplete

func (c FtSearchParamsNameValue) NameValue(name string, value string) FtSearchParamsNameValue {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsNameValue)
}

func (c FtSearchParamsNameValue) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchParamsNameValue) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchParamsNargs Incomplete

func (c FtSearchParamsNargs) NameValue() FtSearchParamsNameValue {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsNameValue)
}

type FtSearchParamsParams Incomplete

func (c FtSearchParamsParams) Nargs(nargs int64) FtSearchParamsNargs {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsNargs)
}

type FtSearchPayload Incomplete

func (c FtSearchPayload) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchPayload) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchPayload) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchPayload) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchPayload) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchQuery Incomplete

func (c FtSearchQuery) Nocontent() FtSearchNocontent {
	_ = "STUB: not implemented"
	return *new(FtSearchNocontent)
}

func (c FtSearchQuery) Verbatim() FtSearchVerbatim {
	_ = "STUB: not implemented"
	return *new(FtSearchVerbatim)
}

func (c FtSearchQuery) Nostopwords() FtSearchNostopwords {
	_ = "STUB: not implemented"
	return *new(FtSearchNostopwords)
}

func (c FtSearchQuery) Withscores() FtSearchWithscores {
	_ = "STUB: not implemented"
	return *new(FtSearchWithscores)
}

func (c FtSearchQuery) Withpayloads() FtSearchWithpayloads {
	_ = "STUB: not implemented"
	return *new(FtSearchWithpayloads)
}

func (c FtSearchQuery) Withsortkeys() FtSearchWithsortkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchWithsortkeys)
}

func (c FtSearchQuery) Filter(numericField string) FtSearchFilterFilter {
	_ = "STUB: not implemented"
	return *new(FtSearchFilterFilter)
}

func (c FtSearchQuery) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchQuery) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchQuery) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchQuery) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchQuery) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchQuery) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchQuery) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchQuery) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchQuery) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchQuery) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchQuery) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchQuery) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchQuery) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchQuery) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchQuery) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchQuery) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchQuery) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchQuery) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchQuery) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchReturnIdentifiersAs Incomplete

func (c FtSearchReturnIdentifiersAs) Identifier(identifier string) FtSearchReturnIdentifiersIdentifier {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnIdentifiersIdentifier)
}

func (c FtSearchReturnIdentifiersAs) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchReturnIdentifiersAs) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchReturnIdentifiersAs) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchReturnIdentifiersAs) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchReturnIdentifiersAs) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchReturnIdentifiersAs) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchReturnIdentifiersAs) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchReturnIdentifiersAs) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchReturnIdentifiersAs) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchReturnIdentifiersAs) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchReturnIdentifiersAs) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchReturnIdentifiersAs) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchReturnIdentifiersAs) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchReturnIdentifiersAs) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchReturnIdentifiersAs) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchReturnIdentifiersIdentifier Incomplete

func (c FtSearchReturnIdentifiersIdentifier) As(property string) FtSearchReturnIdentifiersAs {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnIdentifiersAs)
}

func (c FtSearchReturnIdentifiersIdentifier) Identifier(identifier string) FtSearchReturnIdentifiersIdentifier {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnIdentifiersIdentifier)
}

func (c FtSearchReturnIdentifiersIdentifier) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchReturnIdentifiersIdentifier) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchReturnIdentifiersIdentifier) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchReturnIdentifiersIdentifier) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchReturnIdentifiersIdentifier) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchReturnIdentifiersIdentifier) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchReturnIdentifiersIdentifier) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchReturnIdentifiersIdentifier) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchReturnIdentifiersIdentifier) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchReturnIdentifiersIdentifier) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchReturnIdentifiersIdentifier) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchReturnIdentifiersIdentifier) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchReturnIdentifiersIdentifier) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchReturnIdentifiersIdentifier) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchReturnIdentifiersIdentifier) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchReturnReturn Incomplete

func (c FtSearchReturnReturn) Identifier(identifier string) FtSearchReturnIdentifiersIdentifier {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnIdentifiersIdentifier)
}

type FtSearchScorer Incomplete

func (c FtSearchScorer) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchScorer) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchScorer) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchScorer) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchScorer) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchScorer) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchScorer) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchSlop Incomplete

func (c FtSearchSlop) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchSlop) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchSlop) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchSlop) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchSlop) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchSlop) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchSlop) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchSlop) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchSlop) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchSlop) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchSlop) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchSlop) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchSortbyOrderAsc Incomplete

func (c FtSearchSortbyOrderAsc) Withcount() FtSearchSortbyWithcount {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbyWithcount)
}

func (c FtSearchSortbyOrderAsc) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchSortbyOrderAsc) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchSortbyOrderAsc) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchSortbyOrderAsc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchSortbyOrderDesc Incomplete

func (c FtSearchSortbyOrderDesc) Withcount() FtSearchSortbyWithcount {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbyWithcount)
}

func (c FtSearchSortbyOrderDesc) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchSortbyOrderDesc) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchSortbyOrderDesc) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchSortbyOrderDesc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchSortbySortby Incomplete

func (c FtSearchSortbySortby) Asc() FtSearchSortbyOrderAsc {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbyOrderAsc)
}

func (c FtSearchSortbySortby) Desc() FtSearchSortbyOrderDesc {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbyOrderDesc)
}

func (c FtSearchSortbySortby) Withcount() FtSearchSortbyWithcount {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbyWithcount)
}

func (c FtSearchSortbySortby) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchSortbySortby) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchSortbySortby) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchSortbySortby) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchSortbyWithcount Incomplete

func (c FtSearchSortbyWithcount) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchSortbyWithcount) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchSortbyWithcount) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchSortbyWithcount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchSummarizeFieldsField Incomplete

func (c FtSearchSummarizeFieldsField) Field(field ...string) FtSearchSummarizeFieldsField {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeFieldsField)
}

func (c FtSearchSummarizeFieldsField) Frags(num int64) FtSearchSummarizeFrags {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeFrags)
}

func (c FtSearchSummarizeFieldsField) Len(fragsize int64) FtSearchSummarizeLen {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeLen)
}

func (c FtSearchSummarizeFieldsField) Separator(separator string) FtSearchSummarizeSeparator {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSeparator)
}

func (c FtSearchSummarizeFieldsField) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchSummarizeFieldsField) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchSummarizeFieldsField) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchSummarizeFieldsField) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchSummarizeFieldsField) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchSummarizeFieldsField) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchSummarizeFieldsField) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchSummarizeFieldsField) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchSummarizeFieldsField) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchSummarizeFieldsField) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchSummarizeFieldsField) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchSummarizeFieldsField) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchSummarizeFieldsField) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchSummarizeFieldsField) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchSummarizeFieldsFields Incomplete

func (c FtSearchSummarizeFieldsFields) Field(field ...string) FtSearchSummarizeFieldsField {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeFieldsField)
}

type FtSearchSummarizeFrags Incomplete

func (c FtSearchSummarizeFrags) Len(fragsize int64) FtSearchSummarizeLen {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeLen)
}

func (c FtSearchSummarizeFrags) Separator(separator string) FtSearchSummarizeSeparator {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSeparator)
}

func (c FtSearchSummarizeFrags) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchSummarizeFrags) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchSummarizeFrags) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchSummarizeFrags) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchSummarizeFrags) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchSummarizeFrags) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchSummarizeFrags) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchSummarizeFrags) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchSummarizeFrags) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchSummarizeFrags) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchSummarizeFrags) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchSummarizeFrags) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchSummarizeFrags) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchSummarizeFrags) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchSummarizeLen Incomplete

func (c FtSearchSummarizeLen) Separator(separator string) FtSearchSummarizeSeparator {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSeparator)
}

func (c FtSearchSummarizeLen) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchSummarizeLen) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchSummarizeLen) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchSummarizeLen) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchSummarizeLen) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchSummarizeLen) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchSummarizeLen) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchSummarizeLen) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchSummarizeLen) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchSummarizeLen) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchSummarizeLen) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchSummarizeLen) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchSummarizeLen) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchSummarizeLen) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchSummarizeSeparator Incomplete

func (c FtSearchSummarizeSeparator) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchSummarizeSeparator) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchSummarizeSeparator) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchSummarizeSeparator) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchSummarizeSeparator) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchSummarizeSeparator) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchSummarizeSeparator) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchSummarizeSeparator) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchSummarizeSeparator) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchSummarizeSeparator) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchSummarizeSeparator) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchSummarizeSeparator) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchSummarizeSeparator) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchSummarizeSeparator) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchSummarizeSummarize Incomplete

func (c FtSearchSummarizeSummarize) Fields(count string) FtSearchSummarizeFieldsFields {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeFieldsFields)
}

func (c FtSearchSummarizeSummarize) Frags(num int64) FtSearchSummarizeFrags {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeFrags)
}

func (c FtSearchSummarizeSummarize) Len(fragsize int64) FtSearchSummarizeLen {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeLen)
}

func (c FtSearchSummarizeSummarize) Separator(separator string) FtSearchSummarizeSeparator {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSeparator)
}

func (c FtSearchSummarizeSummarize) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchSummarizeSummarize) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchSummarizeSummarize) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchSummarizeSummarize) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchSummarizeSummarize) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchSummarizeSummarize) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchSummarizeSummarize) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchSummarizeSummarize) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchSummarizeSummarize) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchSummarizeSummarize) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchSummarizeSummarize) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchSummarizeSummarize) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchSummarizeSummarize) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchSummarizeSummarize) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSearchTagsInorder Incomplete

func (c FtSearchTagsInorder) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchTagsInorder) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchTagsInorder) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchTagsInorder) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchTagsInorder) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchTagsInorder) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchTagsInorder) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchTagsInorder) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchTagsInorder) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchTagsInorder) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchTimeout Incomplete

func (c FtSearchTimeout) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchTimeout) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchTimeout) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchTimeout) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchTimeout) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchTimeout) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchTimeout) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchTimeout) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchTimeout) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchTimeout) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchVerbatim Incomplete

func (c FtSearchVerbatim) Nostopwords() FtSearchNostopwords {
	_ = "STUB: not implemented"
	return *new(FtSearchNostopwords)
}

func (c FtSearchVerbatim) Withscores() FtSearchWithscores {
	_ = "STUB: not implemented"
	return *new(FtSearchWithscores)
}

func (c FtSearchVerbatim) Withpayloads() FtSearchWithpayloads {
	_ = "STUB: not implemented"
	return *new(FtSearchWithpayloads)
}

func (c FtSearchVerbatim) Withsortkeys() FtSearchWithsortkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchWithsortkeys)
}

func (c FtSearchVerbatim) Filter(numericField string) FtSearchFilterFilter {
	_ = "STUB: not implemented"
	return *new(FtSearchFilterFilter)
}

func (c FtSearchVerbatim) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchVerbatim) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchVerbatim) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchVerbatim) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchVerbatim) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchVerbatim) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchVerbatim) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchVerbatim) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchVerbatim) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchVerbatim) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchVerbatim) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchVerbatim) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchVerbatim) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchVerbatim) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchVerbatim) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchVerbatim) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchVerbatim) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchVerbatim) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchVerbatim) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchWithpayloads Incomplete

func (c FtSearchWithpayloads) Withsortkeys() FtSearchWithsortkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchWithsortkeys)
}

func (c FtSearchWithpayloads) Filter(numericField string) FtSearchFilterFilter {
	_ = "STUB: not implemented"
	return *new(FtSearchFilterFilter)
}

func (c FtSearchWithpayloads) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchWithpayloads) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchWithpayloads) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchWithpayloads) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchWithpayloads) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchWithpayloads) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchWithpayloads) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchWithpayloads) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchWithpayloads) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchWithpayloads) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchWithpayloads) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchWithpayloads) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchWithpayloads) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchWithpayloads) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchWithpayloads) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchWithpayloads) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchWithpayloads) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchWithpayloads) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchWithpayloads) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchWithscores Incomplete

func (c FtSearchWithscores) Withpayloads() FtSearchWithpayloads {
	_ = "STUB: not implemented"
	return *new(FtSearchWithpayloads)
}

func (c FtSearchWithscores) Withsortkeys() FtSearchWithsortkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchWithsortkeys)
}

func (c FtSearchWithscores) Filter(numericField string) FtSearchFilterFilter {
	_ = "STUB: not implemented"
	return *new(FtSearchFilterFilter)
}

func (c FtSearchWithscores) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchWithscores) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchWithscores) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchWithscores) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchWithscores) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchWithscores) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchWithscores) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchWithscores) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchWithscores) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchWithscores) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchWithscores) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchWithscores) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchWithscores) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchWithscores) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchWithscores) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchWithscores) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchWithscores) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchWithscores) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchWithscores) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSearchWithsortkeys Incomplete

func (c FtSearchWithsortkeys) Filter(numericField string) FtSearchFilterFilter {
	_ = "STUB: not implemented"
	return *new(FtSearchFilterFilter)
}

func (c FtSearchWithsortkeys) Geofilter(geoField string) FtSearchGeoFilterGeofilter {
	_ = "STUB: not implemented"
	return *new(FtSearchGeoFilterGeofilter)
}

func (c FtSearchWithsortkeys) Inkeys(count string) FtSearchInKeysInkeys {
	_ = "STUB: not implemented"
	return *new(FtSearchInKeysInkeys)
}

func (c FtSearchWithsortkeys) Infields(count string) FtSearchInFieldsInfields {
	_ = "STUB: not implemented"
	return *new(FtSearchInFieldsInfields)
}

func (c FtSearchWithsortkeys) Return(count string) FtSearchReturnReturn {
	_ = "STUB: not implemented"
	return *new(FtSearchReturnReturn)
}

func (c FtSearchWithsortkeys) Summarize() FtSearchSummarizeSummarize {
	_ = "STUB: not implemented"
	return *new(FtSearchSummarizeSummarize)
}

func (c FtSearchWithsortkeys) Highlight() FtSearchHighlightHighlight {
	_ = "STUB: not implemented"
	return *new(FtSearchHighlightHighlight)
}

func (c FtSearchWithsortkeys) Slop(slop int64) FtSearchSlop {
	_ = "STUB: not implemented"
	return *new(FtSearchSlop)
}

func (c FtSearchWithsortkeys) Timeout(timeout int64) FtSearchTimeout {
	_ = "STUB: not implemented"
	return *new(FtSearchTimeout)
}

func (c FtSearchWithsortkeys) Inorder() FtSearchTagsInorder {
	_ = "STUB: not implemented"
	return *new(FtSearchTagsInorder)
}

func (c FtSearchWithsortkeys) Language(language string) FtSearchLanguage {
	_ = "STUB: not implemented"
	return *new(FtSearchLanguage)
}

func (c FtSearchWithsortkeys) Expander(expander string) FtSearchExpander {
	_ = "STUB: not implemented"
	return *new(FtSearchExpander)
}

func (c FtSearchWithsortkeys) Scorer(scorer string) FtSearchScorer {
	_ = "STUB: not implemented"
	return *new(FtSearchScorer)
}

func (c FtSearchWithsortkeys) Explainscore() FtSearchExplainscore {
	_ = "STUB: not implemented"
	return *new(FtSearchExplainscore)
}

func (c FtSearchWithsortkeys) Payload(payload string) FtSearchPayload {
	_ = "STUB: not implemented"
	return *new(FtSearchPayload)
}

func (c FtSearchWithsortkeys) Sortby(sortby string) FtSearchSortbySortby {
	_ = "STUB: not implemented"
	return *new(FtSearchSortbySortby)
}

func (c FtSearchWithsortkeys) Limit() FtSearchLimitLimit {
	_ = "STUB: not implemented"
	return *new(FtSearchLimitLimit)
}

func (c FtSearchWithsortkeys) Params() FtSearchParamsParams {
	_ = "STUB: not implemented"
	return *new(FtSearchParamsParams)
}

func (c FtSearchWithsortkeys) Dialect(dialect int64) FtSearchDialect {
	_ = "STUB: not implemented"
	return *new(FtSearchDialect)
}

func (c FtSearchWithsortkeys) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSpellcheck Incomplete

func (b Builder) FtSpellcheck() (c FtSpellcheck) {
	_ = "STUB: not implemented"
	return *new(FtSpellcheck)
}

func (c FtSpellcheck) Index(index string) FtSpellcheckIndex {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckIndex)
}

type FtSpellcheckDialect Incomplete

func (c FtSpellcheckDialect) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSpellcheckDistance Incomplete

func (c FtSpellcheckDistance) TermsInclude() FtSpellcheckTermsTermsInclude {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckTermsTermsInclude)
}

func (c FtSpellcheckDistance) TermsExclude() FtSpellcheckTermsTermsExclude {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckTermsTermsExclude)
}

func (c FtSpellcheckDistance) Dialect(dialect int64) FtSpellcheckDialect {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckDialect)
}

func (c FtSpellcheckDistance) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSpellcheckIndex Incomplete

func (c FtSpellcheckIndex) Query(query string) FtSpellcheckQuery {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckQuery)
}

type FtSpellcheckQuery Incomplete

func (c FtSpellcheckQuery) Distance(distance int64) FtSpellcheckDistance {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckDistance)
}

func (c FtSpellcheckQuery) TermsInclude() FtSpellcheckTermsTermsInclude {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckTermsTermsInclude)
}

func (c FtSpellcheckQuery) TermsExclude() FtSpellcheckTermsTermsExclude {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckTermsTermsExclude)
}

func (c FtSpellcheckQuery) Dialect(dialect int64) FtSpellcheckDialect {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckDialect)
}

func (c FtSpellcheckQuery) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSpellcheckTermsDictionary Incomplete

func (c FtSpellcheckTermsDictionary) Terms(terms ...string) FtSpellcheckTermsTerms {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckTermsTerms)
}

func (c FtSpellcheckTermsDictionary) Dialect(dialect int64) FtSpellcheckDialect {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckDialect)
}

func (c FtSpellcheckTermsDictionary) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSpellcheckTermsTerms Incomplete

func (c FtSpellcheckTermsTerms) Terms(terms ...string) FtSpellcheckTermsTerms {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckTermsTerms)
}

func (c FtSpellcheckTermsTerms) Dialect(dialect int64) FtSpellcheckDialect {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckDialect)
}

func (c FtSpellcheckTermsTerms) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FtSpellcheckTermsTermsExclude Incomplete

func (c FtSpellcheckTermsTermsExclude) Dictionary(dictionary string) FtSpellcheckTermsDictionary {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckTermsDictionary)
}

type FtSpellcheckTermsTermsInclude Incomplete

func (c FtSpellcheckTermsTermsInclude) Dictionary(dictionary string) FtSpellcheckTermsDictionary {
	_ = "STUB: not implemented"
	return *new(FtSpellcheckTermsDictionary)
}

type FtSyndump Incomplete

func (b Builder) FtSyndump() (c FtSyndump) { _ = "STUB: not implemented"; return *new(FtSyndump) }

func (c FtSyndump) Index(index string) FtSyndumpIndex {
	_ = "STUB: not implemented"
	return *new(FtSyndumpIndex)
}

type FtSyndumpIndex Incomplete

func (c FtSyndumpIndex) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtSynupdate Incomplete

func (b Builder) FtSynupdate() (c FtSynupdate) { _ = "STUB: not implemented"; return *new(FtSynupdate) }

func (c FtSynupdate) Index(index string) FtSynupdateIndex {
	_ = "STUB: not implemented"
	return *new(FtSynupdateIndex)
}

type FtSynupdateIndex Incomplete

func (c FtSynupdateIndex) SynonymGroupId(synonymGroupId string) FtSynupdateSynonymGroupId {
	_ = "STUB: not implemented"
	return *new(FtSynupdateSynonymGroupId)
}

type FtSynupdateSkipinitialscan Incomplete

func (c FtSynupdateSkipinitialscan) Term(term ...string) FtSynupdateTerm {
	_ = "STUB: not implemented"
	return *new(FtSynupdateTerm)
}

type FtSynupdateSynonymGroupId Incomplete

func (c FtSynupdateSynonymGroupId) Skipinitialscan() FtSynupdateSkipinitialscan {
	_ = "STUB: not implemented"
	return *new(FtSynupdateSkipinitialscan)
}

func (c FtSynupdateSynonymGroupId) Term(term ...string) FtSynupdateTerm {
	_ = "STUB: not implemented"
	return *new(FtSynupdateTerm)
}

type FtSynupdateTerm Incomplete

func (c FtSynupdateTerm) Term(term ...string) FtSynupdateTerm {
	_ = "STUB: not implemented"
	return *new(FtSynupdateTerm)
}

func (c FtSynupdateTerm) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtTagvals Incomplete

func (b Builder) FtTagvals() (c FtTagvals) { _ = "STUB: not implemented"; return *new(FtTagvals) }

func (c FtTagvals) Index(index string) FtTagvalsIndex {
	_ = "STUB: not implemented"
	return *new(FtTagvalsIndex)
}

type FtTagvalsFieldName Incomplete

func (c FtTagvalsFieldName) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FtTagvalsIndex Incomplete

func (c FtTagvalsIndex) FieldName(fieldName string) FtTagvalsFieldName {
	_ = "STUB: not implemented"
	return *new(FtTagvalsFieldName)
}
