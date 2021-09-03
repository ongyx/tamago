package tamago

type Entry struct {
	offset uint16
	region *Region
}

type MMU struct {
	entries []Entry

	cart           *Cart
	render         *Render
	ram, oam, hram *RAM // ram = external + work ram
	io             *IO
}

func NewMMU(rr *Renderer) *MMU {
	m := &MMU{
		cart:   NewCart(),
		render: NewRender(rr),
		ram:    NewRam(0x4000),
		oam:    NewRam(0x100),
		hram:   NewRam(0x80),
	}

	m.io = NewIO(m.render)
	m.entries = []MMUEntry{
		{0x7fff, s.cart},
		{0x9fff, s.render},
		{0xfdff, s.ram},
		{0xfeff, s.oam},
		{0xff7f, s.io},
		{0xfffe, s.hram},
		{0xffff, s.io},
	}

	return m
}

// Return the region addr is in and it's relative offset in the region.
func (m *MMU) region(addr uint16) (region *Region, offset uint16) {
	offset := uint16(0)

	for _, e := range m {
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
