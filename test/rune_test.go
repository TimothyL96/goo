package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestFromRune(t *testing.T) {
	want := 'a'
	get := goo.FromRune(want)

	if get.ToRune() != want {
		t.Errorf("FromRune(%U) = %U; want %U", want, get, want)
	}
}

func TestToRune(t *testing.T) {
	want := 'a'
	gBool := goo.Rune(want)
	get := gBool.ToRune()

	if get != want {
		t.Errorf("ToRune() = %U; want %U", get, want)
	}
}
