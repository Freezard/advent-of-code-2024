package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day1_2() {
	var listLeft = make(map[int]int)
	var listRight = make(map[int]int)

	file, err := os.Open("day1-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		locationIDLeft, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		locationIDRight, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		listLeft[locationIDLeft] = listLeft[locationIDLeft] + 1
		listRight[locationIDRight] = listRight[locationIDRight] + 1
	}

	var totalSimilarityScore int
	for locationID, count := range listLeft {
		similarityScore := locationID * count * listRight[locationID]
		totalSimilarityScore += similarityScore
	}

	fmt.Println(totalSimilarityScore)
}
