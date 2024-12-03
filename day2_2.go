package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day2_2() {
	totalSafeLevels := 0

	file, err := os.Open("day2-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var report []int

		for _, str := range line {
			level, err := strconv.Atoi(str)
			if err != nil {
				return
			}
			report = append(report, level)
		}

		if isIncreasingSafely(report, true) || isDecreasingSafely(report, true) {
			totalSafeLevels += 1
		}
	}
	fmt.Println(totalSafeLevels)
}

func isIncreasingSafely(report []int, allowRemoveLevel bool) bool {
	for i := 1; i < len(report); {
		if report[i]-report[i-1] < 1 || report[i]-report[i-1] > 3 {
			if allowRemoveLevel {
				withoutCurrentLevel := append(append([]int{}, report[:i]...), report[i+1:]...)
				withoutPreviousLevel := append(append([]int{}, report[:i-1]...), report[i:]...)

				return isIncreasingSafely(withoutCurrentLevel, false) || isIncreasingSafely(withoutPreviousLevel, false)
			} else {
				return false
			}
		}
		i++
	}
	return true
}

func isDecreasingSafely(report []int, allowRemoveLevel bool) bool {
	for i := 1; i < len(report); {
		if report[i-1]-report[i] < 1 || report[i-1]-report[i] > 3 {
			if allowRemoveLevel {
				withoutCurrentLevel := append(append([]int{}, report[:i]...), report[i+1:]...)
				withoutPreviousLevel := append(append([]int{}, report[:i-1]...), report[i:]...)

				return isDecreasingSafely(withoutCurrentLevel, false) || isDecreasingSafely(withoutPreviousLevel, false)
			} else {
				return false
			}
		}
		i++
	}
	return true
}
