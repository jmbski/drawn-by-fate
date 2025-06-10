package core

import (
	cmp "drawn-by-fate/internal/components"

	"github.com/mlange-42/ark/ecs"
)

type PlayerEntity = *ecs.Map[cmp.Player]

type WorldContext struct {
	
}

func InitializeWorld() *ecs.World {
	world := ecs.NewWorld()

	return &world
}
