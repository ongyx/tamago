package tamago

import (
	"math"
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

func (r *Register) setZeroIfNeeded(reg *uint) {
	if *reg == 0 {
		r.Setf(Flag.Zero)
	} else {
		r.Clearf(Flag.Zero)
	}
}

func (r *Register) setHalfCarryIfNeeded(reg *uint, value uint) {
	if ((*reg & 0x0f) + (value & 0x0f)) > 0x0f {
		r.Setf(Flag.HalfCarry)
	} else {
		r.Clearf(Flag.HalfCarry)
	}
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
	fullReg := uint(*reg)
	fullValue := uint(value)

	result := fullReg + fullValue

	if result > math.MaxUint8 {
		// overflow, keep only the lower 8 bits
		r.Setf(Flag.Carry)
		result = uint8(result & math.MaxUint8)

	} else {
		r.Clearf(Flag.Carry)
	}

	*reg = result

	r.setZeroIfNeeded(reg)
	r.setHalfCarryIfNeeded(fullReg, fullValue)
	r.Clearf(Flag.Negative)
}

func (r *Register) AddShort(reg *uint16, value uint16) {
	fullReg := uint(*reg)
	fullValue := uint(value)

	result := fullReg + fullValue

	if result > math.MaxUint16 {
		// overflow, keep only the lower 16 bits
		r.Setf(Flag.Carry)
		result = uint16(result & math.MaxUint16)

	} else {
		r.Clearf(Flag.Carry)
	}

	*reg = result

	r.setHalfCarryIfNeeded(reg, value)
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
