package main

import (
	"errors"
	"reflect"
)

// Node represents a given point on a map
// g is teh total distance of the node from the start
// h is the estimated distance of the node from the ending
// f is the total value of the node (g+h)
type node struct {
	Parent   *node
	Position *Position
	g        int
	h        int
	f        int
}

func newNode(parent *node, position *Position) *node {
	return &node{
		Parent:   parent,
		Position: position,
		g:        0,
		h:        0,
		f:        0,
	}
}

func (n *node) isEqual(other *node) bool {
	return (n.Position.X == other.Position.X && n.Position.Y == other.Position.Y)
}

type AStar struct{}

// GetPath takes a level, the starting position and an ending position (the goal) and returns
// a list of Positions which is the path between the points.
func (as AStar) GetPath(level Level, start *Position, end *Position) []Position {
	gd := NewGameData()
	openList := make([]*node, 0)
	closedList := make([]*node, 0)

	// create our starting point
	startNode := newNode(nil, start)
	startNode.g = 0
	startNode.h = 0
	startNode.f = 0

	endNodePlaceholder := newNode(nil, end)

	openList = append(openList, startNode)

	for {
		if len(openList) == 0 {
			break
		}

		currentNode := openList[0]
		currentIndex := 0

		for index, item := range openList {
			if item.f < currentNode.f {
				currentNode = item
				currentIndex = index
			}
		}

		openList = append(openList[:currentIndex], openList[currentIndex+1:]...)
		closedList = append(closedList, currentNode)

		if currentNode.isEqual(endNodePlaceholder) {
			path := make([]Position, 0)
			current := currentNode
			for {
				if current == nil {
					break
				}
				path = append(path, *current.Position)
				current = current.Parent
			}

			reverseSlice(path)
			return path
		}

		edges := make([]*node, 0)

		if currentNode.Position.Y > 0 {
			tile := level.Tiles[level.GetIndexFromXY(currentNode.Position.X, currentNode.Position.Y-1)]

			if tile.TileType != WALL {
				upNodePosition := Position{
					X: currentNode.Position.X,
					Y: currentNode.Position.Y - 1,
				}

				newNode := newNode(currentNode, &upNodePosition)
				edges = append(edges, newNode)
			}
		}
	}
}

func reverseSlice(data interface{}) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		panic(errors.New("data must be a slice type"))
	}

	valueLen := value.Len()
	for i := 0; i <= int((valueLen-1)/2); i++ {
		reverseIndex := valueLen - 1 - i
		tmp := value.Index(reverseIndex).Interface()
		value.Index(reverseIndex).Set(value.Index(i))
		value.Index(i).Set(reflect.ValueOf(tmp))
	}
}
