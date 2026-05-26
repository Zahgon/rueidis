package rueidislimiter

import "time"

type RateLimitOption struct {
	limit  int64
	window time.Duration
}

func WithCustomRateLimit(limit int, window time.Duration) RateLimitOption {
	_ = "STUB: not implemented"
	return *new(RateLimitOption)
}
