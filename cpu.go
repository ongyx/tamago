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

func (c *CPU) fetch() uint8 {
	b := c.state.Read(c.state.PC)
	c.state.PC++
	return b
}

func (c *CPU) step() {
	if c.state.stopped {
		return
	}

	var ins Instruction

	opcode := c.fetch()
	if opcode == 0xCB {
		opcode = c.fetch()
		ins = c.xtable[opcode]
	} else {
		ins = c.table[opcode]
	}

	buf := make([]uint8, 2)
	for i := 1; i <= ins.length; i++ {
		buf = append(buf, c.fetch())
	}

	ins.fn(c.state, NewValue(buf))
}
