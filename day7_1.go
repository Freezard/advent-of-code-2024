package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day7_1() {
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
		validEquation := checkEquation(convertedEquation, testValue, 0)
		if validEquation {
			totalResult += testValue
		}
	}

	fmt.Println(totalResult)
}

func checkEquation(equation []int, testValue int, currentValue int) bool {
	if currentValue == 0 {
		currentValue = equation[0]
		equation = equation[1:]
	}

	if len(equation) == 0 {
		return currentValue == testValue
	}

	product := currentValue * equation[0]

	if product == testValue {
		return true
	} else if product > testValue {
		sum := currentValue + equation[0]
		if sum == testValue {
			return true
		} else if sum > testValue {
			return false
		} else {
			return checkEquation(equation[1:], testValue, sum)
		}
	} else {
		if checkEquation(equation[1:], testValue, product) {
			return true
		} else if checkEquation(equation[1:], testValue, currentValue+equation[0]) {
			return true
		}
	}
	return false
}

func convertToNumbers(equation []string) []int {
	var convertedEquation []int
	for _, number := range equation {
		convertedNumber, _ := strconv.Atoi(number)
		convertedEquation = append(convertedEquation, convertedNumber)
	}
	return convertedEquation
}
