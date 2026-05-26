// Code generated DO NOT EDIT

package cmds

import (
	"time"
)

type Append Incomplete

func (b Builder) Append() (c Append) { _ = "STUB: not implemented"; return *new(Append) }

func (c Append) Key(key string) AppendKey { _ = "STUB: not implemented"; return *new(AppendKey) }

type AppendKey Incomplete

func (c AppendKey) Value(value string) AppendValue {
	_ = "STUB: not implemented"
	return *new(AppendValue)
}

type AppendValue Incomplete

func (c AppendValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Decr Incomplete

func (b Builder) Decr() (c Decr) { _ = "STUB: not implemented"; return *new(Decr) }

func (c Decr) Key(key string) DecrKey { _ = "STUB: not implemented"; return *new(DecrKey) }

type DecrKey Incomplete

func (c DecrKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Decrby Incomplete

func (b Builder) Decrby() (c Decrby) { _ = "STUB: not implemented"; return *new(Decrby) }

func (c Decrby) Key(key string) DecrbyKey { _ = "STUB: not implemented"; return *new(DecrbyKey) }

type DecrbyDecrement Incomplete

func (c DecrbyDecrement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type DecrbyKey Incomplete

func (c DecrbyKey) Decrement(decrement int64) DecrbyDecrement {
	_ = "STUB: not implemented"
	return *new(DecrbyDecrement)
}

type Delex Incomplete

func (b Builder) Delex() (c Delex) { _ = "STUB: not implemented"; return *new(Delex) }

func (c Delex) Key(key string) DelexKey { _ = "STUB: not implemented"; return *new(DelexKey) }

type DelexConditionIfdeq Incomplete

func (c DelexConditionIfdeq) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type DelexConditionIfdne Incomplete

func (c DelexConditionIfdne) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type DelexConditionIfeq Incomplete

func (c DelexConditionIfeq) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type DelexConditionIfne Incomplete

func (c DelexConditionIfne) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type DelexKey Incomplete

func (c DelexKey) Ifeq(ifeq string) DelexConditionIfeq {
	_ = "STUB: not implemented"
	return *new(DelexConditionIfeq)
}

func (c DelexKey) Ifne(ifne string) DelexConditionIfne {
	_ = "STUB: not implemented"
	return *new(DelexConditionIfne)
}

func (c DelexKey) Ifdeq(ifdeq string) DelexConditionIfdeq {
	_ = "STUB: not implemented"
	return *new(DelexConditionIfdeq)
}

func (c DelexKey) Ifdne(ifdne string) DelexConditionIfdne {
	_ = "STUB: not implemented"
	return *new(DelexConditionIfdne)
}

func (c DelexKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Delifeq Incomplete

func (b Builder) Delifeq() (c Delifeq) { _ = "STUB: not implemented"; return *new(Delifeq) }

func (c Delifeq) Key(key string) DelifeqKey { _ = "STUB: not implemented"; return *new(DelifeqKey) }

type DelifeqKey Incomplete

func (c DelifeqKey) Value(value string) DelifeqValue {
	_ = "STUB: not implemented"
	return *new(DelifeqValue)
}

type DelifeqValue Incomplete

func (c DelifeqValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Digest Incomplete

func (b Builder) Digest() (c Digest) { _ = "STUB: not implemented"; return *new(Digest) }

func (c Digest) Key(key string) DigestKey { _ = "STUB: not implemented"; return *new(DigestKey) }

type DigestKey Incomplete

func (c DigestKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Get Incomplete

func (b Builder) Get() (c Get) { _ = "STUB: not implemented"; return *new(Get) }

func (c Get) Key(key string) GetKey { _ = "STUB: not implemented"; return *new(GetKey) }

type GetKey Incomplete

func (c GetKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GetKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Getdel Incomplete

func (b Builder) Getdel() (c Getdel) { _ = "STUB: not implemented"; return *new(Getdel) }

func (c Getdel) Key(key string) GetdelKey { _ = "STUB: not implemented"; return *new(GetdelKey) }

type GetdelKey Incomplete

func (c GetdelKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Getex Incomplete

func (b Builder) Getex() (c Getex) { _ = "STUB: not implemented"; return *new(Getex) }

func (c Getex) Key(key string) GetexKey { _ = "STUB: not implemented"; return *new(GetexKey) }

type GetexExpirationExSecTyped Incomplete

func (c GetexExpirationExSecTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GetexExpirationExSeconds Incomplete

func (c GetexExpirationExSeconds) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GetexExpirationExatTimestamp Incomplete

func (c GetexExpirationExatTimestamp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GetexExpirationExatTimestampTyped Incomplete

func (c GetexExpirationExatTimestampTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GetexExpirationPersist Incomplete

func (c GetexExpirationPersist) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GetexExpirationPxMilliseconds Incomplete

func (c GetexExpirationPxMilliseconds) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GetexExpirationPxMsTyped Incomplete

func (c GetexExpirationPxMsTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GetexExpirationPxatMillisecondsTimestamp Incomplete

func (c GetexExpirationPxatMillisecondsTimestamp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GetexExpirationPxatMsTimestampTyped Incomplete

func (c GetexExpirationPxatMsTimestampTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GetexKey Incomplete

func (c GetexKey) ExSeconds(seconds int64) GetexExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(GetexExpirationExSeconds)
}

func (c GetexKey) PxMilliseconds(milliseconds int64) GetexExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(GetexExpirationPxMilliseconds)
}

func (c GetexKey) ExatTimestamp(timestamp int64) GetexExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(GetexExpirationExatTimestamp)
}

func (c GetexKey) PxatMillisecondsTimestamp(millisecondsTimestamp int64) GetexExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(GetexExpirationPxatMillisecondsTimestamp)
}

func (c GetexKey) Persist() GetexExpirationPersist {
	_ = "STUB: not implemented"
	return *new(GetexExpirationPersist)
}

func (c GetexKey) Ex(duration time.Duration) GetexExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(GetexExpirationExSecTyped)
}

func (c GetexKey) Px(duration time.Duration) GetexExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(GetexExpirationPxMsTyped)
}

func (c GetexKey) Exat(timestamp time.Time) GetexExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(GetexExpirationExatTimestampTyped)
}

func (c GetexKey) Pxat(timestamp time.Time) GetexExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(GetexExpirationPxatMsTimestampTyped)
}

func (c GetexKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Getrange Incomplete

func (b Builder) Getrange() (c Getrange) { _ = "STUB: not implemented"; return *new(Getrange) }

func (c Getrange) Key(key string) GetrangeKey { _ = "STUB: not implemented"; return *new(GetrangeKey) }

type GetrangeEnd Incomplete

func (c GetrangeEnd) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GetrangeEnd) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GetrangeKey Incomplete

func (c GetrangeKey) Start(start int64) GetrangeStart {
	_ = "STUB: not implemented"
	return *new(GetrangeStart)
}

type GetrangeStart Incomplete

func (c GetrangeStart) End(end int64) GetrangeEnd {
	_ = "STUB: not implemented"
	return *new(GetrangeEnd)
}

type Getset Incomplete

func (b Builder) Getset() (c Getset) { _ = "STUB: not implemented"; return *new(Getset) }

func (c Getset) Key(key string) GetsetKey { _ = "STUB: not implemented"; return *new(GetsetKey) }

type GetsetKey Incomplete

func (c GetsetKey) Value(value string) GetsetValue {
	_ = "STUB: not implemented"
	return *new(GetsetValue)
}

type GetsetValue Incomplete

func (c GetsetValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Incr Incomplete

func (b Builder) Incr() (c Incr) { _ = "STUB: not implemented"; return *new(Incr) }

func (c Incr) Key(key string) IncrKey { _ = "STUB: not implemented"; return *new(IncrKey) }

type IncrKey Incomplete

func (c IncrKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Incrby Incomplete

func (b Builder) Incrby() (c Incrby) { _ = "STUB: not implemented"; return *new(Incrby) }

func (c Incrby) Key(key string) IncrbyKey { _ = "STUB: not implemented"; return *new(IncrbyKey) }

type IncrbyIncrement Incomplete

func (c IncrbyIncrement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type IncrbyKey Incomplete

func (c IncrbyKey) Increment(increment int64) IncrbyIncrement {
	_ = "STUB: not implemented"
	return *new(IncrbyIncrement)
}

type Incrbyfloat Incomplete

func (b Builder) Incrbyfloat() (c Incrbyfloat) { _ = "STUB: not implemented"; return *new(Incrbyfloat) }

func (c Incrbyfloat) Key(key string) IncrbyfloatKey {
	_ = "STUB: not implemented"
	return *new(IncrbyfloatKey)
}

type IncrbyfloatIncrement Incomplete

func (c IncrbyfloatIncrement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type IncrbyfloatKey Incomplete

func (c IncrbyfloatKey) Increment(increment float64) IncrbyfloatIncrement {
	_ = "STUB: not implemented"
	return *new(IncrbyfloatIncrement)
}

type Lcs Incomplete

func (b Builder) Lcs() (c Lcs) { _ = "STUB: not implemented"; return *new(Lcs) }

func (c Lcs) Key1(key1 string) LcsKey1 { _ = "STUB: not implemented"; return *new(LcsKey1) }

type LcsIdx Incomplete

func (c LcsIdx) Minmatchlen(len int64) LcsMinmatchlen {
	_ = "STUB: not implemented"
	return *new(LcsMinmatchlen)
}

func (c LcsIdx) Withmatchlen() LcsWithmatchlen {
	_ = "STUB: not implemented"
	return *new(LcsWithmatchlen)
}

func (c LcsIdx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LcsKey1 Incomplete

func (c LcsKey1) Key2(key2 string) LcsKey2 { _ = "STUB: not implemented"; return *new(LcsKey2) }

type LcsKey2 Incomplete

func (c LcsKey2) Len() LcsLen { _ = "STUB: not implemented"; return *new(LcsLen) }

func (c LcsKey2) Idx() LcsIdx { _ = "STUB: not implemented"; return *new(LcsIdx) }

func (c LcsKey2) Minmatchlen(len int64) LcsMinmatchlen {
	_ = "STUB: not implemented"
	return *new(LcsMinmatchlen)
}

func (c LcsKey2) Withmatchlen() LcsWithmatchlen {
	_ = "STUB: not implemented"
	return *new(LcsWithmatchlen)
}

func (c LcsKey2) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LcsLen Incomplete

func (c LcsLen) Idx() LcsIdx { _ = "STUB: not implemented"; return *new(LcsIdx) }

func (c LcsLen) Minmatchlen(len int64) LcsMinmatchlen {
	_ = "STUB: not implemented"
	return *new(LcsMinmatchlen)
}

func (c LcsLen) Withmatchlen() LcsWithmatchlen {
	_ = "STUB: not implemented"
	return *new(LcsWithmatchlen)
}

func (c LcsLen) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LcsMinmatchlen Incomplete

func (c LcsMinmatchlen) Withmatchlen() LcsWithmatchlen {
	_ = "STUB: not implemented"
	return *new(LcsWithmatchlen)
}

func (c LcsMinmatchlen) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LcsWithmatchlen Incomplete

func (c LcsWithmatchlen) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Mget Incomplete

func (b Builder) Mget() (c Mget) { _ = "STUB: not implemented"; return *new(Mget) }

func (c Mget) Key(key ...string) MgetKey { _ = "STUB: not implemented"; return *new(MgetKey) }

type MgetKey Incomplete

func (c MgetKey) Key(key ...string) MgetKey { _ = "STUB: not implemented"; return *new(MgetKey) }

func (c MgetKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c MgetKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Mset Incomplete

func (b Builder) Mset() (c Mset) { _ = "STUB: not implemented"; return *new(Mset) }

func (c Mset) KeyValue() MsetKeyValue { _ = "STUB: not implemented"; return *new(MsetKeyValue) }

type MsetKeyValue Incomplete

func (c MsetKeyValue) KeyValue(key string, value string) MsetKeyValue {
	_ = "STUB: not implemented"
	return *new(MsetKeyValue)
}

func (c MsetKeyValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Msetex Incomplete

func (b Builder) Msetex() (c Msetex) { _ = "STUB: not implemented"; return *new(Msetex) }

func (c Msetex) Numkeys(numkeys int64) MsetexNumkeys {
	_ = "STUB: not implemented"
	return *new(MsetexNumkeys)
}

type MsetexConditionNx Incomplete

func (c MsetexConditionNx) ExSeconds(seconds int64) MsetexExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExSeconds)
}

func (c MsetexConditionNx) PxMilliseconds(milliseconds int64) MsetexExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxMilliseconds)
}

func (c MsetexConditionNx) ExatTimestamp(timestamp int64) MsetexExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExatTimestamp)
}

func (c MsetexConditionNx) PxatMillisecondsTimestamp(millisecondsTimestamp int64) MsetexExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxatMillisecondsTimestamp)
}

func (c MsetexConditionNx) Keepttl() MsetexExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationKeepttl)
}

func (c MsetexConditionNx) Ex(duration time.Duration) MsetexExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExSecTyped)
}

func (c MsetexConditionNx) Px(duration time.Duration) MsetexExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxMsTyped)
}

func (c MsetexConditionNx) Exat(timestamp time.Time) MsetexExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExatTimestampTyped)
}

func (c MsetexConditionNx) Pxat(timestamp time.Time) MsetexExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxatMsTimestampTyped)
}

func (c MsetexConditionNx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MsetexConditionXx Incomplete

func (c MsetexConditionXx) ExSeconds(seconds int64) MsetexExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExSeconds)
}

func (c MsetexConditionXx) PxMilliseconds(milliseconds int64) MsetexExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxMilliseconds)
}

func (c MsetexConditionXx) ExatTimestamp(timestamp int64) MsetexExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExatTimestamp)
}

func (c MsetexConditionXx) PxatMillisecondsTimestamp(millisecondsTimestamp int64) MsetexExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxatMillisecondsTimestamp)
}

func (c MsetexConditionXx) Keepttl() MsetexExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationKeepttl)
}

func (c MsetexConditionXx) Ex(duration time.Duration) MsetexExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExSecTyped)
}

func (c MsetexConditionXx) Px(duration time.Duration) MsetexExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxMsTyped)
}

func (c MsetexConditionXx) Exat(timestamp time.Time) MsetexExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExatTimestampTyped)
}

func (c MsetexConditionXx) Pxat(timestamp time.Time) MsetexExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxatMsTimestampTyped)
}

func (c MsetexConditionXx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MsetexExpirationExSecTyped Incomplete

func (c MsetexExpirationExSecTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type MsetexExpirationExSeconds Incomplete

func (c MsetexExpirationExSeconds) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type MsetexExpirationExatTimestamp Incomplete

func (c MsetexExpirationExatTimestamp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type MsetexExpirationExatTimestampTyped Incomplete

func (c MsetexExpirationExatTimestampTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type MsetexExpirationKeepttl Incomplete

func (c MsetexExpirationKeepttl) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type MsetexExpirationPxMilliseconds Incomplete

func (c MsetexExpirationPxMilliseconds) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type MsetexExpirationPxMsTyped Incomplete

func (c MsetexExpirationPxMsTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type MsetexExpirationPxatMillisecondsTimestamp Incomplete

func (c MsetexExpirationPxatMillisecondsTimestamp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type MsetexExpirationPxatMsTimestampTyped Incomplete

func (c MsetexExpirationPxatMsTimestampTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type MsetexKeyValue Incomplete

func (c MsetexKeyValue) KeyValue(key string, value string) MsetexKeyValue {
	_ = "STUB: not implemented"
	return *new(MsetexKeyValue)
}

func (c MsetexKeyValue) Nx() MsetexConditionNx {
	_ = "STUB: not implemented"
	return *new(MsetexConditionNx)
}

func (c MsetexKeyValue) Xx() MsetexConditionXx {
	_ = "STUB: not implemented"
	return *new(MsetexConditionXx)
}

func (c MsetexKeyValue) ExSeconds(seconds int64) MsetexExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExSeconds)
}

func (c MsetexKeyValue) PxMilliseconds(milliseconds int64) MsetexExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxMilliseconds)
}

func (c MsetexKeyValue) ExatTimestamp(timestamp int64) MsetexExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExatTimestamp)
}

func (c MsetexKeyValue) PxatMillisecondsTimestamp(millisecondsTimestamp int64) MsetexExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxatMillisecondsTimestamp)
}

func (c MsetexKeyValue) Keepttl() MsetexExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationKeepttl)
}

func (c MsetexKeyValue) Ex(duration time.Duration) MsetexExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExSecTyped)
}

func (c MsetexKeyValue) Px(duration time.Duration) MsetexExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxMsTyped)
}

func (c MsetexKeyValue) Exat(timestamp time.Time) MsetexExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationExatTimestampTyped)
}

func (c MsetexKeyValue) Pxat(timestamp time.Time) MsetexExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(MsetexExpirationPxatMsTimestampTyped)
}

func (c MsetexKeyValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MsetexNumkeys Incomplete

func (c MsetexNumkeys) KeyValue() MsetexKeyValue {
	_ = "STUB: not implemented"
	return *new(MsetexKeyValue)
}

type Msetnx Incomplete

func (b Builder) Msetnx() (c Msetnx) { _ = "STUB: not implemented"; return *new(Msetnx) }

func (c Msetnx) KeyValue() MsetnxKeyValue { _ = "STUB: not implemented"; return *new(MsetnxKeyValue) }

type MsetnxKeyValue Incomplete

func (c MsetnxKeyValue) KeyValue(key string, value string) MsetnxKeyValue {
	_ = "STUB: not implemented"
	return *new(MsetnxKeyValue)
}

func (c MsetnxKeyValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Psetex Incomplete

func (b Builder) Psetex() (c Psetex) { _ = "STUB: not implemented"; return *new(Psetex) }

func (c Psetex) Key(key string) PsetexKey { _ = "STUB: not implemented"; return *new(PsetexKey) }

type PsetexKey Incomplete

func (c PsetexKey) Milliseconds(milliseconds int64) PsetexMilliseconds {
	_ = "STUB: not implemented"
	return *new(PsetexMilliseconds)
}

type PsetexMilliseconds Incomplete

func (c PsetexMilliseconds) Value(value string) PsetexValue {
	_ = "STUB: not implemented"
	return *new(PsetexValue)
}

type PsetexValue Incomplete

func (c PsetexValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Set Incomplete

func (b Builder) Set() (c Set) { _ = "STUB: not implemented"; return *new(Set) }

func (c Set) Key(key string) SetKey { _ = "STUB: not implemented"; return *new(SetKey) }

type SetConditionIfdeq Incomplete

func (c SetConditionIfdeq) Get() SetGet { _ = "STUB: not implemented"; return *new(SetGet) }

func (c SetConditionIfdeq) ExSeconds(seconds int64) SetExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSeconds)
}

func (c SetConditionIfdeq) PxMilliseconds(milliseconds int64) SetExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMilliseconds)
}

func (c SetConditionIfdeq) ExatTimestamp(timestamp int64) SetExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestamp)
}

func (c SetConditionIfdeq) PxatMillisecondsTimestamp(millisecondsTimestamp int64) SetExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMillisecondsTimestamp)
}

func (c SetConditionIfdeq) Keepttl() SetExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(SetExpirationKeepttl)
}

func (c SetConditionIfdeq) Ex(duration time.Duration) SetExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSecTyped)
}

func (c SetConditionIfdeq) Px(duration time.Duration) SetExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMsTyped)
}

func (c SetConditionIfdeq) Exat(timestamp time.Time) SetExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestampTyped)
}

func (c SetConditionIfdeq) Pxat(timestamp time.Time) SetExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMsTimestampTyped)
}

func (c SetConditionIfdeq) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SetConditionIfdne Incomplete

func (c SetConditionIfdne) Get() SetGet { _ = "STUB: not implemented"; return *new(SetGet) }

func (c SetConditionIfdne) ExSeconds(seconds int64) SetExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSeconds)
}

func (c SetConditionIfdne) PxMilliseconds(milliseconds int64) SetExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMilliseconds)
}

func (c SetConditionIfdne) ExatTimestamp(timestamp int64) SetExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestamp)
}

func (c SetConditionIfdne) PxatMillisecondsTimestamp(millisecondsTimestamp int64) SetExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMillisecondsTimestamp)
}

func (c SetConditionIfdne) Keepttl() SetExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(SetExpirationKeepttl)
}

func (c SetConditionIfdne) Ex(duration time.Duration) SetExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSecTyped)
}

func (c SetConditionIfdne) Px(duration time.Duration) SetExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMsTyped)
}

func (c SetConditionIfdne) Exat(timestamp time.Time) SetExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestampTyped)
}

func (c SetConditionIfdne) Pxat(timestamp time.Time) SetExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMsTimestampTyped)
}

func (c SetConditionIfdne) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SetConditionIfeq Incomplete

func (c SetConditionIfeq) Get() SetGet { _ = "STUB: not implemented"; return *new(SetGet) }

func (c SetConditionIfeq) ExSeconds(seconds int64) SetExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSeconds)
}

func (c SetConditionIfeq) PxMilliseconds(milliseconds int64) SetExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMilliseconds)
}

func (c SetConditionIfeq) ExatTimestamp(timestamp int64) SetExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestamp)
}

func (c SetConditionIfeq) PxatMillisecondsTimestamp(millisecondsTimestamp int64) SetExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMillisecondsTimestamp)
}

func (c SetConditionIfeq) Keepttl() SetExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(SetExpirationKeepttl)
}

func (c SetConditionIfeq) Ex(duration time.Duration) SetExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSecTyped)
}

func (c SetConditionIfeq) Px(duration time.Duration) SetExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMsTyped)
}

func (c SetConditionIfeq) Exat(timestamp time.Time) SetExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestampTyped)
}

func (c SetConditionIfeq) Pxat(timestamp time.Time) SetExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMsTimestampTyped)
}

func (c SetConditionIfeq) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SetConditionIfne Incomplete

func (c SetConditionIfne) Get() SetGet { _ = "STUB: not implemented"; return *new(SetGet) }

func (c SetConditionIfne) ExSeconds(seconds int64) SetExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSeconds)
}

func (c SetConditionIfne) PxMilliseconds(milliseconds int64) SetExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMilliseconds)
}

func (c SetConditionIfne) ExatTimestamp(timestamp int64) SetExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestamp)
}

func (c SetConditionIfne) PxatMillisecondsTimestamp(millisecondsTimestamp int64) SetExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMillisecondsTimestamp)
}

func (c SetConditionIfne) Keepttl() SetExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(SetExpirationKeepttl)
}

func (c SetConditionIfne) Ex(duration time.Duration) SetExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSecTyped)
}

func (c SetConditionIfne) Px(duration time.Duration) SetExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMsTyped)
}

func (c SetConditionIfne) Exat(timestamp time.Time) SetExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestampTyped)
}

func (c SetConditionIfne) Pxat(timestamp time.Time) SetExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMsTimestampTyped)
}

func (c SetConditionIfne) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SetConditionNx Incomplete

func (c SetConditionNx) Get() SetGet { _ = "STUB: not implemented"; return *new(SetGet) }

func (c SetConditionNx) ExSeconds(seconds int64) SetExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSeconds)
}

func (c SetConditionNx) PxMilliseconds(milliseconds int64) SetExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMilliseconds)
}

func (c SetConditionNx) ExatTimestamp(timestamp int64) SetExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestamp)
}

func (c SetConditionNx) PxatMillisecondsTimestamp(millisecondsTimestamp int64) SetExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMillisecondsTimestamp)
}

func (c SetConditionNx) Keepttl() SetExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(SetExpirationKeepttl)
}

func (c SetConditionNx) Ex(duration time.Duration) SetExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSecTyped)
}

func (c SetConditionNx) Px(duration time.Duration) SetExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMsTyped)
}

func (c SetConditionNx) Exat(timestamp time.Time) SetExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestampTyped)
}

func (c SetConditionNx) Pxat(timestamp time.Time) SetExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMsTimestampTyped)
}

func (c SetConditionNx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SetConditionXx Incomplete

func (c SetConditionXx) Get() SetGet { _ = "STUB: not implemented"; return *new(SetGet) }

func (c SetConditionXx) ExSeconds(seconds int64) SetExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSeconds)
}

func (c SetConditionXx) PxMilliseconds(milliseconds int64) SetExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMilliseconds)
}

func (c SetConditionXx) ExatTimestamp(timestamp int64) SetExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestamp)
}

func (c SetConditionXx) PxatMillisecondsTimestamp(millisecondsTimestamp int64) SetExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMillisecondsTimestamp)
}

func (c SetConditionXx) Keepttl() SetExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(SetExpirationKeepttl)
}

func (c SetConditionXx) Ex(duration time.Duration) SetExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSecTyped)
}

func (c SetConditionXx) Px(duration time.Duration) SetExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMsTyped)
}

func (c SetConditionXx) Exat(timestamp time.Time) SetExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestampTyped)
}

func (c SetConditionXx) Pxat(timestamp time.Time) SetExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMsTimestampTyped)
}

func (c SetConditionXx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SetExpirationExSecTyped Incomplete

func (c SetExpirationExSecTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SetExpirationExSeconds Incomplete

func (c SetExpirationExSeconds) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SetExpirationExatTimestamp Incomplete

func (c SetExpirationExatTimestamp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SetExpirationExatTimestampTyped Incomplete

func (c SetExpirationExatTimestampTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SetExpirationKeepttl Incomplete

func (c SetExpirationKeepttl) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SetExpirationPxMilliseconds Incomplete

func (c SetExpirationPxMilliseconds) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SetExpirationPxMsTyped Incomplete

func (c SetExpirationPxMsTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SetExpirationPxatMillisecondsTimestamp Incomplete

func (c SetExpirationPxatMillisecondsTimestamp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SetExpirationPxatMsTimestampTyped Incomplete

func (c SetExpirationPxatMsTimestampTyped) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SetGet Incomplete

func (c SetGet) ExSeconds(seconds int64) SetExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSeconds)
}

func (c SetGet) PxMilliseconds(milliseconds int64) SetExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMilliseconds)
}

func (c SetGet) ExatTimestamp(timestamp int64) SetExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestamp)
}

func (c SetGet) PxatMillisecondsTimestamp(millisecondsTimestamp int64) SetExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMillisecondsTimestamp)
}

func (c SetGet) Keepttl() SetExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(SetExpirationKeepttl)
}

func (c SetGet) Ex(duration time.Duration) SetExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSecTyped)
}

func (c SetGet) Px(duration time.Duration) SetExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMsTyped)
}

func (c SetGet) Exat(timestamp time.Time) SetExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestampTyped)
}

func (c SetGet) Pxat(timestamp time.Time) SetExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMsTimestampTyped)
}

func (c SetGet) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SetKey Incomplete

func (c SetKey) Value(value string) SetValue { _ = "STUB: not implemented"; return *new(SetValue) }

type SetValue Incomplete

func (c SetValue) Nx() SetConditionNx { _ = "STUB: not implemented"; return *new(SetConditionNx) }

func (c SetValue) Xx() SetConditionXx { _ = "STUB: not implemented"; return *new(SetConditionXx) }

func (c SetValue) Ifeq(ifeq string) SetConditionIfeq {
	_ = "STUB: not implemented"
	return *new(SetConditionIfeq)
}

func (c SetValue) Ifne(ifne string) SetConditionIfne {
	_ = "STUB: not implemented"
	return *new(SetConditionIfne)
}

func (c SetValue) Ifdeq(ifdeq string) SetConditionIfdeq {
	_ = "STUB: not implemented"
	return *new(SetConditionIfdeq)
}

func (c SetValue) Ifdne(ifdne string) SetConditionIfdne {
	_ = "STUB: not implemented"
	return *new(SetConditionIfdne)
}

func (c SetValue) Get() SetGet { _ = "STUB: not implemented"; return *new(SetGet) }

func (c SetValue) ExSeconds(seconds int64) SetExpirationExSeconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSeconds)
}

func (c SetValue) PxMilliseconds(milliseconds int64) SetExpirationPxMilliseconds {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMilliseconds)
}

func (c SetValue) ExatTimestamp(timestamp int64) SetExpirationExatTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestamp)
}

func (c SetValue) PxatMillisecondsTimestamp(millisecondsTimestamp int64) SetExpirationPxatMillisecondsTimestamp {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMillisecondsTimestamp)
}

func (c SetValue) Keepttl() SetExpirationKeepttl {
	_ = "STUB: not implemented"
	return *new(SetExpirationKeepttl)
}

func (c SetValue) Ex(duration time.Duration) SetExpirationExSecTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExSecTyped)
}

func (c SetValue) Px(duration time.Duration) SetExpirationPxMsTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxMsTyped)
}

func (c SetValue) Exat(timestamp time.Time) SetExpirationExatTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationExatTimestampTyped)
}

func (c SetValue) Pxat(timestamp time.Time) SetExpirationPxatMsTimestampTyped {
	_ = "STUB: not implemented"
	return *new(SetExpirationPxatMsTimestampTyped)
}

func (c SetValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Setex Incomplete

func (b Builder) Setex() (c Setex) { _ = "STUB: not implemented"; return *new(Setex) }

func (c Setex) Key(key string) SetexKey { _ = "STUB: not implemented"; return *new(SetexKey) }

type SetexKey Incomplete

func (c SetexKey) Seconds(seconds int64) SetexSeconds {
	_ = "STUB: not implemented"
	return *new(SetexSeconds)
}

type SetexSeconds Incomplete

func (c SetexSeconds) Value(value string) SetexValue {
	_ = "STUB: not implemented"
	return *new(SetexValue)
}

type SetexValue Incomplete

func (c SetexValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Setnx Incomplete

func (b Builder) Setnx() (c Setnx) { _ = "STUB: not implemented"; return *new(Setnx) }

func (c Setnx) Key(key string) SetnxKey { _ = "STUB: not implemented"; return *new(SetnxKey) }

type SetnxKey Incomplete

func (c SetnxKey) Value(value string) SetnxValue {
	_ = "STUB: not implemented"
	return *new(SetnxValue)
}

type SetnxValue Incomplete

func (c SetnxValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Setrange Incomplete

func (b Builder) Setrange() (c Setrange) { _ = "STUB: not implemented"; return *new(Setrange) }

func (c Setrange) Key(key string) SetrangeKey { _ = "STUB: not implemented"; return *new(SetrangeKey) }

type SetrangeKey Incomplete

func (c SetrangeKey) Offset(offset int64) SetrangeOffset {
	_ = "STUB: not implemented"
	return *new(SetrangeOffset)
}

type SetrangeOffset Incomplete

func (c SetrangeOffset) Value(value string) SetrangeValue {
	_ = "STUB: not implemented"
	return *new(SetrangeValue)
}

type SetrangeValue Incomplete

func (c SetrangeValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Strlen Incomplete

func (b Builder) Strlen() (c Strlen) { _ = "STUB: not implemented"; return *new(Strlen) }

func (c Strlen) Key(key string) StrlenKey { _ = "STUB: not implemented"; return *new(StrlenKey) }

type StrlenKey Incomplete

func (c StrlenKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c StrlenKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }
