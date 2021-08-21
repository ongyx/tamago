package tamago

var (
	Carry, HalfCarry, Negative, Zero uint8

	Carry     = 1 << 4
	HalfCarry = 1 << 5
	Negative  = 1 << 6
	Zero      = 1 << 7
)

// The flag register F stores info on the previous instruction that was executed.
// Only bits 4-7 are used, bits 0-2 are unused.
type Flags struct {
	reg *uint8
}

func NewFlags(reg *uint8) *Flags {
	return &Flags{reg: reg}
}

// Check if a flag is set in the flag register.
func (f *Flags) Has(flag uint8) {
	return (*f.reg & flag) != 0
}

// Set a flag in the flag register.
func (f *Flags) Set(flag uint8) {
	*f.reg |= flag
}

// Clear a flag in the flag register.
func (f *Flags) Clear(flag uint8) {
	*f.reg &^= flag
}

// Clear all flags.
func (f *Flags) ClearAll() {
	*f.reg = 0
}

// Clear all flags except for one.
func (f *Flags) ClearAllExcept(flag uint8) {
	*f.reg &= flag
}

// Increment a uint8 register and set flags as necessary.
func (f *Flags) inc(reg *uint8) {

	if (*reg & 0x0f) == 0x0f {
		// third bit is about to be carried over to fourth bit
		f.Set(HalfCarry)
	} else {
		f.Clear(HalfCarry)
	}

	*reg++

	f.setIfZero(*reg)
	f.Clear(Negative)

}

// Decrement a uint8 register and set flags as necessary.
func (f *Flags) dec(reg *uint8) {

	if (*reg & 0x0f) == 0 {
		f.Set(HalfCarry)
	} else {
		f.Clear(HalfCarry)
	}

	*reg--

	f.setIfZero(*reg)
	f.Set(Negative)

}

func (f *Flags) setIfZero(v uint8) {
	if v == 0 {
		f.Set(Zero)
	} else {
		f.Clear(Zero)
	}
}

func (f *Flags) setIfCarry(v uint8) {
	if v != 0 {
		f.Set(Carry)
	} else {
		f.Clear(Carry)
	}
}

func (f *Flags) setIfHalfCarry(v1, v2 uint) {
	if ((*reg & 0x0f) + (value & 0x0f)) > 0x0f {
		f.Set(HalfCarry)
	} else {
		f.Clear(HalfCarry)
	}
}
