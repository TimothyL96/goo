package goo

import (
	"strings"
	"unsafe"
)

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
