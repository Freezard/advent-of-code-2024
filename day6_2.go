package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day6_2() {
	var mapLayout [][]string
	var startPosition Position
	startDirection := "UP"

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
		for _, char := range line {
			row = append(row, string(char))
		}
		mapLayout = append(mapLayout, row)
		yPos++
	}

	startPosition = findStartPosition(mapLayout)
	visitedMap := markMapAsVisited(startPosition, startDirection, mapLayout)
	totalInfiniteLoops := getTotalInfiniteLoops(startPosition, startDirection, visitedMap)

	fmt.Println(totalInfiniteLoops)
}

func markMapAsVisited(currPosition Position, currDirection string, mapLayout [][]string) [][]string {
	walking := true
	for walking {
		nextObstacle := getNextObstacle(currPosition, currDirection, mapLayout)
		if nextObstacle == "" {
			break
		} else if nextObstacle == "." {
			currPosition = getNextPosition(currPosition, currDirection)
			markPositionAsVisited(currPosition, mapLayout)
		} else if nextObstacle == "X" {
			currPosition = getNextPosition(currPosition, currDirection)
		} else if nextObstacle == "#" {
			currDirection = getNextDirection(currDirection, nextObstacle)
		}
	}
	return mapLayout
}

func checkForInfiniteLoop(currPosition Position, currDirection string, mapLayout [][]string) bool {
	for i := 0; i <= 10000; i++ {
		nextObstacle := getNextObstacle(currPosition, currDirection, mapLayout)
		if nextObstacle == "" {
			break
		} else if nextObstacle == "." || nextObstacle == "X" || nextObstacle == "^" {
			currPosition = getNextPosition(currPosition, currDirection)
		} else if nextObstacle == "#" {
			currDirection = getNextDirection(currDirection, nextObstacle)
		}
		/* Bad way of checking infinite loop.
		Should check if any position has been visited multiple times from the same direction */
		if i == 10000 {
			return true
		}
	}
	return false
}

func getTotalInfiniteLoops(startPosition Position, startDirection string, mapLayout [][]string) int {
	totalInfiniteLoops := 0
	for i := range mapLayout {
		for j := range mapLayout[i] {
			if mapLayout[i][j] == "X" {
				addObjectToPosition(Position{j, i}, mapLayout)
				infiniteLoop := checkForInfiniteLoop(startPosition, startDirection, mapLayout)
				if infiniteLoop {
					totalInfiniteLoops++
				}
				markPositionAsVisited(Position{j, i}, mapLayout)
			}
		}
	}
	return totalInfiniteLoops
}

func addObjectToPosition(position Position, mapLayout [][]string) {
	mapLayout[position.Y][position.X] = "#"
}
