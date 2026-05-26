package rueidis

import (
	"container/list"
	"context"
	"sync"
	"time"
	"unsafe"
)

const (
	entrySize    = int(unsafe.Sizeof(cacheEntry{})) + int(unsafe.Sizeof(&cacheEntry{}))
	keyCacheSize = int(unsafe.Sizeof(keyCache{})) + int(unsafe.Sizeof(&keyCache{}))
	elementSize  = int(unsafe.Sizeof(list.Element{})) + int(unsafe.Sizeof(&list.Element{}))
	stringSSize  = int(unsafe.Sizeof(""))

	entryBaseSize = (keyCacheSize + entrySize + elementSize + stringSSize*2) * 3 / 2
	entryMinSize  = entryBaseSize + messageStructSize

	moveThreshold = uint32(1024 - 1)
)

type cacheEntry struct {
	err  error
	ch   chan struct{}
	kc   *keyCache
	cmd  string
	val  RedisMessage
	size int
}

func (e *cacheEntry) Wait(ctx context.Context) (RedisMessage, error) {
	_ = "STUB: not implemented"
	return *new(RedisMessage), nil
}

type keyCache struct {
	cache map[string]*list.Element
	key   string
	hits  uint32
	miss  uint32
}

var _ CacheStore = (*lru)(nil)

type lru struct {
	store map[string]*keyCache
	list  *list.List
	mu    sync.RWMutex
	size  int
	max   int
}

func newLRU(opt CacheStoreOption) CacheStore { _ = "STUB: not implemented"; return *new(CacheStore) }

func (c *lru) Flight(key, cmd string, ttl time.Duration, now time.Time) (v RedisMessage, ce CacheEntry) {
	_ = "STUB: not implemented"
	return *new(RedisMessage), *new(CacheEntry)
}

func (c *lru) Flights(now time.Time, multi []CacheableTTL, results []RedisResult, entries map[int]CacheEntry) (missed []int) {
	_ = "STUB: not implemented"
	return nil
}

func (c *lru) Update(key, cmd string, value RedisMessage) (pxat int64) {
	_ = "STUB: not implemented"
	return 0
}

// server side ttl should only shorten client side ttl

// do not delete pending entries

func (c *lru) Cancel(key, cmd string, err error) { _ = "STUB: not implemented"; return }

func (c *lru) GetTTL(key, cmd string) (ttl time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *lru) purge(key string, kc *keyCache) { _ = "STUB: not implemented"; return }

// do not delete pending entries

func (c *lru) Delete(keys []RedisMessage) { _ = "STUB: not implemented"; return }

func (c *lru) Close(err error) { _ = "STUB: not implemented"; return }
