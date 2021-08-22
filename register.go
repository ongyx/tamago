package tamago

import (
	"math"
)

// The Game Boy has six 16-bit registers (AF, BC, DE, HL, SP, PC),
// of which the first four can be written to as general-purpose registers.
// They can be used seperately as eight 8-bit registers (A, B, C, D, E, F, H, L).
//
// NOTE:
// Even though the register is named AF, BC, etc.,
// the high byte is the first 8-bit register (i.e B) while the low byte is the second 8-bit register (i.e C).
type Register struct {
	Hi, Lo uint8
}

func NewRegister() *Register {
	return &Register{}
}

func (r *Register) Get() uint16 {
	return U16.From(r.Lo, r.Hi)
}

func (r *Register) Set(v uint16) {
	U16.To(&r.Lo, &r.Hi, v)
}

func (r *Register) Inc() {
	r.Set(r.Get() + 1)
}

func (r *Register) Dec() {
	r.Set(r.Get() - 1)
}

// Add a value to the register and set flags as necessary.
func (r *Register) Add(v uint16, fl *Flags) {
	result := uint(r.Get()) + uint(v)

	// The result of the bit clear will be more than 1 if there is overflow.
	fl.setIfCarry(result &^ math.MaxUint16)

	// Truncate any overflow.
	r.Set(uint16(result & math.MaxUint16))

	fl.setIfHalfCarry(uint(r.Get()), uint(v))
	fl.Clear(Negative)
}
