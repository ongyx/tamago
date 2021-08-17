package tamago

import (
	"strconv"
)

var Flag = struct {
	Carry, HalfCarry, Negative, Zero uint8
}{
	1 << 4,
	1 << 5,
	1 << 6,
	1 << 7,
}

func validify(flag uint8) {
	switch flag {
	case Flag.Carry, Flag.HalfCarry, Flag.Negative, Flag.Zero:
	default:
		panic("invalid flag " + strconv.Itoa(flag))
	}
}

type Register struct {
	A, B, C, D, E, F, H, L uint8
	SP, PC                 uint16
}

// Set a flag
func (r *Register) Setf(flag uint8) {
	validify(flag)

	r.F |= flag
}

// Clear a flag.
func (r *Register) Clearf(flag uint8) {
	validify(flag)

	r.F &^= flag
}

func (r *Register) setZeroIfNeeded(reg *uint8) {
	if *reg == 0 {
		r.Setf(Flag.Zero)
	} else {
		r.Clearf(Flag.Zero)
	}
}

func (r *Register) Increment(reg *uint8) {
	if (*reg & 0x0f) == 0x0f {
		// third bit is about to be carried over to fourth bit
		r.Setf(Flag.HalfCarry)
	} else {
		r.Clearf(Flag.HalfCarry)
	}

	*reg++

	r.setZeroIfNeeded(reg)
	r.Clearf(Flag.Negative)
}

func (r *Register) Decrement(reg *uint8) {
	if (*reg & 0x0f) == 0 {
		r.Setf(Flag.HalfCarry)
	} else {
		r.Clearf(Flag.HalfCarry)
	}

	*reg--

	r.setZeroIfNeeded(reg)
	r.Setf(Flag.Negative)
}

func (r *Register) Add(reg *uint8, value uint8) {
	result := uint16(*reg) + uint16(value)

	if (result & 0xff00) != 0 {
		// overflow above 0xff
		r.Setf(Flag.Carry)
	}

	// truncate all bits above 0xff
	*reg = uint8(result & 0xff)

	r.setZeroIfNeeded(reg)
	r.Clearf(Flag.Negative)
}

func (r *Register) AF() uint16 {
	return getShort(&r.A, &r.F)
}

func (r *Register) SetAF(value uint16) {
	setShort(&r.F, &r.A, value)
}

func (r *Register) BC() uint16 {
	return getShort(&r.C, &r.B)
}

func (r *Register) SetBC(value uint16) {
	setShort(&r.C, &r.B, value)
}

func (r *Register) DE() uint16 {
	return getShort(&r.E, &r.D)
}

func (r *Register) SetDE(value uint16) {
	setShort(&r.E, &r.D, value)
}

func (r *Register) HL() uint16 {
	return getShort(&r.L, &r.H)
}

func (r *Register) SetHL(value uint16) {
	setShort(&r.L, &r.H, value)
}

func (r *Register) increment(reg *uint8) {}
