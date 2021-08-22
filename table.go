package tamago

type Table struct {
	ins []Instruction
}

func nop(s *State, v Value) {}

// This table contains all 512 opcodes used in the Game Boy
// (256 normal, 256 prefixed by 0xCB).

var table = &Table{
	ins: []Instruction{

		// 0x00
		{
			asm:    "NOP",
			length: 0,
			cycles: 1,

			fn: nop,
		},

		// 0x01
		{
			asm:    "LD BC,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.BC.Set(v.U16())
			},
		},

		// 0x02
		{
			asm:    "LD (BC),A",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.WriteTo(s.BC, s.AF.Hi)
			},
		},

		// 0x03
		{
			asm:    "INC BC",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.BC.Inc()
			},
		},

		// 0x04
		{
			asm:    "INC B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.inc(&s.BC.Hi)
			},
		},

		// 0x05
		{
			asm:    "DEC B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.dec(&s.BC.Hi)
			},
		},

		// 0x06
		{
			asm:    "LD B,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.BC.Hi = v.U8()
			},
		},

		// 0x07
		{
			asm:    "RLCA",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				// Take only the first bit and shift all the way to the right
				carry := (s.AF.Hi & 0x80) >> 7
				s.fl.setIfCarry(carry)

				// Shift A left 1 bit and add back the carry.
				s.AF.Hi <<= 1
				s.AF.Hi += carry

				// Clear all other flags.
				s.fl.ClearAllExcept(Carry)
			},
		},

		// 0x08
		{
			asm:    "LD (u16),SP",
			length: 2,
			cycles: 5,

			fn: func(s *State, v Value) {
				s.WriteShort(s.SP, v.U16())
			},
		},

		// 0x09
		{
			asm:    "ADD HL,BC",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Add(s.BC.Get(), s.fl)
			},
		},

		// 0x0a
		{
			asm:    "LD A,(BC)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.ReadFrom(s.BC)
			},
		},

		// 0x0b
		{
			asm:    "DEC BC",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.BC.Dec()
			},
		},

		// 0x0c
		{
			asm:    "INC C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.inc(&s.BC.Lo)
			},
		},

		// 0x0d
		{
			asm:    "DEC C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.dec(&s.BC.Lo)
			},
		},

		// 0x0e
		{
			asm:    "LD C,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.BC.Lo = v.U8()
			},
		},

		// 0x0f
		{
			asm:    "RRCA",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				// Take only the last bit
				carry := s.AF.Hi & 0x01
				s.fl.setIfCarry(carry)

				// Shift A right 1 bit and put back the bits.
				s.AF.Hi >>= 1
				if carry != 0 {
					s.AF.Hi |= 0x80
				}

				// Clear all other flags.
				s.fl.ClearAllExcept(Carry)
			},
		},

		// 0x10
		{
			asm:    "STOP",
			length: 1,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.stopped = true
			},
		},

		// 0x11
		{
			asm:    "LD DE,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.DE.Set(v.U16())
			},
		},

		// 0x12
		{
			asm:    "LD (DE),A",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.WriteTo(s.DE, s.AF.Hi)
			},
		},

		// 0x13
		{
			asm:    "INC DE",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.DE.Inc()
			},
		},

		// 0x14
		{
			asm:    "INC D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.inc(&s.DE.Hi)
			},
		},

		// 0x15
		{
			asm:    "DEC D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.dec(&s.DE.Hi)
			},
		},

		// 0x16
		{
			asm:    "LD D,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.DE.Hi = v.U8()
			},
		},

		// 0x17
		{
			asm:    "RLA",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				carry := 0

				if s.fl.Has(Carry) {
					carry = 1
				}

				s.fl.setIfCarry(s.AF.Hi & 0x80)

				s.AF.Hi <<= 1
				s.AF.Hi += carry

				s.fl.ClearAllExcept(Carry)
			},
		},

		// 0x18
		{
			asm:    "JR i8",
			length: 1,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.PC += v.S8()
			},
		},

		// 0x19
		{
			asm:    "ADD HL,DE",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Add(s.DE.Get(), s.fl)
			},
		},

		// 0x1a
		{
			asm:    "LD A,(DE)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.ReadFrom(s.DE)
			},
		},

		// 0x1b
		{
			asm:    "DEC DE",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.DE.Dec()
			},
		},

		// 0x1c
		{
			asm:    "INC E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.inc(&s.DE.Lo)
			},
		},

		// 0x1d
		{
			asm:    "DEC E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.dec(&s.DE.Lo)
			},
		},

		// 0x1e
		{
			asm:    "LD E,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.DE.Lo = v.U8()
			},
		},

		// 0x1f
		{
			asm:    "RRA",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				carryBit := 0
				if s.fl.Has(Carry) {
					carryBit = 1
				}

				carry := carryBit << 7

				s.fl.setIfCarry(s.AF.Hi & 0x01)

				s.AF.Hi >>= 1
				s.AF.Hi += carry

				s.fl.ClearAllExcept(Carry)
			},
		},

		// 0x20
		{
			asm:    "JR NZ,i8",
			length: 1,
			cycles: 0,

			fn: func(s *State, v Value) {
				s.JumpIf(!s.fl.Has(Zero), v)
			},
		},

		// 0x21
		{
			asm:    "LD HL,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.HL.Set(v.U16())
			},
		},

		// 0x22
		{
			asm:    "LD (HL+),A",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.WriteTo(s.HL, s.AF.Hi)
				s.HL.Inc()
			},
		},

		// 0x23
		{
			asm:    "INC HL",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Inc()
			},
		},

		// 0x24
		{
			asm:    "INC H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.inc(&s.HL.Hi)
			},
		},

		// 0x25
		{
			asm:    "DEC H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.dec(&s.BC.Hi)
			},
		},

		// 0x26
		{
			asm:    "LD H,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Hi = v.U8()
			},
		},

		// 0x27
		{
			asm:    "DAA",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				a := uint16(s.AF.Hi)

				if s.fl.Has(Negative) {

					if s.fl.Has(HalfCarry) {
						a = (s - 0x06) & 0xff
					}

					if s.fl.Has(Carry) {
						a -= 0x60
					}

				} else {

					if s.fl.Has(HalfCarry) || (a&0x0f) > 9 {
						a += 0x06
					}

					if s.fl.Has(Carry) || s > 0x9f {
						a += 0x60
					}

				}

				s.AF.Hi = uint8(a & 0xff)

				s.fl.Clear(HalfCarry)
				s.fl.setIfZero(a)
				s.fl.setIfCarry(a &^ 0xff)
			},
		},

		// 0x28
		{
			asm:    "JR Z,i8",
			length: 1,
			cycles: 0,

			fn: func(s *State, v Value) {
				s.JumpIf(s.fl.Has(Zero), v)
			},
		},

		// 0x29
		{
			asm:    "ADD HL,HL",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Add(s.HL.Get(), s.fl)
			},
		},

		// 0x2a
		{
			asm:    "LD A,(HL+)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.ReadFrom(s.HL)
				s.HL.Inc()
			},
		},

		// 0x2b
		{
			asm:    "DEC HL",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Dec()
			},
		},

		// 0x2c
		{
			asm:    "INC L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.inc(&s.HL.Lo)
			},
		},

		// 0x2d
		{
			asm:    "DEC L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.dec(&s.HL.Lo)
			},
		},

		// 0x2e
		{
			asm:    "LD L,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Lo = v.U8()
			},
		},

		// 0x2f
		{
			asm:    "CPL",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.AF.Hi = ^s.AF.Hi
				s.fl.Set(Negative | HalfCarry)
			},
		},

		// 0x30
		{
			asm:    "JR NC,i8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.JumpIf(!s.fl.Has(Carry), v)
			},
		},

		// 0x31
		{
			asm:    "LD SP,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.SP = v.U16()
			},
		},

		// 0x32
		{
			asm:    "LD (HL-),A",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.WriteTo(s.HL, s.AF.Hi)
				s.HL.Dec()
			},
		},

		// 0x33
		{
			asm:    "INC SP",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.SP++
			},
		},

		// 0x34
		{
			asm:    "INC (HL)",
			length: 0,
			cycles: 3,

			fn: func(s *State, v Value) {
				b := s.ReadFrom(s.HL)
				s.fl.inc(&b)
				s.WriteTo(s.HL, b)
			},
		},

		// 0x35
		{
			asm:    "DEC (HL)",
			length: 0,
			cycles: 3,

			fn: func(s *State, v Value) {
				b := s.ReadFrom(s.HL)
				s.fl.dec(&b)
				s.WriteTo(s.HL, b)
			},
		},

		// 0x36
		{
			asm:    "LD (HL),u8",
			length: 1,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.WriteTo(s.HL, v.U8())
			},
		},

		// 0x37
		{
			asm:    "SCF",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.Set(Carry)
				s.fl.Clear(Negative | HalfCarry)
			},
		},

		// 0x38
		{
			asm:    "JR C,i8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.JumpIf(s.fl.Has(Carry), v)
			},
		},

		// 0x39
		{
			asm:    "ADD HL,SP",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Add(s.SP, s.fl)
			},
		},

		// 0x3a
		{
			asm:    "LD A,(HL-)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.ReadFrom(s.HL)
				s.HL.Dec()
			},
		},

		// 0x3b
		{
			asm:    "DEC SP",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.SP--
			},
		},

		// 0x3c
		{
			asm:    "INC A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.inc(&s.AF.Hi)
			},
		},

		// 0x3d
		{
			asm:    "DEC A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.dec(&s.AF.Hi)
			},
		},

		// 0x3e
		{
			asm:    "LD A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = v.U8()
			},
		},

		// 0x3f
		{
			asm:    "CCF",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.Flip(Carry)
			},
		},

		// 0x40
		{
			asm:    "LD B,B",
			length: 0,
			cycles: 1,

			fn: nop,
		},

		// 0x41
		{
			asm:    "LD B,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Hi = s.BC.Lo
			},
		},

		// 0x42
		{
			asm:    "LD B,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Hi = s.DE.Hi
			},
		},

		// 0x43
		{
			asm:    "LD B,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Hi = s.DE.Lo
			},
		},

		// 0x44
		{
			asm:    "LD B,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Hi = s.HL.Hi
			},
		},

		// 0x45
		{
			asm:    "LD B,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Hi = s.HL.Lo
			},
		},

		// 0x46
		{
			asm:    "LD B,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.BC.Hi = s.ReadFrom(s.HL)
			},
		},

		// 0x47
		{
			asm:    "LD B,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Hi = s.AF.Hi
			},
		},

		// 0x48
		{
			asm:    "LD C,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.BC.Hi
			},
		},

		// 0x49
		{
			asm:    "LD C,C",
			length: 0,
			cycles: 1,

			fn: nop,
		},

		// 0x4a
		{
			asm:    "LD C,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.DE.Hi
			},
		},

		// 0x4b
		{
			asm:    "LD C,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.DE.Lo
			},
		},

		// 0x4c
		{
			asm:    "LD C,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.HL.Hi
			},
		},

		// 0x4d
		{
			asm:    "LD C,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.HL.Lo
			},
		},

		// 0x4e
		{
			asm:    "LD C,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.ReadFrom(s.HL)
			},
		},

		// 0x4f
		{
			asm:    "LD C,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.AF.Hi
			},
		},

		// 0x50
		{
			asm:    "LD D,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Hi = s.BC.Hi
			},
		},

		// 0x51
		{
			asm:    "LD D,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Hi = s.BC.Lo
			},
		},

		// 0x52
		{
			asm:    "LD D,D",
			length: 0,
			cycles: 1,

			fn: nop,
		},

		// 0x53
		{
			asm:    "LD D,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Hi = s.DE.Lo
			},
		},

		// 0x54
		{
			asm:    "LD D,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Hi = s.HL.Hi
			},
		},

		// 0x55
		{
			asm:    "LD D,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Hi = s.HL.Lo
			},
		},

		// 0x56
		{
			asm:    "LD D,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.DE.Hi = s.ReadFrom(s.HL)
			},
		},

		// 0x57
		{
			asm:    "LD D,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Hi = s.AF.Hi
			},
		},

		// 0x58
		{
			asm:    "LD E,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Lo = s.BC.Hi
			},
		},

		// 0x59
		{
			asm:    "LD E,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Lo = s.BC.Lo
			},
		},

		// 0x5a
		{
			asm:    "LD E,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Lo = s.DE.Hi
			},
		},

		// 0x5b
		{
			asm:    "LD E,E",
			length: 0,
			cycles: 1,

			fn: nop,
		},

		// 0x5c
		{
			asm:    "LD E,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Lo = s.HL.Hi
			},
		},

		// 0x5d
		{
			asm:    "LD E,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Lo = s.HL.Lo
			},
		},

		// 0x5e
		{
			asm:    "LD E,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.DE.Lo = s.ReadFrom(s.HL)
			},
		},

		// 0x5f
		{
			asm:    "LD E,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Lo = s.AF.Hi
			},
		},

		// 0x60
		{
			asm:    "LD H,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Hi = s.BC.Hi
			},
		},

		// 0x61
		{
			asm:    "LD H,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Hi = s.BC.Lo
			},
		},

		// 0x62
		{
			asm:    "LD H,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Hi = s.DE.Hi
			},
		},

		// 0x63
		{
			asm:    "LD H,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Hi = s.DE.Lo
			},
		},

		// 0x64
		{
			asm:    "LD H,H",
			length: 0,
			cycles: 1,

			fn: nop,
		},

		// 0x65
		{
			asm:    "LD H,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Hi = s.HL.Lo
			},
		},

		// 0x66
		{
			asm:    "LD H,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Hi = s.ReadFrom(s.HL)
			},
		},

		// 0x67
		{
			asm:    "LD H,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Hi = s.AF.Hi
			},
		},

		// 0x68
		{
			asm:    "LD L,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Lo = s.BC.Hi
			},
		},

		// 0x69
		{
			asm:    "LD L,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Lo = s.BC.Lo
			},
		},

		// 0x6a
		{
			asm:    "LD L,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Lo = s.DE.Hi
			},
		},

		// 0x6b
		{
			asm:    "LD L,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Lo = s.DE.Lo
			},
		},

		// 0x6c
		{
			asm:    "LD L,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Lo = s.HL.Hi
			},
		},

		// 0x6d
		{
			asm:    "LD L,L",
			length: 0,
			cycles: 1,

			fn: nop,
		},

		// 0x6e
		{
			asm:    "LD L,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Lo = s.ReadFrom(s.HL)
			},
		},

		// 0x6f
		{
			asm:    "LD L,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Lo = s.AF.Hi
			},
		},

		// 0x70
		{
			asm:    "LD (HL),B",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.WriteTo(s.HL, s.BC.Hi)
			},
		},

		// 0x71
		{
			asm:    "LD (HL),C",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.WriteTo(s.HL, s.BC.Lo)
			},
		},

		// 0x72
		{
			asm:    "LD (HL),D",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.WriteTo(s.HL, s.DE.Hi)
			},
		},

		// 0x73
		{
			asm:    "LD (HL),E",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.WriteTo(s.HL, s.DE.Lo)
			},
		},

		// 0x74
		{
			asm:    "LD (HL),H",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.WriteTo(s.HL, s.HL.Hi)
			},
		},

		// 0x75
		{
			asm:    "LD (HL),L",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.WriteTo(s.HL, s.HL.Lo)
			},
		},

		// 0x76
		{
			asm:    "HALT",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				// TODO
			},
		},

		// 0x77
		{
			asm:    "LD (HL),A",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.WriteTo(s.HL, s.AF.Hi)
			},
		},

		// 0x78
		{
			asm:    "LD A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.BC.Hi
			},
		},

		// 0x79
		{
			asm:    "LD A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.BC.Lo
			},
		},

		// 0x7a
		{
			asm:    "LD A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.DE.Hi
			},
		},

		// 0x7b
		{
			asm:    "LD A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.DE.Lo
			},
		},

		// 0x7c
		{
			asm:    "LD A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.HL.Hi
			},
		},

		// 0x7d
		{
			asm:    "LD A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.HL.Lo
			},
		},

		// 0x7e
		{
			asm:    "LD A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.ReadFrom(s.HL)
			},
		},

		// 0x7f
		{
			asm:    "LD A,A",
			length: 0,
			cycles: 1,

			fn: nop,
		},

		// 0x80
		{
			asm:    "ADD A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.add(&s.AF.Hi, s.BC.Hi)
			},
		},

		// 0x81
		{
			asm:    "ADD A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.add(&s.AF.Hi, s.BC.Lo)
			},
		},

		// 0x82
		{
			asm:    "ADD A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.add(&s.AF.Hi, s.DE.Hi)
			},
		},

		// 0x83
		{
			asm:    "ADD A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.add(&s.AF.Hi, s.DE.Lo)
			},
		},

		// 0x84
		{
			asm:    "ADD A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.add(&s.AF.Hi, s.HL.Hi)
			},
		},

		// 0x85
		{
			asm:    "ADD A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.add(&s.AF.Hi, s.HL.Lo)
			},
		},

		// 0x86
		{
			asm:    "ADD A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.add(&s.AF.Hi, s.ReadFrom(s.HL))
			},
		},

		// 0x87
		{
			asm:    "ADD A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.add(&s.AF.Hi, s.AF.Hi)
			},
		},

		// 0x88
		{
			asm:    "ADC A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.adc(s.BC.Hi)
			},
		},

		// 0x89
		{
			asm:    "ADC A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.adc(s.BC.Lo)
			},
		},

		// 0x8a
		{
			asm:    "ADC A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.adc(s.DE.Hi)
			},
		},

		// 0x8b
		{
			asm:    "ADC A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.adc(s.DE.Lo)
			},
		},

		// 0x8c
		{
			asm:    "ADC A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.adc(s.HL.Hi)
			},
		},

		// 0x8d
		{
			asm:    "ADC A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.adc(s.HL.Lo)
			},
		},

		// 0x8e
		{
			asm:    "ADC A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.adc(s.ReadFrom(s.HL))
			},
		},

		// 0x8f
		{
			asm:    "ADC A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.adc(s.AF.Hi)
			},
		},

		// 0x90
		{
			asm:    "SUB A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x91
		{
			asm:    "SUB A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x92
		{
			asm:    "SUB A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x93
		{
			asm:    "SUB A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x94
		{
			asm:    "SUB A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x95
		{
			asm:    "SUB A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x96
		{
			asm:    "SUB A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x97
		{
			asm:    "SUB A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x98
		{
			asm:    "SBC A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x99
		{
			asm:    "SBC A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9a
		{
			asm:    "SBC A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9b
		{
			asm:    "SBC A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9c
		{
			asm:    "SBC A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9d
		{
			asm:    "SBC A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9e
		{
			asm:    "SBC A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0x9f
		{
			asm:    "SBC A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xa0
		{
			asm:    "AND A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xa1
		{
			asm:    "AND A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xa2
		{
			asm:    "AND A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xa3
		{
			asm:    "AND A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xa4
		{
			asm:    "AND A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xa5
		{
			asm:    "AND A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xa6
		{
			asm:    "AND A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xa7
		{
			asm:    "AND A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xa8
		{
			asm:    "XOR A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xa9
		{
			asm:    "XOR A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xaa
		{
			asm:    "XOR A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xab
		{
			asm:    "XOR A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xac
		{
			asm:    "XOR A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xad
		{
			asm:    "XOR A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xae
		{
			asm:    "XOR A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xaf
		{
			asm:    "XOR A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xb0
		{
			asm:    "OR A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xb1
		{
			asm:    "OR A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xb2
		{
			asm:    "OR A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xb3
		{
			asm:    "OR A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xb4
		{
			asm:    "OR A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xb5
		{
			asm:    "OR A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xb6
		{
			asm:    "OR A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xb7
		{
			asm:    "OR A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xb8
		{
			asm:    "CP A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xb9
		{
			asm:    "CP A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xba
		{
			asm:    "CP A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xbb
		{
			asm:    "CP A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xbc
		{
			asm:    "CP A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xbd
		{
			asm:    "CP A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xbe
		{
			asm:    "CP A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xbf
		{
			asm:    "CP A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xc0
		{
			asm:    "RET NZ",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xc1
		{
			asm:    "POP BC",
			length: 0,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xc2
		{
			asm:    "JP NZ,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xc3
		{
			asm:    "JP u16",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xc4
		{
			asm:    "CALL NZ,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xc5
		{
			asm:    "PUSH BC",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xc6
		{
			asm:    "ADD A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xc7
		{
			asm:    "RST 00h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xc8
		{
			asm:    "RET Z",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xc9
		{
			asm:    "RET",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xca
		{
			asm:    "JP Z,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xcb
		{
			asm:    "PREFIX CB",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xcc
		{
			asm:    "CALL Z,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xcd
		{
			asm:    "CALL u16",
			length: 2,
			cycles: 6,

			fn: func(s *State, v Value) {

			},
		},

		// 0xce
		{
			asm:    "ADC A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xcf
		{
			asm:    "RST 08h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xd0
		{
			asm:    "RET NC",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xd1
		{
			asm:    "POP DE",
			length: 0,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xd2
		{
			asm:    "JP NC,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xd3
		{
			asm:    "UNUSED",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {

			},
		},

		// 0xd4
		{
			asm:    "CALL NC,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xd5
		{
			asm:    "PUSH DE",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xd6
		{
			asm:    "SUB A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xd7
		{
			asm:    "RST 10h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xd8
		{
			asm:    "RET C",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xd9
		{
			asm:    "RETI",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xda
		{
			asm:    "JP C,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xdb
		{
			asm:    "UNUSED",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {

			},
		},

		// 0xdc
		{
			asm:    "CALL C,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xdd
		{
			asm:    "UNUSED",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {

			},
		},

		// 0xde
		{
			asm:    "SBC A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xdf
		{
			asm:    "RST 18h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xe0
		{
			asm:    "LD (FF00+u8),A",
			length: 1,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xe1
		{
			asm:    "POP HL",
			length: 0,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xe2
		{
			asm:    "LD (FF00+C),A",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xe3
		{
			asm:    "UNUSED",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {

			},
		},

		// 0xe4
		{
			asm:    "UNUSED",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {

			},
		},

		// 0xe5
		{
			asm:    "PUSH HL",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xe6
		{
			asm:    "AND A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xe7
		{
			asm:    "RST 20h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xe8
		{
			asm:    "ADD SP,i8",
			length: 1,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xe9
		{
			asm:    "JP HL",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xea
		{
			asm:    "LD (u16),A",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xeb
		{
			asm:    "UNUSED",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {

			},
		},

		// 0xec
		{
			asm:    "UNUSED",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {

			},
		},

		// 0xed
		{
			asm:    "UNUSED",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {

			},
		},

		// 0xee
		{
			asm:    "XOR A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xef
		{
			asm:    "RST 28h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xf0
		{
			asm:    "LD A,(FF00+u8)",
			length: 1,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xf1
		{
			asm:    "POP AF",
			length: 0,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xf2
		{
			asm:    "LD A,(FF00+C)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xf3
		{
			asm:    "DI",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xf4
		{
			asm:    "UNUSED",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {

			},
		},

		// 0xf5
		{
			asm:    "PUSH AF",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xf6
		{
			asm:    "OR A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xf7
		{
			asm:    "RST 30h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xf8
		{
			asm:    "LD HL,SP+i8",
			length: 1,
			cycles: 3,

			fn: func(s *State, v Value) {

			},
		},

		// 0xf9
		{
			asm:    "LD SP,HL",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xfa
		{
			asm:    "LD A,(u16)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},

		// 0xfb
		{
			asm:    "EI",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {

			},
		},

		// 0xfc
		{
			asm:    "UNUSED",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {

			},
		},

		// 0xfd
		{
			asm:    "UNUSED",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {

			},
		},

		// 0xfe
		{
			asm:    "CP A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {

			},
		},

		// 0xff
		{
			asm:    "RST 38h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {

			},
		},
	},
}
