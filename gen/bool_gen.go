//go:build ignore

package main

import (
	"reflect"
)

func main() {
	types := []reflect.Type{
		reflect.TypeOf((*bool)(nil)).Elem(),
	}

	Generator("bool", types)
}
