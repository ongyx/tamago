package tamago

type fn func(s *State, v uint16)

// An instruction to execute in the CPU.
// assembly is the instruction in assembly as text (i.e nop, etc.)
// cycles is the amount of transistor cycles taken to execute the instruction.
// fn is the function to run when this instruction is executed.
type Instruction struct {
	assembly string
	cycles   uint8
	length   uint8
	fn       *fn
}
