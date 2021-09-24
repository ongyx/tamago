package tamago

import (
	"encoding/binary"
	"log"
	"os"
	"strconv"
)

var (
	Endian = binary.LittleEndian
	logger = log.New(os.Stdout, "", 0)

	U16 = u16{}
)

func tobit(b bool) uint8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func tohex(n int64) string {
	return strconv.FormatInt(n, 16)
}

type u16 struct{}

// Convert two bytes into a single 16-bit unsigned int.
func (u16) From(lo, hi uint8) uint16 {
	return Endian.Uint16([]byte{lo, hi})
}

// Set a single 16-bit unsigned int into two bytes.
func (u16) To(lo, hi *uint8, v uint16) {
	buf := make([]byte, 2)
	Endian.PutUint16(buf, v)
	*lo = buf[0]
	*hi = buf[1]
}
