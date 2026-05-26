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
	"context"
	"reflect"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/internal/cmds"
)

const (
	KeepTTL           = -1
	BitCountIndexByte = "BYTE"
	BitCountIndexBit  = "BIT"
)

var Nil = rueidis.Nil

type Cmdable interface {
	CoreCmdable
	Cache(ttl time.Duration) CacheCompat

	Subscribe(ctx context.Context, channels ...string) PubSub
	PSubscribe(ctx context.Context, patterns ...string) PubSub
	SSubscribe(ctx context.Context, channels ...string) PubSub

	Watch(ctx context.Context, fn func(Tx) error, keys ...string) error

	// ForEachMaster concurrently calls the fn on each master node in the cluster.
	// It returns the first error if any.
	ForEachMaster(ctx context.Context, fn func(ctx context.Context, client Cmdable) error) error

	Client() rueidis.Client
}

type CoreCmdable interface {
	Command(ctx context.Context) *CommandsInfoCmd
	CommandList(ctx context.Context, filter FilterBy) *StringSliceCmd
	CommandGetKeys(ctx context.Context, commands ...any) *StringSliceCmd
	CommandGetKeysAndFlags(ctx context.Context, commands ...any) *KeyFlagsCmd
	ClientGetName(ctx context.Context) *StringCmd
	Echo(ctx context.Context, message any) *StringCmd
	Ping(ctx context.Context) *StatusCmd
	Quit(ctx context.Context) *StatusCmd
	Del(ctx context.Context, keys ...string) *IntCmd
	Unlink(ctx context.Context, keys ...string) *IntCmd
	Dump(ctx context.Context, key string) *StringCmd
	Exists(ctx context.Context, keys ...string) *IntCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
	ExpireTime(ctx context.Context, key string) *DurationCmd
	ExpireNX(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireXX(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireGT(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	ExpireLT(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	Keys(ctx context.Context, pattern string) *StringSliceCmd
	Migrate(ctx context.Context, host string, port int64, key string, db int64, timeout time.Duration) *StatusCmd
	Move(ctx context.Context, key string, db int64) *BoolCmd
	ObjectRefCount(ctx context.Context, key string) *IntCmd
	ObjectEncoding(ctx context.Context, key string) *StringCmd
	ObjectIdleTime(ctx context.Context, key string) *DurationCmd
	Persist(ctx context.Context, key string) *BoolCmd
	PExpire(ctx context.Context, key string, expiration time.Duration) *BoolCmd
	PExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd
	PExpireTime(ctx context.Context, key string) *DurationCmd
	PTTL(ctx context.Context, key string) *DurationCmd
	RandomKey(ctx context.Context) *StringCmd
	Rename(ctx context.Context, key, newkey string) *StatusCmd
	RenameNX(ctx context.Context, key, newkey string) *BoolCmd
	Restore(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd
	RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd
	Sort(ctx context.Context, key string, sort Sort) *StringSliceCmd
	SortRO(ctx context.Context, key string, sort Sort) *StringSliceCmd
	SortStore(ctx context.Context, key, store string, sort Sort) *IntCmd
	SortInterfaces(ctx context.Context, key string, sort Sort) *SliceCmd
	Touch(ctx context.Context, keys ...string) *IntCmd
	TTL(ctx context.Context, key string) *DurationCmd
	Type(ctx context.Context, key string) *StatusCmd
	Append(ctx context.Context, key, value string) *IntCmd
	Decr(ctx context.Context, key string) *IntCmd
	DecrBy(ctx context.Context, key string, decrement int64) *IntCmd
	Get(ctx context.Context, key string) *StringCmd
	GetRange(ctx context.Context, key string, start, end int64) *StringCmd
	GetSet(ctx context.Context, key string, value any) *StringCmd
	GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd
	GetDel(ctx context.Context, key string) *StringCmd
	Incr(ctx context.Context, key string) *IntCmd
	IncrBy(ctx context.Context, key string, value int64) *IntCmd
	IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd
	MGet(ctx context.Context, keys ...string) *SliceCmd
	MSet(ctx context.Context, values ...any) *StatusCmd
	MSetNX(ctx context.Context, values ...any) *BoolCmd
	Set(ctx context.Context, key string, value any, expiration time.Duration) *StatusCmd
	SetArgs(ctx context.Context, key string, value any, a SetArgs) *StatusCmd
	SetEX(ctx context.Context, key string, value any, expiration time.Duration) *StatusCmd
	SetNX(ctx context.Context, key string, value any, expiration time.Duration) *BoolCmd
	SetXX(ctx context.Context, key string, value any, expiration time.Duration) *BoolCmd
	SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd
	StrLen(ctx context.Context, key string) *IntCmd
	Copy(ctx context.Context, sourceKey string, destKey string, db int64, replace bool) *IntCmd

	GetBit(ctx context.Context, key string, offset int64) *IntCmd
	SetBit(ctx context.Context, key string, offset int64, value int64) *IntCmd
	BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd
	BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd
	BitOpNot(ctx context.Context, destKey string, key string) *IntCmd
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd
	BitPosSpan(ctx context.Context, key string, bit int64, start, end int64, span string) *IntCmd
	BitField(ctx context.Context, key string, args ...any) *IntSliceCmd
	BitFieldRO(ctx context.Context, key string, values ...any) *IntSliceCmd

	Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd
	ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	HScanNoValues(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd

	HDel(ctx context.Context, key string, fields ...string) *IntCmd
	HExists(ctx context.Context, key, field string) *BoolCmd
	HGet(ctx context.Context, key, field string) *StringCmd
	HGetAll(ctx context.Context, key string) *StringStringMapCmd
	HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd
	HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd
	HKeys(ctx context.Context, key string) *StringSliceCmd
	HLen(ctx context.Context, key string) *IntCmd
	HStrLen(ctx context.Context, key, field string) *IntCmd
	HMGet(ctx context.Context, key string, fields ...string) *SliceCmd
	HSet(ctx context.Context, key string, values ...any) *IntCmd
	HMSet(ctx context.Context, key string, values ...any) *BoolCmd
	HSetNX(ctx context.Context, key, field string, value any) *BoolCmd
	HVals(ctx context.Context, key string) *StringSliceCmd
	HRandField(ctx context.Context, key string, count int64) *StringSliceCmd
	HRandFieldWithValues(ctx context.Context, key string, count int64) *KeyValueSliceCmd
	HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *IntSliceCmd
	HExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd
	HPExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *IntSliceCmd
	HPExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd
	HExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *IntSliceCmd
	HExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd
	HPExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *IntSliceCmd
	HPExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd
	HPersist(ctx context.Context, key string, fields ...string) *IntSliceCmd
	HExpireTime(ctx context.Context, key string, fields ...string) *IntSliceCmd
	HPExpireTime(ctx context.Context, key string, fields ...string) *IntSliceCmd
	HTTL(ctx context.Context, key string, fields ...string) *IntSliceCmd
	HPTTL(ctx context.Context, key string, fields ...string) *IntSliceCmd
	HGetDel(ctx context.Context, key string, fields ...string) *StringSliceCmd
	HGetEX(ctx context.Context, key string, fields ...string) *StringSliceCmd
	HGetEXWithArgs(ctx context.Context, key string, options *HGetEXOptions, fields ...string) *StringSliceCmd
	HSetEX(ctx context.Context, key string, fieldsAndValues ...string) *IntCmd
	HSetEXWithArgs(ctx context.Context, key string, options *HSetEXOptions, fieldsAndValues ...string) *IntCmd

	BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
	BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *KeyValuesCmd
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd
	LCS(ctx context.Context, q *LCSQuery) *LCSCmd
	LIndex(ctx context.Context, key string, index int64) *StringCmd
	LInsert(ctx context.Context, key, op string, pivot, value any) *IntCmd
	LInsertBefore(ctx context.Context, key string, pivot, value any) *IntCmd
	LInsertAfter(ctx context.Context, key string, pivot, value any) *IntCmd
	LLen(ctx context.Context, key string) *IntCmd
	LMPop(ctx context.Context, direction string, count int64, keys ...string) *KeyValuesCmd
	LPop(ctx context.Context, key string) *StringCmd
	LPopCount(ctx context.Context, key string, count int64) *StringSliceCmd
	LPos(ctx context.Context, key string, value string, args LPosArgs) *IntCmd
	LPosCount(ctx context.Context, key string, value string, count int64, args LPosArgs) *IntSliceCmd
	LPush(ctx context.Context, key string, values ...any) *IntCmd
	LPushX(ctx context.Context, key string, values ...any) *IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	LRem(ctx context.Context, key string, count int64, value any) *IntCmd
	LSet(ctx context.Context, key string, index int64, value any) *StatusCmd
	LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd
	RPop(ctx context.Context, key string) *StringCmd
	RPopCount(ctx context.Context, key string, count int64) *StringSliceCmd
	RPopLPush(ctx context.Context, source, destination string) *StringCmd
	RPush(ctx context.Context, key string, values ...any) *IntCmd
	RPushX(ctx context.Context, key string, values ...any) *IntCmd
	LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd
	BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *StringCmd

	SAdd(ctx context.Context, key string, members ...any) *IntCmd
	SCard(ctx context.Context, key string) *IntCmd
	SDiff(ctx context.Context, keys ...string) *StringSliceCmd
	SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd
	SInter(ctx context.Context, keys ...string) *StringSliceCmd
	SInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd
	SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd
	SIsMember(ctx context.Context, key string, member any) *BoolCmd
	SMIsMember(ctx context.Context, key string, members ...any) *BoolSliceCmd
	SMembers(ctx context.Context, key string) *StringSliceCmd
	SMembersMap(ctx context.Context, key string) *StringStructMapCmd
	SMove(ctx context.Context, source, destination string, member any) *BoolCmd
	SPop(ctx context.Context, key string) *StringCmd
	SPopN(ctx context.Context, key string, count int64) *StringSliceCmd
	SRandMember(ctx context.Context, key string) *StringCmd
	SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd
	SRem(ctx context.Context, key string, members ...any) *IntCmd
	SUnion(ctx context.Context, keys ...string) *StringSliceCmd
	SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd

	XAdd(ctx context.Context, a XAddArgs) *StringCmd
	XDel(ctx context.Context, stream string, ids ...string) *IntCmd
	XLen(ctx context.Context, stream string) *IntCmd
	XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd
	XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd
	XRevRange(ctx context.Context, stream string, start, stop string) *XMessageSliceCmd
	XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *XMessageSliceCmd
	XRead(ctx context.Context, a XReadArgs) *XStreamSliceCmd
	XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd
	XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd
	XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd
	XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd
	XGroupDestroy(ctx context.Context, stream, group string) *IntCmd
	XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
	XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
	XReadGroup(ctx context.Context, a XReadGroupArgs) *XStreamSliceCmd
	XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd
	XPending(ctx context.Context, stream, group string) *XPendingCmd
	XPendingExt(ctx context.Context, a XPendingExtArgs) *XPendingExtCmd
	XClaim(ctx context.Context, a XClaimArgs) *XMessageSliceCmd
	XClaimJustID(ctx context.Context, a XClaimArgs) *StringSliceCmd
	XAutoClaim(ctx context.Context, a XAutoClaimArgs) *XAutoClaimCmd
	XAutoClaimJustID(ctx context.Context, a XAutoClaimArgs) *XAutoClaimJustIDCmd
	XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd
	XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd
	XTrimMinID(ctx context.Context, key string, minID string) *IntCmd
	XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd
	XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd
	XInfoStream(ctx context.Context, key string) *XInfoStreamCmd
	XInfoStreamFull(ctx context.Context, key string, count int64) *XInfoStreamFullCmd
	XInfoConsumers(ctx context.Context, key string, group string) *XInfoConsumersCmd
	XCfgSet(ctx context.Context, a XCfgSetArgs) *StatusCmd

	BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
	BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
	BZMPop(ctx context.Context, timeout time.Duration, order string, count int64, keys ...string) *ZSliceWithKeyCmd

	ZAdd(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddLT(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddGT(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddNX(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddXX(ctx context.Context, key string, members ...Z) *IntCmd
	ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd
	ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd
	ZCard(ctx context.Context, key string) *IntCmd
	ZCount(ctx context.Context, key, min, max string) *IntCmd
	ZLexCount(ctx context.Context, key, min, max string) *IntCmd
	ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd
	ZInter(ctx context.Context, store ZStore) *StringSliceCmd
	ZInterWithScores(ctx context.Context, store ZStore) *ZSliceCmd
	ZInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd
	ZInterStore(ctx context.Context, destination string, store ZStore) *IntCmd
	ZMPop(ctx context.Context, order string, count int64, keys ...string) *ZSliceWithKeyCmd
	ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd
	ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd
	ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd
	ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	ZRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd
	ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd
	ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd
	ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd
	ZRank(ctx context.Context, key, member string) *IntCmd
	ZRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd
	ZRem(ctx context.Context, key string, members ...any) *IntCmd
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd
	ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd
	ZRemRangeByLex(ctx context.Context, key, min, max string) *IntCmd
	ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
	ZRevRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRevRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd
	ZRevRank(ctx context.Context, key, member string) *IntCmd
	ZRevRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd
	ZScore(ctx context.Context, key, member string) *FloatCmd
	ZUnionStore(ctx context.Context, dest string, store ZStore) *IntCmd
	ZRandMember(ctx context.Context, key string, count int64) *StringSliceCmd
	ZRandMemberWithScores(ctx context.Context, key string, count int64) *ZSliceCmd
	ZUnion(ctx context.Context, store ZStore) *StringSliceCmd
	ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd
	ZDiff(ctx context.Context, keys ...string) *StringSliceCmd
	ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd
	ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd

	PFAdd(ctx context.Context, key string, els ...any) *IntCmd
	PFCount(ctx context.Context, keys ...string) *IntCmd
	PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd

	BgRewriteAOF(ctx context.Context) *StatusCmd
	BgSave(ctx context.Context) *StatusCmd
	ClientKill(ctx context.Context, ipPort string) *StatusCmd
	ClientKillByFilter(ctx context.Context, keys ...string) *IntCmd
	ClientList(ctx context.Context) *StringCmd
	ClientInfo(ctx context.Context) *ClientInfoCmd
	ClientPause(ctx context.Context, dur time.Duration) *BoolCmd
	ClientUnpause(ctx context.Context) *BoolCmd
	ClientID(ctx context.Context) *IntCmd
	ClientUnblock(ctx context.Context, id int64) *IntCmd
	ClientUnblockWithError(ctx context.Context, id int64) *IntCmd
	ConfigGet(ctx context.Context, parameter string) *StringStringMapCmd
	ConfigResetStat(ctx context.Context) *StatusCmd
	ConfigSet(ctx context.Context, parameter, value string) *StatusCmd
	ConfigRewrite(ctx context.Context) *StatusCmd
	DBSize(ctx context.Context) *IntCmd
	FlushAll(ctx context.Context) *StatusCmd
	FlushAllAsync(ctx context.Context) *StatusCmd
	FlushDB(ctx context.Context) *StatusCmd
	FlushDBAsync(ctx context.Context) *StatusCmd
	Info(ctx context.Context, section ...string) *StringCmd
	LastSave(ctx context.Context) *IntCmd
	Save(ctx context.Context) *StatusCmd
	Shutdown(ctx context.Context) *StatusCmd
	ShutdownSave(ctx context.Context) *StatusCmd
	ShutdownNoSave(ctx context.Context) *StatusCmd
	SlaveOf(ctx context.Context, host, port string) *StatusCmd
	SlowLogGet(ctx context.Context, num int64) *SlowLogCmd
	SlowLogReset(ctx context.Context) *StatusCmd
	Time(ctx context.Context) *TimeCmd
	DebugObject(ctx context.Context, key string) *StringCmd
	ReadOnly(ctx context.Context) *StatusCmd
	ReadWrite(ctx context.Context) *StatusCmd
	MemoryUsage(ctx context.Context, key string, samples ...int64) *IntCmd

	Eval(ctx context.Context, script string, keys []string, args ...any) *Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...any) *Cmd
	EvalRO(ctx context.Context, script string, keys []string, args ...any) *Cmd
	EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...any) *Cmd
	ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd
	ScriptFlush(ctx context.Context) *StatusCmd
	ScriptKill(ctx context.Context) *StatusCmd
	ScriptLoad(ctx context.Context, script string) *StringCmd

	FunctionLoad(ctx context.Context, code string) *StringCmd
	FunctionLoadReplace(ctx context.Context, code string) *StringCmd
	FunctionDelete(ctx context.Context, libName string) *StringCmd
	FunctionFlush(ctx context.Context) *StringCmd
	FunctionKill(ctx context.Context) *StringCmd
	FunctionFlushAsync(ctx context.Context) *StringCmd
	FunctionList(ctx context.Context, q FunctionListQuery) *FunctionListCmd
	FunctionDump(ctx context.Context) *StringCmd
	FunctionRestore(ctx context.Context, libDump string) *StringCmd
	FunctionStats(ctx context.Context) *FunctionStatsCmd
	FCall(ctx context.Context, function string, keys []string, args ...any) *Cmd
	FCallRO(ctx context.Context, function string, keys []string, args ...any) *Cmd

	Publish(ctx context.Context, channel string, message any) *IntCmd
	SPublish(ctx context.Context, channel string, message any) *IntCmd
	PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd
	PubSubNumSub(ctx context.Context, channels ...string) *StringIntMapCmd
	PubSubNumPat(ctx context.Context) *IntCmd
	PubSubShardChannels(ctx context.Context, pattern string) *StringSliceCmd
	PubSubShardNumSub(ctx context.Context, channels ...string) *StringIntMapCmd

	ClusterMyShardID(ctx context.Context) *StringCmd
	ClusterSlots(ctx context.Context) *ClusterSlotsCmd
	ClusterShards(ctx context.Context) *ClusterShardsCmd
	ClusterLinks(ctx context.Context) *ClusterLinksCmd
	ClusterNodes(ctx context.Context) *StringCmd
	ClusterMeet(ctx context.Context, host string, port int64) *StatusCmd
	ClusterForget(ctx context.Context, nodeID string) *StatusCmd
	ClusterReplicate(ctx context.Context, nodeID string) *StatusCmd
	ClusterResetSoft(ctx context.Context) *StatusCmd
	ClusterResetHard(ctx context.Context) *StatusCmd
	ClusterInfo(ctx context.Context) *StringCmd
	ClusterKeySlot(ctx context.Context, key string) *IntCmd
	ClusterGetKeysInSlot(ctx context.Context, slot int64, count int64) *StringSliceCmd
	ClusterCountFailureReports(ctx context.Context, nodeID string) *IntCmd
	ClusterCountKeysInSlot(ctx context.Context, slot int64) *IntCmd
	ClusterDelSlots(ctx context.Context, slots ...int64) *StatusCmd
	ClusterDelSlotsRange(ctx context.Context, min, max int64) *StatusCmd
	ClusterSaveConfig(ctx context.Context) *StatusCmd
	ClusterSlaves(ctx context.Context, nodeID string) *StringSliceCmd
	ClusterFailover(ctx context.Context) *StatusCmd
	ClusterAddSlots(ctx context.Context, slots ...int64) *StatusCmd
	ClusterAddSlotsRange(ctx context.Context, min, max int64) *StatusCmd

	GeoAdd(ctx context.Context, key string, geoLocation ...GeoLocation) *IntCmd
	GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *IntCmd
	GeoRadiusByMember(ctx context.Context, key, member string, query GeoRadiusQuery) *GeoLocationCmd
	GeoRadiusByMemberStore(ctx context.Context, key, member string, query GeoRadiusQuery) *IntCmd
	GeoSearch(ctx context.Context, key string, q GeoSearchQuery) *StringSliceCmd
	GeoSearchLocation(ctx context.Context, key string, q GeoSearchLocationQuery) *GeoLocationCmd
	GeoSearchStore(ctx context.Context, key, store string, q GeoSearchStoreQuery) *IntCmd
	GeoDist(ctx context.Context, key string, member1, member2, unit string) *FloatCmd
	GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd

	ACLDryRun(ctx context.Context, username string, command ...any) *StringCmd
	ACLLog(ctx context.Context, count int64) *ACLLogCmd
	ACLSetUser(ctx context.Context, username string, rules ...string) *StatusCmd
	ACLDelUser(ctx context.Context, username string) *IntCmd
	ACLLogReset(ctx context.Context) *StatusCmd
	ACLList(ctx context.Context) *StringSliceCmd
	ACLCat(ctx context.Context) *StringSliceCmd
	ACLCatArgs(ctx context.Context, options *ACLCatArgs) *StringSliceCmd

	ModuleLoadex(ctx context.Context, conf *ModuleLoadexConfig) *StringCmd
	GearsCmdable
	ProbabilisticCmdable
	TimeseriesCmdable
	JSONCmdable
	SearchCmdable
}

type SearchCmdable interface {
	FT_List(ctx context.Context) *StringSliceCmd
	FTAggregate(ctx context.Context, index string, query string) *MapStringInterfaceCmd
	FTAggregateWithArgs(ctx context.Context, index string, query string, options *FTAggregateOptions) *AggregateCmd
	FTAliasAdd(ctx context.Context, index string, alias string) *StatusCmd
	FTAliasDel(ctx context.Context, alias string) *StatusCmd
	FTAliasUpdate(ctx context.Context, index string, alias string) *StatusCmd
	FTAlter(ctx context.Context, index string, skipInitialScan bool, definition []interface{}) *StatusCmd
	FTConfigGet(ctx context.Context, option string) *MapMapStringInterfaceCmd
	FTConfigSet(ctx context.Context, option string, value interface{}) *StatusCmd
	FTCreate(ctx context.Context, index string, options *FTCreateOptions, schema ...*FieldSchema) *StatusCmd
	FTCursorDel(ctx context.Context, index string, cursorId int) *StatusCmd
	FTCursorRead(ctx context.Context, index string, cursorId int, count int) *MapStringInterfaceCmd
	FTDictAdd(ctx context.Context, dict string, term ...interface{}) *IntCmd
	FTDictDel(ctx context.Context, dict string, term ...interface{}) *IntCmd
	FTDictDump(ctx context.Context, dict string) *StringSliceCmd
	FTDropIndex(ctx context.Context, index string) *StatusCmd
	FTDropIndexWithArgs(ctx context.Context, index string, options *FTDropIndexOptions) *StatusCmd
	FTExplain(ctx context.Context, index string, query string) *StringCmd
	FTExplainWithArgs(ctx context.Context, index string, query string, options *FTExplainOptions) *StringCmd
	FTInfo(ctx context.Context, index string) *FTInfoCmd
	FTSpellCheck(ctx context.Context, index string, query string) *FTSpellCheckCmd
	FTSpellCheckWithArgs(ctx context.Context, index string, query string, options *FTSpellCheckOptions) *FTSpellCheckCmd
	FTSearch(ctx context.Context, index string, query string) *FTSearchCmd
	FTSearchWithArgs(ctx context.Context, index string, query string, options *FTSearchOptions) *FTSearchCmd
	FTSynDump(ctx context.Context, index string) *FTSynDumpCmd
	FTSynUpdate(ctx context.Context, index string, synGroupId interface{}, terms []interface{}) *StatusCmd
	FTSynUpdateWithArgs(ctx context.Context, index string, synGroupId interface{}, options *FTSynUpdateOptions, terms []interface{}) *StatusCmd
	FTTagVals(ctx context.Context, index string, field string) *StringSliceCmd
}

// https://github.com/redis/go-redis/blob/af4872cbd0de349855ce3f0978929c2f56eb995f/probabilistic.go#L10
type ProbabilisticCmdable interface {
	BFAdd(ctx context.Context, key string, element interface{}) *BoolCmd
	BFCard(ctx context.Context, key string) *IntCmd
	BFExists(ctx context.Context, key string, element interface{}) *BoolCmd
	BFInfo(ctx context.Context, key string) *BFInfoCmd
	BFInfoArg(ctx context.Context, key, option string) *BFInfoCmd
	BFInfoCapacity(ctx context.Context, key string) *BFInfoCmd
	BFInfoSize(ctx context.Context, key string) *BFInfoCmd
	BFInfoFilters(ctx context.Context, key string) *BFInfoCmd
	BFInfoItems(ctx context.Context, key string) *BFInfoCmd
	BFInfoExpansion(ctx context.Context, key string) *BFInfoCmd
	BFInsert(ctx context.Context, key string, options *BFInsertOptions, elements ...interface{}) *BoolSliceCmd
	BFMAdd(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd
	BFMExists(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd
	BFReserve(ctx context.Context, key string, errorRate float64, capacity int64) *StatusCmd
	BFReserveExpansion(ctx context.Context, key string, errorRate float64, capacity, expansion int64) *StatusCmd
	BFReserveNonScaling(ctx context.Context, key string, errorRate float64, capacity int64) *StatusCmd
	BFReserveWithArgs(ctx context.Context, key string, options *BFReserveOptions) *StatusCmd
	BFScanDump(ctx context.Context, key string, iterator int64) *ScanDumpCmd
	BFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *StatusCmd

	CFAdd(ctx context.Context, key string, element interface{}) *BoolCmd
	CFAddNX(ctx context.Context, key string, element interface{}) *BoolCmd
	CFCount(ctx context.Context, key string, element interface{}) *IntCmd
	CFDel(ctx context.Context, key string, element interface{}) *BoolCmd
	CFExists(ctx context.Context, key string, element interface{}) *BoolCmd
	CFInfo(ctx context.Context, key string) *CFInfoCmd
	CFInsert(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *BoolSliceCmd
	CFInsertNX(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *IntSliceCmd
	CFMExists(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd
	CFReserve(ctx context.Context, key string, capacity int64) *StatusCmd
	CFReserveWithArgs(ctx context.Context, key string, options *CFReserveOptions) *StatusCmd
	CFReserveExpansion(ctx context.Context, key string, capacity int64, expansion int64) *StatusCmd
	CFReserveBucketSize(ctx context.Context, key string, capacity int64, bucketsize int64) *StatusCmd
	CFReserveMaxIterations(ctx context.Context, key string, capacity int64, maxiterations int64) *StatusCmd
	CFScanDump(ctx context.Context, key string, iterator int64) *ScanDumpCmd
	CFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *StatusCmd

	CMSIncrBy(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd
	CMSInfo(ctx context.Context, key string) *CMSInfoCmd
	CMSInitByDim(ctx context.Context, key string, width, height int64) *StatusCmd
	CMSInitByProb(ctx context.Context, key string, errorRate, probability float64) *StatusCmd
	CMSMerge(ctx context.Context, destKey string, sourceKeys ...string) *StatusCmd
	CMSMergeWithWeight(ctx context.Context, destKey string, sourceKeys map[string]int64) *StatusCmd
	CMSQuery(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd

	TopKAdd(ctx context.Context, key string, elements ...interface{}) *StringSliceCmd
	TopKCount(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd
	TopKIncrBy(ctx context.Context, key string, elements ...interface{}) *StringSliceCmd
	TopKInfo(ctx context.Context, key string) *TopKInfoCmd
	TopKList(ctx context.Context, key string) *StringSliceCmd
	TopKListWithCount(ctx context.Context, key string) *MapStringIntCmd
	TopKQuery(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd
	TopKReserve(ctx context.Context, key string, k int64) *StatusCmd
	TopKReserveWithOptions(ctx context.Context, key string, k int64, width, depth int64, decay float64) *StatusCmd

	TDigestAdd(ctx context.Context, key string, elements ...float64) *StatusCmd
	TDigestByRank(ctx context.Context, key string, rank ...uint64) *FloatSliceCmd
	TDigestByRevRank(ctx context.Context, key string, rank ...uint64) *FloatSliceCmd
	TDigestCDF(ctx context.Context, key string, elements ...float64) *FloatSliceCmd
	TDigestCreate(ctx context.Context, key string) *StatusCmd
	TDigestCreateWithCompression(ctx context.Context, key string, compression int64) *StatusCmd
	TDigestInfo(ctx context.Context, key string) *TDigestInfoCmd
	TDigestMax(ctx context.Context, key string) *FloatCmd
	TDigestMin(ctx context.Context, key string) *FloatCmd
	TDigestMerge(ctx context.Context, destKey string, options *TDigestMergeOptions, sourceKeys ...string) *StatusCmd
	TDigestQuantile(ctx context.Context, key string, elements ...float64) *FloatSliceCmd
	TDigestRank(ctx context.Context, key string, values ...float64) *IntSliceCmd
	TDigestReset(ctx context.Context, key string) *StatusCmd
	TDigestRevRank(ctx context.Context, key string, values ...float64) *IntSliceCmd
	TDigestTrimmedMean(ctx context.Context, key string, lowCutQuantile, highCutQuantile float64) *FloatCmd

	Pipeline() Pipeliner
	Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error)

	TxPipeline() Pipeliner
	TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error)
}

// Align with go-redis
// https://github.com/redis/go-redis/blob/f994ff1cd96299a5c8029ae3403af7b17ef06e8a/gears_commands.go#L9-L19
type GearsCmdable interface {
	TFunctionLoad(ctx context.Context, lib string) *StatusCmd
	TFunctionLoadArgs(ctx context.Context, lib string, options *TFunctionLoadOptions) *StatusCmd
	TFunctionDelete(ctx context.Context, libName string) *StatusCmd
	TFunctionList(ctx context.Context) *MapStringInterfaceSliceCmd
	TFunctionListArgs(ctx context.Context, options *TFunctionListOptions) *MapStringInterfaceSliceCmd
	TFCall(ctx context.Context, libName string, funcName string, numKeys int) *Cmd
	TFCallArgs(ctx context.Context, libName string, funcName string, numKeys int, options *TFCallOptions) *Cmd
	TFCallASYNC(ctx context.Context, libName string, funcName string, numKeys int) *Cmd
	TFCallASYNCArgs(ctx context.Context, libName string, funcName string, numKeys int, options *TFCallOptions) *Cmd
}

type TimeseriesCmdable interface {
	TSAdd(ctx context.Context, key string, timestamp interface{}, value float64) *IntCmd
	TSAddWithArgs(ctx context.Context, key string, timestamp interface{}, value float64, options *TSOptions) *IntCmd
	TSCreate(ctx context.Context, key string) *StatusCmd
	TSCreateWithArgs(ctx context.Context, key string, options *TSOptions) *StatusCmd
	TSAlter(ctx context.Context, key string, options *TSAlterOptions) *StatusCmd
	TSCreateRule(ctx context.Context, sourceKey string, destKey string, aggregator Aggregator, bucketDuration int) *StatusCmd
	TSCreateRuleWithArgs(ctx context.Context, sourceKey string, destKey string, aggregator Aggregator, bucketDuration int, options *TSCreateRuleOptions) *StatusCmd
	TSIncrBy(ctx context.Context, Key string, timestamp float64) *IntCmd
	TSIncrByWithArgs(ctx context.Context, key string, timestamp float64, options *TSIncrDecrOptions) *IntCmd
	TSDecrBy(ctx context.Context, Key string, timestamp float64) *IntCmd
	TSDecrByWithArgs(ctx context.Context, key string, timestamp float64, options *TSIncrDecrOptions) *IntCmd
	TSDel(ctx context.Context, Key string, fromTimestamp int, toTimestamp int) *IntCmd
	TSDeleteRule(ctx context.Context, sourceKey string, destKey string) *StatusCmd
	TSGet(ctx context.Context, key string) *TSTimestampValueCmd
	TSGetWithArgs(ctx context.Context, key string, options *TSGetOptions) *TSTimestampValueCmd
	TSInfo(ctx context.Context, key string) *MapStringInterfaceCmd
	TSInfoWithArgs(ctx context.Context, key string, options *TSInfoOptions) *MapStringInterfaceCmd
	TSMAdd(ctx context.Context, ktvSlices [][]interface{}) *IntSliceCmd
	TSQueryIndex(ctx context.Context, filterExpr []string) *StringSliceCmd
	TSRevRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *TSTimestampValueSliceCmd
	TSRevRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *TSRevRangeOptions) *TSTimestampValueSliceCmd
	TSRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *TSTimestampValueSliceCmd
	TSRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *TSRangeOptions) *TSTimestampValueSliceCmd
	TSMRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *MapStringSliceInterfaceCmd
	TSMRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *TSMRangeOptions) *MapStringSliceInterfaceCmd
	TSMRevRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *MapStringSliceInterfaceCmd
	TSMRevRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *TSMRevRangeOptions) *MapStringSliceInterfaceCmd
	TSMGet(ctx context.Context, filters []string) *MapStringSliceInterfaceCmd
	TSMGetWithArgs(ctx context.Context, filters []string, options *TSMGetOptions) *MapStringSliceInterfaceCmd
}

type JSONCmdable interface {
	JSONArrAppend(ctx context.Context, key, path string, values ...interface{}) *IntSliceCmd
	JSONArrIndex(ctx context.Context, key, path string, value ...interface{}) *IntSliceCmd
	JSONArrIndexWithArgs(ctx context.Context, key, path string, options *JSONArrIndexArgs, value ...interface{}) *IntSliceCmd
	JSONArrInsert(ctx context.Context, key, path string, index int64, values ...interface{}) *IntSliceCmd
	JSONArrLen(ctx context.Context, key, path string) *IntSliceCmd
	JSONArrPop(ctx context.Context, key, path string, index int) *StringSliceCmd
	JSONArrTrim(ctx context.Context, key, path string) *IntSliceCmd
	JSONArrTrimWithArgs(ctx context.Context, key, path string, options *JSONArrTrimArgs) *IntSliceCmd
	JSONClear(ctx context.Context, key, path string) *IntCmd
	JSONDebugMemory(ctx context.Context, key, path string) *IntCmd
	JSONDel(ctx context.Context, key, path string) *IntCmd
	JSONForget(ctx context.Context, key, path string) *IntCmd
	JSONGet(ctx context.Context, key string, paths ...string) *JSONCmd
	JSONGetWithArgs(ctx context.Context, key string, options *JSONGetArgs, paths ...string) *JSONCmd
	JSONMerge(ctx context.Context, key, path string, value string) *StatusCmd
	JSONMSetArgs(ctx context.Context, docs []JSONSetArgs) *StatusCmd
	JSONMSet(ctx context.Context, params ...interface{}) *StatusCmd
	JSONMGet(ctx context.Context, path string, keys ...string) *JSONSliceCmd
	JSONNumIncrBy(ctx context.Context, key, path string, value float64) *JSONCmd
	JSONObjKeys(ctx context.Context, key, path string) *SliceCmd
	JSONObjLen(ctx context.Context, key, path string) *IntPointerSliceCmd
	JSONSet(ctx context.Context, key, path string, value interface{}) *StatusCmd
	JSONSetMode(ctx context.Context, key, path string, value interface{}, mode string) *StatusCmd
	JSONStrAppend(ctx context.Context, key, path, value string) *IntPointerSliceCmd
	JSONStrLen(ctx context.Context, key, path string) *IntPointerSliceCmd
	JSONToggle(ctx context.Context, key, path string) *IntPointerSliceCmd
	JSONType(ctx context.Context, key, path string) *JSONSliceCmd
}

var _ Cmdable = (*Compat)(nil)

type Compat struct {
	client rueidis.Client
	maxp   int
	pOnly  bool
}

// CacheCompat implements commands that support client-side caching.
type CacheCompat struct {
	client rueidis.Client
	ttl    time.Duration
}

// AdapterOption is a functional option type for NewAdapter
type AdapterOption func(c *Compat)

// WithNodeScaleoutLimit sets the maximum parallelism for node scaleout operations.
// If not set, defaults to runtime.GOMAXPROCS(0).
// Values less than 1 will be set to 1.
func WithNodeScaleoutLimit(limit int) AdapterOption {
	_ = "STUB: not implemented"
	return *new(AdapterOption)
}

func NewAdapter(client rueidis.Client, options ...AdapterOption) Cmdable {
	_ = "STUB: not implemented"
	return *new(Cmdable)
}

func (c *Compat) Client() rueidis.Client {
	_ = "STUB: not implemented"

	// ForEachMaster concurrently calls the fn on each master node in the cluster.
	// It returns the first error if any.
	return *new(rueidis.Client)
}

func (c *Compat) ForEachMaster(ctx context.Context, fn func(ctx context.Context, client Cmdable) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Cache(ttl time.Duration) CacheCompat {
	_ = "STUB: not implemented"
	return *new(CacheCompat)
}

func (c *Compat) Command(ctx context.Context) *CommandsInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

type FilterBy struct {
	Module  string
	ACLCat  string
	Pattern string
}

func (c *Compat) CommandList(ctx context.Context, filter FilterBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CommandGetKeys(ctx context.Context, commands ...any) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CommandGetKeysAndFlags(ctx context.Context, commands ...any) *KeyFlagsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClientGetName(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Echo(ctx context.Context, message any) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Ping(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) Quit(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) Del(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Unlink(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Dump(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Exists(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Expire(ctx context.Context, key string, seconds time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ExpireAt(ctx context.Context, key string, timestamp time.Time) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ExpireTime(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ExpireNX(ctx context.Context, key string, seconds time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ExpireXX(ctx context.Context, key string, seconds time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ExpireGT(ctx context.Context, key string, seconds time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ExpireLT(ctx context.Context, key string, seconds time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Keys(ctx context.Context, pattern string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Migrate(ctx context.Context, host string, port int64, key string, db int64, timeout time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Move(ctx context.Context, key string, db int64) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ObjectRefCount(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ObjectEncoding(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ObjectIdleTime(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Persist(ctx context.Context, key string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) PExpire(ctx context.Context, key string, milliseconds time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) PExpireAt(ctx context.Context, key string, millisecondsTimestamp time.Time) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) PExpireTime(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) PTTL(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) RandomKey(ctx context.Context) *StringCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) Rename(ctx context.Context, key, newkey string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) RenameNX(ctx context.Context, key, newkey string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Restore(ctx context.Context, key string, ttl time.Duration, serializedValue string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) RestoreReplace(ctx context.Context, key string, ttl time.Duration, serializedValue string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) sort(command, key string, sort Sort) cmds.Arbitrary {
	_ = "STUB: not implemented"
	return *new(cmds.Arbitrary)
}

func (c *Compat) Sort(ctx context.Context, key string, sort Sort) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SortRO(ctx context.Context, key string, sort Sort) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SortStore(ctx context.Context, key, store string, sort Sort) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SortInterfaces(ctx context.Context, key string, sort Sort) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Touch(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TTL(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Type(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Append(ctx context.Context, key, value string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Decr(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) DecrBy(ctx context.Context, key string, decrement int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Get(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) GetRange(ctx context.Context, key string, start, end int64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) GetSet(ctx context.Context, key string, value any) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

// GetEx An expiration of zero removes the TTL associated with the key (i.e., GETEX key persist).
// Requires Redis >= 6.2.0.
func (c *Compat) GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) GetDel(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Incr(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) IncrBy(ctx context.Context, key string, increment int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) IncrByFloat(ctx context.Context, key string, increment float64) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) MGet(ctx context.Context, keys ...string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) MSet(ctx context.Context, values ...any) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) MSetNX(ctx context.Context, values ...any) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

// Set key value [expiration]
//
// For no expiration use 0.
//
// For KEEPTTL use -1.
//
// For more options, use SetArgs.
func (c *Compat) Set(ctx context.Context, key string, value any, expiration time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SetArgs(ctx context.Context, key string, value any, a SetArgs) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SetEX(ctx context.Context, key string, value any, expiration time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SetNX(ctx context.Context, key string, value any, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SetXX(ctx context.Context, key string, value any, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) StrLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Copy(ctx context.Context, source string, destination string, db int64, replace bool) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) GetBit(ctx context.Context, key string, offset int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SetBit(ctx context.Context, key string, offset int64, value int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BitOpNot(ctx context.Context, destKey string, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BitPosSpan(ctx context.Context, key string, bit, start, end int64, span string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BitField(ctx context.Context, key string, args ...any) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BitFieldRO(ctx context.Context, key string, args ...any) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HScanNoValues(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HDel(ctx context.Context, key string, fields ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HExists(ctx context.Context, key, field string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HGet(ctx context.Context, key, field string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HGetAll(ctx context.Context, key string) *StringStringMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HKeys(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HStrLen(ctx context.Context, key, field string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HMGet(ctx context.Context, key string, fields ...string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// HSet requires Redis v4 for multiple field/value pairs support.
func (c *Compat) HSet(ctx context.Context, key string, values ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// HMSet is a deprecated version of HSet left for compatibility with Redis 3.
func (c *Compat) HMSet(ctx context.Context, key string, values ...any) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HSetNX(ctx context.Context, key, field string, value any) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HVals(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HRandField(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HRandFieldWithValues(ctx context.Context, key string, count int64) *KeyValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HPExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HPExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HPExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HPExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HPersist(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HExpireTime(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HPExpireTime(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HTTL(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HPTTL(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HGetDel(ctx context.Context, key string, fields ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HGetEX(ctx context.Context, key string, fields ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HGetEXWithArgs(ctx context.Context, key string, options *HGetEXOptions, fields ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HSetEX(ctx context.Context, key string, fieldsAndValues ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) HSetEXWithArgs(ctx context.Context, key string, options *HSetEXOptions, fieldsAndValues ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *KeyValuesCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LCS(ctx context.Context, q *LCSQuery) *LCSCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LIndex(ctx context.Context, key string, index int64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LInsert(ctx context.Context, key, op string, pivot, element any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LInsertBefore(ctx context.Context, key string, pivot, element any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LInsertAfter(ctx context.Context, key string, pivot, element any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LMPop(ctx context.Context, direction string, count int64, keys ...string) *KeyValuesCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LPop(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LPopCount(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LPos(ctx context.Context, key string, element string, a LPosArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LPosCount(ctx context.Context, key string, element string, count int64, a LPosArgs) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LPush(ctx context.Context, key string, elements ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LPushX(ctx context.Context, key string, elements ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LRem(ctx context.Context, key string, count int64, element any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LSet(ctx context.Context, key string, index int64, element any) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) RPop(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) RPopCount(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) RPopLPush(ctx context.Context, source, destination string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) RPush(ctx context.Context, key string, elements ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) RPushX(ctx context.Context, key string, elements ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SAdd(ctx context.Context, key string, members ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SDiff(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SInter(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SIsMember(ctx context.Context, key string, member any) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SMIsMember(ctx context.Context, key string, members ...any) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SMembers(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SMembersMap(ctx context.Context, key string) *StringStructMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SMove(ctx context.Context, source, destination string, member any) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SPop(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SPopN(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SRandMember(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SRem(ctx context.Context, key string, members ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SUnion(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XAdd(ctx context.Context, a XAddArgs) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XDel(ctx context.Context, stream string, ids ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XLen(ctx context.Context, stream string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XRevRange(ctx context.Context, stream, stop, start string) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XRevRangeN(ctx context.Context, stream, stop, start string, count int64) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XRead(ctx context.Context, a XReadArgs) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XGroupDestroy(ctx context.Context, stream, group string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XReadGroup(ctx context.Context, a XReadGroupArgs) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XPending(ctx context.Context, stream, group string) *XPendingCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XPendingExt(ctx context.Context, a XPendingExtArgs) *XPendingExtCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XClaim(ctx context.Context, a XClaimArgs) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XClaimJustID(ctx context.Context, a XClaimArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XAutoClaim(ctx context.Context, a XAutoClaimArgs) *XAutoClaimCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XAutoClaimJustID(ctx context.Context, a XAutoClaimArgs) *XAutoClaimJustIDCmd {
	_ = "STUB: not implemented"
	return nil
}

// xTrim If approx is true, add the "~" parameter; otherwise it is the default "=" (redis default).
// example:
//
//	XTRIM key MAXLEN/MINID = threshold LIMIT limit.
//	XTRIM key MAXLEN/MINID ~ threshold LIMIT limit.
//
// The redis-server version is lower than 6.2, please set the limit to 0.
func (c *Compat) xTrim(ctx context.Context, key, strategy string,
	approx bool, threshold string, limit int64,
) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// XTrimMaxLen No `~` rules are used, `limit` cannot be used.
// cmd: XTRIM key MAXLEN maxLen
func (c *Compat) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// XTrimMaxLenApprox LIMIT has a bug, please confirm it and use it.
// issue: https://github.com/redis/redis/issues/9046
// cmd: XTRIM key MAXLEN ~ maxLen LIMIT limit
func (c *Compat) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// XTrimMinID No `~` rules are used, `limit` cannot be used.
// cmd: XTRIM key MINID minID
func (c *Compat) XTrimMinID(ctx context.Context, key string, minID string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// XTrimMinIDApprox LIMIT has a bug, please confirm it and use it.
// issue: https://github.com/redis/redis/issues/9046
// cmd: XTRIM key MINID ~ minID LIMIT limit
func (c *Compat) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XInfoStream(ctx context.Context, key string) *XInfoStreamCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XInfoStreamFull(ctx context.Context, key string, count int64) *XInfoStreamFullCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XInfoConsumers(ctx context.Context, key, group string) *XInfoConsumersCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) XCfgSet(ctx context.Context, a XCfgSetArgs) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// BZPopMax Redis `BZPOPMAX key [key ...] timeout` command.
func (c *Compat) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

// BZPopMin Redis `BZPOPMIN key [key ...] timeout` command.
func (c *Compat) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BZMPop(ctx context.Context, timeout time.Duration, order string, count int64, keys ...string) *ZSliceWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

// ZAdd Redis `ZADD key score member [score member ...]` command.
func (c *Compat) ZAdd(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// ZAddNX Redis `ZADD key NX score member [score member ...]` command.
func (c *Compat) ZAddNX(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// ZAddXX Redis `ZADD key XX score member [score member ...]` command.
func (c *Compat) ZAddXX(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZAddLT(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZAddGT(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) zAddArgs(ctx context.Context, key string, incr bool, args ZAddArgs) rueidis.RedisResult {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

// The GT, LT and NX options are mutually exclusive.

func (c *Compat) ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZCount(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZLexCount(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func zstore(cmd cmds.Arbitrary, store ZStore) cmds.Arbitrary {
	_ = "STUB: not implemented"
	return *new(cmds.Arbitrary)
}

func (c *Compat) ZInter(ctx context.Context, store ZStore) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZInterWithScores(ctx context.Context, store ZStore) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZInterStore(ctx context.Context, destination string, store ZStore) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZMPop(ctx context.Context, order string, count int64, keys ...string) *ZSliceWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) zRangeArgs(withScores bool, z ZRangeArgs) rueidis.Completed {
	_ = "STUB: not implemented"
	return *new(rueidis.Completed)
}

func (c *Compat) ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRank(ctx context.Context, key, member string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRem(ctx context.Context, key string, members ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRemRangeByLex(ctx context.Context, key string, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRevRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRevRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRevRank(ctx context.Context, key, member string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRevRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZScore(ctx context.Context, key, member string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZUnionStore(ctx context.Context, dest string, store ZStore) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZUnion(ctx context.Context, store ZStore) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRandMember(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZRandMemberWithScores(ctx context.Context, key string, count int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZDiff(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) PFAdd(ctx context.Context, key string, els ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) PFCount(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BgRewriteAOF(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BgSave(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) ClientKill(ctx context.Context, ipPort string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClientKillByFilter(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClientList(ctx context.Context) *StringCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) ClientPause(ctx context.Context, dur time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClientUnpause(ctx context.Context) *BoolCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) ClientID(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) ClientUnblock(ctx context.Context, id int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClientUnblockWithError(ctx context.Context, id int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClientInfo(ctx context.Context) *ClientInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ConfigGet(ctx context.Context, parameter string) *StringStringMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ConfigResetStat(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ConfigSet(ctx context.Context, parameter, value string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ConfigRewrite(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) DBSize(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) FlushAll(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) FlushAllAsync(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FlushDB(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) FlushDBAsync(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Info(ctx context.Context, section ...string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) LastSave(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) Save(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) Shutdown(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) ShutdownSave(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ShutdownNoSave(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SlaveOf(ctx context.Context, host, port string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SlowLogGet(ctx context.Context, num int64) *SlowLogCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SlowLogReset(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Time(ctx context.Context) *TimeCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) DebugObject(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ReadOnly(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) ReadWrite(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) MemoryUsage(ctx context.Context, key string, samples ...int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Eval(ctx context.Context, script string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) EvalSha(ctx context.Context, sha1 string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) EvalRO(ctx context.Context, script string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ScriptFlush(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) ScriptKill(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) ScriptLoad(ctx context.Context, script string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FunctionLoad(ctx context.Context, code string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FunctionLoadReplace(ctx context.Context, code string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FunctionDelete(ctx context.Context, libName string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FunctionFlush(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FunctionKill(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FunctionFlushAsync(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FunctionList(ctx context.Context, q FunctionListQuery) *FunctionListCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FunctionDump(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FunctionRestore(ctx context.Context, libDump string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FCall(ctx context.Context, function string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FCallRO(ctx context.Context, function string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Publish(ctx context.Context, channel string, message any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) SPublish(ctx context.Context, channel string, message any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) PubSubNumSub(ctx context.Context, channels ...string) *StringIntMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) PubSubNumPat(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) PubSubShardChannels(ctx context.Context, pattern string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) PubSubShardNumSub(ctx context.Context, channels ...string) *StringIntMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterMyShardID(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterSlots(ctx context.Context) *ClusterSlotsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterShards(ctx context.Context) *ClusterShardsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterNodes(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterLinks(ctx context.Context) *ClusterLinksCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterMeet(ctx context.Context, host string, port int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterForget(ctx context.Context, nodeID string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterReplicate(ctx context.Context, nodeID string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterResetSoft(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterResetHard(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterInfo(ctx context.Context) *StringCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) ClusterKeySlot(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterGetKeysInSlot(ctx context.Context, slot int64, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterCountFailureReports(ctx context.Context, nodeID string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterCountKeysInSlot(ctx context.Context, slot int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterDelSlots(ctx context.Context, slots ...int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterDelSlotsRange(ctx context.Context, min, max int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterSaveConfig(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterSlaves(ctx context.Context, nodeID string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterFailover(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterAddSlots(ctx context.Context, slots ...int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ClusterAddSlotsRange(ctx context.Context, min, max int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) GeoAdd(ctx context.Context, key string, geoLocation ...GeoLocation) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd {
	_ = "STUB: not implemented"
	return nil
}

// GeoRadius is a read-only GEORADIUS_RO command.
func (c *Compat) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

// GeoRadiusStore is a writing GEORADIUS command.
func (c *Compat) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// GeoRadiusByMember is a read-only GEORADIUSBYMEMBER_RO command.
func (c *Compat) GeoRadiusByMember(ctx context.Context, key, member string, query GeoRadiusQuery) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

// GeoRadiusByMemberStore is a writing GEORADIUSBYMEMBER command.
func (c *Compat) GeoRadiusByMemberStore(ctx context.Context, key, member string, query GeoRadiusQuery) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) GeoSearch(ctx context.Context, key string, q GeoSearchQuery) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) GeoSearchLocation(ctx context.Context, key string, q GeoSearchLocationQuery) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) GeoSearchStore(ctx context.Context, src, dest string, q GeoSearchStoreQuery) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) GeoDist(ctx context.Context, key, member1, member2, unit string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FunctionStats(ctx context.Context) *FunctionStatsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ACLDryRun(ctx context.Context, username string, command ...any) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

type ACLCatArgs struct {
	Category string
}

func (c *Compat) ACLCatArgs(ctx context.Context, options *ACLCatArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	// if there is a category passed, build new cmd, if there isn't - use the ACLCat method
	return nil
}

func (c *Compat) ACLLog(ctx context.Context, count int64) *ACLLogCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ACLSetUser(ctx context.Context, username string, rules ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ACLLogReset(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) ACLCat(ctx context.Context) *StringSliceCmd { _ = "STUB: not implemented"; return nil }

func (c *Compat) ACLList(ctx context.Context) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ACLDelUser(ctx context.Context, username string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) doPrimaries(ctx context.Context, fn func(c rueidis.Client) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) doStringCmdPrimaries(ctx context.Context, fn func(c rueidis.Client) rueidis.Completed) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) doIntCmdPrimaries(ctx context.Context, fn func(c rueidis.Client) rueidis.Completed) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TFunctionLoad(ctx context.Context, lib string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TFunctionLoadArgs(ctx context.Context, lib string, options *TFunctionLoadOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TFunctionDelete(ctx context.Context, libName string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TFunctionList(ctx context.Context) *MapStringInterfaceSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TFunctionListArgs(ctx context.Context, options *TFunctionListOptions) *MapStringInterfaceSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TFCall(ctx context.Context, libName string, funcName string, numKeys int) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TFCallArgs(ctx context.Context, libName string, funcName string, numKeys int, options *TFCallOptions) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TFCallASYNC(ctx context.Context, libName string, funcName string, numKeys int) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TFCallASYNCArgs(ctx context.Context, libName string, funcName string, numKeys int, options *TFCallOptions) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFAdd(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFExists(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFInfo(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFInfoArg(ctx context.Context, key, option string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFInfoCapacity(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFInfoSize(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFInfoFilters(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFInfoItems(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFInfoExpansion(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFInsert(ctx context.Context, key string, options *BFInsertOptions, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFMAdd(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFMExists(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFReserve(ctx context.Context, key string, errorRate float64, capacity int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFReserveExpansion(ctx context.Context, key string, errorRate float64, capacity, expansion int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFReserveNonScaling(ctx context.Context, key string, errorRate float64, capacity int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFReserveWithArgs(ctx context.Context, key string, options *BFReserveOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFScanDump(ctx context.Context, key string, iterator int64) *ScanDumpCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) BFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFAdd(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFAddNX(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFCount(ctx context.Context, key string, element interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFDel(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFExists(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFInfo(ctx context.Context, key string) *CFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFInsert(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFInsertNX(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFMExists(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFReserve(ctx context.Context, key string, capacity int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFReserveWithArgs(ctx context.Context, key string, options *CFReserveOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFReserveExpansion(ctx context.Context, key string, capacity int64, expansion int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFReserveBucketSize(ctx context.Context, key string, capacity int64, bucketsize int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFReserveMaxIterations(ctx context.Context, key string, capacity int64, maxiterations int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFScanDump(ctx context.Context, key string, iterator int64) *ScanDumpCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CMSIncrBy(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CMSInfo(ctx context.Context, key string) *CMSInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CMSInitByDim(ctx context.Context, key string, width, height int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CMSInitByProb(ctx context.Context, key string, errorRate, probability float64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CMSMerge(ctx context.Context, destKey string, sourceKeys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) CMSMergeWithWeight(ctx context.Context, destKey string, sourceKeys map[string]int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// weight should be integer
// we convert int64 to float64 to avoid API breaking change

func (c *Compat) CMSQuery(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TopKAdd(ctx context.Context, key string, elements ...interface{}) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TopKCount(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TopKIncrBy(ctx context.Context, key string, elements ...interface{}) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TopKInfo(ctx context.Context, key string) *TopKInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TopKList(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TopKListWithCount(ctx context.Context, key string) *MapStringIntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TopKQuery(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TopKReserve(ctx context.Context, key string, k int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TopKReserveWithOptions(ctx context.Context, key string, k int64, width, depth int64, decay float64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestAdd(ctx context.Context, key string, elements ...float64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestByRank(ctx context.Context, key string, rank ...uint64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestByRevRank(ctx context.Context, key string, rank ...uint64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestCDF(ctx context.Context, key string, elements ...float64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestCreate(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestCreateWithCompression(ctx context.Context, key string, compression int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestInfo(ctx context.Context, key string) *TDigestInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestMax(ctx context.Context, key string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestMin(ctx context.Context, key string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestMerge(ctx context.Context, destKey string, options *TDigestMergeOptions, sourceKeys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestQuantile(ctx context.Context, key string, elements ...float64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestRank(ctx context.Context, key string, values ...float64) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestReset(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestRevRank(ctx context.Context, key string, values ...float64) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) TDigestTrimmedMean(ctx context.Context, key string, lowCutQuantile, highCutQuantile float64) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSAdd - Adds one or more observations to a t-digest sketch.
// For more information - https://redis.io/commands/ts.add/
func (c *Compat) TSAdd(ctx context.Context, key string, timestamp interface{}, value float64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSAddWithArgs - Adds one or more observations to a t-digest sketch.
// This function also allows for specifying additional options such as
// Retention, ChunkSize, Encoding, DuplicatePolicy and Labels.
// For more information - https://redis.io/commands/ts.add/
func (c *Compat) TSAddWithArgs(ctx context.Context, key string, timestamp interface{}, value float64, options *TSOptions) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSCreate - Creates a new time-series key.
// For more information - https://redis.io/commands/ts.create/
func (c *Compat) TSCreate(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSCreateWithArgs - Creates a new time-series key with additional options.
// This function allows for specifying additional options such as
// Retention, ChunkSize, Encoding, DuplicatePolicy and Labels.
// For more information - https://redis.io/commands/ts.create/
func (c *Compat) TSCreateWithArgs(ctx context.Context, key string, options *TSOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// var cmd cmds.Completed

// TSAlter - Alters an existing time-series key with additional options.
// This function allows for specifying additional options such as
// Retention, ChunkSize and DuplicatePolicy.
// For more information - https://redis.io/commands/ts.alter/
func (c *Compat) TSAlter(ctx context.Context, key string, options *TSAlterOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSCreateRule - Creates a compaction rule from sourceKey to destKey.
// For more information - https://redis.io/commands/ts.createrule/
func (c *Compat) TSCreateRule(ctx context.Context, sourceKey string, destKey string, aggregator Aggregator, bucketDuration int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSCreateRuleWithArgs - Creates a compaction rule from sourceKey to destKey with an additional option.
// This function allows for specifying an additional option such as
// AlignTimestamp.
// For more information - https://redis.io/commands/ts.createrule/
func (c *Compat) TSCreateRuleWithArgs(ctx context.Context, sourceKey string, destKey string, aggregator Aggregator, bucketDuration int, options *TSCreateRuleOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSIncrBy - Increments the value of a time-series key by the specified timestamp.
// For more information - https://redis.io/commands/ts.incrby/
// FIXME: timestamp should be addend
func (c *Compat) TSIncrBy(ctx context.Context, Key string, timestamp float64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSIncrByWithArgs - Increments the value of a time-series key by the specified timestamp with additional options.
// This function allows for specifying additional options such as:
// Timestamp, Retention, ChunkSize, Uncompressed and Labels.
// For more information - https://redis.io/commands/ts.incrby/
// FIXME: timestamp should be addend
func (c *Compat) TSIncrByWithArgs(ctx context.Context, key string, timestamp float64, options *TSIncrDecrOptions) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSDecrBy - Decrements the value of a time-series key by the specified timestamp.
// For more information - https://redis.io/commands/ts.decrby/
// FIXME: timestamp should be subtrahend
func (c *Compat) TSDecrBy(ctx context.Context, Key string, timestamp float64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSDecrByWithArgs - Decrements the value of a time-series key by the specified timestamp with additional options.
// This function allows for specifying additional options such as:
// Timestamp, Retention, ChunkSize, Uncompressed and Labels.
// For more information - https://redis.io/commands/ts.decrby/
func (c *Compat) TSDecrByWithArgs(ctx context.Context, key string, timestamp float64, options *TSIncrDecrOptions) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSDel - Deletes a range of samples from a time-series key.
// For more information - https://redis.io/commands/ts.del/
func (c *Compat) TSDel(ctx context.Context, Key string, fromTimestamp int, toTimestamp int) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSDeleteRule - Deletes a compaction rule from sourceKey to destKey.
// For more information - https://redis.io/commands/ts.deleterule/
func (c *Compat) TSDeleteRule(ctx context.Context, sourceKey string, destKey string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSGetWithArgs - Gets the last sample of a time-series key with an additional option.
// This function allows for specifying an additional option such as
// Latest.
// For more information - https://redis.io/commands/ts.get/
func (c *Compat) TSGetWithArgs(ctx context.Context, key string, options *TSGetOptions) *TSTimestampValueCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSGet - Gets the last sample of a time-series key.
// For more information - https://redis.io/commands/ts.get/
func (c *Compat) TSGet(ctx context.Context, key string) *TSTimestampValueCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSInfo - Returns information about a time-series key.
// For more information - https://redis.io/commands/ts.info/
func (c *Compat) TSInfo(ctx context.Context, key string) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSInfoWithArgs - Returns information about a time-series key with additional option.
// This function allows for specifying additional option such as:
// Debug.
// For more information - https://redis.io/commands/ts.info/
func (c *Compat) TSInfoWithArgs(ctx context.Context, key string, options *TSInfoOptions) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

// FIXME: should not accept arg, just append "DEBUG"

// TSMAdd - Adds multiple samples to multiple time-series keys.
// For more information - https://redis.io/commands/ts.madd/
func (c *Compat) TSMAdd(ctx context.Context, ktvSlices [][]interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSQueryIndex - Returns all the keys matching the filter expression.
// For more information - https://redis.io/commands/ts.queryindex/
func (c *Compat) TSQueryIndex(ctx context.Context, filterExpr []string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSRevRange - Returns a range of samples from a time-series key in reverse order.
// For more information - https://redis.io/commands/ts.revrange/
func (c *Compat) TSRevRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSRevRangeWithArgs - Returns a range of samples from a time-series key in reverse order with additional options.
// This function allows for specifying additional options such as:
// Latest, FilterByTS, FilterByValue, Count, Align, Aggregator,
// BucketDuration, BucketTimestamp and Empty.
// For more information - https://redis.io/commands/ts.revrange/
func (c *Compat) TSRevRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *TSRevRangeOptions) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSRange - Returns a range of samples from a time-series key.
// For more information - https://redis.io/commands/ts.range/
func (c *Compat) TSRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSRangeWithArgs - Returns a range of samples from a time-series key with additional options.
// This function allows for specifying additional options such as:
// Latest, FilterByTS, FilterByValue, Count, Align, Aggregator,
// BucketDuration, BucketTimestamp and Empty.
// For more information - https://redis.io/commands/ts.range/
func (c *Compat) TSRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *TSRangeOptions) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSMRange - Returns a range of samples from multiple time-series keys.
// For more information - https://redis.io/commands/ts.mrange/
func (c *Compat) TSMRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSMRangeWithArgs - Returns a range of samples from multiple time-series keys with additional options.
// This function allows for specifying additional options such as:
// Latest, FilterByTS, FilterByValue, WithLabels, SelectedLabels,
// Count, Align, Aggregator, BucketDuration, BucketTimestamp,
// Empty, GroupByLabel and Reducer.
// For more information - https://redis.io/commands/ts.mrange/
func (c *Compat) TSMRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *TSMRangeOptions) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

// FIXME: Wrong API definition: REDUCE

// TSMRevRange - Returns a range of samples from multiple time-series keys in reverse order.
// For more information - https://redis.io/commands/ts.mrevrange/
func (c *Compat) TSMRevRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSMRevRangeWithArgs - Returns a range of samples from multiple time-series keys in reverse order with additional options.
// This function allows for specifying additional options such as:
// Latest, FilterByTS, FilterByValue, WithLabels, SelectedLabels,
// Count, Align, Aggregator, BucketDuration, BucketTimestamp,
// Empty, GroupByLabel and Reducer.
// For more information - https://redis.io/commands/ts.mrevrange/
func (c *Compat) TSMRevRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *TSMRevRangeOptions) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

// FIXME: Wrong API definition: REDUCE

// TSMGet - Returns the last sample of multiple time-series keys.
// For more information - https://redis.io/commands/ts.mget/
func (c *Compat) TSMGet(ctx context.Context, filters []string) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

// TSMGetWithArgs - Returns the last sample of multiple time-series keys with additional options.
// This function allows for specifying additional options such as:
// Latest, WithLabels and SelectedLabels.
// For more information - https://redis.io/commands/ts.mget/
func (c *Compat) TSMGetWithArgs(ctx context.Context, filters []string, options *TSMGetOptions) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

// JSONArrAppend adds the provided JSON values to the end of the array at the given path.
// For more information, see https://redis.io/commands/json.arrappend
func (c *Compat) JSONArrAppend(ctx context.Context, key, path string, values ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// JSONArrIndex searches for the first occurrence of the provided JSON value in the array at the given path.
// For more information, see https://redis.io/commands/json.arrindex
// NOTE: value should have the format value start [stop]
func (c *Compat) JSONArrIndex(ctx context.Context, key, path string, value ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// format: value

// format: value start

// format: value start stop

// JSONArrIndex searches for the first occurrence of the provided JSON value in the array at the given path.
// For more information, see https://redis.io/commands/json.arrindex
func (c *Compat) JSONArrIndexWithArgs(ctx context.Context, key, path string, options *JSONArrIndexArgs, value ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	// FIXME: why value has 1..N ?
	return nil
}

func (c *Compat) JSONArrInsert(ctx context.Context, key, path string, index int64, values ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONArrLen(ctx context.Context, key, path string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONArrPop(ctx context.Context, key, path string, index int) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONArrTrim(ctx context.Context, key, path string) *IntSliceCmd {
	_ = "STUB: not implemented"
	// both default value of start and stop are 0
	// Ref: https://redis.io/commands/json.arrtrim/
	return nil
}

func (c *Compat) JSONArrTrimWithArgs(ctx context.Context, key, path string, options *JSONArrTrimArgs) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONClear(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONDebugMemory(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONDel(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONForget(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONGet(ctx context.Context, key string, paths ...string) *JSONCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONGetWithArgs(ctx context.Context, key string, options *JSONGetArgs, paths ...string) *JSONCmd {
	_ = "STUB: not implemented"
	// _cmd := c.client.B().JsonGet().Key(key).Path(paths...)
	return nil
}

func (c *Compat) JSONMerge(ctx context.Context, key, path string, value string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONMSetArgs(ctx context.Context, docs []JSONSetArgs) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONMSet(ctx context.Context, params ...interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONMGet(ctx context.Context, path string, keys ...string) *JSONSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONNumIncrBy(ctx context.Context, key, path string, value float64) *JSONCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONObjKeys(ctx context.Context, key, path string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONObjLen(ctx context.Context, key, path string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONSet(ctx context.Context, key, path string, value interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// JSONSetMode sets the JSON value at the given path in the given key and allows the mode to be set
// (the mode value must be "XX" or "NX"). The value must be something that can be marshaled to JSON (using encoding/JSON) unless
// the argument is a string or []byte when we assume that it can be passed directly as JSON.
// For more information, see https://redis.io/commands/json.set
func (c *Compat) JSONSetMode(ctx context.Context, key, path string, value interface{}, mode string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONStrAppend(ctx context.Context, key, path, value string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONStrLen(ctx context.Context, key, path string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONToggle(ctx context.Context, key, path string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) JSONType(ctx context.Context, key, path string) *JSONSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) Subscribe(ctx context.Context, channels ...string) PubSub {
	_ = "STUB: not implemented"
	return *new(PubSub)
}

func (c *Compat) SSubscribe(ctx context.Context, channels ...string) PubSub {
	_ = "STUB: not implemented"
	return *new(PubSub)
}

func (c *Compat) PSubscribe(ctx context.Context, patterns ...string) PubSub {
	_ = "STUB: not implemented"
	return *new(PubSub)
}

func (c *Compat) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Compat) Pipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *Compat) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Compat) TxPipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *Compat) Watch(ctx context.Context, fn func(Tx) error, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) FT_List(ctx context.Context) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTAggregate - Performs a search query on an index and applies a series of aggregate transformations to the result.
// The 'index' parameter specifies the index to search, and the 'query' parameter specifies the search query.
// For more information, please refer to the Redis documentation:
// [FT.AGGREGATE]: (https://redis.io/commands/ft.aggregate/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L473
func (c *Compat) FTAggregate(ctx context.Context, index string, query string) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTAggregateWithArgs - Performs a search query on an index and applies a series of aggregate transformations to the result.
// The 'index' parameter specifies the index to search, and the 'query' parameter specifies the search query.
// This function also allows for specifying additional options such as: Verbatim, LoadAll, Load, Timeout, GroupBy, SortBy, SortByMax, Apply, LimitOffset, Limit, Filter, WithCursor, Params, and DialectVersion.
// For more information, please refer to the Redis documentation:
// [FT.AGGREGATE]: (https://redis.io/commands/ft.aggregate/)
// see: go-redis v9.7.0: https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L671
func (c *Compat) FTAggregateWithArgs(ctx context.Context, index string, query string, options *FTAggregateOptions) *AggregateCmd {
	_ = "STUB: not implemented"
	return nil
}

// [VERBATIM]

// [SCORER]

// [ADDSCORES]

// [LOAD count field [field ...]]

// LOAD *

// [TIMEOUT timeout]

// [ APPLY expression AS name [ APPLY expression AS name ...]]

// [ GROUPBY nargs property [property ...] [ REDUCE function nargs arg [arg ...] [AS name] [ REDUCE function nargs arg [arg ...] [AS name] ...]] ...]]

// [ SORTBY nargs [ property ASC | DESC [ property ASC | DESC ...]] [MAX num] [WITHCOUNT]

// count the number of args to be passed in to cmds.FtAggregateQuery(_cmd).Sortby()

// ASC

// DESC

// FIXME: go-redis doesn't provide WITHCOUNT option
// [ LIMIT offset num]

// [FILTER filter]

// [ WITHCURSOR [COUNT read_size] [MAXIDLE idle_time]]

// [ PARAMS nargs name value [ name value ...]]

// [ADDSCORES]: NOTE: go-redis doesn't implement this option.
// [DIALECT dialect]

// FTAliasAdd - Adds an alias to an index.
// The 'index' parameter specifies the index to which the alias is added, and the 'alias' parameter specifies the alias.
// For more information, please refer to the Redis documentation:
// [FT.ALIASADD]: (https://redis.io/commands/ft.aliasadd/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L782
func (c *Compat) FTAliasAdd(ctx context.Context, index string, alias string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTAliasDel - Removes an alias from an index.
// The 'alias' parameter specifies the alias to be removed.
// For more information, please refer to the Redis documentation:
// [FT.ALIASDEL]: (https://redis.io/commands/ft.aliasdel/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L793
func (c *Compat) FTAliasDel(ctx context.Context, alias string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTAliasUpdate - Updates an alias to an index.
// The 'index' parameter specifies the index to which the alias is updated, and the 'alias' parameter specifies the alias.
// If the alias already exists for a different index, it updates the alias to point to the specified index instead.
// For more information, please refer to the Redis documentation:
// [FT.ALIASUPDATE]: (https://redis.io/commands/ft.aliasupdate/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L804
func (c *Compat) FTAliasUpdate(ctx context.Context, index string, alias string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTAlter - Alters the definition of an existing index.
// The 'index' parameter specifies the index to alter, and the 'skipInitialScan' parameter specifies whether to skip the initial scan.
// The 'definition' parameter specifies the new definition for the index.
// For more information, please refer to the Redis documentation:
// [FT.ALTER]: (https://redis.io/commands/ft.alter/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L815
func (c *Compat) FTAlter(ctx context.Context, index string, skipInitialScan bool, definition []interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTConfigGet - Retrieves the value of a RediSearch configuration parameter.
// The 'option' parameter specifies the configuration parameter to retrieve.
// For more information, please refer to the Redis documentation:
// [FT.CONFIG GET]: (https://redis.io/commands/ft.config-get/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L831
func (c *Compat) FTConfigGet(ctx context.Context, option string) *MapMapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTConfigSet - Sets the value of a RediSearch configuration parameter.
// The 'option' parameter specifies the configuration parameter to set, and the 'value' parameter specifies the new value.
// For more information, please refer to the Redis documentation:
// [FT.CONFIG SET]: (https://redis.io/commands/ft.config-set/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L841
func (c *Compat) FTConfigSet(ctx context.Context, option string, value interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTCreate - Creates a new index with the given options and schema.
// The 'index' parameter specifies the name of the index to create.
// The 'options' parameter specifies various options for the index, such as
// whether to index hashes or JSONs, prefixes, filters, default language, score, score field, payload field, etc.
// The 'schema' parameter specifies the schema for the index, which includes the field name, field type, etc.
// For more information, please refer to the Redis documentation:
// [FT.CREATE]: (https://redis.io/commands/ft.create/)
// FTCreate aligns with go-redis v9.7.0.
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L854
func (c *Compat) FTCreate(ctx context.Context, index string, options *FTCreateOptions, schema ...*FieldSchema) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// [ON HASH | JSON]

// [PREFIX count prefix [prefix ...]]

// [FILTER {filter}]

// [LANGUAGE default_lang]

// [LANGUAGE_FIELD lang_attribute]

// [SCORE default_score]

// [SCORE_FIELD score_attribute]

// [PAYLOAD_FIELD payload_attribute]

// [MAXTEXTFIELDS]
// FIXME: in go-redis, FTCreateOptions.MaxTextFields should be bool, not int

// [TEMPORARY seconds]
// FIXME: reudis: Temporary should not be float64

// [NOOFFSETS]

// [NOHL]

// [NOFIELDS]

// [NOFREQS]

// [STOPWORDS count [stopword ...]]

// [SKIPINITIALSCAN]

// 	SCHEMA field_name [AS alias] TEXT | TAG | NUMERIC | GEO | VECTOR | GEOSHAPE [ SORTABLE [UNF]]
//   [NOINDEX] [ field_name [AS alias] TEXT | TAG | NUMERIC | GEO | VECTOR | GEOSHAPE [ SORTABLE [UNF]] [NOINDEX] ...]

// Ref: https://redis.io/docs/latest/develop/interact/search-and-query/advanced-concepts/vectors/#create-a-vector-index

// FIXME: redis doc: PHONETIC not in EBNF definition

// FTCursorDel - Deletes a cursor from an existing index.
// The 'index' parameter specifies the index from which to delete the cursor, and the 'cursorId' parameter specifies the ID of the cursor to delete.
// For more information, please refer to the Redis documentation:
// [FT.CURSOR DEL]: (https://redis.io/commands/ft.cursor-del/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1032
func (c *Compat) FTCursorDel(ctx context.Context, index string, cursorId int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTCursorRead - Reads the next results from an existing cursor.
// The 'index' parameter specifies the index from which to read the cursor, the 'cursorId' parameter specifies the ID of the cursor to read, and the 'count' parameter specifies the number of results to read.
// For more information, please refer to the Redis documentation:
// [FT.CURSOR READ]: (https://redis.io/commands/ft.cursor-read/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1042
func (c *Compat) FTCursorRead(ctx context.Context, index string, cursorId int, count int) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTDictAdd - Adds terms to a dictionary.
// The 'dict' parameter specifies the dictionary to which to add the terms, and the 'term' parameter specifies the terms to add.
// For more information, please refer to the Redis documentation:
// [FT.DICTADD]: (https://redis.io/commands/ft.dictadd/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1056
func (c *Compat) FTDictAdd(ctx context.Context, dict string, term ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTDictDel - Deletes terms from a dictionary.
// The 'dict' parameter specifies the dictionary from which to delete the terms, and the 'term' parameter specifies the terms to delete.
// For more information, please refer to the Redis documentation:
// [FT.DICTDEL]: (https://redis.io/commands/ft.dictdel/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1068
func (c *Compat) FTDictDel(ctx context.Context, dict string, term ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTDictDump - Returns all terms in the specified dictionary.
// The 'dict' parameter specifies the dictionary from which to return the terms.
// For more information, please refer to the Redis documentation:
// [FT.DICTDUMP]: (https://redis.io/commands/ft.dictdump/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1080
func (c *Compat) FTDictDump(ctx context.Context, dict string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTDropIndex - Deletes an index.
// The 'index' parameter specifies the index to delete.
// For more information, please refer to the Redis documentation:
// [FT.DROPINDEX]: (https://redis.io/commands/ft.dropindex/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1090
func (c *Compat) FTDropIndex(ctx context.Context, index string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTDropIndexWithArgs - Deletes an index with options.
// The 'index' parameter specifies the index to delete, and the 'options' parameter specifies the DeleteDocs option for doc deletion.
// For more information, please refer to the Redis documentation:
// [FT.DROPINDEX]: (https://redis.io/commands/ft.dropindex/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1101
func (c *Compat) FTDropIndexWithArgs(ctx context.Context, index string, options *FTDropIndexOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTExplain - Returns the execution plan for a complex query.
// The 'index' parameter specifies the index to query, and the 'query' parameter specifies the query string.
// For more information, please refer to the Redis documentation:
// [FT.EXPLAIN]: (https://redis.io/commands/ft.explain/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1117
func (c *Compat) FTExplain(ctx context.Context, index string, query string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTExplainWithArgs - Returns the execution plan for a complex query with options.
// The 'index' parameter specifies the index to query, the 'query' parameter specifies the query string, and the 'options' parameter specifies the Dialect for the query.
// For more information, please refer to the Redis documentation:
// [FT.EXPLAIN]: (https://redis.io/commands/ft.explain/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1127
func (c *Compat) FTExplainWithArgs(ctx context.Context, index string, query string, options *FTExplainOptions) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTInfo - Retrieves information about an index.
// The 'index' parameter specifies the index to retrieve information about.
// For more information, please refer to the Redis documentation:
// [FT.INFO]: (https://redis.io/commands/ft.info/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1393
func (c *Compat) FTInfo(ctx context.Context, index string) *FTInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTSpellCheck - Checks a query string for spelling errors.
// For more details about a spellcheck query please follow:
// https://redis.io/docs/interact/search-and-query/advanced-concepts/spellcheck/
// For more information, please refer to the Redis documentation:
// [FT.SPELLCHECK]: (https://redis.io/commands/ft.spellcheck/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1404
func (c *Compat) FTSpellCheck(ctx context.Context, index string, query string) *FTSpellCheckCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTSpellCheckWithArgs - Checks a query string for spelling errors with additional options.
// For more details about a spellcheck query please follow:
// https://redis.io/docs/interact/search-and-query/advanced-concepts/spellcheck/
// For more information, please refer to the Redis documentation:
// [FT.SPELLCHECK]: (https://redis.io/commands/ft.spellcheck/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1416
func (c *Compat) FTSpellCheckWithArgs(ctx context.Context, index string, query string, options *FTSpellCheckOptions) *FTSpellCheckCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTSearch - Executes a search query on an index.
// The 'index' parameter specifies the index to search, and the 'query' parameter specifies the search query.
// For more information, please refer to the Redis documentation:
// [FT.SEARCH]: (https://redis.io/commands/ft.search/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1679
func (c *Compat) FTSearch(ctx context.Context, index string, query string) *FTSearchCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTSearchWithArgs - Executes a search query on an index with additional options.
// The 'index' parameter specifies the index to search, the 'query' parameter specifies the search query,
// and the 'options' parameter specifies additional options for the search.
// For more information, please refer to the Redis documentation:
// [FT.SEARCH]: (https://redis.io/commands/ft.search/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1802
func (c *Compat) FTSearchWithArgs(ctx context.Context, index string, query string, options *FTSearchOptions) *FTSearchCmd {
	_ = "STUB: not implemented"
	return nil
}

// [NOCONTENT]

// [VERBATIM]

// [NOSTOPWORDS]

// [WITHSCORES]

// [WITHPAYLOADS]

// [WITHSORTKEYS]

// [FILTER numeric_field min max [ FILTER numeric_field min max ...]]

//  [GEOFILTER geo_field lon lat radius m | km | mi | ft [ GEOFILTER geo_field lon lat radius m | km | mi | ft ...]]

// [INKEYS count key [key ...]]

// [ INFIELDS count field [field ...]]

// [RETURN count identifier [AS property] [ identifier [AS property] ...]]

// FIXME: go-redis doesn't implement SUMMARIZE option
// [SUMMARIZE [ FIELDS count field [field ...]] [FRAGS num] [LEN fragsize] [SEPARATOR separator]]
// [SLOP slop]

// [TIMEOUT timeout]

// [INORDER]

// [LANGUAGE language]

// [EXPANDER expander]

// [SCORER scorer]

// [EXPLAINSCORE]

// [PAYLOAD payload]

// [SORTBY sortby [ ASC | DESC] [WITHCOUNT]]

// [LIMIT offset num]

// [PARAMS nargs name value [ name value ...]]

// [DIALECT dialect]

// FTSynDump - Dumps the contents of a synonym group.
// The 'index' parameter specifies the index to dump.
// For more information, please refer to the Redis documentation:
// [FT.SYNDUMP]: (https://redis.io/commands/ft.syndump/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1987
func (c *Compat) FTSynDump(ctx context.Context, index string) *FTSynDumpCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTSynUpdate - Creates or updates a synonym group with additional terms.
// The 'index' parameter specifies the index to update, the 'synGroupId' parameter specifies the synonym group id, and the 'terms' parameter specifies the additional terms.
// For more information, please refer to the Redis documentation:
// [FT.SYNUPDATE]: (https://redis.io/commands/ft.synupdate/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L1997
func (c *Compat) FTSynUpdate(ctx context.Context, index string, synGroupId interface{}, terms []interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTSynUpdateWithArgs - Creates or updates a synonym group with additional terms and options.
// The 'index' parameter specifies the index to update, the 'synGroupId' parameter specifies the synonym group id, the 'options' parameter specifies additional options for the update, and the 'terms' parameter specifies the additional terms.
// For more information, please refer to the Redis documentation:
// [FT.SYNUPDATE]: (https://redis.io/commands/ft.synupdate/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L2009
func (c *Compat) FTSynUpdateWithArgs(ctx context.Context, index string, synGroupId interface{}, options *FTSynUpdateOptions, terms []interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

// FTTagVals - Returns all distinct values indexed in a tag field.
// The 'index' parameter specifies the index to check, and the 'field' parameter specifies the tag field to retrieve values from.
// For more information, please refer to the Redis documentation:
// [FT.TAGVALS]: (https://redis.io/commands/ft.tagvals/)
// see go-redis v9.7.0 https://github.com/redis/go-redis/blob/v9.7.0/search_commands.go#L2024
func (c *Compat) FTTagVals(ctx context.Context, index string, field string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Compat) ModuleLoadex(ctx context.Context, conf *ModuleLoadexConfig) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BitPosSpan(ctx context.Context, key string, bit, start, end int64, span string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BitFieldRO(ctx context.Context, key string, args ...any) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) EvalRO(ctx context.Context, script string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) FCallRO(ctx context.Context, function string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) GeoDist(ctx context.Context, key, member1, member2, unit string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd {
	_ = "STUB: not implemented"
	return nil
}

// GeoRadius is a read-only GEORADIUS_RO command.
func (c CacheCompat) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

// GeoRadiusByMember is a read-only GEORADIUSBYMEMBER_RO command.
func (c CacheCompat) GeoRadiusByMember(ctx context.Context, key, member string, query GeoRadiusQuery) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) GeoSearch(ctx context.Context, key string, q GeoSearchQuery) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) Get(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) MGet(ctx context.Context, keys ...string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) GetBit(ctx context.Context, key string, offset int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) GetRange(ctx context.Context, key string, start, end int64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) HExists(ctx context.Context, key, field string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) HGet(ctx context.Context, key, field string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) HGetAll(ctx context.Context, key string) *StringStringMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) HKeys(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) HLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) HStrLen(ctx context.Context, key, field string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) HMGet(ctx context.Context, key string, fields ...string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) HVals(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) LIndex(ctx context.Context, key string, index int64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) LLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) LPos(ctx context.Context, key string, element string, a LPosArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) PTTL(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) SCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) SIsMember(ctx context.Context, key string, member any) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) SMIsMember(ctx context.Context, key string, members ...any) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) SMembers(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) SortRO(ctx context.Context, key string, sort Sort) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) StrLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) TTL(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) Type(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZCount(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZLexCount(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) zRangeArgs(withScores bool, z ZRangeArgs) rueidis.Cacheable {
	_ = "STUB: not implemented"
	return *new(rueidis.Cacheable)
}

func (c CacheCompat) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRank(ctx context.Context, key, member string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRevRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRevRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRevRank(ctx context.Context, key, member string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZRevRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) ZScore(ctx context.Context, key, member string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BFExists(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BFInfo(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BFInfoArg(ctx context.Context, key, option string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BFInfoCapacity(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BFInfoSize(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BFInfoFilters(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BFInfoItems(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) BFInfoExpansion(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) CFCount(ctx context.Context, key string, element interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) CFExists(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) CFInfo(ctx context.Context, key string) *CFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) CMSInfo(ctx context.Context, key string) *CMSInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) CMSQuery(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) TopKInfo(ctx context.Context, key string) *TopKInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) TopKList(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) TopKQuery(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) JSONArrIndex(ctx context.Context, key, path string, value ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

// format: value

// format: value start

// format: value start stop

func (c CacheCompat) JSONArrLen(ctx context.Context, key, path string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) JSONGet(ctx context.Context, key string, paths ...string) *JSONCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) JSONMGet(ctx context.Context, path string, keys ...string) *JSONSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) JSONObjKeys(ctx context.Context, key, path string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) JSONObjLen(ctx context.Context, key, path string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) JSONStrLen(ctx context.Context, key, path string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c CacheCompat) JSONType(ctx context.Context, key, path string) *JSONSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func str(arg any) string { _ = "STUB: not implemented"; return "" }

func argsToSlice(src []any) []string { _ = "STUB: not implemented"; return nil }

func argToSlice(arg any) []string { _ = "STUB: not implemented"; return nil }

// scan struct field

func appendStructField(v reflect.Value) []string { _ = "STUB: not implemented"; return nil }

// miss field

// if it's a nil pointer

// if it's a valid pointer

func omitEmpty(opt string) bool { _ = "STUB: not implemented"; return false }

func isEmptyValue(v reflect.Value) bool { _ = "STUB: not implemented"; return false }
