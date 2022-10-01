package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestFromString(t *testing.T) {
	want := "abc"
	get := goo.FromString(want)

	if get.ToString() != want {
		t.Errorf("FromString(%s) = %s; want %s", want, get, want)
	}
}

func TestToString(t *testing.T) {
	want := "abc"
	get := goo.FromString("abc").ToString()

	if get != want {
		t.Errorf("ToString() = %s; want %s", get, want)
	}
}

func TestStringSlice(t *testing.T) {
	want := "bc"
	gString := goo.FromString("abc")
	get := gString[1:]

	if get.ToString() != want {
		t.Errorf("String Slice[1:] %s = %s; want %s", want, get, want)
	}
}
