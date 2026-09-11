package decode

const (
	// Miscellaneous

	NOP Opcode = iota
	PREFIX
	STOP
	HALT

	// Interrupts

	DI
	EI

	// BCD/Carry

	DAA
	SCF
	CPL
	CCF

	// Jumps

	JR_CC_S8
	JP_CC_A16
	JP_HL

	// Stack manipulation

	ADD_SP_S8
	PUSH_R16
	CALL_CC_A16
	RST
	POP_R16
	RET_CC
	RETI

	// 8-bit register arithmetic

	INC_R8
	DEC_R8
	ADD_R8
	ADC_R8
	SUB_R8
	SBC_R8
	AND_R8
	XOR_R8
	OR_R8
	CP_R8

	// 16-bit arithmetic

	INC_R16
	DEC_R16
	ADD_HL_R16

	// Register A bitshifts

	RLCA
	RLA
	RRCA
	RRA

	// Register A arithmetic with 8-bit immediate operand

	ADD_D8
	SUB_D8
	AND_D8
	OR_D8
	ADC_D8
	SBC_D8
	XOR_D8
	CP_D8

	// 8-bit loads

	LD_R8_R8
	LD_R8_D8
	LD_A8_A
	LD_A_A8
	LD_CP_A
	LD_A_CP
	LD_A16_A
	LD_A_A16

	// 16-bit loads

	LD_R16_D16
	LD_R16P_A
	LD_HLPI_A
	LD_HLPD_A
	LD_A_R16P
	LD_A_HLPI
	LD_A_HLPD
	LD_A16_SP
	LD_HL_SP_S8
	LD_SP_HL
)

// An instruction opcode.
//
//go:generate stringer -type=Opcode
type Opcode uint8
