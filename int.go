// Code generated by go generate; DO NOT EDIT.
// This file was generated at
// 2022-09-18 22:15:14.8458565 +0800 +08 m=+0.005949101
package goo


type Int int
type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Uint uint
type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64
type Uintptr uintptr

func FromInt(n int) Int {
	return Int(n)
}

// ToInt converts instance to builtin int
func (n Int) ToInt() int {
	return int(n)
}

func FromInt8(n int8) Int8 {
	return Int8(n)
}

// ToInt8 converts instance to builtin int8
func (n Int8) ToInt8() int8 {
	return int8(n)
}

func FromInt16(n int16) Int16 {
	return Int16(n)
}

// ToInt16 converts instance to builtin int16
func (n Int16) ToInt16() int16 {
	return int16(n)
}

func FromInt32(n int32) Int32 {
	return Int32(n)
}

// ToInt32 converts instance to builtin int32
func (n Int32) ToInt32() int32 {
	return int32(n)
}

func FromInt64(n int64) Int64 {
	return Int64(n)
}

// ToInt64 converts instance to builtin int64
func (n Int64) ToInt64() int64 {
	return int64(n)
}

func FromUint(n uint) Uint {
	return Uint(n)
}

// ToUint converts instance to builtin uint
func (n Uint) ToUint() uint {
	return uint(n)
}

func FromUint8(n uint8) Uint8 {
	return Uint8(n)
}

// ToUint8 converts instance to builtin uint8
func (n Uint8) ToUint8() uint8 {
	return uint8(n)
}

func FromUint16(n uint16) Uint16 {
	return Uint16(n)
}

// ToUint16 converts instance to builtin uint16
func (n Uint16) ToUint16() uint16 {
	return uint16(n)
}

func FromUint32(n uint32) Uint32 {
	return Uint32(n)
}

// ToUint32 converts instance to builtin uint32
func (n Uint32) ToUint32() uint32 {
	return uint32(n)
}

func FromUint64(n uint64) Uint64 {
	return Uint64(n)
}

// ToUint64 converts instance to builtin uint64
func (n Uint64) ToUint64() uint64 {
	return uint64(n)
}

func FromUintptr(n uintptr) Uintptr {
	return Uintptr(n)
}

// ToUintptr converts instance to builtin uintptr
func (n Uintptr) ToUintptr() uintptr {
	return uintptr(n)
}
