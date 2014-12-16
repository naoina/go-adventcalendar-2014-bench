package bench_test

import "testing"

func BenchmarkLenDirect(b *testing.B) {
	n := 1000
	s := make([]string, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s); j++ {
			_ = s[j]
		}
	}
}

func BenchmarkLenCached(b *testing.B) {
	n := 1000
	s := make([]string, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j, length := 0, len(s); j < length; j++ {
			_ = s[j]
		}
	}
}
