package ppu

// A set of 4 color indexes packed into a byte.
type Palette uint8

// Unpacks the palette into its indexes.
func (p Palette) Unpack() (indexes [4]uint8) {
	pv := uint8(p)

	for i := range 8 {
		ci := i / 2
		pi := i % 2
		c := (pv >> i) & 0x1

		indexes[ci] |= c << pi
	}

	return indexes
}

// Unpacks the indexes into the palette.
func (p *Palette) Pack(indexes [4]uint8) {
	var pv uint8

	for i := range 8 {
		ci := i / 2
		pi := i % 2
		c := (indexes[ci] >> pi) & 0x1

		pv |= c << i
	}

	*p = Palette(pv)
}
