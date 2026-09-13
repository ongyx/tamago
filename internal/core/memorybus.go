package core

// Handles read/write accesses to memory addresses.
type MemoryBus struct {
	registers *Registers
	memory    []uint8
}

// Creates a new memory bus.
func NewMemoryBus(registers *Registers) MemoryBus {
	return MemoryBus{
		registers: registers,
		memory:    make([]uint8, 0xFFFF),
	}
}

// Reads a byte from memory.
func (mb *MemoryBus) Read(addr uint16) uint8 {
	switch addr {
	case 0xFF0F:
		return mb.registers.IR.Encode()
	case 0xFFFF:
		return mb.registers.IE.Encode()
	}

	return mb.memory[addr]
}

// Reads a word from memory.
func (mb *MemoryBus) ReadWord(addr uint16) uint16 {
	lo := mb.Read(addr)
	hi := mb.Read(addr + 1)
	return CombineWord(hi, lo)
}

// Writes a byte to memory.
func (mb *MemoryBus) Write(addr uint16, value uint8) {
	switch addr {
	case 0xFF0F:
		mb.registers.IR.Decode(value)
	case 0xFFFF:
		mb.registers.IE.Decode(value)
	}

	mb.memory[addr] = value
}

// Writes a word to memory.
func (mb *MemoryBus) WriteWord(addr uint16, value uint16) {
	hi, lo := SplitWord(value)
	mb.Write(addr, lo)
	mb.Write(addr+1, hi)
}
