package tamago

import "math/rand"

type mmapEntry struct {
	offset uint16
	region *Region
}

type mmap []mmapEntry

// MMU handles reading from and writing to the emulated address bus.
// The contents of the ROM and RAM are stored here.
type MMU struct {
	cart           *Cart
	ppu            *PPU
	ram, oam, hram *RAM // ram = external + work ram
	io             *IO

	mmap mmap
}

func NewMMU() *MMU {
	m := &MMU{
		cart: NewCart(),
		ppu:  NewPPU(),
		ram:  NewRam(0x4000),
		oam:  NewRam(0x100),
		hram: NewRam(0x80),
		io:   NewIO(),
	}
}

// Read the byte at addr.
func (m *MMU) Read(addr uint16) uint8 {

	switch {

	case addr <= 0x7fff:
		return m.rom[addr]

	case addr <= 0x9fff:
		return m.vram[addr-0x8000]

	case addr <= 0xbfff:
		return m.eram[addr-0xa000]

	case addr <= 0xdfff:
		return m.wram[addr-0xc000]

	case addr <= 0xfdff:
		return m.wram[addr-0xe000]

	case addr <= 0xfeff:
		return m.oam[addr-0xfe00]

	// I/O
	case addr == 0xff00:
		return m.input.Poll()

	case addr == 0xff04:
		buf := make([]uint8, 1)
		rand.Read(buf)
		return buf[0]

	case addr == 0xff0f:
		return uint8(m.intr.flags)

	case addr == 0xff40:
		return m.ppu.control

	case addr == 0xff42:
		return m.ppu.scrollY

	case addr == 0xff43:
		return m.ppu.scrollX

	case addr == 0xff44:
		return m.ppu.scanline

	case addr <= 0xff7f:
		// TODO
		return m.io[addr-0xff00]

	case addr == 0xffff:
		return tobit(m.intr.enable)

	case addr <= 0xfffe:
		return m.hram[addr-0xff80]

	}

	return 0
}

// Read the byte at addr, where addr is the register's value.
func (m *MMU) ReadFrom(r *Register) uint8 {
	return m.Read(r.Get())
}

// Read the byte at addr and addr + 1 as a unsigned short.
func (m *MMU) ReadShort(addr uint16) uint16 {
	return Endian.Uint16([]uint8{m.Read(addr), m.Read(addr + 1)})
}

// Write a byte to addr.
func (m *MMU) Write(addr uint16, val uint8) {

	invalid := false

	switch {

	case addr <= 0x7fff:
		invalid = true

	case addr <= 0x9fff:
		m.vram[addr-0x8000] = val

	case addr <= 0xbfff:
		m.eram[addr-0xa000] = val

	case addr <= 0xdfff:
		m.wram[addr-0xc000] = val

	case addr <= 0xfdff:
		m.wram[addr-0xe000] = val

	case addr <= 0xfeff:
		m.oam[addr-0xfe00] = val

	// I/O
	case addr == 0xff40:
		m.ppu.control = val

	case addr == 0xff42:
		m.ppu.scrollY = val

	case addr == 0xff43:
		m.ppu.scrollX = val

	case addr == 0xff46:
		// Copy a region of the cart or RAM to the OAM region.
		dst := uint16(0xfe00)
		src := uint16(val) << 8

		for i := uint16(0); i < 160; i++ {
			m.Write(dst+i, m.Read(src+i))
		}

	case addr == 0xff0f:
		m.intr.flags = val

	case addr == 0xffff:
		if val > 0 {
			m.intr.enable = true
		} else {
			m.intr.enable = false
		}
	}

}

// Write a byte to addr, where addr is the register's value.
func (m *MMU) WriteTo(r *Register, val uint8) {
	m.Write(r.Get(), val)
}

// Write an unsigned short to addr and addr + 1.
func (m *MMU) WriteShort(addr uint16, val uint16) {
	buf := []uint8{}
	Endian.PutUint16(buf, val)

	m.Write(addr, buf[0])
	m.Write(addr+1, buf[1])
}
