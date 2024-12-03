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
	canBeValidLines := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, delimiter)
		data := convertData(parts)
		if dataIsValid(data) {
			validLines++
		} else if dataCanBeValid(data) {
			canBeValidLines++
		}
	}

	fmt.Printf("Total valid lines: %v\n", validLines)
	fmt.Printf("Total valid and can be valid lines: %v\n", validLines+canBeValidLines)
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

// dataIsValid verifies that the data in an array of numbers
// meets the following criteria:
// - The numbers are either all increasing or decreasing
// - Any two adjacent levels differ by at least one and at most three
func dataIsValid(data []float64) bool {
	const treshold = 3

	// Determine the trend: increasing or decreasing
	increasing := data[1] > data[0]
	for i := 1; i < len(data); i++ {
		diff := data[i] - data[i-1]

		// Check for a difference and that it is within the allowed range
		if diff == 0 || math.Abs(float64(diff)) > treshold {
			return false
		}

		// Check if the numbers are consistently increasing or decreasing
		if (increasing && diff < 0) || (!increasing && diff > 0) {
			return false
		}
	}

	return true
}

// dataCanBeValid checks if removing one element can make the data valid
func dataCanBeValid(data []float64) bool {
	for i := range data {
		// Make a new, temporary slice that consists of the elements
		// before and after the current element
		tempData := make([]float64, len(data)-1)
		copy(tempData, data[:i])
		copy(tempData[i:], data[i+1:])

		if dataIsValid(tempData) {
			return true
		}
	}

	return false
}
