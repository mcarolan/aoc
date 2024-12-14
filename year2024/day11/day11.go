package day11

import (
	"aoc/util"
	"fmt"
	"strconv"
)

func blink(stones map[int]int) map[int]int {
	result := make(map[int]int)

	for stone, count := range stones {
		if stone == 0 {
			result[1] += count
		} else {
			digits := fmt.Sprintf("%d", stone)
			if len(digits)%2 == 0 {
				mid := len(digits) / 2
				leftString := digits[:mid]
				rightString := digits[mid:]

				left, err := strconv.Atoi(leftString)
				if err != nil {
					panic(err)
				}

				right, err := strconv.Atoi(rightString)
				if err != nil {
					panic(err)
				}
				result[left] += count
				result[right] += count
			} else {
				result[stone*2024] += count
			}
		}
	}

	return result
}

func SolvePart1(input string) int {
	startingStones := util.AllNumbers(input)

	stones := make(map[int]int)
	for _, stone := range startingStones {
		stones[stone]++
	}

	for range 25 {
		stones = blink(stones)
	}

	result := 0

	for _, count := range stones {
		result += count
	}

	return result
}

func SolvePart2(input string) int {
	startingStones := util.AllNumbers(input)

	stones := make(map[int]int)
	for _, stone := range startingStones {
		stones[stone]++
	}

	for range 75 {
		stones = blink(stones)
	}

	result := 0

	for _, count := range stones {
		result += count
	}

	return result
}
