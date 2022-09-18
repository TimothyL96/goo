package goo

import (
	"strings"
	"unsafe"
)

// Split will split the current string by the given token and return a slice of String
func (s String) Split(sep String) []String {
	strs := strings.Split(string(s), sep.ToString())
	return *(*[]String)(unsafe.Pointer(&strs))
}

// ToUpper will convert all characters of the current String to uppercase
func (s String) ToUpper() String {
	return String(strings.ToUpper(s.ToString()))
}

// ToLower will convert all characters of the current String to lowercase
func (s String) ToLower() String {
	return String(strings.ToLower(s.ToString()))
}

// Trim will remove the given characters from the String instance
func (s String) Trim(cutset String) String {
	return String(strings.Trim(s.ToString(), cutset.ToString()))
}

// Length will return builtin int type
func (n String) Length() int {
	return len(n)
}
