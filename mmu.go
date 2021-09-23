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
	oam     [0xa0]uint8
	hram    [0x80]uint8

	ir     *Interrupt
	input  *Input
	render *Render

	hasBoot, hasROM bool
}

func NewMMU() *MMU {
	m := &MMU{
		ir:    NewInterrupt(),
		input: NewInput(),
	}
	m.render = NewRender(m.vram[:], m.oam[:])

	return m
}

func (m *MMU) Read(addr uint16) uint8 {
	logger.Printf("reading from addr 0x%x", addr)

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
		return m.ir.requested

	case addr == 0xff40:
		return m.render.lcdc.uint8

	case addr == 0xff42:
		return m.render.sy

	case addr == 0xff43:
		return m.render.sx

	case addr == 0xff44:
		return m.render.line

	case addr <= 0xff7f:
		// unimplemented i/o

	/*
		I/O end
	*/

	// high ram
	case addr <= 0xfffe:
		return m.hram[addr-0xff80]

	// interrupt enable
	case addr == 0xffff:
		return m.ir.enabled

	}

	logger.Printf("unimplemented read from addr 0x%x", addr)
	return 0
}

func (m *MMU) Write(addr uint16, val uint8) {
	logger.Printf("writing to addr 0x%x", addr)

	impl := true

	switch {

	// rom
	case addr <= 0x7fff:
		impl = false

	// video ram
	case addr <= 0x9fff:
		m.vram[addr-0x8000] = val

		if addr <= 0x97ff {
			m.render.updateTile(addr - 0x8000)
		}

	// external + work ram
	case addr <= 0xdfff:
		m.ram[addr-0xa000] = val

	// echo ram
	case addr <= 0xfdff:
		m.ram[addr-0xe000] = val

	// oam (sprite ram)
	case addr <= 0xfe9f:
		m.oam[addr-0xfe00] = val
		m.render.updateSprite(addr - 0xfe00)

	// unused
	case addr <= 0xfeff:
		impl = false

	/*
		I/O start
	*/

	case addr == 0xff00:
		m.input.Select(val)

	case addr == 0xff0f:
		m.ir.requested = val

	case addr == 0xff40:
		m.render.lcdc.uint8 = val

	case addr == 0xff42:
		m.render.sy = val

	case addr == 0xff43:
		m.render.sx = val

	case addr == 0xff44:
		m.render.line = val

	// Background palette
	case addr == 0xff47:
		m.render.updatePalette(&m.render.bg, val)

	// Object palette 1
	case addr == 0xff48:
		m.render.updatePalette(&m.render.obj0, val)

	case addr == 0xff49:
		m.render.updatePalette(&m.render.obj1, val)

	case addr <= 0xff80:
		// unimplemented i/o
		impl = false

	/*
		I/O end
	*/

	// high ram
	case addr <= 0xfffe:
		m.hram[addr-0xff80] = val

	// interrupt enable
	case addr == 0xffff:
		m.ir.enabled = val

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

// Load a cartriage from a reader.
func (m *MMU) LoadFrom(rom io.Reader) error {
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

// Load a cartriage from a filename.
func (m *MMU) Load(rom string) error {
	f, err := os.Open(rom)
	defer f.Close()

	if err != nil {
		return err
	}

	if e := m.LoadFrom(f); e != nil {
		return e
	}

	return nil
}

// Load a bootrom from a reader.
func (m *MMU) LoadBootFrom(rom io.Reader) error {
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

// Load a bootrom from a filename.
func (m *MMU) LoadBoot(rom string) error {
	f, err := os.Open(rom)
	defer f.Close()

	if err != nil {
		return err
	}

	if e := m.LoadBootFrom(f); e != nil {
		return e
	}

	return nil
}

// Check if a rom/bootrom has been loaded.
func (m *MMU) Loaded() bool {
	return m.hasBoot || m.hasROM
}
