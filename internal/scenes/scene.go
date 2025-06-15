/* Scenes are equivalent to Levels in the roguelike Ark test app

 */

package scenes

import (
	"drawn-by-fate/internal/config"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var cfg = config.CurrentSettings

// This will probably need some kind of reference to entities
type GameScene struct {
}

func NewGameScene() *GameScene {
	return &GameScene{
		//		player: player.Player{X: float64(cfg.ScreenWidth / 2), Y: float64(cfg.ScreenHeight / 2), Speed: 2.0},
	}
}

func (s *GameScene) Update() error {
	//s.player.Update()
	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black) // black background
	//s.player.Draw(screen)
}
