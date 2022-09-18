package goo

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestNewInt(t *testing.T) {
	want := 123
	get := goo.FromInt(want)

	if get.ToInt() != want {
		t.Errorf("NewString(%d) = %d; want %d", want, get, want)
	}
}

func TestToInt(t *testing.T) {
	want := 123
	gInt := goo.Int(want)
	get := gInt.ToInt()

	if get != want {
		t.Errorf("NewString(%d) = %d; want %d", want, get, want)
	}
}
