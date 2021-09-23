package main

import (
	"flag"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ongyx/tamago"
)

var (
	rom, bootrom string
)

func init() {
	flag.StringVar(&rom, "rom", "", "rom file")
	flag.StringVar(&bootrom, "bootrom", "", "bootrom file")
}

func main() {
	game := tamago.NewGame()

	flag.Parse()

	if bootrom != "" {
		game.S.LoadBoot(bootrom)
	}

	if rom != "" {
		game.S.Load(rom)
	}

	ebiten.SetWindowSize(256, 256)
	ebiten.SetWindowTitle("tamago")
	if err := ebiten.RunGame(game); err != nil {
		fmt.Println(err)
	}

}
