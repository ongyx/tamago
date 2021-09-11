package tamago

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Command is a shell command.
// help is the command description.
// nargs is the number of args required for the command. If less than 0, the command accepts any number of arguments.
// fn is a callback function to execute when a command is invoked.
type Command struct {
	help  string
	nargs int
	fn    func(args []string) error
}

// Shell provides a minimal prompt to run some commands.
type Shell struct {
	cmds    map[string]Command
	scanner *bufio.Scanner
}

// Create a new shell.
func NewShell() *Shell {
	sh := &Shell{
		cmds:    make(map[string]Command),
		scanner: bufio.NewScanner(os.Stdin),
	}

	sh.Register("help", Command{
		help:  "show this help message",
		nargs: 0,
		fn: func(args []string) error {
			fmt.Println(sh.Help())
			return nil
		},
	})

	sh.Register("exit", Command{
		help:  "leave the shell",
		nargs: 0,
	})

	return sh
}

// Register a command with name.
func (sh *Shell) Register(name string, cmd Command) {
	sh.cmds[name] = cmd
}

// Return the helptext of all commands as a string.
func (sh *Shell) Help() string {
	var sb strings.Builder

	for name, cmd := range sh.cmds {
		sb.WriteString(name + "\n\t" + cmd.help)
	}

	return sb.String()
}

// Start prompting the user for command input, with p as the prompt text.
// If an error was returned from a command, it will be returned.
func (sh *Shell) Prompt(p string) error {
	for {
		fmt.Print(p)
		sh.scanner.Scan()

		args := strings.Fields(sh.scanner.Text())
		nargs := len(args) - 1

		if nargs == -1 {
			continue
		}

		name := args[0]

		if cmd, ok := sh.cmds[name]; ok {

			if nargs != cmd.nargs && cmd.nargs >= 0 {
				fmt.Printf("%s expects %d args, got %d args", name, cmd.nargs, nargs)
				continue
			}

			if name == "exit" {
				break
			}

			var slice []string
			if nargs == 0 {
				slice = args
			} else {
				slice = args[1:]
			}

			if err := cmd.fn(slice); err != nil {
				return err
			}

		} else {
			fmt.Println("Invalid command: " + args[0])
			continue
		}

	}

	return nil
}
