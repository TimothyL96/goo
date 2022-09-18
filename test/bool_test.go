package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestFromBool(t *testing.T) {
	want := true
	get := goo.FromBool(true)

	if get.ToBool() != want {
		t.Errorf("FromBool(%t) = %t; want %t", want, get, want)
	}
}

func TestToBool(t *testing.T) {
	want := true
	gBool := goo.Bool(want)
	get := gBool.ToBool()

	if get != want {
		t.Errorf("ToBool() = %t; want %t", get, want)
	}
}
