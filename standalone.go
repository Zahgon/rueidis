package rueidis

import (
	"context"
	"sync/atomic"
	"time"
)

func newStandaloneClient(opt *ClientOption, connFn connFn, retryer retryHandler) (*standalone, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// close primary if any replica fails

type standalone struct {
	retryer        retryHandler
	toReplicas     func(Completed) bool
	nodeSelector   func(uint16, []NodeInfo) int
	primary        atomic.Pointer[singleClient]
	connFn         connFn
	opt            *ClientOption
	redirectCall   call
	replicas       []*singleClient
	nodes          []NodeInfo
	enableRedirect bool
}

func (s *standalone) B() Builder { _ = "STUB: not implemented"; return *new(Builder) }

func (s *standalone) pick(slot uint16) *singleClient { _ = "STUB: not implemented"; return nil }

func (s *standalone) redirectToPrimary(addr string) error {
	_ = "STUB: not implemented"
	// Create a new connection to the redirect address
	return nil
}

// Create a new primary client with the redirect connection

// Atomically swap the primary and close the old one

func (s *standalone) handleRedirect(ctx context.Context, err error) (error, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *standalone) Do(ctx context.Context, cmd Completed) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (s *standalone) DoMulti(ctx context.Context, multi ...Completed) (resp []RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (s *standalone) Receive(ctx context.Context, subscribe Completed, fn func(msg PubSubMessage)) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *standalone) Close() { _ = "STUB: not implemented"; return }

func (s *standalone) DoCache(ctx context.Context, cmd Cacheable, ttl time.Duration) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (s *standalone) DoMultiCache(ctx context.Context, multi ...CacheableTTL) (resp []RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (s *standalone) DoStream(ctx context.Context, cmd Completed) RedisResultStream {
	_ = "STUB: not implemented"
	return *new(RedisResultStream)
}

func (s *standalone) DoMultiStream(ctx context.Context, multi ...Completed) MultiRedisResultStream {
	_ = "STUB: not implemented"
	return *new(MultiRedisResultStream)
}

func (s *standalone) Dedicated(fn func(DedicatedClient) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *standalone) Dedicate() (client DedicatedClient, cancel func()) {
	_ = "STUB: not implemented"
	return *new(DedicatedClient), nil
}

func (s *standalone) Nodes() map[string]Client { _ = "STUB: not implemented"; return nil }

func (s *standalone) Mode() ClientMode { _ = "STUB: not implemented"; return *new(ClientMode) }
