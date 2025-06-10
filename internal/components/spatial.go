package components

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
