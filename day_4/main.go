package main

import (
	"bufio"
	"fmt"
	"os"
)

// directions is a 2D array of x,y coordinates
var directions = [][2]int{
	{0, 1},   // Right
	{1, 0},   // Down
	{0, -1},  // Left
	{-1, 0},  // Up
	{1, 1},   // Diagonal Down-Right
	{-1, -1}, // Diagonal Up-Left
	{1, -1},  // Diagonal Down-Left
	{-1, 1},  // Diagonal Up-Right
}

func readGridFromFile(filename string) [][]rune {
	file, _ := os.Open(filename)
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	return grid
}

func countWordOccurrencesInGrid(word string, grid [][]rune) int {
	wordRunes := []rune(word)
	wordLen := len(wordRunes)
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for _, dir := range directions {
				x, y := dir[0], dir[1]
				found := true

				for i := 0; i < wordLen; i++ {
					next_row, next_col := row+x*i, col+y*i

					if !withinBounds(rows, cols, next_row, next_col) ||
						grid[next_row][next_col] != wordRunes[i] {

						found = false
						break
					}
				}

				if found {
					count++
				}
			}
		}
	}

	return count
}

func withinBounds(rows, cols, next_row, next_col int) bool {
	if next_row < 0 || next_col < 0 || next_row >= rows || next_col >= cols {
		return false
	}

	return true
}

func countOccurrencesInXPattern(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if string(grid[r][c]) == "A" {
				if diagLeft(r, c, grid) && diagRight(r, c, grid) {
					count += 1
				}
			}

		}
	}

	return count
}

func diagLeft(x, y int, grid [][]rune) bool {
	upper_right := string(grid[x-1][y+1])
	lower_left := string(grid[x+1][y-1])
	if (upper_right == "M" && lower_left == "S") ||
		(upper_right == "S" && lower_left == "M") {

		return true
	}

	return false
}

func diagRight(x, y int, grid [][]rune) bool {
	upper_left := string(grid[x-1][y-1])
	lower_right := string(grid[x+1][y+1])
	if (upper_left == "M" && lower_right == "S") ||
		(upper_left == "S" && lower_right == "M") {

		return true
	}

	return false
}

func main() {
	filename := "input.txt"
	grid := readGridFromFile(filename)

	word := "XMAS"
	result := countWordOccurrencesInGrid(word, grid)
	result2 := countOccurrencesInXPattern(grid)

	fmt.Printf("The word '%s' appears %d times in the grid.\n", word, result)
	fmt.Printf("The word '%s' appears in an X-pattern %d times in the grid.\n", "MAS", result2)
}
