package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day4_2() {
	var searchField [][]string
	totalXMAS := 0

	file, err := os.Open("day4-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, char := range line {
			row = append(row, string(char))
		}
		searchField = append(searchField, row)
	}

	for i, row := range searchField {
		for j := range row {
			if searchField[i][j] == string("A") {
				if searchXMAS(searchField, i, j) {
					totalXMAS += 1
				}
			}
		}
	}

	fmt.Println(totalXMAS)
}

func searchXMAS(searchField [][]string, row int, col int) bool {
	if col+1 >= len(searchField[row]) || col-1 < 0 || row+1 >= len(searchField[col]) || row-1 < 0 {
		return false
	}

	if (searchField[row-1][col-1] == string("M") && searchField[row+1][col+1] == string("S") ||
		searchField[row-1][col-1] == string("S") && searchField[row+1][col+1] == string("M")) &&
		(searchField[row-1][col+1] == string("M") && searchField[row+1][col-1] == string("S") ||
			searchField[row-1][col+1] == string("S") && searchField[row+1][col-1] == string("M")) {
		return true
	}

	return false
}
