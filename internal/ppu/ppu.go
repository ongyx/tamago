package ppu

const (
	// The start address of VRAM.
	VRAMStart = 0x8000
	// The end address of VRAM.
	VRAMEnd = 0xA000

	tileSetStart = 0x8000
	tileSetEnd   = 0x9800
	tileMapStart = 0x9800
	tileMapEnd   = 0xA000

	tileSize    = 16
	tileCount   = 384
	tileMapSize = tileMapEnd - tileMapStart
)

// The result of the PPU ticking.
type PPUTickResult struct {
	// Should the screen be rendered?
	DoRender bool
	// Should an interrupt occur?
	DoInterrupt bool
}

// Holds tile data and processes reads and writes to VRAM.
type PPU struct {
	// The PPU state.
	State PPUState
	// The LCD registers.
	Registers Registers

	tiles [tileCount]Tile
	maps  [tileMapSize]uint8
}

// Creates a new PPU.
func New() PPU {
	return PPU{State: NewPPUState(), Registers: NewRegisters()}
}

// Ticks the PPU by a number of M-cycles.
func (p *PPU) Tick(cycles uint8) PPUTickResult {
	om := p.State.Mode()
	ol := p.State.Line()
	rdr := p.State.Tick(cycles)
	m := p.State.Mode()
	l := p.State.Line()

	var irq bool

	if l != ol {
		// Scanline has changed, check if the LYC interrupt is necessary.
		eq := l == p.Registers.LYC
		irq = p.Registers.STAT.IsLYEQLYCInterruptEnabled() && eq

		p.Registers.STAT.SetLYEQLYC(eq)
		p.Registers.LY = l
	}

	if m != om {
		// Mode has changed, check if an interrupt is necessary.
		switch m {
		case ModeHBlank:
			irq = p.Registers.STAT.IsHBlankInterruptEnabled()
		case ModeVBlank:
			irq = p.Registers.STAT.IsVBlankInterruptEnabled()
		case ModeOAMRead:
			irq = p.Registers.STAT.IsOAMInterruptEnabled()
		}
	}

	p.Registers.STAT.SetMode(m)

	return PPUTickResult{DoRender: rdr, DoInterrupt: irq}
}

// Reads an address in VRAM. The address must be between [VRAMStart] and [VRAMEnd].
func (p *PPU) ReadVRAM(addr uint16) uint8 {
	if addr >= tileSetStart && addr < tileSetEnd {
		rel := addr - tileSetStart
		idx := rel / tileSize
		off := rel % tileSize
		return p.tiles[idx].Read(off)
	} else if addr >= tileMapStart && addr < tileMapEnd {
		rel := addr - tileMapStart
		return p.maps[rel]
	} else {
		panic("addr is out of bounds")
	}
}

// Writes a value to the address in VRAM. The address must be between [VRAMStart] and [VRAMEnd].
func (p *PPU) WriteVRAM(addr uint16, v uint8) {
	if addr >= tileSetStart && addr < tileSetEnd {
		rel := addr - tileSetStart
		idx := rel / tileSize
		off := rel % tileSize
		p.tiles[idx].Write(off, v)
	} else if addr >= tileMapStart && addr < tileMapEnd {
		rel := addr - tileMapStart
		p.maps[rel] = v
	} else {
		panic("addr is out of bounds")
	}
}
