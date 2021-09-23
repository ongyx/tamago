package tamago

import (
	"fmt"
	"io"

	"github.com/hajimehoshi/ebiten/v2"
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

func NewState() *State {
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

	s.MMU = NewMMU()

	s.fl = NewFlags(s.AF)
	s.clock = NewClock()

	return s
}

func (s *State) fetch() uint8 {
	b := s.Read(s.PC)
	s.PC++
	return b
}

func (s *State) step(screen *ebiten.Image) {
	var ins Instruction

	opcode := s.fetch()
	if opcode == 0xcb {
		opcode = s.fetch()
		ins = cbops[opcode]
	} else {
		ins = ops[opcode]
	}

	buf := make([]uint8, 2)
	for i := 0; i < ins.length; i++ {
		buf[i] = s.fetch()
	}
	value := NewValue(buf)

	if s.PC == 0x100 {
		s.hasBoot = false
	}

	logger.Printf("[0x%x] executing %s", s.PC, ins.Asm(value))

	ins.fn(s, value)
	s.clock.Step(ins.cycles)
	s.render.step(s.clock)

	// handle interrupts
	ir := s.ir.todo()
	if ir != 0 {
		s.Push(s.PC)

		var cycles int
		switch {

		case (ir & VBlank) != 0:
			s.render.fb.CopyInto(screen)
			s.PC = 0x40
			cycles += 3

		case (ir & LCDStat) != 0:
			s.PC = 0x48
			cycles += 3

		case (ir & Timer) != 0:
			s.PC = 0x50
			cycles += 3

		case (ir & Serial) != 0:
			s.PC = 0x58
			cycles += 3

		case (ir & Joypad) != 0:
			s.PC = 0x60
			cycles += 3

		}

		s.clock.Step(cycles)
	}
}

/*
	Stack/jump functions
*/

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
