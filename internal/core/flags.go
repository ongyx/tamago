package core

const (
	bitZero      = 1 << 7
	bitSubtract  = 1 << 6
	bitHalfCarry = 1 << 5
	bitCarry     = 1 << 4
)

// Represents the flags set by the ALU during arithmetic operations.
// This is also known as the F register.
type Flags struct {
	// Was the result zero?
	Zero bool
	// Was the operation a subtraction?
	Subtract bool
	// Did the 3rd bit carry?
	HalfCarry bool
	// Did the 7th bit carry?
	Carry bool
}

// Encodes the flags to a byte for reading from the F register.
// The lowest 4 bits are always 0.
func (f *Flags) Encode() uint8 {
	var v uint8

	if f.Zero {
		v |= bitZero
	}
	if f.Subtract {
		v |= bitSubtract
	}
	if f.HalfCarry {
		v |= bitHalfCarry
	}
	if f.Carry {
		v |= bitCarry
	}

	return v
}

// Decodes the flags from a byte for writing to the F register.
// The lowest 4 bits are always ignored.
func (f *Flags) Decode(v uint8) {
	f.Zero = (v & bitZero) != 0
	f.Subtract = (v & bitSubtract) != 0
	f.HalfCarry = (v & bitHalfCarry) != 0
	f.Carry = (v & bitCarry) != 0
}
