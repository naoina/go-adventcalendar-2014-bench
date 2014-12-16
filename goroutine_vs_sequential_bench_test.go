package bench_test

import (
	"sync"
	"testing"
)

func BenchmarkGoroutine(b *testing.B) {
	n := 10
	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(n)
		for j := 0; j < n; j++ {
			go func() {
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSequential(b *testing.B) {
	n := 10
	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(n)
		for j := 0; j < n; j++ {
			func() {
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
