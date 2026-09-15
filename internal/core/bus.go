package core

import (
	"github.com/ongyx/tamago/internal/ppu"
	. "github.com/ongyx/tamago/internal/util"
)

const (
	// The start address of ROM.
	ROMStart uint16 = 0x0000
	// The end address of ROM.
	ROMEnd uint16 = 0x8000
)

// Handles memory address reads/writes.
type Bus struct {
	registers *Registers
	ppu       ppu.PPU
	rom       *Cart
	ram       [0x6000]uint8
}

// Creates a new bus.
func NewBus(rs *Registers) Bus {
	return Bus{
		registers: rs,
		ppu:       ppu.New(),
	}
}

// Loads a cartridge into the bus as read-only memory.
func (b *Bus) LoadCart(c *Cart) {
	b.rom = c
}

// Ejects the cartridge from the bus, if any.
func (b *Bus) EjectCart() {
	b.rom = nil
}

// Reads a byte from memory.
func (b *Bus) Read(addr uint16) uint8 {
	switch {
	case addr >= ROMStart && addr < ROMEnd:
		if b.rom != nil {
			return b.rom.Data[addr]
		} else {
			return 0x0
		}
	case addr >= ppu.VRAMStart && addr < ppu.VRAMEnd:
		return b.ppu.ReadVRAM(addr)
	case addr >= ppu.RegisterStart && addr < ppu.RegisterEnd:
		return b.ppu.Registers.Read(addr)
	case addr == 0xFF0F:
		return b.registers.IR.Encode()
	case addr == 0xFFFF:
		return b.registers.IE.Encode()
	default:
		return b.ram[addr]
	}
}

// Reads a word from memory.
func (b *Bus) ReadWord(addr uint16) uint16 {
	lo := b.Read(addr)
	hi := b.Read(addr + 1)
	return CombineWord(hi, lo)
}

// Writes a byte to memory.
func (b *Bus) Write(addr uint16, value uint8) {
	switch {
	case addr >= ROMStart && addr < ROMEnd:
		// ROM is not writable.
		return
	case addr >= ppu.VRAMStart && addr < ppu.VRAMEnd:
		b.ppu.WriteVRAM(addr, value)
	case addr >= ppu.RegisterStart && addr < ppu.RegisterEnd:
		b.ppu.Registers.Write(addr, value)
	case addr == 0xFF0F:
		b.registers.IR.Decode(value)
	case addr == 0xFFFF:
		b.registers.IE.Decode(value)
	default:
		b.ram[addr] = value
	}
}

// Writes a word to memory.
func (b *Bus) WriteWord(addr uint16, value uint16) {
	hi, lo := SplitWord(value)
	b.Write(addr, lo)
	b.Write(addr+1, hi)
}
