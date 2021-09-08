package tamago

import (
	"regexp"
	"strings"
)

var (
	nop    = func(s *State, v Value) {}
	unused = &Instruction{
		asm:    "UNUSED",
		length: 0,
		cycles: 0,
		fn:     nil,
	}

	ValueRegex = regexp.MustCompile(`([ui]8)`)
)

// An instruction to execute in the CPU.
//
// assembly is the instruction in assembly as text (i.e nop, etc.)
//
// length is the number of bytes the operand takes (0 for no operand, 1 for uint8/int8, 2 for uint16).
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

// Return the instruction in assembly formatted with the operand value.
func (i *Instruction) Asm(v Value) string {
	switch i.length {

	case 1:
		t := ValueRegex.FindString(i.asm)

		var num int64

		if t == "u8" {
			num = int64(v.U8())
		} else if t == "i8" {
			num = int64(v.S8())
		}

		return ValueRegex.ReplaceAllLiteralString(i.asm, "$"+tohex(num))

	case 2:
		return strings.Replace(i.asm, "u16", "$"+tohex(int64(v.U16())), 1)

	default:
		return i.asm // no formatting required

	}
}
