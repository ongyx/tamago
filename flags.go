package tamago

import (
	"math"
)

type f flag

const (
	Carry flag = 1 << (iota + 4)
	HalfCarry
	Negative
	Zero
)

// The flag register F (as part of the 16-bit pair AF) stores info on the previous instruction that was executed.
// Only bits 4-7 are used, bits 0-2 are unused.
type Flags struct {
	r *Register
}

func NewFlags(r *Register) *Flags {
	return &Flags{r: r}
}

// Check if a flag is set in the register.
func (fl *Flags) Has(f flag) {
	return (fl.r.Lo & f) != 0
}

// Set a flag in the register.
func (fl *Flags) Set(f flag) {
	fl.r.Lo |= f
}

// Flip a flag in the register.
func (fl *Flags) Flip(f flag) {
	fl.r.Lo ^= f
}

// Clear a flag in the register.
func (fl *Flags) Clear(f flag) {
	fl.r.Lo &^= f
}

// Clear all flags.
func (fl *Flags) ClearAll() {
	fl.r.Lo = 0
}

// Clear all flags except for one.
func (fl *Flags) ClearAllExcept(f flag) {
	fl.r.Lo &= f
}

func (fl *Flags) inc(reg *uint8) {

	if (*reg & 0x0f) == 0x0f {
		// third bit is about to be carried over to fourth bit
		fl.Set(HalfCarry)
	} else {
		fl.Clear(HalfCarry)
	}

	*reg++

	fl.setIfZero(*reg)
	fl.Clear(Negative)

}

func (fl *Flags) dec(reg *uint8) {

	if (*reg & 0x0f) == 0 {
		fl.Set(HalfCarry)
	} else {
		fl.Clear(HalfCarry)
	}

	*reg--

	fl.setIfZero(*reg)
	fl.Set(Negative)

}

func (fl *Flags) add(reg *uint8, v uint8) {
	result = uint16(*reg) + uint16(v)

	fl.setIfCarry(result &^ math.MaxUint8)

	*reg = uint8(result & math.MaxUint8)

	fl.setIfZero(*reg)
	fl.setIfHalfCarry(uint(*reg), uint(v))
	fl.Clear(Negative)
}

func (fl *Flags) sub(v uint8) {
	a := fl.r.Hi

	fl.Set(Negative)

	fl.setIfCarry(bit(v > a))
	fl.setIfHalfCarry(uint(v), uint(a))

	fl.r.Hi -= v

	fl.setIfZero(fl.r.Hi)
}

func (fl *Flags) adc(v uint8) {
	a := fl.r.Hi

	v += bit(fl.Has(Carry))

	result := uint16(a) + uint16(v)

	fl.setIfCarry(result &^ math.MaxUint8)
	fl.setIfZero(result)
	fl.setIfHalfCarry(uint(v), uint(a))

	fl.Set(Negative)

	fl.r.Hi = uint8(result & math.MaxUint8)
}

func (fl *Flags) sbc(v uint8) {
	a := fl.r.Hi

	v += bit(fl.Has(Carry))

	fl.Set(Negative)

	fl.setIfCarry(bit(v > a))
	fl.setIfZero(bit(v == a))
	fl.setIfHalfCarry(uint(v), uint(a))

	fl.r.Hi -= v
}

func (fl *Flags) setIfZero(v uint8) {
	if v == 0 {
		fl.Set(Zero)
	} else {
		fl.Clear(Zero)
	}
}

func (fl *Flags) setIfCarry(v uint8) {
	if v != 0 {
		fl.Set(Carry)
	} else {
		fl.Clear(Carry)
	}
}

func (fl *Flags) setIfHalfCarry(v1, v2 uint) {
	if ((v1 & 0x0f) + (v2 & 0x0f)) > 0x0f {
		fl.Set(HalfCarry)
	} else {
		fl.Clear(HalfCarry)
	}
}
