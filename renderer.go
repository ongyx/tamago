package tamago

// A renderer handles the drawing of pixels on a display (or internal framebuffer).
type Renderer interface {
	Region
	Status() *Status
}

// Status represents the current state of a renderer.
type Status struct {
	control, scrollX, scrollY, scanline, tick uint8
}
