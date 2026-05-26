package rueidis

import (
	"context"
	"errors"
	"sync"
	"time"
)

// errAcquireComplete is a special error used to indicate that the Acquire operation has completed successfully
var errAcquireComplete = errors.New("acquire complete")

func newPool(cap int, dead wire, cleanup time.Duration, minSize int, makeFn func(context.Context) wire) *pool {
	_ = "STUB: not implemented"
	return nil
}

type pool struct {
	dead    wire
	cond    *sync.Cond
	timer   *time.Timer
	make    func(ctx context.Context) wire
	list    []wire
	cleanup time.Duration
	size    int
	minSize int
	cap     int
	down    bool
	timerOn bool
}

func (p *pool) Acquire(ctx context.Context) (v wire) {
	_ = "STUB: not implemented"

	// Set up ctx handling when waiting for an available connection
	return *new(wire)
}

// no need to broadcast if the poolCtx is cancelled explicitly.

// unlock before start to make a new wire
// allowing others to make wires concurrently instead of waiting in line

func (p *pool) Store(v wire) { _ = "STUB: not implemented"; return }

func (p *pool) Close() { _ = "STUB: not implemented"; return }

func (p *pool) startTimerIfNeeded() { _ = "STUB: not implemented"; return }

func (p *pool) removeIdleConns() { _ = "STUB: not implemented"; return }

func (p *pool) stopTimer() { _ = "STUB: not implemented"; return }
