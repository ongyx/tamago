package decode

const (
	RLC PrefixOpcode = iota
	RRC
	RL
	RR
	SLA
	SRA
	SWAP
	SRL
	BIT
	RES
	SET
)

// An instruction opcode prefixed with 0xCB.
//
//go:generate stringer -type PrefixOpcode
type PrefixOpcode uint8
