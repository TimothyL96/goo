package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func BenchmarkItoa(b *testing.B) {
	gInt := goo.FromInt(1)

	for i := 0; i < b.N; i++ {
		_ = gInt.Itoa()
	}
}

func BenchmarkAtoi(b *testing.B) {
	gString := goo.FromString("1")

	for i := 0; i < b.N; i++ {
		_, _ = gString.Atoi()
	}
}
