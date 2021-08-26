package tamago

// Value wraps a byte slice to pass to an instruction.
// This struct should not be passed by pointer since it is meant to be inmmutable.
type Value struct {
	raw []byte
}

func NewValue(buf []byte) Value {
	return Value{raw: buf}
}

func (v Value) S8() int8 {
	return int8(v.raw[0])
}

func (v Value) U8() uint8 {
	return uint8(v.raw[0])
}

func (v Value) U16() uint16 {
	return U16.From(v.raw[0], v.raw[1])
}
