package goo

import (
	"strings"
	"unsafe"
)

func (s String) Split(sep String) []String {
	strs := strings.Split(string(s), sep.ToString())
	return *(*[]String)(unsafe.Pointer(&strs))
}

func (s String) ToUpper() String {
	return String(strings.ToUpper(s.ToString()))
}

func (s String) ToLower() String {
	return String(strings.ToLower(s.ToString()))
}

func (s String) Trim(cutset String) String {
	return String(strings.Trim(s.ToString(), cutset.ToString()))
}

// Length will return builtin int type
func (n String) Length() int {
	return len(n)
}
