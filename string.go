package goo

import "strings"

type String string

func NewString(s string) String {
	return String(s)
}

func (s *String) ToString() string {
	return string(*s)
}

func (s *String) ToUpper() String {
	return String(strings.ToUpper(string(*s)))
}

func (s *String) ToLower() String {
	return String(strings.ToLower(string(*s)))
}
