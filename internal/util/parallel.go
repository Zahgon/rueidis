package util

import (
	"sync"
)

func ParallelKeys[K comparable, V any](maxp int, p map[K]V, fn func(k K)) {
	_ = "STUB: not implemented"
	return
}

func ParallelVals[K comparable, V any](maxp int, p map[K]V, fn func(k V)) {
	_ = "STUB: not implemented"
	return
}

func worker[V any](wg *sync.WaitGroup, ch chan V, fn func(k V)) { _ = "STUB: not implemented"; return }

func closeThenParallel[V any](maxp int, ch chan V, fn func(k V)) { _ = "STUB: not implemented"; return }
