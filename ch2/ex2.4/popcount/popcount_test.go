package popcount_test

import (
	"testing"

	"github.com/vmorris/mygopl/ch2/ex2.4/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount3(0x1234567890ABCDEF)
	}
}
