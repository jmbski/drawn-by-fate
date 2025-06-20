package components

import "github.com/mlange-42/ark/ecs"

type BaseMovementSpeed struct {
	Value float64
}

// TODO: Might not be needed
// type Movable struct{}
type Movable = ecs.Map3[BaseMovementSpeed, Position, Velocity]

type AiMovePattern string

// AI Movement Patterns
const (
	AiMvRandom    AiMovePattern = "ai_random"
	AiMvToNearest AiMovePattern = "ai_to_nearest"
)
