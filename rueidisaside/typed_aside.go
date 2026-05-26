package rueidisaside

import (
	"context"
	"time"
)

// TypedCacheAsideClient is an interface that provides a typed cache-aside client.
// It allows you to cache and retrieve values of a specific type T.
type TypedCacheAsideClient[T any] interface {
	Get(ctx context.Context, ttl time.Duration, key string, fn func(ctx context.Context, key string) (val *T, err error)) (val *T, err error)
	Del(ctx context.Context, key string) error
	Client() CacheAsideClient
}

// typedCacheAsideClient is an implementation of the TypedCacheAsideClient interface.
// It provides a typed cache-aside client that allows caching and retrieving values of a specific type T.
type typedCacheAsideClient[T any] struct {
	client       CacheAsideClient
	serializer   func(*T) (string, error)
	deserializer func(string) (*T, error)
}

// NewTypedCacheAsideClient creates a new TypedCacheAsideClient instance that provides a typed cache-aside client.
// The client, serializer, and deserializer functions are used to interact with the underlying cache.
// The serializer function is used to convert the provided value of type T to a string, and the deserializer function
// is used to convert the cached string value back to the original type T.
func NewTypedCacheAsideClient[T any](
	client CacheAsideClient,
	serializer func(*T) (string, error),
	deserializer func(string) (*T, error),
) TypedCacheAsideClient[T] {
	_ = "STUB: not implemented"
	return nil
}

// Get retrieves a value of type T from the cache or fetches it using the provided
// function and stores it in the cache. The value is cached for the specified TTL.
// If the value cannot be retrieved or deserialized, an error is returned.
func (c typedCacheAsideClient[T]) Get(ctx context.Context, ttl time.Duration, key string, fn func(ctx context.Context, key string) (val *T, err error)) (val *T, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Del deletes the value associated with the given key from the cache.
func (c typedCacheAsideClient[T]) Del(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// Client returns the underlying CacheAsideClient instance used by the TypedCacheAsideClient.
func (c typedCacheAsideClient[T]) Client() CacheAsideClient {
	_ = "STUB: not implemented"
	return *new(CacheAsideClient)
}
