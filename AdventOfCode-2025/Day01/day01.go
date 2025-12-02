package main

import (
	"aoc/utils"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

type wheel struct {
	dialNow   int
	numOfZero int
}

// wheel helper functions

func newWheel() wheel {
	newWheel := wheel{50, 0}
	return newWheel
}

// rotate wheel, advanced method is true if we're in Part2 and we have the password method 0x434C49434B enabled
func rotateWheel(myWheel wheel, direction string, times int, advancedMethod bool) wheel {
	dirTimes := utils.Mod(times, 100)
	if direction == "L" {
		dirTimes *= -1
	}

	tempDial := myWheel.dialNow
	tempNumOfZero := myWheel.numOfZero

	if advancedMethod {
		tempNumOfZero += abs(times) / 100
	}

	tempDial = utils.Mod(tempDial+dirTimes, 100)

	if tempDial == 0 && dirTimes != 0 {
		tempNumOfZero += 1
	} else if advancedMethod && myWheel.dialNow != 0 && (dirTimes+myWheel.dialNow < 0 || myWheel.dialNow+dirTimes > 99) {
		tempNumOfZero += 1
	}

	myWheel.dialNow = tempDial
	myWheel.numOfZero = tempNumOfZero
	return myWheel

}

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

	lockWheel := wheel{50, 0}
	for _, v := range input {
		direction := string(v[0])
		times, err := strconv.Atoi(v[1:])
		utils.Check(err)
		lockWheel = rotateWheel(lockWheel, direction, times, false)
	}
	fmt.Println("Password is", lockWheel.numOfZero)
}

func part2(input []string) {
	fmt.Print("Part2 ")

	lockWheel := newWheel()
	for _, v := range input {
		direction := string(v[0])
		times, err := strconv.Atoi(v[1:])
		utils.Check(err)
		lockWheel = rotateWheel(lockWheel, direction, times, true)
	}
	fmt.Println("Password is", lockWheel.numOfZero)
}
