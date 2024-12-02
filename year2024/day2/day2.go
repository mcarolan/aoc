package day2

import (
	"aoc/util"
)

func allIncreasingOrDecreasing(numbers []int) bool {
	if len(numbers) < 2 {
		return true
	}
	increasing := numbers[1] > numbers[0]
	for i := 1; i < len(numbers); i++ {
		if (increasing && numbers[i] <= numbers[i-1]) || (!increasing && numbers[i] >= numbers[i-1]) {
			return false
		}
	}
	return true
}

func adjacentDifferenceCorrect(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		if i > 0 && util.Abs(numbers[i]-numbers[i-1]) > 3 {
			return false
		}
		if i < len(numbers)-1 && util.Abs(numbers[i]-numbers[i+1]) > 3 {
			return false
		}
	}
	return true
}

func removeOne(input []int) <-chan []int {
	ch := make(chan []int)

	go func() {
		for i := range len(input) {
			next := make([]int, 0)
			next = append(next, input[:i]...)
			next = append(next, input[i+1:]...)
			ch <- next
		}
		close(ch)
	}()

	return ch
}

func SolvePart1(input string) int {
	reports := util.Map(util.Lines(input), util.AllNumbers)

	counter := 0
	for _, report := range reports {
		if allIncreasingOrDecreasing(report) && adjacentDifferenceCorrect(report) {
			counter++
		}
	}
	return counter
}

func SolvePart2(input string) int {
	reports := util.Map(util.Lines(input), util.AllNumbers)

	counter := 0
	for _, report := range reports {
		if allIncreasingOrDecreasing(report) && adjacentDifferenceCorrect(report) {
			counter++
			continue
		} else {
			for newReport := range removeOne(report) {
				if allIncreasingOrDecreasing(newReport) && adjacentDifferenceCorrect(newReport) {
					counter++
					break
				}
			}
		}
	}
	return counter
}
