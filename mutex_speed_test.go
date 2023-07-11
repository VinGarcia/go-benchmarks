package main

import (
	"sync"
	"testing"
)

var someValue int

func BenchmarkMutex(b *testing.B) {
	mux := &sync.Mutex{}
	b.Run("using a mutex", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			mux.Lock()
			someValue = n
			mux.Unlock()
		}
	})

	b.Run("not using a mutex", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			someValue = n
		}
	})
}
