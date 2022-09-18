package goo

import "strconv"

func (i Int) Itoa() String {
	return FromString(strconv.Itoa(i.ToInt()))
}

func (s String) Atoi() (Int, error) {
	v, e := strconv.Atoi(s.ToString())
	return FromInt(v), e
}
