// Code generated DO NOT EDIT

package cmds

type Xack Incomplete

func (b Builder) Xack() (c Xack) { _ = "STUB: not implemented"; return *new(Xack) }

func (c Xack) Key(key string) XackKey { _ = "STUB: not implemented"; return *new(XackKey) }

type XackGroup Incomplete

func (c XackGroup) Id(id ...string) XackId { _ = "STUB: not implemented"; return *new(XackId) }

type XackId Incomplete

func (c XackId) Id(id ...string) XackId { _ = "STUB: not implemented"; return *new(XackId) }

func (c XackId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XackKey Incomplete

func (c XackKey) Group(group string) XackGroup { _ = "STUB: not implemented"; return *new(XackGroup) }

type Xackdel Incomplete

func (b Builder) Xackdel() (c Xackdel) { _ = "STUB: not implemented"; return *new(Xackdel) }

func (c Xackdel) Key(key string) XackdelKey { _ = "STUB: not implemented"; return *new(XackdelKey) }

type XackdelActionAcked Incomplete

func (c XackdelActionAcked) Ids() XackdelIdsBlockIds {
	_ = "STUB: not implemented"
	return *new(XackdelIdsBlockIds)
}

type XackdelActionDelref Incomplete

func (c XackdelActionDelref) Ids() XackdelIdsBlockIds {
	_ = "STUB: not implemented"
	return *new(XackdelIdsBlockIds)
}

type XackdelActionKeepref Incomplete

func (c XackdelActionKeepref) Ids() XackdelIdsBlockIds {
	_ = "STUB: not implemented"
	return *new(XackdelIdsBlockIds)
}

type XackdelGroup Incomplete

func (c XackdelGroup) Keepref() XackdelActionKeepref {
	_ = "STUB: not implemented"
	return *new(XackdelActionKeepref)
}

func (c XackdelGroup) Delref() XackdelActionDelref {
	_ = "STUB: not implemented"
	return *new(XackdelActionDelref)
}

func (c XackdelGroup) Acked() XackdelActionAcked {
	_ = "STUB: not implemented"
	return *new(XackdelActionAcked)
}

func (c XackdelGroup) Ids() XackdelIdsBlockIds {
	_ = "STUB: not implemented"
	return *new(XackdelIdsBlockIds)
}

type XackdelIdsBlockId Incomplete

func (c XackdelIdsBlockId) Id(id ...string) XackdelIdsBlockId {
	_ = "STUB: not implemented"
	return *new(XackdelIdsBlockId)
}

func (c XackdelIdsBlockId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XackdelIdsBlockIds Incomplete

func (c XackdelIdsBlockIds) Numids(numids int64) XackdelIdsBlockNumids {
	_ = "STUB: not implemented"
	return *new(XackdelIdsBlockNumids)
}

type XackdelIdsBlockNumids Incomplete

func (c XackdelIdsBlockNumids) Id(id ...string) XackdelIdsBlockId {
	_ = "STUB: not implemented"
	return *new(XackdelIdsBlockId)
}

type XackdelKey Incomplete

func (c XackdelKey) Group(group string) XackdelGroup {
	_ = "STUB: not implemented"
	return *new(XackdelGroup)
}

type Xadd Incomplete

func (b Builder) Xadd() (c Xadd) { _ = "STUB: not implemented"; return *new(Xadd) }

func (c Xadd) Key(key string) XaddKey { _ = "STUB: not implemented"; return *new(XaddKey) }

type XaddConditionAcked Incomplete

func (c XaddConditionAcked) Idmpauto() XaddIdmpIdmpautoIdmpauto {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpautoIdmpauto)
}

func (c XaddConditionAcked) Idmp() XaddIdmpIdmpIdmp {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpIdmp)
}

func (c XaddConditionAcked) Maxlen() XaddTrimStrategyMaxlen {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMaxlen)
}

func (c XaddConditionAcked) Minid() XaddTrimStrategyMinid {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMinid)
}

func (c XaddConditionAcked) Id(id string) XaddId { _ = "STUB: not implemented"; return *new(XaddId) }

type XaddConditionDelref Incomplete

func (c XaddConditionDelref) Idmpauto() XaddIdmpIdmpautoIdmpauto {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpautoIdmpauto)
}

func (c XaddConditionDelref) Idmp() XaddIdmpIdmpIdmp {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpIdmp)
}

func (c XaddConditionDelref) Maxlen() XaddTrimStrategyMaxlen {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMaxlen)
}

func (c XaddConditionDelref) Minid() XaddTrimStrategyMinid {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMinid)
}

func (c XaddConditionDelref) Id(id string) XaddId { _ = "STUB: not implemented"; return *new(XaddId) }

type XaddConditionKeepref Incomplete

func (c XaddConditionKeepref) Idmpauto() XaddIdmpIdmpautoIdmpauto {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpautoIdmpauto)
}

func (c XaddConditionKeepref) Idmp() XaddIdmpIdmpIdmp {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpIdmp)
}

func (c XaddConditionKeepref) Maxlen() XaddTrimStrategyMaxlen {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMaxlen)
}

func (c XaddConditionKeepref) Minid() XaddTrimStrategyMinid {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMinid)
}

func (c XaddConditionKeepref) Id(id string) XaddId { _ = "STUB: not implemented"; return *new(XaddId) }

type XaddFieldValue Incomplete

func (c XaddFieldValue) FieldValue(field string, value string) XaddFieldValue {
	_ = "STUB: not implemented"
	return *new(XaddFieldValue)
}

func (c XaddFieldValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XaddId Incomplete

func (c XaddId) FieldValue() XaddFieldValue { _ = "STUB: not implemented"; return *new(XaddFieldValue) }

type XaddIdmpIdmpIdmp Incomplete

func (c XaddIdmpIdmpIdmp) Pid(pid string) XaddIdmpIdmpPid {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpPid)
}

type XaddIdmpIdmpIid Incomplete

func (c XaddIdmpIdmpIid) Maxlen() XaddTrimStrategyMaxlen {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMaxlen)
}

func (c XaddIdmpIdmpIid) Minid() XaddTrimStrategyMinid {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMinid)
}

func (c XaddIdmpIdmpIid) Id(id string) XaddId { _ = "STUB: not implemented"; return *new(XaddId) }

type XaddIdmpIdmpPid Incomplete

func (c XaddIdmpIdmpPid) Iid(iid string) XaddIdmpIdmpIid {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpIid)
}

type XaddIdmpIdmpautoIdmpauto Incomplete

func (c XaddIdmpIdmpautoIdmpauto) Pid(pid string) XaddIdmpIdmpautoPid {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpautoPid)
}

type XaddIdmpIdmpautoPid Incomplete

func (c XaddIdmpIdmpautoPid) Maxlen() XaddTrimStrategyMaxlen {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMaxlen)
}

func (c XaddIdmpIdmpautoPid) Minid() XaddTrimStrategyMinid {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMinid)
}

func (c XaddIdmpIdmpautoPid) Id(id string) XaddId { _ = "STUB: not implemented"; return *new(XaddId) }

type XaddKey Incomplete

func (c XaddKey) Nomkstream() XaddNomkstream {
	_ = "STUB: not implemented"
	return *new(XaddNomkstream)
}

func (c XaddKey) Keepref() XaddConditionKeepref {
	_ = "STUB: not implemented"
	return *new(XaddConditionKeepref)
}

func (c XaddKey) Delref() XaddConditionDelref {
	_ = "STUB: not implemented"
	return *new(XaddConditionDelref)
}

func (c XaddKey) Acked() XaddConditionAcked {
	_ = "STUB: not implemented"
	return *new(XaddConditionAcked)
}

func (c XaddKey) Idmpauto() XaddIdmpIdmpautoIdmpauto {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpautoIdmpauto)
}

func (c XaddKey) Idmp() XaddIdmpIdmpIdmp { _ = "STUB: not implemented"; return *new(XaddIdmpIdmpIdmp) }

func (c XaddKey) Maxlen() XaddTrimStrategyMaxlen {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMaxlen)
}

func (c XaddKey) Minid() XaddTrimStrategyMinid {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMinid)
}

func (c XaddKey) Id(id string) XaddId { _ = "STUB: not implemented"; return *new(XaddId) }

type XaddNomkstream Incomplete

func (c XaddNomkstream) Keepref() XaddConditionKeepref {
	_ = "STUB: not implemented"
	return *new(XaddConditionKeepref)
}

func (c XaddNomkstream) Delref() XaddConditionDelref {
	_ = "STUB: not implemented"
	return *new(XaddConditionDelref)
}

func (c XaddNomkstream) Acked() XaddConditionAcked {
	_ = "STUB: not implemented"
	return *new(XaddConditionAcked)
}

func (c XaddNomkstream) Idmpauto() XaddIdmpIdmpautoIdmpauto {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpautoIdmpauto)
}

func (c XaddNomkstream) Idmp() XaddIdmpIdmpIdmp {
	_ = "STUB: not implemented"
	return *new(XaddIdmpIdmpIdmp)
}

func (c XaddNomkstream) Maxlen() XaddTrimStrategyMaxlen {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMaxlen)
}

func (c XaddNomkstream) Minid() XaddTrimStrategyMinid {
	_ = "STUB: not implemented"
	return *new(XaddTrimStrategyMinid)
}

func (c XaddNomkstream) Id(id string) XaddId { _ = "STUB: not implemented"; return *new(XaddId) }

type XaddTrimLimit Incomplete

func (c XaddTrimLimit) Id(id string) XaddId { _ = "STUB: not implemented"; return *new(XaddId) }

type XaddTrimOperatorAlmost Incomplete

func (c XaddTrimOperatorAlmost) Threshold(threshold string) XaddTrimThreshold {
	_ = "STUB: not implemented"
	return *new(XaddTrimThreshold)
}

type XaddTrimOperatorExact Incomplete

func (c XaddTrimOperatorExact) Threshold(threshold string) XaddTrimThreshold {
	_ = "STUB: not implemented"
	return *new(XaddTrimThreshold)
}

type XaddTrimStrategyMaxlen Incomplete

func (c XaddTrimStrategyMaxlen) Exact() XaddTrimOperatorExact {
	_ = "STUB: not implemented"
	return *new(XaddTrimOperatorExact)
}

func (c XaddTrimStrategyMaxlen) Almost() XaddTrimOperatorAlmost {
	_ = "STUB: not implemented"
	return *new(XaddTrimOperatorAlmost)
}

func (c XaddTrimStrategyMaxlen) Threshold(threshold string) XaddTrimThreshold {
	_ = "STUB: not implemented"
	return *new(XaddTrimThreshold)
}

type XaddTrimStrategyMinid Incomplete

func (c XaddTrimStrategyMinid) Exact() XaddTrimOperatorExact {
	_ = "STUB: not implemented"
	return *new(XaddTrimOperatorExact)
}

func (c XaddTrimStrategyMinid) Almost() XaddTrimOperatorAlmost {
	_ = "STUB: not implemented"
	return *new(XaddTrimOperatorAlmost)
}

func (c XaddTrimStrategyMinid) Threshold(threshold string) XaddTrimThreshold {
	_ = "STUB: not implemented"
	return *new(XaddTrimThreshold)
}

type XaddTrimThreshold Incomplete

func (c XaddTrimThreshold) Limit(count int64) XaddTrimLimit {
	_ = "STUB: not implemented"
	return *new(XaddTrimLimit)
}

func (c XaddTrimThreshold) Id(id string) XaddId { _ = "STUB: not implemented"; return *new(XaddId) }

type Xautoclaim Incomplete

func (b Builder) Xautoclaim() (c Xautoclaim) { _ = "STUB: not implemented"; return *new(Xautoclaim) }

func (c Xautoclaim) Key(key string) XautoclaimKey {
	_ = "STUB: not implemented"
	return *new(XautoclaimKey)
}

type XautoclaimConsumer Incomplete

func (c XautoclaimConsumer) MinIdleTime(minIdleTime string) XautoclaimMinIdleTime {
	_ = "STUB: not implemented"
	return *new(XautoclaimMinIdleTime)
}

type XautoclaimCount Incomplete

func (c XautoclaimCount) Justid() XautoclaimJustid {
	_ = "STUB: not implemented"
	return *new(XautoclaimJustid)
}

func (c XautoclaimCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XautoclaimGroup Incomplete

func (c XautoclaimGroup) Consumer(consumer string) XautoclaimConsumer {
	_ = "STUB: not implemented"
	return *new(XautoclaimConsumer)
}

type XautoclaimJustid Incomplete

func (c XautoclaimJustid) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XautoclaimKey Incomplete

func (c XautoclaimKey) Group(group string) XautoclaimGroup {
	_ = "STUB: not implemented"
	return *new(XautoclaimGroup)
}

type XautoclaimMinIdleTime Incomplete

func (c XautoclaimMinIdleTime) Start(start string) XautoclaimStart {
	_ = "STUB: not implemented"
	return *new(XautoclaimStart)
}

type XautoclaimStart Incomplete

func (c XautoclaimStart) Count(count int64) XautoclaimCount {
	_ = "STUB: not implemented"
	return *new(XautoclaimCount)
}

func (c XautoclaimStart) Justid() XautoclaimJustid {
	_ = "STUB: not implemented"
	return *new(XautoclaimJustid)
}

func (c XautoclaimStart) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Xcfgset Incomplete

func (b Builder) Xcfgset() (c Xcfgset) { _ = "STUB: not implemented"; return *new(Xcfgset) }

func (c Xcfgset) Key(key string) XcfgsetKey { _ = "STUB: not implemented"; return *new(XcfgsetKey) }

type XcfgsetIdmpDurationBlockDuration Incomplete

func (c XcfgsetIdmpDurationBlockDuration) IdmpMaxsize() XcfgsetIdmpMaxsizeBlockIdmpMaxsizeTokenIdmpMaxsize {
	_ = "STUB: not implemented"
	return *new(XcfgsetIdmpMaxsizeBlockIdmpMaxsizeTokenIdmpMaxsize)
}

func (c XcfgsetIdmpDurationBlockDuration) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type XcfgsetIdmpDurationBlockIdmpDurationTokenIdmpDuration Incomplete

func (c XcfgsetIdmpDurationBlockIdmpDurationTokenIdmpDuration) Duration(duration int64) XcfgsetIdmpDurationBlockDuration {
	_ = "STUB: not implemented"
	return *new(XcfgsetIdmpDurationBlockDuration)
}

type XcfgsetIdmpMaxsizeBlockIdmpMaxsizeTokenIdmpMaxsize Incomplete

func (c XcfgsetIdmpMaxsizeBlockIdmpMaxsizeTokenIdmpMaxsize) Maxsize(maxsize int64) XcfgsetIdmpMaxsizeBlockMaxsize {
	_ = "STUB: not implemented"
	return *new(XcfgsetIdmpMaxsizeBlockMaxsize)
}

type XcfgsetIdmpMaxsizeBlockMaxsize Incomplete

func (c XcfgsetIdmpMaxsizeBlockMaxsize) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type XcfgsetKey Incomplete

func (c XcfgsetKey) IdmpDuration() XcfgsetIdmpDurationBlockIdmpDurationTokenIdmpDuration {
	_ = "STUB: not implemented"
	return *new(XcfgsetIdmpDurationBlockIdmpDurationTokenIdmpDuration)
}

func (c XcfgsetKey) IdmpMaxsize() XcfgsetIdmpMaxsizeBlockIdmpMaxsizeTokenIdmpMaxsize {
	_ = "STUB: not implemented"
	return *new(XcfgsetIdmpMaxsizeBlockIdmpMaxsizeTokenIdmpMaxsize)
}

func (c XcfgsetKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Xclaim Incomplete

func (b Builder) Xclaim() (c Xclaim) { _ = "STUB: not implemented"; return *new(Xclaim) }

func (c Xclaim) Key(key string) XclaimKey { _ = "STUB: not implemented"; return *new(XclaimKey) }

type XclaimConsumer Incomplete

func (c XclaimConsumer) MinIdleTime(minIdleTime string) XclaimMinIdleTime {
	_ = "STUB: not implemented"
	return *new(XclaimMinIdleTime)
}

type XclaimForce Incomplete

func (c XclaimForce) Justid() XclaimJustid { _ = "STUB: not implemented"; return *new(XclaimJustid) }

func (c XclaimForce) Lastid() XclaimLastid { _ = "STUB: not implemented"; return *new(XclaimLastid) }

func (c XclaimForce) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XclaimGroup Incomplete

func (c XclaimGroup) Consumer(consumer string) XclaimConsumer {
	_ = "STUB: not implemented"
	return *new(XclaimConsumer)
}

type XclaimId Incomplete

func (c XclaimId) Id(id ...string) XclaimId { _ = "STUB: not implemented"; return *new(XclaimId) }

func (c XclaimId) Idle(ms int64) XclaimIdle { _ = "STUB: not implemented"; return *new(XclaimIdle) }

func (c XclaimId) Time(msUnixTime int64) XclaimTime {
	_ = "STUB: not implemented"
	return *new(XclaimTime)
}

func (c XclaimId) Retrycount(count int64) XclaimRetrycount {
	_ = "STUB: not implemented"
	return *new(XclaimRetrycount)
}

func (c XclaimId) Force() XclaimForce { _ = "STUB: not implemented"; return *new(XclaimForce) }

func (c XclaimId) Justid() XclaimJustid { _ = "STUB: not implemented"; return *new(XclaimJustid) }

func (c XclaimId) Lastid() XclaimLastid { _ = "STUB: not implemented"; return *new(XclaimLastid) }

func (c XclaimId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XclaimIdle Incomplete

func (c XclaimIdle) Time(msUnixTime int64) XclaimTime {
	_ = "STUB: not implemented"
	return *new(XclaimTime)
}

func (c XclaimIdle) Retrycount(count int64) XclaimRetrycount {
	_ = "STUB: not implemented"
	return *new(XclaimRetrycount)
}

func (c XclaimIdle) Force() XclaimForce { _ = "STUB: not implemented"; return *new(XclaimForce) }

func (c XclaimIdle) Justid() XclaimJustid { _ = "STUB: not implemented"; return *new(XclaimJustid) }

func (c XclaimIdle) Lastid() XclaimLastid { _ = "STUB: not implemented"; return *new(XclaimLastid) }

func (c XclaimIdle) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XclaimJustid Incomplete

func (c XclaimJustid) Lastid() XclaimLastid { _ = "STUB: not implemented"; return *new(XclaimLastid) }

func (c XclaimJustid) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XclaimKey Incomplete

func (c XclaimKey) Group(group string) XclaimGroup {
	_ = "STUB: not implemented"
	return *new(XclaimGroup)
}

type XclaimLastid Incomplete

func (c XclaimLastid) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XclaimMinIdleTime Incomplete

func (c XclaimMinIdleTime) Id(id ...string) XclaimId {
	_ = "STUB: not implemented"
	return *new(XclaimId)
}

type XclaimRetrycount Incomplete

func (c XclaimRetrycount) Force() XclaimForce { _ = "STUB: not implemented"; return *new(XclaimForce) }

func (c XclaimRetrycount) Justid() XclaimJustid {
	_ = "STUB: not implemented"
	return *new(XclaimJustid)
}

func (c XclaimRetrycount) Lastid() XclaimLastid {
	_ = "STUB: not implemented"
	return *new(XclaimLastid)
}

func (c XclaimRetrycount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XclaimTime Incomplete

func (c XclaimTime) Retrycount(count int64) XclaimRetrycount {
	_ = "STUB: not implemented"
	return *new(XclaimRetrycount)
}

func (c XclaimTime) Force() XclaimForce { _ = "STUB: not implemented"; return *new(XclaimForce) }

func (c XclaimTime) Justid() XclaimJustid { _ = "STUB: not implemented"; return *new(XclaimJustid) }

func (c XclaimTime) Lastid() XclaimLastid { _ = "STUB: not implemented"; return *new(XclaimLastid) }

func (c XclaimTime) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Xdel Incomplete

func (b Builder) Xdel() (c Xdel) { _ = "STUB: not implemented"; return *new(Xdel) }

func (c Xdel) Key(key string) XdelKey { _ = "STUB: not implemented"; return *new(XdelKey) }

type XdelId Incomplete

func (c XdelId) Id(id ...string) XdelId { _ = "STUB: not implemented"; return *new(XdelId) }

func (c XdelId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XdelKey Incomplete

func (c XdelKey) Id(id ...string) XdelId { _ = "STUB: not implemented"; return *new(XdelId) }

type Xdelex Incomplete

func (b Builder) Xdelex() (c Xdelex) { _ = "STUB: not implemented"; return *new(Xdelex) }

func (c Xdelex) Key(key string) XdelexKey { _ = "STUB: not implemented"; return *new(XdelexKey) }

type XdelexConditionAcked Incomplete

func (c XdelexConditionAcked) Ids() XdelexIds { _ = "STUB: not implemented"; return *new(XdelexIds) }

type XdelexConditionDelref Incomplete

func (c XdelexConditionDelref) Ids() XdelexIds { _ = "STUB: not implemented"; return *new(XdelexIds) }

type XdelexConditionKeepref Incomplete

func (c XdelexConditionKeepref) Ids() XdelexIds { _ = "STUB: not implemented"; return *new(XdelexIds) }

type XdelexId Incomplete

func (c XdelexId) Id(id ...string) XdelexId { _ = "STUB: not implemented"; return *new(XdelexId) }

func (c XdelexId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XdelexIds Incomplete

func (c XdelexIds) Numids(numids int64) XdelexNumids {
	_ = "STUB: not implemented"
	return *new(XdelexNumids)
}

type XdelexKey Incomplete

func (c XdelexKey) Keepref() XdelexConditionKeepref {
	_ = "STUB: not implemented"
	return *new(XdelexConditionKeepref)
}

func (c XdelexKey) Delref() XdelexConditionDelref {
	_ = "STUB: not implemented"
	return *new(XdelexConditionDelref)
}

func (c XdelexKey) Acked() XdelexConditionAcked {
	_ = "STUB: not implemented"
	return *new(XdelexConditionAcked)
}

func (c XdelexKey) Ids() XdelexIds { _ = "STUB: not implemented"; return *new(XdelexIds) }

type XdelexNumids Incomplete

func (c XdelexNumids) Id(id ...string) XdelexId { _ = "STUB: not implemented"; return *new(XdelexId) }

type XgroupCreate Incomplete

func (b Builder) XgroupCreate() (c XgroupCreate) {
	_ = "STUB: not implemented"
	return *new(XgroupCreate)
}

func (c XgroupCreate) Key(key string) XgroupCreateKey {
	_ = "STUB: not implemented"
	return *new(XgroupCreateKey)
}

type XgroupCreateEntriesread Incomplete

func (c XgroupCreateEntriesread) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type XgroupCreateGroup Incomplete

func (c XgroupCreateGroup) Id(id string) XgroupCreateId {
	_ = "STUB: not implemented"
	return *new(XgroupCreateId)
}

type XgroupCreateId Incomplete

func (c XgroupCreateId) Mkstream() XgroupCreateMkstream {
	_ = "STUB: not implemented"
	return *new(XgroupCreateMkstream)
}

func (c XgroupCreateId) Entriesread(entriesRead int64) XgroupCreateEntriesread {
	_ = "STUB: not implemented"
	return *new(XgroupCreateEntriesread)
}

func (c XgroupCreateId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XgroupCreateKey Incomplete

func (c XgroupCreateKey) Group(group string) XgroupCreateGroup {
	_ = "STUB: not implemented"
	return *new(XgroupCreateGroup)
}

type XgroupCreateMkstream Incomplete

func (c XgroupCreateMkstream) Entriesread(entriesRead int64) XgroupCreateEntriesread {
	_ = "STUB: not implemented"
	return *new(XgroupCreateEntriesread)
}

func (c XgroupCreateMkstream) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XgroupCreateconsumer Incomplete

func (b Builder) XgroupCreateconsumer() (c XgroupCreateconsumer) {
	_ = "STUB: not implemented"
	return *new(XgroupCreateconsumer)
}

func (c XgroupCreateconsumer) Key(key string) XgroupCreateconsumerKey {
	_ = "STUB: not implemented"
	return *new(XgroupCreateconsumerKey)
}

type XgroupCreateconsumerConsumer Incomplete

func (c XgroupCreateconsumerConsumer) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type XgroupCreateconsumerGroup Incomplete

func (c XgroupCreateconsumerGroup) Consumer(consumer string) XgroupCreateconsumerConsumer {
	_ = "STUB: not implemented"
	return *new(XgroupCreateconsumerConsumer)
}

type XgroupCreateconsumerKey Incomplete

func (c XgroupCreateconsumerKey) Group(group string) XgroupCreateconsumerGroup {
	_ = "STUB: not implemented"
	return *new(XgroupCreateconsumerGroup)
}

type XgroupDelconsumer Incomplete

func (b Builder) XgroupDelconsumer() (c XgroupDelconsumer) {
	_ = "STUB: not implemented"
	return *new(XgroupDelconsumer)
}

func (c XgroupDelconsumer) Key(key string) XgroupDelconsumerKey {
	_ = "STUB: not implemented"
	return *new(XgroupDelconsumerKey)
}

type XgroupDelconsumerConsumername Incomplete

func (c XgroupDelconsumerConsumername) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type XgroupDelconsumerGroup Incomplete

func (c XgroupDelconsumerGroup) Consumername(consumername string) XgroupDelconsumerConsumername {
	_ = "STUB: not implemented"
	return *new(XgroupDelconsumerConsumername)
}

type XgroupDelconsumerKey Incomplete

func (c XgroupDelconsumerKey) Group(group string) XgroupDelconsumerGroup {
	_ = "STUB: not implemented"
	return *new(XgroupDelconsumerGroup)
}

type XgroupDestroy Incomplete

func (b Builder) XgroupDestroy() (c XgroupDestroy) {
	_ = "STUB: not implemented"
	return *new(XgroupDestroy)
}

func (c XgroupDestroy) Key(key string) XgroupDestroyKey {
	_ = "STUB: not implemented"
	return *new(XgroupDestroyKey)
}

type XgroupDestroyGroup Incomplete

func (c XgroupDestroyGroup) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XgroupDestroyKey Incomplete

func (c XgroupDestroyKey) Group(group string) XgroupDestroyGroup {
	_ = "STUB: not implemented"
	return *new(XgroupDestroyGroup)
}

type XgroupHelp Incomplete

func (b Builder) XgroupHelp() (c XgroupHelp) { _ = "STUB: not implemented"; return *new(XgroupHelp) }

func (c XgroupHelp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XgroupSetid Incomplete

func (b Builder) XgroupSetid() (c XgroupSetid) { _ = "STUB: not implemented"; return *new(XgroupSetid) }

func (c XgroupSetid) Key(key string) XgroupSetidKey {
	_ = "STUB: not implemented"
	return *new(XgroupSetidKey)
}

type XgroupSetidEntriesread Incomplete

func (c XgroupSetidEntriesread) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type XgroupSetidGroup Incomplete

func (c XgroupSetidGroup) Id(id string) XgroupSetidId {
	_ = "STUB: not implemented"
	return *new(XgroupSetidId)
}

type XgroupSetidId Incomplete

func (c XgroupSetidId) Entriesread(entriesRead int64) XgroupSetidEntriesread {
	_ = "STUB: not implemented"
	return *new(XgroupSetidEntriesread)
}

func (c XgroupSetidId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XgroupSetidKey Incomplete

func (c XgroupSetidKey) Group(group string) XgroupSetidGroup {
	_ = "STUB: not implemented"
	return *new(XgroupSetidGroup)
}

type XinfoConsumers Incomplete

func (b Builder) XinfoConsumers() (c XinfoConsumers) {
	_ = "STUB: not implemented"
	return *new(XinfoConsumers)
}

func (c XinfoConsumers) Key(key string) XinfoConsumersKey {
	_ = "STUB: not implemented"
	return *new(XinfoConsumersKey)
}

type XinfoConsumersGroup Incomplete

func (c XinfoConsumersGroup) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XinfoConsumersKey Incomplete

func (c XinfoConsumersKey) Group(group string) XinfoConsumersGroup {
	_ = "STUB: not implemented"
	return *new(XinfoConsumersGroup)
}

type XinfoGroups Incomplete

func (b Builder) XinfoGroups() (c XinfoGroups) { _ = "STUB: not implemented"; return *new(XinfoGroups) }

func (c XinfoGroups) Key(key string) XinfoGroupsKey {
	_ = "STUB: not implemented"
	return *new(XinfoGroupsKey)
}

type XinfoGroupsKey Incomplete

func (c XinfoGroupsKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XinfoHelp Incomplete

func (b Builder) XinfoHelp() (c XinfoHelp) { _ = "STUB: not implemented"; return *new(XinfoHelp) }

func (c XinfoHelp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XinfoStream Incomplete

func (b Builder) XinfoStream() (c XinfoStream) { _ = "STUB: not implemented"; return *new(XinfoStream) }

func (c XinfoStream) Key(key string) XinfoStreamKey {
	_ = "STUB: not implemented"
	return *new(XinfoStreamKey)
}

type XinfoStreamFullCount Incomplete

func (c XinfoStreamFullCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XinfoStreamFullFull Incomplete

func (c XinfoStreamFullFull) Count(count int64) XinfoStreamFullCount {
	_ = "STUB: not implemented"
	return *new(XinfoStreamFullCount)
}

func (c XinfoStreamFullFull) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XinfoStreamKey Incomplete

func (c XinfoStreamKey) Full() XinfoStreamFullFull {
	_ = "STUB: not implemented"
	return *new(XinfoStreamFullFull)
}

func (c XinfoStreamKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Xlen Incomplete

func (b Builder) Xlen() (c Xlen) { _ = "STUB: not implemented"; return *new(Xlen) }

func (c Xlen) Key(key string) XlenKey { _ = "STUB: not implemented"; return *new(XlenKey) }

type XlenKey Incomplete

func (c XlenKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Xnack Incomplete

func (b Builder) Xnack() (c Xnack) { _ = "STUB: not implemented"; return *new(Xnack) }

func (c Xnack) Key(key string) XnackKey { _ = "STUB: not implemented"; return *new(XnackKey) }

type XnackCount Incomplete

func (c XnackCount) Force() XnackForce { _ = "STUB: not implemented"; return *new(XnackForce) }

func (c XnackCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XnackForce Incomplete

func (c XnackForce) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XnackGroup Incomplete

func (c XnackGroup) Silent() XnackModeSilent {
	_ = "STUB: not implemented"
	return *new(XnackModeSilent)
}

func (c XnackGroup) Fail() XnackModeFail { _ = "STUB: not implemented"; return *new(XnackModeFail) }

func (c XnackGroup) Fatal() XnackModeFatal { _ = "STUB: not implemented"; return *new(XnackModeFatal) }

type XnackIdsId Incomplete

func (c XnackIdsId) Id(id ...string) XnackIdsId { _ = "STUB: not implemented"; return *new(XnackIdsId) }

func (c XnackIdsId) Count(count int64) XnackCount {
	_ = "STUB: not implemented"
	return *new(XnackCount)
}

func (c XnackIdsId) Force() XnackForce { _ = "STUB: not implemented"; return *new(XnackForce) }

func (c XnackIdsId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XnackIdsIds Incomplete

func (c XnackIdsIds) Numids(numids int64) XnackIdsNumids {
	_ = "STUB: not implemented"
	return *new(XnackIdsNumids)
}

type XnackIdsNumids Incomplete

func (c XnackIdsNumids) Id(id ...string) XnackIdsId {
	_ = "STUB: not implemented"
	return *new(XnackIdsId)
}

type XnackKey Incomplete

func (c XnackKey) Group(group string) XnackGroup {
	_ = "STUB: not implemented"
	return *new(XnackGroup)
}

type XnackModeFail Incomplete

func (c XnackModeFail) Ids() XnackIdsIds { _ = "STUB: not implemented"; return *new(XnackIdsIds) }

type XnackModeFatal Incomplete

func (c XnackModeFatal) Ids() XnackIdsIds { _ = "STUB: not implemented"; return *new(XnackIdsIds) }

type XnackModeSilent Incomplete

func (c XnackModeSilent) Ids() XnackIdsIds { _ = "STUB: not implemented"; return *new(XnackIdsIds) }

type Xpending Incomplete

func (b Builder) Xpending() (c Xpending) { _ = "STUB: not implemented"; return *new(Xpending) }

func (c Xpending) Key(key string) XpendingKey { _ = "STUB: not implemented"; return *new(XpendingKey) }

type XpendingFiltersConsumer Incomplete

func (c XpendingFiltersConsumer) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type XpendingFiltersCount Incomplete

func (c XpendingFiltersCount) Consumer(consumer string) XpendingFiltersConsumer {
	_ = "STUB: not implemented"
	return *new(XpendingFiltersConsumer)
}

func (c XpendingFiltersCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XpendingFiltersEnd Incomplete

func (c XpendingFiltersEnd) Count(count int64) XpendingFiltersCount {
	_ = "STUB: not implemented"
	return *new(XpendingFiltersCount)
}

type XpendingFiltersIdle Incomplete

func (c XpendingFiltersIdle) Start(start string) XpendingFiltersStart {
	_ = "STUB: not implemented"
	return *new(XpendingFiltersStart)
}

type XpendingFiltersStart Incomplete

func (c XpendingFiltersStart) End(end string) XpendingFiltersEnd {
	_ = "STUB: not implemented"
	return *new(XpendingFiltersEnd)
}

type XpendingGroup Incomplete

func (c XpendingGroup) Idle(minIdleTime int64) XpendingFiltersIdle {
	_ = "STUB: not implemented"
	return *new(XpendingFiltersIdle)
}

func (c XpendingGroup) Start(start string) XpendingFiltersStart {
	_ = "STUB: not implemented"
	return *new(XpendingFiltersStart)
}

func (c XpendingGroup) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XpendingKey Incomplete

func (c XpendingKey) Group(group string) XpendingGroup {
	_ = "STUB: not implemented"
	return *new(XpendingGroup)
}

type Xrange Incomplete

func (b Builder) Xrange() (c Xrange) { _ = "STUB: not implemented"; return *new(Xrange) }

func (c Xrange) Key(key string) XrangeKey { _ = "STUB: not implemented"; return *new(XrangeKey) }

type XrangeCount Incomplete

func (c XrangeCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XrangeEnd Incomplete

func (c XrangeEnd) Count(count int64) XrangeCount {
	_ = "STUB: not implemented"
	return *new(XrangeCount)
}

func (c XrangeEnd) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XrangeKey Incomplete

func (c XrangeKey) Start(start string) XrangeStart {
	_ = "STUB: not implemented"
	return *new(XrangeStart)
}

type XrangeStart Incomplete

func (c XrangeStart) End(end string) XrangeEnd { _ = "STUB: not implemented"; return *new(XrangeEnd) }

type Xread Incomplete

func (b Builder) Xread() (c Xread) { _ = "STUB: not implemented"; return *new(Xread) }

func (c Xread) Count(count int64) XreadCount { _ = "STUB: not implemented"; return *new(XreadCount) }

func (c Xread) Block(milliseconds int64) XreadBlock {
	_ = "STUB: not implemented"
	return *new(XreadBlock)
}

func (c Xread) Streams() XreadStreams { _ = "STUB: not implemented"; return *new(XreadStreams) }

type XreadBlock Incomplete

func (c XreadBlock) Streams() XreadStreams { _ = "STUB: not implemented"; return *new(XreadStreams) }

type XreadCount Incomplete

func (c XreadCount) Block(milliseconds int64) XreadBlock {
	_ = "STUB: not implemented"
	return *new(XreadBlock)
}

func (c XreadCount) Streams() XreadStreams { _ = "STUB: not implemented"; return *new(XreadStreams) }

type XreadId Incomplete

func (c XreadId) Id(id ...string) XreadId { _ = "STUB: not implemented"; return *new(XreadId) }

func (c XreadId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XreadKey Incomplete

func (c XreadKey) Key(key ...string) XreadKey { _ = "STUB: not implemented"; return *new(XreadKey) }

func (c XreadKey) Id(id ...string) XreadId { _ = "STUB: not implemented"; return *new(XreadId) }

type XreadStreams Incomplete

func (c XreadStreams) Key(key ...string) XreadKey { _ = "STUB: not implemented"; return *new(XreadKey) }

type Xreadgroup Incomplete

func (b Builder) Xreadgroup() (c Xreadgroup) { _ = "STUB: not implemented"; return *new(Xreadgroup) }

func (c Xreadgroup) Group(group string, consumer string) XreadgroupGroup {
	_ = "STUB: not implemented"
	return *new(XreadgroupGroup)
}

type XreadgroupBlock Incomplete

func (c XreadgroupBlock) Noack() XreadgroupNoack {
	_ = "STUB: not implemented"
	return *new(XreadgroupNoack)
}

func (c XreadgroupBlock) Claim(claim string) XreadgroupClaim {
	_ = "STUB: not implemented"
	return *new(XreadgroupClaim)
}

func (c XreadgroupBlock) Streams() XreadgroupStreams {
	_ = "STUB: not implemented"
	return *new(XreadgroupStreams)
}

type XreadgroupClaim Incomplete

func (c XreadgroupClaim) Streams() XreadgroupStreams {
	_ = "STUB: not implemented"
	return *new(XreadgroupStreams)
}

type XreadgroupCount Incomplete

func (c XreadgroupCount) Block(milliseconds int64) XreadgroupBlock {
	_ = "STUB: not implemented"
	return *new(XreadgroupBlock)
}

func (c XreadgroupCount) Noack() XreadgroupNoack {
	_ = "STUB: not implemented"
	return *new(XreadgroupNoack)
}

func (c XreadgroupCount) Claim(claim string) XreadgroupClaim {
	_ = "STUB: not implemented"
	return *new(XreadgroupClaim)
}

func (c XreadgroupCount) Streams() XreadgroupStreams {
	_ = "STUB: not implemented"
	return *new(XreadgroupStreams)
}

type XreadgroupGroup Incomplete

func (c XreadgroupGroup) Count(count int64) XreadgroupCount {
	_ = "STUB: not implemented"
	return *new(XreadgroupCount)
}

func (c XreadgroupGroup) Block(milliseconds int64) XreadgroupBlock {
	_ = "STUB: not implemented"
	return *new(XreadgroupBlock)
}

func (c XreadgroupGroup) Noack() XreadgroupNoack {
	_ = "STUB: not implemented"
	return *new(XreadgroupNoack)
}

func (c XreadgroupGroup) Claim(claim string) XreadgroupClaim {
	_ = "STUB: not implemented"
	return *new(XreadgroupClaim)
}

func (c XreadgroupGroup) Streams() XreadgroupStreams {
	_ = "STUB: not implemented"
	return *new(XreadgroupStreams)
}

type XreadgroupId Incomplete

func (c XreadgroupId) Id(id ...string) XreadgroupId {
	_ = "STUB: not implemented"
	return *new(XreadgroupId)
}

func (c XreadgroupId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XreadgroupKey Incomplete

func (c XreadgroupKey) Key(key ...string) XreadgroupKey {
	_ = "STUB: not implemented"
	return *new(XreadgroupKey)
}

func (c XreadgroupKey) Id(id ...string) XreadgroupId {
	_ = "STUB: not implemented"
	return *new(XreadgroupId)
}

type XreadgroupNoack Incomplete

func (c XreadgroupNoack) Claim(claim string) XreadgroupClaim {
	_ = "STUB: not implemented"
	return *new(XreadgroupClaim)
}

func (c XreadgroupNoack) Streams() XreadgroupStreams {
	_ = "STUB: not implemented"
	return *new(XreadgroupStreams)
}

type XreadgroupStreams Incomplete

func (c XreadgroupStreams) Key(key ...string) XreadgroupKey {
	_ = "STUB: not implemented"
	return *new(XreadgroupKey)
}

type Xrevrange Incomplete

func (b Builder) Xrevrange() (c Xrevrange) { _ = "STUB: not implemented"; return *new(Xrevrange) }

func (c Xrevrange) Key(key string) XrevrangeKey {
	_ = "STUB: not implemented"
	return *new(XrevrangeKey)
}

type XrevrangeCount Incomplete

func (c XrevrangeCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XrevrangeEnd Incomplete

func (c XrevrangeEnd) Start(start string) XrevrangeStart {
	_ = "STUB: not implemented"
	return *new(XrevrangeStart)
}

type XrevrangeKey Incomplete

func (c XrevrangeKey) End(end string) XrevrangeEnd {
	_ = "STUB: not implemented"
	return *new(XrevrangeEnd)
}

type XrevrangeStart Incomplete

func (c XrevrangeStart) Count(count int64) XrevrangeCount {
	_ = "STUB: not implemented"
	return *new(XrevrangeCount)
}

func (c XrevrangeStart) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Xsetid Incomplete

func (b Builder) Xsetid() (c Xsetid) { _ = "STUB: not implemented"; return *new(Xsetid) }

func (c Xsetid) Key(key string) XsetidKey { _ = "STUB: not implemented"; return *new(XsetidKey) }

type XsetidEntriesadded Incomplete

func (c XsetidEntriesadded) Maxdeletedid(maxDeletedEntryId string) XsetidMaxdeletedid {
	_ = "STUB: not implemented"
	return *new(XsetidMaxdeletedid)
}

func (c XsetidEntriesadded) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XsetidKey Incomplete

func (c XsetidKey) LastId(lastId string) XsetidLastId {
	_ = "STUB: not implemented"
	return *new(XsetidLastId)
}

type XsetidLastId Incomplete

func (c XsetidLastId) Entriesadded(entriesAdded int64) XsetidEntriesadded {
	_ = "STUB: not implemented"
	return *new(XsetidEntriesadded)
}

func (c XsetidLastId) Maxdeletedid(maxDeletedEntryId string) XsetidMaxdeletedid {
	_ = "STUB: not implemented"
	return *new(XsetidMaxdeletedid)
}

func (c XsetidLastId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XsetidMaxdeletedid Incomplete

func (c XsetidMaxdeletedid) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Xtrim Incomplete

func (b Builder) Xtrim() (c Xtrim) { _ = "STUB: not implemented"; return *new(Xtrim) }

func (c Xtrim) Key(key string) XtrimKey { _ = "STUB: not implemented"; return *new(XtrimKey) }

type XtrimKey Incomplete

func (c XtrimKey) Maxlen() XtrimTrimStrategyMaxlen {
	_ = "STUB: not implemented"
	return *new(XtrimTrimStrategyMaxlen)
}

func (c XtrimKey) Minid() XtrimTrimStrategyMinid {
	_ = "STUB: not implemented"
	return *new(XtrimTrimStrategyMinid)
}

type XtrimTrimLimit Incomplete

func (c XtrimTrimLimit) Keepref() XtrimTrimReferenceKeepref {
	_ = "STUB: not implemented"
	return *new(XtrimTrimReferenceKeepref)
}

func (c XtrimTrimLimit) Delref() XtrimTrimReferenceDelref {
	_ = "STUB: not implemented"
	return *new(XtrimTrimReferenceDelref)
}

func (c XtrimTrimLimit) Acked() XtrimTrimReferenceAcked {
	_ = "STUB: not implemented"
	return *new(XtrimTrimReferenceAcked)
}

func (c XtrimTrimLimit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type XtrimTrimOperatorAlmost Incomplete

func (c XtrimTrimOperatorAlmost) Threshold(threshold string) XtrimTrimThreshold {
	_ = "STUB: not implemented"
	return *new(XtrimTrimThreshold)
}

type XtrimTrimOperatorExact Incomplete

func (c XtrimTrimOperatorExact) Threshold(threshold string) XtrimTrimThreshold {
	_ = "STUB: not implemented"
	return *new(XtrimTrimThreshold)
}

type XtrimTrimReferenceAcked Incomplete

func (c XtrimTrimReferenceAcked) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type XtrimTrimReferenceDelref Incomplete

func (c XtrimTrimReferenceDelref) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type XtrimTrimReferenceKeepref Incomplete

func (c XtrimTrimReferenceKeepref) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type XtrimTrimStrategyMaxlen Incomplete

func (c XtrimTrimStrategyMaxlen) Exact() XtrimTrimOperatorExact {
	_ = "STUB: not implemented"
	return *new(XtrimTrimOperatorExact)
}

func (c XtrimTrimStrategyMaxlen) Almost() XtrimTrimOperatorAlmost {
	_ = "STUB: not implemented"
	return *new(XtrimTrimOperatorAlmost)
}

func (c XtrimTrimStrategyMaxlen) Threshold(threshold string) XtrimTrimThreshold {
	_ = "STUB: not implemented"
	return *new(XtrimTrimThreshold)
}

type XtrimTrimStrategyMinid Incomplete

func (c XtrimTrimStrategyMinid) Exact() XtrimTrimOperatorExact {
	_ = "STUB: not implemented"
	return *new(XtrimTrimOperatorExact)
}

func (c XtrimTrimStrategyMinid) Almost() XtrimTrimOperatorAlmost {
	_ = "STUB: not implemented"
	return *new(XtrimTrimOperatorAlmost)
}

func (c XtrimTrimStrategyMinid) Threshold(threshold string) XtrimTrimThreshold {
	_ = "STUB: not implemented"
	return *new(XtrimTrimThreshold)
}

type XtrimTrimThreshold Incomplete

func (c XtrimTrimThreshold) Limit(count int64) XtrimTrimLimit {
	_ = "STUB: not implemented"
	return *new(XtrimTrimLimit)
}

func (c XtrimTrimThreshold) Keepref() XtrimTrimReferenceKeepref {
	_ = "STUB: not implemented"
	return *new(XtrimTrimReferenceKeepref)
}

func (c XtrimTrimThreshold) Delref() XtrimTrimReferenceDelref {
	_ = "STUB: not implemented"
	return *new(XtrimTrimReferenceDelref)
}

func (c XtrimTrimThreshold) Acked() XtrimTrimReferenceAcked {
	_ = "STUB: not implemented"
	return *new(XtrimTrimReferenceAcked)
}

func (c XtrimTrimThreshold) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }
