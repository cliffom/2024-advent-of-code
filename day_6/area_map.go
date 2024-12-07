package main

import "fmt"

type AreaMap struct {
	Contents [][]rune
}

func (m *AreaMap) Dimensions() [2]int {
	return [2]int{len(m.Contents) - 1, len(m.Contents[0]) - 1}
}

func (m *AreaMap) Draw() {
	for _, j := range m.Contents {
		for _, k := range j {
			fmt.Printf("%v", string(k))
		}
		fmt.Printf("\n")
	}
}

func (m *AreaMap) PositionIsOutOfBounds(postion [2]int) bool {
	x := postion[0]
	y := postion[1]

	if x < 0 || y < 0 || x > m.Dimensions()[0] || y > m.Dimensions()[1] {
		return true
	}

	return false
}

func (m *AreaMap) PositionIsOnBorder(position [2]int) bool {
	dimensions := m.Dimensions()
	x := position[0]
	y := position[1]
	width := dimensions[0]
	height := dimensions[1]

	return x == 0 || y == 0 || x == width || y == height

}

func (m *AreaMap) PositionIsOccupied(position [2]int) bool {
	return m.ContentsAtPosition(position) == rune(int('#'))
}

func (m *AreaMap) ContentsAtPosition(position [2]int) rune {
	x := position[0]
	y := position[1]

	return m.Contents[x][y]
}

func (m *AreaMap) SetContentsAtPosition(position [2]int, value rune) {
	x := position[0]
	y := position[1]
	m.Contents[x][y] = value
}

func (m *AreaMap) MarkPositionVisited(position [2]int) {
	x := position[0]
	y := position[1]

	m.SetContentsAtPosition([2]int{x, y}, rune(int('X')))
}

func (m *AreaMap) DistinctPositionsVisited() int {
	count := 0
	for _, rows := range m.Contents {
		for _, cols := range rows {
			if string(cols) == "X" {
				count += 1
			}
		}
	}

	return count
}
