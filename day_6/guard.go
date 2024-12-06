package main

const guardMovementDirections = 4

type Guard struct {
	CurrentDirection int
	CurrentPosition  [2]int
	Map              AreaMap
}

func (g *Guard) MovementModifier() [2]int {
	coordinates := [guardMovementDirections][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	return coordinates[g.CurrentDirection]
}

func (g *Guard) CurrentFrame() rune {
	guardFrames := [guardMovementDirections]rune{
		rune(int('^')), // Up
		rune(int('>')), // Right
		rune(int('v')), // Down
		rune(int('<')), // Left
	}

	return guardFrames[g.CurrentDirection]
}

func (g *Guard) GetCurrentPosition() [2]int {
	x, y := g.CurrentPosition[0], g.CurrentPosition[1]
	return [2]int{x, y}
}

func (g *Guard) GetNextPosition() [2]int {
	position := g.GetCurrentPosition()
	nextPosition := [2]int{position[0] + g.MovementModifier()[0], position[1] + g.MovementModifier()[1]}

	return nextPosition
}

func (g *Guard) SetPosition(position [2]int) {
	x, y := position[0], position[1]
	g.CurrentPosition = [2]int{x, y}
	g.Map.SetContentsAtPosition(g.CurrentPosition, g.CurrentFrame())
}

func (g *Guard) ChangeDirection() {
	g.CurrentDirection += 1
	if g.CurrentDirection >= guardMovementDirections {
		g.CurrentDirection = 0
	}
}

func (g *Guard) Move() {
	position := g.GetCurrentPosition()
	nextPosition := g.GetNextPosition()

	g.Map.MarkPositionVisited(position)

	if g.Map.PositionIsOutOfBounds(nextPosition) {
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
