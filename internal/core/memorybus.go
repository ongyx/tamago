package core

// Handles read/write accesses to memory addresses.
type MemoryBus struct {
	memory []uint8
}

// Creates a new memory bus.
func NewMemoryBus() MemoryBus {
	return MemoryBus{memory: make([]uint8, 0xFFFF)}
}

// Reads a byte from memory.
func (mb *MemoryBus) Read(addr uint16) uint8 {
	return mb.memory[addr]
}

// Reads a word from memory.
func (mb *MemoryBus) ReadWord(addr uint16) uint16 {
	lo := mb.Read(addr)
	hi := mb.Read(addr + 1)
	return CombineWord(lo, hi)
}

// Writes a byte to memory.
func (mb *MemoryBus) Write(addr uint16, value uint8) {
	mb.memory[addr] = value
}

// Writes a word to memory.
func (mb *MemoryBus) WriteWord(addr uint16, value uint16) {
	lo, hi := SplitWord(value)
	mb.Write(addr, lo)
	mb.Write(addr+1, hi)
}
