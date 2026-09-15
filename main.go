package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ongyx/tamago/internal/core"
	"github.com/ongyx/tamago/internal/util"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <path to ROM file>\n", os.Args[0])
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("error: failed to load ROM: %s", err)
	}
	cart := core.NewCart(data)

	ebiten.SetWindowTitle("tamago")
	ebiten.SetWindowSize(util.ScreenWidth*screenScale, util.ScreenHeight*screenScale)

	cpu := core.NewCPU()
	cpu.LoadCart(cart)
	g := NewGame(cpu)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("error: failed to run game: %s", err)
	}
}
