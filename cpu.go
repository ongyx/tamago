package tamago

import (
	"os"
	"os/signal"
)

type CPU struct {
	state *State
	// Lookup table for instructions
	table  []Instruction
	xtable []Instruction
}

func NewCPU(rr Renderer) *CPU {
	return &CPU{
		state:  NewState(rr),
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

func (c *CPU) Run() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	for {
		select {

		case sig := <-ch:
			// ctrl-c
			logger.Println("caught signal " + sig.String())
			return

		default:
			if !c.state.stopped {
				c.step()
			}

		}
	}
}

func (c *CPU) LoadROM(name string) error {
	f, err := os.Open(name)
	defer f.Close()

	if err != nil {
		return err
	}

	if e := c.state.cart.Load(f); e != nil {
		return e
	}

	return nil
}
