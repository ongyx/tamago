package ppu

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Provides debug functions for a PPU.
type Debugger struct {
	ppu *PPU
}

// Creates a new debugger for the PPU.
func NewDebugger(ppu *PPU) *Debugger {
	return &Debugger{ppu: ppu}
}

// Dumps the tiles to an image with the background palette.
func (d *Debugger) DumpTiles() *ebiten.Image {
	// 16 tiles by 24 tiles in size.
	img := ebiten.NewImage(16*TileWidthHeight, (tileCount/16)*TileWidthHeight)

	var buf []uint8

	for idx, tile := range d.ppu.tiles {
		dx := TileWidthHeight * (idx % 16)
		dy := TileWidthHeight * (idx / 16)

		buf = tile.Dump(d.ppu.Registers.BGP, buf)

		sub := img.SubImage(image.Rect(dx, dy, dx+8, dy+8)).(*ebiten.Image)
		sub.WritePixels(buf)
	}

	return img
}
