package tamago

var mode = struct {
	HBlank, VBlank, OAM, VRAM uint8
}{
	0, 1, 2, 3,
}

// Render wraps a renderer to draw sprites from VRAM.
type Render struct {
	control, scrollX, scrollY, mode, line uint8
	tick                                  int

	vram []uint8
	oam  []uint8

	rr Renderer
}

func NewRender(vram, oam []uint8, rr Renderer) *Render {
	return &Render{
		vram: vram,
		oam:  oam,
		rr:   rr,
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

			// TODO: renderscan
		}

	}
}
