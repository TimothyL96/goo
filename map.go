package goo

import "golang.org/x/exp/constraints"

type Map[T constraints.Ordered, V any] map[T]V

// FromMap converts builtin map of any time to goo.Map
func FromMap[T constraints.Ordered, V any](m map[T]V) Map[T, V] {
	return m
}

// ToMap converts goo.Map object to builtin map
func (m Map[T, V]) ToMap() map[T]V {
	return m
}
