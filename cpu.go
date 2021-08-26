package tamago

type CPU struct {
	state *State
	// Lookup table for instructions
	table  []Instruction
	xtable []Instruction
}

func NewCPU() *CPU {
	return &CPU{
		state:  NewState(),
		table:  ops[:],
		xtable: cbops[:],
	}
}
