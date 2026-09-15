package ppu

const (
	// The PPU is in HBlank mode.
	ModeHBlank Mode = iota
	// The PPU is in VBlank mode.
	ModeVBlank
	// The PPU is in OAM read mode.
	ModeOAMRead
	// The PPU is in VRAM read mode.
	ModeVRAMRead
)

// A PPU mode.
type Mode uint8
