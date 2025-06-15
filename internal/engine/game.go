package engine

import (
	"drawn-by-fate/internal/config"
	"drawn-by-fate/internal/core"
	"drawn-by-fate/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
)

var currentScene = scenes.CurrentScene()
var cfg = config.CurrentSettings

type Game struct{
	World *ecs.World
	Context *core.WorldContext
}

func (g *Game) Update() error {
	return currentScene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	currentScene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return cfg.ScreenWidth, cfg.ScreenHeight
}
