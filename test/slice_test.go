package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestFromSlice1(t *testing.T) {
	want := []int{1, 2, 3}
	get := goo.FromSlice([]int{1, 2, 3})

	for k, v := range get.ToSlice() {
		if v != want[k] {
			t.Errorf("FromSlice(%+v) = %+v; want %+v", want, get, want)
		}
	}
}

func TestFromSlice2(t *testing.T) {
	want := []string{"a", "b", "c"}
	get := goo.FromSlice([]string{"a", "b", "c"})

	for k, v := range get.ToSlice() {
		if v != want[k] {
			t.Errorf("FromSlice(%+v) = %+v; want %+v", want, get, want)
		}
	}
}

func TestFromSlice3(t *testing.T) {
	want := []goo.String{"a", "b", "c"}
	get := goo.FromSlice([]goo.String{"a", "b", "c"})

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

func TestNewSlice(t *testing.T) {
	want := []int{1, 2, 3}
	get := goo.NewSlice(1, 2, 3)

	for k, v := range get {
		if v != want[k] {
			t.Errorf("NewSlice() = %+v; want %+v", get, want)
		}
	}
}

func TestNewSliceGoo(t *testing.T) {
	want := []goo.Int{1, 2, 3}
	get := goo.NewSlice(1, 2, 3)

	for k, v := range get {
		if v != want[k].ToInt() {
			t.Errorf("NewSlice() = %+v; want %+v", get, want)
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

func TestSliceLength(t *testing.T) {
	want := 3
	get := goo.FromSlice([]int{1, 2, 3}).Length()

	if get != want {
		t.Errorf("Length() = %d; want %v", get, want)
	}
}

func TestSliceUniqueInt(t *testing.T) {
	want := []int{1, 2, 3}
	toGet := []int{1, 1, 1, 2, 3, 3}
	get := goo.FromSlice(toGet).Unique()

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueIntGoo(t *testing.T) {
	want := []goo.Int{1, 2, 3}
	toGet := []goo.Int{1, 1, 1, 2, 3, 3}
	get := goo.FromSlice(toGet).Unique()

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueString(t *testing.T) {
	want := []string{"a", "b", "c"}
	toGet := []string{"a", "a", "b", "b", "b", "b", "c"}
	get := goo.FromSlice(toGet).Unique()

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueStringGoo(t *testing.T) {
	want := []goo.String{"a", "b", "c"}
	toGet := []goo.String{"a", "a", "b", "b", "b", "b", "c"}
	get := goo.FromSlice(toGet).Unique()

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueFloat(t *testing.T) {
	want := []float64{11.11, 22.22, 33.33}
	toGet := []float64{11.11, 11.11, 11.11, 22.22, 22.22, 33.33}
	get := goo.FromSlice(toGet).Unique()

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueFloatGoo(t *testing.T) {
	want := []goo.Float64{11.11, 22.22, 33.33}
	toGet := []goo.Float64{11.11, 11.11, 11.11, 22.22, 22.22, 33.33}
	get := goo.FromSlice(toGet).Unique()

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueSlice(t *testing.T) {
	a := []int{1}
	b := []int{1}
	c := []int{1}

	want := 7
	toGet := [][]int{a, a, a, a, b, b, c}
	get := goo.FromSlice(toGet)

	get = get.Unique()

	if len(get.ToSlice()) != want {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), want)
	}
}

func TestSliceUniqueFunc(t *testing.T) {
	a := func() {}
	b := func() {}
	c := func() {}

	want := 7
	toGet := []func(){a, a, a, a, b, b, c}
	get := goo.FromSlice(toGet)

	get = get.Unique()

	if len(get.ToSlice()) != want {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), want)
	}
}

func TestSliceUniqueMap(t *testing.T) {
	a := make(map[int]struct{})
	b := make(map[int]struct{})
	c := make(map[int]struct{})

	want := 7
	toGet := []map[int]struct{}{a, a, a, a, b, b, c}
	get := goo.FromSlice(toGet)

	get = get.Unique()

	if len(get.ToSlice()) != want {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), want)
	}
}

func TestSliceUniqueIntFunction(t *testing.T) {
	want := []int{1, 2, 3}
	toGet := []int{1, 1, 1, 2, 3, 3}
	get := goo.Unique(toGet)

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueIntLengthZero(t *testing.T) {
	want := []int{}
	toGet := []int{}
	get := goo.Unique(toGet)

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueIntLengthOne(t *testing.T) {
	want := []int{1}
	toGet := []int{1}
	get := goo.Unique(toGet)

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}

func TestSliceUniqueStringFunction(t *testing.T) {
	want := []goo.String{"a", "b", "c"}
	toGet := []goo.String{"a", "a", "b", "b", "b", "b", "c"}
	get := goo.Unique(toGet)

	if len(get.ToSlice()) != len(want) {
		t.Errorf("Unique(%+v) = len %d; want %v", toGet, len(get), len(want))
	}
}
