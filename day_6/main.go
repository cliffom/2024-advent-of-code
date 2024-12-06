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

func getAreaMapFromInput(filename string) ([][]rune, [2]int) {
	file, _ := os.Open(filename)
	defer file.Close()

	var areaMap [][]rune
	var startPos [2]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := make([]rune, 0)
		areaMap = append(areaMap, runes)
		areaMapSize := len(areaMap)
		for i, char := range line {
			areaMap[areaMapSize-1] = append(areaMap[areaMapSize-1], char)
			if string(char) == "^" {
				startPos = [2]int{areaMapSize - 1, i}
			}
		}

	}

	return areaMap, startPos
}

func main() {
	areaMap, startPos := getAreaMapFromInput("input.txt")
	fmt.Printf("Guard starting position: %v\n", startPos)

	guard := Guard{
		CurrentDirection: Up,
		CurrentPosition:  startPos,
		AreaMap:          areaMap,
	}

	for !guard.ExitedArea() {
		guard.Move()
	}

	distinctGuardPositions := 0
	for _, rows := range guard.AreaMap {
		for _, cols := range rows {
			if string(cols) == "X" {
				distinctGuardPositions += 1
			}
		}
	}

	fmt.Printf("The guard visited %v distinct positions.\n", distinctGuardPositions)
}
