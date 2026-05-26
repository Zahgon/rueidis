package rueidis

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/rueidis/internal/cmds"
)

type connFn func(dst string, opt *ClientOption) conn
type dialFn func(ctx context.Context, dst string, opt *ClientOption) (net.Conn, error)
type wireFn func(ctx context.Context) wire

type singleconnect struct {
	w wire
	e error
	g sync.WaitGroup
}

type conn interface {
	Do(ctx context.Context, cmd Completed) RedisResult
	DoCache(ctx context.Context, cmd Cacheable, ttl time.Duration) RedisResult
	DoMulti(ctx context.Context, multi ...Completed) *redisresults
	DoMultiCache(ctx context.Context, multi ...CacheableTTL) *redisresults
	Receive(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error
	DoStream(ctx context.Context, cmd Completed) RedisResultStream
	DoMultiStream(ctx context.Context, multi ...Completed) MultiRedisResultStream
	Info() map[string]RedisMessage
	Version() int
	AZ() string
	Error() error
	Close()
	Dial() error
	Override(conn)
	Acquire(ctx context.Context) wire
	Store(w wire)
	Addr() string
	SetOnCloseHook(func(error))
	OptInCmd() cmds.Completed
}

var _ conn = (*mux)(nil)

type muxwire struct {
	wire atomic.Value
	sc   *singleconnect
	mu   sync.Mutex
}

type mux struct {
	init     wire
	dead     wire
	clhks    atomic.Value
	dpool    *pool
	spool    *pool
	wireFn   wireFn
	dst      string
	muxwires []muxwire
	maxp     int
	maxm     int

	usePool bool
	optIn   bool
}

func makeMux(dst string, option *ClientOption, dialFn dialFn) *mux {
	_ = "STUB: not implemented"
	return nil
}

func newMux(dst string, option *ClientOption, init, dead wire, wireFn wireFn, wireNoBgFn wireFn) *mux {
	_ = "STUB: not implemented"
	return nil
}

func isOptIn(opts []string) bool { _ = "STUB: not implemented"; return false }

func (m *mux) OptInCmd() cmds.Completed { _ = "STUB: not implemented"; return *new(cmds.Completed) }

func (m *mux) SetOnCloseHook(fn func(error)) { _ = "STUB: not implemented"; return }

func (m *mux) setCloseHookOnWire(i uint16, w wire) { _ = "STUB: not implemented"; return }

func (m *mux) Override(cc conn) { _ = "STUB: not implemented"; return }

// bind the new m to the old w

func (m *mux) _pipe(ctx context.Context, i uint16) (w wire, err error) {
	_ = "STUB: not implemented"
	return *new(wire), nil
}

func (m *mux) pipe(ctx context.Context, i uint16) wire {
	_ = "STUB: not implemented"
	return *new(wire)
}

// this should never be nil

func (m *mux) Dial() error { _ = "STUB: not implemented"; return nil }

func (m *mux) Info() map[string]RedisMessage { _ = "STUB: not implemented"; return nil }

func (m *mux) Version() int { _ = "STUB: not implemented"; return 0 }

func (m *mux) AZ() string { _ = "STUB: not implemented"; return "" }

func (m *mux) Error() error { _ = "STUB: not implemented"; return nil }

func (m *mux) DoStream(ctx context.Context, cmd Completed) RedisResultStream {
	_ = "STUB: not implemented"
	return *new(RedisResultStream)
}

func (m *mux) DoMultiStream(ctx context.Context, multi ...Completed) MultiRedisResultStream {
	_ = "STUB: not implemented"
	return *new(MultiRedisResultStream)
}

func (m *mux) Do(ctx context.Context, cmd Completed) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (m *mux) DoMulti(ctx context.Context, multi ...Completed) (resp *redisresults) {
	_ = "STUB: not implemented"
	return nil
}

// mark the first cmd as blocked if one of them is blocked to shortcut later check.

// use a dedicated connection if the pipeline is too large

func (m *mux) blocking(pool *pool, ctx context.Context, cmd Completed) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

// abort the wire if blocking command return early (ex. context.DeadlineExceeded)

func (m *mux) blockingMulti(pool *pool, ctx context.Context, cmd []Completed) (resp *redisresults) {
	_ = "STUB: not implemented"
	return nil
}

// abort the wire if blocking command return early (ex. context.DeadlineExceeded)

func (m *mux) pipeline(ctx context.Context, cmd Completed) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (m *mux) pipelineMulti(ctx context.Context, cmd []Completed) (resp *redisresults) {
	_ = "STUB: not implemented"
	return nil
}

func (m *mux) DoCache(ctx context.Context, cmd Cacheable, ttl time.Duration) RedisResult {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (m *mux) DoMultiCache(ctx context.Context, multi ...CacheableTTL) (results *redisresults) {
	_ = "STUB: not implemented"
	return nil
}

func (m *mux) doMultiCache(ctx context.Context, slot uint16, multi []CacheableTTL) (resps *redisresults) {
	_ = "STUB: not implemented"
	return nil
}

func (m *mux) Receive(ctx context.Context, subscribe Completed, fn func(message PubSubMessage)) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mux) Acquire(ctx context.Context) wire { _ = "STUB: not implemented"; return *new(wire) }

func (m *mux) Store(w wire) { _ = "STUB: not implemented"; return }

func (m *mux) Close() { _ = "STUB: not implemented"; return }

func (m *mux) Addr() string { _ = "STUB: not implemented"; return "" }

func isBroken(err error, w wire) bool { _ = "STUB: not implemented"; return false }

func slotfn(n int, ks uint16, noreply bool) uint16 { _ = "STUB: not implemented"; return 0 }
