package tamago

// A region represents a part of the address bus.
type Region interface {
	Read(addr uint16) uint8
	Write(addr uint16, val uint8)
}
