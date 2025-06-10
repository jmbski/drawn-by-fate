package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct{}
type Combatant struct{}
type Health struct {
	MaxHealth     int
	CurrentHealth int
}

type Position struct {
	X int
	Y int
}

func (p *Position) IsEqual(other *Position) bool {
	return (p.X == other.X && p.Y == other.Y)
}

func (p *Position) GetManhattanDistance(other *Position) int {
	xDist := math.Abs(float64(p.X - other.X))
	yDist := math.Abs(float64(p.Y - other.Y))
	return int(xDist) + int(yDist)
}

type Renderable struct {
	Image *ebiten.Image
}

type Movable struct{}

type Name struct {
	Label string
}
type Monster struct {
}

type MeleeWeapon struct {
	Name       string
	MinDamage  int
	MaxDamage  int
	ToHitBonus int
}

type Armor struct {
	Name       string
	Defense    int
	ArmorClass int
}

type UserMessage struct {
	AttackMessage    string
	DeadMessage      string
	GameStateMessage string
}

/* type EcsQuery struct {
	Results *ecs.QueryResult
}

func (q *EcsQuery) GetHealth() *Health {
	return q.Results.Components[health].(*Health)
}

func (q *EcsQuery) GetArmor() *Armor {
	return q.Results.Components[armor].(*Armor)
}

func (q *EcsQuery) GetWeapon() *MeleeWeapon {
	return q.Results.Components[meleeWeapon].(*MeleeWeapon)
}

func GetHealth(p *ecs.QueryResult) *Health {
	return p.Components[health].(*Health)
}

func GetComponent[T any](q *ecs.QueryResult, c *ecs.Component) T {
	return q.Components[c].(T)
}

// Probably won't use, but will keep for educational purposes
func GetComponent2[T any](q *ecs.QueryResult, c *ecs.Component, t ...T) T {
	return q.Components[c].(T)
} */
