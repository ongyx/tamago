package core

import "github.com/ongyx/tamago/internal/ppu"

// Provides debug functions for a CPU.
type Debugger struct {
	cpu *CPU
}

// Creates a debugger for the CPU.
func NewDebugger(cpu *CPU) *Debugger {
	return &Debugger{cpu: cpu}
}

// Returns the PPU debugger.
func (d *Debugger) PPU() *ppu.Debugger {
	return ppu.NewDebugger(&d.cpu.bus.PPU)
}
