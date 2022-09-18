package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestFromFloat32(t *testing.T) {
	want := float32(123.321)
	get := goo.FromFloat32(want)

	if get.ToFloat32() != want {
		t.Errorf("FromFloat32(%f) = %f; want %f", want, get, want)
	}
}

func TestToFloat32(t *testing.T) {
	want := float32(123.321)
	gInt := goo.Float32(want)
	get := gInt.ToFloat32()

	if get != want {
		t.Errorf("ToFloat32() = %f; want %f", get, want)
	}
}

func TestFromFloat64(t *testing.T) {
	want := 123.321
	get := goo.FromFloat64(want)

	if get.ToFloat64() != want {
		t.Errorf("FromFloat64(%f) = %f; want %f", want, get, want)
	}
}

func TestToFloat64(t *testing.T) {
	want := 123.321
	gInt := goo.Float64(want)
	get := gInt.ToFloat64()

	if get != want {
		t.Errorf("ToFloat64() = %f; want %f", get, want)
	}
}
