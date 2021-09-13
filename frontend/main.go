package main

import (
	"flag"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

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
	game := tamago.NewGame()

	flag.Parse()

	// TODO: integrate debug with game
	if debug {
		game.C.DebugRun()
		return
	}

	if bootrom != "" {
		game.C.LoadBoot(bootrom)
	}

	if rom != "" {
		game.C.Load(rom)
	}

	ebiten.SetWindowSize(256, 256)
	ebiten.SetWindowTitle("tamago")
	if err := ebiten.RunGame(game); err != nil {
		fmt.Println(err)
	}

}
