package scenes

import (
	"drawn-by-fate/internal/config"
	"drawn-by-fate/internal/player"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var cfg = config.CurrentSettings

type GameScene struct {
	player player.Player
}

func NewGameScene() *GameScene {
	return &GameScene{
		player: player.Player{X: float32(cfg.ScreenWidth / 2), Y: float32(cfg.ScreenHeight / 2), Speed: 2.0},
	}
}

func (s *GameScene) Update() error {
	s.player.Update()
	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black) // black background
	s.player.Draw(screen)
}

