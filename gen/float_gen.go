//go:build ignore

package main

import (
	"reflect"
)

func main() {
	types := []reflect.Type{reflect.TypeOf((*float32)(nil)).Elem(),
		reflect.TypeOf((*float64)(nil)).Elem(),
	}

	Generator("float", types)
}
