package day1

import (
	"aoc/util"
	"slices"
)

func SolvePart1(input string) int {
	numbers := util.Map(util.Lines(input), util.AllNumbers)

	firstColumn := make([]int, len(numbers))
	secondColumn := make([]int, len(numbers))

	for i := range len(numbers) {
		firstColumn[i] = numbers[i][0]
		secondColumn[i] = numbers[i][1]
	}

	slices.Sort(firstColumn)
	slices.Sort(secondColumn)

	result := 0

	for i := range len(numbers) {
		result += util.Abs(firstColumn[i] - secondColumn[i])
	}

	return result
}

func SolvePart2(input string) int {
	numbers := util.Map(util.Lines(input), util.AllNumbers)

	firstColumn := make([]int, len(numbers))
	secondColumnOccurances := make(map[int]int)

	for i, pair := range numbers {
		firstColumn[i] = pair[0]
		secondColumnOccurances[pair[1]]++
	}

	result := 0

	for _, n := range firstColumn {
		result += n * secondColumnOccurances[n]
	}

	return result
}
