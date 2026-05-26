package rueidislimiter

import (
	"context"
	"errors"
	"time"

	"github.com/redis/rueidis"
)

var (
	ErrInvalidTokens   = errors.New("number of tokens must be non-negative")
	ErrInvalidResponse = errors.New("invalid response from Redis")
	ErrInvalidLimit    = errors.New("limit must be positive")
	ErrInvalidWindow   = errors.New("window must be positive")
	ErrNilBuilder      = errors.New("client builder is required")
)

type Result struct {
	Allowed   bool
	Remaining int64
	ResetAtMs int64
}

type RateLimiterClient interface {
	Check(ctx context.Context, identifier string, options ...RateLimitOption) (Result, error)
	Allow(ctx context.Context, identifier string, options ...RateLimitOption) (Result, error)
	AllowN(ctx context.Context, identifier string, n int64, options ...RateLimitOption) (Result, error)
	Limit() int
}

const (
	PlaceholderPrefix = "rueidislimiter"
	keyDelimOpen      = ":{"
	keyDelimClose     = "}"
)

type rateLimiter struct {
	client           rueidis.Client
	keyPrefix        string
	defaultRateLimit RateLimitOption
}

type RateLimiterOption struct {
	ClientBuilder func(option rueidis.ClientOption) (rueidis.Client, error)
	KeyPrefix     string
	ClientOption  rueidis.ClientOption
	Limit         int
	Window        time.Duration
}

func NewRateLimiter(option RateLimiterOption) (RateLimiterClient, error) {
	_ = "STUB: not implemented"
	return *new(RateLimiterClient), nil
}

func (l *rateLimiter) Limit() int { _ = "STUB: not implemented"; return 0 }

func (l *rateLimiter) Check(ctx context.Context, identifier string, options ...RateLimitOption) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (l *rateLimiter) Allow(ctx context.Context, identifier string, options ...RateLimitOption) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (l *rateLimiter) AllowN(ctx context.Context, identifier string, n int64, options ...RateLimitOption) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

var rateLimitScript = rueidis.NewLuaScript(`
local rate_limit_key = KEYS[1]
local increment_amount = tonumber(ARGV[1])
local next_expires_at = tonumber(ARGV[2])
local current_time = tonumber(ARGV[3])
local expires_at_key = KEYS[2]
local expires_at = tonumber(redis.call("get", expires_at_key))
if not expires_at or expires_at < current_time then
  redis.call("set", rate_limit_key, 0, "pxat", next_expires_at + 1000)
  redis.call("set", expires_at_key, next_expires_at, "pxat", next_expires_at + 1000)
  expires_at = next_expires_at
end
local current = redis.call("incrby", rate_limit_key, increment_amount)
return { current, expires_at }
`)
