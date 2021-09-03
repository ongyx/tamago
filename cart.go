package tamago

import (
	"errors"
	"io"
)

var TooLargeErr = errors.New("ROM is too large! (size > 32762 bytes)")

// A cart is a read-only memory (ROM) region that contains a game's code/sprites.
// TODO: bank switching?
type Cart struct {
	*RAM
}

func NewCart() *Cart {
	return &Cart{NewRAM(0x8000)}
}

// Load a cartriage from a file and return an error if any.
func (c *Cart) Load(rom io.Reader) error {
	buf, err := io.ReadAll(rom)
	if err != nil {
		return err
	}

	// TODO: bank switching?
	if len(buf) > len(c.buf) {
		return TooLargeErr
	}

	copy(buf, c.buf)

	return nil
}

func (c *Cart) Write(addr uint16, val uint8) {
	logger.Printf("can't write to cart! addr %d, value %d", addr, val)
}
