//go:build ignore

package main

import (
	"reflect"
)

func main() {
	types := []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()}

	Generator("string", types)
}
