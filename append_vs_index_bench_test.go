package bench_test

import "testing"

func BenchmarkFillSliceByAppend(b *testing.B) {
	n := 100
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, n)
		for j := 0; j < n; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkFillSliceByIndex(b *testing.B) {
	n := 100
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, n)
		for j := 0; j < n; j++ {
			s[j] = j
		}
	}
}
