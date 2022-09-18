package goo

import "strconv"

// Itoa converts current Integer to ASCII string
func (i Int) Itoa() String {
	return FromString(strconv.Itoa(i.ToInt()))
}

// Atoi converts current String to Integer
func (s String) Atoi() (Int, error) {
	v, e := strconv.Atoi(s.ToString())
	return FromInt(v), e
}
