package day1

import (
	"strconv"
	"strings"
)

func SolvePart1(input string) (int, error) {
	numbers, err := inputToInts(input)

	if err != nil {
		return 0, err
	}

	return countIncreases(numbers), nil
}

func SolvePart2(input string) (int, error) {
	numbers, err := inputToInts(input)

	if err != nil {
		return 0, err
	}

	var windowSums []int

	for i := 2; i < len(numbers); i += 1 {
		windowSum := numbers[i] + numbers[i-1] + numbers[i-2]
		windowSums = append(windowSums, windowSum)
	}

	return countIncreases(windowSums), nil
}

func inputToInts(input string) ([]int, error) {
	lines := strings.Split(input, "\n")
	var numbers []int

	for _, line := range lines {
		number, err := strconv.Atoi(line)

		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}
	return numbers, nil
}

func countIncreases(numbers []int) int {
	var increases int
	var last int

	for i, number := range numbers {
		if i > 0 && number > last {
			increases += 1
		}
		last = number
	}

	return increases
}
