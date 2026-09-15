package util

// Gets the nth bit in the byte value. This panics if n is more than 7.
func GetBit(v uint8, n uint8) uint8 {
	if n > 7 {
		panic("n must be in the range [0, 7]")
	}

	return (v >> n) & 0x1
}

// Equivalent to [GetBit], but interprets the bit as a bool.
func HasBit(v uint8, n uint8) bool {
	return GetBit(v, n) != 0
}

// Sets or clears the nth bit in the byte value. This panics if n is more than 7.
func SetBit(v uint8, n uint8, b bool) uint8 {
	if n > 7 {
		panic("n must be in the range [0, 7]")
	}

	if b {
		return v | 1<<n
	} else {
		return v &^ (1 << n)
	}
}

// Splits a word into two bytes.
func SplitWord(v uint16) (hi, lo uint8) {
	return uint8(v >> 8), uint8(v & 0xFF)
}

// Combines a word into two bytes.
func CombineWord(hi uint8, lo uint8) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}
