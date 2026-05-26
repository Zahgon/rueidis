package rueidis

import (
	"context"
	"sync"
	"time"
)

type call struct {
	ts time.Time
	ch chan struct{}
	cn int
	mu sync.Mutex
}

func (c *call) Do(ctx context.Context, fn func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *call) LazyDo(threshold time.Duration, fn func() error) { _ = "STUB: not implemented"; return }

func (c *call) do(ch chan struct{}, fn func() error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *call) suppressing() int { _ = "STUB: not implemented"; return 0 }
