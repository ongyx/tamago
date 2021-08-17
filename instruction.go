package tamago

type execute func(cpu *CPU, value interface{})

// An instruction to execute in the CPU.
// Assembly is the instruction in assembly as text (i.e nop, etc.)
// Length is the total length in bytes of the instruction's arguments (0-2).
// Cycles is the amount of transistor cycles taken to execute the instruction.
// Execute is the function to run when this instruction is executed where cpu is the CPU and value may be nil (Length = 0), a uint8 (1), or a uint16 (2).
type Instruction struct {
	Assembly string
	Length   uint8
	Cycles   uint8
	Execute  execute
}

var Table = [256]Instruction{

	// 0x0
	{
		"NOP", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x1
	{
		"LD BC,u16", 2, 12,
		func(cpu *CPU, value interface{}) {
			cpu.reg.SetBC(value.(uint16))
		},
	},

	// 0x2
	{
		"LD (BC),A", 0, 8,
		func(cpu *CPU, value interface{}) {
			cpu.Write(cpu.reg.BC(), cpu.reg.A)
		},
	},

	// 0x3
	{
		"INC BC", 0, 8,
		func(cpu *CPU, value interface{}) {
			cpu.reg.SetBC(cpu.reg.BC() + 1)
		},
	},

	// 0x4
	{
		"INC B", 0, 4,
		func(cpu *CPU, value interface{}) {
			cpu.reg.Increment(&cpu.reg.B)
		},
	},

	// 0x5
	{
		"DEC B", 0, 4,
		func(cpu *CPU, value interface{}) {
			cpu.reg.Decrement(&cpu.reg.B)
		},
	},

	// 0x6
	{
		"LD B,u8", 1, 8,
		func(cpu *CPU, value interface{}) {
			cpu.reg.B = value.(uint8)
		},
	},

	// 0x7
	{
		"RLCA", 0, 4,
		func(cpu *CPU, value interface{}) {
			// Take only the first bit and shift all the way to the right
			carry := (cpu.reg.A & 0x80) >> 7

			if carry != 0 {
				// carry is now 0b1, so set the carry flag
				cpu.Setf(Flag.Carry)
			} else {
				cpu.Clearf(Flag.Carry)
			}

			// Shift A left 1 bit and add back the carry.
			cpu.reg.A <<= 1
			cpu.reg.A += carry

			// Clear all other flags.
			cpu.ClearFlag(HalfCarryFlag | NegativeFlag | ZeroFlag)
		},
	},

	// 0x8
	{
		"LD (u16),SP", 2, 20,
		func(cpu *CPU, value interface{}) {
			cpu.WriteShort(cpu.reg.SP, value.(uint16))
		},
	},

	// 0x9
	{
		"ADD HL,BC", 0, 8,
		func(cpu *CPU, value interface{}) {
			hl := cpu.reg.HL()
			cpu.reg.AddShort(&hl, cpu.reg.BC())

			cpu.reg.SetHL(hl)
		},
	},

	// 0xa
	{
		"LD A,(BC)", 0, 8,
		func(cpu *CPU, value interface{}) {
			cpu.reg.A = cpu.Read(value.(uint8))
		},
	},

	// 0xb
	{
		"DEC BC", 0, 8,
		func(cpu *CPU, value interface{}) {
			cpu.reg.SetBC(cpu.reg.BC() - 1)
		},
	},

	// 0xc
	{
		"INC C", 0, 4,
		func(cpu *CPU, value interface{}) {
			cpu.reg.Increment(&cpu.reg.C)
		},
	},

	// 0xd
	{
		"DEC C", 0, 4,
		func(cpu *CPU, value interface{}) {
			cpu.reg.Decrement(&cpu.reg.C)
		},
	},

	// 0xe
	{
		"LD C,u8", 1, 8,
		func(cpu *CPU, value interface{}) {
			cpu.reg.C = value.(uint8)
		},
	},

	// 0xf
	{
		"RRCA", 0, 4,
		func(cpu *CPU, value interface{}) {
			// Take only the last bit
			carry := cpu.reg.A & 0x01

			if carry != 0 {
				// carry is now 0b1, so set the carry flag
				cpu.Setf(Flag.Carry)
			} else {
				cpu.Clearf(Flag.Carry)
			}

			// Shift A right 1 bit and put back the bits.
			cpu.reg.A >>= 1
			if carry != 0 {
				cpu.reg.A |= 0x80
			}

			// Clear all other flags.
			cpu.ClearFlag(Flag.HalfCarryFlag | Flag.Negative | Flag.Zero)
		},
	},

	// 0x10
	{
		"STOP", 1, 4,
		func(cpu *CPU, value interface{}) {
			cpu.stopped = true
		},
	},

	// 0x11
	{
		"LD DE,u16", 2, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x12
	{
		"LD (DE),A", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x13
	{
		"INC DE", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x14
	{
		"INC D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x15
	{
		"DEC D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x16
	{
		"LD D,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x17
	{
		"RLA", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x18
	{
		"JR i8", 1, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x19
	{
		"ADD HL,DE", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x1a
	{
		"LD A,(DE)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x1b
	{
		"DEC DE", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x1c
	{
		"INC E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x1d
	{
		"DEC E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x1e
	{
		"LD E,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x1f
	{
		"RRA", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x20
	{
		"JR NZ,i8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x21
	{
		"LD HL,u16", 2, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x22
	{
		"LD (HL+),A", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x23
	{
		"INC HL", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x24
	{
		"INC H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x25
	{
		"DEC H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x26
	{
		"LD H,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x27
	{
		"DAA", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x28
	{
		"JR Z,i8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x29
	{
		"ADD HL,HL", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x2a
	{
		"LD A,(HL+)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x2b
	{
		"DEC HL", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x2c
	{
		"INC L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x2d
	{
		"DEC L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x2e
	{
		"LD L,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x2f
	{
		"CPL", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x30
	{
		"JR NC,i8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x31
	{
		"LD SP,u16", 2, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x32
	{
		"LD (HL-),A", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x33
	{
		"INC SP", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x34
	{
		"INC (HL)", 0, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x35
	{
		"DEC (HL)", 0, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x36
	{
		"LD (HL),u8", 1, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x37
	{
		"SCF", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x38
	{
		"JR C,i8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x39
	{
		"ADD HL,SP", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x3a
	{
		"LD A,(HL-)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x3b
	{
		"DEC SP", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x3c
	{
		"INC A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x3d
	{
		"DEC A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x3e
	{
		"LD A,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x3f
	{
		"CCF", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x40
	{
		"LD B,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x41
	{
		"LD B,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x42
	{
		"LD B,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x43
	{
		"LD B,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x44
	{
		"LD B,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x45
	{
		"LD B,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x46
	{
		"LD B,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x47
	{
		"LD B,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x48
	{
		"LD C,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x49
	{
		"LD C,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x4a
	{
		"LD C,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x4b
	{
		"LD C,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x4c
	{
		"LD C,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x4d
	{
		"LD C,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x4e
	{
		"LD C,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x4f
	{
		"LD C,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x50
	{
		"LD D,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x51
	{
		"LD D,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x52
	{
		"LD D,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x53
	{
		"LD D,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x54
	{
		"LD D,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x55
	{
		"LD D,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x56
	{
		"LD D,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x57
	{
		"LD D,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x58
	{
		"LD E,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x59
	{
		"LD E,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x5a
	{
		"LD E,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x5b
	{
		"LD E,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x5c
	{
		"LD E,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x5d
	{
		"LD E,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x5e
	{
		"LD E,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x5f
	{
		"LD E,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x60
	{
		"LD H,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x61
	{
		"LD H,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x62
	{
		"LD H,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x63
	{
		"LD H,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x64
	{
		"LD H,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x65
	{
		"LD H,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x66
	{
		"LD H,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x67
	{
		"LD H,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x68
	{
		"LD L,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x69
	{
		"LD L,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x6a
	{
		"LD L,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x6b
	{
		"LD L,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x6c
	{
		"LD L,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x6d
	{
		"LD L,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x6e
	{
		"LD L,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x6f
	{
		"LD L,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x70
	{
		"LD (HL),B", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x71
	{
		"LD (HL),C", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x72
	{
		"LD (HL),D", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x73
	{
		"LD (HL),E", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x74
	{
		"LD (HL),H", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x75
	{
		"LD (HL),L", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x76
	{
		"HALT", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x77
	{
		"LD (HL),A", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x78
	{
		"LD A,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x79
	{
		"LD A,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x7a
	{
		"LD A,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x7b
	{
		"LD A,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x7c
	{
		"LD A,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x7d
	{
		"LD A,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x7e
	{
		"LD A,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x7f
	{
		"LD A,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x80
	{
		"ADD A,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x81
	{
		"ADD A,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x82
	{
		"ADD A,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x83
	{
		"ADD A,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x84
	{
		"ADD A,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x85
	{
		"ADD A,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x86
	{
		"ADD A,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x87
	{
		"ADD A,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x88
	{
		"ADC A,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x89
	{
		"ADC A,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x8a
	{
		"ADC A,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x8b
	{
		"ADC A,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x8c
	{
		"ADC A,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x8d
	{
		"ADC A,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x8e
	{
		"ADC A,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x8f
	{
		"ADC A,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x90
	{
		"SUB A,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x91
	{
		"SUB A,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x92
	{
		"SUB A,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x93
	{
		"SUB A,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x94
	{
		"SUB A,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x95
	{
		"SUB A,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x96
	{
		"SUB A,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x97
	{
		"SUB A,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x98
	{
		"SBC A,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x99
	{
		"SBC A,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x9a
	{
		"SBC A,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x9b
	{
		"SBC A,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x9c
	{
		"SBC A,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x9d
	{
		"SBC A,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x9e
	{
		"SBC A,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0x9f
	{
		"SBC A,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xa0
	{
		"AND A,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xa1
	{
		"AND A,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xa2
	{
		"AND A,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xa3
	{
		"AND A,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xa4
	{
		"AND A,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xa5
	{
		"AND A,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xa6
	{
		"AND A,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xa7
	{
		"AND A,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xa8
	{
		"XOR A,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xa9
	{
		"XOR A,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xaa
	{
		"XOR A,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xab
	{
		"XOR A,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xac
	{
		"XOR A,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xad
	{
		"XOR A,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xae
	{
		"XOR A,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xaf
	{
		"XOR A,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xb0
	{
		"OR A,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xb1
	{
		"OR A,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xb2
	{
		"OR A,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xb3
	{
		"OR A,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xb4
	{
		"OR A,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xb5
	{
		"OR A,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xb6
	{
		"OR A,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xb7
	{
		"OR A,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xb8
	{
		"CP A,B", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xb9
	{
		"CP A,C", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xba
	{
		"CP A,D", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xbb
	{
		"CP A,E", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xbc
	{
		"CP A,H", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xbd
	{
		"CP A,L", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xbe
	{
		"CP A,(HL)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xbf
	{
		"CP A,A", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xc0
	{
		"RET NZ", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xc1
	{
		"POP BC", 0, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xc2
	{
		"JP NZ,u16", 2, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xc3
	{
		"JP u16", 2, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xc4
	{
		"CALL NZ,u16", 2, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xc5
	{
		"PUSH BC", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xc6
	{
		"ADD A,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xc7
	{
		"RST 00h", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xc8
	{
		"RET Z", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xc9
	{
		"RET", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xca
	{
		"JP Z,u16", 2, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xcb
	{
		"PREFIX CB", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xcc
	{
		"CALL Z,u16", 2, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xcd
	{
		"CALL u16", 2, 24,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xce
	{
		"ADC A,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xcf
	{
		"RST 08h", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xd0
	{
		"RET NC", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xd1
	{
		"POP DE", 0, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xd2
	{
		"JP NC,u16", 2, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xd3
	{
		"UNUSED", 0, 0,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xd4
	{
		"CALL NC,u16", 2, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xd5
	{
		"PUSH DE", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xd6
	{
		"SUB A,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xd7
	{
		"RST 10h", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xd8
	{
		"RET C", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xd9
	{
		"RETI", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xda
	{
		"JP C,u16", 2, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xdb
	{
		"UNUSED", 0, 0,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xdc
	{
		"CALL C,u16", 2, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xdd
	{
		"UNUSED", 0, 0,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xde
	{
		"SBC A,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xdf
	{
		"RST 18h", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xe0
	{
		"LD (FF00+u8),A", 1, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xe1
	{
		"POP HL", 0, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xe2
	{
		"LD (FF00+C),A", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xe3
	{
		"UNUSED", 0, 0,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xe4
	{
		"UNUSED", 0, 0,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xe5
	{
		"PUSH HL", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xe6
	{
		"AND A,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xe7
	{
		"RST 20h", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xe8
	{
		"ADD SP,i8", 1, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xe9
	{
		"JP HL", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xea
	{
		"LD (u16),A", 2, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xeb
	{
		"UNUSED", 0, 0,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xec
	{
		"UNUSED", 0, 0,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xed
	{
		"UNUSED", 0, 0,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xee
	{
		"XOR A,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xef
	{
		"RST 28h", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xf0
	{
		"LD A,(FF00+u8)", 1, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xf1
	{
		"POP AF", 0, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xf2
	{
		"LD A,(FF00+C)", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xf3
	{
		"DI", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xf4
	{
		"UNUSED", 0, 0,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xf5
	{
		"PUSH AF", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xf6
	{
		"OR A,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xf7
	{
		"RST 30h", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xf8
	{
		"LD HL,SP+i8", 1, 12,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xf9
	{
		"LD SP,HL", 0, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xfa
	{
		"LD A,(u16)", 2, 16,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xfb
	{
		"EI", 0, 4,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xfc
	{
		"UNUSED", 0, 0,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xfd
	{
		"UNUSED", 0, 0,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xfe
	{
		"CP A,u8", 1, 8,
		func(cpu *CPU, value interface{}) {},
	},

	// 0xff
	{
		"RST 38h", 0, 16,
		func(cpu *CPU, value interface{}) {},
	},
}
