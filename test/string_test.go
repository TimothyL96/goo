package goo

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestNewString(t *testing.T) {
	want := "abc"
	get := goo.NewString(want)

	if get.ToString() != want {
		t.Errorf("NewString(%s) = %s; want %s", want, get, want)
	}
}
