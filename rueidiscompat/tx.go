package rueidiscompat

import (
	"context"
	"errors"
	"time"

	"github.com/redis/rueidis"
)

var TxFailedErr = errors.New("redis: transaction failed")

var _ Pipeliner = (*TxPipeline)(nil)

type rePipeline = Pipeline

func newTxPipeline(real rueidis.Client) *TxPipeline { _ = "STUB: not implemented"; return nil }

type TxPipeline struct {
	*rePipeline
}

func (c *TxPipeline) Exec(ctx context.Context) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *TxPipeline) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *TxPipeline) Pipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

func (c *TxPipeline) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *TxPipeline) TxPipeline() Pipeliner { _ = "STUB: not implemented"; return *new(Pipeliner) }

var _ rueidis.Client = (*txproxy)(nil)

type txproxy struct {
	rueidis.CoreClient
}

func (p *txproxy) DoCache(_ context.Context, _ rueidis.Cacheable, _ time.Duration) (resp rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

func (p *txproxy) DoMultiCache(_ context.Context, _ ...rueidis.CacheableTTL) (resp []rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (p *txproxy) DoStream(_ context.Context, _ rueidis.Completed) rueidis.RedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResultStream)
}

func (p *txproxy) DoMultiStream(_ context.Context, _ ...rueidis.Completed) rueidis.MultiRedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.MultiRedisResultStream)
}

func (p *txproxy) Dedicated(_ func(rueidis.DedicatedClient) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *txproxy) Dedicate() (client rueidis.DedicatedClient, cancel func()) {
	_ = "STUB: not implemented"
	return *new(rueidis.DedicatedClient), nil
}

func (p *txproxy) Nodes() map[string]rueidis.Client { _ = "STUB: not implemented"; return nil }

func (p *txproxy) Mode() rueidis.ClientMode {
	_ = "STUB: not implemented"
	return *new(rueidis.ClientMode)
}

type Tx interface {
	CoreCmdable
	Watch(ctx context.Context, keys ...string) *StatusCmd
	Unwatch(ctx context.Context, keys ...string) *StatusCmd
	Close(ctx context.Context) error
}

func newTx(client rueidis.DedicatedClient, cancel func()) *tx {
	_ = "STUB: not implemented"
	return nil
}

type tx struct {
	CoreCmdable
	cancel func()
}

func (t *tx) Watch(ctx context.Context, keys ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (t *tx) Unwatch(ctx context.Context, _ ...string) *StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func (t *tx) Close(_ context.Context) error { _ = "STUB: not implemented"; return nil }
