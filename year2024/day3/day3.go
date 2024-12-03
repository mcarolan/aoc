package day3

import (
	"regexp"
	"strconv"
)

const DISABLE_INSTRUCTION = "don't()"
const ENABLE_INSTRUCTION = "do()"

func SolvePart1(input string) int {
	multiplicationMatcher := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := multiplicationMatcher.FindAllStringSubmatch(input, -1)

	result := 0

	for _, match := range matches {
		a, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		result += a * b
	}
	return result
}

func SolvePart2(input string) int {
	multiplicationMatcher := regexp.MustCompile(`(don't\(\))|(do\(\))|mul\((\d+),(\d+)\)`)
	matches := multiplicationMatcher.FindAllStringSubmatch(input, -1)

	result := 0
	isEnabled := true

	for _, match := range matches {
		if match[0] == DISABLE_INSTRUCTION {
			isEnabled = false
			continue
		} else if match[0] == ENABLE_INSTRUCTION {
			isEnabled = true
			continue
		}

		if !isEnabled {
			continue
		}

		a, err := strconv.Atoi(match[3])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(match[4])
		if err != nil {
			panic(err)
		}
		result += a * b
	}
	return result
}
