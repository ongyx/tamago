package tamago

import (
	"encoding/binary"
)

func getShort(r1 *uint8, r2 *uint8) uint16 {
	return binary.LittleEndian.Uint16([]uint8{*r1, *r2})
}

func setShort(r1 *uint8, r2 *uint8, val uint16) {
	buf := []uint8{}
	binary.LittleEndian.PutUint16(buf, val)

	*r1 = buf[0]
	*r2 = buf[1]
}
