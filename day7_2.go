package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day7_2() {
	var equations [][]string
	totalResult := 0

	file, err := os.Open("day7-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		equations = append(equations, line)
	}

	for _, equation := range equations {
		testValue, _ := strconv.Atoi(strings.Trim(equation[0], ":"))

		convertedEquation := convertToNumbers(equation[1:])
		validEquation := checkEquationPart2(convertedEquation, testValue, 0)
		if validEquation {
			totalResult += testValue
		}
	}

	fmt.Println(totalResult)
}

func checkEquationPart2(equation []int, testValue int, currentValue int) bool {
	if currentValue == 0 {
		currentValue = equation[0]
		equation = equation[1:]
	}

	if len(equation) == 0 {
		return currentValue == testValue
	}

	if currentValue > testValue {
		return false
	}

	nextValue := equation[0]
	remainingEquation := equation[1:]

	if checkEquationPart2(remainingEquation, testValue, currentValue+nextValue) {
		return true
	}

	if checkEquationPart2(remainingEquation, testValue, currentValue*nextValue) {
		return true
	}

	concatString := strconv.Itoa(currentValue) + strconv.Itoa(nextValue)
	concatValue, _ := strconv.Atoi(concatString)
	return checkEquationPart2(remainingEquation, testValue, concatValue)
}
