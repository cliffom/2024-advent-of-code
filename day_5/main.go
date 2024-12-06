package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInputFromFile(filename string) ([][2]int, [][]int) {
	file, _ := os.Open(filename)
	defer file.Close()

	var orderingRules [][2]int
	var updates [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "|") {
			orderingRules = append(orderingRules, getOrderingRules(text))
		} else if strings.Contains(text, ",") {
			updates = append(updates, getUpdateData(text))
		}
	}

	return orderingRules, updates
}

func getOrderingRules(data string) [2]int {
	strings := strings.Split(data, "|")
	num1, _ := strconv.Atoi(strings[0])
	num2, _ := strconv.Atoi(strings[1])
	return [2]int{num1, num2}
}

func getUpdateData(data string) []int {
	strings := strings.Split(data, ",")
	var updates []int
	for _, v := range strings {
		update, _ := strconv.Atoi(v)
		updates = append(updates, update)
	}

	return updates
}

func main() {
	orderingRules, updates := readInputFromFile("input.txt")

	var sumValid, sumInvalid int
	for _, update := range updates {
		updateIsValid := validateUpdate(update, orderingRules)
		if updateIsValid {
			middlePageNumber := (len(update) / 2)
			sumValid += update[middlePageNumber]
		} else {
			correctedUpdate := correctInvalidUpdate(update, orderingRules)
			middlePageNumber := (len(correctedUpdate) / 2)
			sumInvalid += correctedUpdate[middlePageNumber]
		}
	}

	log.Printf("The sum of the middle page numbers for valid entries is: %v", sumValid)
	log.Printf("The sum of the middle page numbers for invalid entries is: %v", sumInvalid)
}

func validateUpdate(update []int, orderingRules [][2]int) bool {
	var validate func(index int) bool

	validate = func(index int) bool {
		if index >= len(update)-1 {
			return true
		}

		ruleToValidate := [2]int{update[index], update[index+1]}
		isValid := false
		for _, rule := range orderingRules {
			if ruleToValidate == rule {
				isValid = true
				break
			}
		}

		if !isValid {
			return false
		}

		return validate(index + 1)
	}

	return validate(0)
}

func correctInvalidUpdate(update []int, orderingRules [][2]int) []int {
	var repairUpdate func(update []int, index int) []int

	repairUpdate = func(update []int, index int) []int {
		if index >= len(update)-1 {
			return update
		}

		ruleToValidateReversed := [2]int{update[index+1], update[index]}

		for _, rule := range orderingRules {
			if ruleToValidateReversed == rule {
				update[index], update[index+1] = update[index+1], update[index]
				return repairUpdate(update, 0)
			}
		}

		return repairUpdate(update, index+1)
	}

	return repairUpdate(update, 0) // Start the recursion
}
