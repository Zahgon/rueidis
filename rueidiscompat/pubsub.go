// Copyright (c) 2013 The github.com/go-redis/redis Authors.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
// * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
// * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package rueidiscompat

import (
	"context"
	"sync"
	"time"

	"github.com/redis/rueidis"
)

type PubSub interface {
	Close() error
	Subscribe(ctx context.Context, channels ...string) error
	PSubscribe(ctx context.Context, patterns ...string) error
	SSubscribe(ctx context.Context, channels ...string) error
	Unsubscribe(ctx context.Context, channels ...string) error
	PUnsubscribe(ctx context.Context, patterns ...string) error
	SUnsubscribe(ctx context.Context, channels ...string) error
	Ping(ctx context.Context, payload ...string) error
	ReceiveTimeout(ctx context.Context, timeout time.Duration) (any, error)
	Receive(ctx context.Context) (any, error)
	ReceiveMessage(ctx context.Context) (*Message, error)
	Channel(opts ...ChannelOption) <-chan *Message
	ChannelWithSubscriptions(opts ...ChannelOption) <-chan any
	String() string
}

type ChannelOption func(c *chopt)

type chopt struct {
	chanSize int
}

// WithChannelSize specifies the Go chan size that is used to buffer incoming messages.
// The default is 1000 messages.
func WithChannelSize(size int) ChannelOption { _ = "STUB: not implemented"; return *new(ChannelOption) }

// WithChannelHealthCheckInterval is an empty ChannelOption to keep compatibility
func WithChannelHealthCheckInterval(_ time.Duration) ChannelOption {
	_ = "STUB: not implemented"
	return *

	// WithChannelSendTimeout is an empty ChannelOption to keep compatibility
	new(ChannelOption)
}

func WithChannelSendTimeout(_ time.Duration) ChannelOption {
	_ = "STUB: not implemented"
	return *

	// Subscription received after a successful subscription to a channel.
	new(ChannelOption)
}

type Subscription struct {
	// Can be "subscribe", "unsubscribe", "psubscribe" or "punsubscribe".
	Kind string
	// Channel name we have subscribed to.
	Channel string
	// Number of channels we are currently subscribed to.
	Count int
}

func (m *Subscription) String() string { _ = "STUB: not implemented"; return "" }

// Message received as a result of a PUBLISH command issued by another client.
type Message struct {
	Channel      string
	Pattern      string
	Payload      string
	PayloadSlice []string
}

func (m *Message) String() string { _ = "STUB: not implemented"; return "" }

func newPubSub(client rueidis.Client) *pubsub { _ = "STUB: not implemented"; return nil }

type pubsub struct {
	rc      rueidis.Client
	mc      rueidis.DedicatedClient
	mcancel func()

	channels  map[string]bool
	patterns  map[string]bool
	schannels map[string]bool

	allCh chan any
	msgCh chan *Message
	mu    sync.Mutex
}

func (p *pubsub) mconn() rueidis.DedicatedClient {
	_ = "STUB: not implemented"
	return *new(rueidis.DedicatedClient)
}

func (p *pubsub) Close() error { _ = "STUB: not implemented"; return nil }

func (p *pubsub) Subscribe(ctx context.Context, channels ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pubsub) PSubscribe(ctx context.Context, patterns ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pubsub) SSubscribe(ctx context.Context, channels ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pubsub) Unsubscribe(ctx context.Context, channels ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pubsub) PUnsubscribe(ctx context.Context, patterns ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pubsub) SUnsubscribe(ctx context.Context, channels ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pubsub) Ping(_ context.Context, _ ...string) error {
	_ = "STUB: not implemented"
	// we already ping the connection periodically by default
	return nil
}

func (p *pubsub) reset() { _ = "STUB: not implemented"; return }

func (p *pubsub) resubscribe(ctx context.Context) rueidis.DedicatedClient {
	_ = "STUB: not implemented"
	return *new(rueidis.DedicatedClient)
}

func (p *pubsub) ReceiveTimeout(ctx context.Context, timeout time.Duration) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (p *pubsub) Receive(_ context.Context) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (p *pubsub) ReceiveMessage(_ context.Context) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *pubsub) Channel(opts ...ChannelOption) <-chan *Message {
	_ = "STUB: not implemented"
	return nil
}

func (p *pubsub) ChannelWithSubscriptions(opts ...ChannelOption) <-chan any {
	_ = "STUB: not implemented"
	return nil
}

func (p *pubsub) String() string { _ = "STUB: not implemented"; return "" }

func mapKeys(m map[string]bool) []string { _ = "STUB: not implemented"; return nil }
