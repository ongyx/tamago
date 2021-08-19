package tamago

import "math"

// The Game Boy has six 16-bit registers (AF, BC, DE, HL, SP, PC),
// of which the first four can be written to as general-purpose registers.
// They can be used seperately as eight 8-bit registers (A, B, C, D, E, F, H, L).
//
// NOTE:
// Even though the register is named AF, BC, etc.,
// the low byte is the second 8-bit register (i.e C) while the high byte is the first 8-bit register (i.e B).
type Register struct {
	High, Low uint8
}

// Get the register as an unsigned 16-bit int.
func (r *Register) Value() uint16 {
	return Endian.Uint16([]uint8{r.Low, r.High})
}

// Set the register with an unsigned 16-bit int.
func (r *Register) SetValue(v uint16) {
	buf := []uint8{}
	Endian.PutUint16(buf, v)

	r.Low = buf[0]
	r.High = buf[1]
}

// Increment the register.
func (r *Register) Increment() {
	r.SetValue(r.Value() + 1)
}

// Decrement the register.
func (r *Register) Decrement() {
	r.SetValue(r.Value() - 1)
}

// Add a value to the register and set flags as necessary.
func (r *Register) Add(v uint16, flags *Flags) {
	original := uint(r.Value())
	value := uint(v)
	result := original + value

	// The result of the bit clear will be more than 1 if there is overflow.
	flags.setIfCarry(result &^ math.MaxUint16)

	// Truncate any overflow.
	r.SetValue(uint16(result & math.MaxUint16))

	flags.setIfHalfCarry(original, value)
	flags.Clear(Negative)
}

// Increment the low byte and set flags as necessary.
func (r *Register) IncrementLow(flags *Flags) {

	if (r.Low & 0x0f) == 0x0f {
		// third bit is about to be carried over to fourth bit
		flags.Set(HalfCarry)
	} else {
		flags.Clear(HalfCarry)
	}

	r.Low++

	flags.setIfZero(r.Low)
	flags.Clear(Negative)

}

// Increment the low byte and set flags as necessary.
func (r *Register) DecrementLow(flags *Flags) {

	if (*reg & 0x0f) == 0 {
		flags.Set(HalfCarry)
	} else {
		flags.Clear(HalfCarry)
	}

	r.Low--

	flags.setIfZero(r.Low)
	flags.Set(Negative)

}

// Increment the high byte and set flags as necessary.
func (r *Register) IncrementHigh(flags *Flags) {

	if (r.High & 0x0f) == 0x0f {
		// third bit is about to be carried over to fourth bit
		flags.Set(HalfCarry)
	} else {
		flags.Clear(HalfCarry)
	}

	r.High++

	flags.setIfZero(r.High)
	flags.Clear(Negative)

}

// Increment the high byte and set flags as necessary.
func (r *Register) DecrementHigh(flags *Flags) {

	if (r.High & 0x0f) == 0 {
		flags.Set(HalfCarry)
	} else {
		flags.Clear(HalfCarry)
	}

	r.High--

	flags.setIfZero(r.High)
	flags.Set(Negative)

}
