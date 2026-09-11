package decode

// A prefixed instruction to execute.
type PrefixInstruction struct {
	// The opcode.
	Opcode PrefixOpcode
	// The source and destination 8-bit register.
	Operand Register
	// The bit to manipulate, for the [BIT], [RES], and [SET] operations.
	Bit uint8
}

// Decodes a prefix instruction from a raw opcode.
func DecodePrefixInstruction(op uint8) PrefixInstruction {
	var pins PrefixInstruction

	if op < 0x40 {
		// The first 64 instructions correspond to the first 8 prefix opcodes, so the cast here is safe.
		pins.Opcode = PrefixOpcode(op / 8)
		pins.Operand = r8Order[op%8]
	} else {
		// The BIT, RES, and SET groups have 64 instructions each, with 8 sub-groups for each bit to manipulate.
		offset := op - 0x40
		pins.Opcode = PrefixOpcode(uint8(BIT) + offset/64)
		pins.Operand = r8Order[offset%8]
		pins.Bit = (offset % 0x40) / 8
	}

	return pins
}
