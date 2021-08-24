package tamago

// State represents the current state of the emulation at some point in time.
type State struct {
	AF, BC, DE, HL *Register
	SP, PC         uint16

	fl      *Flags
	intr    *Interrupt
	clock   *Clock
	mmu     *MMU
	gpu     *GPU
	stopped bool
}

func NewState() *State {
	s := &State{}
	s.AF = NewRegister()
	s.BC = NewRegister()
	s.DE = NewRegister()
	s.HL = NewRegister()

	s.fl = NewFlags(s.AF)
	s.intr = NewInterrupt()
	s.clock = NewClock()
	s.mmu = NewMMU()
	s.gpu = NewGPU()
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
		s.Clock.Step(3)
	} else {
		s.Clock.Step(2)
	}
}

/*
	I/O functions
*/

// Read the byte at addr.
func (s *State) Read(addr uint16) uint8 {
	return s.mmu.read(addr)
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
	s.mmu.write(addr, val)
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
	s.Write(addr, buf[1])
}

/*
	Stack functions
*/

// Push a value onto the stack.
func (s *State) Push(v uint16) {
	s.SP -= 2
	s.WriteShort(s.SP.Get(), v)
}

// Pop a value from the stack.
func (s *State) Pop() uint16 {
	v := s.ReadShort(s.SP.Get())
	s.SP += 2

	return v
}
