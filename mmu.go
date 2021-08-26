package tamago

// MMU handles reading from and writing to the emulated address bus.
// The contents of the ROM and RAM are stored here.
// However, this should not be used directly: use State.Read and State.Write instead.
type MMU struct {
	rom  [0x8000]uint8
	vram [0x2000]uint8
	eram [0x2000]uint8
	wram [0x2000]uint8
	oam  [0x100]uint8
	io   [0x100]uint8
	hram [0x80]uint8
}

func NewMMU() *MMU {
	return &MMU{}
}

func (m *MMU) pointer(addr uint16) *uint8 {
	var v *uint8
	*v = 0

	switch {

	case addr <= 0x7fff:
		v = &m.rom[addr]

	case addr <= 0x9fff:
		v = &m.vram[addr-0x8000]

	case addr <= 0xbfff:
		v = &m.eram[addr-0xa000]

	case addr <= 0xdfff:
		v = &m.wram[addr-0xc000]

	case addr <= 0xfdff:
		v = &m.wram[addr-0xe000]

	case addr <= 0xfeff:
		v = &m.oam[addr-0xfe00]

	case addr <= 0xff7f:
		// TODO
		v = &m.io[addr-0xff00]

	case addr <= 0xfffe:
		v = &m.hram[addr-0xff80]

	case addr == 0xffff:
		// TODO

	}

	return v
}

func (m *MMU) read(addr uint16) uint8 {
	ptr := m.pointer(addr)
	return *ptr
}

func (m *MMU) write(addr uint16, val uint8) {
	ptr := m.pointer(addr)
	*ptr = val
}
