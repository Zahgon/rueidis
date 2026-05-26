package rueidis

import (
	"context"
	"time"
)

const (
	defaultMaxRetries    = 20
	defaultMaxRetryDelay = 1 * time.Second
)

// RetryDelayFn returns the delay that should be used before retrying the
// attempt. Will return a negative delay if the delay could not be determined or does not retry.
type RetryDelayFn func(attempts int, cmd Completed, err error) time.Duration

// defaultRetryDelayFn delays the next retry exponentially without considering the error.
// Max delay is 1 second.
// This "Equal Jitter" delay produced by this implementation is not monotonic increasing. ref: https://aws.amazon.com/ko/blogs/architecture/exponential-backoff-and-jitter/
func defaultRetryDelayFn(attempts int, _ Completed, _ error) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type retryHandler interface {
	// RetryDelay returns the delay that should be used before retrying the
	// attempt. Will return a negative delay if the delay could not be determined or does
	// not retry.
	// If the delay is zero, the next retry should be attempted immediately.
	RetryDelay(attempts int, cmd Completed, err error) time.Duration

	// WaitForRetry waits until the next retry should be attempted.
	WaitForRetry(ctx context.Context, duration time.Duration)

	// WaitOrSkipRetry waits until the next retry should be attempted
	// or returns false if the command should not be retried.
	// Returns false immediately if the command should not be retried.
	// Returns true after the delay if the command should be retried.
	WaitOrSkipRetry(ctx context.Context, attempts int, cmd Completed, err error) bool
}

type retryer struct {
	RetryDelayFn RetryDelayFn
}

var _ retryHandler = (*retryer)(nil)

func newRetryer(retryDelayFn RetryDelayFn) *retryer { _ = "STUB: not implemented"; return nil }

func (r *retryer) RetryDelay(attempts int, cmd Completed, err error) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *retryer) WaitForRetry(ctx context.Context, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (r *retryer) WaitOrSkipRetry(
	ctx context.Context, attempts int, cmd Completed, err error,
) bool {
	_ = "STUB: not implemented"
	return false
}
