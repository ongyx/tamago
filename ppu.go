package tamago

type PPU struct {
	control, scrollX, scrollY, scanline, tick uint8
}

func NewPPU() *PPU {
	return &PPU{}
}
