package tamago

import (
	"image/color"
)

var (
	White     = color.RGBA{255, 255, 255, 0}
	LightGrey = color.RGBA{192, 192, 192, 0}
	DarkGrey  = color.RGBA{96, 96, 96, 0}
	Black     = color.RGBA{0, 0, 0, 0}

	DefaultPalette = Palette{White, LightGrey, DarkGrey, Black}
)

type Palette [4]color.RGBA
