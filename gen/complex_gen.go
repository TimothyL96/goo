//go:build ignore

package main

import (
	"reflect"
)

func main() {
	types := []reflect.Type{
		reflect.TypeOf((*complex128)(nil)).Elem(),
		reflect.TypeOf((*complex64)(nil)).Elem(),
	}

	Generator("complex", types)
}
