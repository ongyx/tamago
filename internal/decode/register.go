package decode

const (
	// No register is specified.
	RegisterNone Register = iota
	// General register A.
	RegisterA
	// General register B.
	RegisterB
	// General register C.
	RegisterC
	// General register D.
	RegisterD
	// General register E.
	RegisterE
	// General register H.
	RegisterH
	// General register L.
	RegisterL
	// Refers to the byte pointed at by HL.
	//
	// This is technically not a register, but is included here for ease of decoding.
	RegisterHLP
)

// An 8-bit, byte-sized register.
//
//go:generate stringer -type=Register
type Register uint8
