package tamago

type tile [8][8]uint8

var mode = struct {
	HBlank, VBlank, OAM, VRAM uint8
}{
	0, 1, 2, 3,
}

// Renderer handles the drawing of pixels on a display (or internal framebuffer).
type Renderer interface {
	Write(x, y int, c Colour)
}

// Render wraps a renderer to draw sprites from VRAM.
type Render struct {
	scrollX, scrollY, mode, line uint8
	tick                         int
	lcdc                         *Bits
	tileset                      [384]tile
	palette                      Palette

	vram []uint8
	oam  []uint8

	rr Renderer
}

func NewRender(vram, oam []uint8, rr Renderer) *Render {
	return &Render{
		lcdc:    &Bits{0},
		palette: DefaultPalette,
		vram:    vram,
		oam:     oam,
		rr:      rr,
	}
}

// Given an address in VRAM (8000-9fff) and the value being written to it, update the tileset.
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

func (r *Render) tile(addr uint16) int {
	tile := int(r.vram[addr])

	if !r.lcdc.At(4) && tile < 128 {
		// tile map starts at $8800, use signed addressing
		tile += 256
	}

	return tile
}

// Render the next scanline.
func (r *Render) scanline() {
	var offset uint16

	// background tile map offset
	if r.lcdc.At(3) {
		offset = 0x1c00
	} else {
		offset = 0x1800
	}

	offset += uint16(r.line+r.scrollY) / 8

	line := r.scrollX / 8
	y := (r.line + r.scrollY) & 0x7
	x := r.scrollX & 0x7

	tile := r.tile(offset + uint16(line))

	for i := 0; i < 160; i++ {
		colour := r.palette[r.tileset[tile][y][x]]
		r.rr.Write(i, int(r.line), colour)

		x++

		if x == 8 {
			x = 0
			line = (line + 1) & 31
			tile = r.tile(offset + uint16(line))
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
