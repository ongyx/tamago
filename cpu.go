package tamago

import (
	"errors"
	"fmt"
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

// Start intepreting instructions from the (boot)rom.
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

// Start a basic debug shell.
func (c *CPU) DebugRun() {
	var loaded bool

	shell := NewShell()

	shell.Register("load", Command{
		"load a rom by filename",
		1,
		func(args []string) error {
			if err := c.Load(args[0]); err != nil {
				fmt.Println(err)
			} else {
				loaded = true
			}

			return nil
		},
	})

	shell.Register("loadboot", Command{
		"load a bootrom by filename",
		1,
		func(args []string) error {
			if err := c.LoadBoot(args[0]); err != nil {
				fmt.Println(err)
			} else {
				loaded = true
			}

			return nil
		},
	})

	shell.Register("step", Command{
		"execute the next instruction",
		0,
		func(args []string) error {
			if !loaded {
				fmt.Println("no rom loaded!")
			} else {
				c.step()
			}

			return nil
		},
	})

	shell.Register("show", Command{
		"show the register state",
		0,
		func(args []string) error {
			fmt.Println(c.String())
			return nil
		},
	})

	shell.Register("dump", Command{
		"dump the rom and ram to disk",
		0,
		func(args []string) error {
			return c.DebugDump()
		},
	})

	shell.Register("peek", Command{
		"peek the next instruction to execute",
		0,
		func(args []string) error {
			fmt.Println(c.table[c.PC].asm)
			return nil
		},
	})

	if err := shell.Prompt("> "); err != nil {
		fmt.Println(err)
	}
}
