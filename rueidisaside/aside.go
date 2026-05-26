package rueidisaside

import (
	"context"
	"sync"
	"time"

	"github.com/redis/rueidis"
)

type ClientOption struct {
	// ClientBuilder can be used to modify rueidis.Client used by Locker
	ClientBuilder func(option rueidis.ClientOption) (rueidis.Client, error)
	ClientOption  rueidis.ClientOption
	ClientTTL     time.Duration // TTL for the client marker, refreshed every 1/2 TTL. Defaults to 10s. The marker allows other clients to know if this client is still alive.
	UseLuaLock    bool
}

type CacheAsideClient interface {
	Get(ctx context.Context, ttl time.Duration, key string, fn func(ctx context.Context, key string) (val string, err error)) (val string, err error)
	Del(ctx context.Context, key string) error
	Client() rueidis.Client
	Close()
}

func NewClient(option ClientOption) (cc CacheAsideClient, err error) {
	_ = "STUB: not implemented"
	return *new(CacheAsideClient), nil
}

type Client struct {
	client     rueidis.Client
	ctx        context.Context
	waits      map[string]chan struct{}
	cancel     context.CancelFunc
	id         string
	ttl        time.Duration
	mu         sync.Mutex
	useLuaLock bool
}

func (c *Client) onInvalidation(messages []rueidis.RedisMessage) { _ = "STUB: not implemented"; return }

func (c *Client) register(key string) (ch chan struct{}) { _ = "STUB: not implemented"; return nil }

func (c *Client) refresh(id string) { _ = "STUB: not implemented"; return }

// client id has changed, abort this goroutine

func (c *Client) keepalive() (id string, err error) { _ = "STUB: not implemented"; return "", nil }

// randStr generates a 24-byte long, random string.
func randStr() string { _ = "STUB: not implemented"; return "" }

func (c *Client) Get(ctx context.Context, ttl time.Duration, key string, fn func(ctx context.Context, key string) (val string, err error)) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// cache miss, prepare to populate the value by fn()

// acquire client id

// successfully set client id on the key as a lock
// attach TTL pointer to context for potential modification via OverrideCacheTTL

// failed to populate the value, release the lock.

// the client who held the lock has gone, release the lock.

func (c *Client) Del(ctx context.Context, key string) error { _ = "STUB: not implemented"; return nil }

// Client exports the underlying rueidis.Client
func (c *Client) Client() rueidis.Client { _ = "STUB: not implemented"; return *new(rueidis.Client) }

func (c *Client) Close() { _ = "STUB: not implemented"; return }

const PlaceholderPrefix = "rueidisid:"

type ctxKey struct{}

var ttlKey = ctxKey{}

// OverrideCacheTTL sets a custom TTL for the cache entry being populated in the current context.
// It can be called in the callback function passed to CacheAsideClient.Get() to customize
// the TTL based on the data being cached.
func OverrideCacheTTL(ctx context.Context, ttl time.Duration) { _ = "STUB: not implemented"; return }

var (
	delkey      = rueidis.NewLuaScript(`if redis.call("GET",KEYS[1]) == ARGV[1] then return redis.call("DEL",KEYS[1]) else return 0 end`)
	setkey      = rueidis.NewLuaScript(`if redis.call("GET",KEYS[1]) == ARGV[1] then return redis.call("SET",KEYS[1],ARGV[2],"PX",ARGV[3]) else return 0 end`)
	acquireLock = rueidis.NewLuaScript(`if redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2]) then return nil else return redis.call("GET", KEYS[1]) end`)
)
