package main

const (
	guardMovementDirections = 4
	upRune                  = '^'
	rightRune               = '>'
	downRune                = 'v'
	leftRune                = '<'
)

type Position struct {
	X, Y      int
	Direction int // "up", "right", "down", "left"
}

type Guard struct {
	CurrentDirection int
	CurrentPosition  [2]int
	Map              AreaMap
	Positions        map[Position]bool
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
		upRune,
		rightRune,
		downRune,
		leftRune,
	}

	return guardFrames[g.CurrentDirection]
}

func (g *Guard) GetCurrentPosition() [2]int {
	x, y := g.CurrentPosition[0], g.CurrentPosition[1]
	return [2]int{x, y}
}

func (g *Guard) GetNextPosition() [2]int {
	x, y := g.CurrentPosition[0], g.CurrentPosition[1]
	dX, dY := g.MovementModifier()[0], g.MovementModifier()[1]
	nextPosition := [2]int{x + dX, y + dY}

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

	if g.Map.PositionIsOutOfBounds(nextPosition) {
		return
	}

	// check for obstacle
	if g.Map.PositionIsOccupied(nextPosition) {
		g.ChangeDirection()
		return
	}

	g.MarkPositionVisited(position)
	g.SetPosition(nextPosition)
}

func (g *Guard) MarkPositionVisited(position [2]int) {
	pos := Position{
		X:         position[0],
		Y:         position[1],
		Direction: g.CurrentDirection,
	}

	g.Positions[pos] = true
	g.Map.MarkPositionVisited(position)
}

// InMapArea checks to see if a guard made it to the edge of a map
// We only concern about the outer edges of a map and not beyond
// since if a guard makes it to the edge of a map, they will make it
// to the next area
func (g *Guard) InMapArea() bool {
	position := g.GetCurrentPosition()

	if g.Map.PositionIsOnBorder(position) {
		g.Map.MarkPositionVisited(position)
		return false
	}

	return true
}

func (g *Guard) CheckForLoop() bool {
	position := g.GetCurrentPosition()
	if (g.Positions[Position{X: position[0], Y: position[1], Direction: g.CurrentDirection}]) {
		return true
	}

	return false
}
