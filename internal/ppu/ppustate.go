package ppu

const (
	hBlankLen   = 204
	vBlankLen   = 456
	oamReadLen  = 80
	vramReadLen = 172

	vBlankLineStart = 143
	vBlankLineEnd   = vBlankLineStart + 10
)

// Contains the PPU state.
type PPUState struct {
	mode   Mode
	cycles int
	line   uint8
}

// Creates a new PPU state.
func NewPPUState() PPUState {
	return PPUState{mode: ModeHBlank}
}

// Returns the current mode.
func (l *PPUState) Mode() Mode {
	return l.mode
}

// Returns the number of M-cycles elapsed.
func (l *PPUState) Cycles() int {
	return l.cycles
}

// Returns the current scanline.
func (l *PPUState) Line() uint8 {
	return l.line
}

// Ticks the PPU state by a number of M-cycles. If the screen should be rendered, true is returned.
func (l *PPUState) Tick(cycles uint8) (render bool) {
	l.cycles += int(cycles)

	switch l.mode {
	case ModeHBlank:
		if l.cycles >= hBlankLen {
			l.cycles = 0
			l.line++

			if l.line == vBlankLineStart {
				// Vblank after frame has finished rendering
				l.mode = ModeVBlank
				render = true
			} else {
				l.mode = ModeOAMRead
			}
		}
	case ModeVBlank:
		if l.cycles >= vBlankLen {
			l.cycles = 0
			l.line++

			if l.line > vBlankLineEnd {
				l.mode = ModeOAMRead
				l.line = 0
			}
		}
	case ModeOAMRead:
		if l.cycles >= oamReadLen {
			l.cycles = 0
			l.mode = ModeVRAMRead
		}
	case ModeVRAMRead:
		if l.cycles >= vramReadLen {
			l.cycles = 0
			l.mode = ModeHBlank
		}
	}

	return render
}
