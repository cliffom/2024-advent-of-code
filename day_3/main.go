package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// We want to match input items that look like
// mul(x,y) where x and 7 are numbers.
// Ex: mul(2,3)
var re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var unrefined_ouput, refined_output int
	var builder strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	instructions := builder.String()
	unrefined_ouput += sumOfValidUnrefinedInstructions(instructions)
	refined_output += sumOfValidRefinedInstructions(instructions)

	log.Printf("Sum of unrefined instructions: %v", unrefined_ouput)
	log.Printf("Sum of refined instructions: %v", refined_output)
}

// sumOfValidUnrefinedInstructions iterates through all valid instructions
// and sums the value of the product of each pair
func sumOfValidUnrefinedInstructions(data string) int {
	instructions := re.FindAllStringSubmatch(data, -1)

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

// sumOfValidRefinedInstructions iterates through all valid
// instructions and disables processing of instructions
// that follow a `don't` instruction. Instructions
// after `do` are processed normally.
func sumOfValidRefinedInstructions(line string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Find all matches
	matches := re.FindAllStringSubmatchIndex(line, -1)

	// Filter matches not between "don't" and "do"
	result := []string{}
	for _, match := range matches {
		start := match[0]
		end := match[1]

		substr := line[:start]
		if !strings.Contains(substr, "don't") || strings.LastIndex(substr, "don't") < strings.LastIndex(substr, "do") {
			result = append(result, line[start:end])
		}
	}

	sum := 0
	for _, match := range result {
		sum += sumOfValidUnrefinedInstructions(match)
	}

	return sum
}
