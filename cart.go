package tamago

// A cart is a read-only memory (ROM) region that contains a game's code/sprites.
// TODO: bank switching?
type Cart struct {
	*RAM
}

func NewCart() *Cart {
	return &Cart(NewRAM(0x8000))
}

func (c *Cart) Write(addr uint16, val uint8) {
	logger.Printf("can't write to cart! addr %d, value %d", addr, val)
}
