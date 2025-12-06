package main

import (
	"aoc/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day03")
	// inputTxt := "test-input.txt"
	inputTxt := "input.txt"
	dat, err := os.ReadFile(inputTxt)
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	partPtr := flag.Int("part", 0, "Choose part to run")
	flag.Parse()

	if *partPtr == 1 {
		part1(lines)
	} else if *partPtr == 2 {
		part2(lines)
	} else {
		part1(lines)
		part2(lines)
	}
}

func part1(input []string) {
	fmt.Print("Part1 ")
	rollMap := createRollMap(input)

	count := 0
	for i, line := range rollMap {
		for j, val := range line {
			rollsAround := 0
			if val == "." {
				continue
			}
			// check all directions
			// i-1, j-1
			if i-1 >= 0 && j-1 >= 0 && rollMap[i-1][j-1] == "@" {
				rollsAround += 1
			}

			// i-1, j
			if i-1 >= 0 && rollMap[i-1][j] == "@" {
				rollsAround += 1
			}

			// i-1, j+1
			if i-1 >= 0 && j+1 < len(line) && rollMap[i-1][j+1] == "@" {
				rollsAround += 1
			}

			// i, j-1
			if j-1 >= 0 && rollMap[i][j-1] == "@" {
				rollsAround += 1
			}

			// i, j+1
			if j+1 < len(line) && rollMap[i][j+1] == "@" {
				rollsAround += 1
			}

			// i+1, j-1
			if i+1 < len(rollMap) && j-1 >= 0 && rollMap[i+1][j-1] == "@" {
				rollsAround += 1
			}

			// i+1, j
			if i+1 < len(rollMap) && rollMap[i+1][j] == "@" {
				rollsAround += 1
			}

			// i+1, j+1
			if i+1 < len(rollMap) && j+1 < len(line) && rollMap[i+1][j+1] == "@" {
				rollsAround += 1
			}

			if rollsAround < 4 {
				count += 1
			}

		}
	}
	fmt.Println("Number of accessible rolls is", count)

}

func part2(input []string) {
	fmt.Print("Part2 ")
	fmt.Println()
}

// Helper functions

func createRollMap(input []string) [][]string {
	var rollMap [][]string
	for _, line := range input {
		var temp []string
		for _, val := range line {
			temp = append(temp, string(val))
		}
		rollMap = append(rollMap, temp)
	}
	return rollMap
}
