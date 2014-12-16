package bench_test

import "testing"

func BenchmarkMemAllocOndemand(b *testing.B) {
	n := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]string, 0)
		for j := 0; j < n; j++ {
			s = append(s, "alice")
		}
	}
}

func BenchmarkMemAllocAllBeforeUsing(b *testing.B) {
	n := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]string, 0, n)
		for j := 0; j < n; j++ {
			s = append(s, "alice")
		}
	}
}
