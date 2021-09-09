package tamago

import (
	"fmt"
	"io"
)

const debug = `
AF: %v
BC: %v
DE: %v
HL: %v
SP: %d
PC: %d
clock (t cycles): %d
stopped: %t
`

// State represents the current state of the emulation at some point in time.
type State struct {
	*MMU

	AF, BC, DE, HL *Register
	SP, PC         uint16

	fl    *Flags
	clock *Clock

	stopped bool
}

func NewState(rr Renderer) *State {
	s := &State{
		// The DMG bootrom assigns these values to the registers.
		// https://gbdev.io/pandocs/Power_Up_Sequence.html#cpu-registers
		AF: &Register{0x01, 0xb0},
		BC: &Register{0x00, 0x13},
		DE: &Register{0x00, 0xd8},
		HL: &Register{0x01, 0x4d},
		SP: 0xfffe,
		PC: 0x100,
	}

	s.MMU = NewMMU(rr)

	s.fl = NewFlags(s.AF)
	s.clock = NewClock()

	return s
}

// Jump the program counter to a relative offset from the current address.
func (s *State) Jump(offset int8) {
	if offset < 0 {
		s.PC -= uint16(-offset)
	} else {
		s.PC += uint16(offset)
	}
}

// Jump only if cond is met and advance the clock accordingly.
func (s *State) JumpIf(cond bool, v Value) {
	if cond {
		s.Jump(v.S8())
		s.clock.Step(3)
	} else {
		s.clock.Step(2)
	}
}

/*
	Stack functions
*/

// Push a value onto the stack.
func (s *State) Push(v uint16) {
	s.SP -= 2
	s.WriteShort(s.SP, v)
}

// Pop a value from the stack.
func (s *State) Pop() uint16 {
	v := s.ReadShort(s.SP)
	s.SP += 2

	return v
}

/*
	Misc functions
*/

// Show the contents of the registers and dump the contents of memory for debugging.
func (s *State) String() string {
	return fmt.Sprintf(debug, s.AF, s.BC, s.DE, s.HL, s.SP, s.PC, s.clock.t, s.stopped)
}

// This shadows the embedded MMU functions so the program counter can be set correctly.
func (s *State) LoadBootFrom(rom io.Reader) error {
	if err := s.MMU.LoadBootFrom(rom); err != nil {
		return err
	}

	s.PC = 0

	return nil
}

func (s *State) LoadBoot(rom string) error {
	if err := s.MMU.LoadBoot(rom); err != nil {
		return err
	}

	s.PC = 0

	return nil
}
