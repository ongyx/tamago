package tamago

import (
	"image/color"
)

const (
	renderWidth  = 160
	renderHeight = 144
)

var mode = struct {
	HBlank, VBlank, OAM, VRAM uint8
}{
	0, 1, 2, 3,
}

// Render draws tiles from VRAM onto a screen.
// sx and sy are the offsets of the display (size 160x144)
// from (0,0) at the top left of the background map (size 256x256).
type Render struct {
	sx, sy, mode, line uint8
	tick               int
	lcdc               *Bits

	tileset    Tileset
	spriteData SpriteData

	bg, obj0, obj1 Palette

	fb *Framebuffer

	vram []uint8
	oam  []uint8
}

func NewRender(vram, oam []uint8) *Render {
	return &Render{
		lcdc: &Bits{0},
		bg:   DefaultPalette,
		obj0: DefaultPalette,
		obj1: DefaultPalette,
		fb:   NewFramebuffer(renderWidth, renderHeight),
		vram: vram,
		oam:  oam,
	}
}

// Given an offset in VRAM, update the tileset.
func (r *Render) updateTile(offset uint16) {
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

// Given an offset in OAM, update the sprite data.
func (r *Render) updateSprite(offset uint16) {
	val := r.oam[offset]
	sprite := &r.spriteData[offset/4]

	switch offset % 4 {
	case 0:
		sprite.y = val
	case 1:
		sprite.x = val
	case 2:
		sprite.tile = val
	case 3:
		sprite.options.uint8 = val
	}
}

// Update a palette (bg, obj0 or obj1) with a value
func (r *Render) updatePalette(p *Palette, v uint8) {
	for i := 0; i < 4; i++ {
		colour := &p[i]

		// select each set of 2 bits from 0 to 7.
		switch (v >> (i * 2)) & 0x3 {
		case 0:
			*colour = White
		case 1:
			*colour = LightGrey
		case 2:
			*colour = DarkGrey
		case 3:
			*colour = Black
		}
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

// Render the next scanline on screen.
func (r *Render) scanline() {
	var scanline [160]*color.RGBA

	dy := int(r.line)

	// background enable
	if r.lcdc.At(0) {
		offset := r.offset()
		tile := r.tile(offset)

		// Calculate the x and y offset of the current pixel in the tile.
		y := (r.sy + r.line) % 8
		x := r.sx % 8

		for dx := range scanline {
			scanline[dx] = &r.bg[tile[y][x]]

			x++
			if x == 8 {
				// Reached the end of this tile, get the tile referred to by the next offset.
				x = 0
				offset++
				tile = r.tile(offset)
			}
		}
	}

	// object enable
	if r.lcdc.At(1) {
		for i := 0; i < 40; i++ {
			sprite := r.spriteData[i]

			var sx, sy int
			sx = int(sprite.x) - 8
			sy = int(sprite.y) - 16

			if sy <= dy && (sy+8) > dy {
				var p Palette

				// palette select
				if sprite.options.At(4) {
					p = r.obj1
				} else {
					p = r.obj0
				}

				var rowIndex int

				// y-flip
				if sprite.options.At(6) {
					rowIndex = 7 - (dy - sy)
				} else {
					rowIndex = dy - sy
				}
				row := r.tileset[sprite.tile][rowIndex]

				for x := 0; x < 8; x++ {
					pos := sx + x

					if pos >= 0 && pos < 160 && (!sprite.options.At(7) || *scanline[pos] == White) {

						index := x
						if sprite.options.At(5) {
							index = 7 - x
						}

						colour := p[row[index]]
						if colour != White {
							scanline[sx] = &colour
						}

						sx++

					}
				}
			}
		}
	}

	for dx, colour := range scanline {
		if colour == nil {
			colour = &White
		}
		logger.Printf("drawing %v at (%d,%d)", colour, dx, dy)
		r.fb.Write(dx, dy, colour)
	}

}

// Given the emulation state clock, update the render.
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
