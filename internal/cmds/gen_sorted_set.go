// Code generated DO NOT EDIT

package cmds

type Bzmpop Incomplete

func (b Builder) Bzmpop() (c Bzmpop) { _ = "STUB: not implemented"; return *new(Bzmpop) }

func (c Bzmpop) Timeout(timeout float64) BzmpopTimeout {
	_ = "STUB: not implemented"
	return *new(BzmpopTimeout)
}

type BzmpopCount Incomplete

func (c BzmpopCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BzmpopKey Incomplete

func (c BzmpopKey) Key(key ...string) BzmpopKey { _ = "STUB: not implemented"; return *new(BzmpopKey) }

func (c BzmpopKey) Min() BzmpopWhereMin { _ = "STUB: not implemented"; return *new(BzmpopWhereMin) }

func (c BzmpopKey) Max() BzmpopWhereMax { _ = "STUB: not implemented"; return *new(BzmpopWhereMax) }

type BzmpopNumkeys Incomplete

func (c BzmpopNumkeys) Key(key ...string) BzmpopKey {
	_ = "STUB: not implemented"
	return *new(BzmpopKey)
}

type BzmpopTimeout Incomplete

func (c BzmpopTimeout) Numkeys(numkeys int64) BzmpopNumkeys {
	_ = "STUB: not implemented"
	return *new(BzmpopNumkeys)
}

type BzmpopWhereMax Incomplete

func (c BzmpopWhereMax) Count(count int64) BzmpopCount {
	_ = "STUB: not implemented"
	return *new(BzmpopCount)
}

func (c BzmpopWhereMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BzmpopWhereMin Incomplete

func (c BzmpopWhereMin) Count(count int64) BzmpopCount {
	_ = "STUB: not implemented"
	return *new(BzmpopCount)
}

func (c BzmpopWhereMin) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Bzpopmax Incomplete

func (b Builder) Bzpopmax() (c Bzpopmax) { _ = "STUB: not implemented"; return *new(Bzpopmax) }

func (c Bzpopmax) Key(key ...string) BzpopmaxKey {
	_ = "STUB: not implemented"
	return *new(BzpopmaxKey)
}

type BzpopmaxKey Incomplete

func (c BzpopmaxKey) Key(key ...string) BzpopmaxKey {
	_ = "STUB: not implemented"
	return *new(BzpopmaxKey)
}

func (c BzpopmaxKey) Timeout(timeout float64) BzpopmaxTimeout {
	_ = "STUB: not implemented"
	return *new(BzpopmaxTimeout)
}

type BzpopmaxTimeout Incomplete

func (c BzpopmaxTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Bzpopmin Incomplete

func (b Builder) Bzpopmin() (c Bzpopmin) { _ = "STUB: not implemented"; return *new(Bzpopmin) }

func (c Bzpopmin) Key(key ...string) BzpopminKey {
	_ = "STUB: not implemented"
	return *new(BzpopminKey)
}

type BzpopminKey Incomplete

func (c BzpopminKey) Key(key ...string) BzpopminKey {
	_ = "STUB: not implemented"
	return *new(BzpopminKey)
}

func (c BzpopminKey) Timeout(timeout float64) BzpopminTimeout {
	_ = "STUB: not implemented"
	return *new(BzpopminTimeout)
}

type BzpopminTimeout Incomplete

func (c BzpopminTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zadd Incomplete

func (b Builder) Zadd() (c Zadd) { _ = "STUB: not implemented"; return *new(Zadd) }

func (c Zadd) Key(key string) ZaddKey { _ = "STUB: not implemented"; return *new(ZaddKey) }

type ZaddChangeCh Incomplete

func (c ZaddChangeCh) Incr() ZaddIncrementIncr {
	_ = "STUB: not implemented"
	return *new(ZaddIncrementIncr)
}

func (c ZaddChangeCh) ScoreMember() ZaddScoreMember {
	_ = "STUB: not implemented"
	return *new(ZaddScoreMember)
}

type ZaddComparisonGt Incomplete

func (c ZaddComparisonGt) Ch() ZaddChangeCh { _ = "STUB: not implemented"; return *new(ZaddChangeCh) }

func (c ZaddComparisonGt) Incr() ZaddIncrementIncr {
	_ = "STUB: not implemented"
	return *new(ZaddIncrementIncr)
}

func (c ZaddComparisonGt) ScoreMember() ZaddScoreMember {
	_ = "STUB: not implemented"
	return *new(ZaddScoreMember)
}

type ZaddComparisonLt Incomplete

func (c ZaddComparisonLt) Ch() ZaddChangeCh { _ = "STUB: not implemented"; return *new(ZaddChangeCh) }

func (c ZaddComparisonLt) Incr() ZaddIncrementIncr {
	_ = "STUB: not implemented"
	return *new(ZaddIncrementIncr)
}

func (c ZaddComparisonLt) ScoreMember() ZaddScoreMember {
	_ = "STUB: not implemented"
	return *new(ZaddScoreMember)
}

type ZaddConditionNx Incomplete

func (c ZaddConditionNx) Gt() ZaddComparisonGt {
	_ = "STUB: not implemented"
	return *new(ZaddComparisonGt)
}

func (c ZaddConditionNx) Lt() ZaddComparisonLt {
	_ = "STUB: not implemented"
	return *new(ZaddComparisonLt)
}

func (c ZaddConditionNx) Ch() ZaddChangeCh { _ = "STUB: not implemented"; return *new(ZaddChangeCh) }

func (c ZaddConditionNx) Incr() ZaddIncrementIncr {
	_ = "STUB: not implemented"
	return *new(ZaddIncrementIncr)
}

func (c ZaddConditionNx) ScoreMember() ZaddScoreMember {
	_ = "STUB: not implemented"
	return *new(ZaddScoreMember)
}

type ZaddConditionXx Incomplete

func (c ZaddConditionXx) Gt() ZaddComparisonGt {
	_ = "STUB: not implemented"
	return *new(ZaddComparisonGt)
}

func (c ZaddConditionXx) Lt() ZaddComparisonLt {
	_ = "STUB: not implemented"
	return *new(ZaddComparisonLt)
}

func (c ZaddConditionXx) Ch() ZaddChangeCh { _ = "STUB: not implemented"; return *new(ZaddChangeCh) }

func (c ZaddConditionXx) Incr() ZaddIncrementIncr {
	_ = "STUB: not implemented"
	return *new(ZaddIncrementIncr)
}

func (c ZaddConditionXx) ScoreMember() ZaddScoreMember {
	_ = "STUB: not implemented"
	return *new(ZaddScoreMember)
}

type ZaddIncrementIncr Incomplete

func (c ZaddIncrementIncr) ScoreMember() ZaddScoreMember {
	_ = "STUB: not implemented"
	return *new(ZaddScoreMember)
}

type ZaddKey Incomplete

func (c ZaddKey) Nx() ZaddConditionNx { _ = "STUB: not implemented"; return *new(ZaddConditionNx) }

func (c ZaddKey) Xx() ZaddConditionXx { _ = "STUB: not implemented"; return *new(ZaddConditionXx) }

func (c ZaddKey) Gt() ZaddComparisonGt { _ = "STUB: not implemented"; return *new(ZaddComparisonGt) }

func (c ZaddKey) Lt() ZaddComparisonLt { _ = "STUB: not implemented"; return *new(ZaddComparisonLt) }

func (c ZaddKey) Ch() ZaddChangeCh { _ = "STUB: not implemented"; return *new(ZaddChangeCh) }

func (c ZaddKey) Incr() ZaddIncrementIncr {
	_ = "STUB: not implemented"
	return *new(ZaddIncrementIncr)
}

func (c ZaddKey) ScoreMember() ZaddScoreMember {
	_ = "STUB: not implemented"
	return *new(ZaddScoreMember)
}

type ZaddScoreMember Incomplete

func (c ZaddScoreMember) ScoreMember(score float64, member string) ZaddScoreMember {
	_ = "STUB: not implemented"
	return *new(ZaddScoreMember)
}

func (c ZaddScoreMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zcard Incomplete

func (b Builder) Zcard() (c Zcard) { _ = "STUB: not implemented"; return *new(Zcard) }

func (c Zcard) Key(key string) ZcardKey { _ = "STUB: not implemented"; return *new(ZcardKey) }

type ZcardKey Incomplete

func (c ZcardKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZcardKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Zcount Incomplete

func (b Builder) Zcount() (c Zcount) { _ = "STUB: not implemented"; return *new(Zcount) }

func (c Zcount) Key(key string) ZcountKey { _ = "STUB: not implemented"; return *new(ZcountKey) }

type ZcountKey Incomplete

func (c ZcountKey) Min(min string) ZcountMin { _ = "STUB: not implemented"; return *new(ZcountMin) }

type ZcountMax Incomplete

func (c ZcountMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZcountMax) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZcountMin Incomplete

func (c ZcountMin) Max(max string) ZcountMax { _ = "STUB: not implemented"; return *new(ZcountMax) }

type Zdiff Incomplete

func (b Builder) Zdiff() (c Zdiff) { _ = "STUB: not implemented"; return *new(Zdiff) }

func (c Zdiff) Numkeys(numkeys int64) ZdiffNumkeys {
	_ = "STUB: not implemented"
	return *new(ZdiffNumkeys)
}

type ZdiffKey Incomplete

func (c ZdiffKey) Key(key ...string) ZdiffKey { _ = "STUB: not implemented"; return *new(ZdiffKey) }

func (c ZdiffKey) Withscores() ZdiffWithscores {
	_ = "STUB: not implemented"
	return *new(ZdiffWithscores)
}

func (c ZdiffKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZdiffNumkeys Incomplete

func (c ZdiffNumkeys) Key(key ...string) ZdiffKey { _ = "STUB: not implemented"; return *new(ZdiffKey) }

type ZdiffWithscores Incomplete

func (c ZdiffWithscores) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zdiffstore Incomplete

func (b Builder) Zdiffstore() (c Zdiffstore) { _ = "STUB: not implemented"; return *new(Zdiffstore) }

func (c Zdiffstore) Destination(destination string) ZdiffstoreDestination {
	_ = "STUB: not implemented"
	return *new(ZdiffstoreDestination)
}

type ZdiffstoreDestination Incomplete

func (c ZdiffstoreDestination) Numkeys(numkeys int64) ZdiffstoreNumkeys {
	_ = "STUB: not implemented"
	return *new(ZdiffstoreNumkeys)
}

type ZdiffstoreKey Incomplete

func (c ZdiffstoreKey) Key(key ...string) ZdiffstoreKey {
	_ = "STUB: not implemented"
	return *new(ZdiffstoreKey)
}

func (c ZdiffstoreKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZdiffstoreNumkeys Incomplete

func (c ZdiffstoreNumkeys) Key(key ...string) ZdiffstoreKey {
	_ = "STUB: not implemented"
	return *new(ZdiffstoreKey)
}

type Zincrby Incomplete

func (b Builder) Zincrby() (c Zincrby) { _ = "STUB: not implemented"; return *new(Zincrby) }

func (c Zincrby) Key(key string) ZincrbyKey { _ = "STUB: not implemented"; return *new(ZincrbyKey) }

type ZincrbyIncrement Incomplete

func (c ZincrbyIncrement) Member(member string) ZincrbyMember {
	_ = "STUB: not implemented"
	return *new(ZincrbyMember)
}

type ZincrbyKey Incomplete

func (c ZincrbyKey) Increment(increment float64) ZincrbyIncrement {
	_ = "STUB: not implemented"
	return *new(ZincrbyIncrement)
}

type ZincrbyMember Incomplete

func (c ZincrbyMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zinter Incomplete

func (b Builder) Zinter() (c Zinter) { _ = "STUB: not implemented"; return *new(Zinter) }

func (c Zinter) Numkeys(numkeys int64) ZinterNumkeys {
	_ = "STUB: not implemented"
	return *new(ZinterNumkeys)
}

type ZinterAggregateCount Incomplete

func (c ZinterAggregateCount) Withscores() ZinterWithscores {
	_ = "STUB: not implemented"
	return *new(ZinterWithscores)
}

func (c ZinterAggregateCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZinterAggregateMax Incomplete

func (c ZinterAggregateMax) Withscores() ZinterWithscores {
	_ = "STUB: not implemented"
	return *new(ZinterWithscores)
}

func (c ZinterAggregateMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZinterAggregateMin Incomplete

func (c ZinterAggregateMin) Withscores() ZinterWithscores {
	_ = "STUB: not implemented"
	return *new(ZinterWithscores)
}

func (c ZinterAggregateMin) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZinterAggregateSum Incomplete

func (c ZinterAggregateSum) Withscores() ZinterWithscores {
	_ = "STUB: not implemented"
	return *new(ZinterWithscores)
}

func (c ZinterAggregateSum) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZinterKey Incomplete

func (c ZinterKey) Key(key ...string) ZinterKey { _ = "STUB: not implemented"; return *new(ZinterKey) }

func (c ZinterKey) Weights(weight ...int64) ZinterWeights {
	_ = "STUB: not implemented"
	return *new(ZinterWeights)
}

func (c ZinterKey) AggregateSum() ZinterAggregateSum {
	_ = "STUB: not implemented"
	return *new(ZinterAggregateSum)
}

func (c ZinterKey) AggregateMin() ZinterAggregateMin {
	_ = "STUB: not implemented"
	return *new(ZinterAggregateMin)
}

func (c ZinterKey) AggregateMax() ZinterAggregateMax {
	_ = "STUB: not implemented"
	return *new(ZinterAggregateMax)
}

func (c ZinterKey) AggregateCount() ZinterAggregateCount {
	_ = "STUB: not implemented"
	return *new(ZinterAggregateCount)
}

func (c ZinterKey) Withscores() ZinterWithscores {
	_ = "STUB: not implemented"
	return *new(ZinterWithscores)
}

func (c ZinterKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZinterNumkeys Incomplete

func (c ZinterNumkeys) Key(key ...string) ZinterKey {
	_ = "STUB: not implemented"
	return *new(ZinterKey)
}

type ZinterWeights Incomplete

func (c ZinterWeights) Weights(weight ...int64) ZinterWeights {
	_ = "STUB: not implemented"
	return *new(ZinterWeights)
}

func (c ZinterWeights) AggregateSum() ZinterAggregateSum {
	_ = "STUB: not implemented"
	return *new(ZinterAggregateSum)
}

func (c ZinterWeights) AggregateMin() ZinterAggregateMin {
	_ = "STUB: not implemented"
	return *new(ZinterAggregateMin)
}

func (c ZinterWeights) AggregateMax() ZinterAggregateMax {
	_ = "STUB: not implemented"
	return *new(ZinterAggregateMax)
}

func (c ZinterWeights) AggregateCount() ZinterAggregateCount {
	_ = "STUB: not implemented"
	return *new(ZinterAggregateCount)
}

func (c ZinterWeights) Withscores() ZinterWithscores {
	_ = "STUB: not implemented"
	return *new(ZinterWithscores)
}

func (c ZinterWeights) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZinterWithscores Incomplete

func (c ZinterWithscores) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zintercard Incomplete

func (b Builder) Zintercard() (c Zintercard) { _ = "STUB: not implemented"; return *new(Zintercard) }

func (c Zintercard) Numkeys(numkeys int64) ZintercardNumkeys {
	_ = "STUB: not implemented"
	return *new(ZintercardNumkeys)
}

type ZintercardKey Incomplete

func (c ZintercardKey) Key(key ...string) ZintercardKey {
	_ = "STUB: not implemented"
	return *new(ZintercardKey)
}

func (c ZintercardKey) Limit(limit int64) ZintercardLimit {
	_ = "STUB: not implemented"
	return *new(ZintercardLimit)
}

func (c ZintercardKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZintercardLimit Incomplete

func (c ZintercardLimit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZintercardNumkeys Incomplete

func (c ZintercardNumkeys) Key(key ...string) ZintercardKey {
	_ = "STUB: not implemented"
	return *new(ZintercardKey)
}

type Zinterstore Incomplete

func (b Builder) Zinterstore() (c Zinterstore) { _ = "STUB: not implemented"; return *new(Zinterstore) }

func (c Zinterstore) Destination(destination string) ZinterstoreDestination {
	_ = "STUB: not implemented"
	return *new(ZinterstoreDestination)
}

type ZinterstoreAggregateCount Incomplete

func (c ZinterstoreAggregateCount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ZinterstoreAggregateMax Incomplete

func (c ZinterstoreAggregateMax) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ZinterstoreAggregateMin Incomplete

func (c ZinterstoreAggregateMin) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ZinterstoreAggregateSum Incomplete

func (c ZinterstoreAggregateSum) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ZinterstoreDestination Incomplete

func (c ZinterstoreDestination) Numkeys(numkeys int64) ZinterstoreNumkeys {
	_ = "STUB: not implemented"
	return *new(ZinterstoreNumkeys)
}

type ZinterstoreKey Incomplete

func (c ZinterstoreKey) Key(key ...string) ZinterstoreKey {
	_ = "STUB: not implemented"
	return *new(ZinterstoreKey)
}

func (c ZinterstoreKey) Weights(weight ...int64) ZinterstoreWeights {
	_ = "STUB: not implemented"
	return *new(ZinterstoreWeights)
}

func (c ZinterstoreKey) AggregateSum() ZinterstoreAggregateSum {
	_ = "STUB: not implemented"
	return *new(ZinterstoreAggregateSum)
}

func (c ZinterstoreKey) AggregateMin() ZinterstoreAggregateMin {
	_ = "STUB: not implemented"
	return *new(ZinterstoreAggregateMin)
}

func (c ZinterstoreKey) AggregateMax() ZinterstoreAggregateMax {
	_ = "STUB: not implemented"
	return *new(ZinterstoreAggregateMax)
}

func (c ZinterstoreKey) AggregateCount() ZinterstoreAggregateCount {
	_ = "STUB: not implemented"
	return *new(ZinterstoreAggregateCount)
}

func (c ZinterstoreKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZinterstoreNumkeys Incomplete

func (c ZinterstoreNumkeys) Key(key ...string) ZinterstoreKey {
	_ = "STUB: not implemented"
	return *new(ZinterstoreKey)
}

type ZinterstoreWeights Incomplete

func (c ZinterstoreWeights) Weights(weight ...int64) ZinterstoreWeights {
	_ = "STUB: not implemented"
	return *new(ZinterstoreWeights)
}

func (c ZinterstoreWeights) AggregateSum() ZinterstoreAggregateSum {
	_ = "STUB: not implemented"
	return *new(ZinterstoreAggregateSum)
}

func (c ZinterstoreWeights) AggregateMin() ZinterstoreAggregateMin {
	_ = "STUB: not implemented"
	return *new(ZinterstoreAggregateMin)
}

func (c ZinterstoreWeights) AggregateMax() ZinterstoreAggregateMax {
	_ = "STUB: not implemented"
	return *new(ZinterstoreAggregateMax)
}

func (c ZinterstoreWeights) AggregateCount() ZinterstoreAggregateCount {
	_ = "STUB: not implemented"
	return *new(ZinterstoreAggregateCount)
}

func (c ZinterstoreWeights) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zlexcount Incomplete

func (b Builder) Zlexcount() (c Zlexcount) { _ = "STUB: not implemented"; return *new(Zlexcount) }

func (c Zlexcount) Key(key string) ZlexcountKey {
	_ = "STUB: not implemented"
	return *new(ZlexcountKey)
}

type ZlexcountKey Incomplete

func (c ZlexcountKey) Min(min string) ZlexcountMin {
	_ = "STUB: not implemented"
	return *new(ZlexcountMin)
}

type ZlexcountMax Incomplete

func (c ZlexcountMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZlexcountMax) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZlexcountMin Incomplete

func (c ZlexcountMin) Max(max string) ZlexcountMax {
	_ = "STUB: not implemented"
	return *new(ZlexcountMax)
}

type Zmpop Incomplete

func (b Builder) Zmpop() (c Zmpop) { _ = "STUB: not implemented"; return *new(Zmpop) }

func (c Zmpop) Numkeys(numkeys int64) ZmpopNumkeys {
	_ = "STUB: not implemented"
	return *new(ZmpopNumkeys)
}

type ZmpopCount Incomplete

func (c ZmpopCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZmpopKey Incomplete

func (c ZmpopKey) Key(key ...string) ZmpopKey { _ = "STUB: not implemented"; return *new(ZmpopKey) }

func (c ZmpopKey) Min() ZmpopWhereMin { _ = "STUB: not implemented"; return *new(ZmpopWhereMin) }

func (c ZmpopKey) Max() ZmpopWhereMax { _ = "STUB: not implemented"; return *new(ZmpopWhereMax) }

type ZmpopNumkeys Incomplete

func (c ZmpopNumkeys) Key(key ...string) ZmpopKey { _ = "STUB: not implemented"; return *new(ZmpopKey) }

type ZmpopWhereMax Incomplete

func (c ZmpopWhereMax) Count(count int64) ZmpopCount {
	_ = "STUB: not implemented"
	return *new(ZmpopCount)
}

func (c ZmpopWhereMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZmpopWhereMin Incomplete

func (c ZmpopWhereMin) Count(count int64) ZmpopCount {
	_ = "STUB: not implemented"
	return *new(ZmpopCount)
}

func (c ZmpopWhereMin) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zmscore Incomplete

func (b Builder) Zmscore() (c Zmscore) { _ = "STUB: not implemented"; return *new(Zmscore) }

func (c Zmscore) Key(key string) ZmscoreKey { _ = "STUB: not implemented"; return *new(ZmscoreKey) }

type ZmscoreKey Incomplete

func (c ZmscoreKey) Member(member ...string) ZmscoreMember {
	_ = "STUB: not implemented"
	return *new(ZmscoreMember)
}

type ZmscoreMember Incomplete

func (c ZmscoreMember) Member(member ...string) ZmscoreMember {
	_ = "STUB: not implemented"
	return *new(ZmscoreMember)
}

func (c ZmscoreMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZmscoreMember) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Zpopmax Incomplete

func (b Builder) Zpopmax() (c Zpopmax) { _ = "STUB: not implemented"; return *new(Zpopmax) }

func (c Zpopmax) Key(key string) ZpopmaxKey { _ = "STUB: not implemented"; return *new(ZpopmaxKey) }

type ZpopmaxCount Incomplete

func (c ZpopmaxCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZpopmaxKey Incomplete

func (c ZpopmaxKey) Count(count int64) ZpopmaxCount {
	_ = "STUB: not implemented"
	return *new(ZpopmaxCount)
}

func (c ZpopmaxKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zpopmin Incomplete

func (b Builder) Zpopmin() (c Zpopmin) { _ = "STUB: not implemented"; return *new(Zpopmin) }

func (c Zpopmin) Key(key string) ZpopminKey { _ = "STUB: not implemented"; return *new(ZpopminKey) }

type ZpopminCount Incomplete

func (c ZpopminCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZpopminKey Incomplete

func (c ZpopminKey) Count(count int64) ZpopminCount {
	_ = "STUB: not implemented"
	return *new(ZpopminCount)
}

func (c ZpopminKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zrandmember Incomplete

func (b Builder) Zrandmember() (c Zrandmember) { _ = "STUB: not implemented"; return *new(Zrandmember) }

func (c Zrandmember) Key(key string) ZrandmemberKey {
	_ = "STUB: not implemented"
	return *new(ZrandmemberKey)
}

type ZrandmemberKey Incomplete

func (c ZrandmemberKey) Count(count int64) ZrandmemberOptionsCount {
	_ = "STUB: not implemented"
	return *new(ZrandmemberOptionsCount)
}

func (c ZrandmemberKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZrandmemberOptionsCount Incomplete

func (c ZrandmemberOptionsCount) Withscores() ZrandmemberOptionsWithscores {
	_ = "STUB: not implemented"
	return *new(ZrandmemberOptionsWithscores)
}

func (c ZrandmemberOptionsCount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ZrandmemberOptionsWithscores Incomplete

func (c ZrandmemberOptionsWithscores) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type Zrange Incomplete

func (b Builder) Zrange() (c Zrange) { _ = "STUB: not implemented"; return *new(Zrange) }

func (c Zrange) Key(key string) ZrangeKey { _ = "STUB: not implemented"; return *new(ZrangeKey) }

type ZrangeKey Incomplete

func (c ZrangeKey) Min(min string) ZrangeMin { _ = "STUB: not implemented"; return *new(ZrangeMin) }

type ZrangeLimit Incomplete

func (c ZrangeLimit) Withscores() ZrangeWithscores {
	_ = "STUB: not implemented"
	return *new(ZrangeWithscores)
}

func (c ZrangeLimit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrangeLimit) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrangeMax Incomplete

func (c ZrangeMax) Byscore() ZrangeSortbyByscore {
	_ = "STUB: not implemented"
	return *new(ZrangeSortbyByscore)
}

func (c ZrangeMax) Bylex() ZrangeSortbyBylex {
	_ = "STUB: not implemented"
	return *new(ZrangeSortbyBylex)
}

func (c ZrangeMax) Rev() ZrangeRev { _ = "STUB: not implemented"; return *new(ZrangeRev) }

func (c ZrangeMax) Limit(offset int64, count int64) ZrangeLimit {
	_ = "STUB: not implemented"
	return *new(ZrangeLimit)
}

func (c ZrangeMax) Withscores() ZrangeWithscores {
	_ = "STUB: not implemented"
	return *new(ZrangeWithscores)
}

func (c ZrangeMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrangeMax) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrangeMin Incomplete

func (c ZrangeMin) Max(max string) ZrangeMax { _ = "STUB: not implemented"; return *new(ZrangeMax) }

type ZrangeRev Incomplete

func (c ZrangeRev) Limit(offset int64, count int64) ZrangeLimit {
	_ = "STUB: not implemented"
	return *new(ZrangeLimit)
}

func (c ZrangeRev) Withscores() ZrangeWithscores {
	_ = "STUB: not implemented"
	return *new(ZrangeWithscores)
}

func (c ZrangeRev) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrangeRev) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrangeSortbyBylex Incomplete

func (c ZrangeSortbyBylex) Rev() ZrangeRev { _ = "STUB: not implemented"; return *new(ZrangeRev) }

func (c ZrangeSortbyBylex) Limit(offset int64, count int64) ZrangeLimit {
	_ = "STUB: not implemented"
	return *new(ZrangeLimit)
}

func (c ZrangeSortbyBylex) Withscores() ZrangeWithscores {
	_ = "STUB: not implemented"
	return *new(ZrangeWithscores)
}

func (c ZrangeSortbyBylex) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrangeSortbyBylex) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrangeSortbyByscore Incomplete

func (c ZrangeSortbyByscore) Rev() ZrangeRev { _ = "STUB: not implemented"; return *new(ZrangeRev) }

func (c ZrangeSortbyByscore) Limit(offset int64, count int64) ZrangeLimit {
	_ = "STUB: not implemented"
	return *new(ZrangeLimit)
}

func (c ZrangeSortbyByscore) Withscores() ZrangeWithscores {
	_ = "STUB: not implemented"
	return *new(ZrangeWithscores)
}

func (c ZrangeSortbyByscore) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrangeSortbyByscore) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrangeWithscores Incomplete

func (c ZrangeWithscores) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrangeWithscores) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Zrangebylex Incomplete

func (b Builder) Zrangebylex() (c Zrangebylex) { _ = "STUB: not implemented"; return *new(Zrangebylex) }

func (c Zrangebylex) Key(key string) ZrangebylexKey {
	_ = "STUB: not implemented"
	return *new(ZrangebylexKey)
}

type ZrangebylexKey Incomplete

func (c ZrangebylexKey) Min(min string) ZrangebylexMin {
	_ = "STUB: not implemented"
	return *new(ZrangebylexMin)
}

type ZrangebylexLimit Incomplete

func (c ZrangebylexLimit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrangebylexLimit) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrangebylexMax Incomplete

func (c ZrangebylexMax) Limit(offset int64, count int64) ZrangebylexLimit {
	_ = "STUB: not implemented"
	return *new(ZrangebylexLimit)
}

func (c ZrangebylexMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrangebylexMax) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrangebylexMin Incomplete

func (c ZrangebylexMin) Max(max string) ZrangebylexMax {
	_ = "STUB: not implemented"
	return *new(ZrangebylexMax)
}

type Zrangebyscore Incomplete

func (b Builder) Zrangebyscore() (c Zrangebyscore) {
	_ = "STUB: not implemented"
	return *new(Zrangebyscore)
}

func (c Zrangebyscore) Key(key string) ZrangebyscoreKey {
	_ = "STUB: not implemented"
	return *new(ZrangebyscoreKey)
}

type ZrangebyscoreKey Incomplete

func (c ZrangebyscoreKey) Min(min string) ZrangebyscoreMin {
	_ = "STUB: not implemented"
	return *new(ZrangebyscoreMin)
}

type ZrangebyscoreLimit Incomplete

func (c ZrangebyscoreLimit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrangebyscoreLimit) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrangebyscoreMax Incomplete

func (c ZrangebyscoreMax) Withscores() ZrangebyscoreWithscores {
	_ = "STUB: not implemented"
	return *new(ZrangebyscoreWithscores)
}

func (c ZrangebyscoreMax) Limit(offset int64, count int64) ZrangebyscoreLimit {
	_ = "STUB: not implemented"
	return *new(ZrangebyscoreLimit)
}

func (c ZrangebyscoreMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrangebyscoreMax) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrangebyscoreMin Incomplete

func (c ZrangebyscoreMin) Max(max string) ZrangebyscoreMax {
	_ = "STUB: not implemented"
	return *new(ZrangebyscoreMax)
}

type ZrangebyscoreWithscores Incomplete

func (c ZrangebyscoreWithscores) Limit(offset int64, count int64) ZrangebyscoreLimit {
	_ = "STUB: not implemented"
	return *new(ZrangebyscoreLimit)
}

func (c ZrangebyscoreWithscores) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c ZrangebyscoreWithscores) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type Zrangestore Incomplete

func (b Builder) Zrangestore() (c Zrangestore) { _ = "STUB: not implemented"; return *new(Zrangestore) }

func (c Zrangestore) Dst(dst string) ZrangestoreDst {
	_ = "STUB: not implemented"
	return *new(ZrangestoreDst)
}

type ZrangestoreDst Incomplete

func (c ZrangestoreDst) Src(src string) ZrangestoreSrc {
	_ = "STUB: not implemented"
	return *new(ZrangestoreSrc)
}

type ZrangestoreLimit Incomplete

func (c ZrangestoreLimit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZrangestoreMax Incomplete

func (c ZrangestoreMax) Byscore() ZrangestoreSortbyByscore {
	_ = "STUB: not implemented"
	return *new(ZrangestoreSortbyByscore)
}

func (c ZrangestoreMax) Bylex() ZrangestoreSortbyBylex {
	_ = "STUB: not implemented"
	return *new(ZrangestoreSortbyBylex)
}

func (c ZrangestoreMax) Rev() ZrangestoreRev {
	_ = "STUB: not implemented"
	return *new(ZrangestoreRev)
}

func (c ZrangestoreMax) Limit(offset int64, count int64) ZrangestoreLimit {
	_ = "STUB: not implemented"
	return *new(ZrangestoreLimit)
}

func (c ZrangestoreMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZrangestoreMin Incomplete

func (c ZrangestoreMin) Max(max string) ZrangestoreMax {
	_ = "STUB: not implemented"
	return *new(ZrangestoreMax)
}

type ZrangestoreRev Incomplete

func (c ZrangestoreRev) Limit(offset int64, count int64) ZrangestoreLimit {
	_ = "STUB: not implemented"
	return *new(ZrangestoreLimit)
}

func (c ZrangestoreRev) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZrangestoreSortbyBylex Incomplete

func (c ZrangestoreSortbyBylex) Rev() ZrangestoreRev {
	_ = "STUB: not implemented"
	return *new(ZrangestoreRev)
}

func (c ZrangestoreSortbyBylex) Limit(offset int64, count int64) ZrangestoreLimit {
	_ = "STUB: not implemented"
	return *new(ZrangestoreLimit)
}

func (c ZrangestoreSortbyBylex) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ZrangestoreSortbyByscore Incomplete

func (c ZrangestoreSortbyByscore) Rev() ZrangestoreRev {
	_ = "STUB: not implemented"
	return *new(ZrangestoreRev)
}

func (c ZrangestoreSortbyByscore) Limit(offset int64, count int64) ZrangestoreLimit {
	_ = "STUB: not implemented"
	return *new(ZrangestoreLimit)
}

func (c ZrangestoreSortbyByscore) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ZrangestoreSrc Incomplete

func (c ZrangestoreSrc) Min(min string) ZrangestoreMin {
	_ = "STUB: not implemented"
	return *new(ZrangestoreMin)
}

type Zrank Incomplete

func (b Builder) Zrank() (c Zrank) { _ = "STUB: not implemented"; return *new(Zrank) }

func (c Zrank) Key(key string) ZrankKey { _ = "STUB: not implemented"; return *new(ZrankKey) }

type ZrankKey Incomplete

func (c ZrankKey) Member(member string) ZrankMember {
	_ = "STUB: not implemented"
	return *new(ZrankMember)
}

type ZrankMember Incomplete

func (c ZrankMember) Withscore() ZrankWithscore {
	_ = "STUB: not implemented"
	return *new(ZrankWithscore)
}

func (c ZrankMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrankMember) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrankWithscore Incomplete

func (c ZrankWithscore) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrankWithscore) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Zrem Incomplete

func (b Builder) Zrem() (c Zrem) { _ = "STUB: not implemented"; return *new(Zrem) }

func (c Zrem) Key(key string) ZremKey { _ = "STUB: not implemented"; return *new(ZremKey) }

type ZremKey Incomplete

func (c ZremKey) Member(member ...string) ZremMember {
	_ = "STUB: not implemented"
	return *new(ZremMember)
}

type ZremMember Incomplete

func (c ZremMember) Member(member ...string) ZremMember {
	_ = "STUB: not implemented"
	return *new(ZremMember)
}

func (c ZremMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zremrangebylex Incomplete

func (b Builder) Zremrangebylex() (c Zremrangebylex) {
	_ = "STUB: not implemented"
	return *new(Zremrangebylex)
}

func (c Zremrangebylex) Key(key string) ZremrangebylexKey {
	_ = "STUB: not implemented"
	return *new(ZremrangebylexKey)
}

type ZremrangebylexKey Incomplete

func (c ZremrangebylexKey) Min(min string) ZremrangebylexMin {
	_ = "STUB: not implemented"
	return *new(ZremrangebylexMin)
}

type ZremrangebylexMax Incomplete

func (c ZremrangebylexMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZremrangebylexMin Incomplete

func (c ZremrangebylexMin) Max(max string) ZremrangebylexMax {
	_ = "STUB: not implemented"
	return *new(ZremrangebylexMax)
}

type Zremrangebyrank Incomplete

func (b Builder) Zremrangebyrank() (c Zremrangebyrank) {
	_ = "STUB: not implemented"
	return *new(Zremrangebyrank)
}

func (c Zremrangebyrank) Key(key string) ZremrangebyrankKey {
	_ = "STUB: not implemented"
	return *new(ZremrangebyrankKey)
}

type ZremrangebyrankKey Incomplete

func (c ZremrangebyrankKey) Start(start int64) ZremrangebyrankStart {
	_ = "STUB: not implemented"
	return *new(ZremrangebyrankStart)
}

type ZremrangebyrankStart Incomplete

func (c ZremrangebyrankStart) Stop(stop int64) ZremrangebyrankStop {
	_ = "STUB: not implemented"
	return *new(ZremrangebyrankStop)
}

type ZremrangebyrankStop Incomplete

func (c ZremrangebyrankStop) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zremrangebyscore Incomplete

func (b Builder) Zremrangebyscore() (c Zremrangebyscore) {
	_ = "STUB: not implemented"
	return *new(Zremrangebyscore)
}

func (c Zremrangebyscore) Key(key string) ZremrangebyscoreKey {
	_ = "STUB: not implemented"
	return *new(ZremrangebyscoreKey)
}

type ZremrangebyscoreKey Incomplete

func (c ZremrangebyscoreKey) Min(min string) ZremrangebyscoreMin {
	_ = "STUB: not implemented"
	return *new(ZremrangebyscoreMin)
}

type ZremrangebyscoreMax Incomplete

func (c ZremrangebyscoreMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZremrangebyscoreMin Incomplete

func (c ZremrangebyscoreMin) Max(max string) ZremrangebyscoreMax {
	_ = "STUB: not implemented"
	return *new(ZremrangebyscoreMax)
}

type Zrevrange Incomplete

func (b Builder) Zrevrange() (c Zrevrange) { _ = "STUB: not implemented"; return *new(Zrevrange) }

func (c Zrevrange) Key(key string) ZrevrangeKey {
	_ = "STUB: not implemented"
	return *new(ZrevrangeKey)
}

type ZrevrangeKey Incomplete

func (c ZrevrangeKey) Start(start int64) ZrevrangeStart {
	_ = "STUB: not implemented"
	return *new(ZrevrangeStart)
}

type ZrevrangeStart Incomplete

func (c ZrevrangeStart) Stop(stop int64) ZrevrangeStop {
	_ = "STUB: not implemented"
	return *new(ZrevrangeStop)
}

type ZrevrangeStop Incomplete

func (c ZrevrangeStop) Withscores() ZrevrangeWithscores {
	_ = "STUB: not implemented"
	return *new(ZrevrangeWithscores)
}

func (c ZrevrangeStop) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrevrangeStop) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrevrangeWithscores Incomplete

func (c ZrevrangeWithscores) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrevrangeWithscores) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Zrevrangebylex Incomplete

func (b Builder) Zrevrangebylex() (c Zrevrangebylex) {
	_ = "STUB: not implemented"
	return *new(Zrevrangebylex)
}

func (c Zrevrangebylex) Key(key string) ZrevrangebylexKey {
	_ = "STUB: not implemented"
	return *new(ZrevrangebylexKey)
}

type ZrevrangebylexKey Incomplete

func (c ZrevrangebylexKey) Max(max string) ZrevrangebylexMax {
	_ = "STUB: not implemented"
	return *new(ZrevrangebylexMax)
}

type ZrevrangebylexLimit Incomplete

func (c ZrevrangebylexLimit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrevrangebylexLimit) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrevrangebylexMax Incomplete

func (c ZrevrangebylexMax) Min(min string) ZrevrangebylexMin {
	_ = "STUB: not implemented"
	return *new(ZrevrangebylexMin)
}

type ZrevrangebylexMin Incomplete

func (c ZrevrangebylexMin) Limit(offset int64, count int64) ZrevrangebylexLimit {
	_ = "STUB: not implemented"
	return *new(ZrevrangebylexLimit)
}

func (c ZrevrangebylexMin) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrevrangebylexMin) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Zrevrangebyscore Incomplete

func (b Builder) Zrevrangebyscore() (c Zrevrangebyscore) {
	_ = "STUB: not implemented"
	return *new(Zrevrangebyscore)
}

func (c Zrevrangebyscore) Key(key string) ZrevrangebyscoreKey {
	_ = "STUB: not implemented"
	return *new(ZrevrangebyscoreKey)
}

type ZrevrangebyscoreKey Incomplete

func (c ZrevrangebyscoreKey) Max(max string) ZrevrangebyscoreMax {
	_ = "STUB: not implemented"
	return *new(ZrevrangebyscoreMax)
}

type ZrevrangebyscoreLimit Incomplete

func (c ZrevrangebyscoreLimit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrevrangebyscoreLimit) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrevrangebyscoreMax Incomplete

func (c ZrevrangebyscoreMax) Min(min string) ZrevrangebyscoreMin {
	_ = "STUB: not implemented"
	return *new(ZrevrangebyscoreMin)
}

type ZrevrangebyscoreMin Incomplete

func (c ZrevrangebyscoreMin) Withscores() ZrevrangebyscoreWithscores {
	_ = "STUB: not implemented"
	return *new(ZrevrangebyscoreWithscores)
}

func (c ZrevrangebyscoreMin) Limit(offset int64, count int64) ZrevrangebyscoreLimit {
	_ = "STUB: not implemented"
	return *new(ZrevrangebyscoreLimit)
}

func (c ZrevrangebyscoreMin) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrevrangebyscoreMin) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrevrangebyscoreWithscores Incomplete

func (c ZrevrangebyscoreWithscores) Limit(offset int64, count int64) ZrevrangebyscoreLimit {
	_ = "STUB: not implemented"
	return *new(ZrevrangebyscoreLimit)
}

func (c ZrevrangebyscoreWithscores) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c ZrevrangebyscoreWithscores) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type Zrevrank Incomplete

func (b Builder) Zrevrank() (c Zrevrank) { _ = "STUB: not implemented"; return *new(Zrevrank) }

func (c Zrevrank) Key(key string) ZrevrankKey { _ = "STUB: not implemented"; return *new(ZrevrankKey) }

type ZrevrankKey Incomplete

func (c ZrevrankKey) Member(member string) ZrevrankMember {
	_ = "STUB: not implemented"
	return *new(ZrevrankMember)
}

type ZrevrankMember Incomplete

func (c ZrevrankMember) Withscore() ZrevrankWithscore {
	_ = "STUB: not implemented"
	return *new(ZrevrankWithscore)
}

func (c ZrevrankMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrevrankMember) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type ZrevrankWithscore Incomplete

func (c ZrevrankWithscore) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZrevrankWithscore) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Zscan Incomplete

func (b Builder) Zscan() (c Zscan) { _ = "STUB: not implemented"; return *new(Zscan) }

func (c Zscan) Key(key string) ZscanKey { _ = "STUB: not implemented"; return *new(ZscanKey) }

type ZscanCount Incomplete

func (c ZscanCount) Noscores() ZscanNoscores { _ = "STUB: not implemented"; return *new(ZscanNoscores) }

func (c ZscanCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZscanCursor Incomplete

func (c ZscanCursor) Match(pattern string) ZscanMatch {
	_ = "STUB: not implemented"
	return *new(ZscanMatch)
}

func (c ZscanCursor) Count(count int64) ZscanCount {
	_ = "STUB: not implemented"
	return *new(ZscanCount)
}

func (c ZscanCursor) Noscores() ZscanNoscores {
	_ = "STUB: not implemented"
	return *new(ZscanNoscores)
}

func (c ZscanCursor) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZscanKey Incomplete

func (c ZscanKey) Cursor(cursor uint64) ZscanCursor {
	_ = "STUB: not implemented"
	return *new(ZscanCursor)
}

type ZscanMatch Incomplete

func (c ZscanMatch) Count(count int64) ZscanCount {
	_ = "STUB: not implemented"
	return *new(ZscanCount)
}

func (c ZscanMatch) Noscores() ZscanNoscores { _ = "STUB: not implemented"; return *new(ZscanNoscores) }

func (c ZscanMatch) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZscanNoscores Incomplete

func (c ZscanNoscores) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zscore Incomplete

func (b Builder) Zscore() (c Zscore) { _ = "STUB: not implemented"; return *new(Zscore) }

func (c Zscore) Key(key string) ZscoreKey { _ = "STUB: not implemented"; return *new(ZscoreKey) }

type ZscoreKey Incomplete

func (c ZscoreKey) Member(member string) ZscoreMember {
	_ = "STUB: not implemented"
	return *new(ZscoreMember)
}

type ZscoreMember Incomplete

func (c ZscoreMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c ZscoreMember) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Zunion Incomplete

func (b Builder) Zunion() (c Zunion) { _ = "STUB: not implemented"; return *new(Zunion) }

func (c Zunion) Numkeys(numkeys int64) ZunionNumkeys {
	_ = "STUB: not implemented"
	return *new(ZunionNumkeys)
}

type ZunionAggregateCount Incomplete

func (c ZunionAggregateCount) Withscores() ZunionWithscores {
	_ = "STUB: not implemented"
	return *new(ZunionWithscores)
}

func (c ZunionAggregateCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZunionAggregateMax Incomplete

func (c ZunionAggregateMax) Withscores() ZunionWithscores {
	_ = "STUB: not implemented"
	return *new(ZunionWithscores)
}

func (c ZunionAggregateMax) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZunionAggregateMin Incomplete

func (c ZunionAggregateMin) Withscores() ZunionWithscores {
	_ = "STUB: not implemented"
	return *new(ZunionWithscores)
}

func (c ZunionAggregateMin) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZunionAggregateSum Incomplete

func (c ZunionAggregateSum) Withscores() ZunionWithscores {
	_ = "STUB: not implemented"
	return *new(ZunionWithscores)
}

func (c ZunionAggregateSum) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZunionKey Incomplete

func (c ZunionKey) Key(key ...string) ZunionKey { _ = "STUB: not implemented"; return *new(ZunionKey) }

func (c ZunionKey) Weights(weight ...int64) ZunionWeights {
	_ = "STUB: not implemented"
	return *new(ZunionWeights)
}

func (c ZunionKey) AggregateSum() ZunionAggregateSum {
	_ = "STUB: not implemented"
	return *new(ZunionAggregateSum)
}

func (c ZunionKey) AggregateMin() ZunionAggregateMin {
	_ = "STUB: not implemented"
	return *new(ZunionAggregateMin)
}

func (c ZunionKey) AggregateMax() ZunionAggregateMax {
	_ = "STUB: not implemented"
	return *new(ZunionAggregateMax)
}

func (c ZunionKey) AggregateCount() ZunionAggregateCount {
	_ = "STUB: not implemented"
	return *new(ZunionAggregateCount)
}

func (c ZunionKey) Withscores() ZunionWithscores {
	_ = "STUB: not implemented"
	return *new(ZunionWithscores)
}

func (c ZunionKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZunionNumkeys Incomplete

func (c ZunionNumkeys) Key(key ...string) ZunionKey {
	_ = "STUB: not implemented"
	return *new(ZunionKey)
}

type ZunionWeights Incomplete

func (c ZunionWeights) Weights(weight ...int64) ZunionWeights {
	_ = "STUB: not implemented"
	return *new(ZunionWeights)
}

func (c ZunionWeights) AggregateSum() ZunionAggregateSum {
	_ = "STUB: not implemented"
	return *new(ZunionAggregateSum)
}

func (c ZunionWeights) AggregateMin() ZunionAggregateMin {
	_ = "STUB: not implemented"
	return *new(ZunionAggregateMin)
}

func (c ZunionWeights) AggregateMax() ZunionAggregateMax {
	_ = "STUB: not implemented"
	return *new(ZunionAggregateMax)
}

func (c ZunionWeights) AggregateCount() ZunionAggregateCount {
	_ = "STUB: not implemented"
	return *new(ZunionAggregateCount)
}

func (c ZunionWeights) Withscores() ZunionWithscores {
	_ = "STUB: not implemented"
	return *new(ZunionWithscores)
}

func (c ZunionWeights) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZunionWithscores Incomplete

func (c ZunionWithscores) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Zunionstore Incomplete

func (b Builder) Zunionstore() (c Zunionstore) { _ = "STUB: not implemented"; return *new(Zunionstore) }

func (c Zunionstore) Destination(destination string) ZunionstoreDestination {
	_ = "STUB: not implemented"
	return *new(ZunionstoreDestination)
}

type ZunionstoreAggregateCount Incomplete

func (c ZunionstoreAggregateCount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ZunionstoreAggregateMax Incomplete

func (c ZunionstoreAggregateMax) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ZunionstoreAggregateMin Incomplete

func (c ZunionstoreAggregateMin) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ZunionstoreAggregateSum Incomplete

func (c ZunionstoreAggregateSum) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ZunionstoreDestination Incomplete

func (c ZunionstoreDestination) Numkeys(numkeys int64) ZunionstoreNumkeys {
	_ = "STUB: not implemented"
	return *new(ZunionstoreNumkeys)
}

type ZunionstoreKey Incomplete

func (c ZunionstoreKey) Key(key ...string) ZunionstoreKey {
	_ = "STUB: not implemented"
	return *new(ZunionstoreKey)
}

func (c ZunionstoreKey) Weights(weight ...int64) ZunionstoreWeights {
	_ = "STUB: not implemented"
	return *new(ZunionstoreWeights)
}

func (c ZunionstoreKey) AggregateSum() ZunionstoreAggregateSum {
	_ = "STUB: not implemented"
	return *new(ZunionstoreAggregateSum)
}

func (c ZunionstoreKey) AggregateMin() ZunionstoreAggregateMin {
	_ = "STUB: not implemented"
	return *new(ZunionstoreAggregateMin)
}

func (c ZunionstoreKey) AggregateMax() ZunionstoreAggregateMax {
	_ = "STUB: not implemented"
	return *new(ZunionstoreAggregateMax)
}

func (c ZunionstoreKey) AggregateCount() ZunionstoreAggregateCount {
	_ = "STUB: not implemented"
	return *new(ZunionstoreAggregateCount)
}

func (c ZunionstoreKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ZunionstoreNumkeys Incomplete

func (c ZunionstoreNumkeys) Key(key ...string) ZunionstoreKey {
	_ = "STUB: not implemented"
	return *new(ZunionstoreKey)
}

type ZunionstoreWeights Incomplete

func (c ZunionstoreWeights) Weights(weight ...int64) ZunionstoreWeights {
	_ = "STUB: not implemented"
	return *new(ZunionstoreWeights)
}

func (c ZunionstoreWeights) AggregateSum() ZunionstoreAggregateSum {
	_ = "STUB: not implemented"
	return *new(ZunionstoreAggregateSum)
}

func (c ZunionstoreWeights) AggregateMin() ZunionstoreAggregateMin {
	_ = "STUB: not implemented"
	return *new(ZunionstoreAggregateMin)
}

func (c ZunionstoreWeights) AggregateMax() ZunionstoreAggregateMax {
	_ = "STUB: not implemented"
	return *new(ZunionstoreAggregateMax)
}

func (c ZunionstoreWeights) AggregateCount() ZunionstoreAggregateCount {
	_ = "STUB: not implemented"
	return *new(ZunionstoreAggregateCount)
}

func (c ZunionstoreWeights) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }
