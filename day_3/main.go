package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// We want to match input items that look like
// mul(x,y) where x and 7 are numbers.
// Ex: mul(2,3)
var re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	output := 0
	for scanner.Scan() {
		line := scanner.Text()
		instructions := re.FindAllStringSubmatch(line, -1)
		output += sumOfValidInstructions(instructions)
	}

	log.Printf("Sum of valid instructions: %v", output)
}

// sumOfValidInstructions iterates through all valid instructions
// and sums the value of the product of each pair
func sumOfValidInstructions(instructions [][]string) int {
	sum := 0
	// each entry in matches[] will contain the string matched,
	// the first digit, and the second digit. We care about indices 1 and 2
	for _, instruction := range instructions {
		num1, _ := strconv.Atoi(instruction[1])
		num2, _ := strconv.Atoi(instruction[2])
		sum += num1 * num2
	}

	return sum
}
