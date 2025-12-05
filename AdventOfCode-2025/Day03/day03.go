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
	sum := 0
	for _, line := range input {
		bank := []rune(line)
		var firstDigit, secondDigit rune

		max := utils.MaxRune(bank)
		if max != len(bank)-1 {
			firstDigit = bank[max]
			tempBank := bank[max+1:]
			max = utils.MaxRune(bank[max+1:])
			secondDigit = tempBank[max]
		} else {
			secondDigit = bank[max]
			max = utils.MaxRune(bank[:max])
			firstDigit = bank[max]
		}
		numStr := []rune{firstDigit, secondDigit}

		num, err := strconv.Atoi(string(numStr))
		utils.Check(err)
		sum += num
	}

	fmt.Println("The sum of voltage is", sum)
}

func part2(input []string) {
	fmt.Print("Part2 ")

	sum := 0

	for _, line := range input {
		bank := []rune(line)
		var digits []rune

		// find maximum digit that has at least 11 digits before, then 10 etc.
		for i := 12; i >= 1; i-- {
			temp := bank[:len(bank)-i+1]
			max := utils.MaxRune(temp)

			digits = append(digits, bank[max])
			bank = bank[max+1:]
		}

		numStr := string(digits)
		num, err := strconv.Atoi(numStr)
		utils.Check(err)
		sum += num
	}

	fmt.Println("The sum of voltage is", sum)
}
