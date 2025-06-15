package components

import "github.com/mlange-42/ark/ecs"

type BaseMovementSpeed struct {
	BaseSpeed float64
}

// TODO: Might not be needed
//type Movable struct{}
type Movable = ecs.Map3[BaseMovementSpeed,Position,Velocity]
