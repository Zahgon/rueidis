package rueidislock

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/redis/rueidis"
)

// LockerOption should be passed to NewLocker to construct a Locker
type LockerOption struct {
	// ClientBuilder can be used to modify rueidis.Client used by Locker
	ClientBuilder func(option rueidis.ClientOption) (rueidis.Client, error)
	// KeyPrefix is the prefix of the redis key for locks. The default value is "rueidislock".
	KeyPrefix string
	// ClientOption is passed to rueidis.NewClient or LockerOption.ClientBuilder to build a rueidis.Client
	ClientOption rueidis.ClientOption
	// KeyValidity is the validity duration of locks and will be extended periodically by the ExtendInterval. The default value is 5s.
	KeyValidity time.Duration
	// ExtendInterval is the interval to extend KeyValidity. Default value is 1/2 of KeyValidity.
	ExtendInterval time.Duration
	// TryNextAfter is the timeout duration before trying the next redis key for locks. The default value is 20ms.
	TryNextAfter time.Duration
	// KeyMajority is at least how many redis keys in a total of KeyMajority*2-1 should be acquired to be a valid lock.
	// The default value is 2.
	KeyMajority int32
	// NoLoopTracking will use NOLOOP in the CLIENT TRACKING command to avoid unnecessary notifications and thus have better performance.
	// This can only be enabled if all your redis nodes >= 7.0.5. (https://github.com/redis/redis/pull/11052)
	NoLoopTracking bool
	// Use SET PX instead of SET PXAT when acquiring locks to be compatible with Redis < 6.2
	FallbackSETPX bool
}

// Locker is the interface of rueidislock
type Locker interface {
	// WithContext acquires a distributed redis lock by name by waiting for it. It may return ErrLockerClosed.
	WithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error)
	// TryWithContext tries to acquire a distributed redis lock by name without waiting. It may return ErrNotLocked.
	TryWithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error)
	// ForceWithContext takes over a distributed redis lock by canceling the original holder. It may return ErrNotLocked.
	ForceWithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error)
	// Client exports the underlying rueidis.Client
	Client() rueidis.Client
	// Close closes the underlying rueidis.Client
	Close()
}

// NewLocker creates the distributed Locker backed by redis client side caching
func NewLocker(option LockerOption) (Locker, error) {
	_ = "STUB: not implemented"
	return *new(Locker), nil
}

// this ensures the CSC goes to the same connection.

type locker struct {
	client   rueidis.Client
	gates    map[string]*gate
	prefix   string
	validity time.Duration
	interval time.Duration
	timeout  time.Duration
	mu       sync.RWMutex
	majority int32
	totalcnt int32
	noloop   bool
	nocsc    bool
	setpx    bool
}

type gate struct {
	ch  chan struct{}
	csc []chan struct{}
	w   int
}

func makegate(size int32) *gate { _ = "STUB: not implemented"; return nil }

func random() string { _ = "STUB: not implemented"; return "" }

func keyname(prefix, name string, i int32) string { _ = "STUB: not implemented"; return "" }

func (m *locker) acquire(ctx context.Context, key, val string, duration time.Duration, deadline time.Time, force bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *locker) script(ctx context.Context, script *rueidis.Lua, key, val string, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *locker) getgate(name string) (g *gate) { _ = "STUB: not implemented"; return nil }

func (m *locker) removegate(g *gate, name string) { _ = "STUB: not implemented"; return }

func (m *locker) onInvalidations(messages []rueidis.RedisMessage) {
	_ = "STUB: not implemented"
	return
}

func (m *locker) try(ctx context.Context, cancel context.CancelFunc, name string, g *gate, force bool) (context.CancelFunc, error) {
	_ = "STUB: not implemented"
	return *new(context.CancelFunc), nil
}

func (m *locker) tryonce(ctx context.Context, name string, force bool) (context.Context, context.CancelFunc, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc), nil
}

func (m *locker) ForceWithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc), nil
}

func (m *locker) TryWithContext(ctx context.Context, name string) (context.Context, context.CancelFunc, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc), nil
}

func (m *locker) WithContext(src context.Context, name string) (context.Context, context.CancelFunc, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc), nil
}

func (m *locker) Client() rueidis.Client { _ = "STUB: not implemented"; return *new(rueidis.Client) }

func (m *locker) Close() { _ = "STUB: not implemented"; return }

var (
	delkey = rueidis.NewLuaScriptRetryable(`if redis.call("GET",KEYS[1]) == ARGV[1] then return redis.call("DEL",KEYS[1]) end;return 0`)
	extend = rueidis.NewLuaScriptRetryable(`if redis.call("GET",KEYS[1]) == ARGV[1] then local r = redis.call("PEXPIREAT",KEYS[1],ARGV[2]);redis.call("GET",KEYS[1]);return r end;return 0`)
	acqms  = rueidis.NewLuaScriptRetryable(`local r = redis.call("SET",KEYS[1],ARGV[1],"NX","PX",ARGV[2]);redis.call("GET",KEYS[1]);return r`)
	acqat  = rueidis.NewLuaScriptRetryable(`local r = redis.call("SET",KEYS[1],ARGV[1],"NX","PXAT",ARGV[2]);redis.call("GET",KEYS[1]);return r`)
	fcqms  = rueidis.NewLuaScriptRetryable(`local r = redis.call("SET",KEYS[1],ARGV[1],"PX",ARGV[2]);redis.call("GET",KEYS[1]);return r`)
	fcqat  = rueidis.NewLuaScriptRetryable(`local r = redis.call("SET",KEYS[1],ARGV[1],"PXAT",ARGV[2]);redis.call("GET",KEYS[1]);return r`)
)

// ErrNotLocked is returned from the Locker.TryWithContext when it fails
var ErrNotLocked = errors.New("not locked")

// ErrLockerClosed is returned from the Locker.WithContext when the Locker is closed
var ErrLockerClosed = errors.New("locker closed")
