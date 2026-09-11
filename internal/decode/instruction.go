package decode

import (
	"errors"
)

// The 8-bit registers in the order of the `INC R8`, `DEC R8`, `LD R8, R8`, `LD R8, D8`, `<op> A, R8`, and prefixed instructions.
var r8Order = [...]Register{
	RegisterB,
	RegisterC,
	RegisterD,
	RegisterE,
	RegisterH,
	RegisterL,
	RegisterHLP,
	RegisterA,
}

// The 16-bit registers in the order of the `LD R16, D16`, `ADD HL, R16`, `INC R16` and `DEC 16` instructions.
var r16Order = [...]WordRegister{
	WordRegisterBC,
	WordRegisterDE,
	WordRegisterHL,
	WordRegisterSP,
}

// The 16-bit registers in the order of the `PUSH R16` and `POP R16` instructions.
var r16StackOrder = [...]WordRegister{
	WordRegisterBC,
	WordRegisterDE,
	WordRegisterHL,
	WordRegisterAF,
}

// The opcodes of the `<op> A, R8` instructions.
var opAR8Codes = [...]Opcode{
	ADD_R8,
	ADC_D8,
	SUB_R8,
	SBC_R8,
	AND_R8,
	XOR_R8,
	OR_R8,
	CP_R8,
}

// The opcodes of the `<op> A, D8` instructions.
var opAD8Codes = [...]Opcode{
	ADD_D8,
	ADC_D8,
	SUB_D8,
	SBC_D8,
	AND_D8,
	XOR_D8,
	OR_D8,
	CP_D8,
}

// Error returned by [DecodeInstruction] when the opcode is illegal.
var ErrIllegalOpcode = errors.New("illegal opcode")

// An instruction to execute.
type Instruction struct {
	// The instruction opcode.
	Opcode Opcode
	// The destination 8-bit register, if any.
	DstOperand Register
	// The source 8-bit register, if any.
	SrcOperand Register
	// The source/destination 16-bit register, if any.
	WordOperand WordRegister
	// The jump, call, or return condition, if any.
	Condition Condition
	// The RST vector, if any.
	RSTVector RSTVector
}

// Decodes an instruction from a raw opcode. If the opcode is illegal, [ErrIllegalOpcode] is returned.
func DecodeInstruction(op uint8) (ins Instruction, err error) {
	switch op {
	// Miscellaneous
	case 0x00:
		ins.Opcode = NOP
	case 0xCB:
		ins.Opcode = PREFIX
	case 16:
		ins.Opcode = STOP
	case 0x76:
		ins.Opcode = HALT

	// Interrupts
	case 0xF3:
		ins.Opcode = DI
	case 0xFB:
		ins.Opcode = EI

	// BCD/Carry
	case 0x27:
		ins.Opcode = DAA
	case 0x37:
		ins.Opcode = SCF
	case 0x2F:
		ins.Opcode = CPL
	case 0x3F:
		ins.Opcode = CCF

	// Jumps
	case 0x18, 0x20, 0x28, 0x30, 0x38:
		ins.Opcode = JR_CC_S8

		switch op {
		case 0x18:
			ins.Condition = ConditionNone
		case 0x20:
			ins.Condition = ConditionNotZero
		case 0x28:
			ins.Condition = ConditionZero
		case 0x30:
			ins.Condition = ConditionNotCarry
		case 0x38:
			ins.Condition = ConditionCarry
		}
	case 0xC2, 0xC3, 0xCA, 0xD2, 0xDA:
		ins.Opcode = JP_CC_A16

		switch op {
		case 0xC2:
			ins.Condition = ConditionNotZero
		case 0xC3:
			ins.Condition = ConditionNone
		case 0xCA:
			ins.Condition = ConditionZero
		case 0xD2:
			ins.Condition = ConditionNotCarry
		case 0xDA:
			ins.Condition = ConditionCarry
		}
	case 0xE9:
		ins.Opcode = JP_HL

	// Stack manipulation
	case 0xE8:
		ins.Opcode = ADD_SP_S8
	case 0xC5, 0xD5, 0xE5, 0xF5:
		ins.Opcode = PUSH_R16
		ins.WordOperand = r16StackOrder[(op-0xC5)/16]
	case 0xC4, 0xCC, 0xCD, 0xD4, 0xDC:
		ins.Opcode = CALL_CC_A16

		switch op {
		case 0xC4:
			ins.Condition = ConditionNotZero
		case 0xCC:
			ins.Condition = ConditionZero
		case 0xCD:
			ins.Condition = ConditionNone
		case 0xD4:
			ins.Condition = ConditionNotCarry
		case 0xDC:
			ins.Condition = ConditionCarry
		}
	case 0xC7, 0xCF, 0xD7, 0xDF, 0xE7, 0xEF, 0xF7, 0xFF:
		ins.Opcode = RST
		ins.RSTVector = RSTVector(op - 0xC7)
	case 0xC1, 0xD1, 0xE1, 0xF1:
		ins.Opcode = POP_R16
		ins.WordOperand = r16StackOrder[(op-0xC1)/16]
	case 0xC0, 0xC8, 0xC9, 0xD0, 0xD8:
		ins.Opcode = RET_CC

		switch op {
		case 0xC0:
			ins.Condition = ConditionNotZero
		case 0xC8:
			ins.Condition = ConditionZero
		case 0xC9:
			ins.Condition = ConditionNone
		case 0xD0:
			ins.Condition = ConditionNotCarry
		case 0xD8:
			ins.Condition = ConditionCarry
		}
	case 0xD9:
		ins.Opcode = RETI

	// 8-bit register arithmetic (<op> A, R8 is handled by default case)
	case 0x4, 0xC, 0x14, 0x1C, 0x24, 0x2C, 0x34, 0x3C:
		ins.Opcode = INC_R8
		ins.SrcOperand = r8Order[(op-0x4)/8]
	case 0x5, 0xD, 0x15, 0x1D, 0x25, 0x2D, 0x35, 0x3D:
		ins.Opcode = DEC_R8
		ins.SrcOperand = r8Order[(op-0x5)/8]

	// 16-bit arithmetic
	case 0x3, 0x13, 0x23, 0x33:
		ins.Opcode = INC_R16
		ins.WordOperand = r16Order[(op-0x3)/16]
	case 0xB, 0x1B, 0x2B, 0x3B:
		ins.Opcode = DEC_R16
		ins.WordOperand = r16Order[(op-0xB)/16]
	case 0x9, 0x19, 0x29, 0x39:
		ins.Opcode = ADD_HL_R16
		ins.WordOperand = r16Order[(op-0x9)/16]

	// Register A bitshifts
	case 0x7:
		ins.Opcode = RLCA
	case 0x17:
		ins.Opcode = RLA
	case 0xF:
		ins.Opcode = RRCA
	case 0x1F:
		ins.Opcode = RRA

	// Register A arithmetic with 8-bit immediate operand
	case 0xC6, 0xCE, 0xD6, 0xDE, 0xE6, 0xEE, 0xF6, 0xFE:
		ins.Opcode = opAD8Codes[(op-0xC6)/8]

	// 8-bit loads (LD R8 R8 is handled by default case)
	case 0x6, 0xE, 0x16, 0x1E, 0x26, 0x2E, 0x36, 0x3E:
		ins.Opcode = LD_R8_D8
		ins.DstOperand = r8Order[(op-0x6)/8]
	case 0xE0:
		ins.Opcode = LD_A8_A
	case 0xF0:
		ins.Opcode = LD_A_A8
	case 0xE2:
		ins.Opcode = LD_CP_A
	case 0xF2:
		ins.Opcode = LD_A_CP
	case 0xEA:
		ins.Opcode = LD_A16_A
	case 0xFA:
		ins.Opcode = LD_A_A16

	// 16-bit loads
	case 0x1, 0x11, 0x21, 0x31:
		ins.Opcode = LD_R16_D16
		ins.WordOperand = r16Order[(op-0x1)/16]
	case 0x2, 0x12:
		ins.Opcode = LD_R16P_A
		ins.WordOperand = r16Order[(op-0x2)/16]
	case 0x22:
		ins.Opcode = LD_HLPI_A
	case 0x32:
		ins.Opcode = LD_HLPD_A
	case 0xA, 0x1A:
		ins.Opcode = LD_A_R16P
		ins.WordOperand = r16Order[(op-0xA)/16]
	case 0x2A:
		ins.Opcode = LD_A_HLPI
	case 0x3A:
		ins.Opcode = LD_A_HLPD
	case 0x8:
		ins.Opcode = LD_A16_SP
	case 0xF8:
		ins.Opcode = LD_HL_SP_S8
	case 0xF9:
		ins.Opcode = LD_SP_HL

	default:
		// <op> A, R8 and LD R8, R8 can be decoded more concisely because they reside in the continguous range [0x40, 0xC0).
		if op >= 0x40 && op < 0x80 {
			offset := op - 0x40
			ins.Opcode = LD_R8_R8
			// The group signifies the destination register to load to.
			ins.DstOperand = r8Order[offset/8]
			ins.SrcOperand = r8Order[offset%8]
		} else if op >= 0x80 && op < 0xC0 {
			offset := op - 0x80
			// Each ALU operation takes up a group of 8 instructions, one for each register.
			ins.Opcode = opAR8Codes[offset/8]
			ins.SrcOperand = r8Order[offset%8]
		} else {
			err = ErrIllegalOpcode
		}
	}

	return ins, err
}
