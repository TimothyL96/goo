package goo

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
		t.Errorf("ToString(%s) = %s; want %s", want, get, want)
	}
}

func TestSplit(t *testing.T) {
	toSplit := "a.b.c"
	want := []string{"a", "b", "c"}
	gString := goo.FromString(toSplit)
	get := gString.Split(".")

	if len(want) != len(get) {
		t.Errorf("Size of list %d; want %d", len(get), len(want))
	}

	for k := range want {
		if want[k] != get[k].ToString() {
			t.Errorf("Split(%s) of index %d = %s; want %s", toSplit, k, get[k], want[k])
		}
	}
}

func TestToUpper(t *testing.T) {
	want := "ABC"
	gString := goo.FromString("abc")
	get := gString.ToUpper()

	if get.ToString() != want {
		t.Errorf("ToUpper(%s) = %s; want %s", want, get, want)
	}
}

func TestToLower(t *testing.T) {
	want := "abc"
	gString := goo.FromString("ABC")
	get := gString.ToLower()

	if get.ToString() != want {
		t.Errorf("ToLower(%s) = %s; want %s", want, get, want)
	}
}

func TestSlice(t *testing.T) {
	want := "bc"
	gString := goo.FromString("abc")
	get := gString[1:]

	if get.ToString() != want {
		t.Errorf("Slice[1:] %s = %s; want %s", want, get, want)
	}
}
