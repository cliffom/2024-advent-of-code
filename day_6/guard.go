package main

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
	Map              AreaMap
}

func (g *Guard) GetCurrentPosition() [2]int {
	return [2]int{g.CurrentPosition[0], g.CurrentPosition[1]}
}

func (g *Guard) GetNextPosition() [2]int {
	position := g.GetCurrentPosition()
	x := position[0] + coordinates[g.CurrentDirection][0]
	y := position[1] + coordinates[g.CurrentDirection][1]

	return [2]int{x, y}
}

func (g *Guard) SetPosition(position [2]int) {
	x := position[0]
	y := position[1]
	g.CurrentPosition = [2]int{x, y}
	g.Map.SetContentsAtPosition(g.CurrentPosition, guardFrames[g.CurrentDirection])
}

func (g *Guard) ChangeDirection() {
	g.CurrentDirection += 1
	if g.CurrentDirection >= len(coordinates) {
		g.CurrentDirection = 0
	}
}

func (g *Guard) Move() {
	position := g.GetCurrentPosition()
	nextPosition := g.GetNextPosition()

	g.Map.MarkPositionVisited(position)

	if g.Map.PositionIsOutOfBounds(position) {
		return
	}

	// check for obstacle
	if g.Map.PositionIsOccupied(nextPosition) {
		g.ChangeDirection()
		return
	}

	g.SetPosition(nextPosition)
}

func (g *Guard) InArea() bool {
	position := g.GetCurrentPosition()
	dimensions := g.Map.Dimensions()

	x := position[0]
	y := position[1]

	if x == dimensions[0]-1 || y == dimensions[1]-1 {
		g.Map.MarkPositionVisited(position)
		return false
	}

	return true
}
