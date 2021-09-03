package tamago

// A renderer handles the drawing of pixels on a display (or internal framebuffer).
type Renderer interface {
	Region
}

// A render is an abstract wrapper around a renderer that emulates a Gameboy's display.
type Render struct {
	Renderer

	control, scrollX, scrollY, scanline, tick uint8
}

func NewRender(rr *Renderer) *Render {
	return &Render{Renderer: rr}
}
