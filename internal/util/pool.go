package util

import (
	"sync"
)

type Container interface {
	Capacity() int
	ResetLen(n int)
}

func NewPool[T Container](fn func(capacity int) T) *Pool[T] { _ = "STUB: not implemented"; return nil }

type Pool[T Container] struct {
	sp sync.Pool
	fn func(capacity int) T
}

func (p *Pool[T]) Get(length, capacity int) T { _ = "STUB: not implemented"; return *new(T) }

func (p *Pool[T]) Put(s T) { _ = "STUB: not implemented"; return }
