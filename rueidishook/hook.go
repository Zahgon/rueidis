package rueidishook

import (
	"context"
	"time"

	"github.com/redis/rueidis"
)

var _ rueidis.Client = (*hookclient)(nil)

// Hook allows user to intercept rueidis.Client by using WithHook
type Hook interface {
	Do(client rueidis.Client, ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult)
	DoMulti(client rueidis.Client, ctx context.Context, multi ...rueidis.Completed) (resps []rueidis.RedisResult)
	DoCache(client rueidis.Client, ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult)
	DoMultiCache(client rueidis.Client, ctx context.Context, multi ...rueidis.CacheableTTL) (resps []rueidis.RedisResult)
	Receive(client rueidis.Client, ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error)
	DoStream(client rueidis.Client, ctx context.Context, cmd rueidis.Completed) rueidis.RedisResultStream
	DoMultiStream(client rueidis.Client, ctx context.Context, multi ...rueidis.Completed) rueidis.MultiRedisResultStream
}

// WithHook wraps rueidis.Client with Hook and allows the user to intercept rueidis.Client
func WithHook(client rueidis.Client, hook Hook) rueidis.Client {
	_ = "STUB: not implemented"
	return *new(rueidis.Client)
}

type hookclient struct {
	client rueidis.Client
	hook   Hook
}

func (c *hookclient) B() rueidis.Builder { _ = "STUB: not implemented"; return *new(rueidis.Builder) }

func (c *hookclient) Do(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

func (c *hookclient) DoMulti(ctx context.Context, multi ...rueidis.Completed) (resp []rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (c *hookclient) DoCache(ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

func (c *hookclient) DoMultiCache(ctx context.Context, multi ...rueidis.CacheableTTL) (resps []rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (c *hookclient) DoStream(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResultStream)
}

func (c *hookclient) DoMultiStream(ctx context.Context, multi ...rueidis.Completed) rueidis.MultiRedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.MultiRedisResultStream)
}

func (c *hookclient) Dedicated(fn func(rueidis.DedicatedClient) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *hookclient) Dedicate() (rueidis.DedicatedClient, func()) {
	_ = "STUB: not implemented"
	return *new(rueidis.DedicatedClient), nil
}

func (c *hookclient) Receive(ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *hookclient) Nodes() map[string]rueidis.Client { _ = "STUB: not implemented"; return nil }

func (c *hookclient) Mode() rueidis.ClientMode {
	_ = "STUB: not implemented"
	return *new(rueidis.ClientMode)
}

func (c *hookclient) Close() { _ = "STUB: not implemented"; return }

var _ rueidis.DedicatedClient = (*dedicated)(nil)

type dedicated struct {
	client *extended
	hook   Hook
}

func (d *dedicated) B() rueidis.Builder { _ = "STUB: not implemented"; return *new(rueidis.Builder) }

func (d *dedicated) Do(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

func (d *dedicated) DoMulti(ctx context.Context, multi ...rueidis.Completed) (resp []rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) Receive(ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) SetPubSubHooks(hooks rueidis.PubSubHooks) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) SetOnInvalidations(fn func([]rueidis.RedisMessage)) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) Close() { _ = "STUB: not implemented"; return }

var _ rueidis.Client = (*extended)(nil)

type extended struct {
	rueidis.DedicatedClient
}

func (e *extended) DoCache(ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

func (e *extended) DoMultiCache(ctx context.Context, multi ...rueidis.CacheableTTL) (resp []rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (c *extended) DoStream(ctx context.Context, cmd rueidis.Completed) rueidis.RedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResultStream)
}

func (c *extended) DoMultiStream(ctx context.Context, multi ...rueidis.Completed) rueidis.MultiRedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.MultiRedisResultStream)
}

func (e *extended) Dedicated(fn func(rueidis.DedicatedClient) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *extended) Dedicate() (client rueidis.DedicatedClient, cancel func()) {
	_ = "STUB: not implemented"
	return *new(rueidis.DedicatedClient), nil
}

func (e *extended) Nodes() map[string]rueidis.Client { _ = "STUB: not implemented"; return nil }

func (e *extended) Mode() rueidis.ClientMode {
	_ = "STUB: not implemented"
	return *new(rueidis.ClientMode)
}

type result struct {
	err error
	val rueidis.RedisMessage
}

func NewErrorResult(err error) rueidis.RedisResult {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

type stream struct {
	p *int
	w *int
	e error
	n int
}

func NewErrorResultStream(err error) rueidis.RedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResultStream)
}
