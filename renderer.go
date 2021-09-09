package tamago

// Renderer handles the drawing of pixels on a display (or internal framebuffer).
type Renderer interface {
	Write(x, y int, c Colour)
}
