package main

import (
	"bufio"
	"log"
	"os"
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
	log.Printf("Start position: %v", startPos)
	for _, v := range areaMap {
		log.Println(string(v))
	}
}
