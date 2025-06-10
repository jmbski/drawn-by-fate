package components

import "drawn-by-fate/internal/consts"

type Offense struct {
	AttackDamage float32
	AttackSpeed  float32
	Range        float32
	Area         float32
	Type         consts.AttackType
}

type Defense struct {
	Armor   float32
	Evasion float32
}

type Health struct {
	Max     float32
	Current float32
}
