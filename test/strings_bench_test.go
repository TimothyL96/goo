package test

import (
	"strings"
	"testing"

	"github.com/timothyl96/goo"
)

func BenchmarkSplit(b *testing.B) {
	str := goo.FromString("a.b.c")

	for i := 0; i < b.N; i++ {
		_ = str.Split(".")
	}
}

func BenchmarkStdSplit(b *testing.B) {
	str := "a.b.c"

	for i := 0; i < b.N; i++ {
		_ = strings.Split(str, ".")
	}
}

func BenchmarkToUpper(b *testing.B) {
	str := goo.FromString("abc")

	for i := 0; i < b.N; i++ {
		_ = str.ToUpper()
	}
}

func BenchmarkStdToUpper(b *testing.B) {
	str := "abc"

	for i := 0; i < b.N; i++ {
		_ = strings.ToUpper(str)
	}
}

func BenchmarkToLower(b *testing.B) {
	str := goo.FromString("abc")

	for i := 0; i < b.N; i++ {
		_ = str.ToLower()
	}
}

func BenchmarkStdToLower(b *testing.B) {
	str := "abc"

	for i := 0; i < b.N; i++ {
		_ = strings.ToLower(str)
	}
}

func BenchmarkTrim(b *testing.B) {
	gString := goo.FromString(",.ABC.,")

	for i := 0; i < b.N; i++ {
		_ = gString.Trim(",.")
	}
}

func BenchmarkStdTrim(b *testing.B) {
	str := ",.ABC.,"

	for i := 0; i < b.N; i++ {
		_ = strings.Trim(str, ",.")
	}
}

func BenchmarkLength(b *testing.B) {
	gString := goo.FromString(",.ABC.,")

	for i := 0; i < b.N; i++ {
		_ = gString.Length()
	}
}

func BenchmarkStdLength(b *testing.B) {
	str := ",.ABC.,"

	for i := 0; i < b.N; i++ {
		_ = len(str)
	}
}
