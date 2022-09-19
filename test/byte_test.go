package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestFromByte(t *testing.T) {
	var want byte = 97
	get := goo.FromByte(want)

	if get.ToByte() != want {
		t.Errorf("FromByte(%c) = %c; want %c", want, get, want)
	}
}

func TestToByte(t *testing.T) {
	var want byte = 97
	gBool := goo.Byte(want)
	get := gBool.ToByte()

	if get != want {
		t.Errorf("ToByte() = %c; want %c", get, want)
	}
}
