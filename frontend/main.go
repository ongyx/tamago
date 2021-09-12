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

type DummyRenderer struct{}

func (rr DummyRenderer) Write(x, y int, c tamago.Colour) {}

func main() {
	cpu := tamago.NewCPU(DummyRenderer{})

	flag.Parse()

	if debug {
		cpu.DebugRun()
	}

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
