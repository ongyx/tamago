package tamago

type entry struct {
	offset uint16
	region Region
}

type MMU struct {
	entries []entry

	cart           *Cart
	rr             Renderer
	ram, oam, hram *RAM // ram = external + work ram
	io             *IO
}

func NewMMU(rr Renderer) *MMU {
	m := &MMU{
		cart: NewCart(),
		rr:   rr,
		ram:  NewRAM(0x4000),
		oam:  NewRAM(0x100),
		hram: NewRAM(0x80),
	}

	m.io = NewIO(m.rr)
	m.entries = []entry{
		{0x7fff, m.cart},
		{0x9fff, m.rr},
		{0xfdff, m.ram},
		{0xfeff, m.oam},
		{0xff7f, m.io},
		{0xfffe, m.hram},
		{0xffff, m.io},
	}

	return m
}

// Return the region addr is in and it's relative offset in the region.
func (m *MMU) region(addr uint16) (Region, uint16) {
	offset := uint16(0)

	for _, e := range m.entries {
		if addr <= e.offset {
			return e.region, addr - offset
		}
		offset = e.offset
	}

	return nil, 0
}

func (m *MMU) Read(addr uint16) uint8 {
	if r, o := m.region(addr); r != nil {
		return r.Read(o)
	}

	logger.Printf("MMU: unimplemented read for addr 0x%x", addr)
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

func (m *MMU) Write(addr uint16, val uint8) {
	if r, o := m.region(addr); r != nil {
		r.Write(o, val)
	}

	logger.Printf("MMU: unimplemented write for addr 0x%x (val 0x%x)", addr, val)
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
