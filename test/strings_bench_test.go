package test

import (
	"strings"
	"testing"

	"github.com/timothyl96/goo"
)

func BenchmarkStringSplit(b *testing.B) {
	str := goo.FromString("a.b.c")

	for i := 0; i < b.N; i++ {
		_ = str.Split(".")
	}
}

func BenchmarkStringStdSplit(b *testing.B) {
	str := "a.b.c"

	for i := 0; i < b.N; i++ {
		_ = strings.Split(str, ".")
	}
}

func BenchmarkStringToUpper(b *testing.B) {
	str := goo.FromString("abc")

	for i := 0; i < b.N; i++ {
		_ = str.ToUpper()
	}
}

func BenchmarkStringStdToUpper(b *testing.B) {
	str := "abc"

	for i := 0; i < b.N; i++ {
		_ = strings.ToUpper(str)
	}
}

func BenchmarkStringToLower(b *testing.B) {
	str := goo.FromString("abc")

	for i := 0; i < b.N; i++ {
		_ = str.ToLower()
	}
}

func BenchmarkStringStdToLower(b *testing.B) {
	str := "abc"

	for i := 0; i < b.N; i++ {
		_ = strings.ToLower(str)
	}
}

func BenchmarkStringTrim(b *testing.B) {
	gString := goo.FromString(",.ABC.,")

	for i := 0; i < b.N; i++ {
		_ = gString.Trim(",.")
	}
}

func BenchmarkStringStdTrim(b *testing.B) {
	str := ",.ABC.,"

	for i := 0; i < b.N; i++ {
		_ = strings.Trim(str, ",.")
	}
}

func BenchmarkStringLength(b *testing.B) {
	gString := goo.FromString(",.ABCDEFGOERINGROENGEROGNERIGOERGREGREGZXXZC ZX.,")

	for i := 0; i < b.N; i++ {
		_ = gString.Length()
	}
}

func BenchmarkStringStdLength(b *testing.B) {
	str := ",.ABCDEFGOERINGROENGEROGNERIGOERGREGREGZXXZC ZX.,"

	for i := 0; i < b.N; i++ {
		_ = len(str)
	}
}

func BenchmarkStringHasPrefix(b *testing.B) {
	gString := goo.FromString(",.ABC.,")

	for i := 0; i < b.N; i++ {
		_ = gString.HasPrefix(",.")
	}
}

func BenchmarkStringStdHasPrefix(b *testing.B) {
	str := ",.ABC.,"

	for i := 0; i < b.N; i++ {
		_ = strings.HasPrefix(str, ",.")
	}
}
