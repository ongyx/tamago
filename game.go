package main

import (
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ongyx/tamago/internal/core"
	. "github.com/ongyx/tamago/internal/util"
)

const (
	screenScale = 4
)

var _ ebiten.Game = (*Game)(nil)

// The game entrypoint.
type Game struct {
	cpu          *core.CPU
	renderBuffer []uint8
}

func NewGame(cpu *core.CPU) *Game {
	return &Game{cpu: cpu}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if err := g.dumpTiles(); err != nil {
			return err
		}

		return ebiten.Termination
	}

	cycles := 0

	for {
		// Tick until the screen should be rendered.
		r, err := g.cpu.Tick()
		if err != nil {
			return err
		}

		cycles += int(r.Cycles)

		if r.Render {
			break
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderBuffer = g.cpu.Render(g.renderBuffer)
	// Blit the buffer onto the screen directly.
	screen.WritePixels(g.renderBuffer)
}

func (g *Game) Layout(ow, oh int) (sw, sh int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) dumpTiles() error {
	log.Println("dumping tiles to file")

	dbg := core.NewDebugger(g.cpu).PPU()
	img := dbg.DumpTiles()

	f, err := os.Create("./tiles.png")
	if err != nil {
		return err
	}

	if err := png.Encode(f, img); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
