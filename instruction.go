package tamago

// An instruction to execute in the CPU.
//
// assembly is the instruction in assembly as text (i.e nop, etc.)
//
// cycles is the amount of machine cycles taken to execute the instruction.
// If it is 0, the instruction has to manually increment cycles.
//
// fn is the function to run when this instruction is executed.
// If it is nil, the instruction is unused.
type Instruction struct {
	asm    string
	length uint8
	cycles uint8
	fn     func(s *State, v uint16)
}
