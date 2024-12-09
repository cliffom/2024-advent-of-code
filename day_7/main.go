package main

import (
	"bufio"
	"fmt"
	"os"
)

func getCalibrationEquationsFromInput(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()

	var calibrationEquations []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		calibrationEquations = append(calibrationEquations, scanner.Text())
	}

	return calibrationEquations
}

func main() {
	sum := 0
	refinedSum := 0
	calibrationEquations := getCalibrationEquationsFromInput("input.txt")

	for _, calibrationEquation := range calibrationEquations {
		ce := &CalibrationEquation{}
		if !ce.FromInputData(calibrationEquation) {
			continue
		}

		if ce.IsValid(false) {
			sum += ce.Result
		}

		if ce.IsValid(true) {
			refinedSum += ce.Result
		}
	}

	fmt.Println("The sum of valid calibration results is:", sum)
	fmt.Println("The sum of valid and refined calibration results is:", refinedSum)
}
