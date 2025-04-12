package main

import (
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var position *ecs.Component
var renderable *ecs.Component
var monster *ecs.Component

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

	// Get first Room
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()
	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()
	monster = manager.NewComponent()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			Image: playerImg,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
		})

	for _, room := range startingLevel.Rooms {
		if room.X1 != startingRoom.X1 {
			mX, mY := room.Center()
			manager.NewEntity().
				AddComponent(monster, &Monster{Name: "Skeleton"}).
				AddComponent(renderable, &Renderable{Image: skellyImg}).
				AddComponent(position, &Position{X: mX, Y: mY})
		}
	}

	players := ecs.BuildTag(player, position)
	tags["players"] = players

	monsters := ecs.BuildTag(monster, position)
	tags["monsters"] = monsters

	renderables := ecs.BuildTag(renderable, position)
	tags["renderables"] = renderables

	return manager, tags
}
