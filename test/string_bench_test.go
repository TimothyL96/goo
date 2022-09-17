package goo

import (
	"strings"
	"testing"

	"github.com/timothyl96/goo"
)

func BenchmarkToUpper(b *testing.B) {
	str := goo.NewString("abc")

	for i := 0; i < b.N; i++ {
		str.ToUpper()
	}
}

func BenchmarkStdToUpper(b *testing.B) {
	str := "abc"

	for i := 0; i < b.N; i++ {
		strings.ToUpper(str)
	}
}

func BenchmarkToLower(b *testing.B) {
	str := goo.NewString("abc")

	for i := 0; i < b.N; i++ {
		str = str.ToLower()
	}
}

func BenchmarkStdToLower(b *testing.B) {
	str := "abc"

	for i := 0; i < b.N; i++ {
		str = strings.ToLower(str)
	}
}
