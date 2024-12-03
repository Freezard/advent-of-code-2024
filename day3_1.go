package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func day3_1() {
	sumMultiplications := 0

	file, err := os.ReadFile("day3-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)

	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		digit1, _ := strconv.Atoi(match[1])
		digit2, _ := strconv.Atoi(match[2])
		sumMultiplications += digit1 * digit2
	}

	fmt.Println(sumMultiplications)
}
