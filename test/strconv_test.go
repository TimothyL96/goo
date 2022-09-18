package test

import (
	"testing"

	"github.com/timothyl96/goo"
)

func TestItoa(t *testing.T) {
	want := "1"
	gInt := goo.FromInt(1)
	get := gInt.Itoa()

	if want != get.ToString() {
		t.Errorf("Itoa() = %s; want %s", get, want)
	}
}

func TestAtoi(t *testing.T) {
	want := 1
	gString := goo.FromString("1")
	get, err := gString.Atoi()

	if err != nil {
		t.Errorf("Atoi() error: %e", err)
	}

	if want != get.ToInt() {
		t.Errorf("Atoi() = %d; want %d", get, want)
	}
}
