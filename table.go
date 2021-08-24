package tamago

type Table struct {
	ins []Instruction
}

func (t *Table) Lookup(opcode uint8) Instruction {
	return t.ins[opcode]
}
