package ppu

const (
	// The maximum offset for reading from/writing to a tile, since each tile takes up 16 bytes in memory.
	MaxTileOffset = 16
)

// A tile 8x8 pixels in size.
//
// Pixels in a tile are encoded in memory as 2 bits split across 2 bytes (lo, hi),
// where lo.7 contains the LSB, hi.7 contains the MSB, and so on for each tile in the row. (lo, hi) represents the whole row.
//
// Note that pixels are read from left to right: bit 7 contains data for pixel 1, bit 6 contains data for pixel 2, etc.
type Tile [8][8]uint8

// Reads a byte value at offset. A panic occurs if the offset exceeds [MaxTileOffset].
func (t *Tile) Read(off uint16) uint8 {
	if off >= MaxTileOffset {
		panic("off exceeds MaxTileOffset")
	}

	r := off / 2
	// Even offsets address the LSB, odd offsets address the MSB.
	b := off % 2

	var v uint8
	for i := range 8 {
		// Copy the LSB or MSB from the row into the value.
		v |= t[r][7-i] & (1 << b)
	}

	return v
}

// Writes a byte value at offset. A panic occurs if the offset exceeds [MaxTileOffset].
func (t *Tile) Write(off uint16, v uint8) {
	if off >= MaxTileOffset {
		panic("off exceeds MaxTileOffset")
	}

	r := off / 2
	b := off % 2

	for i := range 8 {
		t[r][7-i] |= v & (1 << b)
	}
}
