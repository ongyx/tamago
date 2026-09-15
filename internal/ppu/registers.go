package ppu

import (
	"fmt"
)

const (
	// The start address of the LCD registers.
	RegisterStart = 0xFF40
	// The end address of the LCD registers.
	RegisterEnd = 0xFF4C

	lcdRegisterSize = RegisterEnd - RegisterStart

	lcdcAddr = 0xFF40
	statAddr = 0xFF41
	lyAddr   = 0xFF44
	lycAddr  = 0xFF45
	scyAddr  = 0xFF42
	scxAddr  = 0xFF43
	wyAddr   = 0xFF4A
	wxAddr   = 0xFF4B
	bgpAddr  = 0xFF47
	obp1Addr = 0xFF48
	obp2Addr = 0xFF49
)

// LCD registers used by the PPU.
type Registers struct {
	LCDC Control
	STAT Stat
	LY   uint8
	LYC  uint8
	SCY  uint8
	SCX  uint8
	WY   uint8
	WX   uint8
	BGP  Palette
	OBP1 Palette
	OBP2 Palette
}

// Creates a new set of LCD registers.
func NewRegisters() Registers {
	return Registers{}
}

// Returns the viewpoint coordinates as a point.
func (r *Registers) ViewportCoords() Point {
	return Point{r.SCX, r.SCY}
}

// Returns the window coordinates as a point.
func (r *Registers) WindowCoords() Point {
	x := r.WX
	if x >= 7 {
		x -= 7
	} else {
		x = 0
	}
	y := r.WY

	return Point{x, y}
}

// Returns the background palette.
func (r *Registers) BackgroundPalette() [4]uint8 {
	return r.BGP.Unpack()
}

// Returns the sprite paltette at the given index. If index is not 0 or 1, a panic occurs.
func (r *Registers) SpritePalette(index int) [4]uint8 {
	switch index {
	case 0:
		return r.OBP1.Unpack()
	case 1:
		return r.OBP2.Unpack()
	default:
		panic("index must be 0 or 1")
	}
}

// Reads a value from the LCD register at the address. The address must be between [RegisterStart] and [RegisterEnd].
func (r *Registers) Read(addr uint16) uint8 {
	switch addr {
	case lcdcAddr:
		return uint8(r.LCDC)
	case statAddr:
		return uint8(r.STAT)
	case lyAddr:
		return r.LY
	case lycAddr:
		return r.LYC
	case scyAddr:
		return r.SCY
	case scxAddr:
		return r.SCX
	case wyAddr:
		return r.WY
	case wxAddr:
		return r.WX
	case bgpAddr:
		return uint8(r.BGP)
	case obp1Addr:
		return uint8(r.OBP1)
	case obp2Addr:
		return uint8(r.OBP2)
	default:
		panic(fmt.Sprintf("unimplemented read from %#x", addr))
	}
}

// Writes a value to the LCD register at the address. The address must be between [RegisterStart] and [RegisterEnd].
func (r *Registers) Write(addr uint16, v uint8) {
	switch addr {
	case lcdcAddr:
		r.LCDC = Control(v)
	case statAddr:
		// MSB is ignored.
		r.STAT = Stat(v & 0x7F)
	case lyAddr:
		r.LY = v
	case lycAddr:
		r.LYC = v
	case scyAddr:
		r.SCY = v
	case scxAddr:
		r.SCX = v
	case wyAddr:
		r.WY = v
	case wxAddr:
		r.WX = v
	case bgpAddr:
		r.BGP = Palette(v)
	case obp1Addr:
		r.OBP1 = Palette(v)
	case obp2Addr:
		r.OBP2 = Palette(v)
	default:
		panic(fmt.Sprintf("unimplemented write to %#x", addr))
	}
}
