package decode

const (
	WordRegisterNone WordRegister = iota
	WordRegisterBC
	WordRegisterDE
	WordRegisterHL
	WordRegisterAF
	WordRegisterSP
)

// A 16-bit wide (word) register.
//
//go:generate stringer -type=WordRegister
type WordRegister uint16
