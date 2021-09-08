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
func (fl *Flags) Has(f flag) bool {
	return (fl.r.Lo & uint8(f)) != 0
}

// Set a flag in the register.
func (fl *Flags) Set(f flag) {
	fl.r.Lo |= uint8(f)
}

// Flip a flag in the register.
func (fl *Flags) Flip(f flag) {
	fl.r.Lo ^= uint8(f)
}

// Clear a flag in the register.
func (fl *Flags) Clear(f flag) {
	fl.r.Lo &^= uint8(f)
}

// Clear all flags.
func (fl *Flags) ClearAll() {
	fl.r.Lo = 0
}

// Clear all flags except for one.
func (fl *Flags) ClearAllExcept(f flag) {
	fl.r.Lo &= uint8(f)
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

	var i int

	switch iv := v.(type) {
	case uint8:
		i = int(iv)
	case uint16:
		i = int(iv)
	case int:
		i = iv
	}

	if i > 0 {
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
	fl.setIfHalfCarry((*reg & 0x0f) == 0x0F)

	*reg++

	fl.setIfZero(*reg)

	fl.Clear(Negative)
}

func (fl *Flags) dec(reg *uint8) {
	fl.setIfHalfCarry((*reg & 0x0f) != 0)

	*reg--

	fl.setIfZero(*reg)

	fl.Set(Negative)
}

func (fl *Flags) add(reg *uint8, v uint8) {
	result := uint16(*reg) + uint16(v)

	*reg = uint8(result & math.MaxUint8)

	fl.setIfCarry(result &^ math.MaxUint8)
	fl.setIfHalfCarry(((*reg & 0x0f) + (v & 0x0F)) > 0x0F)
	fl.setIfZero(*reg)

	fl.Clear(Negative)
}

func (fl *Flags) adc(v uint8) {
	a := &fl.r.Hi

	v += tobit(fl.Has(Carry))

	result := uint16(*a) + uint16(v)

	fl.Set(Negative)

	fl.setIfCarry(result &^ math.MaxUint8)
	fl.setIfZero(uint8(result & math.MaxUint8))

	fl.setIfHalfCarry(((v & 0x0f) + (*a & 0x0F)) > 0x0F)

	*a = uint8(result & math.MaxUint8)
}

func (fl *Flags) sub(v uint8) {
	a := &fl.r.Hi

	*a -= v

	fl.setIfCarry(tobit(v > *a))
	fl.setIfHalfCarry((v & 0x0f) > (*a & 0x0F))
	fl.setIfZero(*a)

	fl.Set(Negative)
}

func (fl *Flags) sbc(v uint8) {
	a := &fl.r.Hi

	v += tobit(fl.Has(Carry))

	fl.Set(Negative)

	fl.setIfCarry(tobit(v > *a))
	fl.setIfZero(tobit(v == *a))
	fl.setIfHalfCarry((v & 0x0f) > (*a & 0x0F))

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

	fl.setIfZero(*a)

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

	case (v & 0x0f) > (a & 0x0F):
		f |= HalfCarry

	}

	fl.Set(f)
}

/*
	0xcb-prefix bitwise functions
*/

func (fl *Flags) rlc(v *uint8) {
	carry := *v & 0x80

	fl.setIfCarry(carry)

	*v <<= 1
	*v += carry >> 7

	fl.setIfZero(*v)
	fl.Clear(HalfCarry | Negative)
}

func (fl *Flags) rrc(v *uint8) {
	carry := *v & 0x01

	// Shift right by one and put back the carry bit.
	*v >>= 1
	*v |= (carry << 7)

	fl.setIfCarry(carry)
	fl.setIfZero(*v)

	fl.Clear(HalfCarry | Negative)
}

func (fl *Flags) rl(v *uint8) {
	carry := tobit(fl.Has(Carry))

	// Check if the 7th bit is about to be discarded.
	fl.setIfCarry(*v & 0x80)

	*v <<= 1
	*v += carry

	fl.setIfZero(*v)

	fl.Clear(HalfCarry | Negative)
}

func (fl *Flags) rr(v *uint8) {
	// Shift v right and add back the carry if any.
	*v >>= 1
	*v |= tobit(fl.Has(Carry)) << 7

	fl.setIfCarry(*v & 0x01)
	fl.setIfZero(*v)

	fl.Clear(HalfCarry | Negative)
}

func (fl *Flags) sla(v *uint8) {
	fl.setIfCarry(*v & 0x80)

	*v <<= 1

	fl.setIfZero(*v)

	fl.Clear(HalfCarry | Negative)
}

func (fl *Flags) sra(v *uint8) {
	fl.setIfCarry(*v & 0x01)

	// only keep the 7th bit
	// i.e 10000010 -> 11000001
	*v = (*v & 0x80) | (*v >> 1)

	fl.setIfZero(*v)

	fl.Clear(HalfCarry | Negative)
}

func (fl *Flags) swap(v *uint8) {
	lo := *v & 0x0f
	hi := *v & 0xf0
	*v = lo<<4 | hi>>4

	fl.setIfZero(*v)
	fl.Clear(Carry | HalfCarry | Negative)
}

func (fl *Flags) srl(v *uint8) {
	fl.setIfCarry(*v & 0x01)

	*v >>= 1

	fl.setIfZero(*v)

	fl.Clear(HalfCarry | Negative)
}

func (fl *Flags) bit(pos uint8, v *uint8) {
	fl.setIfZero(*v & (1 << pos))

	fl.Set(HalfCarry)
	fl.Clear(Negative)
}

func (fl *Flags) res(pos uint8, v *uint8) {
	*v &^= (1 << pos)
}

func (fl *Flags) set(pos uint8, v *uint8) {
	*v |= (1 << pos)
}
