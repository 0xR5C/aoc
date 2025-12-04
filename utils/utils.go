package utils

import (
	"strconv"
)

func Works() int {
	return 0
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ArrayAtoi(arr []string) []int {
	numArr := make([]int, len(arr))
	for i := range arr {
		val, ok := strconv.Atoi(arr[i])
		Check(ok)
		numArr[i] = val
	}
	return numArr
}

func AbsDiff(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func MaxRune(ar []rune) int {
	if len(ar) == 0 {
		return -1
	}
	max := 0
	for i, v := range ar {
		if ar[max] < v {
			max = i
		}
	}
	return max
}
