package tamago

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	renderWidth  = 160
	renderHeight = 144
)

type (
	Tile    [8][8]uint8
	Tileset [384]Tile
	Palette [4]color.Color
)

var (
	mode = struct {
		HBlank, VBlank, OAM, VRAM uint8
	}{
		0, 1, 2, 3,
	}

	White     = color.RGBA{255, 255, 255, 0}
	LightGray = color.RGBA{192, 192, 192, 0}
	DarkGray  = color.RGBA{96, 96, 96, 0}
	Black     = color.RGBA{0, 0, 0, 0}

	DefaultPalette = Palette{White, LightGray, DarkGray, Black}
)

// Render draws tiles from VRAM onto a screen.
// sx and sy are the offsets of the display (size 160x144)
// from (0,0) at the top left of the background map (size 256x256).
type Render struct {
	sx, sy, mode, line uint8
	tick               int
	lcdc               *Bits

	tileset Tileset
	palette Palette
	screen  *ebiten.Image

	vram []uint8
	oam  []uint8
}

func NewRender(vram, oam []uint8) *Render {
	screen := ebiten.NewImage(renderWidth, renderHeight)

	return &Render{
		lcdc:    &Bits{0},
		palette: DefaultPalette,
		screen:  screen,
		vram:    vram,
		oam:     oam,
	}
}

// Given an address in VRAM (8000-97ff) and the value being written to it, update the tileset.
func (r *Render) update(addr uint16) {
	offset := addr - 0x8000
	index := offset / 16

	lo := Bits{r.vram[offset]}
	hi := Bits{r.vram[offset+1]}

	y := (offset % 16) / 2
	for x := 0; x < 8; x++ {
		var pixel uint8

		// leftmost pixel starts at 7th bit
		if lo.At(7 - x) {
			pixel += 1
		}

		if hi.At(7 - x) {
			pixel += 2
		}

		r.tileset[index][y][x] = pixel
	}
}

// Return the VRAM offset of the current tile.
func (r *Render) offset() uint16 {
	var base uint16

	// Get the background map's base address.
	if r.lcdc.At(3) {
		base = 0x1c00
	} else {
		base = 0x1800
	}

	// Since each tile is 8x8 pixels, the address of the tile reference in VRAM is:
	// map offset + ((y offset + display offset) / 8) + (x offset / 8)
	// because all offsets are in pixels.
	return base + uint16((r.sy+r.line)/8) + uint16(r.sx/8)
}

// Return the tile pointed to by the VRAM offset.
func (r *Render) tile(offset uint16) Tile {
	idx := int(r.vram[offset])

	if !r.lcdc.At(4) && idx < 128 {
		idx += 256
	}

	return r.tileset[idx]
}

// Render the next scanline.
func (r *Render) scanline() {
	offset := r.offset()
	tile := r.tile(offset)

	// Calculate the x and y offset of the current pixel in the tile.
	y := (r.sy + r.line) % 8
	x := r.sx % 8

	dy := int(r.line)
	for dx := 0; dx < 160; x++ {
		logger.Printf("drawing pixel at (%d,%d) with tile offset (%d,%d)", dx, dy, x, y)
		// Draw the pixel at the offsets in the tile.
		colour := r.palette[tile[y][x]]
		r.screen.Set(dx, dy, colour)

		x++
		if x == 8 {
			// Reached the end of this tile, get the tile referred to by the next offset.
			x = 0
			offset++
			tile = r.tile(offset)
		}
	}
}

// Update the GPU timing, given the main state's clock.
func (r *Render) step(c *Clock) {
	r.tick += c.t

	switch r.mode {

	case mode.HBlank:
		if r.tick >= 204 {
			r.tick = 0
			r.line++

			if r.line == 143 {
				// last scanline, vblank
				r.mode = mode.VBlank
			} else {
				r.mode = mode.OAM
			}
		}

	case mode.VBlank:
		if r.tick >= 456 {
			r.tick = 0
			r.line++

			if r.line >= 153 {
				r.mode = mode.OAM
				r.line = 0
			}
		}

	case mode.OAM:
		if r.tick >= 80 {
			r.tick = 0
			r.mode = mode.VRAM
		}

	case mode.VRAM:
		if r.tick >= 172 {
			r.tick = 0
			r.mode = 0

			r.scanline()
		}

	}
}
