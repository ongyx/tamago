package tamago

import (
	"math"
)

type flag uint8

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

/*
	Public functions
*/

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

/*
	Calculation flag helpers
*/

func (fl *Flags) setIfZero(v uint8) {
	if v == 0 {
		fl.Set(Zero)
	} else {
		fl.Clear(Zero)
	}
}

func (fl *Flags) setIfCarry(v interface{}) {
	if int(v) != 0 {
		fl.Set(Carry)
	} else {
		fl.Clear(Carry)
	}
}

func (fl *Flags) setIfHalfCarry(b bool) {
	if b {
		fl.Set(HalfCarry)
	} else {
		fl.Clear(HalfCarry)
	}
}

/*
	Arithmetic functions
*/

func (fl *Flags) inc(reg *uint8) {
	// third bit is about to be carried over to fourth bit
	fl.setIfHalfCarry((*reg & 0x0F) == 0x0F)

	*reg++

	fl.setIfZero(*reg)

	fl.Clear(Negative)
}

func (fl *Flags) dec(reg *uint8) {
	fl.setIfHalfCarry((*reg & 0x0F) != 0)

	*reg--

	fl.setIfZero(*reg)

	fl.Set(Negative)
}

func (fl *Flags) add(reg *uint8, v uint8) {
	result = uint16(*reg) + uint16(v)

	*reg = uint8(result & math.MaxUint8)

	fl.setIfCarry(result &^ math.MaxUint8)
	fl.setIfHalfCarry(((*reg & 0x0F) + (v & 0x0F)) > 0x0F)
	fl.setIfZero(*reg)

	fl.Clear(Negative)
}

func (fl *Flags) adc(v uint8) {
	a := &fl.r.Hi

	v += bit(fl.Has(Carry))

	result := uint16(*a) + uint16(v)

	fl.Set(Negative)

	fl.setIfCarry(result &^ math.MaxUint8)
	fl.setIfZero(result)

	fl.setIfHalfCarry(((v & 0x0F) + (*a & 0x0F)) > 0x0F)

	*a = uint8(result & math.MaxUint8)
}

func (fl *Flags) sub(v uint8) {
	a := &fl.r.Hi

	*a -= v

	fl.setIfCarry(bit(v > *a))
	fl.setIfHalfCarry((v & 0x0F) > (*a & 0x0F))
	fl.setIfZero(*a)

	fl.Set(Negative)
}

func (fl *Flags) sbc(v uint8) {
	a := &fl.r.Hi

	v += bit(fl.Has(Carry))

	fl.Set(Negative)

	fl.setIfCarry(bit(v > *a))
	fl.setIfZero(bit(v == *a))
	fl.setIfHalfCarry((v & 0x0F) > (*a & 0x0F))

	*a -= v
}

/*
	Bitwise functions
*/

func (fl *Flags) and(v uint8) {
	a := &fl.r.Hi

	*a &= v

	fl.setIfZero(*a)

	fl.Set(HalfCarry)
	fl.Clear(Carry | Negative)
}

func (fl *Flags) or(v uint8) {
	a := &fl.r.Hi

	*a |= v

	fl.setIfZero(*a)

	fl.Clear(Carry | HalfCarry | Negative)
}

func (fl *Flags) xor(v uint8) {
	a := &fl.r.Hi

	*a ^= v

	fl.setIfZero(a)

	fl.Clear(Carry | HalfCarry | Negative)
}

func (fl *Flags) cmp(v uint8) {
	var f flag

	fl.ClearAll()

	switch a := fl.r.Hi; true {

	case a == v:
		f |= Zero
		fallthrough

	case v > a:
		f |= Carry
		fallthrough

	case (v & 0x0F) > (a & 0x0F):
		f |= HalfCarry

	}

	fl.Set(f)
}
