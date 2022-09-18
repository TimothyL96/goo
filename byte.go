package goo

type Byte byte

// FromByte converts builtin byte to Byte
func FromByte(n byte) Byte {
	return Byte(n)
}

// ToUint8 converts instance to builtin uint8
func (n Byte) ToByte() byte {
	return byte(n)
}
