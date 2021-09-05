package tamago

// A renderer handles the drawing of pixels on a display (or internal framebuffer).
type Renderer interface {
	Read(x, y int) Colour
	Write(x, y int, c Colour)
}

type Render struct {
	control, scrollX, scrollY, scanline uint8
	tick                                uint

	// Slice reference to the actual array.
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
