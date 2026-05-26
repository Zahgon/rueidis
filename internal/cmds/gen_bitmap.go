// Code generated DO NOT EDIT

package cmds

type Bitcount Incomplete

func (b Builder) Bitcount() (c Bitcount) { _ = "STUB: not implemented"; return *new(Bitcount) }

func (c Bitcount) Key(key string) BitcountKey { _ = "STUB: not implemented"; return *new(BitcountKey) }

type BitcountIndexEnd Incomplete

func (c BitcountIndexEnd) Byte() BitcountIndexIndexUnitByte {
	_ = "STUB: not implemented"
	return *new(BitcountIndexIndexUnitByte)
}

func (c BitcountIndexEnd) Bit() BitcountIndexIndexUnitBit {
	_ = "STUB: not implemented"
	return *new(BitcountIndexIndexUnitBit)
}

func (c BitcountIndexEnd) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c BitcountIndexEnd) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type BitcountIndexIndexUnitBit Incomplete

func (c BitcountIndexIndexUnitBit) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c BitcountIndexIndexUnitBit) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type BitcountIndexIndexUnitByte Incomplete

func (c BitcountIndexIndexUnitByte) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c BitcountIndexIndexUnitByte) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type BitcountIndexStart Incomplete

func (c BitcountIndexStart) End(end int64) BitcountIndexEnd {
	_ = "STUB: not implemented"
	return *new(BitcountIndexEnd)
}

type BitcountKey Incomplete

func (c BitcountKey) Start(start int64) BitcountIndexStart {
	_ = "STUB: not implemented"
	return *new(BitcountIndexStart)
}

func (c BitcountKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c BitcountKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Bitfield Incomplete

func (b Builder) Bitfield() (c Bitfield) { _ = "STUB: not implemented"; return *new(Bitfield) }

func (c Bitfield) Key(key string) BitfieldKey { _ = "STUB: not implemented"; return *new(BitfieldKey) }

type BitfieldKey Incomplete

func (c BitfieldKey) Get(encoding string, offset int64) BitfieldOperationGet {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationGet)
}

func (c BitfieldKey) OverflowWrap() BitfieldOperationWriteOverflowWrap {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowWrap)
}

func (c BitfieldKey) OverflowSat() BitfieldOperationWriteOverflowSat {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowSat)
}

func (c BitfieldKey) OverflowFail() BitfieldOperationWriteOverflowFail {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowFail)
}

func (c BitfieldKey) Set(encoding string, offset int64, value int64) BitfieldOperationWriteSetSet {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetSet)
}

func (c BitfieldKey) Incrby(encoding string, offset int64, increment int64) BitfieldOperationWriteSetIncrby {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetIncrby)
}

func (c BitfieldKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BitfieldOperationGet Incomplete

func (c BitfieldOperationGet) OverflowWrap() BitfieldOperationWriteOverflowWrap {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowWrap)
}

func (c BitfieldOperationGet) OverflowSat() BitfieldOperationWriteOverflowSat {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowSat)
}

func (c BitfieldOperationGet) OverflowFail() BitfieldOperationWriteOverflowFail {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowFail)
}

func (c BitfieldOperationGet) Set(encoding string, offset int64, value int64) BitfieldOperationWriteSetSet {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetSet)
}

func (c BitfieldOperationGet) Incrby(encoding string, offset int64, increment int64) BitfieldOperationWriteSetIncrby {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetIncrby)
}

func (c BitfieldOperationGet) Get(encoding string, offset int64) BitfieldOperationGet {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationGet)
}

func (c BitfieldOperationGet) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BitfieldOperationWriteOverflowFail Incomplete

func (c BitfieldOperationWriteOverflowFail) Set(encoding string, offset int64, value int64) BitfieldOperationWriteSetSet {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetSet)
}

func (c BitfieldOperationWriteOverflowFail) Incrby(encoding string, offset int64, increment int64) BitfieldOperationWriteSetIncrby {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetIncrby)
}

type BitfieldOperationWriteOverflowSat Incomplete

func (c BitfieldOperationWriteOverflowSat) Set(encoding string, offset int64, value int64) BitfieldOperationWriteSetSet {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetSet)
}

func (c BitfieldOperationWriteOverflowSat) Incrby(encoding string, offset int64, increment int64) BitfieldOperationWriteSetIncrby {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetIncrby)
}

type BitfieldOperationWriteOverflowWrap Incomplete

func (c BitfieldOperationWriteOverflowWrap) Set(encoding string, offset int64, value int64) BitfieldOperationWriteSetSet {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetSet)
}

func (c BitfieldOperationWriteOverflowWrap) Incrby(encoding string, offset int64, increment int64) BitfieldOperationWriteSetIncrby {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetIncrby)
}

type BitfieldOperationWriteSetIncrby Incomplete

func (c BitfieldOperationWriteSetIncrby) Get(encoding string, offset int64) BitfieldOperationGet {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationGet)
}

func (c BitfieldOperationWriteSetIncrby) OverflowWrap() BitfieldOperationWriteOverflowWrap {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowWrap)
}

func (c BitfieldOperationWriteSetIncrby) OverflowSat() BitfieldOperationWriteOverflowSat {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowSat)
}

func (c BitfieldOperationWriteSetIncrby) OverflowFail() BitfieldOperationWriteOverflowFail {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowFail)
}

func (c BitfieldOperationWriteSetIncrby) Set(encoding string, offset int64, value int64) BitfieldOperationWriteSetSet {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetSet)
}

func (c BitfieldOperationWriteSetIncrby) Incrby(encoding string, offset int64, increment int64) BitfieldOperationWriteSetIncrby {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetIncrby)
}

func (c BitfieldOperationWriteSetIncrby) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type BitfieldOperationWriteSetSet Incomplete

func (c BitfieldOperationWriteSetSet) Incrby(encoding string, offset int64, increment int64) BitfieldOperationWriteSetIncrby {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetIncrby)
}

func (c BitfieldOperationWriteSetSet) Get(encoding string, offset int64) BitfieldOperationGet {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationGet)
}

func (c BitfieldOperationWriteSetSet) OverflowWrap() BitfieldOperationWriteOverflowWrap {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowWrap)
}

func (c BitfieldOperationWriteSetSet) OverflowSat() BitfieldOperationWriteOverflowSat {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowSat)
}

func (c BitfieldOperationWriteSetSet) OverflowFail() BitfieldOperationWriteOverflowFail {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteOverflowFail)
}

func (c BitfieldOperationWriteSetSet) Set(encoding string, offset int64, value int64) BitfieldOperationWriteSetSet {
	_ = "STUB: not implemented"
	return *new(BitfieldOperationWriteSetSet)
}

func (c BitfieldOperationWriteSetSet) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type BitfieldRo Incomplete

func (b Builder) BitfieldRo() (c BitfieldRo) { _ = "STUB: not implemented"; return *new(BitfieldRo) }

func (c BitfieldRo) Key(key string) BitfieldRoKey {
	_ = "STUB: not implemented"
	return *new(BitfieldRoKey)
}

type BitfieldRoGet Incomplete

func (c BitfieldRoGet) Get(encoding string, offset int64) BitfieldRoGet {
	_ = "STUB: not implemented"
	return *new(BitfieldRoGet)
}

func (c BitfieldRoGet) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c BitfieldRoGet) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type BitfieldRoKey Incomplete

func (c BitfieldRoKey) Get() BitfieldRoGet { _ = "STUB: not implemented"; return *new(BitfieldRoGet) }

func (c BitfieldRoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c BitfieldRoKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Bitop Incomplete

func (b Builder) Bitop() (c Bitop) { _ = "STUB: not implemented"; return *new(Bitop) }

func (c Bitop) And() BitopOperationAnd { _ = "STUB: not implemented"; return *new(BitopOperationAnd) }

func (c Bitop) Or() BitopOperationOr { _ = "STUB: not implemented"; return *new(BitopOperationOr) }

func (c Bitop) Xor() BitopOperationXor { _ = "STUB: not implemented"; return *new(BitopOperationXor) }

func (c Bitop) Not() BitopOperationNot { _ = "STUB: not implemented"; return *new(BitopOperationNot) }

func (c Bitop) Diff() BitopOperationDiff {
	_ = "STUB: not implemented"
	return *new(BitopOperationDiff)
}

func (c Bitop) Diff1() BitopOperationDiff1 {
	_ = "STUB: not implemented"
	return *new(BitopOperationDiff1)
}

func (c Bitop) Andor() BitopOperationAndor {
	_ = "STUB: not implemented"
	return *new(BitopOperationAndor)
}

func (c Bitop) One() BitopOperationOne { _ = "STUB: not implemented"; return *new(BitopOperationOne) }

type BitopDestkey Incomplete

func (c BitopDestkey) Key(key ...string) BitopKey { _ = "STUB: not implemented"; return *new(BitopKey) }

type BitopKey Incomplete

func (c BitopKey) Key(key ...string) BitopKey { _ = "STUB: not implemented"; return *new(BitopKey) }

func (c BitopKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BitopOperationAnd Incomplete

func (c BitopOperationAnd) Destkey(destkey string) BitopDestkey {
	_ = "STUB: not implemented"
	return *new(BitopDestkey)
}

type BitopOperationAndor Incomplete

func (c BitopOperationAndor) Destkey(destkey string) BitopDestkey {
	_ = "STUB: not implemented"
	return *new(BitopDestkey)
}

type BitopOperationDiff Incomplete

func (c BitopOperationDiff) Destkey(destkey string) BitopDestkey {
	_ = "STUB: not implemented"
	return *new(BitopDestkey)
}

type BitopOperationDiff1 Incomplete

func (c BitopOperationDiff1) Destkey(destkey string) BitopDestkey {
	_ = "STUB: not implemented"
	return *new(BitopDestkey)
}

type BitopOperationNot Incomplete

func (c BitopOperationNot) Destkey(destkey string) BitopDestkey {
	_ = "STUB: not implemented"
	return *new(BitopDestkey)
}

type BitopOperationOne Incomplete

func (c BitopOperationOne) Destkey(destkey string) BitopDestkey {
	_ = "STUB: not implemented"
	return *new(BitopDestkey)
}

type BitopOperationOr Incomplete

func (c BitopOperationOr) Destkey(destkey string) BitopDestkey {
	_ = "STUB: not implemented"
	return *new(BitopDestkey)
}

type BitopOperationXor Incomplete

func (c BitopOperationXor) Destkey(destkey string) BitopDestkey {
	_ = "STUB: not implemented"
	return *new(BitopDestkey)
}

type Bitpos Incomplete

func (b Builder) Bitpos() (c Bitpos) { _ = "STUB: not implemented"; return *new(Bitpos) }

func (c Bitpos) Key(key string) BitposKey { _ = "STUB: not implemented"; return *new(BitposKey) }

type BitposBit Incomplete

func (c BitposBit) Start(start int64) BitposIndexStart {
	_ = "STUB: not implemented"
	return *new(BitposIndexStart)
}

func (c BitposBit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c BitposBit) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type BitposIndexEndIndexEnd Incomplete

func (c BitposIndexEndIndexEnd) Byte() BitposIndexEndIndexIndexUnitByte {
	_ = "STUB: not implemented"
	return *new(BitposIndexEndIndexIndexUnitByte)
}

func (c BitposIndexEndIndexEnd) Bit() BitposIndexEndIndexIndexUnitBit {
	_ = "STUB: not implemented"
	return *new(BitposIndexEndIndexIndexUnitBit)
}

func (c BitposIndexEndIndexEnd) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c BitposIndexEndIndexEnd) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type BitposIndexEndIndexIndexUnitBit Incomplete

func (c BitposIndexEndIndexIndexUnitBit) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c BitposIndexEndIndexIndexUnitBit) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type BitposIndexEndIndexIndexUnitByte Incomplete

func (c BitposIndexEndIndexIndexUnitByte) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c BitposIndexEndIndexIndexUnitByte) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type BitposIndexStart Incomplete

func (c BitposIndexStart) End(end int64) BitposIndexEndIndexEnd {
	_ = "STUB: not implemented"
	return *new(BitposIndexEndIndexEnd)
}

func (c BitposIndexStart) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c BitposIndexStart) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type BitposKey Incomplete

func (c BitposKey) Bit(bit int64) BitposBit { _ = "STUB: not implemented"; return *new(BitposBit) }

type Getbit Incomplete

func (b Builder) Getbit() (c Getbit) { _ = "STUB: not implemented"; return *new(Getbit) }

func (c Getbit) Key(key string) GetbitKey { _ = "STUB: not implemented"; return *new(GetbitKey) }

type GetbitKey Incomplete

func (c GetbitKey) Offset(offset int64) GetbitOffset {
	_ = "STUB: not implemented"
	return *new(GetbitOffset)
}

type GetbitOffset Incomplete

func (c GetbitOffset) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GetbitOffset) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Setbit Incomplete

func (b Builder) Setbit() (c Setbit) { _ = "STUB: not implemented"; return *new(Setbit) }

func (c Setbit) Key(key string) SetbitKey { _ = "STUB: not implemented"; return *new(SetbitKey) }

type SetbitKey Incomplete

func (c SetbitKey) Offset(offset int64) SetbitOffset {
	_ = "STUB: not implemented"
	return *new(SetbitOffset)
}

type SetbitOffset Incomplete

func (c SetbitOffset) Value(value int64) SetbitValue {
	_ = "STUB: not implemented"
	return *new(SetbitValue)
}

type SetbitValue Incomplete

func (c SetbitValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }
