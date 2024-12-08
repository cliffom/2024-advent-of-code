package main

import (
	"regexp"
	"strconv"
	"strings"
)

type CalibrationEquation struct {
	Result   int
	Operands []int
}

func (ce *CalibrationEquation) FromInputData(input string) bool {
	pattern := `(\d+):\s((?:\d+\s?)+)`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(input)

	if len(match) <= 2 {
		return false
	}

	result, _ := strconv.Atoi(match[1])
	numStrs := strings.Fields(match[2])

	var numbers []int
	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		numbers = append(numbers, num)
	}

	ce.Result = result
	ce.Operands = numbers

	return true
}

func (ce *CalibrationEquation) IsValid() bool {
	numberOfOperands := len(ce.Operands)

	var search func(index, current int) bool
	search = func(index, current int) bool {
		if index == numberOfOperands {
			return current == ce.Result
		}

		if search(index+1, current+ce.Operands[index]) {
			return true
		}

		if search(index+1, current*ce.Operands[index]) {
			return true
		}
		return false
	}

	return search(1, ce.Operands[0])
}
