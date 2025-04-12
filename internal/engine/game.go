package engine

import (
	"drawn-by-fate/internal/config"
	"drawn-by-fate/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

var currentScene = scenes.CurrentScene()
var cfg = config.CurrentSettings

type Game struct{}

func (g *Game) Update() error {
	return currentScene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	currentScene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return cfg.ScreenWidth, cfg.ScreenHeight
}
