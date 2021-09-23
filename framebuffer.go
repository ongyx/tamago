package tamago

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Framebuffer provides a efficient way to manipulate pixels.
type Framebuffer struct {
	width, height int
	pixels        []byte
}

// Create a new framebuffer.
// Width and height must be positive.
func NewFramebuffer(width, height int) *Framebuffer {
	if !((width > 0) || (height > 0)) {
		return nil
	}

	return &Framebuffer{
		width:  width,
		height: height,
		pixels: make([]byte, width*height*4),
	}
}

// Get the total size of the framebuffer.
func (fb *Framebuffer) Size() int {
	return fb.width * fb.height * 4
}

// Write the colour c at the position (x, y).
// 0 <= x < width and 0 <= y < height must be true or a panic will occur.
func (fb *Framebuffer) Write(x, y int, c *color.RGBA) {
	index := x * y * 4

	if index > fb.Size() || index < 0 {
		panic(fmt.Sprintf("index (%d, %d) out of bounds!", x, y))
	}

	fb.pixels[index] = c.R
	fb.pixels[index+1] = c.G
	fb.pixels[index+2] = c.B
	fb.pixels[index+3] = c.A
}

// Copy the contents of the framebuffer into a screen.
func (fb *Framebuffer) CopyInto(screen *ebiten.Image) {
	screen.ReplacePixels(fb.pixels)
}
