package core

import (
	"math"
	"math/bits"
)

// An arithmetic logic unit (ALU) provides arithmetic operations to the CPU, setting flags where appropriate.
type ALU struct {
	registers *Registers
}

// Creates a new ALU.
func NewALU(rs *Registers) ALU {
	return ALU{rs}
}

// Increments the byte value.
func (alu *ALU) Inc(v uint8) uint8 {
	r := v + 1

	alu.registers.F.Zero = r == 0
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = (v&0xF)+1 > 0xF

	return r
}

// Decrements the byte value.
func (alu *ALU) Dec(v uint8) uint8 {
	r := v - 1

	alu.registers.F.Zero = r == 0
	alu.registers.F.Subtract = true
	alu.registers.F.HalfCarry = int16(v&0xF)-1 < 0

	return r
}

// Adds an byte value to the register A. If carry is true, the carry bit is added.
func (alu *ALU) Add(v uint8, carry bool) {
	var cr uint8
	if carry && alu.registers.F.Carry {
		cr = 1
	}

	r := alu.registers.A + v + cr

	alu.registers.F.Zero = r == 0
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = (alu.registers.A&0xF)+(v&0xF)+cr > 0xF
	alu.registers.F.Carry = int16(alu.registers.A)+int16(v)+int16(cr) > math.MaxUint8

	alu.registers.A = r
}

// Subtracts an byte value from the register A. If carry is true, the carry bit is subtracted.
func (alu *ALU) Sub(v uint8, carry bool) {
	var cr uint8
	if carry && alu.registers.F.Carry {
		cr = 1
	}

	r := alu.registers.A - v - cr

	alu.registers.F.Zero = r == 0
	alu.registers.F.Subtract = true
	alu.registers.F.HalfCarry = int16(alu.registers.A&0xF)-int16(v&0xF)-int16(cr) < 0
	alu.registers.F.Carry = int16(alu.registers.A)-int16(v)-int16(cr) < 0

	alu.registers.A = v
}

// ANDs an byte value with the register A.
func (alu *ALU) And(v uint8) {
	result := alu.registers.A & v

	alu.registers.F.Zero = result == 0
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = true
	alu.registers.F.Carry = false

	alu.registers.A = result
}

// XORs an byte value with the register A.
func (alu *ALU) Xor(v uint8) {
	result := alu.registers.A ^ v

	alu.registers.F.Zero = result == 0
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = true
	alu.registers.F.Carry = false

	alu.registers.A = result
}

// ORs an byte value with the register A.
func (alu *ALU) Or(v uint8) {
	result := alu.registers.A | v

	alu.registers.F.Zero = result == 0
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = false
	alu.registers.F.Carry = false

	alu.registers.A = result
}

// Compares an byte value with the register A. This is equivalent to calling [ALU.Sub](v, false), but the result is discarded.
func (alu *ALU) Cp(v uint8) {
	a := alu.registers.A
	alu.Sub(v, false)
	alu.registers.A = a
}

// Rotates the byte value to the left, storing bit 7 in the carry flag.
// If throughCarry is true, the previous carry flag is stored in bit 0, otherwise bit 7 is stored in bit 0.
// If setZero is true, the zero flag is set.
func (alu *ALU) RotateLeft(v uint8, throughCarry bool, setZero bool) uint8 {
	msb := v >> 7

	var r uint8
	if throughCarry {
		var cr uint8
		if alu.registers.F.Carry {
			cr = 1
		}

		r = v<<1 | cr
	} else {
		r = bits.RotateLeft8(v, 1)
	}

	alu.registers.F.Zero = setZero && r == 0
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = false
	alu.registers.F.Carry = msb == 1

	return r
}

// Rotates the byte value to the right, storing bit 0 in the carry flag.
// If throughCarry is true, the previous carry flag is stored in bit 7, otherwise bit 0 is stored in bit 7.
// If setZero is true, the zero flag is set.
func (alu *ALU) RotateRight(v uint8, throughCarry bool, setZero bool) uint8 {
	lsb := v & 0x1

	var r uint8
	if throughCarry {
		var cr uint8
		if alu.registers.F.Carry {
			cr = 1
		}

		r = v>>1 | cr<<7
	} else {
		r = bits.RotateLeft8(v, -1)
	}

	alu.registers.F.Zero = setZero && r == 0
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = false
	alu.registers.F.Carry = lsb == 1

	return r
}

// Shifts the byte value to the left, clearing bit 0.
func (alu *ALU) ShiftLeft(v uint8) uint8 {
	msb := v >> 7
	r := v << 1

	alu.registers.F.Zero = r == 0
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = false
	alu.registers.F.Carry = msb == 1

	return r
}

// Shifts the byte value to the right, clearing bit 7.
// If preserveMsb is true, the value of bit 7 is preserved.
func (alu *ALU) ShiftRight(v uint8, preserveMsb bool) uint8 {
	lsb := v & 0x1
	r := v >> 1

	if preserveMsb {
		r |= v & (1 << 7)
	}

	alu.registers.F.Zero = r == 0
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = false
	alu.registers.F.Carry = lsb == 1

	return r
}

// Swaps the 4 upper bits with the 4 lower bits and vice versa.
func (alu *ALU) Swap(v uint8) uint8 {
	r := (v >> 4) | (v << 4)

	alu.registers.F.Zero = r == 0
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = false
	alu.registers.F.Carry = false

	return r
}

// Tests a bit by index within the byte. If the bit is not set, the zero flag is set to true.
func (alu *ALU) TestBit(v uint8, i uint8) {
	if i > 7 {
		panic("index must be in the range [0, 7]")
	}

	r := v & (1 << i)

	alu.registers.F.Zero = r == 0
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = true
}

// Clears a bit by index within the byte value.
func (alu *ALU) ClearBit(v uint8, i uint8) uint8 {
	if i > 7 {
		panic("index must be in the range [0, 7]")
	}

	return v &^ (1 << i)
}

// Sets a bit by index within the byte value.
func (alu *ALU) SetBit(v uint8, i uint8) uint8 {
	if i > 7 {
		panic("index must be in the range [0, 7]")
	}

	return v | (1 << i)
}

// Adjusts the value of register A to be in binary coded decimal (BCD).
func (alu *ALU) AdjustBCDA() {
	a := alu.registers.A
	r := a

	// Taking reference from https://rgbds.gbdev.io/docs/v1.0.3/gbz80.7#DAA...
	if alu.registers.F.Subtract {
		if alu.registers.F.HalfCarry {
			r -= 0x6
		}

		if alu.registers.F.Carry {
			r -= 0x60
		}
	} else {
		if alu.registers.F.HalfCarry || (a&0xF) > 0x9 {
			r += 0x6
		}

		if alu.registers.F.Carry || a > 0x99 {
			r += 0x60
			alu.registers.F.Carry = true
		}
	}

	alu.registers.F.Zero = r == 0
	alu.registers.F.HalfCarry = false

	alu.registers.A = r
}

// Takes the one's complement of register A and stores it back.
func (alu *ALU) ComplementA() {
	r := ^alu.registers.A

	alu.registers.F.Subtract = true
	alu.registers.F.HalfCarry = true

	alu.registers.A = r
}

// Sets the carry flag.
func (alu *ALU) SetCarry() {
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = false
	alu.registers.F.Carry = true
}

// Flips the carry flag.
func (alu *ALU) FlipCarry() {
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = false
	alu.registers.F.Carry = !alu.registers.F.Carry
}

// Adds a signed offset to the stack pointer.
func (alu *ALU) AddSP(offset int8) uint16 {
	sp := alu.registers.SP
	r := uint16(int16(sp) + int16(offset))

	alu.registers.F.Zero = false
	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = sp&0xF > r&0xF
	alu.registers.F.Carry = sp&0xFF > r&0xFF

	return r
}

// Adds a word value to register HL.
func (alu *ALU) AddHL(v uint16) {
	hl := alu.registers.HL()
	r := hl + v

	alu.registers.F.Subtract = false
	alu.registers.F.HalfCarry = hl&0xFF > r&0xFF
	alu.registers.F.Carry = hl > r

	alu.registers.SetHL(r)
}
