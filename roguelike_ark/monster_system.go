package main

import (
	"github.com/norendren/go-fov/fov"
)

func UpdateMonster(game *Game) {
	l := game.Map.CurrentLevel
	playerPosition := Position{}

	pQuery := game.Context.PlayerPositions.Query()
	mQuery := game.Context.MobPosHps.Query()

	for pQuery.Next() {
		_, pos := pQuery.Get()
		playerPosition.X = pos.X
		playerPosition.Y = pos.Y
	}

	for mQuery.Next() {
		_, pos, health := mQuery.Get()

		monsterSees := fov.New()
		monsterSees.Compute(l, pos.X, pos.Y, 8)

		if monsterSees.IsVisible(playerPosition.X, playerPosition.Y) {
			if pos.GetManhattanDistance(&playerPosition) == 1 {
				AttackSystem(game, pos, &playerPosition)
				if health.CurrentHealth <= 0 {
					t := l.Tiles[l.GetIndexFromXY(pos.X, pos.Y)]
					t.Blocked = false
				}
			} else {
				astar := AStar{}
				path := astar.GetPath(l, pos, &playerPosition)

				if len(path) > 1 {
					nextTile := l.Tiles[l.GetIndexFromXY(path[1].X, path[1].Y)]
					if !nextTile.Blocked {
						l.Tiles[l.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
						pos.X = path[1].X
						pos.Y = path[1].Y
						nextTile.Blocked = true
					}
				}
			}
		}
	}

	game.Turn = PlayerTurn
}
