package core

// Splits a word into two bytes.
func SplitWord(v uint16) (lo, hi uint8) {
	return uint8(v >> 8), uint8(v & 0xFF)
}

// Combines a word into two bytes.
func CombineWord(lo uint8, hi uint8) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}
