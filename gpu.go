package tamago

type GPU struct {
	control, scrollX, scrollY, scanline, tick uint8
}

func NewGPU() *GPU {
	return &GPU{}
}
