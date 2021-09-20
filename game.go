package tamago

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	C *CPU
}

func NewGame() *Game {
	return &Game{NewCPU()}
}

func (g *Game) Update() error {
	if !(g.C.hasBoot || g.C.hasROM) {
		return NoROMErr
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.C.step()
	g.C.render.fb.CopyInto(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return renderWidth, renderHeight
}
