package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const delimiter = " "

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validLines := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, delimiter)
		data := convertData(parts)
		if dataIsValid(data) {
			validLines++
		}
	}

	fmt.Println(validLines)
}

// convertData takes an array of strings and returns an array of []float64s
// we assume the input data is clean and everything can be cast properly
func convertData(data []string) []float64 {
	floats := make([]float64, len(data))

	for i, part := range data {
		value, _ := strconv.ParseFloat(part, 64)
		floats[i] = value
	}

	return floats
}

// dataIsValid verifies that for the data in an array of ints
// meets the following criteria:
// - The numbers are either all increasing or decreasing
// - Any two adjacent levels differ by at least one and at most three
func dataIsValid(data []float64) bool {
	const treshold = 3

	// Determine the trend: increasing or decreasing
	increasing := data[1] > data[0]
	for i := 1; i < len(data); i++ {
		diff := data[i] - data[i-1]

		// Check if the difference is within the allowed range using math.Abs
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return false
		}

		// Check if the numbers are consistently increasing or decreasing
		if (increasing && diff < 0) || (!increasing && diff > 0) {
			return false
		}
	}

	return true
}
