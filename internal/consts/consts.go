package consts

type ActorType int
type AttackType int

// Actor Type enum
const (
	PlayerType ActorType = iota
	MonsterType
	SummonableType
	AllyType
	NPCType
)

// Attack Type enum
const (
	MeleeAttack AttackType = iota
	RangedAttack
)
