package core

import (
	cmp "drawn-by-fate/internal/components"

	"github.com/mlange-42/ark/ecs"
)

type PlayerEntity = *ecs.Map[cmp.Player]

type WorldContext struct {
	Renderables *ecs.Filter1[cmp.Renderable]
	Players     *ecs.Filter1[cmp.Player]
	Movable     *ecs.Filter1[cmp.Movable]
}

func InitializeWorld() *ecs.World {
	world := ecs.NewWorld()

	return &world
}
