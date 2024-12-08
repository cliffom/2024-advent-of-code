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
	calibrationEquations := getCalibrationEquationsFromInput("input.txt")

	for _, calibrationEquation := range calibrationEquations {
		ce := &CalibrationEquation{}
		if !ce.FromInputData(calibrationEquation) {
			continue
		}

		if ce.IsValid() {
			sum += ce.Result
		}
	}

	fmt.Println("The sum of valid calibration results is:", sum)
}
