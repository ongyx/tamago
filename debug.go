package tamago

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	banner = "tamago debug shell (type 'help' for help)"
	help   = `
help                 Show this message.
load <filename>      Load a rom by filename.
loadboot <filename>  Load a bootrom by filename.
step                 Execute the next instruction.
show                 Show the state of the registers.
dump                 Dump the ROM and RAM to disk.
peek                 Show the next instruction to execute.
exit                 Leave the shell.
`
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

type DummyRenderer struct{}

func (dr *DummyRenderer) Read(x, y int) Colour     { return Colour{0, 0, 0, 0} }
func (dr *DummyRenderer) Write(x, y int, c Colour) {}

func input(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

func Shell() {
	dr := &DummyRenderer{}
	cpu := NewCPU(dr)

	loaded := false

	fmt.Println(banner)

loop:
	for {
		cmd := input("> ")
		if cmd == "" {
			continue
		}

		cmds := strings.Split(cmd, " ")
		for len(cmds) < 2 {
			cmds = append(cmds, "")
		}

		switch cmds[0] {

		case "help":
			fmt.Println(help)

		case "load":
			if err := cpu.Load(cmds[1]); err != nil {
				fmt.Println(err)
			} else {
				loaded = true
			}

		case "loadboot":
			if err := cpu.LoadBoot(cmds[1]); err != nil {
				fmt.Println(err)
			} else {
				loaded = true
			}

		case "step":
			if !loaded {
				fmt.Println("no rom loaded!")
			} else {
				cpu.step()
			}

		case "show":
			fmt.Println(cpu.state.String())

		case "dump":
			cpu.state.DebugDump()

		case "peek":
			fmt.Println(cpu.table[cpu.state.PC].asm)

		case "exit":
			break loop

		case "":
		default:
			fmt.Println("invalid command!")

		}
	}
}
