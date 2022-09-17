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

func TestToString(t *testing.T) {
	want := "abc"
	gString := (goo.String)(want)
	get := gString.ToString()

	if get != want {
		t.Errorf("NewString(%s) = %s; want %s", want, get, want)
	}
}

func TestToUpper(t *testing.T) {
	want := "ABC"
	gString := goo.NewString("abc")
	get := gString.ToUpper()

	if get.ToString() != want {
		t.Errorf("NewString(%s) = %s; want %s", want, get, want)
	}
}

func TestToLower(t *testing.T) {
	want := "abc"
	gString := goo.NewString("ABC")
	get := gString.ToLower()

	if get.ToString() != want {
		t.Errorf("NewString(%s) = %s; want %s", want, get, want)
	}
}

func TestSlice(t *testing.T) {
	want := "bc"
	gString := goo.NewString("abc")
	get := gString[1:]

	if get.ToString() != want {
		t.Errorf("NewString(%s) = %s; want %s", want, get, want)
	}
}
