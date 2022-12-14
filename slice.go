package goo

import (
	"log"
	"reflect"
)

type Slice[T any] []T

// FromSlice converts builtin slice of any time to goo.Slice object
func FromSlice[T any](slice []T) Slice[T] {
	return slice
}

// NewSlice creates a new slice
func NewSlice[T any](elements ...T) Slice[T] {
	return elements
}

// ToSlice converts goo.Slice object to builtin slice of any time
func (s Slice[T]) ToSlice() []T {
	return s
}

// Append returns a new goo.Slice object with the input elements appended at the end
// Not yet tested if type is a slice or map
func (s Slice[T]) Append(elements ...T) Slice[T] {
	return append(s, elements...)
}

// Length returns the builtin length of slice
func (s Slice[T]) Length() int {
	return len(s)
}

// Unique return unique values from a list
// Does not work for map, slice, function as it is shallow checking and
// they will always have different pointers in a slice
func (s Slice[T]) Unique() Slice[T] {
	if s.Length() <= 1 {
		return s
	}

	// Check with reflection package, but it is just converting pointer type
	var checkType T

	switch reflect.ValueOf(checkType).Kind() {
	case reflect.Slice, reflect.Map, reflect.Func:
		log.Println("Deep compare not supported. Returning the slice")
		return s
	}

	// Unique for comparable types: Integer, Float, String (constraints.Ordered)
	m := make(map[any]struct{})

	for k := range s {
		m[s[k]] = struct{}{}
	}

	newSlice := make(Slice[T], len(m))

	i := 0
	for k := range m {
		newSlice[i] = k.(T)
		i += 1
	}

	return newSlice
}

// ForEach loop through each element and perform an action with the given function
func (s Slice[T]) ForEach(exec func(T)) {
	for k := range s {
		exec(s[k])
	}
}
