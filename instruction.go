package tamago

var (
	nop    = func(s *State, v Value) {}
	unused = &Instruction{
		asm:    "UNUSED",
		length: 0,
		cycles: 0,
		fn:     nil,
	}
)

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
	length int
	cycles int
	fn     func(s *State, v Value)
}
