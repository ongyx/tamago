package tamago

// State represents the current state of the emulation at some point in time.
type State struct {
	*MMU

	AF, BC, DE, HL *Register
	SP, PC         uint16

	bus     [0xFFFF]uint8
	stopped bool

	fl    *Flags
	clock *Clock
}

func NewState() *State {
	s := &State{}

	s.MMU = NewMMU()

	// The DMG bootrom assigns these values to the registers.
	// https://gbdev.io/pandocs/Power_Up_Sequence.html#cpu-registers
	s.AF = &Register{0x01, 0xb0}
	s.BC = &Register{0x00, 0x13}
	s.DE = &Register{0x00, 0xd8}
	s.HL = &Register{0x01, 0x4d}
	s.SP = 0xfffe
	s.PC = 0x100

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
