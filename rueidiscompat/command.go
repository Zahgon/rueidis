// Copyright (c) 2013 The github.com/go-redis/redis Authors.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
// * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package rueidiscompat

import (
	"time"

	"github.com/redis/rueidis"
)

type Cmder interface {
	SetErr(error)
	Err() error
	from(result rueidis.RedisResult)
}

type baseCmd[T any] struct {
	err        error
	val        T
	rawVal     any
	isCacheHit bool
}

func (cmd *baseCmd[T]) setIsCacheHit(val bool) { _ = "STUB: not implemented"; return }

func (cmd *baseCmd[T]) IsCacheHit() bool { _ = "STUB: not implemented"; return false }

func (cmd *baseCmd[T]) SetVal(val T) { _ = "STUB: not implemented"; return }

func (cmd *baseCmd[T]) Val() T { _ = "STUB: not implemented"; return *new(T) }

func (cmd *baseCmd[T]) SetRawVal(rawVal any) { _ = "STUB: not implemented"; return }

func (cmd *baseCmd[T]) RawVal() any { _ = "STUB: not implemented"; return *new(any) }

func (cmd *baseCmd[T]) SetErr(err error) { _ = "STUB: not implemented"; return }

func (cmd *baseCmd[T]) Err() error { _ = "STUB: not implemented"; return nil }

func (cmd *baseCmd[T]) Result() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func (cmd *baseCmd[T]) RawResult() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

type Cmd struct {
	baseCmd[any]
}

func (cmd *Cmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newCmd(res rueidis.RedisResult) *Cmd { _ = "STUB: not implemented"; return nil }

func (cmd *Cmd) Text() (string, error) { _ = "STUB: not implemented"; return "", nil }

func toString(val any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (cmd *Cmd) Int() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *Cmd) Int64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func toInt64(val any) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *Cmd) Uint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func toUint64(val any) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *Cmd) Float32() (float32, error) { _ = "STUB: not implemented"; return 0, nil }

func toFloat32(val any) (float32, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *Cmd) Float64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func toFloat64(val any) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *Cmd) Bool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func toBool(val any) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (cmd *Cmd) Slice() ([]any, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) StringSlice() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) Int64Slice() ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) Uint64Slice() ([]uint64, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) Float32Slice() ([]float32, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) Float64Slice() ([]float64, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *Cmd) BoolSlice() ([]bool, error) { _ = "STUB: not implemented"; return nil, nil }

type StringCmd struct {
	baseCmd[string]
}

func (cmd *StringCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newStringCmd(res rueidis.RedisResult) *StringCmd { _ = "STUB: not implemented"; return nil }

func (cmd *StringCmd) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *StringCmd) Bool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (cmd *StringCmd) Int() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *StringCmd) Int64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *StringCmd) Uint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *StringCmd) Float32() (float32, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *StringCmd) Float64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cmd *StringCmd) Time() (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (cmd *StringCmd) String() string { _ = "STUB: not implemented"; return "" }

type BoolCmd struct {
	baseCmd[bool]
}

func (cmd *BoolCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newBoolCmd(res rueidis.RedisResult) *BoolCmd { _ = "STUB: not implemented"; return nil }

type IntCmd struct {
	baseCmd[int64]
}

func (cmd *IntCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newIntCmd(res rueidis.RedisResult) *IntCmd { _ = "STUB: not implemented"; return nil }

func (cmd *IntCmd) Uint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

type DurationCmd struct {
	baseCmd[time.Duration]
	precision time.Duration
}

func (cmd *DurationCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newDurationCmd(res rueidis.RedisResult, precision time.Duration) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

type StatusCmd = StringCmd

func newStatusCmd(res rueidis.RedisResult) *StatusCmd { _ = "STUB: not implemented"; return nil }

type SliceCmd struct {
	baseCmd[[]any]
	keys []string
	json bool
}

func (cmd *SliceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

// for JSON.OBJKEYS

// convert to any which underlying type is []any

// keep the old behavior the same as before (don't handle error while parsing v as string)

// newSliceCmd returns SliceCmd according to input arguments, if the caller is JSONObjKeys,
// set isJSONObjKeys to true.
func newSliceCmd(res rueidis.RedisResult, isJSONObjKeys bool, keys ...string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// Scan scans the results from the map into a destination struct. The map keys
// are matched in the Redis struct fields by the `redis:"field"` tag.
// NOTE: result from JSON.ObjKeys should not call this.
func (cmd *SliceCmd) Scan(dst any) error { _ = "STUB: not implemented"; return nil }

type StringSliceCmd struct {
	baseCmd[[]string]
}

func (cmd *StringSliceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newStringSliceCmd(res rueidis.RedisResult) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type IntSliceCmd struct {
	err        error
	val        []int64
	isCacheHit bool
}

func (cmd *IntSliceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newIntSliceCmd(res rueidis.RedisResult) *IntSliceCmd { _ = "STUB: not implemented"; return nil }

func (cmd *IntSliceCmd) setIsCacheHit(isCacheHit bool) { _ = "STUB: not implemented"; return }

func (cmd *IntSliceCmd) SetVal(val []int64) { _ = "STUB: not implemented"; return }

func (cmd *IntSliceCmd) SetErr(err error) { _ = "STUB: not implemented"; return }

func (cmd *IntSliceCmd) Val() []int64 { _ = "STUB: not implemented"; return nil }

func (cmd *IntSliceCmd) Err() error { _ = "STUB: not implemented"; return nil }

func (cmd *IntSliceCmd) Result() ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *IntSliceCmd) IsCacheHit() bool { _ = "STUB: not implemented"; return false }

type BoolSliceCmd struct {
	baseCmd[[]bool]
}

func (cmd *BoolSliceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newBoolSliceCmd(res rueidis.RedisResult) *BoolSliceCmd { _ = "STUB: not implemented"; return nil }

type FloatSliceCmd struct {
	baseCmd[[]float64]
}

func (cmd *FloatSliceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newFloatSliceCmd(res rueidis.RedisResult) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type ZSliceCmd struct {
	baseCmd[[]Z]
	single bool
}

func (cmd *ZSliceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newZSliceCmd(res rueidis.RedisResult) *ZSliceCmd { _ = "STUB: not implemented"; return nil }

func newZSliceSingleCmd(res rueidis.RedisResult) *ZSliceCmd { _ = "STUB: not implemented"; return nil }

type FloatCmd struct {
	baseCmd[float64]
}

func (cmd *FloatCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newFloatCmd(res rueidis.RedisResult) *FloatCmd { _ = "STUB: not implemented"; return nil }

type ScanCmd struct {
	err    error
	keys   []string
	cursor uint64
}

func (cmd *ScanCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newScanCmd(res rueidis.RedisResult) *ScanCmd { _ = "STUB: not implemented"; return nil }

func (cmd *ScanCmd) SetVal(keys []string, cursor uint64) { _ = "STUB: not implemented"; return }

func (cmd *ScanCmd) Val() (keys []string, cursor uint64) { _ = "STUB: not implemented"; return nil, 0 }

func (cmd *ScanCmd) SetErr(err error) { _ = "STUB: not implemented"; return }

func (cmd *ScanCmd) Err() error { _ = "STUB: not implemented"; return nil }

func (cmd *ScanCmd) Result() (keys []string, cursor uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

type KeyValue struct {
	Key   string
	Value string
}

type KeyValueSliceCmd struct {
	baseCmd[[]KeyValue]
}

func (cmd *KeyValueSliceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newKeyValueSliceCmd(res rueidis.RedisResult) *KeyValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type KeyValuesCmd struct {
	err error
	val rueidis.KeyValues
}

func (cmd *KeyValuesCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newKeyValuesCmd(res rueidis.RedisResult) *KeyValuesCmd { _ = "STUB: not implemented"; return nil }

func (cmd *KeyValuesCmd) SetVal(key string, val []string) { _ = "STUB: not implemented"; return }

func (cmd *KeyValuesCmd) SetErr(err error) { _ = "STUB: not implemented"; return }

func (cmd *KeyValuesCmd) Val() (string, []string) { _ = "STUB: not implemented"; return "", nil }

func (cmd *KeyValuesCmd) Err() error { _ = "STUB: not implemented"; return nil }

func (cmd *KeyValuesCmd) Result() (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

type KeyFlags struct {
	Key   string
	Flags []string
}

type KeyFlagsCmd struct {
	baseCmd[[]KeyFlags]
}

func (cmd *KeyFlagsCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newKeyFlagsCmd(res rueidis.RedisResult) *KeyFlagsCmd { _ = "STUB: not implemented"; return nil }

type ZSliceWithKeyCmd struct {
	err error
	key string
	val []Z
}

func (cmd *ZSliceWithKeyCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newZSliceWithKeyCmd(res rueidis.RedisResult) *ZSliceWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ZSliceWithKeyCmd) SetVal(key string, val []Z) { _ = "STUB: not implemented"; return }

func (cmd *ZSliceWithKeyCmd) SetErr(err error) { _ = "STUB: not implemented"; return }

func (cmd *ZSliceWithKeyCmd) Val() (string, []Z) { _ = "STUB: not implemented"; return "", nil }

func (cmd *ZSliceWithKeyCmd) Err() error { _ = "STUB: not implemented"; return nil }

func (cmd *ZSliceWithKeyCmd) Result() (string, []Z, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

type StringStringMapCmd struct {
	baseCmd[map[string]string]
}

func (cmd *StringStringMapCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newStringStringMapCmd(res rueidis.RedisResult) *StringStringMapCmd {
	_ = "STUB: not implemented"
	return nil
}

// Scan scans the results from the map into a destination struct. The map keys
// are matched in the Redis struct fields by the `redis:"field"` tag.
func (cmd *StringStringMapCmd) Scan(dest interface{}) error { _ = "STUB: not implemented"; return nil }

type StringIntMapCmd struct {
	baseCmd[map[string]int64]
}

func (cmd *StringIntMapCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newStringIntMapCmd(res rueidis.RedisResult) *StringIntMapCmd {
	_ = "STUB: not implemented"
	return nil
}

type StringStructMapCmd struct {
	baseCmd[map[string]struct{}]
}

func (cmd *StringStructMapCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newStringStructMapCmd(res rueidis.RedisResult) *StringStructMapCmd {
	_ = "STUB: not implemented"
	return nil
}

type XMessageSliceCmd struct {
	baseCmd[[]XMessage]
}

func (cmd *XMessageSliceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newXMessageSliceCmd(res rueidis.RedisResult) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func newXMessage(r rueidis.XRangeEntry) XMessage { _ = "STUB: not implemented"; return *new(XMessage) }

type XStream struct {
	Stream   string
	Messages []XMessage
}

type XStreamSliceCmd struct {
	baseCmd[[]XStream]
}

func (cmd *XStreamSliceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newXStreamSliceCmd(res rueidis.RedisResult) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type XPending struct {
	Consumers map[string]int64
	Lower     string
	Higher    string
	Count     int64
}

type XPendingCmd struct {
	baseCmd[XPending]
}

func (cmd *XPendingCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newXPendingCmd(res rueidis.RedisResult) *XPendingCmd { _ = "STUB: not implemented"; return nil }

type XPendingExt struct {
	ID         string
	Consumer   string
	Idle       time.Duration
	RetryCount int64
}

type XPendingExtCmd struct {
	baseCmd[[]XPendingExt]
}

func (cmd *XPendingExtCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newXPendingExtCmd(res rueidis.RedisResult) *XPendingExtCmd {
	_ = "STUB: not implemented"
	return nil
}

type XAutoClaimCmd struct {
	err   error
	start string
	val   []XMessage
}

func (cmd *XAutoClaimCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newXAutoClaimCmd(res rueidis.RedisResult) *XAutoClaimCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XAutoClaimCmd) SetVal(val []XMessage, start string) { _ = "STUB: not implemented"; return }

func (cmd *XAutoClaimCmd) SetErr(err error) { _ = "STUB: not implemented"; return }

func (cmd *XAutoClaimCmd) Val() (messages []XMessage, start string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func (cmd *XAutoClaimCmd) Err() error { _ = "STUB: not implemented"; return nil }

func (cmd *XAutoClaimCmd) Result() (messages []XMessage, start string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

type XAutoClaimJustIDCmd struct {
	err   error
	start string
	val   []string
}

func (cmd *XAutoClaimJustIDCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newXAutoClaimJustIDCmd(res rueidis.RedisResult) *XAutoClaimJustIDCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *XAutoClaimJustIDCmd) SetVal(val []string, start string) {
	_ = "STUB: not implemented"
	return
}

func (cmd *XAutoClaimJustIDCmd) SetErr(err error) { _ = "STUB: not implemented"; return }

func (cmd *XAutoClaimJustIDCmd) Val() (ids []string, start string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func (cmd *XAutoClaimJustIDCmd) Err() error { _ = "STUB: not implemented"; return nil }

func (cmd *XAutoClaimJustIDCmd) Result() (ids []string, start string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

type XInfoGroup struct {
	Name            string
	LastDeliveredID string
	Consumers       int64
	Pending         int64
	EntriesRead     int64
	Lag             int64
}

type XInfoGroupsCmd struct {
	baseCmd[[]XInfoGroup]
}

func (cmd *XInfoGroupsCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newXInfoGroupsCmd(res rueidis.RedisResult) *XInfoGroupsCmd {
	_ = "STUB: not implemented"
	return nil
}

type XInfoStream struct {
	FirstEntry           XMessage
	LastEntry            XMessage
	LastGeneratedID      string
	MaxDeletedEntryID    string
	RecordedFirstEntryID string
	Length               int64
	RadixTreeKeys        int64
	RadixTreeNodes       int64
	Groups               int64
	EntriesAdded         int64
	IDMPDuration         int64
	IDMPMaxSize          int64
	PIDsTracked          int64
	IIDsTracked          int64
	IIDsAdded            int64
	IIDsDuplicates       int64
}
type XInfoStreamCmd struct {
	baseCmd[XInfoStream]
}

func (cmd *XInfoStreamCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newXInfoStreamCmd(res rueidis.RedisResult) *XInfoStreamCmd {
	_ = "STUB: not implemented"
	return nil
}

type XInfoStreamConsumerPending struct {
	DeliveryTime  time.Time
	ID            string
	DeliveryCount int64
}

type XInfoStreamGroupPending struct {
	DeliveryTime  time.Time
	ID            string
	Consumer      string
	DeliveryCount int64
}

type XInfoStreamConsumer struct {
	SeenTime time.Time
	Name     string
	Pending  []XInfoStreamConsumerPending
	PelCount int64
}

type XInfoStreamGroup struct {
	Name            string
	LastDeliveredID string
	Pending         []XInfoStreamGroupPending
	Consumers       []XInfoStreamConsumer
	EntriesRead     int64
	Lag             int64
	PelCount        int64
}

type XInfoStreamFull struct {
	LastGeneratedID      string
	MaxDeletedEntryID    string
	RecordedFirstEntryID string
	Entries              []XMessage
	Groups               []XInfoStreamGroup
	Length               int64
	RadixTreeKeys        int64
	RadixTreeNodes       int64
	EntriesAdded         int64
	IDMPDuration         int64
	IDMPMaxSize          int64
	PIDsTracked          int64
	IIDsTracked          int64
	IIDsAdded            int64
	IIDsDuplicates       int64
}

type XInfoStreamFullCmd struct {
	baseCmd[XInfoStreamFull]
}

func (cmd *XInfoStreamFullCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newXInfoStreamFullCmd(res rueidis.RedisResult) *XInfoStreamFullCmd {
	_ = "STUB: not implemented"
	return nil
}

func readStreamGroups(res rueidis.RedisMessage) ([]XInfoStreamGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readXInfoStreamGroupPending(res rueidis.RedisMessage) ([]XInfoStreamGroupPending, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readXInfoStreamConsumers(res rueidis.RedisMessage) ([]XInfoStreamConsumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type XInfoConsumer struct {
	Name    string
	Pending int64
	Idle    time.Duration
}
type XInfoConsumersCmd struct {
	baseCmd[[]XInfoConsumer]
}

func (cmd *XInfoConsumersCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newXInfoConsumersCmd(res rueidis.RedisResult) *XInfoConsumersCmd {
	_ = "STUB: not implemented"
	return nil
}

// Z represents sorted set member.
type Z struct {
	Member string
	Score  float64
}

// ZWithKey represents a sorted set member including the name of the key where it was popped.
type ZWithKey struct {
	Key string
	Z
}

// ZStore is used as an arg to ZInter/ZInterStore and ZUnion/ZUnionStore.
type ZStore struct {
	Aggregate string
	Keys      []string
	Weights   []int64
}

type ZWithKeyCmd struct {
	baseCmd[ZWithKey]
}

func (cmd *ZWithKeyCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newZWithKeyCmd(res rueidis.RedisResult) *ZWithKeyCmd { _ = "STUB: not implemented"; return nil }

type RankScore struct {
	Rank  int64
	Score float64
}

type RankWithScoreCmd struct {
	baseCmd[RankScore]
}

func (cmd *RankWithScoreCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newRankWithScoreCmd(res rueidis.RedisResult) *RankWithScoreCmd {
	_ = "STUB: not implemented"
	return nil
}

type TimeCmd struct {
	baseCmd[time.Time]
}

func (cmd *TimeCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newTimeCmd(res rueidis.RedisResult) *TimeCmd { _ = "STUB: not implemented"; return nil }

type ClusterNode struct {
	ID   string
	Addr string
}

type ClusterSlot struct {
	Nodes []ClusterNode
	Start int64
	End   int64
}

type ClusterSlotsCmd struct {
	baseCmd[[]ClusterSlot]
}

func (cmd *ClusterSlotsCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newClusterSlotsCmd(res rueidis.RedisResult) *ClusterSlotsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ClusterShardsCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newClusterShardsCmd(res rueidis.RedisResult) *ClusterShardsCmd {
	_ = "STUB: not implemented"
	return nil
}

type SlotRange struct {
	Start int64
	End   int64
}
type Node struct {
	ID                string
	Endpoint          string
	IP                string
	Hostname          string
	Role              string
	Health            string
	Port              int64
	TLSPort           int64
	ReplicationOffset int64
}
type ClusterShard struct {
	Slots []SlotRange
	Nodes []Node
}

type ClusterShardsCmd struct {
	baseCmd[[]ClusterShard]
}

type GeoPos struct {
	Longitude, Latitude float64
}

type GeoPosCmd struct {
	baseCmd[[]*GeoPos]
}

func (cmd *GeoPosCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newGeoPosCmd(res rueidis.RedisResult) *GeoPosCmd { _ = "STUB: not implemented"; return nil }

type GeoLocationCmd struct {
	baseCmd[[]rueidis.GeoLocation]
}

func (cmd *GeoLocationCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newGeoLocationCmd(res rueidis.RedisResult) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

type CommandInfo struct {
	Name        string
	Flags       []string
	ACLFlags    []string
	Arity       int64
	FirstKeyPos int64
	LastKeyPos  int64
	StepCount   int64
	ReadOnly    bool
}

type CommandsInfoCmd struct {
	baseCmd[map[string]CommandInfo]
}

func (cmd *CommandsInfoCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newCommandsInfoCmd(res rueidis.RedisResult) *CommandsInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

type HExpireArgs struct {
	NX bool
	XX bool
	GT bool
	LT bool
}

// ExpirationType represents an expiration option for the HGETEX command.
type HGetEXExpirationType string

const (
	HGetEXExpirationEX      HGetEXExpirationType = "EX"
	HGetEXExpirationPX      HGetEXExpirationType = "PX"
	HGetEXExpirationEXAT    HGetEXExpirationType = "EXAT"
	HGetEXExpirationPXAT    HGetEXExpirationType = "PXAT"
	HGetEXExpirationPERSIST HGetEXExpirationType = "PERSIST"
)

type HSetEXCondition string

const (
	HSetEXFNX HSetEXCondition = "FNX"
	HSetEXFXX HSetEXCondition = "FXX"
)

type HGetEXOptions struct {
	ExpirationType HGetEXExpirationType
	ExpirationVal  int64
}

type HSetEXExpirationType string

const (
	HSetEXExpirationEX      HSetEXExpirationType = "EX"
	HSetEXExpirationPX      HSetEXExpirationType = "PX"
	HSetEXExpirationEXAT    HSetEXExpirationType = "EXAT"
	HSetEXExpirationPXAT    HSetEXExpirationType = "PXAT"
	HSetEXExpirationKEEPTTL HSetEXExpirationType = "KEEPTTL"
)

type HSetEXOptions struct {
	Condition      HSetEXCondition
	ExpirationType HSetEXExpirationType
	ExpirationVal  int64
}

type Sort struct {
	By     string
	Order  string
	Get    []string
	Offset int64
	Count  int64
	Alpha  bool
}

// SetArgs provides arguments for the SetArgs function.
type SetArgs struct {
	ExpireAt time.Time
	Mode     string
	TTL      time.Duration
	Get      bool
	KeepTTL  bool
}

type BitCount struct {
	Unit       string // Stores BIT or BYTE
	Start, End int64
}

//type BitPos struct {
//	BitCount
//	Byte bool
//}

type BitFieldArg struct {
	Encoding string
	Offset   int64
}

type BitField struct {
	Get       *BitFieldArg
	Set       *BitFieldArg
	IncrBy    *BitFieldArg
	Overflow  string
	Increment int64
}

type LPosArgs struct {
	Rank, MaxLen int64
}

// Note: MaxLen/MaxLenApprox and MinID are in conflict, only one of them can be used.
type XAddArgs struct {
	Values         any
	Stream         string
	MinID          string
	ID             string
	MaxLen         int64
	Limit          int64
	NoMkStream     bool
	Approx         bool
	ProducerID     string
	IdempotentID   string
	IdempotentAuto bool
}

type XCfgSetArgs struct {
	Stream   string
	Duration int64
	MaxSize  int64
}

type XReadArgs struct {
	Streams []string // list of streams
	Count   int64
	Block   time.Duration
}

type XReadGroupArgs struct {
	Group    string
	Consumer string
	Streams  []string // list of streams
	Count    int64
	Block    time.Duration
	NoAck    bool
}

type XPendingExtArgs struct {
	Stream   string
	Group    string
	Start    string
	End      string
	Consumer string
	Idle     time.Duration
	Count    int64
}

type XClaimArgs struct {
	Stream   string
	Group    string
	Consumer string
	Messages []string
	MinIdle  time.Duration
}

type XAutoClaimArgs struct {
	Stream   string
	Group    string
	Start    string
	Consumer string
	MinIdle  time.Duration
	Count    int64
}

type XMessage struct {
	Values map[string]any
	ID     string
}

// Note: The GT, LT and NX options are mutually exclusive.
type ZAddArgs struct {
	Members []Z
	NX      bool
	XX      bool
	LT      bool
	GT      bool
	Ch      bool
}

// ZRangeArgs is all the options of the ZRange command.
// In version> 6.2.0, you can replace the(cmd):
//
//	ZREVRANGE,
//	ZRANGEBYSCORE,
//	ZREVRANGEBYSCORE,
//	ZRANGEBYLEX,
//	ZREVRANGEBYLEX.
//
// Please pay attention to your redis-server version.
//
// Rev, ByScore, ByLex and Offset+Count options require redis-server 6.2.0 and higher.
type ZRangeArgs struct {
	Start   any
	Stop    any
	Key     string
	Offset  int64
	Count   int64
	ByScore bool
	ByLex   bool
	Rev     bool
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

type GeoLocation = rueidis.GeoLocation

// GeoRadiusQuery is used with GeoRadius to query geospatial index.
type GeoRadiusQuery struct {
	Unit        string
	Sort        string
	Store       string
	StoreDist   string
	Radius      float64
	Count       int64
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
}

// GeoSearchQuery is used for GEOSearch/GEOSearchStore command query.
type GeoSearchQuery struct {
	Member     string
	RadiusUnit string
	BoxUnit    string
	Sort       string
	Longitude  float64
	Latitude   float64
	Radius     float64
	BoxWidth   float64
	BoxHeight  float64
	Count      int64
	CountAny   bool
}

type GeoSearchLocationQuery struct {
	GeoSearchQuery

	WithCoord bool
	WithDist  bool
	WithHash  bool
}

type GeoSearchStoreQuery struct {
	GeoSearchQuery

	// When using the StoreDist option, the command stores the items in a
	// sorted set populated with their distance from the center of the circle or box,
	// as a floating-point number, in the same unit specified for that shape.
	StoreDist bool
}

func (q *GeoRadiusQuery) args() []string { _ = "STUB: not implemented"; return nil }

func (q *GeoSearchQuery) args() []string { _ = "STUB: not implemented"; return nil }

func (q *GeoSearchLocationQuery) args() []string { _ = "STUB: not implemented"; return nil }

type Function struct {
	Name        string
	Description string
	Flags       []string
}

type Library struct {
	Name      string
	Engine    string
	Code      string
	Functions []Function
}

type FunctionListQuery struct {
	LibraryNamePattern string
	WithCode           bool
}

type FunctionListCmd struct {
	baseCmd[[]Library]
}

func (cmd *FunctionListCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newFunctionListCmd(res rueidis.RedisResult) *FunctionListCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FunctionListCmd) First() (*Library, error) { _ = "STUB: not implemented"; return nil, nil }

func usePrecise(dur time.Duration) bool { _ = "STUB: not implemented"; return false }

func formatMs(dur time.Duration) int64 { _ = "STUB: not implemented"; return 0 }

// too small, truncate too 1ms

func formatSec(dur time.Duration) int64 { _ = "STUB: not implemented"; return 0 }

// too small, truncate too 1s

// https://github.com/redis/go-redis/blob/f994ff1cd96299a5c8029ae3403af7b17ef06e8a/gears_commands.go#L21C1-L35C2
type TFunctionLoadOptions struct {
	Config  string
	Replace bool
}

type TFunctionListOptions struct {
	Library  string
	Verbose  int
	Withcode bool
}

type TFCallOptions struct {
	Keys      []string
	Arguments []string
}

type MapStringInterfaceSliceCmd struct {
	baseCmd[[]map[string]any]
}

func (cmd *MapStringInterfaceSliceCmd) from(res rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return
}

func newMapStringInterfaceSliceCmd(res rueidis.RedisResult) *MapStringInterfaceSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type BFInsertOptions struct {
	Capacity   int64
	Error      float64
	Expansion  int64
	NonScaling bool
	NoCreate   bool
}

type BFReserveOptions struct {
	Capacity   int64
	Error      float64
	Expansion  int64
	NonScaling bool
}

type CFReserveOptions struct {
	Capacity      int64
	BucketSize    int64
	MaxIterations int64
	Expansion     int64
}

type CFInsertOptions struct {
	Capacity int64
	NoCreate bool
}

type BFInfo struct {
	Capacity      int64 `redis:"Capacity"`
	Size          int64 `redis:"Size"`
	Filters       int64 `redis:"Number of filters"`
	ItemsInserted int64 `redis:"Number of items inserted"`
	ExpansionRate int64 `redis:"Expansion rate"`
}

type BFInfoCmd struct {
	baseCmd[BFInfo]
}

func (cmd *BFInfoCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newBFInfoCmd(res rueidis.RedisResult) *BFInfoCmd { _ = "STUB: not implemented"; return nil }

type ScanDump struct {
	Data string
	Iter int64
}

type ScanDumpCmd struct {
	baseCmd[ScanDump]
}

func (cmd *ScanDumpCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newScanDumpCmd(res rueidis.RedisResult) *ScanDumpCmd { _ = "STUB: not implemented"; return nil }

type CFInfo struct {
	Size             int64 `redis:"Size"`
	NumBuckets       int64 `redis:"Number of buckets"`
	NumFilters       int64 `redis:"Number of filters"`
	NumItemsInserted int64 `redis:"Number of items inserted"`
	NumItemsDeleted  int64 `redis:"Number of items deleted"`
	BucketSize       int64 `redis:"Bucket size"`
	ExpansionRate    int64 `redis:"Expansion rate"`
	MaxIteration     int64 `redis:"Max iterations"`
}

type CFInfoCmd struct {
	baseCmd[CFInfo]
}

func (cmd *CFInfoCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newCFInfoCmd(res rueidis.RedisResult) *CFInfoCmd { _ = "STUB: not implemented"; return nil }

type CMSInfo struct {
	Width int64 `redis:"width"`
	Depth int64 `redis:"depth"`
	Count int64 `redis:"count"`
}

type CMSInfoCmd struct {
	baseCmd[CMSInfo]
}

func (cmd *CMSInfoCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newCMSInfoCmd(res rueidis.RedisResult) *CMSInfoCmd { _ = "STUB: not implemented"; return nil }

type TopKInfo struct {
	K     int64   `redis:"k"`
	Width int64   `redis:"width"`
	Depth int64   `redis:"depth"`
	Decay float64 `redis:"decay"`
}

type TopKInfoCmd struct {
	baseCmd[TopKInfo]
}

func (cmd *TopKInfoCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

// args of strconv.FormatFloat is copied from cmds.TopkReserveParamsDepth.Decay

func newTopKInfoCmd(res rueidis.RedisResult) *TopKInfoCmd { _ = "STUB: not implemented"; return nil }

type MapStringIntCmd struct {
	baseCmd[map[string]int64]
}

func (cmd *MapStringIntCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newMapStringIntCmd(res rueidis.RedisResult) *MapStringIntCmd {
	_ = "STUB: not implemented"
	return nil
}

// Ref: https://redis.io/commands/tdigest.info/
type TDigestInfo struct {
	Compression       int64 `redis:"Compression"`
	Capacity          int64 `redis:"Capacity"`
	MergedNodes       int64 `redis:"Merged nodes"`
	UnmergedNodes     int64 `redis:"UnmergedNodes"`
	MergedWeight      int64 `redis:"MergedWeight"`
	UnmergedWeight    int64 `redis:"Unmerged weight"`
	Observations      int64 `redis:"Observations"`
	TotalCompressions int64 `redis:"Total compressions"`
	MemoryUsage       int64 `redis:"Memory usage"`
}

type TDigestInfoCmd struct {
	baseCmd[TDigestInfo]
}

func (cmd *TDigestInfoCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newTDigestInfoCmd(res rueidis.RedisResult) *TDigestInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

type TDigestMergeOptions struct {
	Compression int64
	Override    bool
}

type TSOptions struct {
	Labels          map[string]string
	Encoding        string
	DuplicatePolicy string
	Retention       int
	ChunkSize       int
}
type TSIncrDecrOptions struct {
	Labels       map[string]string
	Timestamp    int64
	Retention    int
	ChunkSize    int
	Uncompressed bool
}

type TSAlterOptions struct {
	Labels          map[string]string
	DuplicatePolicy string
	Retention       int
	ChunkSize       int
}

type TSCreateRuleOptions struct {
	AlignTimestamp int64
}

type TSGetOptions struct {
	Latest bool
}

type TSInfoOptions struct {
	Debug bool
}
type Aggregator int

const (
	Invalid = Aggregator(iota)
	Avg
	Sum
	Min
	Max
	Range
	Count
	First
	Last
	StdP
	StdS
	VarP
	VarS
	Twa
)

func (a Aggregator) String() string { _ = "STUB: not implemented"; return "" }

type TSTimestampValue struct {
	Timestamp int64
	Value     float64
}
type TSTimestampValueCmd struct {
	baseCmd[TSTimestampValue]
}

func (cmd *TSTimestampValueCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newTSTimestampValueCmd(res rueidis.RedisResult) *TSTimestampValueCmd {
	_ = "STUB: not implemented"
	return nil
}

type MapStringInterfaceCmd struct {
	baseCmd[map[string]any]
}

func (cmd *MapStringInterfaceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newMapStringInterfaceCmd(res rueidis.RedisResult) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

type TSTimestampValueSliceCmd struct {
	baseCmd[[]TSTimestampValue]
}

func (cmd *TSTimestampValueSliceCmd) from(res rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return
}

func newTSTimestampValueSliceCmd(res rueidis.RedisResult) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type MapStringSliceInterfaceCmd struct {
	baseCmd[map[string][]any]
}

func (cmd *MapStringSliceInterfaceCmd) from(res rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return
}

func newMapStringSliceInterfaceCmd(res rueidis.RedisResult) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

type TSRangeOptions struct {
	Align           interface{}
	BucketTimestamp interface{}
	FilterByTS      []int
	FilterByValue   []int
	Count           int
	Aggregator      Aggregator
	BucketDuration  int
	Latest          bool
	Empty           bool
}

type TSRevRangeOptions struct {
	Align           interface{}
	BucketTimestamp interface{}
	FilterByTS      []int
	FilterByValue   []int
	Count           int
	Aggregator      Aggregator
	BucketDuration  int
	Latest          bool
	Empty           bool
}

type TSMRangeOptions struct {
	Align           interface{}
	BucketTimestamp interface{}
	GroupByLabel    interface{}
	Reducer         interface{}
	FilterByTS      []int
	FilterByValue   []int
	SelectedLabels  []interface{}
	Count           int
	Aggregator      Aggregator
	BucketDuration  int
	Latest          bool
	WithLabels      bool
	Empty           bool
}

type TSMRevRangeOptions struct {
	Align           interface{}
	BucketTimestamp interface{}
	GroupByLabel    interface{}
	Reducer         interface{}
	FilterByTS      []int
	FilterByValue   []int
	SelectedLabels  []interface{}
	Count           int
	Aggregator      Aggregator
	BucketDuration  int
	Latest          bool
	WithLabels      bool
	Empty           bool
}

type TSMGetOptions struct {
	SelectedLabels []interface{}
	Latest         bool
	WithLabels     bool
}

type JSONSetArgs struct {
	Value interface{}
	Key   string
	Path  string
}

type JSONArrIndexArgs struct {
	Stop  *int
	Start int
}

type JSONArrTrimArgs struct {
	Stop  *int
	Start int
}

type JSONCmd struct {
	baseCmd[string]
	expanded []any // expanded will be used at JSONCmd.Expanded
	typ      jsonCmdTyp
}

type jsonCmdTyp int

const (
	TYP_STRING jsonCmdTyp = iota
	TYP_ARRAY
)

// https://github.com/redis/go-redis/blob/v9.3.0/json.go#L86
func (cmd *JSONCmd) Val() string {
	_ = "STUB: not implemented"

	// https://github.com/redis/go-redis/blob/v9.3.0/json.go#L105
	return ""
}

func (cmd *JSONCmd) Expanded() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// TYP_ARRAY

func (cmd *JSONCmd) Result() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (cmd *JSONCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

// JSON.GET

// JSON.NUMINCRBY

// we set the marshaled string to cmd.val
// which will be used at cmd.Val()
// and also stored parsed result to cmd.expanded,
// which will be used at cmd.Expanded()

func newJSONCmd(res rueidis.RedisResult) *JSONCmd { _ = "STUB: not implemented"; return nil }

type JSONGetArgs struct {
	Indent  string
	Newline string
	Space   string
}

type IntPointerSliceCmd struct {
	baseCmd[[]*int64]
}

func (cmd *IntPointerSliceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

// newIntPointerSliceCmd initialises an IntPointerSliceCmd
func newIntPointerSliceCmd(res rueidis.RedisResult) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

type JSONSliceCmd struct {
	baseCmd[[]any]
}

func (cmd *JSONSliceCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newJSONSliceCmd(res rueidis.RedisResult) *JSONSliceCmd { _ = "STUB: not implemented"; return nil }

type MapMapStringInterfaceCmd struct {
	baseCmd[map[string]any]
}

func (cmd *MapMapStringInterfaceCmd) from(res rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return
}

func newMapMapStringInterfaceCmd(res rueidis.RedisResult) *MapMapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

type FTAggregateResult struct {
	Rows  []AggregateRow
	Total int
}

type AggregateRow struct {
	Fields map[string]any
}

// Each AggregateReducer have different args.
// Please follow https://redis.io/docs/interact/search-and-query/search/aggregations/#supported-groupby-reducers for more information.
type FTAggregateReducer struct {
	As      string
	Args    []interface{}
	Reducer SearchAggregator
}

type FTAggregateGroupBy struct {
	Fields []interface{}
	Reduce []FTAggregateReducer
}

type FTAggregateSortBy struct {
	FieldName string
	Asc       bool
	Desc      bool
}

type FTAggregateApply struct {
	Field string
	As    string
}

type FTAggregateLoad struct {
	Field string
	As    string
}

type FTAggregateWithCursor struct {
	Count   int
	MaxIdle int
}

type FTAggregateOptions struct {
	WithCursorOptions *FTAggregateWithCursor
	Params            map[string]interface{}
	Filter            string
	Scorer            string
	Load              []FTAggregateLoad
	GroupBy           []FTAggregateGroupBy
	SortBy            []FTAggregateSortBy
	Apply             []FTAggregateApply
	Timeout           int
	SortByMax         int
	LimitOffset       int
	Limit             int
	DialectVersion    int
	Verbatim          bool
	LoadAll           bool
	WithCursor        bool
	AddScores         bool
}

type AggregateCmd struct {
	baseCmd[*FTAggregateResult]
}

func (cmd *AggregateCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

// is RESP2 array

// Ref: https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L584
func processAggregateResult(data []interface{}) (*FTAggregateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newAggregateCmd(res rueidis.RedisResult) *AggregateCmd { _ = "STUB: not implemented"; return nil }

type FTCreateOptions struct {
	Filter          string
	DefaultLanguage string
	LanguageField   string
	ScoreField      string
	PayloadField    string
	Prefix          []any
	StopWords       []any
	Score           float64
	MaxTextFields   int
	Temporary       int
	OnHash          bool
	OnJSON          bool
	NoOffsets       bool
	NoHL            bool
	NoFields        bool
	NoFreqs         bool
	SkipInitialScan bool
}

type SearchAggregator int

const (
	SearchInvalid = SearchAggregator(iota)
	SearchAvg
	SearchSum
	SearchMin
	SearchMax
	SearchCount
	SearchCountDistinct
	SearchCountDistinctish
	SearchStdDev
	SearchQuantile
	SearchToList
	SearchFirstValue
	SearchRandomSample
)

func (a SearchAggregator) String() string { _ = "STUB: not implemented"; return "" }

type SearchFieldType int

const (
	SearchFieldTypeInvalid = SearchFieldType(iota)
	SearchFieldTypeNumeric
	SearchFieldTypeTag
	SearchFieldTypeText
	SearchFieldTypeGeo
	SearchFieldTypeVector
	SearchFieldTypeGeoShape
)

func (t SearchFieldType) String() string { _ = "STUB: not implemented"; return "" }

type FieldSchema struct {
	VectorArgs        *FTVectorArgs
	FieldName         string
	As                string
	PhoneticMatcher   string
	Separator         string
	GeoShapeFieldType string
	FieldType         SearchFieldType
	Weight            float64
	Sortable          bool
	UNF               bool
	NoStem            bool
	NoIndex           bool
	CaseSensitive     bool
	WithSuffixtrie    bool
	IndexEmpty        bool
	IndexMissing      bool
}

type FTVectorArgs struct {
	FlatOptions *FTFlatOptions
	HNSWOptions *FTHNSWOptions
}

type FTFlatOptions struct {
	Type            string
	DistanceMetric  string
	Dim             int
	InitialCapacity int
	BlockSize       int
}

type FTHNSWOptions struct {
	Type                   string
	DistanceMetric         string
	Dim                    int
	InitialCapacity        int
	MaxEdgesPerNode        int
	MaxAllowedEdgesPerNode int
	EFRunTime              int
	Epsilon                float64
}

type SpellCheckTerms struct {
	Dictionary string
	Include    bool
	Exclude    bool
}

type FTSearchFilter struct {
	FieldName any
	Min       any
	Max       any
}

type FTSearchGeoFilter struct {
	FieldName string
	Unit      string
	Longitude float64
	Latitude  float64
	Radius    float64
}

type FTSearchReturn struct {
	FieldName string
	As        string
}

type FTSearchSortBy struct {
	FieldName string
	Asc       bool
	Desc      bool
}

type FTDropIndexOptions struct {
	DeleteDocs bool
}

type FTExplainOptions struct {
	Dialect string
}

type IndexErrors struct {
	LastIndexingError    string
	LastIndexingErrorKey string
	IndexingFailures     int `redis:"indexing failures"`
}

type FTAttribute struct {
	Identifier      string
	Attribute       string
	Type            string
	PhoneticMatcher string
	Weight          float64
	Sortable        bool
	NoStem          bool
	NoIndex         bool
	UNF             bool
	CaseSensitive   bool
	WithSuffixtrie  bool
}

type CursorStats struct {
	GlobalIdle    int
	GlobalTotal   int
	IndexCapacity int
	IndexTotal    int
}

type FieldStatistic struct {
	Identifier  string
	Attribute   string
	IndexErrors IndexErrors
}

type GCStats struct {
	AverageCycleTimeMs   string `redis:"average_cycle_time_ms"`
	BytesCollected       int    `redis:"bytes_collected"`
	TotalMsRun           int    `redis:"total_ms_run"`
	TotalCycles          int    `redis:"total_cycles"`
	LastRunTimeMs        int    `redis:"last_run_time_ms"`
	GCNumericTreesMissed int    `redis:"gc_numeric_trees_missed"`
	GCBlocksDenied       int    `redis:"gc_blocks_denied"`
}

type IndexDefinition struct {
	KeyType      string
	Prefixes     []string
	DefaultScore float64
}

type FTInfoResult struct {
	DialectStats             map[string]int   `redis:"dialect_stats"`
	IndexErrors              IndexErrors      `redis:"Index Errors"`
	BytesPerRecordAvg        string           `redis:"bytes_per_record_avg"`
	IndexName                string           `redis:"index_name"`
	OffsetBitsPerRecordAvg   string           `redis:"offset_bits_per_record_avg"`
	OffsetsPerTermAvg        string           `redis:"offsets_per_term_avg"`
	RecordsPerDocAvg         string           `redis:"records_per_doc_avg"`
	Attributes               []FTAttribute    `redis:"attributes"`
	FieldStatistics          []FieldStatistic `redis:"field statistics"`
	IndexOptions             []string         `redis:"index_options"`
	IndexDefinition          IndexDefinition  `redis:"index_definition"`
	GCStats                  GCStats          `redis:"gc_stats"`
	CursorStats              CursorStats      `redis:"cursor_stats"`
	Cleaning                 int              `redis:"cleaning"`
	DocTableSizeMB           float64          `redis:"doc_table_size_mb"`
	GeoshapesSzMB            float64          `redis:"geoshapes_sz_mb"`
	HashIndexingFailures     int              `redis:"hash_indexing_failures"`
	Indexing                 int              `redis:"indexing"`
	InvertedSzMB             float64          `redis:"inverted_sz_mb"`
	KeyTableSizeMB           float64          `redis:"key_table_size_mb"`
	MaxDocID                 int              `redis:"max_doc_id"`
	NumDocs                  int              `redis:"num_docs"`
	NumRecords               int              `redis:"num_records"`
	NumTerms                 int              `redis:"num_terms"`
	NumberOfUses             int              `redis:"number_of_uses"`
	OffsetVectorsSzMB        float64          `redis:"offset_vectors_sz_mb"`
	PercentIndexed           float64          `redis:"percent_indexed"`
	SortableValuesSizeMB     float64          `redis:"sortable_values_size_mb"`
	TagOverheadSzMB          float64          `redis:"tag_overhead_sz_mb"`
	TextOverheadSzMB         float64          `redis:"text_overhead_sz_mb"`
	TotalIndexMemorySzMB     float64          `redis:"total_index_memory_sz_mb"`
	TotalIndexingTime        int              `redis:"total_indexing_time"`
	TotalInvertedIndexBlocks int              `redis:"total_inverted_index_blocks"`
	VectorIndexSzMB          float64          `redis:"vector_index_sz_mb"`
}

type FTInfoCmd struct {
	baseCmd[FTInfoResult]
}

// Ref: https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1143
func parseFTInfo(data map[string]interface{}) (FTInfoResult, error) {
	_ = "STUB: not implemented"
	return *

	// Manually parse each field from the map
	new(FTInfoResult), nil
}

func (cmd *FTInfoCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newFTInfoCmd(res rueidis.RedisResult) *FTInfoCmd { _ = "STUB: not implemented"; return nil }

type FTSpellCheckOptions struct {
	Terms    *FTSpellCheckTerms
	Distance int
	Dialect  int
}

type FTSpellCheckTerms struct {
	Inclusion  string // Either "INCLUDE" or "EXCLUDE"
	Dictionary string
	Terms      []interface{}
}

type SpellCheckResult struct {
	Term        string
	Suggestions []SpellCheckSuggestion
}

type SpellCheckSuggestion struct {
	Suggestion string
	Score      float64
}

type FTSpellCheckCmd struct{ baseCmd[[]SpellCheckResult] }

func (cmd *FTSpellCheckCmd) Val() []SpellCheckResult { _ = "STUB: not implemented"; return nil }

func (cmd *FTSpellCheckCmd) Result() ([]SpellCheckResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FTSpellCheckCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

// is RESP3 map

// map key: suggestion, score

// is RESP2 array

func newFTSpellCheckCmd(res rueidis.RedisResult) *FTSpellCheckCmd {
	_ = "STUB: not implemented"
	return nil
}

func parseFTSpellCheck(data []interface{}) ([]SpellCheckResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Document struct {
	Score   *float64
	Payload *string
	SortKey *string
	Fields  map[string]string
	ID      string
}

type FTSearchResult struct {
	Docs  []Document
	Total int64
}

type FTSearchOptions struct {
	Params          map[string]interface{}
	Language        string
	Expander        string
	Scorer          string
	Payload         string
	Filters         []FTSearchFilter
	GeoFilter       []FTSearchGeoFilter
	InKeys          []interface{}
	InFields        []interface{}
	Return          []FTSearchReturn
	SortBy          []FTSearchSortBy
	Slop            int
	Timeout         int
	LimitOffset     int
	Limit           int
	DialectVersion  int
	NoContent       bool
	Verbatim        bool
	NoStopWords     bool
	WithScores      bool
	WithPayloads    bool
	WithSortKeys    bool
	InOrder         bool
	ExplainScore    bool
	SortByWithCount bool
}

type FTSearchCmd struct {
	baseCmd[FTSearchResult]
	options *FTSearchOptions
}

// Ref: https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1541
func (cmd *FTSearchCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

// is RESP3 map

// doc.ID = resultArr[i+1].String()

// is RESP2 array

func newFTSearchCmd(res rueidis.RedisResult, options *FTSearchOptions) *FTSearchCmd {
	_ = "STUB: not implemented"
	return nil
}

type FTSynUpdateOptions struct {
	SkipInitialScan bool
}

type FTSynDumpCmd struct {
	baseCmd[[]FTSynDumpResult]
}

func (cmd *FTSynDumpCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

// is RESP3 map

// is RESP2 array

func newFTSynDumpCmd(res rueidis.RedisResult) *FTSynDumpCmd { _ = "STUB: not implemented"; return nil }

type FTSynDumpResult struct {
	Term     string
	Synonyms []string
}

// ClientFlags is redis-server client flags, copy from redis/src/server.h (redis 7.0)
type ClientFlags uint64

const (
	ClientSlave            ClientFlags = 1 << 0  /* This client is a replica */
	ClientMaster           ClientFlags = 1 << 1  /* This client is a master */
	ClientMonitor          ClientFlags = 1 << 2  /* This client is a slave monitor, see MONITOR */
	ClientMulti            ClientFlags = 1 << 3  /* This client is in a MULTI context */
	ClientBlocked          ClientFlags = 1 << 4  /* The client is waiting in a blocking operation */
	ClientDirtyCAS         ClientFlags = 1 << 5  /* Watched keys modified. EXEC will fail. */
	ClientCloseAfterReply  ClientFlags = 1 << 6  /* Close after writing the entire reply. */
	ClientUnBlocked        ClientFlags = 1 << 7  /* This client was unblocked and is stored in server.unblocked_clients */
	ClientScript           ClientFlags = 1 << 8  /* This is a non-connected client used by Lua */
	ClientAsking           ClientFlags = 1 << 9  /* Client issued the ASKING command */
	ClientCloseASAP        ClientFlags = 1 << 10 /* Close this client ASAP */
	ClientUnixSocket       ClientFlags = 1 << 11 /* Client connected via Unix domain socket */
	ClientDirtyExec        ClientFlags = 1 << 12 /* EXEC will fail for errors while queueing */
	ClientMasterForceReply ClientFlags = 1 << 13 /* Queue replies even if is master */
	ClientForceAOF         ClientFlags = 1 << 14 /* Force AOF propagation of current cmd. */
	ClientForceRepl        ClientFlags = 1 << 15 /* Force replication of the current cmd. */
	ClientPrePSync         ClientFlags = 1 << 16 /* Instance don't understand PSYNC. */
	ClientReadOnly         ClientFlags = 1 << 17 /* Cluster client is in the read-only state. */
	ClientPubSub           ClientFlags = 1 << 18 /* Client is in Pub/Sub mode. */
	ClientPreventAOFProp   ClientFlags = 1 << 19 /* Don't propagate to AOF. */
	ClientPreventReplProp  ClientFlags = 1 << 20 /* Don't propagate to slaves. */
	ClientPreventProp      ClientFlags = ClientPreventAOFProp | ClientPreventReplProp
	ClientPendingWrite     ClientFlags = 1 << 21 /* Client has output to send, but a-write handler is yet not installed. */
	ClientReplyOff         ClientFlags = 1 << 22 /* Don't send replies to a client. */
	ClientReplySkipNext    ClientFlags = 1 << 23 /* Set ClientREPLY_SKIP for the next cmd */
	ClientReplySkip        ClientFlags = 1 << 24 /* Don't send just this reply. */
	ClientLuaDebug         ClientFlags = 1 << 25 /* Run EVAL in debug mode. */
	ClientLuaDebugSync     ClientFlags = 1 << 26 /* EVAL debugging without fork() */
	ClientModule           ClientFlags = 1 << 27 /* Non-connected client used by some module. */
	ClientProtected        ClientFlags = 1 << 28 /* Client should not be freed for now. */
	ClientExecutingCommand ClientFlags = 1 << 29 /* Indicates that the client is currently in the process of handling
	   a command. usually this will be marked only during call()
	   however, blocked clients might have this flag kept until they
	   will try to reprocess the command. */
	ClientPendingCommand      ClientFlags = 1 << 30 /* Indicates the client has a fully * parsed command ready for execution. */
	ClientTracking            ClientFlags = 1 << 31 /* Client enabled key tracking in order to perform client side caching. */
	ClientTrackingBrokenRedir ClientFlags = 1 << 32 /* Target client is invalid. */
	ClientTrackingBCAST       ClientFlags = 1 << 33 /* Tracking in BCAST mode. */
	ClientTrackingOptIn       ClientFlags = 1 << 34 /* Tracking in opt-in mode. */
	ClientTrackingOptOut      ClientFlags = 1 << 35 /* Tracking in opt-out mode. */
	ClientTrackingCaching     ClientFlags = 1 << 36 /* CACHING yes/no was given, depending on opt-in/opt-out mode. */
	ClientTrackingNoLoop      ClientFlags = 1 << 37 /* Don't send invalidation messages about writes performed by myself.*/
	ClientInTimeoutTable      ClientFlags = 1 << 38 /* This client is in the timeout table. */
	ClientProtocolError       ClientFlags = 1 << 39 /* Protocol error chatting with it. */
	ClientCloseAfterCommand   ClientFlags = 1 << 40 /* Close after executing commands * and writing the entire reply. */
	ClientDenyBlocking        ClientFlags = 1 << 41 /* Indicate that the client should not be blocked. currently, turned on inside MULTI, Lua, RM_Call, and AOF client */
	ClientReplRDBOnly         ClientFlags = 1 << 42 /* This client is a replica that only wants RDB without a replication buffer. */
	ClientNoEvict             ClientFlags = 1 << 43 /* This client is protected against client memory eviction. */
	ClientAllowOOM            ClientFlags = 1 << 44 /* Client used by RM_Call is allowed to fully execute scripts even when in OOM */
	ClientNoTouch             ClientFlags = 1 << 45 /* This client will not touch LFU/LRU stats. */
	ClientPushing             ClientFlags = 1 << 46 /* This client is pushing notifications. */
)

// ClientInfo is redis-server ClientInfo
type ClientInfo struct {
	Addr               string        // address/port of the client
	LAddr              string        // address/port of a local address client connected to (bind address)
	Name               string        // the name set by the client with CLIENT SETNAME
	Events             string        // file descriptor events (see below)
	LastCmd            string        // cmd, last command played
	User               string        // the authenticated username of the client
	LibName            string        // redis version 7.2, client library name
	LibVer             string        // redis version 7.2, client library version
	ID                 int64         // redis version 2.8.12, a unique 64-bit client ID
	FD                 int64         // file descriptor corresponding to the socket
	Age                time.Duration // total duration of the connection in seconds
	Idle               time.Duration // idle time of the connection in seconds
	Flags              ClientFlags   // client flags (see below)
	DB                 int           // current database ID
	Sub                int           // number of channel subscriptions
	PSub               int           // number of pattern matching subscriptions
	SSub               int           // redis version 7.0.3, number of shard channel subscriptions
	Multi              int           // number of commands in a MULTI/EXEC context
	Watch              int           // redis version 7.4 RC1, the number of keys this client is currently watching.
	QueryBuf           int           // qbuf, query buffer length (0 means no query pending)
	QueryBufFree       int           // qbuf-free, free space of the query buffer (0 means the buffer is full)
	ArgvMem            int           // incomplete arguments for the next command (already extracted from query buffer)
	MultiMem           int           // redis version 7.0, memory is used up by buffered multi commands
	BufferSize         int           // rbs, usable size of buffer
	BufferPeak         int           // rbp, peak used size of buffer in the last 5 sec interval
	OutputBufferLength int           // obl, output buffer length
	OutputListLength   int           // oll, output list length (replies are queued in this list when the buffer is full)
	OutputMemory       int           // omem, output buffer memory usage
	TotalMemory        int           // tot-mem, total memory consumed by this client in its various buffers
	Redir              int64         // client id of current client tracking redirection
	Resp               int           // redis version 7.0, client RESP protocol version
	TotalNetIn         int64         // tot-net-in, total network bytes read from the client connection
	TotalNetOut        int64         // tot-net-out, total network bytes sent to the client connection
	TotalCmds          int64         // tot-cmds, number of commands executed by the client connection
}

type ClientInfoCmd struct {
	baseCmd[*ClientInfo]
}

func newClientInfoCmd(res rueidis.RedisResult) *ClientInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ClientInfoCmd) SetVal(val *ClientInfo) { _ = "STUB: not implemented"; return }

func (cmd *ClientInfoCmd) Val() *ClientInfo { _ = "STUB: not implemented"; return nil }

func (cmd *ClientInfoCmd) Result() (*ClientInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func stringToClientInfo(txt string) (*ClientInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fmt.Sscanf() cannot handle null values
func (cmd *ClientInfoCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

type ACLLogEntry struct {
	Count                int64
	Reason               string
	Context              string
	Object               string
	Username             string
	AgeSeconds           float64
	ClientInfo           *ClientInfo
	EntryID              int64
	TimestampCreated     int64
	TimestampLastUpdated int64
}

type ACLLogCmd struct {
	baseCmd[[]*ACLLogEntry]
}

func (cmd *ACLLogCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newACLLogCmd(res rueidis.RedisResult) *ACLLogCmd { _ = "STUB: not implemented"; return nil }

// ModuleLoadexConfig struct is used to specify the arguments for the MODULE LOADEX command of redis.
// `MODULE LOADEX path [CONFIG name value [CONFIG name value ...]] [ARGS args [args ...]]`
type ModuleLoadexConfig struct {
	Path string
	Conf map[string]interface{}
	Args []interface{}
}

type ClusterLink struct {
	Direction           string
	Node                string
	CreateTime          int64
	Events              string
	SendBufferAllocated int64
	SendBufferUsed      int64
}

// ClusterLinksCmd represents the response structure for ClusterLinks.
type ClusterLinksCmd struct {
	val []ClusterLink
	err error
}

func (c *ClusterLinksCmd) SetErr(err error) { _ = "STUB: not implemented"; return }

func (c *ClusterLinksCmd) Err() error { _ = "STUB: not implemented"; return nil }

func (cmd *ClusterLinksCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newClusterLinksCmd(resp rueidis.RedisResult) *ClusterLinksCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *ClusterLinksCmd) SetVal(val []ClusterLink) { _ = "STUB: not implemented"; return }

func (cmd *ClusterLinksCmd) Val() []ClusterLink { _ = "STUB: not implemented"; return nil }

func (cmd *ClusterLinksCmd) Result() ([]ClusterLink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type SlowLog struct {
	ID         int64
	Time       time.Time
	Duration   time.Duration
	Args       []string
	ClientAddr string
	ClientName string
}

type SlowLogCmd struct {
	baseCmd[[]*SlowLog]
}

func (cmd *SlowLogCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

func newSlowLogCmd(res rueidis.RedisResult) *SlowLogCmd { _ = "STUB: not implemented"; return nil }

// LCSQuery is a parameter used for the LCS command
type LCSQuery struct {
	Key1         string
	Key2         string
	Len          bool
	Idx          bool
	MinMatchLen  int
	WithMatchLen bool
}

// LCSMatch is the result set of the LCS command
type LCSMatch struct {
	MatchString string
	Matches     []LCSMatchedPosition
	Len         int64
}

type LCSMatchedPosition struct {
	Key1 LCSPosition
	Key2 LCSPosition

	// only for withMatchLen is true
	MatchLen int64
}

type LCSPosition struct {
	Start int64
	End   int64
}

type LCSCmd struct {
	baseCmd[*LCSMatch]

	// 1: match string
	// 2: match len
	// 3: match idx LCSMatch
	readType uint8
}

func newLCSCmd(res rueidis.RedisResult, readType uint8) *LCSCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *LCSCmd) SetVal(val *LCSMatch) { _ = "STUB: not implemented"; return }

func (cmd *LCSCmd) SetErr(err error) { _ = "STUB: not implemented"; return }

func (cmd *LCSCmd) Val() *LCSMatch { _ = "STUB: not implemented"; return nil }

func (cmd *LCSCmd) Err() error { _ = "STUB: not implemented"; return nil }

func (cmd *LCSCmd) Result() (*LCSMatch, error) { _ = "STUB: not implemented"; return nil, nil }

func (cmd *LCSCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

// match string

// match len

// read LCSMatch

// Validate length (should have exactly 2 keys: "matches" and "len")

// read matches or len field

// read an array of matched positions

// read match length

func (cmd *LCSCmd) readMatchedPositions(res rueidis.RedisMessage) ([]LCSMatchedPosition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read match length if WithMatchLen is true

func (cmd *LCSCmd) readPosition(res rueidis.RedisMessage) (LCSPosition, error) {
	_ = "STUB: not implemented"
	return *new(LCSPosition), nil
}

type FunctionStats struct {
	Engines   []Engine
	isRunning bool
	rs        RunningScript
	allrs     []RunningScript
}

func (fs *FunctionStats) Running() bool { _ = "STUB: not implemented"; return false }

func (fs *FunctionStats) RunningScript() (RunningScript, bool) {
	_ = "STUB: not implemented"
	return *new(RunningScript), false
}

func (fs *FunctionStats) AllRunningScripts() []RunningScript { _ = "STUB: not implemented"; return nil }

type FunctionStatsCmd struct {
	baseCmd[FunctionStats]
}

func newFunctionStatsCmd(res rueidis.RedisResult) *FunctionStatsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (cmd *FunctionStatsCmd) from(res rueidis.RedisResult) { _ = "STUB: not implemented"; return }

type RunningScript struct {
	Name     string
	Command  []string
	Duration time.Duration
}

func (cmd *FunctionStatsCmd) parseRunningScript(msg rueidis.RedisMessage) (RunningScript, bool, error) {
	_ = "STUB: not implemented"
	return *new(RunningScript), false, nil
}

type Engine struct {
	Language       string
	LibrariesCount int64
	FunctionsCount int64
}

func (cmd *FunctionStatsCmd) parseEngines(msg rueidis.RedisMessage) ([]Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cmd *FunctionStatsCmd) parseRunningScripts(msg rueidis.RedisMessage) ([]RunningScript, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
