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

func day5_2() {
	rules := make(map[string][]string)
	sumFixedUpdates := 0

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
			if !valid {
				fixedPageUpdate := fixPageUpdate(pageUpdate, rules)

				middleNumber := fixedPageUpdate[(len(fixedPageUpdate)-1)/2]
				middleNumberInt, _ := strconv.Atoi(middleNumber)
				sumFixedUpdates += middleNumberInt
			}
		}
	}

	fmt.Println(sumFixedUpdates)
}

func fixPageUpdate(pageUpdate []string, rules map[string][]string) []string {
	fixedPageUpdate := make([]string, len(pageUpdate))

	for i, currPageNumber := range pageUpdate {
		pageNumberIndex := 0
		for _, pageNumber := range pageUpdate {
			if slices.Contains(rules[pageUpdate[i]], pageNumber) {
				pageNumberIndex += 1
			}
		}
		fixedPageUpdate[pageNumberIndex] = currPageNumber
	}
	return fixedPageUpdate
}
