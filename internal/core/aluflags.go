package core

const (
	aluFlagZero      = 1 << 7
	aluFlagSubtract  = 1 << 6
	aluFlagHalfCarry = 1 << 5
	aluFlagCarry     = 1 << 4
)

// Contains the flags set by the ALU during arithmetic operations.
// This is also known as the F register.
type ALUFlags struct {
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
func (f *ALUFlags) Encode() uint8 {
	var v uint8

	if f.Zero {
		v |= aluFlagZero
	}
	if f.Subtract {
		v |= aluFlagSubtract
	}
	if f.HalfCarry {
		v |= aluFlagHalfCarry
	}
	if f.Carry {
		v |= aluFlagCarry
	}

	return v
}

// Decodes the flags from a byte for writing to the F register.
// The lowest 4 bits are always ignored.
func (f *ALUFlags) Decode(v uint8) {
	f.Zero = (v & aluFlagZero) != 0
	f.Subtract = (v & aluFlagSubtract) != 0
	f.HalfCarry = (v & aluFlagHalfCarry) != 0
	f.Carry = (v & aluFlagCarry) != 0
}
