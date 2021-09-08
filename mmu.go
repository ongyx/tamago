package tamago

import (
	"errors"
	"io"
	"math/rand"
	"os"
)

var TooLargeErr = errors.New("ROM is too large!")

type MMU struct {
	bootrom [0x100]uint8
	rom     [0x8000]uint8
	vram    [0x2000]uint8
	ram     [0x4000]uint8
	// oam should only be 160 bytes (9f),
	// so 0 is returned for any read between 9f and ff.
	oam  [0x100]uint8
	hram [0x80]uint8

	ir     *Interrupt
	input  *Input
	render *Render

	hasBoot, hasROM bool
}

func NewMMU(rr Renderer) *MMU {
	m := &MMU{
		ir:    NewInterrupt(),
		input: NewInput(),
	}
	m.render = NewRender(m.vram[:], m.oam[:], rr)

	return m
}

func (m *MMU) Read(addr uint16) uint8 {
	switch {

	// rom
	case addr <= 0x7fff:
		if m.hasBoot && addr < 0x100 {
			// reading from bootrom
			return m.bootrom[addr]
		}

		// TODO: Bank switching?
		return m.rom[addr]

	// video ram
	case addr <= 0x9fff:
		return m.vram[addr-0x8000]

	// external + work ram
	case addr <= 0xdfff:
		return m.ram[addr-0xa000]

	// echo ram
	case addr <= 0xfdff:
		return m.ram[addr-0xe000]

	// oam (sprite ram)
	case addr <= 0xfe9f:
		return m.oam[addr-0xfe00]

	// unused
	case addr <= 0xfeff:

	/*
		I/O start
	*/

	case addr == 0xff00:
		return m.input.Poll()

	case addr == 0xff04:
		// TODO: Implement actual timer?
		return uint8(rand.Intn(256))

	case addr == 0xff0f:
		return uint8(m.ir.flags)

	case addr == 0xff40:
		return m.render.control

	case addr == 0xff42:
		return m.render.scrollY

	case addr == 0xff43:
		return m.render.scrollX

	case addr == 0xff44:
		return m.render.scanline

	case addr <= 0xff7f:
		// unimplemented i/o
		return 0x0

	/*
		I/O end
	*/

	// high ram
	case addr <= 0xfffe:
		return m.hram[addr-0xff80]

	// interrupt enable
	case addr == 0xffff:
		return m.ir.enable

	}

	logger.Printf("unimplemented read from addr 0x%x", addr)
	return 0
}

func (m *MMU) Write(addr uint16, val uint8) {
	impl := true

	switch {

	// rom
	case addr <= 0x7fff:
		impl = false

	// video ram
	case addr <= 0x9fff:
		m.vram[addr-0x8000] = val

	// external + work ram
	case addr <= 0xdfff:
		m.ram[addr-0xa000] = val

	// echo ram
	case addr <= 0xfdff:
		m.ram[addr-0xe000] = val

	// oam (sprite ram)
	case addr <= 0xfe9f:
		m.oam[addr-0xfe00] = val

	// unused
	case addr <= 0xfeff:
		impl = false

	/*
		I/O start
	*/

	case addr == 0xff00:
		m.input.Select(val)

	case addr == 0xff0f:
		m.ir.flags = iflag(val)

	case addr == 0xff40:
		m.render.control = val

	case addr == 0xff42:
		m.render.scrollY = val

	case addr == 0xff43:
		m.render.scrollX = val

	case addr == 0xff44:
		m.render.scanline = val

	case addr <= 0xff80:
		// unimplemented i/o

	/*
		I/O end
	*/

	// high ram
	case addr <= 0xfffe:
		m.hram[addr-0xff80] = val

	// interrupt enable
	case addr == 0xffff:
		m.ir.enable = val

	}

	if !impl {
		logger.Printf("unimplemented write to addr 0x%x (val 0x%x)", addr, val)
	}
}

// Read the byte at addr, where addr is the register's value.
func (m *MMU) ReadFrom(r *Register) uint8 {
	return m.Read(r.Get())
}

// Read the byte at addr and addr + 1 as a unsigned short.
func (m *MMU) ReadShort(addr uint16) uint16 {
	return Endian.Uint16([]uint8{m.Read(addr), m.Read(addr + 1)})
}

// Write a byte to addr, where addr is the register's value.
func (m *MMU) WriteTo(r *Register, val uint8) {
	m.Write(r.Get(), val)
}

// Write an unsigned short to addr and addr + 1.
func (m *MMU) WriteShort(addr uint16, val uint16) {
	buf := make([]uint8, 2)
	Endian.PutUint16(buf, val)

	m.Write(addr, buf[0])
	m.Write(addr+1, buf[1])
}

// Load a cartriage from a file.
func (m *MMU) Load(rom io.Reader) error {
	buf, err := io.ReadAll(rom)
	if err != nil {
		return err
	}

	// TODO: bank switching?
	if len(buf) > len(m.rom) {
		return TooLargeErr
	}

	copy(m.rom[:], buf)
	m.hasROM = true

	return nil
}

// Load a bootrom from a file.
func (m *MMU) LoadBoot(rom io.Reader) error {
	buf, err := io.ReadAll(rom)
	if err != nil {
		return err
	}

	if len(buf) > 256 {
		return TooLargeErr
	}

	copy(m.bootrom[:], buf)
	m.hasBoot = true

	return nil
}

// Dump the contents of the ROM and RAM into seperate files for debugging.
func (m *MMU) DebugDump() error {
	contents := []struct {
		name    string
		content []uint8
	}{
		{"bootrom", m.bootrom[:]},
		{"rom", m.rom[:]},
		{"vram", m.vram[:]},
		{"ram", m.ram[:]},
		{"oam", m.oam[:]},
		{"hram", m.hram[:]},
	}

	for _, c := range contents {
		if err := os.WriteFile(c.name, c.content, 0644); err != nil {
			return err
		}
	}

	return nil
}
