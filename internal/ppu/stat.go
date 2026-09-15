package ppu

import . "github.com/ongyx/tamago/internal/util"

const (
	statLYEQLYC = iota + 2
	statHBlankInterrupt
	statVBlankInterrupt
	statOAMInterrupt
	statLYEQLYCInterrupt

	// Lowest 2 bits.
	statModeMask = 0x3
)

// Represents the LCD Stat (STAT) register.
type Stat uint8

// The PPU mode.
func (s Stat) Mode() Mode {
	return Mode(s & statModeMask)
}

// Sets the PPU mode.
func (s *Stat) SetMode(m Mode) {
	sv := uint8(*s)
	mv := uint8(m)
	*s = Stat((sv &^ statModeMask) | (mv & statModeMask))
}

// Is LY equal to LYC?
func (s Stat) IsLYEQLYC() bool {
	return HasBit(uint8(s), statLYEQLYC)
}

// Sets if LY is equal to LYC.
func (s *Stat) SetLYEQLYC(b bool) {
	sv := uint8(*s)
	*s = Stat(SetBit(sv, statLYEQLYC, b))
}

// Is the HBLank interrupted enabled?
func (s Stat) IsHBlankInterruptEnabled() bool {
	return HasBit(uint8(s), statHBlankInterrupt)
}

// Is the VBlank interrupt enabled?
func (s Stat) IsVBlankInterruptEnabled() bool {
	return HasBit(uint8(s), statVBlankInterrupt)
}

// Is the OAM interrupt enabled?
func (s Stat) IsOAMInterruptEnabled() bool {
	return HasBit(uint8(s), statOAMInterrupt)
}

// Is the LY equal to LYC interrupt enabled?
func (s Stat) IsLYEQLYCInterruptEnabled() bool {
	return HasBit(uint8(s), statLYEQLYCInterrupt)
}
