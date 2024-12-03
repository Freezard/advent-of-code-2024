package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day2_1() {
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

		isIncreasingSafely := slices.IsSortedFunc(report, func(a, b int) int {
			if a-b >= 1 && a-b <= 3 {
				return 1
			}
			return -1
		})
		isDecreasingSafely := slices.IsSortedFunc(report, func(a, b int) int {
			if b-a >= 1 && b-a <= 3 {
				return 1
			}
			return -1
		})
		if isIncreasingSafely || isDecreasingSafely {
			totalSafeLevels += 1
		}
	}

	fmt.Println(totalSafeLevels)
}
