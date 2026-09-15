package core

// A cartridge loaded from disk.
type Cart struct {
	data []byte
}

// Creates a new cartridge from the data buffer. This retains a reference to it.
func NewCart(data []byte) *Cart {
	return &Cart{data}
}

// Reads a byte value at the address in the cart.
func (c *Cart) Read(addr uint16) uint8 {
	return c.data[addr]
}
