package main

import (
	"flag"
	"fmt"

	"github.com/ongyx/tamago"
)

var (
	debug        bool
	rom, bootrom string
)

func init() {
	flag.BoolVar(&debug, "debug", false, "enter debug shell")
	flag.StringVar(&rom, "rom", "", "rom file")
	flag.StringVar(&bootrom, "bootrom", "", "bootrom file")
}

func main() {
	flag.Parse()

	if debug {
		tamago.Shell()
	} else {
		cpu := tamago.NewCPU(&tamago.DummyRenderer{})

		if bootrom != "" {
			cpu.LoadBoot(bootrom)
		}

		if rom != "" {
			cpu.Load(rom)
		}

		if err := cpu.Run(); err != nil {
			fmt.Println(err)
		}
	}
}
