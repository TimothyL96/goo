package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestFromSlice(t *testing.T) {
	want := []int{1, 2, 3}
	get := goo.FromSlice([]int{1, 2, 3})

	for k, v := range get.ToSlice() {
		if v != want[k] {
			t.Errorf("FromSlice(%+v) = %+v; want %+v", want, get, want)
		}
	}
}

func TestToSlice(t *testing.T) {
	want := []int{1, 2, 3}
	get := goo.FromSlice([]int{1, 2, 3}).ToSlice()

	for k, v := range get {
		if v != want[k] {
			t.Errorf("ToSlice() = %+v; want %+v", get, want)
		}
	}
}

func TestAppend(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6}
	toAppend := []int{4, 5, 6}
	get := goo.FromSlice([]int{1, 2, 3}).Append(toAppend...)

	for k, v := range get.ToSlice() {
		if v != want[k] {
			t.Errorf("Append(%+v) = %+v; want %+v", want, get, want)
		}
	}
}
