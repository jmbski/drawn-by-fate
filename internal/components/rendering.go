package components

import "github.com/hajimehoshi/ebiten/v2"

type Renderable struct {
	Image *ebiten.Image
	X     float64 // exact X coordinate position
	Y     float64 // exact Y coordinate position
}

func (r *Renderable) Value() (x, y *float64) {
	return &r.X, &r.Y
}

func NewRenderable(image *ebiten.Image, x, y float64) *Renderable {
	return &Renderable{
		Image: image,
		X:     x,
		Y:     y,
	}
}

// 2D position coordinate
type Position struct {
	X float64
	Y float64
}

func (p *Position) Value() (x, y *float64) {
	return &p.X, &p.Y
}

func NewPosition(x, y float64) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

// Velocity of movement
type Velocity struct {
	X float64
	Y float64
}

func (p *Velocity) Value() (x, y *float64) {
	return &p.X, &p.Y
}

func NewVelocity(x, y float64) *Velocity {
	return &Velocity{
		X: x,
		Y: y,
	}
}
