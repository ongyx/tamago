package ppu

import (
	. "github.com/ongyx/tamago/internal/util"
)

const (
	// The start address of VRAM.
	VRAMStart = 0x8000
	// The end address of VRAM.
	VRAMEnd = 0xA000

	tileSetStart = 0x8000
	tileSetEnd   = 0x9800
	tileMapStart = 0x9800
	tileMapEnd   = 0xA000

	tileByteSize     = 16
	tileCount        = 384
	tileMapSize      = tileMapEnd - tileMapStart
	tileMapTableSize = tileMapSize / 2

	tileRows    = ScreenHeight / 8
	tileColumns = ScreenWidth / 8
	layerWidth  = 32
)

// The result of the PPU ticking.
type PPUTickResult struct {
	// Should the screen be rendered?
	Render bool
	// Should an interrupt occur?
	Interrupt bool
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

	return PPUTickResult{Render: rdr, Interrupt: irq}
}

// Renders the screen to the buffer of at least [util.DisplayBuffer] length.
func (p *PPU) Render(buffer []uint8) {
	if p.Registers.LCDC.IsBackgroundLayerEnabled() {
		p.renderBackground(buffer)
	}
}

func (p *PPU) renderBackground(buffer []uint8) {
	offset := int(p.Registers.LCDC.BackgroundTileMap()) * tileMapTableSize
	colors := p.Registers.BackgroundPalette()

	for ty := range tileRows {
		for tx := range tileColumns {
			// Fetch the map, then the tile index in the map, then the tile.
			mn := ty*layerWidth + tx
			ti := p.maps[offset+mn]

			var ati int
			if p.Registers.LCDC.BackgroundWindowTileSet() == 0 {
				// The tile index must be interpreted as a signed 8-bit offset to reach the last 128 tiles.
				ati = 256 + int(int8(ti))
			} else {
				// The tile index is below 256.
				ati = int(ti)
			}

			tile := p.tiles[ati]

			for y := range 8 {
				row := tile[y]
				py := 8*ty + y

				for x := range 8 {
					px := 8*tx + x
					cell := row[x]
					color := DefaultColorPalette[colors[cell]]

					bi := 4 * (py*ScreenWidth + px)
					for i := range 4 {
						buffer[bi+i] = color[i]
					}
				}
			}
		}
	}
}

// Reads an address in VRAM. The address must be between [VRAMStart] and [VRAMEnd].
func (p *PPU) ReadVRAM(addr uint16) uint8 {
	if addr >= tileSetStart && addr < tileSetEnd {
		rel := addr - tileSetStart
		idx := rel / tileByteSize
		off := rel % tileByteSize
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
		idx := rel / tileByteSize
		off := rel % tileByteSize
		p.tiles[idx].Write(off, v)
	} else if addr >= tileMapStart && addr < tileMapEnd {
		rel := addr - tileMapStart
		p.maps[rel] = v
	} else {
		panic("addr is out of bounds")
	}
}
