package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var searchField [][]string
var searchWord = "XMAS"
var totalWords = 0

func day4_1() {
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
			if searchField[i][j] == string(searchWord[0]) {
				searchRight(i, j)
				searchLeft(i, j)
				searchUp(i, j)
				searchDown(i, j)
				searchRightUp(i, j)
				searchRightDown(i, j)
				searchLeftUp(i, j)
				searchLeftDown(i, j)
			}
		}
	}

	fmt.Println(totalWords)
}

func searchRight(row int, col int) {
	if col+3 >= len(searchField[row]) {
		return
	}

	for i := 1; i < len(searchWord); i++ {
		if searchField[row][col+i] != string(searchWord[i]) {
			return
		}
	}
	totalWords += 1
}

func searchLeft(row int, col int) {
	if col-3 <= 0 {
		return
	}

	for i := 1; i < len(searchWord); i++ {
		if searchField[row][col-i] != string(searchWord[i]) {
			return
		}
	}
	totalWords += 1
}

func searchUp(row int, col int) {
	if row-3 < 0 {
		return
	}

	for i := 1; i < len(searchWord); i++ {
		if searchField[row-i][col] != string(searchWord[i]) {
			return
		}
	}
	totalWords += 1
}

func searchDown(row int, col int) {
	if row+3 >= len(searchField[col]) {
		return
	}

	for i := 1; i < len(searchWord); i++ {
		if searchField[row+i][col] != string(searchWord[i]) {
			return
		}
	}
	totalWords += 1
}

func searchRightUp(row int, col int) {
	if col+3 >= len(searchField[row]) || row-3 < 0 {
		return
	}

	for i := 1; i < len(searchWord); i++ {
		if searchField[row-i][col+i] != string(searchWord[i]) {
			return
		}
	}
	totalWords += 1
}

func searchRightDown(row int, col int) {
	if col+3 >= len(searchField[row]) || row+3 >= len(searchField[col]) {
		return
	}

	for i := 1; i < len(searchWord); i++ {
		if searchField[row+i][col+i] != string(searchWord[i]) {
			return
		}
	}
	totalWords += 1
}

func searchLeftUp(row int, col int) {
	if col-3 < 0 || row-3 < 0 {
		return
	}

	for i := 1; i < len(searchWord); i++ {
		if searchField[row-i][col-i] != string(searchWord[i]) {
			return
		}
	}
	totalWords += 1
}

func searchLeftDown(row int, col int) {
	if col-3 < 0 || row+3 >= len(searchField[col]) {
		return
	}

	for i := 1; i < len(searchWord); i++ {
		if searchField[row+i][col-i] != string(searchWord[i]) {
			return
		}
	}
	totalWords += 1
}
