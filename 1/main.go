package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const delimiter = "   "

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var col1, col2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, delimiter)

		num1, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		num2, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	sort.Ints(col1)
	sort.Ints(col2)
	sum := sumColumns(col1, col2)

	fmt.Println(sum)
}

func sumColumns(column1, column2 []int) int {
	sum := 0
	for i := range column1 {
		diff := column1[i] - column2[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return sum
}
