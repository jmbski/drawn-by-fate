package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func TakePlayerAction(g *Game) {
	turnTaken := false

	players := g.WorldTags["players"]
	x := 0
	y := 0

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		turnTaken = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		x = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	level := g.Map.CurrentLevel

	for _, result := range g.World.Query(players) {
		pos := result.Components[position].(*Position)
		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)

		tile := level.Tiles[index]
		/* if !tile.Blocked {
			pos.X += x
			pos.Y += y
			level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)
		} */
		if !tile.Blocked {
			level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)].Blocked = false

			pos.X += x
			pos.Y += y
			level.Tiles[index].Blocked = true
			level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)

		} else if x != 0 || y != 0 {
			if level.Tiles[index].TileType != WALL {
				monsterPosition := Position{X: pos.X + x, Y: pos.Y + y}
				AttackSystem(g, pos, &monsterPosition)
			}
		}
	}

	if x != 0 || y != 0 || turnTaken {
		g.Turn = GetNextState(g.Turn)
		g.TurnCounter = 0
	}

}
