package day6

import (
	"regexp"
	"strconv"
)

var timerRegex = regexp.MustCompile(`([0-9]+),?`)

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

const (
	DaysPart1  = 80
	DaysPart2  = 256
	ResetTimer = 6
	NewTimer   = 8
)

func simulate(input string, days int) (int, error) {
	timers, err := extractNInts(input, timerRegex, -1)

	if err != nil {
		return 0, err
	}

	state := make([]int, NewTimer+1)

	for _, timer := range timers {
		state[timer]++
	}

	for range days {
		numberToAppend := state[0]
		copy(state, state[1:])
		state[ResetTimer] += numberToAppend
		state[NewTimer] = numberToAppend
	}

	total := 0
	for _, i := range state {
		total += i
	}

	return total, nil
}

func SolvePart1(input string) (int, error) {
	return simulate(input, DaysPart1)
}

func SolvePart2(input string) (int, error) {
	return simulate(input, DaysPart2)
}
