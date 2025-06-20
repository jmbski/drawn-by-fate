package systems

import (
	cmp "drawn-by-fate/internal/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
)

type RenderSystem struct {
	filter *ecs.Filter1[cmp.Renderable]
}

func (s *RenderSystem) Draw(screen *ebiten.Image, world *ecs.World) {
	query := s.filter.Query()
	for query.Next() {
		renderable := query.Get()

		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(renderable.X, renderable.Y)

		screen.DrawImage(renderable.Image, op)

	}
}

func NewRenderSystem(world *ecs.World) *RenderSystem {
	return &RenderSystem{
		filter: ecs.NewFilter1[cmp.Renderable](world).Register(),
	}
}

