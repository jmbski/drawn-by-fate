package player

import (
	cmp "drawn-by-fate/internal/components"

	"github.com/mlange-42/ark/ecs"
)

type PlayerMap = ecs.Map8[
	cmp.Player,
	cmp.BaseMovementSpeed,
	cmp.Combatant,
	cmp.Controllable,
	cmp.Defense,
	cmp.Faction,
	cmp.Health,
	cmp.Experience,
]

type Player = ecs.Map5[
	cmp.Player,
	cmp.BaseMovementSpeed,
	cmp.Movable,
	cmp.Controllable,
	cmp.Renderable,
]

/* type Player struct {
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

func NewPlayer() *Player {
	return &Player{
		X:     1,
		Y:     1,
		Speed: 1,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	//result := (&colorMatrix).Apply(colorWhite)
	//ebitenutil.DrawRect(screen, p.X, p.Y, 16, 16, result)
	vector.DrawFilledRect(screen, p.X, p.Y, 16, 16, color.White, false)
} */
