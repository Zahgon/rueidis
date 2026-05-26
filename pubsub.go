package rueidis

import (
	"sync"
)

// PubSubMessage represents a pubsub message from redis
type PubSubMessage struct {
	// Pattern is only available with pmessage.
	Pattern string
	// Channel is the channel the message belongs to
	Channel string
	// Message is the message content
	Message string
}

// PubSubSubscription represent a pubsub "subscribe", "unsubscribe", "ssubscribe", "sunsubscribe", "psubscribe" or "punsubscribe" event.
type PubSubSubscription struct {
	// Kind is "subscribe", "unsubscribe", "ssubscribe", "sunsubscribe", "psubscribe" or "punsubscribe"
	Kind string
	// Channel is the event subject.
	Channel string
	// Count is the current number of subscriptions for a connection.
	Count int64
}

// PubSubHooks can be registered into DedicatedClient to process pubsub messages without using Client.Receive
type PubSubHooks struct {
	// OnMessage will be called when receiving "message" and "pmessage" event.
	OnMessage func(m PubSubMessage)
	// OnSubscription will be called when receiving "subscribe", "unsubscribe", "psubscribe" and "punsubscribe" event.
	OnSubscription func(s PubSubSubscription)
	// onInvalidations will be called when receiving "invalidate" event.
	onInvalidations func([]RedisMessage)
}

func (h *PubSubHooks) isZero() bool { _ = "STUB: not implemented"; return false }

func newSubs() *subs { _ = "STUB: not implemented"; return nil }

type subs struct {
	chs map[string]chs
	sub map[uint64]*sub
	cnt uint64
	mu  sync.RWMutex
}

type chs struct {
	sub map[uint64]*sub
}

type sub struct {
	ch chan PubSubMessage
	fn func(PubSubSubscription)
	cs []string
}

func (s *subs) Publish(channel string, msg PubSubMessage) { _ = "STUB: not implemented"; return }

func (s *subs) Subscribe(channels []string, fn func(PubSubSubscription)) (ch chan PubSubMessage, cancel func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *subs) remove(id uint64) { _ = "STUB: not implemented"; return }

func (s *subs) Confirm(sub PubSubSubscription) { _ = "STUB: not implemented"; return }

func (s *subs) Unsubscribe(sub PubSubSubscription) { _ = "STUB: not implemented"; return }

func (s *subs) Close() { _ = "STUB: not implemented"; return }
