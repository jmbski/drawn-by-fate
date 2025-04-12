package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X, Y  float32
	Speed float32
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += p.Speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	//result := (&colorMatrix).Apply(colorWhite)
	//ebitenutil.DrawRect(screen, p.X, p.Y, 16, 16, result)
	vector.DrawFilledRect(screen, p.X, p.Y, 16, 16, color.White, false)
}
