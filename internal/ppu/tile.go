package ppu

import (
	"fmt"

	. "github.com/ongyx/tamago/internal/util"
)

const (
	// The maximum offset for reading from/writing to a tile, since each tile takes up 16 bytes in memory.
	MaxTileOffset = 16

	// The width and height of a tile.
	TileWidthHeight = 8
	// The buffer size required for a tile in RGBA format.
	TileBufferSize = TileWidthHeight * TileWidthHeight * 4
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

	r := uint8(off / 2)
	// Even offsets address the LSB, odd offsets address the MSB.
	n := uint8(off % 2)

	var v uint8
	for i := range 8 {
		// Copy the bit from the tile pixel into v, then left shift to make space for the next pixel.
		v |= GetBit(t[r][i], n)
		v <<= 1
	}

	return v
}

// Writes a byte value at offset. A panic occurs if the offset exceeds [MaxTileOffset].
func (t *Tile) Write(off uint16, value uint8) {
	if off >= MaxTileOffset {
		panic("off exceeds MaxTileOffset")
	}

	r := uint8(off / 2)
	n := uint8(off % 2)

	for i := range 8 {
		// Get the corresponding pixel, then set the bit in it.
		v := &t[r][7-i]
		b := HasBit(value, uint8(i))

		*v = SetBit(*v, n, b)
	}
}

// Dumps the tiles as an 8x8 image into the buffer in RGBA format, given the color palette.
//
// If the buffer is nil, a suitably sized one is created.
// If the buffer is too small, a panic occurs.
func (t *Tile) Dump(palette Palette, buffer []uint8) []uint8 {
	if buffer == nil {
		buffer = make([]uint8, TileBufferSize)
	} else if len(buffer) < TileBufferSize {
		panic(fmt.Sprintf("buffer length must be at least %d", TileBufferSize))
	}

	colors := palette.Unpack()

	for y, r := range t {
		for x, c := range r {
			color := DefaultColorPalette[colors[c]]

			idx := 4 * (TileWidthHeight*y + x)
			for o := range 4 {
				buffer[idx+o] = color[o]
			}
		}
	}

	return buffer
}
