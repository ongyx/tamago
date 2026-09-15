package util

const (
	// The width of the emulated screen.
	ScreenWidth = 160
	// The height of the emulated screen.
	ScreenHeight = 144
	// The buffer size required for rendering.
	DisplayBuffer = ScreenWidth * ScreenHeight * 4
)

// The default color palette, white, light gray, dark gray, then black.
var DefaultColorPalette = [4][4]uint8{
	{255, 255, 255, 255},
	{128, 128, 128, 255},
	{64, 64, 64, 255},
	{0, 0, 0, 255},
}
