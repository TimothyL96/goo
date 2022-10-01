package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func BenchmarkSliceAppend(b *testing.B) {
	x := goo.FromSlice([]int{1, 2, 3})

	for i := 0; i < b.N; i++ {
		_ = x.Append(4, 5, 6)
	}
}

func BenchmarkSliceStdAppend(b *testing.B) {
	x := []int{1, 2, 3}

	for i := 0; i < b.N; i++ {
		_ = append(x, 4, 5, 6)
	}
}

func BenchmarkSliceLength(b *testing.B) {
	x := goo.FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})

	for i := 0; i < b.N; i++ {
		_ = x.Length()
	}
}

func BenchmarkSliceStdLength(b *testing.B) {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for i := 0; i < b.N; i++ {
		_ = len(x)
	}
}
