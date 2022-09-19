package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestFromComplex64(t *testing.T) {
	var x float32 = 3.0
	var y float32 = 5.9
	var want complex64 = complex(x, y)
	get := goo.FromComplex64(want)

	if get.ToComplex64() != want {
		t.Errorf("FromComplex64(%f, %f) = %f; want %f", x, y, get, want)
	}
}

func TestToComplex64(t *testing.T) {
	var x float32 = 3.0
	var y float32 = 5.9
	var want complex64 = complex(x, y)
	gComplex := goo.Complex64(want)
	get := gComplex.ToComplex64()

	if get != want {
		t.Errorf("ToComplex64() = %f; want %f", get, want)
	}
}

func TestFromComplex128(t *testing.T) {
	x := 3.0
	y := 5.9
	var want complex128 = complex(x, y)
	get := goo.FromComplex128(want)

	if get.ToComplex128() != want {
		t.Errorf("FromComplex128(%f, %f) = %f; want %f", x, y, get, want)
	}
}

func TestToComplex128(t *testing.T) {
	x := 3.0
	y := 5.9
	var want complex128 = complex(x, y)
	gComplex := goo.Complex128(want)
	get := gComplex.ToComplex128()

	if get != want {
		t.Errorf("ToComplex128() = %f; want %f", get, want)
	}
}
