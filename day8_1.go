package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day8_1() {
	var cityMap [][]string
	antennaPositions := make(map[string][]Position)
	antinodeUniqueLocations := 0

	file, err := os.Open("day8-input.txt")
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
		cityMap = append(cityMap, row)
	}

	for i := range cityMap {
		for j := range cityMap[i] {
			if cityMap[i][j] != "." {
				antennaPositions[cityMap[i][j]] = append(antennaPositions[cityMap[i][j]], Position{j, i})
			}
		}
	}

	for _, antennaPositions := range antennaPositions {
		for _, antennaPos := range antennaPositions {
			for _, otherAntennaPos := range antennaPositions {
				if antennaPos == otherAntennaPos {
					continue
				}

				distance := antennaPos.Subtract(otherAntennaPos)
				mirroredPos := antennaPos.Add(distance)

				if mirroredPos.X < 0 || mirroredPos.X >= len(cityMap[0]) || mirroredPos.Y < 0 || mirroredPos.Y >= len(cityMap) {
					continue
				}

				if cityMap[mirroredPos.Y][mirroredPos.X] == "#" {
					continue
				}

				cityMap[mirroredPos.Y][mirroredPos.X] = "#"

				antinodeUniqueLocations++
			}
		}

	}

	fmt.Println(antinodeUniqueLocations)
}

func (pos1 Position) Add(pos2 Position) Position {
	return Position{
		X: pos1.X + pos2.X,
		Y: pos1.Y + pos2.Y,
	}
}

func (pos1 Position) Subtract(pos2 Position) Position {
	return Position{
		X: pos1.X - pos2.X,
		Y: pos1.Y - pos2.Y,
	}
}
