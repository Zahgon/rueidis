package mock

import (
	"bufio"
	"bytes"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/rueidis"
)

func Result(val rueidis.RedisMessage) rueidis.RedisResult {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

func ErrorResult(err error) rueidis.RedisResult {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

func RedisString(v string) rueidis.RedisMessage {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisMessage)
}

func RedisBlobString(v string) rueidis.RedisMessage {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisMessage)
}

func RedisError(v string) rueidis.RedisMessage {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisMessage)
}

func RedisInt64(v int64) rueidis.RedisMessage {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisMessage)
}

func RedisFloat64(v float64) rueidis.RedisMessage {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisMessage)
}

func RedisBool(v bool) rueidis.RedisMessage {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisMessage)
}

func RedisNil() rueidis.RedisMessage { _ = "STUB: not implemented"; return *new(rueidis.RedisMessage) }

func RedisArray(values ...rueidis.RedisMessage) rueidis.RedisMessage {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisMessage)
}

func RedisMap(kv map[string]rueidis.RedisMessage) rueidis.RedisMessage {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisMessage)
}

func serialize(m message, buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func RedisResultStreamError(err error) rueidis.RedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResultStream)
}

func RedisResultStream(ms ...rueidis.RedisMessage) rueidis.RedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResultStream)
}

func MultiRedisResultStream(ms ...rueidis.RedisMessage) rueidis.MultiRedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.MultiRedisResultStream)
}

func MultiRedisResultStreamError(err error) rueidis.RedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResultStream)
}

type message struct {
	attrs   *rueidis.RedisMessage
	bytes   *byte
	array   *rueidis.RedisMessage
	integer int64
	typ     byte
	ttl     [7]byte
}

func (m *message) string() string { _ = "STUB: not implemented"; return "" }

func (m *message) values() []rueidis.RedisMessage { _ = "STUB: not implemented"; return nil }

func slicemsg(typ byte, values []rueidis.RedisMessage) message {
	_ = "STUB: not implemented"
	return *new(message)
}

func strmsg(typ byte, value string) message { _ = "STUB: not implemented"; return *new(message) }

type result struct {
	err error
	val rueidis.RedisMessage
}

type pool struct {
	dead    any
	cond    *sync.Cond
	timer   *time.Timer
	make    func() any
	list    []any
	cleanup time.Duration
	size    int
	minSize int
	cap     int
	down    bool
	timerOn bool
}

type pipe struct {
	conn            net.Conn
	clhks           atomic.Value // closed hook, invoked after the conn is closed
	queue           any
	cache           any
	pshks           atomic.Pointer[pshks] // pubsub hook, registered by the SetPubSubHooks
	error           atomic.Pointer[errs]
	r               *bufio.Reader
	w               *bufio.Writer
	close           chan struct{}
	onInvalidations func([]rueidis.RedisMessage)
	ssubs           *any // pubsub smessage subscriptions
	nsubs           *any // pubsub  message subscriptions
	psubs           *any // pubsub pmessage subscriptions
	r2p             *any
	pingTimer       *time.Timer // timer for background ping
	lftmTimer       *time.Timer // lifetime timer
	info            map[string]rueidis.RedisMessage
	timeout         time.Duration
	pinggap         time.Duration
	maxFlushDelay   time.Duration
	lftm            time.Duration // lifetime
	wrCounter       atomic.Uint64
	version         int32
	blcksig         int32
	state           int32
	bgState         int32
	r2ps            bool // identify this pipe is used for resp2 pubsub or not
	noNoDelay       bool
	optIn           bool
}

type stream struct {
	p *pool
	w *pipe
	e error
	n int
}

type errs struct{ error }

type pshks struct {
	hooks rueidis.PubSubHooks
	close chan error
}
