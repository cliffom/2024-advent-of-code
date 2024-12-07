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

const startingRune = "^"

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

	fmt.Printf("The guard visited %v distinct positions.\n", guard.Map.DistinctPositionsVisited())

	loopCounter := 0
	for i := 0; i <= areaMap.Dimensions()[0]; i++ {
		for j := 0; j <= areaMap.Dimensions()[1]; j++ {
			areaMap2, startPos2 := getAreaMapFromInput("input.txt")
			areaMap2.SetContentsAtPosition([2]int{i, j}, rune(int('#')))
			guard2 := Guard{
				CurrentDirection: Up,
				CurrentPosition:  startPos2,
				Map:              areaMap2,
				Positions:        make(map[Position]bool),
			}

			for guard2.InMapArea() {
				if !guard2.Move() {
					loopCounter += 1
					break
				}
			}
		}
	}

	fmt.Printf("%v obstacles can be placed to create a loop.\n", loopCounter)
}
