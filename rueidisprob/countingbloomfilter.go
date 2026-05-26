package rueidisprob

import (
	"context"
	"errors"

	"github.com/redis/rueidis"
)

var (
	ErrEmptyCountingBloomFilterName                          = errors.New("name cannot be empty")
	ErrCountingBloomFilterFalsePositiveRateLessThanEqualZero = errors.New("false positive rate cannot be less than or equal to zero")
	ErrCountingBloomFilterFalsePositiveRateGreaterThanOne    = errors.New("false positive rate cannot be greater than 1")
	ErrCountingBloomFilterBitsSizeZero                       = errors.New("bits size cannot be zero")
)

const (
	countingBloomFilterAddMultiScript = `
local itemCount = tonumber(ARGV[1])
local numElements = tonumber(#ARGV) - 1
local filterKey = KEYS[1]
local counterKey = KEYS[2]

for i=2, numElements+1 do
    redis.call('HINCRBY', filterKey, ARGV[i], 1)
end

return redis.call('INCRBY', counterKey, itemCount)
`

	countingBloomFilterRemoveMultiScript = `
local function MergeTables(t1, t2)
	for i=1, #t2 do
		table.insert(t1, t2[i])
	end

	return t1
end

local numElements = tonumber(#ARGV) - 1
local hashIterations = tonumber(ARGV[#ARGV])
local filterKey = KEYS[1]
local counterKey = KEYS[2]

local indexCounter = {}
for i=1, numElements do
	local index = ARGV[i]
	local count = redis.call('HGET', filterKey, index)

	if (not indexCounter[index]) then
		if (not count) then
			indexCounter[index] = 0
		else
			indexCounter[index] = tonumber(count)
		end
	end
end

local decreaseIndexes = {}
local deleteItemCount = 0
for i=1, numElements, hashIterations do
	local isAbleToRemove = true
	local temp = {}
	local rollbackIndex = i

	for j=i, i+hashIterations-1 do
		local index = ARGV[j]

		table.insert(temp, index)
		indexCounter[index] = indexCounter[index] - 1
		
		if indexCounter[index] < 0 then
			isAbleToRemove = false
			rollbackIndex = j
			break
		end
	end

	if isAbleToRemove then
		decreaseIndexes = MergeTables(decreaseIndexes, temp)
		deleteItemCount = deleteItemCount + 1
	else
		for j=i, rollbackIndex do
			local index = ARGV[j]
			
			indexCounter[index] = indexCounter[index] + 1
		end
	end
end

for i=1, #decreaseIndexes do
    redis.call('HINCRBY', filterKey, decreaseIndexes[i], -1)
end

return redis.call('DECRBY', counterKey, deleteItemCount)
`

	countingBloomFilterDeleteScript = `
local filterKey = KEYS[1]
local counterKey = KEYS[2]

redis.call('DEL', filterKey)
redis.call('DEL', counterKey)

return 1
`
)

// CountingBloomFilter based on Hashes.
// CountingBloomFilter uses a 128-bit murmur3 hash function.
type CountingBloomFilter interface {
	// Add adds an item to the Counting Bloom Filter.
	Add(ctx context.Context, key string) error

	// AddMulti adds one or more items to the Counting Bloom Filter.
	// NOTE: If keys are too many, it can block the Redis server for a long time.
	AddMulti(ctx context.Context, keys []string) error

	// Exists checks if an item is in the Counting Bloom Filter.
	Exists(ctx context.Context, key string) (bool, error)

	// ExistsMulti checks if one or more items are in the Counting Bloom Filter.
	// Returns a slice of bool values where each bool indicates
	// whether the corresponding key was found.
	ExistsMulti(ctx context.Context, keys []string) ([]bool, error)

	// Remove removes an item from the Counting Bloom Filter.
	Remove(ctx context.Context, key string) error

	// RemoveMulti removes one or more items from the Counting Bloom Filter.
	// NOTE: If keys are too many, it can block the Redis server for a long time.
	RemoveMulti(ctx context.Context, keys []string) error

	// Delete deletes the Counting Bloom Filter.
	Delete(ctx context.Context) error

	// ItemMinCount returns the minimum count of item in the Counting Bloom Filter.
	// If the item is not in the Counting Bloom Filter, it returns a zero value.
	// A minimum count is not always accurate because of the hash collisions.
	ItemMinCount(ctx context.Context, key string) (uint64, error)

	// ItemMinCountMulti returns the minimum count of items in the Counting Bloom Filter.
	// If the item is not in the Counting Bloom Filter, it returns a zero value.
	// A minimum count is not always accurate because of the hash collisions.
	ItemMinCountMulti(ctx context.Context, keys []string) ([]uint64, error)

	// Count returns count of items in Counting Bloom Filter.
	Count(ctx context.Context) (uint64, error)
}

type countingBloomFilter struct {
	client rueidis.Client

	addMultiScript *rueidis.Lua

	removeMultiScript *rueidis.Lua

	// name is the name of the Counting Bloom Filter.
	// It is used as a key in the Redis.
	name string

	// counter is the name of the counter.
	counter string

	hashIterationString string

	addMultiKeys []string

	removeMultiKeys []string

	// hashIterations is the number of hash functions to use.
	hashIterations uint

	// size is the number of bits to use.
	size uint
}

// NewCountingBloomFilter creates a new Counting Bloom Filter.
// NOTE: 'name:cbf:c' is used as a counter-key in the Redis and
// 'name:cbf' is used as a filter key in the Redis
// to keep track of the number of items in the Counting Bloom Filter for Count method.
func NewCountingBloomFilter(
	client rueidis.Client,
	name string,
	expectedNumberOfItems uint,
	falsePositiveRate float64,
) (CountingBloomFilter, error) {
	_ = "STUB: not implemented"
	return *new(CountingBloomFilter), nil
}

// NOTE: https://redis.io/docs/reference/cluster-spec/#hash-tags

func (f *countingBloomFilter) Add(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *countingBloomFilter) AddMulti(ctx context.Context, keys []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *countingBloomFilter) indexes(keys []string, buf *[]byte) []string {
	_ = "STUB: not implemented"
	return nil
}

func (f *countingBloomFilter) Exists(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *countingBloomFilter) ExistsMulti(ctx context.Context, keys []string) ([]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *countingBloomFilter) Remove(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *countingBloomFilter) RemoveMulti(ctx context.Context, keys []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *countingBloomFilter) Delete(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *countingBloomFilter) ItemMinCount(ctx context.Context, key string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *countingBloomFilter) ItemMinCountMulti(ctx context.Context, keys []string) ([]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *countingBloomFilter) Count(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
