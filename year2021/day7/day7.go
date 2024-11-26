package day7

import (
	"math"
	"regexp"
	"strconv"
)

var numberRegex = regexp.MustCompile(`([0-9]+),?`)

func extractNInts(input string, regex *regexp.Regexp, n int) ([]int, error) {
	var result []int
	matches := regex.FindAllStringSubmatch(input, n)

	if matches == nil {
		return nil, nil
	}

	for _, match := range matches {
		number, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, err
		}
		result = append(result, number)
	}

	return result, nil
}

func parseInput(input string) ([]int, error) {
	return extractNInts(input, numberRegex, -1)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func SolvePart1(input string) (int, error) {
	horizontalPositions, err := parseInput(input)

	if err != nil {
		return 0, err
	}

	minFuel := math.MaxInt

	for _, pos := range horizontalPositions {
		fuel := 0

		for _, i := range horizontalPositions {
			fuel += abs(pos - i)
		}

		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel, nil
}

func SolvePart2(input string) (int, error) {
	horizontalPositions, err := parseInput(input)

	if err != nil {
		return 0, err
	}

	maxPosition := 0
	for _, pos := range horizontalPositions {
		maxPosition = max(maxPosition, pos)
	}
	minFuel := math.MaxInt

	for pos := range maxPosition {
		fuel := 0

		for _, i := range horizontalPositions {
			n := abs(pos - i)
			fuel += (n * (n + 1)) / 2
		}

		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel, nil
}
