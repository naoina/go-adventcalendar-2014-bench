package bench_test

import (
	"regexp"
	"testing"
)

func BenchmarkStringMatchWithRegexp(b *testing.B) {
	s := "0xDeadBeef"
	re := regexp.MustCompile(`^0[xX][0-9A-Fa-f]+$`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		re.MatchString(s)
	}
}

func BenchmarkStringMatchWithoutRegexp(b *testing.B) {
	s := "0xDeadBeef"
	isHexString := func(s string) bool {
		if len(s) < 3 || s[0] != '0' || s[1] != 'x' && s[1] != 'X' {
			return false
		}
		for _, c := range s[2:] {
			if c < '0' || '9' < c && c < 'A' || 'F' < c && c < 'a' || 'f' < c {
				return false
			}
		}
		return true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isHexString(s)
	}
}
