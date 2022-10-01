package goo

import (
	"golang.org/x/exp/constraints"
)

type Map[T constraints.Ordered, V any] map[T]V

// FromMap converts builtin map of any time to goo.Map
func FromMap[T constraints.Ordered, V any](m map[T]V) Map[T, V] {
	return m
}

// NewMap creates a new map with optional capacity
func NewMap[T constraints.Ordered, V any](size ...Int) Map[T, V] {
	if len(size) > 1 {
		panic("Wrong number of arguments in NewMap")
	}

	if len(size) == 1 {
		return make(Map[T, V], size[0])
	}

	return make(Map[T, V])
}

// ToMap converts goo.Map object to builtin map
func (m Map[T, V]) ToMap() map[T]V {
	return m
}
