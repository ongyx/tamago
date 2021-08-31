package tamago

import (
	"math/rand"
)

type IO struct {
	input *Input
	intr  *Interrupt
	ppu   *PPU
}

func NewIO(ppu *PPU) *IO {
	return &IO{
		input: NewInput(),
		intr:  NewInterrupt(),
		ppu:   ppu,
	}
}

func (io *IO) Read(addr uint16) uint8 {
	switch addr {

	case 0x0000:
		return io.input.Poll()

	case 0x0004:
		// TODO: Implement actual timer?
		return uint8(rand.Intn(256))

	case 0x000f:
		return uint8(io.intr.flags)

	case 0x0040:
		return io.ppu.control

	case 0x0042:
		return io.ppu.scrollY

	case 0x0043:
		return io.ppu.scrollX

	case 0x0044:
		return io.ppu.scanline

	case 0x00ff:
		return tobit(io.intr.enable)

	}

	// TODO
	return 0
}

func (io *IO) Write(addr uint16, val uint8) {
	switch addr {

	case 0x0000:
		io.input.Select(val)

	case 0x000f:
		io.intr.flags = val

	case 0x0040:
		io.ppu.control = val

	case 0x0042:
		io.ppu.scrollY = val

	case 0x0043:
		io.ppu.scrollX = val

	case 0x0044:
		io.ppu.scanline = val

	case 0x00ff:
		io.intr.enable = val > 0

	}
}
