// Code generated DO NOT EDIT

package cmds

type Blmove Incomplete

func (b Builder) Blmove() (c Blmove) { _ = "STUB: not implemented"; return *new(Blmove) }

func (c Blmove) Source(source string) BlmoveSource {
	_ = "STUB: not implemented"
	return *new(BlmoveSource)
}

type BlmoveDestination Incomplete

func (c BlmoveDestination) Left() BlmoveWherefromLeft {
	_ = "STUB: not implemented"
	return *new(BlmoveWherefromLeft)
}

func (c BlmoveDestination) Right() BlmoveWherefromRight {
	_ = "STUB: not implemented"
	return *new(BlmoveWherefromRight)
}

type BlmoveSource Incomplete

func (c BlmoveSource) Destination(destination string) BlmoveDestination {
	_ = "STUB: not implemented"
	return *new(BlmoveDestination)
}

type BlmoveTimeout Incomplete

func (c BlmoveTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BlmoveWherefromLeft Incomplete

func (c BlmoveWherefromLeft) Left() BlmoveWheretoLeft {
	_ = "STUB: not implemented"
	return *new(BlmoveWheretoLeft)
}

func (c BlmoveWherefromLeft) Right() BlmoveWheretoRight {
	_ = "STUB: not implemented"
	return *new(BlmoveWheretoRight)
}

type BlmoveWherefromRight Incomplete

func (c BlmoveWherefromRight) Left() BlmoveWheretoLeft {
	_ = "STUB: not implemented"
	return *new(BlmoveWheretoLeft)
}

func (c BlmoveWherefromRight) Right() BlmoveWheretoRight {
	_ = "STUB: not implemented"
	return *new(BlmoveWheretoRight)
}

type BlmoveWheretoLeft Incomplete

func (c BlmoveWheretoLeft) Timeout(timeout float64) BlmoveTimeout {
	_ = "STUB: not implemented"
	return *new(BlmoveTimeout)
}

type BlmoveWheretoRight Incomplete

func (c BlmoveWheretoRight) Timeout(timeout float64) BlmoveTimeout {
	_ = "STUB: not implemented"
	return *new(BlmoveTimeout)
}

type Blmpop Incomplete

func (b Builder) Blmpop() (c Blmpop) { _ = "STUB: not implemented"; return *new(Blmpop) }

func (c Blmpop) Timeout(timeout float64) BlmpopTimeout {
	_ = "STUB: not implemented"
	return *new(BlmpopTimeout)
}

type BlmpopCount Incomplete

func (c BlmpopCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BlmpopKey Incomplete

func (c BlmpopKey) Key(key ...string) BlmpopKey { _ = "STUB: not implemented"; return *new(BlmpopKey) }

func (c BlmpopKey) Left() BlmpopWhereLeft { _ = "STUB: not implemented"; return *new(BlmpopWhereLeft) }

func (c BlmpopKey) Right() BlmpopWhereRight {
	_ = "STUB: not implemented"
	return *new(BlmpopWhereRight)
}

type BlmpopNumkeys Incomplete

func (c BlmpopNumkeys) Key(key ...string) BlmpopKey {
	_ = "STUB: not implemented"
	return *new(BlmpopKey)
}

type BlmpopTimeout Incomplete

func (c BlmpopTimeout) Numkeys(numkeys int64) BlmpopNumkeys {
	_ = "STUB: not implemented"
	return *new(BlmpopNumkeys)
}

type BlmpopWhereLeft Incomplete

func (c BlmpopWhereLeft) Count(count int64) BlmpopCount {
	_ = "STUB: not implemented"
	return *new(BlmpopCount)
}

func (c BlmpopWhereLeft) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BlmpopWhereRight Incomplete

func (c BlmpopWhereRight) Count(count int64) BlmpopCount {
	_ = "STUB: not implemented"
	return *new(BlmpopCount)
}

func (c BlmpopWhereRight) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Blpop Incomplete

func (b Builder) Blpop() (c Blpop) { _ = "STUB: not implemented"; return *new(Blpop) }

func (c Blpop) Key(key ...string) BlpopKey { _ = "STUB: not implemented"; return *new(BlpopKey) }

type BlpopKey Incomplete

func (c BlpopKey) Key(key ...string) BlpopKey { _ = "STUB: not implemented"; return *new(BlpopKey) }

func (c BlpopKey) Timeout(timeout float64) BlpopTimeout {
	_ = "STUB: not implemented"
	return *new(BlpopTimeout)
}

type BlpopTimeout Incomplete

func (c BlpopTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Brpop Incomplete

func (b Builder) Brpop() (c Brpop) { _ = "STUB: not implemented"; return *new(Brpop) }

func (c Brpop) Key(key ...string) BrpopKey { _ = "STUB: not implemented"; return *new(BrpopKey) }

type BrpopKey Incomplete

func (c BrpopKey) Key(key ...string) BrpopKey { _ = "STUB: not implemented"; return *new(BrpopKey) }

func (c BrpopKey) Timeout(timeout float64) BrpopTimeout {
	_ = "STUB: not implemented"
	return *new(BrpopTimeout)
}

type BrpopTimeout Incomplete

func (c BrpopTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Brpoplpush Incomplete

func (b Builder) Brpoplpush() (c Brpoplpush) { _ = "STUB: not implemented"; return *new(Brpoplpush) }

func (c Brpoplpush) Source(source string) BrpoplpushSource {
	_ = "STUB: not implemented"
	return *new(BrpoplpushSource)
}

type BrpoplpushDestination Incomplete

func (c BrpoplpushDestination) Timeout(timeout float64) BrpoplpushTimeout {
	_ = "STUB: not implemented"
	return *new(BrpoplpushTimeout)
}

type BrpoplpushSource Incomplete

func (c BrpoplpushSource) Destination(destination string) BrpoplpushDestination {
	_ = "STUB: not implemented"
	return *new(BrpoplpushDestination)
}

type BrpoplpushTimeout Incomplete

func (c BrpoplpushTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Lindex Incomplete

func (b Builder) Lindex() (c Lindex) { _ = "STUB: not implemented"; return *new(Lindex) }

func (c Lindex) Key(key string) LindexKey { _ = "STUB: not implemented"; return *new(LindexKey) }

type LindexIndex Incomplete

func (c LindexIndex) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c LindexIndex) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type LindexKey Incomplete

func (c LindexKey) Index(index int64) LindexIndex {
	_ = "STUB: not implemented"
	return *new(LindexIndex)
}

type Linsert Incomplete

func (b Builder) Linsert() (c Linsert) { _ = "STUB: not implemented"; return *new(Linsert) }

func (c Linsert) Key(key string) LinsertKey { _ = "STUB: not implemented"; return *new(LinsertKey) }

type LinsertElement Incomplete

func (c LinsertElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LinsertKey Incomplete

func (c LinsertKey) Before() LinsertWhereBefore {
	_ = "STUB: not implemented"
	return *new(LinsertWhereBefore)
}

func (c LinsertKey) After() LinsertWhereAfter {
	_ = "STUB: not implemented"
	return *new(LinsertWhereAfter)
}

type LinsertPivot Incomplete

func (c LinsertPivot) Element(element string) LinsertElement {
	_ = "STUB: not implemented"
	return *new(LinsertElement)
}

type LinsertWhereAfter Incomplete

func (c LinsertWhereAfter) Pivot(pivot string) LinsertPivot {
	_ = "STUB: not implemented"
	return *new(LinsertPivot)
}

type LinsertWhereBefore Incomplete

func (c LinsertWhereBefore) Pivot(pivot string) LinsertPivot {
	_ = "STUB: not implemented"
	return *new(LinsertPivot)
}

type Llen Incomplete

func (b Builder) Llen() (c Llen) { _ = "STUB: not implemented"; return *new(Llen) }

func (c Llen) Key(key string) LlenKey { _ = "STUB: not implemented"; return *new(LlenKey) }

type LlenKey Incomplete

func (c LlenKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c LlenKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Lmove Incomplete

func (b Builder) Lmove() (c Lmove) { _ = "STUB: not implemented"; return *new(Lmove) }

func (c Lmove) Source(source string) LmoveSource {
	_ = "STUB: not implemented"
	return *new(LmoveSource)
}

type LmoveDestination Incomplete

func (c LmoveDestination) Left() LmoveWherefromLeft {
	_ = "STUB: not implemented"
	return *new(LmoveWherefromLeft)
}

func (c LmoveDestination) Right() LmoveWherefromRight {
	_ = "STUB: not implemented"
	return *new(LmoveWherefromRight)
}

type LmoveSource Incomplete

func (c LmoveSource) Destination(destination string) LmoveDestination {
	_ = "STUB: not implemented"
	return *new(LmoveDestination)
}

type LmoveWherefromLeft Incomplete

func (c LmoveWherefromLeft) Left() LmoveWheretoLeft {
	_ = "STUB: not implemented"
	return *new(LmoveWheretoLeft)
}

func (c LmoveWherefromLeft) Right() LmoveWheretoRight {
	_ = "STUB: not implemented"
	return *new(LmoveWheretoRight)
}

type LmoveWherefromRight Incomplete

func (c LmoveWherefromRight) Left() LmoveWheretoLeft {
	_ = "STUB: not implemented"
	return *new(LmoveWheretoLeft)
}

func (c LmoveWherefromRight) Right() LmoveWheretoRight {
	_ = "STUB: not implemented"
	return *new(LmoveWheretoRight)
}

type LmoveWheretoLeft Incomplete

func (c LmoveWheretoLeft) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LmoveWheretoRight Incomplete

func (c LmoveWheretoRight) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Lmpop Incomplete

func (b Builder) Lmpop() (c Lmpop) { _ = "STUB: not implemented"; return *new(Lmpop) }

func (c Lmpop) Numkeys(numkeys int64) LmpopNumkeys {
	_ = "STUB: not implemented"
	return *new(LmpopNumkeys)
}

type LmpopCount Incomplete

func (c LmpopCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LmpopKey Incomplete

func (c LmpopKey) Key(key ...string) LmpopKey { _ = "STUB: not implemented"; return *new(LmpopKey) }

func (c LmpopKey) Left() LmpopWhereLeft { _ = "STUB: not implemented"; return *new(LmpopWhereLeft) }

func (c LmpopKey) Right() LmpopWhereRight { _ = "STUB: not implemented"; return *new(LmpopWhereRight) }

type LmpopNumkeys Incomplete

func (c LmpopNumkeys) Key(key ...string) LmpopKey { _ = "STUB: not implemented"; return *new(LmpopKey) }

type LmpopWhereLeft Incomplete

func (c LmpopWhereLeft) Count(count int64) LmpopCount {
	_ = "STUB: not implemented"
	return *new(LmpopCount)
}

func (c LmpopWhereLeft) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LmpopWhereRight Incomplete

func (c LmpopWhereRight) Count(count int64) LmpopCount {
	_ = "STUB: not implemented"
	return *new(LmpopCount)
}

func (c LmpopWhereRight) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Lpop Incomplete

func (b Builder) Lpop() (c Lpop) { _ = "STUB: not implemented"; return *new(Lpop) }

func (c Lpop) Key(key string) LpopKey { _ = "STUB: not implemented"; return *new(LpopKey) }

type LpopCount Incomplete

func (c LpopCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LpopKey Incomplete

func (c LpopKey) Count(count int64) LpopCount { _ = "STUB: not implemented"; return *new(LpopCount) }

func (c LpopKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Lpos Incomplete

func (b Builder) Lpos() (c Lpos) { _ = "STUB: not implemented"; return *new(Lpos) }

func (c Lpos) Key(key string) LposKey { _ = "STUB: not implemented"; return *new(LposKey) }

type LposCount Incomplete

func (c LposCount) Maxlen(len int64) LposMaxlen { _ = "STUB: not implemented"; return *new(LposMaxlen) }

func (c LposCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c LposCount) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type LposElement Incomplete

func (c LposElement) Rank(rank int64) LposRank { _ = "STUB: not implemented"; return *new(LposRank) }

func (c LposElement) Count(numMatches int64) LposCount {
	_ = "STUB: not implemented"
	return *new(LposCount)
}

func (c LposElement) Maxlen(len int64) LposMaxlen {
	_ = "STUB: not implemented"
	return *new(LposMaxlen)
}

func (c LposElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c LposElement) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type LposKey Incomplete

func (c LposKey) Element(element string) LposElement {
	_ = "STUB: not implemented"
	return *new(LposElement)
}

type LposMaxlen Incomplete

func (c LposMaxlen) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c LposMaxlen) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type LposRank Incomplete

func (c LposRank) Count(numMatches int64) LposCount {
	_ = "STUB: not implemented"
	return *new(LposCount)
}

func (c LposRank) Maxlen(len int64) LposMaxlen { _ = "STUB: not implemented"; return *new(LposMaxlen) }

func (c LposRank) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c LposRank) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Lpush Incomplete

func (b Builder) Lpush() (c Lpush) { _ = "STUB: not implemented"; return *new(Lpush) }

func (c Lpush) Key(key string) LpushKey { _ = "STUB: not implemented"; return *new(LpushKey) }

type LpushElement Incomplete

func (c LpushElement) Element(element ...string) LpushElement {
	_ = "STUB: not implemented"
	return *new(LpushElement)
}

func (c LpushElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LpushKey Incomplete

func (c LpushKey) Element(element ...string) LpushElement {
	_ = "STUB: not implemented"
	return *new(LpushElement)
}

type Lpushx Incomplete

func (b Builder) Lpushx() (c Lpushx) { _ = "STUB: not implemented"; return *new(Lpushx) }

func (c Lpushx) Key(key string) LpushxKey { _ = "STUB: not implemented"; return *new(LpushxKey) }

type LpushxElement Incomplete

func (c LpushxElement) Element(element ...string) LpushxElement {
	_ = "STUB: not implemented"
	return *new(LpushxElement)
}

func (c LpushxElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LpushxKey Incomplete

func (c LpushxKey) Element(element ...string) LpushxElement {
	_ = "STUB: not implemented"
	return *new(LpushxElement)
}

type Lrange Incomplete

func (b Builder) Lrange() (c Lrange) { _ = "STUB: not implemented"; return *new(Lrange) }

func (c Lrange) Key(key string) LrangeKey { _ = "STUB: not implemented"; return *new(LrangeKey) }

type LrangeKey Incomplete

func (c LrangeKey) Start(start int64) LrangeStart {
	_ = "STUB: not implemented"
	return *new(LrangeStart)
}

type LrangeStart Incomplete

func (c LrangeStart) Stop(stop int64) LrangeStop {
	_ = "STUB: not implemented"
	return *new(LrangeStop)
}

type LrangeStop Incomplete

func (c LrangeStop) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c LrangeStop) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Lrem Incomplete

func (b Builder) Lrem() (c Lrem) { _ = "STUB: not implemented"; return *new(Lrem) }

func (c Lrem) Key(key string) LremKey { _ = "STUB: not implemented"; return *new(LremKey) }

type LremCount Incomplete

func (c LremCount) Element(element string) LremElement {
	_ = "STUB: not implemented"
	return *new(LremElement)
}

type LremElement Incomplete

func (c LremElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LremKey Incomplete

func (c LremKey) Count(count int64) LremCount { _ = "STUB: not implemented"; return *new(LremCount) }

type Lset Incomplete

func (b Builder) Lset() (c Lset) { _ = "STUB: not implemented"; return *new(Lset) }

func (c Lset) Key(key string) LsetKey { _ = "STUB: not implemented"; return *new(LsetKey) }

type LsetElement Incomplete

func (c LsetElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LsetIndex Incomplete

func (c LsetIndex) Element(element string) LsetElement {
	_ = "STUB: not implemented"
	return *new(LsetElement)
}

type LsetKey Incomplete

func (c LsetKey) Index(index int64) LsetIndex { _ = "STUB: not implemented"; return *new(LsetIndex) }

type Ltrim Incomplete

func (b Builder) Ltrim() (c Ltrim) { _ = "STUB: not implemented"; return *new(Ltrim) }

func (c Ltrim) Key(key string) LtrimKey { _ = "STUB: not implemented"; return *new(LtrimKey) }

type LtrimKey Incomplete

func (c LtrimKey) Start(start int64) LtrimStart { _ = "STUB: not implemented"; return *new(LtrimStart) }

type LtrimStart Incomplete

func (c LtrimStart) Stop(stop int64) LtrimStop { _ = "STUB: not implemented"; return *new(LtrimStop) }

type LtrimStop Incomplete

func (c LtrimStop) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Rpop Incomplete

func (b Builder) Rpop() (c Rpop) { _ = "STUB: not implemented"; return *new(Rpop) }

func (c Rpop) Key(key string) RpopKey { _ = "STUB: not implemented"; return *new(RpopKey) }

type RpopCount Incomplete

func (c RpopCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RpopKey Incomplete

func (c RpopKey) Count(count int64) RpopCount { _ = "STUB: not implemented"; return *new(RpopCount) }

func (c RpopKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Rpoplpush Incomplete

func (b Builder) Rpoplpush() (c Rpoplpush) { _ = "STUB: not implemented"; return *new(Rpoplpush) }

func (c Rpoplpush) Source(source string) RpoplpushSource {
	_ = "STUB: not implemented"
	return *new(RpoplpushSource)
}

type RpoplpushDestination Incomplete

func (c RpoplpushDestination) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RpoplpushSource Incomplete

func (c RpoplpushSource) Destination(destination string) RpoplpushDestination {
	_ = "STUB: not implemented"
	return *new(RpoplpushDestination)
}

type Rpush Incomplete

func (b Builder) Rpush() (c Rpush) { _ = "STUB: not implemented"; return *new(Rpush) }

func (c Rpush) Key(key string) RpushKey { _ = "STUB: not implemented"; return *new(RpushKey) }

type RpushElement Incomplete

func (c RpushElement) Element(element ...string) RpushElement {
	_ = "STUB: not implemented"
	return *new(RpushElement)
}

func (c RpushElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RpushKey Incomplete

func (c RpushKey) Element(element ...string) RpushElement {
	_ = "STUB: not implemented"
	return *new(RpushElement)
}

type Rpushx Incomplete

func (b Builder) Rpushx() (c Rpushx) { _ = "STUB: not implemented"; return *new(Rpushx) }

func (c Rpushx) Key(key string) RpushxKey { _ = "STUB: not implemented"; return *new(RpushxKey) }

type RpushxElement Incomplete

func (c RpushxElement) Element(element ...string) RpushxElement {
	_ = "STUB: not implemented"
	return *new(RpushxElement)
}

func (c RpushxElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RpushxKey Incomplete

func (c RpushxKey) Element(element ...string) RpushxElement {
	_ = "STUB: not implemented"
	return *new(RpushxElement)
}
