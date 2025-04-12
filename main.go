package main

import (
	"drawn-by-fate/internal/config"
	"drawn-by-fate/internal/engine"
	"drawn-by-fate/internal/scenes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var cfg = config.CurrentSettings

func main() {
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle("Drawn By Fate - Vertical Slice")

	scenes.SwitchScene(*scenes.NewGameScene())

	if err := ebiten.RunGame(&engine.Game{}); err != nil {
		log.Fatal(err)
	}
}
