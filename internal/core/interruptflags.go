package core

const (
	interruptFlagVBlank = 1 << iota
	interruptFlagLCD
	interruptFlagTimer
	interruptFlagSerial
	interruptFlagJoypad
)

// Contains the flags set by the CPU regarding interrupts.
type InterruptFlags struct {
	VBlank bool
	LCD    bool
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
	if f.LCD {
		v |= interruptFlagLCD
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
	f.LCD = (v & interruptFlagLCD) != 0
	f.Timer = (v & interruptFlagTimer) != 0
	f.Serial = (v & interruptFlagSerial) != 0
	f.Joypad = (v & interruptFlagJoypad) != 0
}
