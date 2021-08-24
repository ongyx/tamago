package tamago

import (
	"math"
)

// This table contains the normal instructions used by the Game Boy.
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

		// 0x0A
		{
			asm:    "LD A,(BC)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.ReadFrom(s.BC)
			},
		},

		// 0x0B
		{
			asm:    "DEC BC",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.BC.Dec()
			},
		},

		// 0x0C
		{
			asm:    "INC C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.inc(&s.BC.Lo)
			},
		},

		// 0x0D
		{
			asm:    "DEC C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.dec(&s.BC.Lo)
			},
		},

		// 0x0E
		{
			asm:    "LD C,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.BC.Lo = v.U8()
			},
		},

		// 0x0F
		{
			asm:    "RRCA",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				// Take only the last bit
				carry := s.AF.Hi & 0x80
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

		// 0x1A
		{
			asm:    "LD A,(DE)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.ReadFrom(s.DE)
			},
		},

		// 0x1B
		{
			asm:    "DEC DE",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.DE.Dec()
			},
		},

		// 0x1C
		{
			asm:    "INC E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.inc(&s.DE.Lo)
			},
		},

		// 0x1D
		{
			asm:    "DEC E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.dec(&s.DE.Lo)
			},
		},

		// 0x1E
		{
			asm:    "LD E,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.DE.Lo = v.U8()
			},
		},

		// 0x1F
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
						a = (s - 0x06) & math.MaxUint8
					}

					if s.fl.Has(Carry) {
						a -= 0x60
					}

				} else {

					if s.fl.Has(HalfCarry) || (a&0x0F) > 9 {
						a += 0x06
					}

					if s.fl.Has(Carry) || s > 0x9F {
						a += 0x60
					}

				}

				s.AF.Hi = uint8(a & math.MaxUint8)

				s.fl.Clear(HalfCarry)
				s.fl.setIfZero(a)
				s.fl.setIfCarry(a &^ math.MaxUint8)
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

		// 0x2A
		{
			asm:    "LD A,(HL+)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.ReadFrom(s.HL)
				s.HL.Inc()
			},
		},

		// 0x2B
		{
			asm:    "DEC HL",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Dec()
			},
		},

		// 0x2C
		{
			asm:    "INC L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.inc(&s.HL.Lo)
			},
		},

		// 0x2D
		{
			asm:    "DEC L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.dec(&s.HL.Lo)
			},
		},

		// 0x2E
		{
			asm:    "LD L,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Lo = v.U8()
			},
		},

		// 0x2F
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

		// 0x3A
		{
			asm:    "LD A,(HL-)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.ReadFrom(s.HL)
				s.HL.Dec()
			},
		},

		// 0x3B
		{
			asm:    "DEC SP",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.SP--
			},
		},

		// 0x3C
		{
			asm:    "INC A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.inc(&s.AF.Hi)
			},
		},

		// 0x3D
		{
			asm:    "DEC A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.dec(&s.AF.Hi)
			},
		},

		// 0x3E
		{
			asm:    "LD A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = v.U8()
			},
		},

		// 0x3F
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

		// 0x4A
		{
			asm:    "LD C,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.DE.Hi
			},
		},

		// 0x4B
		{
			asm:    "LD C,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.DE.Lo
			},
		},

		// 0x4C
		{
			asm:    "LD C,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.HL.Hi
			},
		},

		// 0x4D
		{
			asm:    "LD C,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.HL.Lo
			},
		},

		// 0x4E
		{
			asm:    "LD C,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.BC.Lo = s.ReadFrom(s.HL)
			},
		},

		// 0x4F
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

		// 0x5A
		{
			asm:    "LD E,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Lo = s.DE.Hi
			},
		},

		// 0x5B
		{
			asm:    "LD E,E",
			length: 0,
			cycles: 1,

			fn: nop,
		},

		// 0x5C
		{
			asm:    "LD E,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Lo = s.HL.Hi
			},
		},

		// 0x5D
		{
			asm:    "LD E,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.DE.Lo = s.HL.Lo
			},
		},

		// 0x5E
		{
			asm:    "LD E,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.DE.Lo = s.ReadFrom(s.HL)
			},
		},

		// 0x5F
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

		// 0x6A
		{
			asm:    "LD L,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Lo = s.DE.Hi
			},
		},

		// 0x6B
		{
			asm:    "LD L,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Lo = s.DE.Lo
			},
		},

		// 0x6C
		{
			asm:    "LD L,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.HL.Lo = s.HL.Hi
			},
		},

		// 0x6D
		{
			asm:    "LD L,L",
			length: 0,
			cycles: 1,

			fn: nop,
		},

		// 0x6E
		{
			asm:    "LD L,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.HL.Lo = s.ReadFrom(s.HL)
			},
		},

		// 0x6F
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

		// 0x7A
		{
			asm:    "LD A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.DE.Hi
			},
		},

		// 0x7B
		{
			asm:    "LD A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.DE.Lo
			},
		},

		// 0x7C
		{
			asm:    "LD A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.HL.Hi
			},
		},

		// 0x7D
		{
			asm:    "LD A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.HL.Lo
			},
		},

		// 0x7E
		{
			asm:    "LD A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.ReadFrom(s.HL)
			},
		},

		// 0x7F
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

		// 0x8A
		{
			asm:    "ADC A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.adc(s.DE.Hi)
			},
		},

		// 0x8B
		{
			asm:    "ADC A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.adc(s.DE.Lo)
			},
		},

		// 0x8C
		{
			asm:    "ADC A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.adc(s.HL.Hi)
			},
		},

		// 0x8D
		{
			asm:    "ADC A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.adc(s.HL.Lo)
			},
		},

		// 0x8E
		{
			asm:    "ADC A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.adc(s.ReadFrom(s.HL))
			},
		},

		// 0x8F
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
				s.fl.sub(s.BC.Hi)
			},
		},

		// 0x91
		{
			asm:    "SUB A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sub(s.BC.Lo)
			},
		},

		// 0x92
		{
			asm:    "SUB A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sub(s.DE.Hi)
			},
		},

		// 0x93
		{
			asm:    "SUB A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sub(s.DE.Lo)
			},
		},

		// 0x94
		{
			asm:    "SUB A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sub(s.DE.Lo)
			},
		},

		// 0x95
		{
			asm:    "SUB A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sub(s.HL.Lo)
			},
		},

		// 0x96
		{
			asm:    "SUB A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.sub(s.ReadFrom(s.HL))
			},
		},

		// 0x97
		{
			asm:    "SUB A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sub(s.AF.Hi)
			},
		},

		// 0x98
		{
			asm:    "SBC A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sbc(s.BC.Hi)
			},
		},

		// 0x99
		{
			asm:    "SBC A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sbc(s.BC.Lo)
			},
		},

		// 0x9A
		{
			asm:    "SBC A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sbc(s.DE.Hi)
			},
		},

		// 0x9B
		{
			asm:    "SBC A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sbc(s.DE.Lo)
			},
		},

		// 0x9C
		{
			asm:    "SBC A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sbc(s.HL.Hi)
			},
		},

		// 0x9D
		{
			asm:    "SBC A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sbc(s.HL.Lo)
			},
		},

		// 0x9E
		{
			asm:    "SBC A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.sbc(s.ReadFrom(s.HL))
			},
		},

		// 0x9F
		{
			asm:    "SBC A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.sbc(s.AF.Hi)
			},
		},

		// 0xA0
		{
			asm:    "AND A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.and(s.BC.Hi)
			},
		},

		// 0xA1
		{
			asm:    "AND A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.and(s.BC.Lo)
			},
		},

		// 0xA2
		{
			asm:    "AND A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.and(s.DE.Hi)
			},
		},

		// 0xA3
		{
			asm:    "AND A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.and(s.DE.Lo)
			},
		},

		// 0xA4
		{
			asm:    "AND A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.and(s.HL.Hi)
			},
		},

		// 0xA5
		{
			asm:    "AND A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.and(s.HL.Lo)
			},
		},

		// 0xA6
		{
			asm:    "AND A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.and(s.ReadFrom(s.HL))
			},
		},

		// 0xA7
		{
			asm:    "AND A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.and(s.AF.Hi)
			},
		},

		// 0xA8
		{
			asm:    "XOR A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.xor(s.BC.Hi)
			},
		},

		// 0xA9
		{
			asm:    "XOR A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.xor(s.BC.Lo)
			},
		},

		// 0xAA
		{
			asm:    "XOR A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.xor(s.DE.Hi)
			},
		},

		// 0xAB
		{
			asm:    "XOR A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.xor(s.DE.Lo)
			},
		},

		// 0xAC
		{
			asm:    "XOR A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.xor(s.HL.Hi)
			},
		},

		// 0xAD
		{
			asm:    "XOR A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.xor(s.HL.Lo)
			},
		},

		// 0xAE
		{
			asm:    "XOR A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.xor(s.ReadFrom(s.HL))
			},
		},

		// 0xAF
		{
			asm:    "XOR A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.xor(s.AF.Hi)
			},
		},

		// 0xB0
		{
			asm:    "OR A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.or(s.BC.Hi)
			},
		},

		// 0xB1
		{
			asm:    "OR A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.or(s.BC.Lo)
			},
		},

		// 0xB2
		{
			asm:    "OR A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.or(s.DE.Hi)
			},
		},

		// 0xB3
		{
			asm:    "OR A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.or(s.DE.Lo)
			},
		},

		// 0xB4
		{
			asm:    "OR A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.or(s.HL.Hi)
			},
		},

		// 0xB5
		{
			asm:    "OR A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.or(s.HL.Lo)
			},
		},

		// 0xB6
		{
			asm:    "OR A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.or(s.ReadFrom(s.HL))
			},
		},

		// 0xB7
		{
			asm:    "OR A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.or(s.AF.Hi)
			},
		},

		// 0xB8
		{
			asm:    "CP A,B",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.cmp(s.BC.Hi)
			},
		},

		// 0xB9
		{
			asm:    "CP A,C",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.cmp(s.BC.Lo)
			},
		},

		// 0xBA
		{
			asm:    "CP A,D",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.cmp(s.DE.Hi)
			},
		},

		// 0xBB
		{
			asm:    "CP A,E",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.cmp(s.DE.Lo)
			},
		},

		// 0xBC
		{
			asm:    "CP A,H",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.cmp(s.HL.Hi)
			},
		},

		// 0xBD
		{
			asm:    "CP A,L",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.cmp(s.HL.Lo)
			},
		},

		// 0xBE
		{
			asm:    "CP A,(HL)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.cmp(s.ReadFrom(s.HL))
			},
		},

		// 0xBF
		{
			asm:    "CP A,A",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.fl.cmp(s.AF.Hi)
			},
		},

		// 0xC0
		{
			asm:    "RET NZ",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {
				if !s.fl.Has(Zero) {
					s.PC = s.Pop()
					s.clock.Step(5)
				} else {
					s.clock.Step(2)
				}
			},
		},

		// 0xC1
		{
			asm:    "POP BC",
			length: 0,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.BC.Set(s.Pop())
			},
		},

		// 0xC2
		{
			asm:    "JP NZ,u16",
			length: 2,
			cycles: 0,

			fn: func(s *State, v Value) {
				if !s.fl.Has(Zero) {
					s.PC = v.U16()
					s.clock.Step(4)
				} else {
					s.clock.Step(3)
				}
			},
		},

		// 0xC3
		{
			asm:    "JP u16",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.PC = v.U16()
			},
		},

		// 0xC4
		{
			asm:    "CALL NZ,u16",
			length: 2,
			cycles: 0,

			fn: func(s *State, v Value) {
				if !s.fl.Has(Zero) {
					s.Push(s.PC)
					s.PC = v.U16()
					s.clock.Step(6)
				} else {
					s.clock.Step(3)
				}
			},
		},

		// 0xC5
		{
			asm:    "PUSH BC",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.BC.Get())
			},
		},

		// 0xC6
		{
			asm:    "ADD A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.add(&s.AF.Hi, v.U8())
			},
		},

		// 0xC7
		{
			asm:    "RST 00h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.PC)
				s.PC = 0x0000
			},
		},

		// 0xC8
		{
			asm:    "RET Z",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {
				if s.fl.Has(Zero) {
					s.PC = s.Pop()
					s.clock.Step(5)
				} else {
					s.clock.Step(2)
				}
			},
		},

		// 0xC9
		{
			asm:    "RET",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.PC = s.Pop()
			},
		},

		// 0xCA
		{
			asm:    "JP Z,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {
				if s.fl.Has(Zero) {
					s.PC = v.U16()
					s.clock.Step(4)
				} else {
					s.clock.Step(3)
				}
			},
		},

		// 0xCB
		{
			asm:    "PREFIX CB",
			length: 0,
			cycles: 0,

			// 0xCB instructions are offset at 0x100 onwards.
			fn: nop,
		},

		// 0xCC
		{
			asm:    "CALL Z,u16",
			length: 2,
			cycles: 0,

			fn: func(s *State, v Value) {
				if s.fl.Has(Zero) {
					s.Push(s.PC)
					s.PC = v.U16()
					s.clock.Step(6)
				} else {
					s.clock.Step(3)
				}
			},
		},

		// 0xCD
		{
			asm:    "CALL u16",
			length: 2,
			cycles: 6,

			fn: func(s *State, v Value) {
				s.Push(s.PC)
				s.PC = v.U16()
			},
		},

		// 0xCE
		{
			asm:    "ADC A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.adc(v.U8())
			},
		},

		// 0xCF
		{
			asm:    "RST 08h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.PC)
				s.PC = 0x0008
			},
		},

		// 0xD0
		{
			asm:    "RET NC",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {
				if !s.fl.Has(Carry) {
					s.PC = s.Pop()
					s.clock.Step(5)
				} else {
					s.clock.Step(2)
				}
			},
		},

		// 0xD1
		{
			asm:    "POP DE",
			length: 0,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.DE.Set(s.Pop())
			},
		},

		// 0xD2
		{
			asm:    "JP NC,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {
				if !s.fl.Has(Carry) {
					s.PC = v.U16()
					s.clock.Step(4)
				} else {
					s.clock.Step(3)
				}
			},
		},

		// 0xD3
		*unused,

		// 0xD4
		{
			asm:    "CALL NC,u16",
			length: 2,
			cycles: 0,

			fn: func(s *State, v Value) {
				if !s.fl.Has(Carry) {
					s.Push(s.PC)
					s.PC = v.U16()
					s.clock.Step(6)
				} else {
					s.clock.Step(3)
				}
			},
		},

		// 0xD5
		{
			asm:    "PUSH DE",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.DE.Get())
			},
		},

		// 0xD6
		{
			asm:    "SUB A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.sub(v.U8())
			},
		},

		// 0xD7
		{
			asm:    "RST 10h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.PC)
				s.PC = 0x0010
			},
		},

		// 0xD8
		{
			asm:    "RET C",
			length: 0,
			cycles: 0,

			fn: func(s *State, v Value) {
				if s.fl.Has(Carry) {
					s.PC = s.Pop()
					s.clock.Step(5)
				} else {
					s.clock.Step(2)
				}
			},
		},

		// 0xD9
		{
			asm:    "RETI",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.intr.Enable()
				s.PC = s.Pop()
			},
		},

		// 0xDA
		{
			asm:    "JP C,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {
				if s.fl.Has(Carry) {
					s.PC = v.U16()
					s.clock.Step(4)
				} else {
					s.clock.Step(3)
				}
			},
		},

		// 0xDB
		*unused,

		// 0xDC
		{
			asm:    "CALL C,u16",
			length: 2,
			cycles: 3,

			fn: func(s *State, v Value) {
				if s.fl.Has(Carry) {
					s.Push(s.PC)
					s.PC = v.U16()
					s.clock.Step(6)
				} else {
					s.clock.Step(3)
				}
			},
		},

		// 0xDD
		*unused,

		// 0xDE
		{
			asm:    "SBC A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.sbc(v.U8())
			},
		},

		// 0xDF
		{
			asm:    "RST 18h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.PC)
				s.PC = 0x0018
			},
		},

		// 0xE0
		{
			asm:    "LD (FF00+u8),A",
			length: 1,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.Write(uint16(0xFF00)+uint16(v.U8()), s.AF.Hi)
			},
		},

		// 0xE1
		{
			asm:    "POP HL",
			length: 0,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.HL.Set(s.Pop())
			},
		},

		// 0xE2
		{
			asm:    "LD (FF00+C),A",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.Write(uint16(0xFF00)+uint16(s.BC.Lo), s.AF.Hi)
			},
		},

		// 0xE3
		*unused,

		// 0xE4
		*unused,

		// 0xE5
		{
			asm:    "PUSH HL",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.HL.Get())
			},
		},

		// 0xE6
		{
			asm:    "AND A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.and(v.U8())
			},
		},

		// 0xE7
		{
			asm:    "RST 20h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.PC)
				s.PC = 0x0020
			},
		},

		// 0xE8
		{
			asm:    "ADD SP,i8",
			length: 1,
			cycles: 4,

			fn: func(s *State, v Value) {
				result := int(s.SP) + int(v.S8())

				s.fl.setIfCarry(result &^ math.MaxUint16)

				s.SP = result & math.MaxUint16

				s.fl.setIfHalfCarry(((s.SP & 0x0F) + (v.S8() & 0x0F)) > 0x0F)

				s.fl.Clear(Zero | Negative)
			},
		},

		// 0xE9
		{
			asm:    "JP HL",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.PC = s.HL.Get()
			},
		},

		// 0xEA
		{
			asm:    "LD (u16),A",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Write(v.U16(), s.AF.Hi)
			},
		},

		// 0xEB
		*unused,

		// 0xEC
		*unused,

		// 0xED
		*unused,

		// 0xEE
		{
			asm:    "XOR A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.xor(v.U8())
			},
		},

		// 0xEF
		{
			asm:    "RST 28h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.PC)
				s.PC = 0x0028
			},
		},

		// 0xF0
		{
			asm:    "LD A,(FF00+u8)",
			length: 1,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.Read(0xFF00 + v.U8())
			},
		},

		// 0xF1
		{
			asm:    "POP AF",
			length: 0,
			cycles: 3,

			fn: func(s *State, v Value) {
				s.AF.Set(s.Pop())
			},
		},

		// 0xF2
		{
			asm:    "LD A,(FF00+C)",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.Read(0xFF00 + s.BC.Lo)
			},
		},

		// 0xF3
		{
			asm:    "DI",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.intr.Disable()
			},
		},

		// 0xF4
		*unused,

		// 0xF5
		{
			asm:    "PUSH AF",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.AF.Get())
			},
		},

		// 0xF6
		{
			asm:    "OR A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.or(v.U8())
			},
		},

		// 0xF7
		{
			asm:    "RST 30h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.PC)
				s.PC = 0x0030
			},
		},

		// 0xF8
		{
			asm:    "LD HL,SP+i8",
			length: 1,
			cycles: 3,

			fn: func(s *State, v Value) {
				result := int(s.SP) + int(v.S8())

				s.fl.setIfCarry(result &^ math.MaxUint16)
				s.fl.setIfHalfCarry(((s.SP & 0x0F) + (v.S8() & 0x0F)) > 0x0F)

				s.fl.Clear(Zero | Negative)

				s.HL.Set(uint16(result & math.MaxUint16))
			},
		},

		// 0xF9
		{
			asm:    "LD SP,HL",
			length: 0,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.SP = s.HL.Get()
			},
		},

		// 0xFA
		{
			asm:    "LD A,(u16)",
			length: 2,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.AF.Hi = s.Read(v.U16())
			},
		},

		// 0xFB
		{
			asm:    "EI",
			length: 0,
			cycles: 1,

			fn: func(s *State, v Value) {
				s.intr.Enable()
			},
		},

		// 0xFC
		*unused,

		// 0xFD
		*unused,

		// 0xFE
		{
			asm:    "CP A,u8",
			length: 1,
			cycles: 2,

			fn: func(s *State, v Value) {
				s.fl.cmp(v.U8())
			},
		},

		// 0xFF
		{
			asm:    "RST 38h",
			length: 0,
			cycles: 4,

			fn: func(s *State, v Value) {
				s.Push(s.PC)
				s.PC = 0x0038
			},
		},
	},
}
