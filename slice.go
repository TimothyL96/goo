package goo

type Slice[T any] []T

// FromSlice converts builtin slice of any time to goo.Slice object
func FromSlice[T any](slice []T) Slice[T] {
	return slice
}

// ToSlice converts goo.Slice object to builtin slice of any time
func (s Slice[T]) ToSlice() []T {
	return s
}

// Append returns a new goo.Slice object with the input elements appended at the end
func (s Slice[T]) Append(elements ...T) Slice[T] {
	return append(s, elements...)
}
