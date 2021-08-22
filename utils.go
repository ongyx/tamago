package tamago

import (
	"encoding/binary"
)

var (
	Endian = binary.LittleEndian

	U16 = struct{}{}
)

// Convert a boolean into 0 if false or 1 if true.
func bit(b bool) uint8 {
	var v uint8
	if b {
		v = 1
	}

	return v
}

// Convert two bytes into a single 16-bit unsigned int.
func (U16) From(lo, hi uint8) uint16 {
	return Endian.Uint16([]byte{lo, hi})
}

// Set a single 16-bit unsigned int into two bytes.
func (U16) To(lo, hi *uint8, v uint16) {
	buf := []byte{}
	Endian.PutUint16(buf, v)
	*lo = buf[0]
	*hi = buf[1]
}
