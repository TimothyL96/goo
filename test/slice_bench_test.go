package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func BenchmarkAppend(b *testing.B) {
	x := goo.FromSlice([]int{1, 2, 3})

	for i := 0; i < b.N; i++ {
		_ = x.Append(4, 5, 6)
	}
}

func BenchmarkStdAppend(b *testing.B) {
	x := []int{1, 2, 3}

	for i := 0; i < b.N; i++ {
		_ = append(x, 4, 5, 6)
	}
}
