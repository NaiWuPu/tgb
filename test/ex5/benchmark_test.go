package ex5

import (
	"fmt"
	"testing"
)

func Benchmark_Alloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", i)
	}
}
