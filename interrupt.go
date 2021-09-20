package tamago

import (
	"fmt"
)

const (
	VBlank uint8 = 1 << iota
	LCDStat
	Timer
	Serial
	Joypad
)

var priority = [...]uint8{VBlank, LCDStat, Timer, Serial, Joypad}

// An interrupt is an event that is handled before the next instruction executes.
type Interrupt struct {
	master             bool
	enabled, requested uint8
	handlers           map[uint8]func()
}

func NewInterrupt() *Interrupt {
	return &Interrupt{handlers: make(map[uint8]func())}
}

func (ir *Interrupt) register(flag uint8, fn func()) {
	ir.handlers[flag] = fn
}

func (ir *Interrupt) step() {
	flags := ir.enabled & ir.requested

	if ir.master && flags > 0 {
		for _, flag := range priority {

			if fn, ok := ir.handlers[flag]; ok {
				if (flags & flag) != 0 {
					ir.requested &^= flag
					ir.master = false
					fn()
				}

			} else {
				panic(fmt.Sprintf("no handler for flag %d", flag))
			}

		}
	}
}
