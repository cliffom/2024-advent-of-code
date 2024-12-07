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
	}

	for guard.InMapArea() {
		guard.Move()
	}

	//guard.Map.Draw()
	fmt.Printf("The guard visited %v distinct positions.\n", guard.Map.DistinctPositionsVisited())
}
