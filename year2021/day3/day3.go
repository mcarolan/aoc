package day3

import (
	"strconv"
	"strings"
)

func SolvePart1(input string) int64 {
	lines := strings.Split(input, "\n")
	digits := len(lines[0])
	oneCount := make([]int, digits)

	for _, line := range lines {
		for i, char := range line {
			if char == '1' {
				oneCount[i]++
			}
		}
	}

	halfLineCount := len(lines) / 2

	gamma, epsilon := int64(0), int64(0)

	for _, count := range oneCount {
		gamma <<= 1
		epsilon <<= 1
		if count > halfLineCount {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}

	return gamma * epsilon
}

func whittle(lines []string, keepMajority bool) (int64, error) {
	digits := len(lines[0])

	for i := range digits {
		if len(lines) == 1 {
			break
		}
		var oneCount int
		for _, line := range lines {
			if line[i] == '1' {
				oneCount++
			}
		}

		half := len(lines) / 2
		oneMajority := oneCount >= half+(len(lines)%2)

		var filteredLines []string
		for _, line := range lines {
			if (keepMajority && ((oneMajority && line[i] == '1') || (!oneMajority && line[i] == '0'))) ||
				(!keepMajority && ((oneMajority && line[i] == '0') || (!oneMajority && line[i] == '1'))) {
				filteredLines = append(filteredLines, line)
			}
		}
		lines = filteredLines
	}

	return strconv.ParseInt(lines[0], 2, 64)
}

func SolvePart2(input string) (int64, error) {
	lines := strings.Split(input, "\n")
	oxygen, err := whittle(lines, true)
	if err != nil {
		return 0, err
	}
	co, err := whittle(lines, false)
	if err != nil {
		return 0, err
	}
	return oxygen * co, nil
}
