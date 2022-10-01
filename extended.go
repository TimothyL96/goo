package goo

import "golang.org/x/exp/constraints"

// Unique function that is not tied to a slice
func Unique[T constraints.Ordered](s Slice[T]) Slice[T] {
	return s.Unique()
}
