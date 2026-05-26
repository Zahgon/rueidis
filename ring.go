package rueidis

import (
	"context"
	"sync"

	"golang.org/x/sys/cpu"
)

type queue interface {
	PutOne(ctx context.Context, m Completed) (chan RedisResult, error)
	PutMulti(ctx context.Context, m []Completed, resps []RedisResult) (chan RedisResult, error)
	NextWriteCmd() (Completed, []Completed, chan RedisResult)
	WaitForWrite() (Completed, []Completed, chan RedisResult)
	NextResultCh() (Completed, []Completed, chan RedisResult, []RedisResult)
	FinishResult()
}

var _ queue = (*ring)(nil)

func newRing(factor int) *ring { _ = "STUB: not implemented"; return nil }

// this channel can't be buffered

type ring struct {
	resc  *sync.Cond
	store []node // store's size must be 2^N to work with the mask
	_     cpu.CacheLinePad
	write uint32
	_     cpu.CacheLinePad
	read1 uint32
	read2 uint32
	mask  uint32
}

type node struct {
	c1    *sync.Cond
	c2    *sync.Cond
	ch    chan RedisResult
	one   Completed
	multi []Completed
	resps []RedisResult
	mark  uint32
	slept bool
}

func (r *ring) PutOne(_ context.Context, m Completed) (chan RedisResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ring) PutMulti(_ context.Context, m []Completed, resps []RedisResult) (chan RedisResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NextWriteCmd should be only called by one dedicated thread
func (r *ring) NextWriteCmd() (one Completed, multi []Completed, ch chan RedisResult) {
	_ = "STUB: not implemented"
	return *new(Completed), nil, nil
}

// WaitForWrite should be only called by one dedicated thread
func (r *ring) WaitForWrite() (one Completed, multi []Completed, ch chan RedisResult) {
	_ = "STUB: not implemented"
	return *new(Completed), nil, nil
}

// c1 and c2 share the same mutex

// NextResultCh should be only called by one dedicated thread
func (r *ring) NextResultCh() (one Completed, multi []Completed, ch chan RedisResult, resps []RedisResult) {
	_ = "STUB: not implemented"
	return *new(Completed), nil, nil, nil
}

// FinishResult should be only called by one dedicated thread
func (r *ring) FinishResult() { _ = "STUB: not implemented"; return }
