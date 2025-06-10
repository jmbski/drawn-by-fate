package main

import (
	"github.com/mlange-42/ark/ecs"
)

type MobMap = ecs.Map8[Monster, Renderable, Position, Health, MeleeWeapon, Armor, Name, UserMessage]
type PlayerMap = ecs.Map9[Player, Renderable, Movable, Position, Health, MeleeWeapon, Armor, Name, UserMessage]

type WorldContext struct {
	PlayerPositions *ecs.Filter2[Player, Position]
	PlayerHuds      *ecs.Filter4[Player, Health, Armor, MeleeWeapon]
	Combatants      *ecs.Filter6[Position, Armor, Health, Name, MeleeWeapon, UserMessage]
	MobPosHps       *ecs.Filter3[Monster, Position, Health]
	Renderables     *ecs.Filter2[Renderable, Position]
	Messengers      *ecs.Filter1[UserMessage]
	Players         *ecs.Filter1[Player]
}

func NewWorldContext(world *ecs.World) *WorldContext {
	return &WorldContext{
		PlayerPositions: ecs.NewFilter2[Player, Position](world).Register(),
		PlayerHuds:      ecs.NewFilter4[Player, Health, Armor, MeleeWeapon](world).Register(),
		Combatants:      ecs.NewFilter6[Position, Armor, Health, Name, MeleeWeapon, UserMessage](world).Register(),
		MobPosHps:       ecs.NewFilter3[Monster, Position, Health](world).Register(),
		Renderables:     ecs.NewFilter2[Renderable, Position](world).Register(),
		Messengers:      ecs.NewFilter1[UserMessage](world).Register(),
		Players:         ecs.NewFilter1[Player](world).Register(),
	}
}

func GetMobs(world *ecs.World) *MobMap {
	return ecs.NewMap8[Monster, Renderable, Position, Health, MeleeWeapon, Armor, Name, UserMessage](world)
}

func GetPlayers(world *ecs.World) *PlayerMap {
	return ecs.NewMap9[Player, Renderable, Movable, Position, Health, MeleeWeapon, Armor, Name, UserMessage](world)
}

func InitializeWorld(startingLevel Level) (*ecs.World, *WorldContext) {
	world := ecs.NewWorld()

	playerImg := GetImage("player.png")
	skellyImg := GetImage("skelly.png")
	orcImg := GetImage("orc.png")

	// Get first Room
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	players := GetPlayers(&world)
	monsters := GetMobs(&world)

	players.NewEntity(
		&Player{},
		&Renderable{
			Image: playerImg,
		},
		&Movable{},
		&Position{X: x, Y: y},
		&Health{MaxHealth: 30, CurrentHealth: 30},
		&MeleeWeapon{
			Name:       "Battle Axe",
			MinDamage:  10,
			MaxDamage:  20,
			ToHitBonus: 4,
		},
		&Armor{
			Name:       "Plate Armor",
			Defense:    4,
			ArmorClass: 10,
		},
		&Name{Label: "Player"},
		&UserMessage{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		},
	)

	for _, room := range startingLevel.Rooms {
		if room.X1 == startingRoom.X1 {
			continue
		}

		mX, mY := room.Center()

		mobSpawn := GetDiceRoll(2)

		if mobSpawn == 1 {
			monsters.NewEntity(
				&Monster{},
				&Renderable{Image: orcImg},
				&Position{X: mX, Y: mY},
				&Health{MaxHealth: 30, CurrentHealth: 30},
				&MeleeWeapon{
					Name:       "Machete",
					MinDamage:  4,
					MaxDamage:  8,
					ToHitBonus: 1,
				},
				&Armor{Name: "Leather", Defense: 5, ArmorClass: 6},
				&Name{Label: "Orc"},
				&UserMessage{
					AttackMessage:    "",
					DeadMessage:      "",
					GameStateMessage: "",
				},
			)
		} else {
			monsters.NewEntity(
				&Monster{},
				&Renderable{Image: skellyImg},
				&Position{X: mX, Y: mY},
				&Health{MaxHealth: 10, CurrentHealth: 10},
				&MeleeWeapon{
					Name:       "Short Sword",
					MinDamage:  2,
					MaxDamage:  6,
					ToHitBonus: 0,
				},
				&Armor{Name: "Bone", Defense: 3, ArmorClass: 4},
				&Name{Label: "Skeleton"},
				&UserMessage{
					AttackMessage:    "",
					DeadMessage:      "",
					GameStateMessage: "",
				},
			)
		}
	}

	filters := NewWorldContext(&world)
	return &world, filters
}
