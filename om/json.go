package om

import (
	"context"
	"reflect"
	"time"

	"github.com/redis/rueidis"
)

// NewJSONRepository creates a JSONRepository.
// The prefix parameter is used as redis key prefix. The entity stored by the repository will be named in the form of `{prefix}:{id}`
// The schema parameter should be a struct with fields tagged with `redis:",key"`. The `redis:",ver"` tag is optional for optimistic locking.
func NewJSONRepository[T any](prefix string, schema T, client rueidis.Client, opts ...RepositoryOption) Repository[T] {
	_ = "STUB: not implemented"
	return nil
}

var _ Repository[any] = (*JSONRepository[any])(nil)

// JSONRepository is an OM repository backed by RedisJSON.
type JSONRepository[T any] struct {
	schema schema
	typ    reflect.Type
	client rueidis.Client
	prefix string
	idx    string
}

// NewEntity returns an empty entity and will have the `redis:",key"` field be set with ULID automatically.
func (r *JSONRepository[T]) NewEntity() *T { _ = "STUB: not implemented"; return nil }

// Fetch an entity whose name is `{prefix}:{id}`
func (r *JSONRepository[T]) Fetch(ctx context.Context, id string) (v *T, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchCache is like Fetch, but it uses the client side caching mechanism.
func (r *JSONRepository[T]) FetchCache(ctx context.Context, id string, ttl time.Duration) (v *T, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *JSONRepository[T]) decode(record string) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *JSONRepository[T]) toExec(entity *T) (verf reflect.Value, exec rueidis.LuaExec) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), *new(rueidis.LuaExec)
}

// verless, set verf to a dummy value

// Save the entity under the redis key of `{prefix}:{id}`.
// If the entity has a `redis:",ver"` field, it uses optimistic locking to prevent lost updates.
func (r *JSONRepository[T]) Save(ctx context.Context, entity *T) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// SaveMulti batches multiple HashRepository.Save at once
func (r *JSONRepository[T]) SaveMulti(ctx context.Context, entities ...*T) []error {
	_ = "STUB: not implemented"
	return nil
}

// Remove the entity under the redis key of `{prefix}:{id}`.
func (r *JSONRepository[T]) Remove(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// AlterIndex uses FT.ALTER from the RediSearch module to alter index under the name `jsonidx:{prefix}`
// You can use the cmdFn parameter to mutate the index alter command.
func (r *JSONRepository[T]) AlterIndex(ctx context.Context, cmdFn func(alter FtAlterIndex) rueidis.Completed) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateIndex uses FT.CREATE from the RediSearch module to create an inverted index under the name `jsonidx:{prefix}`
// You can use the cmdFn parameter to mutate the index construction command,
// and note that the field name should be specified with JSON path syntax; otherwise, the index may not work as expected.
func (r *JSONRepository[T]) CreateIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) rueidis.Completed) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateAndAliasIndex creates a new index, aliases it, and drops the old index if needed.
func (r *JSONRepository[T]) CreateAndAliasIndex(ctx context.Context, cmdFn func(schema FtCreateSchema) rueidis.Completed) error {
	_ = "STUB: not implemented"
	return nil
}

// DropIndex uses FT.DROPINDEX from the RediSearch module to drop the index whose name is `jsonidx:{prefix}`
func (r *JSONRepository[T]) DropIndex(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Search uses FT.SEARCH from the RediSearch module to search the index whose name is `jsonidx:{prefix}`
// It returns three values:
// 1. total count of match results inside the redis, and note that it might be larger than the returned search result.
// 2. the search result, and note that its length might be smaller than the first return value.
// 3. error if any
// You can use the cmdFn parameter to mutate the search command.
func (r *JSONRepository[T]) Search(ctx context.Context, cmdFn func(search FtSearchIndex) rueidis.Completed) (n int64, s []*T, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// supports dialect 3

// Aggregate performs the FT.AGGREGATE and returns a *AggregateCursor for accessing the results
func (r *JSONRepository[T]) Aggregate(ctx context.Context, cmdFn func(agg FtAggregateIndex) rueidis.Completed) (cursor *AggregateCursor, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IndexName returns the index name used in the FT.CREATE
func (r *JSONRepository[T]) IndexName() string { _ = "STUB: not implemented"; return "" }

var jsonSaveScript = rueidis.NewLuaScript(`
if (ARGV[1] == '')
then
  redis.call('JSON.SET',KEYS[1],'$',ARGV[3])
  if #ARGV == 4 then redis.call('PEXPIREAT',KEYS[1],ARGV[4]) end
  return ARGV[2]
end
local v = redis.call('JSON.GET',KEYS[1],ARGV[1])
if (not v or v == ARGV[2])
then
  redis.call('JSON.SET',KEYS[1],'$',ARGV[3])
  local v = redis.call('JSON.NUMINCRBY',KEYS[1],ARGV[1],1)
  if #ARGV == 4 then redis.call('PEXPIREAT',KEYS[1],ARGV[4]) end
  return v
end
return nil
`)
