package tamago

import (
	"math/rand"
)

type IO struct {
	intr   *Interrupt
	input  *Input
	render *Render
}

func NewIO(render *Render) *IO {
	return &IO{
		intr:   NewInterrupt(),
		input:  NewInput(),
		render: render,
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
		return io.render.control

	case 0x0042:
		return io.render.scrollY

	case 0x0043:
		return io.render.scrollX

	case 0x0044:
		return io.render.scanline

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
		io.render.control = val

	case 0x0042:
		io.render.scrollY = val

	case 0x0043:
		io.render.scrollX = val

	case 0x0044:
		io.render.scanline = val

	case 0x00ff:
		io.intr.enable = val > 0

	}
}
