// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2022-09-19 03:03:32.1658398 +0800 +08 m=+0.006479901
package goo

type String string

// FromString converts builtin type to Goo's string
func FromString(n string) String {
	return String(n)
}

// ToString converts instance to builtin string
func (n String) ToString() string {
	return string(n)
}
