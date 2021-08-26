package tamago

import (
	"encoding/binary"
	"log"
	"os"
)

var (
	Endian = binary.LittleEndian
	logger = log.New(os.Stdout, "", log.Lshortfile)

	U16 = u16{}
)

// Convert a boolean into 0 if false or 1 if true.
func tobit(b bool) uint8 {
	if b {
		return 1
	} else {
		return 0
	}
}

type u16 struct{}

// Convert two bytes into a single 16-bit unsigned int.
func (u16) From(lo, hi uint8) uint16 {
	return Endian.Uint16([]byte{lo, hi})
}

// Set a single 16-bit unsigned int into two bytes.
func (u16) To(lo, hi *uint8, v uint16) {
	buf := []byte{}
	Endian.PutUint16(buf, v)
	*lo = buf[0]
	*hi = buf[1]
}
