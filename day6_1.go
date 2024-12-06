package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Position struct {
	X int
	Y int
}

func day6_1() {
	var mapLayout [][]string
	var currPosition Position
	currDirection := "UP"
	positionsVisited := 0

	file, err := os.Open("day6-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	yPos := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for xPos, char := range line {
			if string(char) == "^" {
				currPosition = Position{xPos, yPos}
			}
			row = append(row, string(char))
		}
		mapLayout = append(mapLayout, row)
		yPos++
	}

	startPosition := findStartPosition(mapLayout)
	markPositionAsVisited(startPosition, mapLayout)
	positionsVisited++

	walking := true
	for walking {
		nextObstacle := getNextObstacle(currPosition, currDirection, mapLayout)
		if nextObstacle == "" {
			break
		} else if nextObstacle == "." {
			currPosition = getNextPosition(currPosition, currDirection)
			markPositionAsVisited(currPosition, mapLayout)
			positionsVisited++
		} else if nextObstacle == "X" {
			currPosition = getNextPosition(currPosition, currDirection)
		} else if nextObstacle == "#" {
			currDirection = getNextDirection(currDirection, nextObstacle)
		}
	}
	fmt.Println(currPosition)
	fmt.Println(positionsVisited)
}

func findStartPosition(mapLayout [][]string) Position {
	for i := range mapLayout {
		for j := range mapLayout[i] {
			if mapLayout[i][j] == "^" {
				return Position{j, i}
			}
		}
	}
	return Position{}
}

func markPositionAsVisited(position Position, mapLayout [][]string) {
	mapLayout[position.Y][position.X] = "X"
}

func getNextDirection(currDirection, nextObstacle string) string {
	var nextDirection string

	if currDirection == "UP" {
		nextDirection = "RIGHT"
	} else if currDirection == "RIGHT" {
		nextDirection = "DOWN"
	} else if currDirection == "DOWN" {
		nextDirection = "LEFT"
	} else if currDirection == "LEFT" {
		nextDirection = "UP"
	}

	return nextDirection
}

func getNextPosition(currPosition Position, currDirection string) Position {
	nextPosition := Position{currPosition.X, currPosition.Y}

	if currDirection == "UP" {
		nextPosition.Y--
	} else if currDirection == "RIGHT" {
		nextPosition.X++
	} else if currDirection == "DOWN" {
		nextPosition.Y++
	} else if currDirection == "LEFT" {
		nextPosition.X--
	}

	return nextPosition
}

func getNextObstacle(currPosition Position, currDirection string, mapLayout [][]string) string {
	var nextObstacle string

	if currDirection == "UP" {
		if currPosition.Y-1 < 0 {
			nextObstacle = ""
		} else {
			nextObstacle = mapLayout[currPosition.Y-1][currPosition.X]
		}
	} else if currDirection == "RIGHT" {
		if currPosition.X+1 >= len(mapLayout[currPosition.X]) {
			nextObstacle = ""
		} else {
			nextObstacle = mapLayout[currPosition.Y][currPosition.X+1]
		}
	} else if currDirection == "DOWN" {
		if currPosition.Y+1 >= len(mapLayout) {
			nextObstacle = ""
		} else {
			nextObstacle = mapLayout[currPosition.Y+1][currPosition.X]
		}
	} else if currDirection == "LEFT" {
		if currPosition.X-1 < 0 {
			nextObstacle = ""
		} else {
			nextObstacle = mapLayout[currPosition.Y][currPosition.X-1]
		}
	}
	return nextObstacle
}
