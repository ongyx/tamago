package tamago

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	NoROMErr = errors.New("no (boot)rom loaded")
)

type Game struct {
	S *State
}

func NewGame() *Game {
	return &Game{NewState()}
}

func (g *Game) Update() error {
	if !(g.S.hasBoot || g.S.hasROM) {
		return NoROMErr
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.S.Update(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return renderWidth, renderHeight
}
