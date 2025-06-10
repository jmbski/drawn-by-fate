package components

import "drawn-by-fate/internal/consts"

// Actor Type Tag Component structs
type Player struct{}
type Monster struct{}
type Summonable struct{}
type Ally struct{}
type NPC struct{}

// Categorizing/Functional Tag Component structs
type Combatant struct{}
type HasAi struct{}
type Controllable struct{}
type Faction struct{}

// Meta Info/Context/Display Component structs
type ActorInfo struct {
	Type consts.ActorType
}
