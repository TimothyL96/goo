package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestFromInt(t *testing.T) {
	want := 123
	get := goo.FromInt(want)

	if get.ToInt() != want {
		t.Errorf("FromInt(%d) = %d; want %d", want, get, want)
	}
}

func TestToInt(t *testing.T) {
	want := 123
	gInt := goo.Int(want)
	get := gInt.ToInt()

	if get != want {
		t.Errorf("ToInt() = %d; want %d", get, want)
	}
}
