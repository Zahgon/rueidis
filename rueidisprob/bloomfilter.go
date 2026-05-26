package rueidisprob

import (
	"context"
	"errors"

	"github.com/redis/rueidis"
)

const (
	// NOTE: https://redis.io/docs/data-types/bitmaps/
	maxSize = 1 << 32
)

const (
	bloomFilterAddMultiScript = `
local hashIterations = tonumber(ARGV[1])
local numElements = tonumber(#ARGV) - 1
local filterKey = KEYS[1]
local counterKey = KEYS[2]

local counter = 0
local oneBits = 0
for i=1, numElements do
	local bitset = redis.call('BITFIELD', filterKey, 'SET', 'u1', ARGV[i+1], '1')

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		if oneBits ~= hashIterations then
			counter = counter + 1
		end

		oneBits = 0
	end
end

return redis.call('INCRBY', counterKey, counter)
`

	bloomFilterExistsMultiScript = `
local hashIterations = tonumber(ARGV[1])
local numElements = tonumber(#ARGV) - 1
local filterKey = KEYS[1]

local result = {}
local oneBits = 0
for i=1, numElements do
	local index = tonumber(ARGV[i+1])
	local bitset = redis.call('BITFIELD', filterKey, 'GET', 'u1', index)

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		table.insert(result, oneBits == hashIterations)

		oneBits = 0
	end
end

return result
`

	bloomFilterExistsMultiReadOnlyScript = `
local hashIterations = tonumber(ARGV[1])
local numElements = tonumber(#ARGV) - 1
local filterKey = KEYS[1]

local result = {}
local oneBits = 0
for i=1, numElements do
	local index = tonumber(ARGV[i+1])
	local bitset = redis.call('BITFIELD_RO', filterKey, 'GET', 'u1', index)

	oneBits = oneBits + bitset[1]
	if i % hashIterations == 0 then
		table.insert(result, oneBits == hashIterations)

		oneBits = 0
	end
end

return result
`

	bloomFilterResetScript = `
local filterKey = KEYS[1]
local counterKey = KEYS[2]

redis.call('SET', filterKey, "")
redis.call('SET', counterKey, 0)

return 1
`

	bloomFilterDeleteScript = `
local filterKey = KEYS[1]
local counterKey = KEYS[2]

redis.call('DEL', filterKey)
redis.call('DEL', counterKey)

return 1
`
)

var (
	ErrEmptyName                          = errors.New("name cannot be empty")
	ErrFalsePositiveRateLessThanEqualZero = errors.New("false positive rate cannot be less than or equal to zero")
	ErrFalsePositiveRateGreaterThanOne    = errors.New("false positive rate cannot be greater than 1")
	ErrBitsSizeZero                       = errors.New("bits size cannot be zero")
	ErrBitsSizeTooLarge                   = errors.New("bits size is too large")
)

// BloomFilterOptions is used to configure BloomFilter.
type BloomFilterOptions struct {
	enableReadOperation bool
}

// BloomFilterOptionFunc is used to configure BloomFilter.
type BloomFilterOptionFunc func(*BloomFilterOptions)

// WithEnableReadOperation enables read operation.
// If enabled, Exists and ExistsMulti methods will be available as read-only operations.
// NOTE: If enabled, minimum redis version should be 7.0.0.
func WithEnableReadOperation(enableReadOperations bool) BloomFilterOptionFunc {
	_ = "STUB: not implemented"
	return *new(BloomFilterOptionFunc)
}

// BloomFilter based on Redis Bitmaps.
// BloomFilter uses a 128-bit murmur3 hash function.
type BloomFilter interface {
	// Add adds an item to the Bloom filter.
	Add(ctx context.Context, key string) error

	// AddMulti adds one or more items to the Bloom filter.
	// NOTE: If keys are too many, it can block the Redis server for a long time.
	AddMulti(ctx context.Context, keys []string) error

	// Exists checks if an item is in the Bloom filter.
	Exists(ctx context.Context, key string) (bool, error)

	// ExistsMulti checks if one or more items are in the Bloom filter.
	// Returns a slice of bool values where each bool indicates whether the corresponding key was found.
	ExistsMulti(ctx context.Context, keys []string) ([]bool, error)

	// Reset resets the Bloom filter.
	Reset(ctx context.Context) error

	// Delete deletes the Bloom filter.
	Delete(ctx context.Context) error

	// Count returns count of items in Bloom filter.
	Count(ctx context.Context) (uint64, error)
}

type bloomFilter struct {
	client rueidis.Client

	addMultiScript *rueidis.Lua

	existsMultiScript *rueidis.Lua

	// name is the name of the Bloom filter.
	// It is used as a key in the Redis.
	name string

	// counter is the name of the counter.
	counter string

	hashIterationString string

	addMultiKeys []string

	existsMultiKeys []string

	// hashIterations is the number of hash functions to use.
	hashIterations uint

	// size is the number of bits to use.
	size uint
}

// NewBloomFilter creates a new Bloom filter.
// NOTE: 'name:c' is used as a counter-key in the Redis
// to keep track of the number of items in the Bloom filter for Count method.
func NewBloomFilter(
	client rueidis.Client,
	name string,
	expectedNumberOfItems uint,
	falsePositiveRate float64,
	opts ...BloomFilterOptionFunc,
) (BloomFilter, error) {
	_ = "STUB: not implemented"
	return *new(BloomFilter), nil
}

// NOTE: https://redis.io/docs/reference/cluster-spec/#hash-tags

func numberOfBloomFilterBits(n uint, r float64) uint { _ = "STUB: not implemented"; return 0 }

func numberOfBloomFilterHashFunctions(s uint, n uint) uint { _ = "STUB: not implemented"; return 0 }

func (c *bloomFilter) Add(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *bloomFilter) AddMulti(ctx context.Context, keys []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *bloomFilter) indexes(keys []string, buf *[]byte) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *bloomFilter) Exists(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *bloomFilter) ExistsMulti(ctx context.Context, keys []string) ([]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *bloomFilter) Reset(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *bloomFilter) Delete(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *bloomFilter) Count(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
