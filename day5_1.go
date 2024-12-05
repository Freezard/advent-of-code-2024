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

func day5_1() {
	rules := make(map[string][]string)
	sumValidUpdates := 0

	file, err := os.Open("day5-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputMode := "rules"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			inputMode = "updates"
			continue
		}

		if inputMode == "rules" {
			rule := strings.Split(line, "|")

			rules[rule[1]] = append(rules[rule[1]], rule[0])
		} else if inputMode == "updates" {
			pageUpdate := strings.Split(line, ",")

			valid := isPageUpdateValid(pageUpdate, rules)
			if valid {
				middleNumber := pageUpdate[(len(pageUpdate)-1)/2]
				middleNumberInt, _ := strconv.Atoi(middleNumber)
				sumValidUpdates += middleNumberInt
			}
		}
	}

	fmt.Println(sumValidUpdates)
}

func isPageUpdateValid(pageUpdate []string, rules map[string][]string) bool {
	for i, pageNumber := range pageUpdate {
		for j := i + 1; j < len(pageUpdate); j++ {
			if !slices.Contains(rules[pageUpdate[j]], pageNumber) {
				return false
			}
		}
	}
	return true
}
