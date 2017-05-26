package popcount_test

import (
	"testing"

	"github.com/vmorris/mygopl/ch2/ex2.5/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount4(0x1234567890ABCDEF)
	}
}
