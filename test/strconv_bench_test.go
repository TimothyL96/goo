package test

import (
	"strconv"
	"testing"

	"github.com/timothyl96/goo"
)

func BenchmarkItoa(b *testing.B) {
	gInt := goo.FromInt(1)

	for i := 0; i < b.N; i++ {
		_ = gInt.Itoa()
	}
}

func BenchmarkStdItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(1)
	}
}

func BenchmarkAtoi(b *testing.B) {
	gString := goo.FromString("1")

	for i := 0; i < b.N; i++ {
		_, _ = gString.Atoi()
	}
}

func BenchmarkStdAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = strconv.Atoi("1")
	}
}
