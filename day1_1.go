package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day1_1() {
	var listLeft []int
	var listRight []int

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

		listLeft = append(listLeft, locationIDLeft)
		listRight = append(listRight, locationIDRight)
	}

	sort.Ints(listLeft)
	sort.Ints(listRight)

	var totalDifference int
	for i := 0; i < len(listLeft); i++ {
		difference := int(math.Abs(float64(listLeft[i] - listRight[i])))
		totalDifference += difference
	}

	fmt.Println(totalDifference)
}
