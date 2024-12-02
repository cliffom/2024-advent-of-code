package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const delimiter = "   "

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var col1, col2 []float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, delimiter)

		num1, _ := strconv.ParseFloat(parts[0], 64)
		num2, _ := strconv.ParseFloat(parts[1], 64)

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	slices.Sort(col1)
	slices.Sort(col2)

	sum := sumColumns(col1, col2)

	fmt.Printf("%.0f\n", sum)
}

func sumColumns(column1, column2 []float64) float64 {
	var sum float64
	for i := range column1 {
		sum += math.Abs(column1[i] - column2[i])
	}
	return sum
}
