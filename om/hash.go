package om

import (
	"context"
	"reflect"
	"time"

	"github.com/redis/rueidis"
)

// NewHashRepository creates a HashRepository.
// The prefix parameter is used as redis key prefix. The entity stored by the repository will be named in the form of `{prefix}:{id}`
// The schema parameter should be a struct with fields tagged with `redis:",key"`. The `redis:",ver"` tag is optional for optimistic locking.
func NewHashRepository[T any](prefix string, schema T, client rueidis.Client, opts ...RepositoryOption) Repository[T] {
	_ = "STUB: not implemented"
	return nil
}

var _ Repository[any] = (*HashRepository[any])(nil)

// HashRepository is an OM repository backed by redis hash.
type HashRepository[T any] struct {
	schema  schema
	typ     reflect.Type
	client  rueidis.Client
	factory *hashConvFactory
	prefix  string
	idx     string
}

// NewEntity returns an empty entity and will have the `redis:",key"` field be set with ULID automatically.
func (r *HashRepository[T]) NewEntity() (entity *T) { _ = "STUB: not implemented"; return nil }

// Fetch an entity whose name is `{prefix}:{id}`
func (r *HashRepository[T]) Fetch(ctx context.Context, id string) (v *T, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchCache is like Fetch, but it uses the client side caching mechanism.
func (r *HashRepository[T]) FetchCache(ctx context.Context, id string, ttl time.Duration) (v *T, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *HashRepository[T]) toExec(entity *T) (verf reflect.Value, exec rueidis.LuaExec) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(rueidis.LuaExec)
}

// verless, set verf to a dummy value

// keep the ver field be the first pair for the hashSaveScript

// Save the entity under the redis key of `{prefix}:{id}`.
// If the entity has a `redis:",ver"` field, it uses optimistic locking to prevent lost updates.
func (r *HashRepository[T]) Save(ctx context.Context, entity *T) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// SaveMulti batches multiple HashRepository.Save at once
func (r *HashRepository[T]) SaveMulti(ctx context.Context, entities ...*T) []error {
	_ = "STUB: not implemented"
	return nil
}

// Remove the entity under the redis key of `{prefix}:{id}`.
func (r *HashRepository[T]) Remove(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// AlterIndex uses FT.ALTER from the RediSearch module to alter index under the name `hashidx:{prefix}`
// You can use the cmdFn parameter to mutate the index alter command.
func (r *HashRepository[T]) AlterIndex(ctx context.Context, cmdFn func(alter FtAlterIndex) rueidis.Completed) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateIndex uses FT.CREATE from the RediSearch module to create an inverted index under the name `hashidx:{prefix}`
// You can use the cmdFn parameter to mutate the index construction command.
func (r *HashRepository[T]) CreateIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) rueidis.Completed) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateAndAliasIndex creates a new index, aliases it, and drops the old index if needed.
func (r *HashRepository[T]) CreateAndAliasIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) rueidis.Completed) error {
	_ = "STUB: not implemented"
	return nil
}

// DropIndex uses FT.DROPINDEX from the RediSearch module to drop the index whose name is `hashidx:{prefix}`
func (r *HashRepository[T]) DropIndex(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Search uses FT.SEARCH from the RediSearch module to search the index whose name is `hashidx:{prefix}`
// It returns three values:
// 1. total count of match results inside the redis, and note that it might be larger than the returned search result.
// 2. the search result, and note that its length might be smaller than the first return value.
// 3. error if any
// You can use the cmdFn parameter to mutate the search command.
func (r *HashRepository[T]) Search(ctx context.Context, cmdFn func(search FtSearchIndex) rueidis.Completed) (n int64, s []*T, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// Aggregate performs the FT.AGGREGATE and returns a *AggregateCursor for accessing the results
func (r *HashRepository[T]) Aggregate(ctx context.Context, cmdFn func(agg FtAggregateIndex) rueidis.Completed) (cursor *AggregateCursor, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IndexName returns the index name used in the FT.CREATE
func (r *HashRepository[T]) IndexName() string { _ = "STUB: not implemented"; return "" }

func (r *HashRepository[T]) fromHash(record map[string]string) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *HashRepository[T]) fromFields(fields map[string]string) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var hashSaveScript = rueidis.NewLuaScript(`
if (ARGV[1] == '')
then
  local e = (#ARGV % 2 == 1) and table.remove(ARGV) or nil
  if redis.call('HSET',KEYS[1],unpack(ARGV))
  then
    if e then redis.call('PEXPIREAT',KEYS[1],e) end
  end
  return ARGV[2]
end
local v = redis.call('HGET',KEYS[1],ARGV[1])
if (not v or v == ARGV[2])
then
  ARGV[2] = tostring(tonumber(ARGV[2])+1)
  local e = (#ARGV % 2 == 1) and table.remove(ARGV) or nil
  if redis.call('HSET',KEYS[1],unpack(ARGV))
  then
    if e then redis.call('PEXPIREAT',KEYS[1],e) end
    return ARGV[2]
  end
end
return nil
`)
