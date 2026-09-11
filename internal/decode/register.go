package decode

const (
	RegisterNone Register = iota
	RegisterA
	RegisterB
	RegisterC
	RegisterD
	RegisterE
	RegisterH
	RegisterL
	RegisterHLP
)

// An 8-bit wide register.
//
//go:generate stringer -type=Register
type Register uint8
