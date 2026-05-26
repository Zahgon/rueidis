package rueidisprob

import (
	"context"
	"errors"
	"time"

	"github.com/redis/rueidis"
)

const (
	slidingBloomFilterInitializeScript = `
local filterKey = KEYS[1]
local nextFilterKey = KEYS[2]
local counterKey = KEYS[3]
local nextCounterKey = KEYS[4]
local lastRotationKey = KEYS[5]
local windowHalf = tonumber(ARGV[1])

if redis.call('EXISTS', filterKey, nextFilterKey, counterKey, nextCounterKey, lastRotationKey) == 0 then
	local time = redis.call('TIME')
	local current_time = tonumber(time[1]) * 1000 + math.floor(tonumber(time[2]) / 1000)

	redis.call('MSET', filterKey, "", counterKey, 0, nextFilterKey, "", nextCounterKey, 0)
	redis.call('SET', lastRotationKey, tostring(current_time), 'PX', windowHalf, 'NX')
end

return 1
`

	slidingBloomFilterAddMultiScript = `
local hashIterations = tonumber(ARGV[1])
local windowHalf = tonumber(ARGV[2])
local numElements = tonumber(#ARGV) - 2

local filterKey = KEYS[1]
local nextFilterKey = KEYS[2]
local counterKey = KEYS[3]
local nextCounterKey = KEYS[4]
local lastRotationKey = KEYS[5]

local time = redis.call('TIME')
local current_time = tonumber(time[1]) * 1000 + math.floor(tonumber(time[2])/1000)
local acquiredLock = redis.call('SET', lastRotationKey, tostring(current_time), 'PX', windowHalf, 'NX')

if acquiredLock then
	redis.call('RENAME', nextFilterKey, filterKey)
	redis.call('RENAME', nextCounterKey, counterKey)
	redis.call('SET', nextFilterKey, "")
	redis.call('SET', nextCounterKey, 0)
end

local counter = 0
local oneBits = 0
for i=1, numElements do
	local bitset = redis.call('BITFIELD', filterKey, 'SET', 'u1', ARGV[i+2], '1')
	redis.call('BITFIELD', nextFilterKey, 'SET', 'u1', ARGV[i+2], '1')

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		if oneBits ~= hashIterations then
			counter = counter + 1
		end

		oneBits = 0
	end
end

redis.call('INCRBY', nextCounterKey, counter)
return redis.call('INCRBY', counterKey, counter)
`
	slidingBloomFilterExistsMultiScript = `
local hashIterations = tonumber(ARGV[1])
local windowHalf = tonumber(ARGV[2])
local numElements = tonumber(#ARGV) - 2

local filterKey = KEYS[1]
local nextFilterKey = KEYS[2]
local counterKey = KEYS[3]
local nextCounterKey = KEYS[4]
local lastRotationKey = KEYS[5]

local time = redis.call('TIME')
local current_time = tonumber(time[1]) * 1000 + math.floor(tonumber(time[2])/1000)
local acquiredLock = redis.call('SET', lastRotationKey, tostring(current_time), 'PX', windowHalf, 'NX')

if acquiredLock then
	redis.call('RENAME', nextFilterKey, filterKey)
	redis.call('RENAME', nextCounterKey, counterKey)
	redis.call('SET', nextFilterKey, "")
	redis.call('SET', nextCounterKey, 0)
end

local result = {}
local oneBits = 0
for i=1, numElements do
	local index = tonumber(ARGV[i+2])
	local bitset = redis.call('BITFIELD', filterKey, 'GET', 'u1', index)

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		table.insert(result, oneBits == hashIterations)

		oneBits = 0
	end
end

return result
`

	slidingBloomFilterExistsReadOnlyMultiScript = `
local hashIterations = tonumber(ARGV[1])
local windowHalf = tonumber(ARGV[2])
local numElements = tonumber(#ARGV) - 2

local filterKey = KEYS[1]
local nextFilterKey = KEYS[2]
local counterKey = KEYS[3]
local nextCounterKey = KEYS[4]
local lastRotationKey = KEYS[5]

local time = redis.call('TIME')
local current_time = tonumber(time[1]) * 1000 + math.floor(tonumber(time[2])/1000)
local acquiredLock = redis.call('SET', lastRotationKey, tostring(current_time), 'PX', windowHalf, 'NX')

if acquiredLock then
	redis.call('RENAME', nextFilterKey, filterKey)
	redis.call('RENAME', nextCounterKey, counterKey)
	redis.call('SET', nextFilterKey, "")
	redis.call('SET', nextCounterKey, 0)
end

local result = {}
local oneBits = 0
for i=1, numElements do
	local index = tonumber(ARGV[i+2])
	local bitset = redis.call('BITFIELD_RO', filterKey, 'GET', 'u1', index)

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		table.insert(result, oneBits == hashIterations)

		oneBits = 0
	end
end

return result
`

	slidingBloomFilterResetScript = `
local filterKey = KEYS[1]
local nextFilterKey = KEYS[2]
local counterKey = KEYS[3]
local nextCounterKey = KEYS[4]

redis.call('RENAME', nextFilterKey, filterKey)
redis.call('RENAME', nextCounterKey, counterKey)
redis.call('SET', nextFilterKey, "")
redis.call('SET', nextCounterKey, 0)
`

	// Redis key suffixes
	counterSuffix      = ":c"
	nextFilterSuffix   = ":n"
	nextCounterSuffix  = ":nc"
	lastRotationSuffix = ":lr"
)

var (
	_                              BloomFilter = (*slidingBloomFilter)(nil)
	ErrWindowSizeLessThanOneSecond             = errors.New("window size cannot be less than 1 second")
)

type SlidingBloomFilterOptions struct {
	enableReadOperation bool
}

type SlidingBloomFilterOptionFunc func(o *SlidingBloomFilterOptions)

func WithReadOnlyExists(enableReadOperations bool) SlidingBloomFilterOptionFunc {
	_ = "STUB: not implemented"
	return *new(SlidingBloomFilterOptionFunc)
}

type slidingBloomFilter struct {
	client rueidis.Client

	addMultiScript *rueidis.Lua

	existsMultiScript *rueidis.Lua

	// Pre-calculated window half in milliseconds
	windowHalfMs string

	// name is the name of the sliding Bloom filter.
	// It is used as a key in the Redis.
	name string

	// counter is the name of the counter.
	counter string

	hashIterationString string

	addMultiKeys []string

	existsMultiKeys []string

	// window is the duration of the sliding window.
	window time.Duration

	// hashIterations is the number of hash functions to use.
	hashIterations uint

	// size is the number of bits to use.
	size uint
}

// NewSlidingBloomFilter creates a new sliding window Bloom filter.
// NOTE: 'name:c' is used as a counter-key in the Redis
// 'name:n' is used as a next filter key in the Redis
// 'name:nc' is used as a next counter key in the Redis
// 'name:lr' is used as a last rotation key in the Redis
// to keep track of the items in the window.
func NewSlidingBloomFilter(
	redisClient rueidis.Client,
	name string,
	expectedNumberOfItems uint,
	falsePositiveRate float64,
	windowSize time.Duration,
	opts ...SlidingBloomFilterOptionFunc,
) (BloomFilter, error) {
	_ = "STUB: not implemented"
	return *new(BloomFilter), nil
}

// NOTE: https://redis.io/docs/reference/cluster-spec/#hash-tags

func (s *slidingBloomFilter) initialize() error { _ = "STUB: not implemented"; return nil }

func (s *slidingBloomFilter) Add(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *slidingBloomFilter) AddMulti(ctx context.Context, keys []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *slidingBloomFilter) indexes(keys []string, buf *[]byte) []string {
	_ = "STUB: not implemented"
	return nil
}

func (s *slidingBloomFilter) Exists(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *slidingBloomFilter) ExistsMulti(ctx context.Context, keys []string) ([]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *slidingBloomFilter) Reset(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *slidingBloomFilter) Delete(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *slidingBloomFilter) Count(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
