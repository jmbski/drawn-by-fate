package core

import (
	cmp "drawn-by-fate/internal/components"
	"drawn-by-fate/internal/loaders"
	"drawn-by-fate/pkg/config"
	"fmt"
	"math/rand"

	"github.com/mlange-42/ark/ecs"
)

type PlayerEntity = *ecs.Map[cmp.Player]

type WorldContext struct {
	Renderables *ecs.Filter1[cmp.Renderable]
	Players     *ecs.Filter1[cmp.Player]
	Movable     *ecs.Filter1[cmp.Movable]
}

func NewWorldContext(world *ecs.World) *WorldContext {
	return &WorldContext{
		Renderables: ecs.NewFilter1[cmp.Renderable](world).Register(),
		Players:     ecs.NewFilter1[cmp.Player](world).Register(),
		Movable:     ecs.NewFilter1[cmp.Movable](world).Register(),
	}
}

func InitializeWorld() *ecs.World {
	world := ecs.NewWorld()
	assets := loaders.NewAssetLoader()

	// TODO: move Players init to a better location
	playerImg, err := assets.GetImageAsset("sprites/player.png")
	if err != nil {
		// TODO: better error logging
		panic(err)
	}

	mobImg, err := assets.GetImageAsset("sprites/orc.png")
	if err != nil {
		// TODO: error handling
		panic(err)
	}

	players := ecs.NewMap5[cmp.Player, cmp.Renderable, cmp.Velocity, cmp.BaseMovementSpeed, cmp.InputControlled](&world)
	w, h := config.GetScreenSize()
	players.NewEntity(
		&cmp.Player{},
		&cmp.Renderable{
			Image: playerImg,
			X:     float64(w) / 2,
			Y:     float64(h) / 2,
		},
		&cmp.Velocity{},
		&cmp.BaseMovementSpeed{Value: 120},
		&cmp.InputControlled{},
	)

	mobs := ecs.NewMap5[cmp.Monster, cmp.Renderable, cmp.Velocity, cmp.BaseMovementSpeed, cmp.AiControlled](&world)
	for range 10 {
		x := rand.Float64() * float64(w)
		y := rand.ExpFloat64() * float64(h)
		fmt.Println(x, y)
		mobs.NewEntity(
			&cmp.Monster{},
			&cmp.Renderable{
				Image: mobImg,
				X:     x,
				Y:     y,
			},
			&cmp.Velocity{},
			&cmp.BaseMovementSpeed{Value: 150},
			&cmp.AiControlled{
				MovePattern: cmp.AiMvToNearest,
			},
		)
	}

	return &world
}
