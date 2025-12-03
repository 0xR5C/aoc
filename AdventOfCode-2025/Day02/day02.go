package main

import (
	"aoc/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day01")
	// inputTxt := "test-input.txt"
	inputTxt := "input.txt"
	dat, err := os.ReadFile(inputTxt)
	utils.Check(err)
	input := string(dat[:])
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var ranges []string
	for _, val := range lines {
		ranges = append(ranges, strings.Split(val, ",")...)
		if ranges[len(ranges)-1] == "" {
			ranges = ranges[:len(ranges)-1]
		}
	}
	partPtr := flag.Int("part", 0, "Choose part to run")
	flag.Parse()

	if *partPtr == 1 {
		part1(ranges)
	} else if *partPtr == 2 {
		part2(ranges)
	} else {
		part1(ranges)
		part2(ranges)
	}
}

func part1(input []string) {
	fmt.Print("Part1 ")
	sum := 0

	for _, r := range input {
		ranges := strings.Split(r, "-")
		begin, err := strconv.Atoi(ranges[0])
		utils.Check(err)
		end, err := strconv.Atoi(ranges[1])
		utils.Check(err)
		for num := begin; num <= end; num++ {
			tempStr := strconv.Itoa(num)
			if checkInvalid(tempStr) {
				sum += num
			}
		}
	}
	fmt.Println("Sum of invalid IDs is", sum)
}

func part2(input []string) {
	fmt.Print("Part2 ")
	fmt.Println()
}

// Helper Functions

// returns true if id is invalid, else false
func checkInvalid(id string) bool {

	substr1 := id[:len(id)/2]
	substr2 := id[len(id)/2:]
	return substr1 == substr2

}
