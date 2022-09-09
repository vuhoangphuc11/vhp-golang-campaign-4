package pkg

import (
	f "fmt"
	"sync"
	t "time"
)

var mu sync.Mutex
var count = 0

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func Map[K, V any](s []K, transform func(K) V) []V {
	rs := make([]V, len(s))
	for i, v := range s {
		rs[i] = transform(v)
	}
	return rs
}

func Numbers() {
	for i := 1; i <= 6; i++ {
		t.Sleep(350 * t.Millisecond)
		f.Printf("%d ", i)
	}
}

func Alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		t.Sleep(400 * t.Millisecond)
		f.Printf("%c ", i)
	}
}
