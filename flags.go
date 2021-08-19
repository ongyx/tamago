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

// Check if a flag is set in the flag register.
func (flags *Flags) Has(flag uint8) {
	return (*flags.reg & flag) != 0
}

// Set a flag in the flag register.
func (flags *Flags) Set(flag uint8) {
	*flags.reg |= flag
}

// Clear a flag in the flag register.
func (flags *Flags) Clear(flag uint8) {
	*flags.reg &^= flag
}

// Clear all flags.
func (flags *Flags) ClearAll() {
	*flags.reg = 0
}

// Clear all flags except for one.
func (flags *Flags) ClearAllExcept(flag uint8) {
	*flags.reg &= flag
}

func (flags *Flags) setIfZero(v uint8) {
	if v == 0 {
		flags.Set(Zero)
	} else {
		flags.Clear(Zero)
	}
}

func (flags *Flags) setIfCarry(v uint8) {
	if v != 0 {
		flags.Set(Carry)
	} else {
		flags.Clear(Carry)
	}
}

func (flags *Flags) setIfHalfCarry(v1, v2 uint) {
	if ((*reg & 0x0f) + (value & 0x0f)) > 0x0f {
		flags.Set(HalfCarry)
	} else {
		flags.Clear(HalfCarry)
	}
}
