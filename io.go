package tamago

import (
	"math/rand"
)

type IO struct {
	rr    Renderer
	rrst  *Status
	intr  *Interrupt
	input *Input
}

func NewIO(rr Renderer) *IO {
	return &IO{
		rr:    rr,
		rrst:  rr.Status(),
		intr:  NewInterrupt(),
		input: NewInput(),
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
		return io.rrst.control

	case 0x0042:
		return io.rrst.scrollY

	case 0x0043:
		return io.rrst.scrollX

	case 0x0044:
		return io.rrst.scanline

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
		io.rrst.control = val

	case 0x0042:
		io.rrst.scrollY = val

	case 0x0043:
		io.rrst.scrollX = val

	case 0x0044:
		io.rrst.scanline = val

	case 0x00ff:
		io.intr.enable = val > 0

	}
}
