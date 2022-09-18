package goo

type Int int
type Int32 int32
type Int64 int64

func FromInt(n int) Int {
	return Int(n)
}

func FromInt64(n int64) Int64 {
	return Int64(n)
}

// ToInt converts instance to builtin int
func (n *Int) ToInt() int {
	return int(*n)
}
