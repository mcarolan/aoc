package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

var lineRegex = regexp.MustCompile(`([0-9]+),([0-9]+) -> ([0-9]+),([0-9]+)`)

func parseInput(input string) ([]line, error) {
	var result []line
	lines := strings.Split(input, "\n")

	for lineNumber, lineString := range lines {
		match := lineRegex.FindStringSubmatch(lineString)

		if match == nil {
			return nil, fmt.Errorf("no match on line %d", lineNumber)
		}

		var numbers []int

		for i := 1; i < len(match); i++ {
			number, err := strconv.Atoi(match[i])

			if err != nil {
				return nil, fmt.Errorf("failed to parse '%s' as int on line %d", match[i], lineNumber)
			}

			numbers = append(numbers, number)
		}

		result = append(result, line{start: point{x: numbers[0], y: numbers[1]}, end: point{x: numbers[2], y: numbers[3]}})
	}
	return result, nil
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func SolvePart1(input string) (int, error) {
	lines, err := parseInput(input)

	if err != nil {
		return 0, err
	}

	var grid = make(map[point]int)

	for _, line := range lines {
		if line.start.x != line.end.x && line.start.y != line.end.y {
			continue
		}

		if line.start.x == line.end.x {
			for y := min(line.start.y, line.end.y); y <= max(line.start.y, line.end.y); y++ {
				grid[point{x: line.start.x, y: y}]++
			}
		} else if line.start.y == line.end.y {
			for x := min(line.start.x, line.end.x); x <= max(line.start.x, line.end.x); x++ {
				grid[point{x: x, y: line.start.y}]++
			}
		}
	}

	result := 0

	for _, p := range grid {
		if p >= 2 {
			result++
		}
	}

	return result, nil
}

func SolvePart2(input string) (int, error) {
	lines, err := parseInput(input)

	if err != nil {
		return 0, err
	}

	var grid = make(map[point]int)

	for _, line := range lines {
		xChange := 0
		if line.end.x > line.start.x {
			xChange = 1
		} else if line.end.x < line.start.x {
			xChange = -1
		}

		yChange := 0
		if line.end.y > line.start.y {
			yChange = 1
		} else if line.end.y < line.start.y {
			yChange = -1
		}

		manhattan := max(abs(line.start.x-line.end.x), abs(line.start.y-line.end.y))

		for d := range manhattan + 1 {
			grid[point{x: line.start.x + (d * xChange), y: line.start.y + (d * yChange)}]++
		}
	}

	result := 0

	for _, p := range grid {
		if p >= 2 {
			result++
		}
	}

	return result, nil
}
