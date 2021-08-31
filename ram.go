package tamago

type RAM struct {
	buf []uint8
}

func NewRAM(size int) *RAM {
	return &RAM{buf: make([]uint8, size)}
}

func (r *RAM) Read(addr uint16) uint8 {
	// TODO: bank switching?
	return r.buf[addr]
}

func (r *RAM) Write(addr uint16, val uint8) {
	r.buf[addr] = val
}
