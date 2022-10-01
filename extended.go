package goo

import "golang.org/x/exp/constraints"

// Unique function that is not tied to a slice
func Unique[T constraints.Ordered](s Slice[T]) Slice[T] {
	return s.Unique()
}

// HasAnyPrefix returns true if the given string matched any of the prefixes
func HasAnyPrefix(s String, prefixes ...string) Bool {
	return s.HasAnyPrefix(prefixes...)
}

func ForEach[T any](s Slice[T], exec func(T)) {
	s.ForEach(exec)
}
