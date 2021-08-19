package tamago

// State represents the current state of the emulation at some point in time.
type State struct {
	AF, BC, DE, HL *Register
	SP, PC         uint16
	Flags          *Flags
	Clock          *Clock
	MMU            *MMU
	GPU            *GPU
	Stopped        bool
}

func (s *State) Read(addr uint16) uint8 {
	return s.MMU.read(addr)
}

func (s *State) ReadShort(addr uint16) uint16 {
	return Endian.Uint16([]uint8{s.Read(addr), s.Read(addr + 1)})
}

func (s *State) Write(addr uint16, val uint8) {
	s.MMU.write(addr)
}

func (s *State) WriteShort(addr uint16, val uint16) {
	buf := []uint8{}
	Endian.PutUint16(buf, val)

	s.Write(addr, buf[0])
	s.Write(addr, buf[1])
}
