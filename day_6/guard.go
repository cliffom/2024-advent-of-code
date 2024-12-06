package main

import (
	"fmt"
)

var coordinates = [4][2]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

var guardFrames = [4]rune{
	rune(int('^')), // Up
	rune(int('>')), // Right
	rune(int('v')), // Down
	rune(int('<')), // Left
}

type Guard struct {
	CurrentDirection int
	CurrentPosition  [2]int
	AreaMap          [][]rune
}

func (g *Guard) GetCurrentPosition() (int, int) {
	return g.CurrentPosition[0], g.CurrentPosition[1]
}

func (g *Guard) GetNextPosition() (int, int) {
	posX, posY := g.GetCurrentPosition()
	nextX := posX + coordinates[g.CurrentDirection][0]
	nextY := posY + coordinates[g.CurrentDirection][1]

	return nextX, nextY
}

func (g *Guard) SetPosition(x, y int) {
	g.CurrentPosition = [2]int{x, y}
	g.AreaMap[x][y] = guardFrames[g.CurrentDirection]
}

func (g *Guard) MarkPositionVisited() {
	x, y := g.GetCurrentPosition()
	g.AreaMap[x][y] = rune(int('X'))
}

func (g *Guard) GetMapSize() (int, int) {
	width := len(g.AreaMap)
	height := len(g.AreaMap[0])

	return width, height
}

func (g *Guard) DrawMap() {
	for _, j := range g.AreaMap {
		for _, k := range j {
			fmt.Printf("%v", string(k))
		}
		fmt.Printf("\n")
	}
}

func (g *Guard) MapPositionIsOccupied(mapX, mapY int) bool {
	return g.AreaMap[mapX][mapY] == rune(int('#'))
}

func (g *Guard) ChangeDirection() {
	g.CurrentDirection += 1
	if g.CurrentDirection >= len(coordinates) {
		g.CurrentDirection = 0
	}
}

func (g *Guard) Move() {
	nextX, nextY := g.GetNextPosition()
	mapWidth, mapHeight := g.GetMapSize()

	g.MarkPositionVisited()

	// check inner bounds
	if nextX < 0 || nextY < 0 {
		return
	}

	// check  outer bounds
	if nextX >= mapWidth || nextY >= mapHeight {
		return
	}

	// check for obstacle
	if g.MapPositionIsOccupied(nextX, nextY) {
		g.ChangeDirection()
		return
	}

	g.SetPosition(nextX, nextY)
}

func (g *Guard) ExitedArea() bool {
	posX := g.CurrentPosition[0]
	posY := g.CurrentPosition[1]

	if posX == len(g.AreaMap)-1 || posY == len((g.AreaMap[0]))-1 {
		g.MarkPositionVisited()
		return true
	}

	return false
}
