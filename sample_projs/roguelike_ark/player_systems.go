package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func TakePlayerAction(g *Game) {
	turnTaken := false

	query := g.Context.Players.Query()
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

	for query.Next() {

		entity := query.Entity()

		player := g.PlayerInstance(entity)
		if player == nil {
			continue
		}

		//pos := result.Components[position].(*Position)
		index := level.GetIndexFromXY(player.Pos.X+x, player.Pos.Y+y)

		tile := level.Tiles[index]
		/* if !tile.Blocked {
			pos.X += x
			pos.Y += y
			level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)
		} */
		if !tile.Blocked {
			level.Tiles[level.GetIndexFromXY(player.Pos.X, player.Pos.Y)].Blocked = false

			player.Pos.X += x
			player.Pos.Y += y
			level.Tiles[index].Blocked = true
			level.PlayerVisible.Compute(level, player.Pos.X, player.Pos.Y, 8)

		} else if x != 0 || y != 0 {
			if level.Tiles[index].TileType != WALL {
				monsterPosition := Position{X: player.Pos.X + x, Y: player.Pos.Y + y}
				AttackSystem(g, player.Pos, &monsterPosition)
			}
		}
	}

	if x != 0 || y != 0 || turnTaken {
		g.Turn = GetNextState(g.Turn)
		g.TurnCounter = 0
	}

}
