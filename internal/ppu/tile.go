package ppu

import (
	"image"

	. "github.com/ongyx/tamago/internal/util"
)

const (
	// The maximum offset for reading from/writing to a tile, since each tile takes up 16 bytes in memory.
	MaxTileOffset = 16
)

var (
	// The width and height of a tile.
	TileWidthHeight = image.Point{8, 8}
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

// Dumps the tiles as an 8x8 image with a palette.
//
// If the image is too small, a panic occurs.
func (t *Tile) Dump(palette Palette, img *image.NRGBA) {
	if img.Bounds().Size() != TileWidthHeight {
		panic("image must be 8x8 pixels in size")
	}

	colors := palette.Unpack()

	for y, r := range t {
		for x, c := range r {
			// PixOffset assumes x and y are absolute coordinates, so both coordinates must be adjusted to be relative to Rect.Min for sub-images.
			pt := img.Rect.Min.Add(image.Pt(x, y))
			idx := img.PixOffset(pt.X, pt.Y)

			color := DefaultColorPalette[colors[c]]
			for o := range 4 {
				img.Pix[idx+o] = color[o]
			}
		}
	}
}
