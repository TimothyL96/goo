// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2022-09-18 20:34:47.3143093 +0800 +08 m=+0.006610901
package goo


type Float32 float32
type Float64 float64

func FromFloat32(n float32) Float32 {
	return Float32(n)
}

// ToFloat32 converts instance to builtin float32
func (n *Float32) ToFloat32() float32 {
	return float32(*n)
}

func FromFloat64(n float64) Float64 {
	return Float64(n)
}

// ToFloat64 converts instance to builtin float64
func (n *Float64) ToFloat64() float64 {
	return float64(*n)
}
