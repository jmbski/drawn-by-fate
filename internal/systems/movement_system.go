package systems

import (
	cmp "drawn-by-fate/internal/components"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
)

// PlayerMovementSystem handles player input and updates player velocity/position.
type PlayerMovementSystem struct {
	inputHandler *InputSystem
	filter       *ecs.Filter4[cmp.Renderable, cmp.Velocity, cmp.BaseMovementSpeed, cmp.InputControlled]
	LatestSpeed  float64
}

// NewPlayerMovementSystem creates a new PlayerMovementSystem.
func NewPlayerMovementSystem(inputSys *InputSystem, world *ecs.World) *PlayerMovementSystem {
	return &PlayerMovementSystem{
		inputHandler: inputSys,
		filter:       ecs.NewFilter4[cmp.Renderable, cmp.Velocity, cmp.BaseMovementSpeed, cmp.InputControlled](world).Register(),
	}
}

// Update processes player movement for entities with Position, Velocity, and InputControlled components.
func (s *PlayerMovementSystem) Update() {
	query := s.filter.Query()
	for query.Next() {
		pos, vel, speed, _ := query.Get() // Get components for the current entity

		// Reset velocity for this tick
		vel.X = 0
		vel.Y = 0

		dx, dy := 0.0, 0.0
		if s.inputHandler.PlayerInput.ActionIsPressed(ActionMoveLeft) {
			dx -= 1
		}
		if s.inputHandler.PlayerInput.ActionIsPressed(ActionMoveRight) {
			dx += 1
		}
		if s.inputHandler.PlayerInput.ActionIsPressed(ActionMoveUp) {
			dy -= 1
		}
		if s.inputHandler.PlayerInput.ActionIsPressed(ActionMoveDown) {
			dy += 1
		}

		s.LatestSpeed = speed.Value
		mag := math.Hypot(dx, dy)
		if mag != 0 {
			dx = (dx / mag) * speed.Value
			dy = (dy / mag) * speed.Value
		}

		vel.X = dx
		vel.Y = dy
		// Update position based on velocity.
		// Ebiten's Update is typically 60 TPS, so delta time is 1/60.
		// For more precise delta time, consider passing it or using a resource in Ark-tools.
		pos.X += vel.X / float64(ebiten.TPS()) // Divide by TPS for per-second speed
		pos.Y += vel.Y / float64(ebiten.TPS())
	}
}

// TODO: Move this somewhere better
type GameSystem interface {
	Update()
}

type SystemMap map[string]GameSystem

func (s SystemMap) UpdateAll() {
	for _, sys := range s {
		sys.Update()
	}
}
