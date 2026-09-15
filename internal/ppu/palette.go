package ppu

// A 4-color palette packed into a byte.
type Palette uint8

// Unpacks the palette into its individual colors.
func (p Palette) Unpack() (colors [4]uint8) {
	pv := uint8(p)

	for i := range 8 {
		ci := i / 2
		pi := i % 2
		c := (pv >> i) & 0x1

		colors[ci] |= c << pi
	}

	return colors
}

// Unpacks the individual colors into the palette.
func (p *Palette) Pack(colors [4]uint8) {
	var pv uint8

	for i := range 8 {
		ci := i / 2
		pi := i % 2
		c := (colors[ci] >> pi) & 0x1

		pv |= c << i
	}

	*p = Palette(pv)
}
