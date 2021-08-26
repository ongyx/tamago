package tamago

// This table contains the extended instructions (prefixed with 0xcb) used by the Game Boy.
var cbops = [256]Instruction{

	// 0x00
	{
		asm:    "RLC B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rlc(&s.BC.Hi)
		},
	},

	// 0x01
	{
		asm:    "RLC C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rlc(&s.BC.Lo)
		},
	},

	// 0x02
	{
		asm:    "RLC D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rlc(&s.DE.Hi)
		},
	},

	// 0x03
	{
		asm:    "RLC E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rlc(&s.DE.Lo)
		},
	},

	// 0x04
	{
		asm:    "RLC H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rlc(&s.HL.Hi)
		},
	},

	// 0x05
	{
		asm:    "RLC L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rlc(&s.HL.Lo)
		},
	},

	// 0x06
	{
		asm:    "RLC (HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.rlc(&b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x07
	{
		asm:    "RLC A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rlc(&s.AF.Hi)
		},
	},

	// 0x08
	{
		asm:    "RRC B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rrc(&s.BC.Hi)
		},
	},

	// 0x09
	{
		asm:    "RRC C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rrc(&s.BC.Lo)
		},
	},

	// 0x0a
	{
		asm:    "RRC D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rrc(&s.DE.Hi)
		},
	},

	// 0x0b
	{
		asm:    "RRC E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rrc(&s.DE.Lo)
		},
	},

	// 0x0c
	{
		asm:    "RRC H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rrc(&s.HL.Hi)
		},
	},

	// 0x0d
	{
		asm:    "RRC L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rrc(&s.HL.Lo)
		},
	},

	// 0x0e
	{
		asm:    "RRC (HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.rrc(&b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x0f
	{
		asm:    "RRC A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rrc(&s.AF.Hi)
		},
	},

	// 0x10
	{
		asm:    "RL B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rl(&s.BC.Hi)
		},
	},

	// 0x11
	{
		asm:    "RL C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rl(&s.BC.Lo)
		},
	},

	// 0x12
	{
		asm:    "RL D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rl(&s.DE.Hi)
		},
	},

	// 0x13
	{
		asm:    "RL E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rl(&s.DE.Lo)
		},
	},

	// 0x14
	{
		asm:    "RL H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rl(&s.HL.Hi)
		},
	},

	// 0x15
	{
		asm:    "RL L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rl(&s.HL.Lo)
		},
	},

	// 0x16
	{
		asm:    "RL (HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.rl(&b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x17
	{
		asm:    "RL A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rl(&s.AF.Hi)
		},
	},

	// 0x18
	{
		asm:    "RR B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rr(&s.BC.Hi)
		},
	},

	// 0x19
	{
		asm:    "RR C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rr(&s.BC.Lo)
		},
	},

	// 0x1a
	{
		asm:    "RR D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rr(&s.DE.Hi)
		},
	},

	// 0x1b
	{
		asm:    "RR E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rr(&s.DE.Lo)
		},
	},

	// 0x1c
	{
		asm:    "RR H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rr(&s.HL.Hi)
		},
	},

	// 0x1d
	{
		asm:    "RR L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rr(&s.HL.Lo)
		},
	},

	// 0x1e
	{
		asm:    "RR (HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.rr(&b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x1f
	{
		asm:    "RR A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.rr(&s.AF.Hi)
		},
	},

	// 0x20
	{
		asm:    "SLA B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sla(&s.BC.Hi)
		},
	},

	// 0x21
	{
		asm:    "SLA C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sla(&s.BC.Lo)
		},
	},

	// 0x22
	{
		asm:    "SLA D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sla(&s.DE.Hi)
		},
	},

	// 0x23
	{
		asm:    "SLA E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sla(&s.DE.Lo)
		},
	},

	// 0x24
	{
		asm:    "SLA H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sla(&s.HL.Hi)
		},
	},

	// 0x25
	{
		asm:    "SLA L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sla(&s.HL.Lo)
		},
	},

	// 0x26
	{
		asm:    "SLA (HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.sla(&b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x27
	{
		asm:    "SLA A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sla(&s.AF.Hi)
		},
	},

	// 0x28
	{
		asm:    "SRA B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sra(&s.BC.Hi)
		},
	},

	// 0x29
	{
		asm:    "SRA C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sra(&s.BC.Lo)
		},
	},

	// 0x2a
	{
		asm:    "SRA D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sra(&s.DE.Hi)
		},
	},

	// 0x2b
	{
		asm:    "SRA E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sra(&s.DE.Lo)
		},
	},

	// 0x2c
	{
		asm:    "SRA H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sra(&s.HL.Hi)
		},
	},

	// 0x2d
	{
		asm:    "SRA L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sra(&s.HL.Lo)
		},
	},

	// 0x2e
	{
		asm:    "SRA (HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.sra(&b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x2f
	{
		asm:    "SRA A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.sra(&s.AF.Hi)
		},
	},

	// 0x30
	{
		asm:    "SWAP B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.swap(&s.BC.Hi)
		},
	},

	// 0x31
	{
		asm:    "SWAP C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.swap(&s.BC.Lo)
		},
	},

	// 0x32
	{
		asm:    "SWAP D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.swap(&s.DE.Hi)
		},
	},

	// 0x33
	{
		asm:    "SWAP E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.swap(&s.DE.Lo)
		},
	},

	// 0x34
	{
		asm:    "SWAP H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.swap(&s.HL.Hi)
		},
	},

	// 0x35
	{
		asm:    "SWAP L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.swap(&s.HL.Lo)
		},
	},

	// 0x36
	{
		asm:    "SWAP (HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.swap(&b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x37
	{
		asm:    "SWAP A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.swap(&s.AF.Hi)
		},
	},

	// 0x38
	{
		asm:    "SRL B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.srl(&s.BC.Hi)
		},
	},

	// 0x39
	{
		asm:    "SRL C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.srl(&s.BC.Lo)
		},
	},

	// 0x3a
	{
		asm:    "SRL D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.srl(&s.DE.Hi)
		},
	},

	// 0x3b
	{
		asm:    "SRL E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.srl(&s.DE.Lo)
		},
	},

	// 0x3c
	{
		asm:    "SRL H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.srl(&s.HL.Hi)
		},
	},

	// 0x3d
	{
		asm:    "SRL L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.srl(&s.HL.Lo)
		},
	},

	// 0x3e
	{
		asm:    "SRL (HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.srl(&b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x3f
	{
		asm:    "SRL A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.srl(&s.AF.Hi)
		},
	},

	// 0x40
	{
		asm:    "BIT 0,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(0, &s.BC.Hi)
		},
	},

	// 0x41
	{
		asm:    "BIT 0,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(0, &s.BC.Lo)
		},
	},

	// 0x42
	{
		asm:    "BIT 0,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(0, &s.DE.Hi)
		},
	},

	// 0x43
	{
		asm:    "BIT 0,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(0, &s.DE.Lo)
		},
	},

	// 0x44
	{
		asm:    "BIT 0,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(0, &s.HL.Hi)
		},
	},

	// 0x45
	{
		asm:    "BIT 0,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(0, &s.HL.Lo)
		},
	},

	// 0x46
	{
		asm:    "BIT 0,(HL)",
		length: 0,
		cycles: 3,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.bit(0, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x47
	{
		asm:    "BIT 0,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(0, &s.AF.Hi)
		},
	},

	// 0x48
	{
		asm:    "BIT 1,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(1, &s.BC.Hi)
		},
	},

	// 0x49
	{
		asm:    "BIT 1,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(1, &s.BC.Lo)
		},
	},

	// 0x4a
	{
		asm:    "BIT 1,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(1, &s.DE.Hi)
		},
	},

	// 0x4b
	{
		asm:    "BIT 1,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(1, &s.DE.Lo)
		},
	},

	// 0x4c
	{
		asm:    "BIT 1,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(1, &s.HL.Hi)
		},
	},

	// 0x4d
	{
		asm:    "BIT 1,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(1, &s.HL.Lo)
		},
	},

	// 0x4e
	{
		asm:    "BIT 1,(HL)",
		length: 0,
		cycles: 3,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.bit(1, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x4f
	{
		asm:    "BIT 1,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(1, &s.AF.Hi)
		},
	},

	// 0x50
	{
		asm:    "BIT 2,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(2, &s.BC.Hi)
		},
	},

	// 0x51
	{
		asm:    "BIT 2,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(2, &s.BC.Lo)
		},
	},

	// 0x52
	{
		asm:    "BIT 2,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(2, &s.DE.Hi)
		},
	},

	// 0x53
	{
		asm:    "BIT 2,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(2, &s.DE.Lo)
		},
	},

	// 0x54
	{
		asm:    "BIT 2,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(2, &s.HL.Hi)
		},
	},

	// 0x55
	{
		asm:    "BIT 2,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(2, &s.HL.Lo)
		},
	},

	// 0x56
	{
		asm:    "BIT 2,(HL)",
		length: 0,
		cycles: 3,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.bit(2, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x57
	{
		asm:    "BIT 2,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(2, &s.AF.Hi)
		},
	},

	// 0x58
	{
		asm:    "BIT 3,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(3, &s.BC.Hi)
		},
	},

	// 0x59
	{
		asm:    "BIT 3,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(3, &s.BC.Lo)
		},
	},

	// 0x5a
	{
		asm:    "BIT 3,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(3, &s.DE.Hi)
		},
	},

	// 0x5b
	{
		asm:    "BIT 3,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(3, &s.DE.Lo)
		},
	},

	// 0x5c
	{
		asm:    "BIT 3,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(3, &s.HL.Hi)
		},
	},

	// 0x5d
	{
		asm:    "BIT 3,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(3, &s.HL.Lo)
		},
	},

	// 0x5e
	{
		asm:    "BIT 3,(HL)",
		length: 0,
		cycles: 3,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.bit(3, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x5f
	{
		asm:    "BIT 3,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(3, &s.AF.Hi)
		},
	},

	// 0x60
	{
		asm:    "BIT 4,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(4, &s.BC.Hi)
		},
	},

	// 0x61
	{
		asm:    "BIT 4,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(4, &s.BC.Lo)
		},
	},

	// 0x62
	{
		asm:    "BIT 4,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(4, &s.DE.Hi)
		},
	},

	// 0x63
	{
		asm:    "BIT 4,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(4, &s.DE.Lo)
		},
	},

	// 0x64
	{
		asm:    "BIT 4,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(4, &s.HL.Hi)
		},
	},

	// 0x65
	{
		asm:    "BIT 4,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(4, &s.HL.Lo)
		},
	},

	// 0x66
	{
		asm:    "BIT 4,(HL)",
		length: 0,
		cycles: 3,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.bit(4, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x67
	{
		asm:    "BIT 4,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(4, &s.AF.Hi)
		},
	},

	// 0x68
	{
		asm:    "BIT 5,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(5, &s.BC.Hi)
		},
	},

	// 0x69
	{
		asm:    "BIT 5,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(5, &s.BC.Lo)
		},
	},

	// 0x6a
	{
		asm:    "BIT 5,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(5, &s.DE.Hi)
		},
	},

	// 0x6b
	{
		asm:    "BIT 5,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(5, &s.DE.Lo)
		},
	},

	// 0x6c
	{
		asm:    "BIT 5,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(5, &s.HL.Hi)
		},
	},

	// 0x6d
	{
		asm:    "BIT 5,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(5, &s.HL.Lo)
		},
	},

	// 0x6e
	{
		asm:    "BIT 5,(HL)",
		length: 0,
		cycles: 3,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.bit(5, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x6f
	{
		asm:    "BIT 5,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(5, &s.AF.Hi)
		},
	},

	// 0x70
	{
		asm:    "BIT 6,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(6, &s.BC.Hi)
		},
	},

	// 0x71
	{
		asm:    "BIT 6,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(6, &s.BC.Lo)
		},
	},

	// 0x72
	{
		asm:    "BIT 6,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(6, &s.DE.Hi)
		},
	},

	// 0x73
	{
		asm:    "BIT 6,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(6, &s.DE.Lo)
		},
	},

	// 0x74
	{
		asm:    "BIT 6,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(6, &s.HL.Hi)
		},
	},

	// 0x75
	{
		asm:    "BIT 6,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(6, &s.HL.Lo)
		},
	},

	// 0x76
	{
		asm:    "BIT 6,(HL)",
		length: 0,
		cycles: 3,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.bit(6, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x77
	{
		asm:    "BIT 6,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(6, &s.AF.Hi)
		},
	},

	// 0x78
	{
		asm:    "BIT 7,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(7, &s.BC.Hi)
		},
	},

	// 0x79
	{
		asm:    "BIT 7,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(7, &s.BC.Lo)
		},
	},

	// 0x7a
	{
		asm:    "BIT 7,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(7, &s.DE.Hi)
		},
	},

	// 0x7b
	{
		asm:    "BIT 7,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(7, &s.DE.Lo)
		},
	},

	// 0x7c
	{
		asm:    "BIT 7,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(7, &s.HL.Hi)
		},
	},

	// 0x7d
	{
		asm:    "BIT 7,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(7, &s.HL.Lo)
		},
	},

	// 0x7e
	{
		asm:    "BIT 7,(HL)",
		length: 0,
		cycles: 3,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.bit(7, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x7f
	{
		asm:    "BIT 7,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.bit(7, &s.AF.Hi)
		},
	},

	// 0x80
	{
		asm:    "RES 0,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(0, &s.BC.Hi)
		},
	},

	// 0x81
	{
		asm:    "RES 0,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(0, &s.BC.Lo)
		},
	},

	// 0x82
	{
		asm:    "RES 0,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(0, &s.DE.Hi)
		},
	},

	// 0x83
	{
		asm:    "RES 0,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(0, &s.DE.Lo)
		},
	},

	// 0x84
	{
		asm:    "RES 0,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(0, &s.HL.Hi)
		},
	},

	// 0x85
	{
		asm:    "RES 0,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(0, &s.HL.Lo)
		},
	},

	// 0x86
	{
		asm:    "RES 0,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.res(0, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x87
	{
		asm:    "RES 0,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(0, &s.AF.Hi)
		},
	},

	// 0x88
	{
		asm:    "RES 1,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(1, &s.BC.Hi)
		},
	},

	// 0x89
	{
		asm:    "RES 1,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(1, &s.BC.Lo)
		},
	},

	// 0x8a
	{
		asm:    "RES 1,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(1, &s.DE.Hi)
		},
	},

	// 0x8b
	{
		asm:    "RES 1,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(1, &s.DE.Lo)
		},
	},

	// 0x8c
	{
		asm:    "RES 1,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(1, &s.HL.Hi)
		},
	},

	// 0x8d
	{
		asm:    "RES 1,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(1, &s.HL.Lo)
		},
	},

	// 0x8e
	{
		asm:    "RES 1,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.res(1, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x8f
	{
		asm:    "RES 1,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(1, &s.AF.Hi)
		},
	},

	// 0x90
	{
		asm:    "RES 2,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(2, &s.BC.Hi)
		},
	},

	// 0x91
	{
		asm:    "RES 2,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(2, &s.BC.Lo)
		},
	},

	// 0x92
	{
		asm:    "RES 2,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(2, &s.DE.Hi)
		},
	},

	// 0x93
	{
		asm:    "RES 2,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(2, &s.DE.Lo)
		},
	},

	// 0x94
	{
		asm:    "RES 2,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(2, &s.HL.Hi)
		},
	},

	// 0x95
	{
		asm:    "RES 2,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(2, &s.HL.Lo)
		},
	},

	// 0x96
	{
		asm:    "RES 2,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.res(2, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x97
	{
		asm:    "RES 2,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(2, &s.AF.Hi)
		},
	},

	// 0x98
	{
		asm:    "RES 3,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(3, &s.BC.Hi)
		},
	},

	// 0x99
	{
		asm:    "RES 3,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(3, &s.BC.Lo)
		},
	},

	// 0x9a
	{
		asm:    "RES 3,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(3, &s.DE.Hi)
		},
	},

	// 0x9b
	{
		asm:    "RES 3,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(3, &s.DE.Lo)
		},
	},

	// 0x9c
	{
		asm:    "RES 3,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(3, &s.HL.Hi)
		},
	},

	// 0x9d
	{
		asm:    "RES 3,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(3, &s.HL.Lo)
		},
	},

	// 0x9e
	{
		asm:    "RES 3,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.res(3, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0x9f
	{
		asm:    "RES 3,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(3, &s.AF.Hi)
		},
	},

	// 0xa0
	{
		asm:    "RES 4,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(4, &s.BC.Hi)
		},
	},

	// 0xa1
	{
		asm:    "RES 4,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(4, &s.BC.Lo)
		},
	},

	// 0xa2
	{
		asm:    "RES 4,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(4, &s.DE.Hi)
		},
	},

	// 0xa3
	{
		asm:    "RES 4,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(4, &s.DE.Lo)
		},
	},

	// 0xa4
	{
		asm:    "RES 4,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(4, &s.HL.Hi)
		},
	},

	// 0xa5
	{
		asm:    "RES 4,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(4, &s.HL.Lo)
		},
	},

	// 0xa6
	{
		asm:    "RES 4,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.res(4, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xa7
	{
		asm:    "RES 4,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(4, &s.AF.Hi)
		},
	},

	// 0xa8
	{
		asm:    "RES 5,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(5, &s.BC.Hi)
		},
	},

	// 0xa9
	{
		asm:    "RES 5,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(5, &s.BC.Lo)
		},
	},

	// 0xaa
	{
		asm:    "RES 5,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(5, &s.DE.Hi)
		},
	},

	// 0xab
	{
		asm:    "RES 5,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(5, &s.DE.Lo)
		},
	},

	// 0xac
	{
		asm:    "RES 5,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(5, &s.HL.Hi)
		},
	},

	// 0xad
	{
		asm:    "RES 5,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(5, &s.HL.Lo)
		},
	},

	// 0xae
	{
		asm:    "RES 5,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.res(5, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xaf
	{
		asm:    "RES 5,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(5, &s.AF.Hi)
		},
	},

	// 0xb0
	{
		asm:    "RES 6,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(6, &s.BC.Hi)
		},
	},

	// 0xb1
	{
		asm:    "RES 6,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(6, &s.BC.Lo)
		},
	},

	// 0xb2
	{
		asm:    "RES 6,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(6, &s.DE.Hi)
		},
	},

	// 0xb3
	{
		asm:    "RES 6,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(6, &s.DE.Lo)
		},
	},

	// 0xb4
	{
		asm:    "RES 6,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(6, &s.HL.Hi)
		},
	},

	// 0xb5
	{
		asm:    "RES 6,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(6, &s.HL.Lo)
		},
	},

	// 0xb6
	{
		asm:    "RES 6,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.res(6, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xb7
	{
		asm:    "RES 6,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(6, &s.AF.Hi)
		},
	},

	// 0xb8
	{
		asm:    "RES 7,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(7, &s.BC.Hi)
		},
	},

	// 0xb9
	{
		asm:    "RES 7,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(7, &s.BC.Lo)
		},
	},

	// 0xba
	{
		asm:    "RES 7,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(7, &s.DE.Hi)
		},
	},

	// 0xbb
	{
		asm:    "RES 7,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(7, &s.DE.Lo)
		},
	},

	// 0xbc
	{
		asm:    "RES 7,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(7, &s.HL.Hi)
		},
	},

	// 0xbd
	{
		asm:    "RES 7,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(7, &s.HL.Lo)
		},
	},

	// 0xbe
	{
		asm:    "RES 7,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.res(7, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xbf
	{
		asm:    "RES 7,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.res(7, &s.AF.Hi)
		},
	},

	// 0xc0
	{
		asm:    "SET 0,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(0, &s.BC.Hi)
		},
	},

	// 0xc1
	{
		asm:    "SET 0,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(0, &s.BC.Lo)
		},
	},

	// 0xc2
	{
		asm:    "SET 0,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(0, &s.DE.Hi)
		},
	},

	// 0xc3
	{
		asm:    "SET 0,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(0, &s.DE.Lo)
		},
	},

	// 0xc4
	{
		asm:    "SET 0,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(0, &s.HL.Hi)
		},
	},

	// 0xc5
	{
		asm:    "SET 0,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(0, &s.HL.Lo)
		},
	},

	// 0xc6
	{
		asm:    "SET 0,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.set(0, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xc7
	{
		asm:    "SET 0,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(0, &s.AF.Hi)
		},
	},

	// 0xc8
	{
		asm:    "SET 1,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(1, &s.BC.Hi)
		},
	},

	// 0xc9
	{
		asm:    "SET 1,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(1, &s.BC.Lo)
		},
	},

	// 0xca
	{
		asm:    "SET 1,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(1, &s.DE.Hi)
		},
	},

	// 0xcb
	{
		asm:    "SET 1,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(1, &s.DE.Lo)
		},
	},

	// 0xcc
	{
		asm:    "SET 1,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(1, &s.HL.Hi)
		},
	},

	// 0xcd
	{
		asm:    "SET 1,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(1, &s.HL.Lo)
		},
	},

	// 0xce
	{
		asm:    "SET 1,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.set(1, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xcf
	{
		asm:    "SET 1,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(1, &s.AF.Hi)
		},
	},

	// 0xd0
	{
		asm:    "SET 2,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(2, &s.BC.Hi)
		},
	},

	// 0xd1
	{
		asm:    "SET 2,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(2, &s.BC.Lo)
		},
	},

	// 0xd2
	{
		asm:    "SET 2,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(2, &s.DE.Hi)
		},
	},

	// 0xd3
	{
		asm:    "SET 2,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(2, &s.DE.Lo)
		},
	},

	// 0xd4
	{
		asm:    "SET 2,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(2, &s.HL.Hi)
		},
	},

	// 0xd5
	{
		asm:    "SET 2,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(2, &s.HL.Lo)
		},
	},

	// 0xd6
	{
		asm:    "SET 2,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.set(2, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xd7
	{
		asm:    "SET 2,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(2, &s.AF.Hi)
		},
	},

	// 0xd8
	{
		asm:    "SET 3,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(3, &s.BC.Hi)
		},
	},

	// 0xd9
	{
		asm:    "SET 3,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(3, &s.BC.Lo)
		},
	},

	// 0xda
	{
		asm:    "SET 3,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(3, &s.DE.Hi)
		},
	},

	// 0xdb
	{
		asm:    "SET 3,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(3, &s.DE.Lo)
		},
	},

	// 0xdc
	{
		asm:    "SET 3,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(3, &s.HL.Hi)
		},
	},

	// 0xdd
	{
		asm:    "SET 3,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(3, &s.HL.Lo)
		},
	},

	// 0xde
	{
		asm:    "SET 3,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.set(3, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xdf
	{
		asm:    "SET 3,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(3, &s.AF.Hi)
		},
	},

	// 0xe0
	{
		asm:    "SET 4,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(4, &s.BC.Hi)
		},
	},

	// 0xe1
	{
		asm:    "SET 4,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(4, &s.BC.Lo)
		},
	},

	// 0xe2
	{
		asm:    "SET 4,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(4, &s.DE.Hi)
		},
	},

	// 0xe3
	{
		asm:    "SET 4,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(4, &s.DE.Lo)
		},
	},

	// 0xe4
	{
		asm:    "SET 4,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(4, &s.HL.Hi)
		},
	},

	// 0xe5
	{
		asm:    "SET 4,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(4, &s.HL.Lo)
		},
	},

	// 0xe6
	{
		asm:    "SET 4,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.set(4, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xe7
	{
		asm:    "SET 4,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(4, &s.AF.Hi)
		},
	},

	// 0xe8
	{
		asm:    "SET 5,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(5, &s.BC.Hi)
		},
	},

	// 0xe9
	{
		asm:    "SET 5,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(5, &s.BC.Lo)
		},
	},

	// 0xea
	{
		asm:    "SET 5,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(5, &s.DE.Hi)
		},
	},

	// 0xeb
	{
		asm:    "SET 5,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(5, &s.DE.Lo)
		},
	},

	// 0xec
	{
		asm:    "SET 5,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(5, &s.HL.Hi)
		},
	},

	// 0xed
	{
		asm:    "SET 5,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(5, &s.HL.Lo)
		},
	},

	// 0xee
	{
		asm:    "SET 5,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.set(5, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xef
	{
		asm:    "SET 5,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(5, &s.AF.Hi)
		},
	},

	// 0xf0
	{
		asm:    "SET 6,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(6, &s.BC.Hi)
		},
	},

	// 0xf1
	{
		asm:    "SET 6,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(6, &s.BC.Lo)
		},
	},

	// 0xf2
	{
		asm:    "SET 6,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(6, &s.DE.Hi)
		},
	},

	// 0xf3
	{
		asm:    "SET 6,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(6, &s.DE.Lo)
		},
	},

	// 0xf4
	{
		asm:    "SET 6,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(6, &s.HL.Hi)
		},
	},

	// 0xf5
	{
		asm:    "SET 6,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(6, &s.HL.Lo)
		},
	},

	// 0xf6
	{
		asm:    "SET 6,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.set(6, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xf7
	{
		asm:    "SET 6,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(6, &s.AF.Hi)
		},
	},

	// 0xf8
	{
		asm:    "SET 7,B",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(7, &s.BC.Hi)
		},
	},

	// 0xf9
	{
		asm:    "SET 7,C",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(7, &s.BC.Lo)
		},
	},

	// 0xfa
	{
		asm:    "SET 7,D",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(7, &s.DE.Hi)
		},
	},

	// 0xfb
	{
		asm:    "SET 7,E",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(7, &s.DE.Lo)
		},
	},

	// 0xfc
	{
		asm:    "SET 7,H",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(7, &s.HL.Hi)
		},
	},

	// 0xfd
	{
		asm:    "SET 7,L",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(7, &s.HL.Lo)
		},
	},

	// 0xfe
	{
		asm:    "SET 7,(HL)",
		length: 0,
		cycles: 4,

		fn: func(s *State, v Value) {
			b := s.ReadFrom(s.HL)
			s.fl.set(7, &b)
			s.WriteTo(s.HL, b)
		},
	},

	// 0xff
	{
		asm:    "SET 7,A",
		length: 0,
		cycles: 2,

		fn: func(s *State, v Value) {
			s.fl.set(7, &s.AF.Hi)
		},
	},
}
