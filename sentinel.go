package rueidis

import (
	"container/list"
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

func newSentinelClient(opt *ClientOption, connFn connFn, retryer retryHandler) (client *sentinelClient, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type sentinelClient struct {
	mConn        atomic.Value
	rConn        atomic.Value
	sConn        conn
	retryHandler retryHandler
	connFn       connFn
	mOpt         *ClientOption
	sOpt         *ClientOption
	rOpt         *ClientOption
	sentinels    *list.List
	mAddr        atomic.Value
	rAddr        atomic.Value
	sAddr        string
	sc           call
	mu           sync.Mutex
	stop         uint32
	cmd          Builder
	retry        bool
	hasLftm      bool
	replica      bool
}

func (c *sentinelClient) B() Builder { _ = "STUB: not implemented"; return *new(Builder) }

func (c *sentinelClient) Do(ctx context.Context, cmd Completed) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

// not recycle cmds if error, since cmds may be used later in the pipe.

func (c *sentinelClient) DoMulti(ctx context.Context, multi ...Completed) []RedisResult {
	_ = "STUB: not implemented"
	return nil
}

// check transaction block, if zero, then not in transaction

// if no error, then check if transaction block

func (c *sentinelClient) DoCache(ctx context.Context, cmd Cacheable, ttl time.Duration) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (c *sentinelClient) DoMultiCache(ctx context.Context, multi ...CacheableTTL) []RedisResult {
	_ = "STUB: not implemented"
	return nil
}

func (c *sentinelClient) Receive(ctx context.Context, subscribe Completed, fn func(msg PubSubMessage)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *sentinelClient) DoStream(ctx context.Context, cmd Completed) RedisResultStream {
	_ = "STUB: not implemented"
	return *new(RedisResultStream)
}

func (c *sentinelClient) DoMultiStream(ctx context.Context, multi ...Completed) MultiRedisResultStream {
	_ = "STUB: not implemented"
	return *new(MultiRedisResultStream)
}

func (c *sentinelClient) Dedicated(fn func(DedicatedClient) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *sentinelClient) Dedicate() (DedicatedClient, func()) {
	_ = "STUB: not implemented"
	return *new(DedicatedClient), nil
}

func (c *sentinelClient) Nodes() map[string]Client { _ = "STUB: not implemented"; return nil }

func (c *sentinelClient) Mode() ClientMode { _ = "STUB: not implemented"; return *new(ClientMode) }

func (c *sentinelClient) Close() { _ = "STUB: not implemented"; return }

func (c *sentinelClient) isRetryable(err error, ctx context.Context) (should bool) {
	_ = "STUB: not implemented"
	return false
}

func (c *sentinelClient) addSentinel(addr string) { _ = "STUB: not implemented"; return }

func (c *sentinelClient) _addSentinel(addr string) { _ = "STUB: not implemented"; return }

func (c *sentinelClient) pick(cmd Completed) (cc conn) {
	_ = "STUB: not implemented"
	return *new(conn)
}

func (c *sentinelClient) pickMulti(sendToReplica bool) (cc conn) {
	_ = "STUB: not implemented"
	return *new(conn)
}

func (c *sentinelClient) sendAllToReplica(cmds []Completed) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *sentinelClient) sendAllToReplicaCache(cmds []CacheableTTL) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *sentinelClient) switchTargetRetry(addr string, isMaster bool) {
	_ = "STUB: not implemented"
	return
}

func (c *sentinelClient) _switchTarget(addr string, isMaster bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *sentinelClient) refreshRetry() { _ = "STUB: not implemented"; return }

func (c *sentinelClient) refresh() (err error) { _ = "STUB: not implemented"; return nil }

func (c *sentinelClient) _refresh() (err error) { _ = "STUB: not implemented"; return nil }

// listWatch returns the server address with sentinels.
// check if the target is master or replica

// listWatch will use sentinel to list the current master,replica address along with sentinel address
func (c *sentinelClient) listWatch(cc conn) (master string, replica string, sentinels []string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil, nil
}

// not recycle cmds if error, since cmds may be used later in the pipe.

// unsubscribe in case there is any previous subscription

// note that in case of failover, every slave in the setup
// will send +slave event individually.

// call refresh to randomly choose a new slave

// we return a random slave address instead of master

func pickReplica(resp RedisResult) (string, error) { _ = "STUB: not implemented"; return "", nil }

// eliminate replicas with the s_down condition

// choose a replica randomly

func newSentinelOpt(opt *ClientOption) *ClientOption { _ = "STUB: not implemented"; return nil }

// https://github.com/redis/rueidis/issues/138

var (
	errNotMaster = errors.New("the redis role is not master")
	errNotSlave  = errors.New("the redis role is not slave")
)
