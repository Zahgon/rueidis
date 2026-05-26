package rueidis

import (
	"time"

	"github.com/redis/rueidis/internal/util"
)

var (
	resultsp = util.NewPool(func(capacity int) *redisresults {
		return &redisresults{s: make([]RedisResult, 0, capacity)}
	})
	mgetcmdsp = util.NewPool(func(capacity int) *mgetcmds {
		return &mgetcmds{s: make([]Completed, 0, capacity)}
	})
	retryp = util.NewPool(func(capacity int) *retry {
		return &retry{
			cIndexes: make([]int, 0, capacity),
			commands: make([]Completed, 0, capacity),
		}
	})
	mgetcachecmdsp = util.NewPool(func(capacity int) *mgetcachecmds {
		return &mgetcachecmds{s: make([]CacheableTTL, 0, capacity)}
	})
	retrycachep = util.NewPool(func(capacity int) *retrycache {
		return &retrycache{
			cIndexes: make([]int, 0, capacity),
			commands: make([]CacheableTTL, 0, capacity),
		}
	})
	batchcachep = util.NewPool(func(capacity int) *batchcache {
		return &batchcache{
			cIndexes: make([]int, 0, capacity),
			commands: make([]CacheableTTL, 0, capacity),
		}
	})
	batchcachemaps = util.NewPool(func(capacity int) *batchcachemap {
		return &batchcachemap{m: make(map[uint16]*batchcache, capacity), n: capacity}
	})
	muxslotsp = util.NewPool(func(capacity int) *muxslots {
		return &muxslots{s: make([]int, 0, capacity)}
	})
	connretryp = util.NewPool(func(capacity int) *connretry {
		return &connretry{m: make(map[conn]*retry, capacity), n: capacity}
	})
	conncountp = util.NewPool(func(capacity int) *conncount {
		return &conncount{m: make(map[conn]int, capacity), n: capacity}
	})
	connretrycachep = util.NewPool(func(capacity int) *connretrycache {
		return &connretrycache{m: make(map[conn]*retrycache, capacity), n: capacity}
	})
)

type muxslots struct {
	s []int
}

func (r *muxslots) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *muxslots) ResetLen(n int) { _ = "STUB: not implemented"; return }

func (r *muxslots) LessThen(n int) bool { _ = "STUB: not implemented"; return false }

type redisresults struct {
	s []RedisResult
}

func (r *redisresults) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *redisresults) ResetLen(n int) { _ = "STUB: not implemented"; return }

type cacheentries struct {
	e map[int]CacheEntry
	c int
}

func (c *cacheentries) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (c *cacheentries) ResetLen(n int) { _ = "STUB: not implemented"; return }

var entriesp = util.NewPool(func(capacity int) *cacheentries {
	return &cacheentries{e: make(map[int]CacheEntry, capacity), c: capacity}
})

type mgetcachecmds struct {
	s []CacheableTTL
}

func (r *mgetcachecmds) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *mgetcachecmds) ResetLen(n int) { _ = "STUB: not implemented"; return }

type mgetcmds struct {
	s []Completed
}

func (r *mgetcmds) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *mgetcmds) ResetLen(n int) { _ = "STUB: not implemented"; return }

type retry struct {
	cIndexes []int
	commands []Completed
	aIndexes []int
	cAskings []Completed
}

func (r *retry) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *retry) ResetLen(n int) { _ = "STUB: not implemented"; return }

type retrycache struct {
	cIndexes []int
	commands []CacheableTTL
	aIndexes []int
	cAskings []CacheableTTL
}

func (r *retrycache) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *retrycache) ResetLen(n int) { _ = "STUB: not implemented"; return }

type batchcache struct {
	cIndexes []int
	commands []CacheableTTL
}

func (r *batchcache) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *batchcache) ResetLen(n int) { _ = "STUB: not implemented"; return }

type batchcachemap struct {
	m map[uint16]*batchcache
	n int
}

func (r *batchcachemap) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *batchcachemap) ResetLen(n int) { _ = "STUB: not implemented"; return }

type conncount struct {
	m map[conn]int
	n int
}

func (r *conncount) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *conncount) ResetLen(n int) { _ = "STUB: not implemented"; return }

type connretry struct {
	m          map[conn]*retry
	n          int
	RetryDelay time.Duration // NOTE: This is not thread-safe.
	Redirects  uint32        // NOTE: This is not thread-safe.
}

func (r *connretry) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *connretry) ResetLen(n int) { _ = "STUB: not implemented"; return }

// No retry.

type connretrycache struct {
	m          map[conn]*retrycache
	n          int
	RetryDelay time.Duration // NOTE: This is not thread-safe.
	Redirects  uint32        // NOTE: This is not thread-safe.
}

func (r *connretrycache) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *connretrycache) ResetLen(n int) { _ = "STUB: not implemented"; return }

// No retry.
