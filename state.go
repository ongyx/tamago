package tamago

import (
	"math/rand"
)

// State represents the current state of the emulation at some point in time.
type State struct {
	AF, BC, DE, HL *Register
	SP, PC         uint16

	mem     [0xFFFF]uint8
	stopped bool

	fl    *Flags
	intr  *Interrupt
	clock *Clock
	gpu   *GPU
	input *Input
}

func NewState() *State {
	s := &State{}

	// The DMG bootrom assigns these values to the registers.
	// https://gbdev.io/pandocs/Power_Up_Sequence.html#cpu-registers
	s.AF = &Register{0x01, 0xb0}
	s.BC = &Register{0x00, 0x13}
	s.DE = &Register{0x00, 0xd8}
	s.HL = &Register{0x01, 0x4d}
	s.SP = 0xfffe
	s.PC = 0x100

	s.fl = NewFlags(s.AF)
	s.intr = NewInterrupt()
	s.clock = NewClock()
	s.gpu = NewGPU()
	s.input = NewInput()

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
	I/O functions
*/

// Read the byte at addr.
func (s *State) Read(addr uint16) uint8 {

	switch addr {

	// I/O
	case 0xff00:
		return s.input.Poll()

	case 0xff04:
		buf := make([]uint8, 1)
		rand.Read(buf)
		return buf[0]

	case 0xff40:
		return s.gpu.control

	case 0xff42:
		return s.gpu.scrollY

	case 0xff43:
		return s.gpu.scrollX

	case 0xff44:
		return s.gpu.scanline

	case 0xff0f:
		return uint8(s.intr.flags)

	case 0xffff:
		return tobit(s.intr.enable)

	}

	return s.mem[addr]
}

// Read the byte at addr, where addr is the register's value.
func (s *State) ReadFrom(r *Register) uint8 {
	return s.Read(r.Get())
}

// Read the byte at addr and addr + 1 as a unsigned short.
func (s *State) ReadShort(addr uint16) uint16 {
	return Endian.Uint16([]uint8{s.Read(addr), s.Read(addr + 1)})
}

// Write a byte to addr.
func (s *State) Write(addr uint16, val uint8) {
	s.mem[addr] = val
}

// Write a byte to addr, where addr is the register's value.
func (s *State) WriteTo(r *Register, val uint8) {
	s.Write(r.Get(), val)
}

// Write an unsigned short to addr and addr + 1.
func (s *State) WriteShort(addr uint16, val uint16) {
	buf := []uint8{}
	Endian.PutUint16(buf, val)

	s.Write(addr, buf[0])
	s.Write(addr+1, buf[1])
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
