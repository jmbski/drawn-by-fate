package main

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var position *ecs.Component
var renderable *ecs.Component
var monster *ecs.Component
var health *ecs.Component
var meleeWeapon *ecs.Component
var armor *ecs.Component
var name *ecs.Component
var userMessage *ecs.Component

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	skellyImg, _, err := ebitenutil.NewImageFromFile("assets/skelly.png")
	if err != nil {
		log.Fatal(err)
	}

	orcImg, _, err := ebitenutil.NewImageFromFile("assets/orc.png")
	if err != nil {
		log.Fatal(err)
	}

	// Get first Room
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()
	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()
	monster = manager.NewComponent()
	health = manager.NewComponent()
	meleeWeapon = manager.NewComponent()
	armor = manager.NewComponent()
	name = manager.NewComponent()
	userMessage = manager.NewComponent()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			Image: playerImg,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
		}).
		AddComponent(health, &Health{
			MaxHealth:     30,
			CurrentHealth: 30,
		}).
		AddComponent(meleeWeapon, &MeleeWeapon{
			Name:       "Battle Axe",
			MinDamage:  10,
			MaxDamage:  20,
			ToHitBonus: 4,
		}).
		AddComponent(armor, &Armor{
			Name:       "Plate Armor",
			Defense:    4,
			ArmorClass: 10,
		}).
		AddComponent(name, &Name{Label: "Player"}).
		AddComponent(userMessage, &UserMessage{
			AttackMessage:    "",
			DeadMessage:      "",
			GameStateMessage: "",
		})

	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()

			// Flip a coin to see what to add
			mobSpawn := GetDiceRoll(2)

			if mobSpawn == 1 {
				manager.NewEntity().
					AddComponent(monster, &Monster{}).
					AddComponent(renderable, &Renderable{Image: orcImg}).
					AddComponent(position, &Position{X: mX, Y: mY}).
					AddComponent(health, &Health{MaxHealth: 30, CurrentHealth: 30}).
					AddComponent(meleeWeapon, &MeleeWeapon{
						Name:       "Machete",
						MinDamage:  4,
						MaxDamage:  8,
						ToHitBonus: 1,
					}).
					AddComponent(armor, &Armor{Name: "Leather", Defense: 5, ArmorClass: 6}).
					AddComponent(name, &Name{Label: "Orc"}).
					AddComponent(userMessage, &UserMessage{
						AttackMessage:    "",
						DeadMessage:      "",
						GameStateMessage: "",
					})
			} else {
				manager.NewEntity().
					AddComponent(monster, &Monster{}).
					AddComponent(renderable, &Renderable{Image: skellyImg}).
					AddComponent(position, &Position{X: mX, Y: mY}).
					AddComponent(health, &Health{MaxHealth: 10, CurrentHealth: 10}).
					AddComponent(meleeWeapon, &MeleeWeapon{
						Name:       "Short Sword",
						MinDamage:  2,
						MaxDamage:  6,
						ToHitBonus: 0,
					}).
					AddComponent(armor, &Armor{Name: "Bone", Defense: 3, ArmorClass: 4}).
					AddComponent(name, &Name{Label: "Skeleton"}).
					AddComponent(userMessage, &UserMessage{
						AttackMessage:    "",
						DeadMessage:      "",
						GameStateMessage: "",
					})

			}
		}
	}

	players := ecs.BuildTag(player, position, health, meleeWeapon, armor, name, userMessage)
	tags["players"] = players

	monsters := ecs.BuildTag(monster, position, health, meleeWeapon, armor, name, userMessage)
	tags["monsters"] = monsters

	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	messengers := ecs.BuildTag(userMessage)
	tags["messengers"] = messengers

	return manager, tags
}
