package rueidis

import (
	"context"
	"time"
)

type singleClient struct {
	conn         conn
	retryHandler retryHandler
	stop         uint32
	cmd          Builder
	retry        bool
	hasLftm      bool
	DisableCache bool
}

func newSingleClient(opt *ClientOption, prev conn, connFn connFn, retryer retryHandler) (*singleClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// return a singleClient instance even when Dial fails if ForceSingleClient is true.

func newSingleClientWithConn(conn conn, builder Builder, retry, disableCache bool, retryer retryHandler, hasLftm bool) *singleClient {
	_ = "STUB: not implemented"
	return nil
}

func (c *singleClient) B() Builder { _ = "STUB: not implemented"; return *new(Builder) }

func (c *singleClient) Do(ctx context.Context, cmd Completed) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

// not recycle cmds if error, since cmds may be used later in the pipe.

func (c *singleClient) DoStream(ctx context.Context, cmd Completed) RedisResultStream {
	_ = "STUB: not implemented"
	return *new(RedisResultStream)
}

func (c *singleClient) DoMultiStream(ctx context.Context, multi ...Completed) MultiRedisResultStream {
	_ = "STUB: not implemented"
	return *new(MultiRedisResultStream)
}

func (c *singleClient) DoMulti(ctx context.Context, multi ...Completed) (resps []RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

// check transaction block, if zero, then not in transaction

// if no error, then check if transaction block

func (c *singleClient) DoMultiCache(ctx context.Context, multi ...CacheableTTL) (resps []RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (c *singleClient) DoCache(ctx context.Context, cmd Cacheable, ttl time.Duration) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (c *singleClient) Receive(ctx context.Context, subscribe Completed, fn func(msg PubSubMessage)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *singleClient) Dedicated(fn func(DedicatedClient) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *singleClient) Dedicate() (DedicatedClient, func()) {
	_ = "STUB: not implemented"
	return *new(DedicatedClient), nil
}

func (c *singleClient) Nodes() map[string]Client { _ = "STUB: not implemented"; return nil }

func (c *singleClient) Mode() ClientMode { _ = "STUB: not implemented"; return *new(ClientMode) }

func (c *singleClient) Close() { _ = "STUB: not implemented"; return }

type dedicatedSingleClient struct {
	conn         conn
	wire         wire
	retryHandler retryHandler
	mark         uint32
	cmd          Builder
	retry        bool
}

func (c *dedicatedSingleClient) B() Builder { _ = "STUB: not implemented"; return *new(Builder) }

func (c *dedicatedSingleClient) Do(ctx context.Context, cmd Completed) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (c *dedicatedSingleClient) DoMulti(ctx context.Context, multi ...Completed) (resp []RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (c *dedicatedSingleClient) Receive(ctx context.Context, subscribe Completed, fn func(msg PubSubMessage)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *dedicatedSingleClient) SetPubSubHooks(hooks PubSubHooks) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dedicatedSingleClient) SetOnInvalidations(fn func([]RedisMessage)) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dedicatedSingleClient) Close() { _ = "STUB: not implemented"; return }

func (c *dedicatedSingleClient) check() error { _ = "STUB: not implemented"; return nil }

func (c *dedicatedSingleClient) release() { _ = "STUB: not implemented"; return }

func (c *singleClient) isRetryable(err error, ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func isRetryable(err error, w wire, ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func allRetryable(multi []Completed) bool { _ = "STUB: not implemented"; return false }

func chooseSlot(multi []Completed) uint16 { _ = "STUB: not implemented"; return 0 }
