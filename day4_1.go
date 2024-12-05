package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var searchWord = "XMAS"
var totalWords = 0

func day4_1() {
	var searchField [][]string

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
				searchRight(searchField, i, j)
				searchLeft(searchField, i, j)
				searchUp(searchField, i, j)
				searchDown(searchField, i, j)
				searchRightUp(searchField, i, j)
				searchRightDown(searchField, i, j)
				searchLeftUp(searchField, i, j)
				searchLeftDown(searchField, i, j)
			}
		}
	}

	fmt.Println(totalWords)
}

func searchRight(searchField [][]string, row int, col int) {
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

func searchLeft(searchField [][]string, row int, col int) {
	if col-3 < 0 {
		return
	}

	for i := 1; i < len(searchWord); i++ {
		if searchField[row][col-i] != string(searchWord[i]) {
			return
		}
	}
	totalWords += 1
}

func searchUp(searchField [][]string, row int, col int) {
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

func searchDown(searchField [][]string, row int, col int) {
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

func searchRightUp(searchField [][]string, row int, col int) {
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

func searchRightDown(searchField [][]string, row int, col int) {
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

func searchLeftUp(searchField [][]string, row int, col int) {
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

func searchLeftDown(searchField [][]string, row int, col int) {
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
