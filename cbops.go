package tamago

// This table contains the extended instructions (prefixed with 0xCB) used by the Game Boy.
var xtable = &Table{
	ins: []Instruction{

		// 0x00
		{
			asm:    "RLC B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x01
		{
			asm:    "RLC C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x02
		{
			asm:    "RLC D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x03
		{
			asm:    "RLC E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x04
		{
			asm:    "RLC H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x05
		{
			asm:    "RLC L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x06
		{
			asm:    "RLC (HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x07
		{
			asm:    "RLC A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x08
		{
			asm:    "RRC B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x09
		{
			asm:    "RRC C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x0A
		{
			asm:    "RRC D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x0B
		{
			asm:    "RRC E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x0C
		{
			asm:    "RRC H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x0D
		{
			asm:    "RRC L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x0E
		{
			asm:    "RRC (HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x0F
		{
			asm:    "RRC A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x10
		{
			asm:    "RL B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x11
		{
			asm:    "RL C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x12
		{
			asm:    "RL D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x13
		{
			asm:    "RL E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x14
		{
			asm:    "RL H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x15
		{
			asm:    "RL L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x16
		{
			asm:    "RL (HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x17
		{
			asm:    "RL A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x18
		{
			asm:    "RR B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x19
		{
			asm:    "RR C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x1A
		{
			asm:    "RR D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x1B
		{
			asm:    "RR E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x1C
		{
			asm:    "RR H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x1D
		{
			asm:    "RR L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x1E
		{
			asm:    "RR (HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x1F
		{
			asm:    "RR A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x20
		{
			asm:    "SLA B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x21
		{
			asm:    "SLA C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x22
		{
			asm:    "SLA D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x23
		{
			asm:    "SLA E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x24
		{
			asm:    "SLA H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x25
		{
			asm:    "SLA L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x26
		{
			asm:    "SLA (HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x27
		{
			asm:    "SLA A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x28
		{
			asm:    "SRA B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x29
		{
			asm:    "SRA C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x2A
		{
			asm:    "SRA D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x2B
		{
			asm:    "SRA E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x2C
		{
			asm:    "SRA H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x2D
		{
			asm:    "SRA L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x2E
		{
			asm:    "SRA (HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x2F
		{
			asm:    "SRA A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x30
		{
			asm:    "SWAP B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x31
		{
			asm:    "SWAP C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x32
		{
			asm:    "SWAP D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x33
		{
			asm:    "SWAP E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x34
		{
			asm:    "SWAP H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x35
		{
			asm:    "SWAP L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x36
		{
			asm:    "SWAP (HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x37
		{
			asm:    "SWAP A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x38
		{
			asm:    "SRL B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x39
		{
			asm:    "SRL C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x3A
		{
			asm:    "SRL D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x3B
		{
			asm:    "SRL E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x3C
		{
			asm:    "SRL H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x3D
		{
			asm:    "SRL L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x3E
		{
			asm:    "SRL (HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x3F
		{
			asm:    "SRL A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x40
		{
			asm:    "BIT 0,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x41
		{
			asm:    "BIT 0,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x42
		{
			asm:    "BIT 0,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x43
		{
			asm:    "BIT 0,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x44
		{
			asm:    "BIT 0,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x45
		{
			asm:    "BIT 0,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x46
		{
			asm:    "BIT 0,(HL)",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0x47
		{
			asm:    "BIT 0,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x48
		{
			asm:    "BIT 1,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x49
		{
			asm:    "BIT 1,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x4A
		{
			asm:    "BIT 1,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x4B
		{
			asm:    "BIT 1,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x4C
		{
			asm:    "BIT 1,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x4D
		{
			asm:    "BIT 1,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x4E
		{
			asm:    "BIT 1,(HL)",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0x4F
		{
			asm:    "BIT 1,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x50
		{
			asm:    "BIT 2,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x51
		{
			asm:    "BIT 2,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x52
		{
			asm:    "BIT 2,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x53
		{
			asm:    "BIT 2,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x54
		{
			asm:    "BIT 2,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x55
		{
			asm:    "BIT 2,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x56
		{
			asm:    "BIT 2,(HL)",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0x57
		{
			asm:    "BIT 2,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x58
		{
			asm:    "BIT 3,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x59
		{
			asm:    "BIT 3,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x5A
		{
			asm:    "BIT 3,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x5B
		{
			asm:    "BIT 3,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x5C
		{
			asm:    "BIT 3,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x5D
		{
			asm:    "BIT 3,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x5E
		{
			asm:    "BIT 3,(HL)",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0x5F
		{
			asm:    "BIT 3,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x60
		{
			asm:    "BIT 4,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x61
		{
			asm:    "BIT 4,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x62
		{
			asm:    "BIT 4,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x63
		{
			asm:    "BIT 4,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x64
		{
			asm:    "BIT 4,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x65
		{
			asm:    "BIT 4,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x66
		{
			asm:    "BIT 4,(HL)",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0x67
		{
			asm:    "BIT 4,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x68
		{
			asm:    "BIT 5,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x69
		{
			asm:    "BIT 5,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x6A
		{
			asm:    "BIT 5,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x6B
		{
			asm:    "BIT 5,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x6C
		{
			asm:    "BIT 5,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x6D
		{
			asm:    "BIT 5,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x6E
		{
			asm:    "BIT 5,(HL)",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0x6F
		{
			asm:    "BIT 5,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x70
		{
			asm:    "BIT 6,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x71
		{
			asm:    "BIT 6,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x72
		{
			asm:    "BIT 6,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x73
		{
			asm:    "BIT 6,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x74
		{
			asm:    "BIT 6,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x75
		{
			asm:    "BIT 6,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x76
		{
			asm:    "BIT 6,(HL)",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0x77
		{
			asm:    "BIT 6,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x78
		{
			asm:    "BIT 7,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x79
		{
			asm:    "BIT 7,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x7A
		{
			asm:    "BIT 7,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x7B
		{
			asm:    "BIT 7,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x7C
		{
			asm:    "BIT 7,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x7D
		{
			asm:    "BIT 7,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x7E
		{
			asm:    "BIT 7,(HL)",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0x7F
		{
			asm:    "BIT 7,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x80
		{
			asm:    "RES 0,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x81
		{
			asm:    "RES 0,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x82
		{
			asm:    "RES 0,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x83
		{
			asm:    "RES 0,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x84
		{
			asm:    "RES 0,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x85
		{
			asm:    "RES 0,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x86
		{
			asm:    "RES 0,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x87
		{
			asm:    "RES 0,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x88
		{
			asm:    "RES 1,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x89
		{
			asm:    "RES 1,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x8A
		{
			asm:    "RES 1,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x8B
		{
			asm:    "RES 1,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x8C
		{
			asm:    "RES 1,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x8D
		{
			asm:    "RES 1,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x8E
		{
			asm:    "RES 1,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x8F
		{
			asm:    "RES 1,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x90
		{
			asm:    "RES 2,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x91
		{
			asm:    "RES 2,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x92
		{
			asm:    "RES 2,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x93
		{
			asm:    "RES 2,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x94
		{
			asm:    "RES 2,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x95
		{
			asm:    "RES 2,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x96
		{
			asm:    "RES 2,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x97
		{
			asm:    "RES 2,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x98
		{
			asm:    "RES 3,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x99
		{
			asm:    "RES 3,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9A
		{
			asm:    "RES 3,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9B
		{
			asm:    "RES 3,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9C
		{
			asm:    "RES 3,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9D
		{
			asm:    "RES 3,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9E
		{
			asm:    "RES 3,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9F
		{
			asm:    "RES 3,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xA0
		{
			asm:    "RES 4,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xA1
		{
			asm:    "RES 4,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xA2
		{
			asm:    "RES 4,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xA3
		{
			asm:    "RES 4,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xA4
		{
			asm:    "RES 4,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xA5
		{
			asm:    "RES 4,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xA6
		{
			asm:    "RES 4,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xA7
		{
			asm:    "RES 4,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xA8
		{
			asm:    "RES 5,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xA9
		{
			asm:    "RES 5,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xAA
		{
			asm:    "RES 5,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xAB
		{
			asm:    "RES 5,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xAC
		{
			asm:    "RES 5,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xAD
		{
			asm:    "RES 5,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xAE
		{
			asm:    "RES 5,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xAF
		{
			asm:    "RES 5,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xB0
		{
			asm:    "RES 6,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xB1
		{
			asm:    "RES 6,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xB2
		{
			asm:    "RES 6,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xB3
		{
			asm:    "RES 6,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xB4
		{
			asm:    "RES 6,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xB5
		{
			asm:    "RES 6,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xB6
		{
			asm:    "RES 6,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xB7
		{
			asm:    "RES 6,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xB8
		{
			asm:    "RES 7,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xB9
		{
			asm:    "RES 7,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xBA
		{
			asm:    "RES 7,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xBB
		{
			asm:    "RES 7,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xBC
		{
			asm:    "RES 7,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xBD
		{
			asm:    "RES 7,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xBE
		{
			asm:    "RES 7,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xBF
		{
			asm:    "RES 7,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xC0
		{
			asm:    "SET 0,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xC1
		{
			asm:    "SET 0,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xC2
		{
			asm:    "SET 0,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xC3
		{
			asm:    "SET 0,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xC4
		{
			asm:    "SET 0,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xC5
		{
			asm:    "SET 0,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xC6
		{
			asm:    "SET 0,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xC7
		{
			asm:    "SET 0,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xC8
		{
			asm:    "SET 1,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xC9
		{
			asm:    "SET 1,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xCA
		{
			asm:    "SET 1,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xCB
		{
			asm:    "SET 1,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xCC
		{
			asm:    "SET 1,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xCD
		{
			asm:    "SET 1,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xCE
		{
			asm:    "SET 1,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xCF
		{
			asm:    "SET 1,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xD0
		{
			asm:    "SET 2,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xD1
		{
			asm:    "SET 2,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xD2
		{
			asm:    "SET 2,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xD3
		{
			asm:    "SET 2,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xD4
		{
			asm:    "SET 2,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xD5
		{
			asm:    "SET 2,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xD6
		{
			asm:    "SET 2,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xD7
		{
			asm:    "SET 2,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xD8
		{
			asm:    "SET 3,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xD9
		{
			asm:    "SET 3,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xDA
		{
			asm:    "SET 3,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xDB
		{
			asm:    "SET 3,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xDC
		{
			asm:    "SET 3,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xDD
		{
			asm:    "SET 3,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xDE
		{
			asm:    "SET 3,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xDF
		{
			asm:    "SET 3,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xE0
		{
			asm:    "SET 4,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xE1
		{
			asm:    "SET 4,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xE2
		{
			asm:    "SET 4,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xE3
		{
			asm:    "SET 4,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xE4
		{
			asm:    "SET 4,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xE5
		{
			asm:    "SET 4,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xE6
		{
			asm:    "SET 4,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xE7
		{
			asm:    "SET 4,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xE8
		{
			asm:    "SET 5,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xE9
		{
			asm:    "SET 5,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xEA
		{
			asm:    "SET 5,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xEB
		{
			asm:    "SET 5,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xEC
		{
			asm:    "SET 5,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xED
		{
			asm:    "SET 5,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xEE
		{
			asm:    "SET 5,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xEF
		{
			asm:    "SET 5,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xF0
		{
			asm:    "SET 6,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xF1
		{
			asm:    "SET 6,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xF2
		{
			asm:    "SET 6,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xF3
		{
			asm:    "SET 6,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xF4
		{
			asm:    "SET 6,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xF5
		{
			asm:    "SET 6,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xF6
		{
			asm:    "SET 6,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xF7
		{
			asm:    "SET 6,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xF8
		{
			asm:    "SET 7,B",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xF9
		{
			asm:    "SET 7,C",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xFA
		{
			asm:    "SET 7,D",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xFB
		{
			asm:    "SET 7,E",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xFC
		{
			asm:    "SET 7,H",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xFD
		{
			asm:    "SET 7,L",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xFE
		{
			asm:    "SET 7,(HL)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xFF
		{
			asm:    "SET 7,A",
			length: 2,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},
	},
}
