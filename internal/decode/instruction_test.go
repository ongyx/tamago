package decode

import "testing"

var illegalOpcodes = []uint8{
	0xD3,
	0xDB,
	0xDD,
	0xE3,
	0xE4,
	0xEB,
	0xEC,
	0xED,
	0xF4,
	0xFC,
	0xFD,
}

var validInstructions = map[uint8]Instruction{
	0x00: {Opcode: NOP},
	0x01: {Opcode: LD_R16_D16, WordOperand: WordRegisterBC},
	0x02: {Opcode: LD_R16P_A, WordOperand: WordRegisterBC},
	0x03: {Opcode: INC_R16, WordOperand: WordRegisterBC},
	0x04: {Opcode: INC_R8, SrcOperand: RegisterB},
	0x05: {Opcode: DEC_R8, SrcOperand: RegisterB},
	0x06: {Opcode: LD_R8_D8, DstOperand: RegisterB},
	0x07: {Opcode: RLCA},
	0x08: {Opcode: LD_A16_SP},
	0x09: {Opcode: ADD_HL_R16, WordOperand: WordRegisterBC},
	0x0A: {Opcode: LD_A_R16P, WordOperand: WordRegisterBC},
	0x0B: {Opcode: DEC_R16, WordOperand: WordRegisterBC},
	0x0C: {Opcode: INC_R8, SrcOperand: RegisterC},
	0x0D: {Opcode: DEC_R8, SrcOperand: RegisterC},
	0x0E: {Opcode: LD_R8_D8, DstOperand: RegisterC},
	0x0F: {Opcode: RRCA},

	0x10: {Opcode: STOP},
	0x11: {Opcode: LD_R16_D16, WordOperand: WordRegisterDE},
	0x12: {Opcode: LD_R16P_A, WordOperand: WordRegisterDE},
	0x13: {Opcode: INC_R16, WordOperand: WordRegisterDE},
	0x14: {Opcode: INC_R8, SrcOperand: RegisterD},
	0x15: {Opcode: DEC_R8, SrcOperand: RegisterD},
	0x16: {Opcode: LD_R8_D8, DstOperand: RegisterD},
	0x17: {Opcode: RLA},
	0x18: {Opcode: JR_CC_S8, Condition: ConditionNone},
	0x19: {Opcode: ADD_HL_R16, WordOperand: WordRegisterDE},
	0x1A: {Opcode: LD_A_R16P, WordOperand: WordRegisterDE},
	0x1B: {Opcode: DEC_R16, WordOperand: WordRegisterDE},
	0x1C: {Opcode: INC_R8, SrcOperand: RegisterE},
	0x1D: {Opcode: DEC_R8, SrcOperand: RegisterE},
	0x1E: {Opcode: LD_R8_D8, DstOperand: RegisterE},
	0x1F: {Opcode: RRA},

	0x20: {Opcode: JR_CC_S8, Condition: ConditionNotZero},
	0x21: {Opcode: LD_R16_D16, WordOperand: WordRegisterHL},
	0x22: {Opcode: LD_HLPI_A},
	0x23: {Opcode: INC_R16, WordOperand: WordRegisterHL},
	0x24: {Opcode: INC_R8, SrcOperand: RegisterH},
	0x25: {Opcode: DEC_R8, SrcOperand: RegisterH},
	0x26: {Opcode: LD_R8_D8, DstOperand: RegisterH},
	0x27: {Opcode: DAA},
	0x28: {Opcode: JR_CC_S8, Condition: ConditionZero},
	0x29: {Opcode: ADD_HL_R16, WordOperand: WordRegisterHL},
	0x2A: {Opcode: LD_A_HLPI},
	0x2B: {Opcode: DEC_R16, WordOperand: WordRegisterHL},
	0x2C: {Opcode: INC_R8, SrcOperand: RegisterL},
	0x2D: {Opcode: DEC_R8, SrcOperand: RegisterL},
	0x2E: {Opcode: LD_R8_D8, DstOperand: RegisterL},
	0x2F: {Opcode: CPL},

	0x30: {Opcode: JR_CC_S8, Condition: ConditionNotCarry},
	0x31: {Opcode: LD_R16_D16, WordOperand: WordRegisterSP},
	0x32: {Opcode: LD_HLPD_A},
	0x33: {Opcode: INC_R16, WordOperand: WordRegisterSP},
	0x34: {Opcode: INC_R8, SrcOperand: RegisterHLP},
	0x35: {Opcode: DEC_R8, SrcOperand: RegisterHLP},
	0x36: {Opcode: LD_R8_D8, DstOperand: RegisterHLP},
	0x37: {Opcode: SCF},
	0x38: {Opcode: JR_CC_S8, Condition: ConditionCarry},
	0x39: {Opcode: ADD_HL_R16, WordOperand: WordRegisterSP},
	0x3A: {Opcode: LD_A_HLPD},
	0x3B: {Opcode: DEC_R16, WordOperand: WordRegisterSP},
	0x3C: {Opcode: INC_R8, SrcOperand: RegisterA},
	0x3D: {Opcode: DEC_R8, SrcOperand: RegisterA},
	0x3E: {Opcode: LD_R8_D8, DstOperand: RegisterA},
	0x3F: {Opcode: CCF},

	0x40: {Opcode: LD_R8_R8, DstOperand: RegisterB, SrcOperand: RegisterB},
	0x41: {Opcode: LD_R8_R8, DstOperand: RegisterB, SrcOperand: RegisterC},
	0x42: {Opcode: LD_R8_R8, DstOperand: RegisterB, SrcOperand: RegisterD},
	0x43: {Opcode: LD_R8_R8, DstOperand: RegisterB, SrcOperand: RegisterE},
	0x44: {Opcode: LD_R8_R8, DstOperand: RegisterB, SrcOperand: RegisterH},
	0x45: {Opcode: LD_R8_R8, DstOperand: RegisterB, SrcOperand: RegisterL},
	0x46: {Opcode: LD_R8_R8, DstOperand: RegisterB, SrcOperand: RegisterHLP},
	0x47: {Opcode: LD_R8_R8, DstOperand: RegisterB, SrcOperand: RegisterA},
	0x48: {Opcode: LD_R8_R8, DstOperand: RegisterC, SrcOperand: RegisterB},
	0x49: {Opcode: LD_R8_R8, DstOperand: RegisterC, SrcOperand: RegisterC},
	0x4A: {Opcode: LD_R8_R8, DstOperand: RegisterC, SrcOperand: RegisterD},
	0x4B: {Opcode: LD_R8_R8, DstOperand: RegisterC, SrcOperand: RegisterE},
	0x4C: {Opcode: LD_R8_R8, DstOperand: RegisterC, SrcOperand: RegisterH},
	0x4D: {Opcode: LD_R8_R8, DstOperand: RegisterC, SrcOperand: RegisterL},
	0x4E: {Opcode: LD_R8_R8, DstOperand: RegisterC, SrcOperand: RegisterHLP},
	0x4F: {Opcode: LD_R8_R8, DstOperand: RegisterC, SrcOperand: RegisterA},

	0x50: {Opcode: LD_R8_R8, DstOperand: RegisterD, SrcOperand: RegisterB},
	0x51: {Opcode: LD_R8_R8, DstOperand: RegisterD, SrcOperand: RegisterC},
	0x52: {Opcode: LD_R8_R8, DstOperand: RegisterD, SrcOperand: RegisterD},
	0x53: {Opcode: LD_R8_R8, DstOperand: RegisterD, SrcOperand: RegisterE},
	0x54: {Opcode: LD_R8_R8, DstOperand: RegisterD, SrcOperand: RegisterH},
	0x55: {Opcode: LD_R8_R8, DstOperand: RegisterD, SrcOperand: RegisterL},
	0x56: {Opcode: LD_R8_R8, DstOperand: RegisterD, SrcOperand: RegisterHLP},
	0x57: {Opcode: LD_R8_R8, DstOperand: RegisterD, SrcOperand: RegisterA},
	0x58: {Opcode: LD_R8_R8, DstOperand: RegisterE, SrcOperand: RegisterB},
	0x59: {Opcode: LD_R8_R8, DstOperand: RegisterE, SrcOperand: RegisterC},
	0x5A: {Opcode: LD_R8_R8, DstOperand: RegisterE, SrcOperand: RegisterD},
	0x5B: {Opcode: LD_R8_R8, DstOperand: RegisterE, SrcOperand: RegisterE},
	0x5C: {Opcode: LD_R8_R8, DstOperand: RegisterE, SrcOperand: RegisterH},
	0x5D: {Opcode: LD_R8_R8, DstOperand: RegisterE, SrcOperand: RegisterL},
	0x5E: {Opcode: LD_R8_R8, DstOperand: RegisterE, SrcOperand: RegisterHLP},
	0x5F: {Opcode: LD_R8_R8, DstOperand: RegisterE, SrcOperand: RegisterA},

	0x60: {Opcode: LD_R8_R8, DstOperand: RegisterH, SrcOperand: RegisterB},
	0x61: {Opcode: LD_R8_R8, DstOperand: RegisterH, SrcOperand: RegisterC},
	0x62: {Opcode: LD_R8_R8, DstOperand: RegisterH, SrcOperand: RegisterD},
	0x63: {Opcode: LD_R8_R8, DstOperand: RegisterH, SrcOperand: RegisterE},
	0x64: {Opcode: LD_R8_R8, DstOperand: RegisterH, SrcOperand: RegisterH},
	0x65: {Opcode: LD_R8_R8, DstOperand: RegisterH, SrcOperand: RegisterL},
	0x66: {Opcode: LD_R8_R8, DstOperand: RegisterH, SrcOperand: RegisterHLP},
	0x67: {Opcode: LD_R8_R8, DstOperand: RegisterH, SrcOperand: RegisterA},
	0x68: {Opcode: LD_R8_R8, DstOperand: RegisterL, SrcOperand: RegisterB},
	0x69: {Opcode: LD_R8_R8, DstOperand: RegisterL, SrcOperand: RegisterC},
	0x6A: {Opcode: LD_R8_R8, DstOperand: RegisterL, SrcOperand: RegisterD},
	0x6B: {Opcode: LD_R8_R8, DstOperand: RegisterL, SrcOperand: RegisterE},
	0x6C: {Opcode: LD_R8_R8, DstOperand: RegisterL, SrcOperand: RegisterH},
	0x6D: {Opcode: LD_R8_R8, DstOperand: RegisterL, SrcOperand: RegisterL},
	0x6E: {Opcode: LD_R8_R8, DstOperand: RegisterL, SrcOperand: RegisterHLP},
	0x6F: {Opcode: LD_R8_R8, DstOperand: RegisterL, SrcOperand: RegisterA},

	0x70: {Opcode: LD_R8_R8, DstOperand: RegisterHLP, SrcOperand: RegisterB},
	0x71: {Opcode: LD_R8_R8, DstOperand: RegisterHLP, SrcOperand: RegisterC},
	0x72: {Opcode: LD_R8_R8, DstOperand: RegisterHLP, SrcOperand: RegisterD},
	0x73: {Opcode: LD_R8_R8, DstOperand: RegisterHLP, SrcOperand: RegisterE},
	0x74: {Opcode: LD_R8_R8, DstOperand: RegisterHLP, SrcOperand: RegisterH},
	0x75: {Opcode: LD_R8_R8, DstOperand: RegisterHLP, SrcOperand: RegisterL},
	0x76: {Opcode: HALT},
	0x77: {Opcode: LD_R8_R8, DstOperand: RegisterHLP, SrcOperand: RegisterA},
	0x78: {Opcode: LD_R8_R8, DstOperand: RegisterA, SrcOperand: RegisterB},
	0x79: {Opcode: LD_R8_R8, DstOperand: RegisterA, SrcOperand: RegisterC},
	0x7A: {Opcode: LD_R8_R8, DstOperand: RegisterA, SrcOperand: RegisterD},
	0x7B: {Opcode: LD_R8_R8, DstOperand: RegisterA, SrcOperand: RegisterE},
	0x7C: {Opcode: LD_R8_R8, DstOperand: RegisterA, SrcOperand: RegisterH},
	0x7D: {Opcode: LD_R8_R8, DstOperand: RegisterA, SrcOperand: RegisterL},
	0x7E: {Opcode: LD_R8_R8, DstOperand: RegisterA, SrcOperand: RegisterHLP},
	0x7F: {Opcode: LD_R8_R8, DstOperand: RegisterA, SrcOperand: RegisterA},

	0xC0: {Opcode: RET_CC, Condition: ConditionNotZero},
	0xC1: {Opcode: POP_R16, WordOperand: WordRegisterBC},
	0xC2: {Opcode: JP_CC_A16, Condition: ConditionNotZero},
	0xC3: {Opcode: JP_CC_A16, Condition: ConditionNone},
	0xC4: {Opcode: CALL_CC_A16, Condition: ConditionNotZero},
	0xC5: {Opcode: PUSH_R16, WordOperand: WordRegisterBC},
	0xC6: {Opcode: ADD_D8},
	0xC7: {Opcode: RST, RSTVector: RST0},
	0xC8: {Opcode: RET_CC, Condition: ConditionZero},
	0xC9: {Opcode: RET_CC, Condition: ConditionNone},
	0xCA: {Opcode: JP_CC_A16, Condition: ConditionZero},
	0xCB: {Opcode: PREFIX},
	0xCC: {Opcode: CALL_CC_A16, Condition: ConditionZero},
	0xCD: {Opcode: CALL_CC_A16, Condition: ConditionNone},
	0xCE: {Opcode: ADC_D8},
	0xCF: {Opcode: RST, RSTVector: RST1},

	0xD0: {Opcode: RET_CC, Condition: ConditionNotCarry},
	0xD1: {Opcode: POP_R16, WordOperand: WordRegisterDE},
	0xD2: {Opcode: JP_CC_A16, Condition: ConditionNotCarry},
	0xD4: {Opcode: CALL_CC_A16, Condition: ConditionNotCarry},
	0xD5: {Opcode: PUSH_R16, WordOperand: WordRegisterDE},
	0xD6: {Opcode: SUB_D8},
	0xD7: {Opcode: RST, RSTVector: RST2},
	0xD8: {Opcode: RET_CC, Condition: ConditionCarry},
	0xD9: {Opcode: RETI},
	0xDA: {Opcode: JP_CC_A16, Condition: ConditionCarry},
	0xDC: {Opcode: CALL_CC_A16, Condition: ConditionCarry},
	0xDE: {Opcode: SBC_D8},
	0xDF: {Opcode: RST, RSTVector: RST3},

	0xE0: {Opcode: LD_A8_A},
	0xE1: {Opcode: POP_R16, WordOperand: WordRegisterHL},
	0xE2: {Opcode: LD_CP_A},
	0xE5: {Opcode: PUSH_R16, WordOperand: WordRegisterHL},
	0xE6: {Opcode: AND_D8},
	0xE7: {Opcode: RST, RSTVector: RST4},
	0xE8: {Opcode: ADD_SP_S8},
	0xE9: {Opcode: JP_HL},
	0xEA: {Opcode: LD_A16_A},
	0xEE: {Opcode: XOR_D8},
	0xEF: {Opcode: RST, RSTVector: RST5},

	0xF0: {Opcode: LD_A_A8},
	0xF1: {Opcode: POP_R16, WordOperand: WordRegisterAF},
	0xF2: {Opcode: LD_A_CP},
	0xF3: {Opcode: DI},
	0xF5: {Opcode: PUSH_R16, WordOperand: WordRegisterAF},
	0xF6: {Opcode: OR_D8},
	0xF7: {Opcode: RST, RSTVector: RST6},
	0xF8: {Opcode: LD_HL_SP_S8},
	0xF9: {Opcode: LD_SP_HL},
	0xFA: {Opcode: LD_A_A16},
	0xFB: {Opcode: EI},
	0xFE: {Opcode: CP_D8},
	0xFF: {Opcode: RST, RSTVector: RST7},
}

func TestIllegalInstructions(t *testing.T) {
	for _, op := range illegalOpcodes {
		_, err := DecodeInstruction(op)
		if err == nil {
			t.Errorf("Opcode %x should be illegal", op)
		}
	}
}

func TestValidInstructions(t *testing.T) {
	// NOTE: Can't use uint8 in the range clause because it overflows.
	for op := range 0xFF + 1 {
		expected, ok := validInstructions[uint8(op)]
		if !ok {
			continue
		}

		got, err := DecodeInstruction(uint8(op))
		if err != nil {
			t.Errorf("Opcode %x should be valid, got err %s", op, err)
		}

		if got != expected {
			t.Errorf("For opcode %#x, expected instruction %v but got %v", op, expected, got)
		}

	}
}
