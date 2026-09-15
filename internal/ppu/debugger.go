package ppu

import (
	"image"
)

const (
	debugTileCountPerRow = 16
)

// Provides debug functions for a PPU.
type Debugger struct {
	ppu *PPU
}

// Creates a new debugger for the PPU.
func NewDebugger(ppu *PPU) *Debugger {
	return &Debugger{ppu: ppu}
}

// Dumps the tiles to an NRGBA image with the background palette.
func (d *Debugger) DumpTiles() *image.NRGBA {
	// 16 tiles by 24 tiles in size.
	bounds := image.Rectangle{
		Max: image.Point{
			X: debugTileCountPerRow * TileWidthHeight.X,
			Y: (tileCount / debugTileCountPerRow) * TileWidthHeight.Y,
		},
	}
	img := image.NewNRGBA(bounds)

	for idx, tile := range d.ppu.tiles {
		px := TileWidthHeight.X * (idx % debugTileCountPerRow)
		py := TileWidthHeight.Y * (idx / debugTileCountPerRow)

		// Dump directly into the image.
		sub := img.SubImage(image.Rect(px, py, px+8, py+8)).(*image.NRGBA)
		tile.Dump(d.ppu.Registers.BGP, sub)
	}

	return img
}
