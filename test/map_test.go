package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestFromMap(t *testing.T) {
	want := make(map[int]int)
	want[0] = 1
	want[1] = 1
	want[2] = 3

	get := goo.FromMap(want)

	for k := range get {
		if _, ok := get[k]; !ok {
			t.Errorf("FromMap(%v) = %v; want %v", want, get, want)
		}
	}
}

func TestNewMap(t *testing.T) {
	want := make(map[int]int)
	want[1] = 1

	get := goo.NewMap[int, int](1)
	get[1] = 1

	for k := range get {
		if _, ok := get[k]; !ok {
			t.Errorf("NewMap() = %v; want %v", get, want)
		}
	}
}

func TestToMap(t *testing.T) {
	want := make(map[int]int)
	want[0] = 1
	want[1] = 1
	want[2] = 3

	gBool := goo.Map[int, int](want)
	get := gBool.ToMap()

	for k := range get {
		if _, ok := get[k]; !ok {
			t.Errorf("ToMap(%v) = %v; want %v", want, get, want)
		}
	}
}
