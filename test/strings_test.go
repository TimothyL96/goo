package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestSplit(t *testing.T) {
	toSplit := "a.b.c"
	want := []string{"a", "b", "c"}
	gString := goo.FromString(toSplit)
	get := gString.Split(".")

	if len(want) != len(get) {
		t.Errorf("Split() Size of list %d; want %d", len(get), len(want))
	}

	for k := range want {
		if want[k] != get[k].ToString() {
			t.Errorf("Split(%s) of index %d = %s; want %s", gString, k, get[k], want[k])
		}
	}
}

func TestToUpper(t *testing.T) {
	want := "ABC"
	gString := goo.FromString("abc")
	get := gString.ToUpper()

	if get.ToString() != want {
		t.Errorf("ToUpper() = %s; want %s", get, want)
	}
}

func TestToLower(t *testing.T) {
	want := "abc"
	gString := goo.FromString("ABC")
	get := gString.ToLower()

	if get.ToString() != want {
		t.Errorf("ToLower() = %s; want %s", get, want)
	}
}

func TestTrim(t *testing.T) {
	want := "ABC"
	gString := goo.FromString(",.ABC.,")
	get := gString.Trim(",.")

	if get.ToString() != want {
		t.Errorf("Trim(%s) = %s; want %s", gString, get, want)
	}
}

func TestLength(t *testing.T) {
	want := 3
	gString := goo.FromString("abc")
	get := gString.Length()

	if get != want {
		t.Errorf("Length() = %d; want %d", get, want)
	}
}
