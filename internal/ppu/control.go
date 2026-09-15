package ppu

import (
	. "github.com/ongyx/tamago/internal/util"
)

const (
	lcdcIsBackgroundWindowLayerEnabled = iota
	lcdcIsSpriteLayerEnabled
	lcdcAreSprites8x16
	lcdcBackgroundTileMap
	lcdcBackgroundWindowTileSet
	lcdcIsWindowLayerEnabled
	lcdcWindowTileMap
	lcdcIsLCDEnabled
)

// Represents the LCD Control (LCDC) register.
type Control uint8

// Is the LCD turned on?
func (c Control) IsLCDEnabled() bool {
	return HasBit(uint8(c), lcdcIsLCDEnabled)
}

// The tile map in use for the window layer.
func (c Control) WindowTileMap() uint8 {
	return GetBit(uint8(c), lcdcWindowTileMap)
}

// Is the window layer enabled?
func (c Control) IsWindowLayerEnabled() bool {
	return HasBit(uint8(c), lcdcIsBackgroundWindowLayerEnabled) && HasBit(uint8(c), lcdcIsWindowLayerEnabled)
}

// Which tile set is being used for the background and window layer?
func (c Control) BackgroundWindowTileSet() uint8 {
	return GetBit(uint8(c), lcdcBackgroundWindowTileSet)
}

// The tile map in use for the background layer.
func (c Control) BackgroundTileMap() uint8 {
	return GetBit(uint8(c), lcdcBackgroundTileMap)
}

// Are sprites 8x16 pixels in size?
func (c Control) AreSprites8x16() bool {
	return HasBit(uint8(c), lcdcAreSprites8x16)
}

// Is the sprite layer enabled?
func (c Control) IsSpriteLayerEnabled() bool {
	return HasBit(uint8(c), lcdcIsSpriteLayerEnabled)
}

// Is the background layer enabled?
func (c Control) IsBackgroundLayerEnabled() bool {
	return HasBit(uint8(c), lcdcIsBackgroundWindowLayerEnabled)
}
