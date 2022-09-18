package goo

import (
	"strings"
	"unsafe"
)

type String string

func FromString(s string) String {
	return String(s)
}

// ToString converts instance to builtin string
func (s String) ToString() string {
	return string(s)
}

func (s String) Split(sep string) []String {
	strs := strings.Split(string(s), sep)
	return *(*[]String)(unsafe.Pointer(&strs))
}

func (s String) ToUpper() String {
	return String(strings.ToUpper(s.ToString()))
}

func (s String) ToLower() String {
	return String(strings.ToLower(s.ToString()))
}
