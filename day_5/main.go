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

	sum := 0
	for _, update := range updates {
		updateIsValid := validateUpdate(update, orderingRules)
		if updateIsValid {
			middlePageNumber := (len(update) / 2)
			sum += update[middlePageNumber]
		}
	}

	log.Printf("The sum of the middle page numbers is: %v", sum)
}

func validateUpdate(update []int, orderingRules [][2]int) bool {
	updateLen := len(update)
	var updateIsValid bool
	for i, page := range update {
		if i+1 < updateLen {
			ruleToValidate := [2]int{page, update[i+1]}
			for _, rule := range orderingRules {
				if ruleToValidate == rule {
					updateIsValid = true
					break
				}
				updateIsValid = false
			}
			if !updateIsValid {
				break
			}
		}
	}

	return updateIsValid
}
