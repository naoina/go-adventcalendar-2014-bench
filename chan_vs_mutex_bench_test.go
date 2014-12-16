package bench_test

import (
	"sync"
	"testing"
)

func BenchmarkExclusiveWithChannel(b *testing.B) {
	c := make(chan struct{}, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c <- struct{}{}
		// do something.
		<-c
	}
}

func BenchmarkExclusiveWithMutex(b *testing.B) {
	mu := new(sync.Mutex)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		// do something.
		mu.Unlock()
	}
}

func BenchmarkSyncWithChannel(b *testing.B) {
	n := 10
	c := make(chan struct{}, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			go func() {
				// do something.
				c <- struct{}{}
			}()
		}
		for j := 0; j < n; j++ {
			<-c
		}
	}
}

func BenchmarkSyncWithWaitGroup(b *testing.B) {
	n := 10
	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(n)
		for j := 0; j < n; j++ {
			go func() {
				// do something.
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
