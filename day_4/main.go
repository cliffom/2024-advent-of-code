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

func countWord(grid [][]rune, word string) int {
	wordRunes := []rune(word)
	wordLen := len(wordRunes)
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// Check in all directions
			for _, dir := range directions {
				dr, dc := dir[0], dir[1]
				found := true

				for i := 0; i < wordLen; i++ {
					nr, nc := r+dr*i, c+dc*i
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != wordRunes[i] {
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

// Function to count the "MAS" X-pattern in the grid
func countXPattern(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if string(grid[r][c]) == "A" {
				// Check corners for S an M
				upper_left := string(grid[r-1][c-1])
				upper_right := string(grid[r-1][c+1])
				lower_left := string(grid[r+1][c-1])
				lower_right := string(grid[r+1][c+1])
				if (upper_left == "M" && lower_right == "S") && (upper_right == "M" && lower_left == "S") ||
					(upper_left == "S" && lower_right == "M") && (upper_right == "S" && lower_left == "M") ||
					(upper_left == "M" && lower_right == "S") && (upper_right == "S" && lower_left == "M") ||
					(upper_left == "S" && lower_right == "M") && (upper_right == "M" && lower_left == "S") {
					count++
				}
			}

		}
	}

	return count
}

func main() {
	filename := "input.txt"
	grid := readGridFromFile(filename)

	word := "XMAS"
	result := countWord(grid, word)
	result2 := countXPattern(grid)
	fmt.Printf("The word '%s' appears %d times in the grid.\n", word, result)
	fmt.Printf("The word '%s' appears in an X-pattern %d times in the grid.\n", "MAS", result2)
}
