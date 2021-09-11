package tamago

// Bits defines bit-level operations on a unsigned 8-bit integer.
type Bits struct {
	uint8
}

// Get the bit at pos as a bool.
func (b *Bits) At(pos int) bool {
	return (b.uint8 & (1 << pos)) > 0
}

// Set the bit at pos.
func (b *Bits) Set(pos int) {
	b.uint8 |= (1 << pos)
}

// Clear the bit at pos.
func (b *Bits) Clear(pos int) {
	b.uint8 &^= (1 << pos)
}
