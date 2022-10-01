package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestSliceUniqueIntFunction(t *testing.T) {
	want := []int{1, 2, 3}
	toGet := []int{1, 1, 1, 2, 3, 3}
	get := goo.Unique(toGet)

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueIntLengthZero(t *testing.T) {
	want := []int{}
	toGet := []int{}
	get := goo.Unique(toGet)

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueIntLengthOne(t *testing.T) {
	want := []int{1}
	toGet := []int{1}
	get := goo.Unique(toGet)

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueStringFunction(t *testing.T) {
	want := []goo.String{"a", "b", "c"}
	toGet := []goo.String{"a", "a", "b", "b", "b", "b", "c"}
	get := goo.Unique(toGet)

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestStringHasAnyPrefix(t *testing.T) {
	s := goo.FromString("-asdasd")
	pre := goo.NewSlice("=", "-", ":", "a")
	get := goo.HasAnyPrefix(s, pre...)
	want := true

	if get.ToBool() != want {
		t.Errorf("HasAnyPrefix(%s, %s) = %t; want %t", s, pre, get, want)
	}
}
