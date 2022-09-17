package strings

import "strings"

type String string

func NewString(s string) String {
	return (String)(s)
}

func (s *String) ToUpper() String {
	return (String)(strings.ToUpper((string)(*s)))
}
