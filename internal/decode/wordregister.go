package decode

const (
	// No word register is specified.
	WordRegisterNone WordRegister = iota
	// General word register BC. This pairs register B and C.
	WordRegisterBC
	// General word register DE. This pairs register D and E.
	WordRegisterDE
	// General word register HL. This pairs register H and L.
	WordRegisterHL
	// General word register AF. This pairs register A and F.
	WordRegisterAF
	// Special word register SP.
	WordRegisterSP
)

// A 16-bit, word-sized register.
//
//go:generate stringer -type=WordRegister
type WordRegister uint16
