package tamago

var (
	White     = Colour{255, 255, 255, 0}
	LightGray = Colour{192, 192, 192, 0}
	DarkGray  = Colour{96, 96, 96, 0}
	Black     = Colour{0, 0, 0, 0}

	DefaultPalette = Palette{White, LightGray, DarkGray, Black}
)

type Colour struct {
	r, g, b, a uint8
}

type Palette [4]Colour
