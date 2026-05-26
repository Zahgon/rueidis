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
	"errors"
	"time"

	"github.com/redis/rueidis"
)

// Pipeliner is a mechanism to realise Redis Pipeline technique.
//
// Pipelining is a technique to extremely speed up processing by packing
// operations to batches, send them at once to Redis and read a reply in a
// single step.
// See https://redis.io/topics/pipelining
//
// Pay attention that Pipeline is not a transaction, so you can get unexpected
// results in case of big pipelines and small read/write timeouts.
// Redis client has retransmission logic in case of timeouts, pipelines
// can be retransmitted, and commands can be executed more than once.
// To avoid this: it is a good idea to use reasonable bigger read/write timeouts
// depends on your batch size and/or use TxPipeline.
type Pipeliner interface {
	Cmdable

	// Len is to obtain the number of commands in the pipeline that have not yet been executed.
	Len() int

	// Do is an API for executing any command.
	// If a certain Redis command is not yet supported, you can use Do to execute it.
	Do(ctx context.Context, args ...interface{}) *Cmd

	// Discard is to discard all commands in the cache that have not yet been executed.
	Discard()

	// Exec is to send all the commands buffered in the pipeline to the redis-server.
	Exec(ctx context.Context) ([]Cmder, error)
}

var _ Pipeliner = (*Pipeline)(nil)
var _ Cmdable = (*Pipeline)(nil)

type proxyresult struct {
	err error
	val rueidis.RedisMessage
}

var placeholder = proxyresult{err: errors.New("the pipeline has not been executed")}

type proxy struct {
	rueidis.Client
	cmds []rueidis.Completed
}

func (p *proxy) Do(_ context.Context, cmd rueidis.Completed) rueidis.RedisResult {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

func newPipeline(real rueidis.Client) *Pipeline { _ = "STUB: not implemented"; return nil }

// Pipeline implements pipelining as described in
// http://redis.io/topics/pipelining.
// Please note: it is not safe for concurrent use by multiple goroutines.
type Pipeline struct {
	comp Compat
	rets []Cmder
}

func (c *Pipeline) Command(ctx context.Context) *CommandsInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CommandList(ctx context.Context, filter FilterBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CommandGetKeys(ctx context.Context, commands ...any) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CommandGetKeysAndFlags(ctx context.Context, commands ...any) *KeyFlagsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClientGetName(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Echo(ctx context.Context, message any) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Ping(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) Quit(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) Del(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Unlink(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Dump(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Exists(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Expire(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ExpireTime(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ExpireNX(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ExpireXX(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ExpireGT(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ExpireLT(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Keys(ctx context.Context, pattern string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Migrate(ctx context.Context, host string, port int64, key string, db int64, timeout time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Move(ctx context.Context, key string, db int64) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ObjectRefCount(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ObjectEncoding(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ObjectIdleTime(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Persist(ctx context.Context, key string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) PExpire(ctx context.Context, key string, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) PExpireAt(ctx context.Context, key string, tm time.Time) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) PExpireTime(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) PTTL(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) RandomKey(ctx context.Context) *StringCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) Rename(ctx context.Context, key, newkey string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) RenameNX(ctx context.Context, key, newkey string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Restore(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Sort(ctx context.Context, key string, sort Sort) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SortRO(ctx context.Context, key string, sort Sort) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SortStore(ctx context.Context, key, store string, sort Sort) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SortInterfaces(ctx context.Context, key string, sort Sort) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Touch(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TTL(ctx context.Context, key string) *DurationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Type(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Append(ctx context.Context, key, value string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Decr(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) DecrBy(ctx context.Context, key string, decrement int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Get(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GetRange(ctx context.Context, key string, start, end int64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GetSet(ctx context.Context, key string, value any) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GetEx(ctx context.Context, key string, expiration time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GetDel(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Incr(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) IncrBy(ctx context.Context, key string, value int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) IncrByFloat(ctx context.Context, key string, value float64) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) MGet(ctx context.Context, keys ...string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) MSet(ctx context.Context, values ...any) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) MSetNX(ctx context.Context, values ...any) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Set(ctx context.Context, key string, value any, expiration time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SetArgs(ctx context.Context, key string, value any, a SetArgs) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SetEX(ctx context.Context, key string, value any, expiration time.Duration) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SetNX(ctx context.Context, key string, value any, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SetXX(ctx context.Context, key string, value any, expiration time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) StrLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Copy(ctx context.Context, sourceKey string, destKey string, db int64, replace bool) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GetBit(ctx context.Context, key string, offset int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SetBit(ctx context.Context, key string, offset int64, value int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BitOpNot(ctx context.Context, destKey string, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BitPosSpan(ctx context.Context, key string, bit int64, start, end int64, span string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BitField(ctx context.Context, key string, args ...any) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BitFieldRO(ctx context.Context, key string, args ...any) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HScanNoValues(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HDel(ctx context.Context, key string, fields ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HExists(ctx context.Context, key, field string) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HGet(ctx context.Context, key, field string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HGetAll(ctx context.Context, key string) *StringStringMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HKeys(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HStrLen(ctx context.Context, key, field string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HMGet(ctx context.Context, key string, fields ...string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HSet(ctx context.Context, key string, values ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HMSet(ctx context.Context, key string, values ...any) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HSetNX(ctx context.Context, key, field string, value any) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HVals(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HRandField(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HRandFieldWithValues(ctx context.Context, key string, count int64) *KeyValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HPExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HPExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HPExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HPExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs HExpireArgs, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HPersist(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HExpireTime(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HPExpireTime(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HTTL(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HPTTL(ctx context.Context, key string, fields ...string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HGetDel(ctx context.Context, key string, fields ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HGetEX(ctx context.Context, key string, fields ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HGetEXWithArgs(ctx context.Context, key string, options *HGetEXOptions, fields ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HSetEX(ctx context.Context, key string, fieldsAndValues ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) HSetEXWithArgs(ctx context.Context, key string, options *HSetEXOptions, fieldsAndValues ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *KeyValuesCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LCS(ctx context.Context, q *LCSQuery) *LCSCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LIndex(ctx context.Context, key string, index int64) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LInsert(ctx context.Context, key, op string, pivot, value any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LInsertBefore(ctx context.Context, key string, pivot, value any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LInsertAfter(ctx context.Context, key string, pivot, value any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LLen(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LMPop(ctx context.Context, direction string, count int64, keys ...string) *KeyValuesCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LPop(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LPopCount(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LPos(ctx context.Context, key string, value string, args LPosArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LPosCount(ctx context.Context, key string, value string, count int64, args LPosArgs) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LPush(ctx context.Context, key string, values ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LPushX(ctx context.Context, key string, values ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LRem(ctx context.Context, key string, count int64, value any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LSet(ctx context.Context, key string, index int64, value any) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) RPop(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) RPopCount(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) RPopLPush(ctx context.Context, source, destination string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) RPush(ctx context.Context, key string, values ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) RPushX(ctx context.Context, key string, values ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SAdd(ctx context.Context, key string, members ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SDiff(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SInter(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SIsMember(ctx context.Context, key string, member any) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SMIsMember(ctx context.Context, key string, members ...any) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SMembers(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SMembersMap(ctx context.Context, key string) *StringStructMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SMove(ctx context.Context, source, destination string, member any) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SPop(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SPopN(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SRandMember(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SRem(ctx context.Context, key string, members ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SUnion(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XAdd(ctx context.Context, a XAddArgs) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XDel(ctx context.Context, stream string, ids ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XLen(ctx context.Context, stream string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XRevRange(ctx context.Context, stream string, start, stop string) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XRead(ctx context.Context, a XReadArgs) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XGroupDestroy(ctx context.Context, stream, group string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XReadGroup(ctx context.Context, a XReadGroupArgs) *XStreamSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XPending(ctx context.Context, stream, group string) *XPendingCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XPendingExt(ctx context.Context, a XPendingExtArgs) *XPendingExtCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XClaim(ctx context.Context, a XClaimArgs) *XMessageSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XClaimJustID(ctx context.Context, a XClaimArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XAutoClaim(ctx context.Context, a XAutoClaimArgs) *XAutoClaimCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XAutoClaimJustID(ctx context.Context, a XAutoClaimArgs) *XAutoClaimJustIDCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XTrimMinID(ctx context.Context, key string, minID string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XInfoStream(ctx context.Context, key string) *XInfoStreamCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XInfoStreamFull(ctx context.Context, key string, count int64) *XInfoStreamFullCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XInfoConsumers(ctx context.Context, key string, group string) *XInfoConsumersCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) XCfgSet(ctx context.Context, a XCfgSetArgs) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BZMPop(ctx context.Context, timeout time.Duration, order string, count int64, keys ...string) *ZSliceWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZAdd(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZAddLT(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZAddGT(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZAddNX(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZAddXX(ctx context.Context, key string, members ...Z) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZCount(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZLexCount(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZInter(ctx context.Context, store ZStore) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZInterWithScores(ctx context.Context, store ZStore) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZInterCard(ctx context.Context, limit int64, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZInterStore(ctx context.Context, destination string, store ZStore) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZMPop(ctx context.Context, order string, count int64, keys ...string) *ZSliceWithKeyCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRank(ctx context.Context, key, member string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRem(ctx context.Context, key string, members ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRemRangeByLex(ctx context.Context, key, min, max string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRevRangeByScore(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRevRangeByLex(ctx context.Context, key string, opt ZRangeBy) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt ZRangeBy) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRevRank(ctx context.Context, key, member string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRevRankWithScore(ctx context.Context, key, member string) *RankWithScoreCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZScore(ctx context.Context, key, member string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZUnionStore(ctx context.Context, dest string, store ZStore) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRandMember(ctx context.Context, key string, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZRandMemberWithScores(ctx context.Context, key string, count int64) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZUnion(ctx context.Context, store ZStore) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZDiff(ctx context.Context, keys ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) PFAdd(ctx context.Context, key string, els ...any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) PFCount(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BgRewriteAOF(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BgSave(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) ClientKill(ctx context.Context, ipPort string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClientKillByFilter(ctx context.Context, keys ...string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClientList(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClientPause(ctx context.Context, dur time.Duration) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClientUnpause(ctx context.Context) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClientID(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) ClientUnblock(ctx context.Context, id int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClientUnblockWithError(ctx context.Context, id int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClientInfo(ctx context.Context) *ClientInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ConfigGet(ctx context.Context, parameter string) *StringStringMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ConfigResetStat(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ConfigSet(ctx context.Context, parameter, value string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ConfigRewrite(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) DBSize(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) FlushAll(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) FlushAllAsync(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FlushDB(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) FlushDBAsync(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Info(ctx context.Context, section ...string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) LastSave(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) Save(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) Shutdown(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) ShutdownSave(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ShutdownNoSave(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SlaveOf(ctx context.Context, host, port string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SlowLogGet(ctx context.Context, num int64) *SlowLogCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SlowLogReset(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Time(ctx context.Context) *TimeCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) DebugObject(ctx context.Context, key string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ReadOnly(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) ReadWrite(ctx context.Context) *StatusCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) MemoryUsage(ctx context.Context, key string, samples ...int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Eval(ctx context.Context, script string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) EvalSha(ctx context.Context, sha1 string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) EvalRO(ctx context.Context, script string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ScriptFlush(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ScriptKill(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ScriptLoad(ctx context.Context, script string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FunctionLoad(ctx context.Context, code string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FunctionLoadReplace(ctx context.Context, code string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FunctionDelete(ctx context.Context, libName string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FunctionFlush(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FunctionKill(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FunctionFlushAsync(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FunctionList(ctx context.Context, q FunctionListQuery) *FunctionListCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FunctionDump(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FunctionRestore(ctx context.Context, libDump string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FCall(ctx context.Context, function string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FCallRO(ctx context.Context, function string, keys []string, args ...any) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Publish(ctx context.Context, channel string, message any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) SPublish(ctx context.Context, channel string, message any) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) PubSubNumSub(ctx context.Context, channels ...string) *StringIntMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) PubSubNumPat(ctx context.Context) *IntCmd { _ = "STUB: not implemented"; return nil }

func (c *Pipeline) PubSubShardChannels(ctx context.Context, pattern string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) PubSubShardNumSub(ctx context.Context, channels ...string) *StringIntMapCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterMyShardID(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterSlots(ctx context.Context) *ClusterSlotsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterShards(ctx context.Context) *ClusterShardsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterNodes(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterLinks(ctx context.Context) *ClusterLinksCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterMeet(ctx context.Context, host string, port int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterForget(ctx context.Context, nodeID string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterReplicate(ctx context.Context, nodeID string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterResetSoft(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterResetHard(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterInfo(ctx context.Context) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterKeySlot(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterGetKeysInSlot(ctx context.Context, slot int64, count int64) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterCountFailureReports(ctx context.Context, nodeID string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterCountKeysInSlot(ctx context.Context, slot int64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterDelSlots(ctx context.Context, slots ...int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterDelSlotsRange(ctx context.Context, min, max int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterSaveConfig(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterSlaves(ctx context.Context, nodeID string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterFailover(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterAddSlots(ctx context.Context, slots ...int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ClusterAddSlotsRange(ctx context.Context, min, max int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GeoAdd(ctx context.Context, key string, geoLocation ...GeoLocation) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query GeoRadiusQuery) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GeoRadiusByMember(ctx context.Context, key, member string, query GeoRadiusQuery) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GeoRadiusByMemberStore(ctx context.Context, key, member string, query GeoRadiusQuery) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GeoSearch(ctx context.Context, key string, q GeoSearchQuery) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GeoSearchLocation(ctx context.Context, key string, q GeoSearchLocationQuery) *GeoLocationCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GeoSearchStore(ctx context.Context, key, store string, q GeoSearchStoreQuery) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GeoDist(ctx context.Context, key string, member1, member2, unit string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FunctionStats(ctx context.Context) *FunctionStatsCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ACLDryRun(ctx context.Context, username string, command ...any) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ACLLog(ctx context.Context, count int64) *ACLLogCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ACLSetUser(ctx context.Context, username string, rules ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ACLDelUser(ctx context.Context, username string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ACLLogReset(ctx context.Context) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ACLCat(ctx context.Context) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ACLList(ctx context.Context) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ACLCatArgs(ctx context.Context, options *ACLCatArgs) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TFunctionLoad(ctx context.Context, lib string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TFunctionLoadArgs(ctx context.Context, lib string, options *TFunctionLoadOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TFunctionDelete(ctx context.Context, libName string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TFunctionList(ctx context.Context) *MapStringInterfaceSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TFunctionListArgs(ctx context.Context, options *TFunctionListOptions) *MapStringInterfaceSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TFCall(ctx context.Context, libName string, funcName string, numKeys int) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TFCallArgs(ctx context.Context, libName string, funcName string, numKeys int, options *TFCallOptions) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TFCallASYNC(ctx context.Context, libName string, funcName string, numKeys int) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TFCallASYNCArgs(ctx context.Context, libName string, funcName string, numKeys int, options *TFCallOptions) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFAdd(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFCard(ctx context.Context, key string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFExists(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFInfo(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFInfoArg(ctx context.Context, key, option string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFInfoCapacity(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFInfoSize(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFInfoFilters(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFInfoItems(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFInfoExpansion(ctx context.Context, key string) *BFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFInsert(ctx context.Context, key string, options *BFInsertOptions, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFMAdd(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFMExists(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFReserve(ctx context.Context, key string, errorRate float64, capacity int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFReserveExpansion(ctx context.Context, key string, errorRate float64, capacity, expansion int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFReserveNonScaling(ctx context.Context, key string, errorRate float64, capacity int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFReserveWithArgs(ctx context.Context, key string, options *BFReserveOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFScanDump(ctx context.Context, key string, iterator int64) *ScanDumpCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) BFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFAdd(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFAddNX(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFCount(ctx context.Context, key string, element interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFDel(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFExists(ctx context.Context, key string, element interface{}) *BoolCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFInfo(ctx context.Context, key string) *CFInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFInsert(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFInsertNX(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFMExists(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFReserve(ctx context.Context, key string, capacity int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFReserveWithArgs(ctx context.Context, key string, options *CFReserveOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFReserveExpansion(ctx context.Context, key string, capacity int64, expansion int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFReserveBucketSize(ctx context.Context, key string, capacity int64, bucketsize int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFReserveMaxIterations(ctx context.Context, key string, capacity int64, maxiterations int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFScanDump(ctx context.Context, key string, iterator int64) *ScanDumpCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CMSIncrBy(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CMSInfo(ctx context.Context, key string) *CMSInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CMSInitByDim(ctx context.Context, key string, width, height int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CMSInitByProb(ctx context.Context, key string, errorRate, probability float64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CMSMerge(ctx context.Context, destKey string, sourceKeys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CMSMergeWithWeight(ctx context.Context, destKey string, sourceKeys map[string]int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) CMSQuery(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TopKAdd(ctx context.Context, key string, elements ...interface{}) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TopKCount(ctx context.Context, key string, elements ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TopKIncrBy(ctx context.Context, key string, elements ...interface{}) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TopKInfo(ctx context.Context, key string) *TopKInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TopKList(ctx context.Context, key string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TopKListWithCount(ctx context.Context, key string) *MapStringIntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TopKQuery(ctx context.Context, key string, elements ...interface{}) *BoolSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TopKReserve(ctx context.Context, key string, k int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TopKReserveWithOptions(ctx context.Context, key string, k int64, width, depth int64, decay float64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestAdd(ctx context.Context, key string, elements ...float64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestByRank(ctx context.Context, key string, rank ...uint64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestByRevRank(ctx context.Context, key string, rank ...uint64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestCDF(ctx context.Context, key string, elements ...float64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestCreate(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestCreateWithCompression(ctx context.Context, key string, compression int64) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestInfo(ctx context.Context, key string) *TDigestInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestMax(ctx context.Context, key string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestMin(ctx context.Context, key string) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestMerge(ctx context.Context, destKey string, options *TDigestMergeOptions, sourceKeys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestQuantile(ctx context.Context, key string, elements ...float64) *FloatSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestRank(ctx context.Context, key string, values ...float64) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestReset(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestRevRank(ctx context.Context, key string, values ...float64) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TDigestTrimmedMean(ctx context.Context, key string, lowCutQuantile, highCutQuantile float64) *FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSAdd(ctx context.Context, key string, timestamp interface{}, value float64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSAddWithArgs(ctx context.Context, key string, timestamp interface{}, value float64, options *TSOptions) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSCreate(ctx context.Context, key string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSCreateWithArgs(ctx context.Context, key string, options *TSOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSAlter(ctx context.Context, key string, options *TSAlterOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSCreateRule(ctx context.Context, sourceKey string, destKey string, aggregator Aggregator, bucketDuration int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSCreateRuleWithArgs(ctx context.Context, sourceKey string, destKey string, aggregator Aggregator, bucketDuration int, options *TSCreateRuleOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSIncrBy(ctx context.Context, Key string, timestamp float64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSIncrByWithArgs(ctx context.Context, key string, timestamp float64, options *TSIncrDecrOptions) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSDecrBy(ctx context.Context, Key string, timestamp float64) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSDecrByWithArgs(ctx context.Context, key string, timestamp float64, options *TSIncrDecrOptions) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSDel(ctx context.Context, Key string, fromTimestamp int, toTimestamp int) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSDeleteRule(ctx context.Context, sourceKey string, destKey string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSGet(ctx context.Context, key string) *TSTimestampValueCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSGetWithArgs(ctx context.Context, key string, options *TSGetOptions) *TSTimestampValueCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSInfo(ctx context.Context, key string) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSInfoWithArgs(ctx context.Context, key string, options *TSInfoOptions) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSMAdd(ctx context.Context, ktvSlices [][]interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSQueryIndex(ctx context.Context, filterExpr []string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSRevRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSRevRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *TSRevRangeOptions) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *TSRangeOptions) *TSTimestampValueSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSMRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSMRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *TSMRangeOptions) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSMRevRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSMRevRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *TSMRevRangeOptions) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSMGet(ctx context.Context, filters []string) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) TSMGetWithArgs(ctx context.Context, filters []string, options *TSMGetOptions) *MapStringSliceInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONArrAppend(ctx context.Context, key, path string, values ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONArrIndex(ctx context.Context, key, path string, value ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONArrIndexWithArgs(ctx context.Context, key, path string, options *JSONArrIndexArgs, value ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONArrInsert(ctx context.Context, key, path string, index int64, values ...interface{}) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONArrLen(ctx context.Context, key, path string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONArrPop(ctx context.Context, key, path string, index int) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONArrTrim(ctx context.Context, key, path string) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONArrTrimWithArgs(ctx context.Context, key, path string, options *JSONArrTrimArgs) *IntSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONClear(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONDebugMemory(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONDel(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONForget(ctx context.Context, key, path string) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONGet(ctx context.Context, key string, paths ...string) *JSONCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONGetWithArgs(ctx context.Context, key string, options *JSONGetArgs, paths ...string) *JSONCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONMerge(ctx context.Context, key, path string, value string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONMSetArgs(ctx context.Context, docs []JSONSetArgs) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONMSet(ctx context.Context, params ...interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONMGet(ctx context.Context, path string, keys ...string) *JSONSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONNumIncrBy(ctx context.Context, key, path string, value float64) *JSONCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONObjKeys(ctx context.Context, key, path string) *SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONObjLen(ctx context.Context, key, path string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONSet(ctx context.Context, key, path string, value interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONSetMode(ctx context.Context, key, path string, value interface{}, mode string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONStrAppend(ctx context.Context, key, path, value string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONStrLen(ctx context.Context, key, path string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONToggle(ctx context.Context, key, path string) *IntPointerSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) JSONType(ctx context.Context, key, path string) *JSONSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FT_List(ctx context.Context) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTAggregate(ctx context.Context, index string, query string) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTAggregateWithArgs(ctx context.Context, index string, query string, options *FTAggregateOptions) *AggregateCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTAliasAdd(ctx context.Context, index string, alias string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTAliasDel(ctx context.Context, alias string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTAliasUpdate(ctx context.Context, index string, alias string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTAlter(ctx context.Context, index string, skipInitialScan bool, definition []interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTConfigGet(ctx context.Context, option string) *MapMapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTConfigSet(ctx context.Context, option string, value interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTCreate(ctx context.Context, index string, options *FTCreateOptions, schema ...*FieldSchema) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTCursorDel(ctx context.Context, index string, cursorId int) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTCursorRead(ctx context.Context, index string, cursorId int, count int) *MapStringInterfaceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTDictAdd(ctx context.Context, dict string, term ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTDictDel(ctx context.Context, dict string, term ...interface{}) *IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTDictDump(ctx context.Context, dict string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTDropIndex(ctx context.Context, index string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTDropIndexWithArgs(ctx context.Context, index string, options *FTDropIndexOptions) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTExplain(ctx context.Context, index string, query string) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTExplainWithArgs(ctx context.Context, index string, query string, options *FTExplainOptions) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTInfo(ctx context.Context, index string) *FTInfoCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTSpellCheck(ctx context.Context, index string, query string) *FTSpellCheckCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTSpellCheckWithArgs(ctx context.Context, index string, query string, options *FTSpellCheckOptions) *FTSpellCheckCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTSearch(ctx context.Context, index string, query string) *FTSearchCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTSearchWithArgs(ctx context.Context, index string, query string, options *FTSearchOptions) *FTSearchCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTSynDump(ctx context.Context, index string) *FTSynDumpCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTSynUpdate(ctx context.Context, index string, synGroupId interface{}, terms []interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTSynUpdateWithArgs(ctx context.Context, index string, synGroupId interface{}, options *FTSynUpdateOptions, terms []interface{}) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) FTTagVals(ctx context.Context, index string, field string) *StringSliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ModuleLoadex(ctx context.Context, conf *ModuleLoadexConfig) *StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Cache(_ time.Duration) CacheCompat {
	_ = "STUB: not implemented"
	return *new(CacheCompat)
}

func (c *Pipeline) Subscribe(_ context.Context, _ ...string) PubSub {
	_ = "STUB: not implemented"
	return *new(PubSub)
}

func (c *Pipeline) PSubscribe(_ context.Context, _ ...string) PubSub {
	_ = "STUB: not implemented"
	return *new(PubSub)
}

func (c *Pipeline) SSubscribe(_ context.Context, _ ...string) PubSub {
	_ = "STUB: not implemented"
	return *new(PubSub)
}

func (c *Pipeline) Watch(_ context.Context, _ func(Tx) error, _ ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) ForEachMaster(_ context.Context, _ func(ctx context.Context, client Cmdable) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Pipeline) Client() rueidis.Client { _ = "STUB: not implemented"; return *new(rueidis.Client) }

// Len returns the number of queued commands.
func (c *Pipeline) Len() int { _ = "STUB: not implemented"; return 0 }

// Do queues the custom command for later execution.
func (c *Pipeline) Do(_ context.Context, args ...interface{}) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

// Discard resets the pipeline and discards queued commands.
func (c *Pipeline) Discard() { _ = "STUB: not implemented"; return }

// Exec executes all previously queued commands using one
// client-server roundtrip.
//
// Exec always returns a list of commands and error of the first failed
//
//	command, if any.
func (c *Pipeline) Exec(ctx context.Context) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Pipeline) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Pipeline) Pipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *Pipeline) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Pipeline) TxPipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }
