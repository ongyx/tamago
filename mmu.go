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

func (m *MMU) read(addr uint16) uint8 {
	return m.bus[addr]
}

func (m *MMU) write(addr uint16, val uint8) {
	m.bus[addr] = val
}
