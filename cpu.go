package tamago

import (
	"errors"
	"os"
	"os/signal"
)

var (
	NoROMErr = errors.New("No bootrom and/or rom loaded!")
)

type CPU struct {
	*State

	// Lookup table for instructions
	table  []Instruction
	xtable []Instruction
}

func NewCPU(rr Renderer) *CPU {
	return &CPU{
		State:  NewState(rr),
		table:  ops[:],
		xtable: cbops[:],
	}
}

func (c *CPU) fetch() uint8 {
	b := c.Read(c.PC)
	c.PC++
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
	for i := 0; i < ins.length; i++ {
		buf[i] = c.fetch()
	}
	value := NewValue(buf)

	if c.PC == 0x100 {
		c.hasBoot = false
	}

	logger.Printf("[0x%x] executing %s", c.PC, ins.Asm(value))

	ins.fn(c.State, value)
	c.clock.Step(ins.cycles)
	c.render.step(c.clock)
}

func (c *CPU) Run() error {
	if !(c.hasBoot || c.hasROM) {
		return NoROMErr
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	for {
		select {

		case sig := <-ch:
			// ctrl-c
			logger.Println("caught signal " + sig.String())
			return nil

		default:
			if !c.stopped {
				c.step()
			}

		}
	}
}
