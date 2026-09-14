package core

const (
	interruptFlagVBlank = 1 << iota
	interruptFlagStat
	interruptFlagTimer
	interruptFlagSerial
	interruptFlagJoypad
)

// Contains the flags set by the CPU regarding interrupts.
type InterruptFlags struct {
	VBlank bool
	Stat   bool
	Timer  bool
	Serial bool
	Joypad bool
}

// Encodes the flags to a byte for reading from the IE or IF register.
// The highest 3 bits are always 0.
func (f *InterruptFlags) Encode() uint8 {
	var v uint8

	if f.VBlank {
		v |= interruptFlagVBlank
	}
	if f.Stat {
		v |= interruptFlagStat
	}
	if f.Timer {
		v |= interruptFlagTimer
	}
	if f.Serial {
		v |= interruptFlagSerial
	}
	if f.Joypad {
		v |= interruptFlagJoypad
	}

	return v
}

// Decodes the flags from a byte for writing to the IE or IF register.
// The highest 3 bits are always ignored.
func (f *InterruptFlags) Decode(v uint8) {
	f.VBlank = (v & interruptFlagVBlank) != 0
	f.Stat = (v & interruptFlagStat) != 0
	f.Timer = (v & interruptFlagTimer) != 0
	f.Serial = (v & interruptFlagSerial) != 0
	f.Joypad = (v & interruptFlagJoypad) != 0
}
