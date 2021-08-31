package tamago

type PPU struct {
	control, scrollX, scrollY, scanline, tick uint8
}

func NewPPU() *PPU {
	return &PPU{}
}

func (p *PPU) Read(addr uint16) uint8 {
	// TODO
}

func (p *PPU) Write(addr uint16, val uint8) {
	// TODO
}
