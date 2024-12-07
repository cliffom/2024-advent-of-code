package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Up int = iota
	Right
	Down
	Left
)

const (
	startingRune = "^"
	obstacleRune = '#'
	occupiedRune = 'X'
)

func getAreaMapFromInput(filename string) (AreaMap, [2]int) {
	file, _ := os.Open(filename)
	defer file.Close()

	var areaMap [][]rune
	var startPos [2]int
	scanner := bufio.NewScanner(file)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		runes := []rune(line)
		areaMap = append(areaMap, runes)

		if startPos == [2]int{0, 0} {
			for col, char := range runes {
				if string(char) == startingRune {
					startPos = [2]int{row, col}
				}
			}
		}
	}

	return AreaMap{Contents: areaMap}, startPos
}

// checkForLoopCausingObstacles looks at the map of a guards path
// and checks to see how many obstacles can be placed to put the guard
// on a path that results in a loop
// I know this is ugly and a brute force attempt. Don't @me
func checkForLoopCausingObstacles(pathMap AreaMap) int {
	areaMap, startPos := getAreaMapFromInput("input.txt")
	loopCounter := 0

	for x := 0; x <= areaMap.Dimensions()[0]; x++ {
		for y := 0; y <= areaMap.Dimensions()[1]; y++ {
			pos := [2]int{x, y}
			cellOnPath := pathMap.ContentsAtPosition(pos) == occupiedRune

			if cellOnPath {
				currentCellContents := areaMap.ContentsAtPosition(pos)
				areaMap.SetContentsAtPosition(pos, obstacleRune)

				guard := Guard{
					CurrentDirection: Up,
					CurrentPosition:  startPos,
					Map:              areaMap,
					Positions:        make(map[Position]bool),
				}

				for guard.InMapArea() {
					if guard.CheckForLoop() {
						loopCounter += 1
						break
					}
					guard.Move()
				}

				// reset map
				areaMap.SetContentsAtPosition(pos, currentCellContents)
			}
		}
	}

	return loopCounter
}

func main() {
	areaMap, startPos := getAreaMapFromInput("input.txt")
	fmt.Printf("Guard starting position: %v\n", startPos)

	guard := Guard{
		CurrentDirection: Up,
		CurrentPosition:  startPos,
		Map:              areaMap,
		Positions:        make(map[Position]bool),
	}

	for guard.InMapArea() {
		guard.Move()
	}

	possibleLoopCausingObstacles := checkForLoopCausingObstacles(guard.Map)

	fmt.Printf("The guard visited %v distinct positions.\n", guard.Map.DistinctPositionsVisited())
	fmt.Printf("%v obstacles can be placed to create a loop.\n", possibleLoopCausingObstacles)
}
