package rueidis

import (
	"context"
	"errors"
	"iter"
	"sync/atomic"
	"time"
)

// MGetCache is a helper that consults the client-side caches with multiple keys by grouping keys within the same slot into multiple GETs
func MGetCache(client Client, ctx context.Context, ttl time.Duration, keys []string) (ret map[string]RedisMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isCacheDisabled(client Client) bool { _ = "STUB: not implemented"; return false }

// MGet is a helper that consults the redis directly with multiple keys by grouping keys within the same slot into MGET or multiple GETs
func MGet(client Client, ctx context.Context, keys []string) (ret map[string]RedisMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MSet is a helper that consults the redis directly with multiple keys by grouping keys within the same slot into MSETs or multiple SETs
func MSet(client Client, ctx context.Context, kvs map[string]string) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

// MDel is a helper that consults the redis directly with multiple keys by grouping keys within the same slot into DELs
func MDel(client Client, ctx context.Context, keys []string) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

// MSetNX is a helper that consults the redis directly with multiple keys by grouping keys within the same slot into MSETNXs or multiple SETNXs
func MSetNX(client Client, ctx context.Context, kvs map[string]string) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

// JsonMGetCache is a helper that consults the client-side caches with multiple keys by grouping keys within the same slot into multiple JSON.GETs
func JsonMGetCache(client Client, ctx context.Context, ttl time.Duration, keys []string, path string) (ret map[string]RedisMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// JsonMGet is a helper that consults redis directly with multiple keys by grouping keys within the same slot into JSON.MGETs or multiple JSON.GETs
func JsonMGet(client Client, ctx context.Context, keys []string, path string) (ret map[string]RedisMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// JsonMSet is a helper that consults redis directly with multiple keys by grouping keys within the same slot into JSON.MSETs or multiple JSON.SETs
func JsonMSet(client Client, ctx context.Context, kvs map[string]string, path string) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

// DecodeSliceOfJSON is a helper that struct-scans each RedisMessage into dest, which must be a slice of the pointer.
func DecodeSliceOfJSON[T any](result RedisResult, dest *[]T) error {
	_ = "STUB: not implemented"
	return nil
}

func clientMGet(client Client, ctx context.Context, cmd Completed, keys []string) (ret map[string]RedisMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func clientMSet(client Client, ctx context.Context, mset string, kvs map[string]string, ret map[string]error) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

func clientJSONMSet(client Client, ctx context.Context, kvs map[string]string, path string, ret map[string]error) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

func clientMDel(client Client, ctx context.Context, keys []string) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

func doMultiCache(cc Client, ctx context.Context, buf *mgetcachecmds, keys []string) (ret map[string]RedisMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// doMultiSet runs DoMulti, recycles each Completed on success, and returns buf to
// mgetcmdsp when every result has no non-Redis error. If any non-Redis error
// occurs (e.g. context deadline), the auto-pipelining writer may still be
// reading buf.s, so buf is not Put back.
func doMultiSet(cc Client, ctx context.Context, buf *mgetcmds) (ret map[string]error) {
	_ = "STUB: not implemented"
	return nil
}

func arrayToKV(m map[string]RedisMessage, arr []RedisMessage, keys []string) map[string]RedisMessage {
	_ = "STUB: not implemented"
	return nil
}

func clusterMGet(client Client, ctx context.Context, keys []string) (ret map[string]RedisMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Map slot -> index in cmds.s

func clusterJsonMGet(client Client, ctx context.Context, keys []string, path string) (ret map[string]RedisMessage, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Map slot -> index in cmds.s

// ErrMSetNXNotSet is used in the MSetNX helper when the underlying MSETNX response is 0.
// Ref: https://redis.io/commands/msetnx/
var ErrMSetNXNotSet = errors.New("MSETNX: no key was set")

type Scanner struct {
	next func(cursor uint64) (ScanEntry, error)
	err  error
}

func NewScanner(next func(cursor uint64) (ScanEntry, error)) *Scanner {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scanner) scan() iter.Seq[[]string] { _ = "STUB: not implemented"; return nil }

func (s *Scanner) Iter() iter.Seq[string] { _ = "STUB: not implemented"; return nil }

func (s *Scanner) Iter2() iter.Seq2[string, string] { _ = "STUB: not implemented"; return nil }

func (s *Scanner) Err() error {
	_ = "STUB: not implemented"

	// PreferReplicaNodeSelector prioritizes reading from any replica using Round-Robin.
	// If no replicas are available, it falls back to the primary.
	return nil
}

func PreferReplicaNodeSelector() ReadNodeSelectorFunc {
	_ = "STUB: not implemented"
	return *new(ReadNodeSelectorFunc)
}

// AZAffinityNodeSelector prioritizes replicas in the same AZ using Round-Robin.
func AZAffinityNodeSelector(clientAZ string) ReadNodeSelectorFunc {
	_ = "STUB: not implemented"
	return *new(ReadNodeSelectorFunc)
}

// AZAffinityReplicasAndPrimaryNodeSelector prioritizes:
// 1. Same-AZ Replicas
// 2. Same-AZ Primary
// 3. Any Replica
// 4. Primary
func AZAffinityReplicasAndPrimaryNodeSelector(clientAZ string) ReadNodeSelectorFunc {
	_ = "STUB: not implemented"
	return *new(ReadNodeSelectorFunc)
}

// Same-AZ Replicas

// Same-AZ Primary

// Any Replica

// newAZSelector creates the internal selector closure with a specific start index.
func newAZSelector(clientAZ string, startIdx int) func(uint16, []NodeInfo) int {
	_ = "STUB: not implemented"
	return nil
}

// Round-Robin on Same-AZ Replicas

// Round-Robin on ALL available nodes

// pickAZ selects a node index from nodes[startIdx:] that matches the clientAZ.
func pickAZ(nodes []NodeInfo, clientAZ string, startIdx int, counter *atomic.Uint32) int {
	_ = "STUB: not implemented"
	return 0
}

// We cap the search at 255 nodes

// Round-Robin Selection
