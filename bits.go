package tamago

type Bits struct {
	v *uint8
}

func NewBits(v *uint8) *Bits {
	return &Bits{v}
}

// Get the bit at pos.
func (b *Bits) Get(pos int) uint8 {
	return *b.v & (1 << pos)
}

// Set the bit at pos.
func (b *Bits) Set(pos int) {
	*b.v |= (1 << pos)
}
