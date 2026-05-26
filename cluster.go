package rueidis

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/redis/rueidis/internal/util"
)

// ErrNoSlot indicates that there is no redis node owning the key slot.
var ErrNoSlot = errors.New("the slot has no redis node")
var ErrReplicaOnlyConflict = errors.New("ReplicaOnly conflicts with SendToReplicas option")
var ErrInvalidShardsRefreshInterval = errors.New("ShardsRefreshInterval must be greater than or equal to 0")
var ErrReplicaOnlyConflictWithReplicaSelector = errors.New("ReplicaOnly conflicts with ReplicaSelector option")
var ErrReplicaOnlyConflictWithReadNodeSelector = errors.New("ReplicaOnly conflicts with ReadNodeSelector option")
var ErrReplicaSelectorConflictWithReadNodeSelector = errors.New("either set ReplicaSelector or ReadNodeSelector, not both")
var ErrSendToReplicasNotSet = errors.New("SendToReplicas must be set when ReplicaSelector is set")

type clusterClient struct {
	wslots       [16384]conn
	retryHandler retryHandler
	opt          *ClientOption
	rOpt         *ClientOption
	conns        map[string]connrole
	connFn       connFn
	stopCh       chan struct{}
	sc           call
	rslots       [][]NodeInfo
	mu           sync.RWMutex
	stop         uint32
	cmd          Builder
	retry        bool
	hasLftm      bool
}

// NOTE: connrole and conn must be initialized at the same time
type connrole struct {
	conn   conn
	hidden bool
	//replica bool <- this field is removed because a server may have mixed roles at the same time in the future. https://github.com/valkey-io/valkey/issues/1372
}

var replicaOnlySelector = func(_ uint16, replicas []NodeInfo) int {
	return util.FastRand(len(replicas))
}

func newClusterClient(opt *ClientOption, connFn connFn, retryer retryHandler) (*clusterClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterClient) init() error {
	if len(c.opt.InitAddress) == 0 {
		return ErrNoAddr
	}
	results := make(chan error, len(c.opt.InitAddress))
	for _, addr := range c.opt.InitAddress {
		cc := c.connFn(addr, c.opt)
		go func(addr string, cc conn) {
			if err := cc.Dial(); err == nil {
				c.mu.Lock()
				if _, ok := c.conns[addr]; ok {
					go cc.Close() // abort the new connection instead of closing the old one, which may already been used
				} else {
					c.conns[addr] = connrole{
						conn: cc,
					}
				}
				c.mu.Unlock()
				results <- nil
			} else {
				results <- err
			}
		}(addr, cc)
	}
	es := make([]error, cap(results))
	for i := 0; i < cap(results); i++ {
		if err := <-results; err == nil {
			return nil
		} else {
			es[i] = err
		}
	}
	return es[0]
}

func (c *clusterClient) refresh(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *clusterClient) lazyRefresh() { _ = "STUB: not implemented"; return }

type clusterslots struct {
	addr  string
	reply RedisResult
	ver   int
}

func (s clusterslots) parse(tls bool) map[string]group { _ = "STUB: not implemented"; return nil }

func getClusterSlots(c conn, timeout time.Duration) clusterslots {
	_ = "STUB: not implemented"
	return *new(clusterslots)
}

func (c *clusterClient) _refresh() (err error) { _ = "STUB: not implemented"; return nil }

// batch CLUSTER SLOTS/CLUSTER SHARDS for every 4 connections

// make sure InitAddress always be present

// lazy init

// batch AZ() for every 4 connections

// exclude master node

// fallback to master

func (c *clusterClient) single() (conn conn) { _ = "STUB: not implemented"; return *new(conn) }

func (c *clusterClient) nodes() []string { _ = "STUB: not implemented"; return nil }

type nodes []NodeInfo

type group struct {
	nodes nodes
	slots [][2]int64
}

func parseEndpoint(fallback, endpoint string, port int64) string {
	_ = "STUB: not implemented"
	return ""
}

// parseSlots - map redis slots for each redis nodes/addresses
// defaultAddr is needed in case the node does not know its own IP
func parseSlots(slots RedisMessage, defaultAddr string) map[string]group {
	_ = "STUB: not implemented"
	return nil
}

// parseShards - map redis shards for each redis nodes/addresses
// defaultAddr is needed in case the node does not know its own IP
func parseShards(shards RedisMessage, defaultAddr string, tls bool) map[string]group {
	_ = "STUB: not implemented"
	return nil
}

func (c *clusterClient) runClusterTopologyRefreshment() { _ = "STUB: not implemented"; return }

func (c *clusterClient) _pick(slot uint16, toReplica bool) (p conn) {
	_ = "STUB: not implemented"
	return *new(conn)
}

func (c *clusterClient) pick(ctx context.Context, slot uint16, toReplica bool) (p conn, err error) {
	_ = "STUB: not implemented"
	return *new(conn), nil
}

func (c *clusterClient) redirectOrNew(addr string, prev conn, slot uint16, mode RedirectMode) conn {
	_ = "STUB: not implemented"
	return *new(conn)
}

// try reconnection if the MOVED redirects to the same host,
// because the same hostname may actually be resolved into another destination
// depending on the fail-over implementation. ex: AWS MemoryDB's resize process.

// MOVED should always point to the primary.

func (c *clusterClient) B() Builder { _ = "STUB: not implemented"; return *new(Builder) }

func (c *clusterClient) Do(ctx context.Context, cmd Completed) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

// not recycle cmds if error, since cmds may be used later in the pipe.

func (c *clusterClient) do(ctx context.Context, cmd Completed) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (c *clusterClient) toReplica(cmd Completed) bool { _ = "STUB: not implemented"; return false }

func (c *clusterClient) _pickMulti(multi []Completed) (retries *connretry, init bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// the default itor[i] is 0

// if all commands have no slots, such as INFO, we pick a non-nil slot.

func (c *clusterClient) pickMulti(ctx context.Context, multi []Completed) (*connretry, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func isMulti(cmd Completed) bool { _ = "STUB: not implemented"; return false }

func isExec(cmd Completed) bool { _ = "STUB: not implemented"; return false }

func (c *clusterClient) doresultfn(
	ctx context.Context, results *redisresults, retries *connretry, mu *sync.Mutex, cc conn, cIndexes []int, commands []Completed, resps []RedisResult, attempts int, hasInit bool,
) (clean bool) {
	_ = "STUB: not implemented"
	return false
}

// find out if there is a transaction block or not.

// a transaction is found.

// the transaction has been added to the retries, go to the next cmd.

// the current cmd is in the processed transaction and has been added to the retries.

func (c *clusterClient) doretry(
	ctx context.Context, cc conn, results *redisresults, retries *connretry, re *retry, mu *sync.Mutex, wg *sync.WaitGroup, attempts int, hasInit bool,
) {
	_ = "STUB: not implemented"
	return
}

// check transaction block, if zero, then not in transaction

// if no error, then check if transaction block

func (c *clusterClient) DoMulti(ctx context.Context, multi ...Completed) []RedisResult {
	_ = "STUB: not implemented"
	return nil
}

// Assume no retry. Because a client retry flag can be set to false.

func fillErrs(n int, err error) (results []RedisResult) { _ = "STUB: not implemented"; return nil }

func (c *clusterClient) doCache(ctx context.Context, cmd Cacheable, ttl time.Duration) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (c *clusterClient) DoCache(ctx context.Context, cmd Cacheable, ttl time.Duration) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (c *clusterClient) askingMulti(cc conn, ctx context.Context, multi []Completed) *redisresults {
	_ = "STUB: not implemented"
	return nil
}

func (c *clusterClient) askingMultiCache(cc conn, ctx context.Context, multi []CacheableTTL) *redisresults {
	_ = "STUB: not implemented"
	return nil
}

// check exec command error only

// if {Cmd} get a RedisError

func (c *clusterClient) _pickMultiCache(multi []CacheableTTL) *connretrycache {
	_ = "STUB: not implemented"
	return nil
}

func (c *clusterClient) pickMultiCache(ctx context.Context, multi []CacheableTTL) (*connretrycache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clusterClient) resultcachefn(
	ctx context.Context, results *redisresults, retries *connretrycache, mu *sync.Mutex, cc conn, cIndexes []int, commands []CacheableTTL, resps []RedisResult, attempts int,
) (clean bool) {
	_ = "STUB: not implemented"
	return false
}

func (c *clusterClient) doretrycache(
	ctx context.Context, cc conn, results *redisresults, retries *connretrycache, re *retrycache, mu *sync.Mutex, wg *sync.WaitGroup, attempts int,
) {
	_ = "STUB: not implemented"
	return
}

func (c *clusterClient) DoMultiCache(ctx context.Context, multi ...CacheableTTL) []RedisResult {
	_ = "STUB: not implemented"
	return nil
}

// Assume no retry. Because a client retry flag can be set to false.

func (c *clusterClient) Receive(ctx context.Context, subscribe Completed, fn func(msg PubSubMessage)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *clusterClient) DoStream(ctx context.Context, cmd Completed) RedisResultStream {
	_ = "STUB: not implemented"
	return *new(RedisResultStream)
}

func (c *clusterClient) DoMultiStream(ctx context.Context, multi ...Completed) MultiRedisResultStream {
	_ = "STUB: not implemented"
	return *new(MultiRedisResultStream)
}

func (c *clusterClient) Dedicated(fn func(DedicatedClient) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *clusterClient) Dedicate() (DedicatedClient, func()) {
	_ = "STUB: not implemented"
	return *new(DedicatedClient), nil
}

func (c *clusterClient) Nodes() map[string]Client { _ = "STUB: not implemented"; return nil }

func (c *clusterClient) Mode() ClientMode { _ = "STUB: not implemented"; return *new(ClientMode) }

func (c *clusterClient) Close() { _ = "STUB: not implemented"; return }

func (c *clusterClient) shouldRefreshRetry(err error, ctx context.Context) (addr string, mode RedirectMode) {
	_ = "STUB: not implemented"
	return "", *new(RedirectMode)
}

type dedicatedClusterClient struct {
	conn         conn
	wire         wire
	retryHandler retryHandler
	client       *clusterClient
	pshks        *pshks
	mu           sync.Mutex
	cmd          Builder
	slot         uint16
	retry        bool
	mark         bool
}

func (c *dedicatedClusterClient) acquire(ctx context.Context, slot uint16) (wire wire, err error) {
	_ = "STUB: not implemented"
	return *new(wire), nil
}

func (c *dedicatedClusterClient) release() { _ = "STUB: not implemented"; return }

func (c *dedicatedClusterClient) B() Builder { _ = "STUB: not implemented"; return *new(Builder) }

func (c *dedicatedClusterClient) Do(ctx context.Context, cmd Completed) (resp RedisResult) {
	_ = "STUB: not implemented"
	return *new(RedisResult)
}

func (c *dedicatedClusterClient) DoMulti(ctx context.Context, multi ...Completed) (resp []RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (c *dedicatedClusterClient) Receive(ctx context.Context, subscribe Completed, fn func(msg PubSubMessage)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *dedicatedClusterClient) SetPubSubHooks(hooks PubSubHooks) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dedicatedClusterClient) SetOnInvalidations(fn func([]RedisMessage)) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dedicatedClusterClient) Close() { _ = "STUB: not implemented"; return }

type RedirectMode int

const (
	RedirectNone RedirectMode = iota
	RedirectMove
	RedirectAsk
	RedirectRetry

	panicMsgCxSlot = "cross slot command in Dedicated is prohibited"
	panicMixCxSlot = "Mixing no-slot and cross slot commands in DoMulti is prohibited"
)
