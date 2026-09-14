package core

import "github.com/ongyx/tamago/internal/decode"

// Executes instructions from memory.
type CPU struct {
	registers *Registers
	bus       Bus
	alu       ALU

	isHalted bool
	eiDelay  int
}

// Creates a new CPU, with register values initialized to their state after the boot ROM finishes.
func NewCPU() *CPU {
	// re https://aquova.net/emudev/gb/11-final-misc.html
	var f ALUFlags
	f.Decode(0xB0)
	rs := &Registers{
		A: 0x1,
		B: 0x0,
		C: 0x13,
		D: 0x00,
		E: 0xD8,
		// Equivalent to 0xB0
		F: f,
		H: 0x1,
		L: 0x4D,
	}

	b := NewBus(rs)
	b.Write(0xFF10, 0x80)
	b.Write(0xFF11, 0xBF)
	b.Write(0xFF12, 0xF3)
	b.Write(0xFF14, 0xBF)
	b.Write(0xFF16, 0x3F)
	b.Write(0xFF19, 0xBF)
	b.Write(0xFF1A, 0x7F)
	b.Write(0xFF1B, 0xFF)
	b.Write(0xFF1C, 0x9F)
	b.Write(0xFF1E, 0xBF)
	b.Write(0xFF20, 0xFF)
	b.Write(0xFF23, 0xBF)
	b.Write(0xFF24, 0x77)
	b.Write(0xFF25, 0xF3)
	b.Write(0xFF26, 0xF1)
	b.Write(0xFF40, 0x91)
	b.Write(0xFF47, 0xFC)
	b.Write(0xFF48, 0xFF)
	b.Write(0xFF49, 0xFF)

	alu := NewALU(rs)

	return &CPU{
		registers: rs,
		bus:       b,
		alu:       alu,
	}
}

// Executes a fetch-decode-execute cycle, returning the number of M-cycles the instruction takes to finish executing.
func (c *CPU) Tick() (cycles int, err error) {
	if c.isHalted {
		// Do nothing.
		cycles = 1
	} else {
		ins, err := decode.DecodeInstruction(c.fetchByte())
		if err != nil {
			return 0, err
		}

		cycles = c.execute(ins)
	}

	if c.eiDelay > 0 {
		c.eiDelay--
		if c.eiDelay == 0 {
			c.enableInterrupt()
		}
	}

	c.handleInterrupt()

	return cycles, err
}

func (c *CPU) execute(ins decode.Instruction) int {
	cycles := 1
	// Generally, instructions involving a load/store with (HL) require 2 cycles to execute.
	if ins.SrcOperand == decode.RegisterHLP || ins.DstOperand == decode.RegisterHLP {
		cycles = 2
	}

	switch ins.Opcode {
	// Miscellaneous
	case decode.NOP:
		// no-op
	case decode.PREFIX:
		pins := decode.DecodePrefixInstruction(c.fetchByte())
		return c.executePrefix(pins)
	case decode.STOP:
		panic("unimplemented")
	case decode.HALT:
		c.isHalted = true

	// Interrupts
	case decode.DI:
		c.disableInterrupt()
	case decode.EI:
		// Due to a hardware quirk, EI only takes effect after the next instruction.
		c.eiDelay = 2

	// BCD/Carry
	case decode.DAA:
		c.alu.AdjustBCDA()
	case decode.SCF:
		c.alu.SetCarry()
	case decode.CPL:
		c.alu.ComplementA()
	case decode.CCF:
		c.alu.FlipCarry()

	// Jumps
	case decode.JR_CC_S8:
		o := c.fetchSigned()
		if c.evaluateCondition(ins.Condition) {
			cycles = 3
			c.jumpRelative(o)
		} else {
			cycles = 2
		}
	case decode.JP_CC_A16:
		addr := c.fetchWord()
		if c.evaluateCondition(ins.Condition) {
			cycles = 4
			c.jump(addr)
		} else {
			cycles = 3
		}
	case decode.JP_HL:
		c.jump(c.registers.HL())

	// Stack manipulation
	case decode.ADD_SP_S8:
		cycles = 4
		o := c.fetchSigned()
		c.registers.SP = c.alu.AddSP(o)
	case decode.PUSH_R16:
		cycles = 4
		v := c.loadWordRegister(ins.WordOperand)
		c.push(v)
	case decode.CALL_CC_A16:
		addr := c.fetchWord()
		if c.evaluateCondition(ins.Condition) {
			cycles = 6
			c.call(addr)
		} else {
			cycles = 3
		}
	case decode.RST:
		cycles = 4
		addr := uint16(ins.RSTVector)
		c.call(addr)
	case decode.POP_R16:
		cycles = 3
		v := c.pop()
		c.storeWordRegister(ins.WordOperand, v)
	case decode.RET_CC:
		if ins.Condition == decode.ConditionNone {
			// RET is slightly faster than RET CC when the branch is taken.
			cycles = 4
			c.ret()
		} else if c.evaluateCondition(ins.Condition) {
			cycles = 5
			c.ret()
		} else {
			cycles = 2
		}
	case decode.RETI:
		cycles = 4
		c.enableInterrupt()
		c.ret()

	// 8-bit register arithmetic
	case decode.INC_R8:
		if ins.SrcOperand == decode.RegisterHLP {
			cycles = 3
		}
		v := c.loadRegister(ins.SrcOperand)
		r := c.alu.Inc(v)
		c.storeRegister(ins.SrcOperand, r)
	case decode.DEC_R8:
		if ins.SrcOperand == decode.RegisterHLP {
			cycles = 3
		}
		v := c.loadRegister(ins.SrcOperand)
		r := c.alu.Dec(v)
		c.storeRegister(ins.SrcOperand, r)
	case decode.ADD_R8:
		v := c.loadRegister(ins.SrcOperand)
		c.alu.Add(v, false)
	case decode.ADC_R8:
		v := c.loadRegister(ins.SrcOperand)
		c.alu.Add(v, true)
	case decode.SUB_R8:
		v := c.loadRegister(ins.SrcOperand)
		c.alu.Sub(v, false)
	case decode.SBC_R8:
		v := c.loadRegister(ins.SrcOperand)
		c.alu.Sub(v, true)
	case decode.AND_R8:
		v := c.loadRegister(ins.SrcOperand)
		c.alu.And(v)
	case decode.XOR_R8:
		v := c.loadRegister(ins.SrcOperand)
		c.alu.Xor(v)
	case decode.OR_R8:
		v := c.loadRegister(ins.SrcOperand)
		c.alu.Or(v)
	case decode.CP_R8:
		v := c.loadRegister(ins.SrcOperand)
		c.alu.Cp(v)

	// 16-bit arithmetic
	case decode.INC_R16:
		cycles = 2
		v := c.loadWordRegister(ins.WordOperand)
		v++
		c.storeWordRegister(ins.WordOperand, v)
	case decode.DEC_R16:
		cycles = 2
		v := c.loadWordRegister(ins.WordOperand)
		v--
		c.storeWordRegister(ins.WordOperand, v)
	case decode.ADD_HL_R16:
		cycles = 2
		v := c.loadWordRegister(ins.WordOperand)
		c.alu.AddHL(v)

	// Register A bitshifts
	case decode.RLCA:
		c.registers.A = c.alu.RotateLeft(c.registers.A, false, false)
	case decode.RLA:
		c.registers.A = c.alu.RotateLeft(c.registers.A, true, false)
	case decode.RRCA:
		c.registers.A = c.alu.RotateRight(c.registers.A, false, false)
	case decode.RRA:
		c.registers.A = c.alu.RotateRight(c.registers.A, true, false)

	// 8-bit register arithmetic with immediate operand
	case decode.ADD_D8:
		cycles = 2
		v := c.fetchByte()
		c.alu.Add(v, false)
	case decode.ADC_D8:
		cycles = 2
		v := c.fetchByte()
		c.alu.Add(v, true)
	case decode.SUB_D8:
		cycles = 2
		v := c.fetchByte()
		c.alu.Sub(v, false)
	case decode.SBC_D8:
		cycles = 2
		v := c.fetchByte()
		c.alu.Sub(v, true)
	case decode.AND_D8:
		cycles = 2
		v := c.fetchByte()
		c.alu.And(v)
	case decode.XOR_D8:
		cycles = 2
		v := c.fetchByte()
		c.alu.Xor(v)
	case decode.OR_D8:
		cycles = 2
		v := c.fetchByte()
		c.alu.Or(v)
	case decode.CP_D8:
		cycles = 2
		v := c.fetchByte()
		c.alu.Cp(v)

	// 8-bit loads
	case decode.LD_R8_R8:
		if ins.SrcOperand == ins.DstOperand {
			// No-op.
			break
		}

		v := c.loadRegister(ins.SrcOperand)
		c.storeRegister(ins.DstOperand, v)
	case decode.LD_R8_D8:
		cycles = 2
		v := c.fetchByte()
		c.storeRegister(ins.DstOperand, v)
	case decode.LD_A8_A:
		cycles = 3
		addr := 0xFF00 + uint16(c.fetchByte())
		c.bus.Write(addr, c.registers.A)
	case decode.LD_A_A8:
		cycles = 3
		addr := 0xFF00 + uint16(c.fetchByte())
		c.registers.A = c.bus.Read(addr)
	case decode.LD_CP_A:
		cycles = 2
		addr := 0xFF00 + uint16(c.registers.C)
		c.bus.Write(addr, c.registers.A)
	case decode.LD_A_CP:
		cycles = 2
		addr := 0xFF00 + uint16(c.registers.C)
		c.registers.A = c.bus.Read(addr)
	case decode.LD_A16_A:
		cycles = 4
		addr := c.fetchWord()
		c.bus.Write(addr, c.registers.A)
	case decode.LD_A_A16:
		cycles = 4
		addr := c.fetchWord()
		c.registers.A = c.bus.Read(addr)
	case decode.LD_R16P_A:
		cycles = 2
		addr := c.loadWordRegister(ins.WordOperand)
		c.bus.Write(addr, c.registers.A)
	case decode.LD_HLPI_A:
		cycles = 2
		hl := c.registers.HL()
		c.bus.Write(hl, c.registers.A)
		c.registers.SetHL(hl + 1)
	case decode.LD_HLPD_A:
		cycles = 2
		hl := c.registers.HL()
		c.bus.Write(hl, c.registers.A)
		c.registers.SetHL(hl - 1)
	case decode.LD_A_R16P:
		cycles = 2
		addr := c.loadWordRegister(ins.WordOperand)
		c.registers.A = c.bus.Read(addr)
	case decode.LD_A_HLPI:
		cycles = 2
		hl := c.registers.HL()
		c.registers.A = c.bus.Read(hl)
		c.registers.SetHL(hl + 1)
	case decode.LD_A_HLPD:
		cycles = 2
		hl := c.registers.HL()
		c.registers.A = c.bus.Read(hl)
		c.registers.SetHL(hl - 1)

	// 16-bit loads
	case decode.LD_R16_D16:
		cycles = 3
		v := c.fetchWord()
		c.storeWordRegister(ins.WordOperand, v)
	case decode.LD_A16_SP:
		cycles = 5
		addr := c.fetchWord()
		c.bus.WriteWord(addr, c.registers.SP)
	case decode.LD_HL_SP_S8:
		cycles = 3
		o := c.fetchSigned()
		c.registers.SetHL(c.alu.AddSP(o))
	case decode.LD_SP_HL:
		cycles = 2
		c.registers.SP = c.registers.HL()

	default:
		panic("unknown opcode " + ins.Opcode.String())
	}

	return cycles
}

func (c *CPU) executePrefix(pins decode.PrefixInstruction) int {
	v := c.loadRegister(pins.Operand)

	// All prefixed instructions take 2 cycles to execute, except for those interacting with (HL).
	cycles := 2
	if pins.Operand == decode.RegisterHLP {
		if pins.Opcode == decode.BIT {
			// BIT doesn't write back to (HL) so it takes only 3 cycles.
			cycles = 3
		} else {
			cycles = 4
		}
	}

	var r uint8
	switch pins.Opcode {
	// The mnemonics are a bit misleading: RLC/RRC does not rotate through the carry bit, but RL/RR does.
	case decode.RLC:
		r = c.alu.RotateLeft(v, false, true)
	case decode.RRC:
		r = c.alu.RotateRight(v, false, true)
	case decode.RL:
		r = c.alu.RotateLeft(v, true, true)
	case decode.RR:
		r = c.alu.RotateRight(v, true, true)
	case decode.SLA:
		r = c.alu.ShiftLeft(v)
	case decode.SRA:
		r = c.alu.ShiftRight(v, true)
	case decode.SWAP:
		r = c.alu.Swap(v)
	case decode.SRL:
		r = c.alu.ShiftRight(v, false)
	case decode.BIT:
		c.alu.TestBit(v, pins.Bit)
	case decode.RES:
		r = c.alu.ClearBit(v, pins.Bit)
	case decode.SET:
		r = c.alu.SetBit(v, pins.Bit)
	default:
		panic("unknown prefix opcode " + pins.Opcode.String())
	}

	c.storeRegister(pins.Operand, r)

	return cycles
}

func (c *CPU) fetchByte() uint8 {
	v := c.bus.Read(c.registers.PC)
	c.registers.PC++
	return v
}

func (c *CPU) fetchSigned() int8 {
	// Interpret the operand as a two's complement binary number.
	return int8(c.fetchByte())
}

func (c *CPU) fetchWord() uint16 {
	lo := c.fetchByte()
	hi := c.fetchByte()
	return CombineWord(hi, lo)
}

func (c *CPU) evaluateCondition(cc decode.Condition) bool {
	switch cc {
	case decode.ConditionNone:
		return true
	case decode.ConditionZero:
		return c.registers.F.Zero
	case decode.ConditionNotZero:
		return !c.registers.F.Zero
	case decode.ConditionCarry:
		return c.registers.F.Carry
	case decode.ConditionNotCarry:
		return !c.registers.F.Carry
	default:
		panic("condition is invalid: " + cc.String())
	}
}

func (c *CPU) loadRegister(r decode.Register) uint8 {
	switch r {
	case decode.RegisterA:
		return c.registers.A
	case decode.RegisterB:
		return c.registers.B
	case decode.RegisterC:
		return c.registers.C
	case decode.RegisterD:
		return c.registers.D
	case decode.RegisterE:
		return c.registers.E
	case decode.RegisterH:
		return c.registers.H
	case decode.RegisterL:
		return c.registers.L
	case decode.RegisterHLP:
		return c.bus.Read(c.registers.HL())
	default:
		panic("register is invalid: " + r.String())
	}
}

func (c *CPU) storeRegister(r decode.Register, v uint8) {
	switch r {
	case decode.RegisterA:
		c.registers.A = v
	case decode.RegisterB:
		c.registers.B = v
	case decode.RegisterC:
		c.registers.C = v
	case decode.RegisterD:
		c.registers.D = v
	case decode.RegisterE:
		c.registers.E = v
	case decode.RegisterH:
		c.registers.H = v
	case decode.RegisterL:
		c.registers.L = v
	case decode.RegisterHLP:
		c.bus.Write(c.registers.HL(), v)
	default:
		panic("register is invalid: " + r.String())
	}
}

func (c *CPU) loadWordRegister(wr decode.WordRegister) uint16 {
	switch wr {
	case decode.WordRegisterBC:
		return c.registers.BC()
	case decode.WordRegisterDE:
		return c.registers.DE()
	case decode.WordRegisterHL:
		return c.registers.HL()
	case decode.WordRegisterAF:
		return c.registers.AF()
	case decode.WordRegisterSP:
		return c.registers.SP
	default:
		panic("word register is invalid: " + wr.String())
	}
}

func (c *CPU) storeWordRegister(wr decode.WordRegister, v uint16) {
	switch wr {
	case decode.WordRegisterBC:
		c.registers.SetBC(v)
	case decode.WordRegisterDE:
		c.registers.SetDE(v)
	case decode.WordRegisterHL:
		c.registers.SetHL(v)
	case decode.WordRegisterAF:
		c.registers.SetAF(v)
	case decode.WordRegisterSP:
		c.registers.SP = v
	default:
		panic("word register is invalid: " + wr.String())
	}
}

func (c *CPU) jump(addr uint16) {
	c.registers.PC = addr
}

func (c *CPU) jumpRelative(offset int8) {
	addr := uint16(int16(c.registers.PC) + int16(offset))
	c.jump(addr)
}

func (c *CPU) push(v uint16) {
	hi, lo := SplitWord(v)
	// Stack grows toward a lower address on the Game Boy.
	c.registers.SP--
	c.bus.Write(c.registers.SP, hi)
	c.registers.SP--
	c.bus.Write(c.registers.SP, lo)
}

func (c *CPU) pop() uint16 {
	lo := c.bus.Read(c.registers.SP)
	c.registers.SP++
	hi := c.bus.Read(c.registers.SP)
	c.registers.SP++

	return CombineWord(hi, lo)
}

func (c *CPU) call(addr uint16) {
	c.push(c.registers.PC)
	c.jump(addr)
}

func (c *CPU) ret() {
	addr := c.pop()
	c.jump(addr)
}

func (c *CPU) enableInterrupt() {
	c.registers.IME = true
}

func (c *CPU) disableInterrupt() {
	c.registers.IME = false
}

func (c *CPU) handleInterrupt() {
	iv := c.registers.CheckInterrupt()
	if iv != InterruptVectorNone {
		// Interrupts wake up the CPU from halt, even if IME is false.
		c.isHalted = false

		if c.registers.IME {
			c.disableInterrupt()
			// Continue execution from the interrupt handler.
			c.call(uint16(iv))
		}
	}
}
