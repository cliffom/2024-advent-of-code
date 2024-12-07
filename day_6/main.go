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
	for scanner.Scan() {
		line := scanner.Text()
		runes := make([]rune, 0)
		areaMap = append(areaMap, runes)
		areaMapSize := len(areaMap) - 1
		for i, char := range line {
			areaMap[areaMapSize] = append(areaMap[areaMapSize], char)
			if string(char) == startingRune {
				startPos = [2]int{areaMapSize, i}
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

	guard.Map.Draw()
	fmt.Printf("The guard visited %v distinct positions.\n", guard.Map.DistinctPositionsVisited())
}
