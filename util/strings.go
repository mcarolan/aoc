package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RuneToInt(r rune) int {
	if r < '0' || r > '9' {
		panic(fmt.Sprintf(`invalid rune for digit %c`, r))
	}
	return int(r - '0')
}

func Identity[T any](value T) T {
	return value
}

func Lines(s string) []string {
	return strings.Split(s, "\n")
}

var numberRegex = regexp.MustCompile(`(-?\d)+`)

func AllNumbers(line string) []int {
	matches := numberRegex.FindAllStringSubmatch(line, -1)

	return Map(matches, func(match []string) int {
		i, err := strconv.Atoi(match[1])

		if err != nil {
			panic(err)
		}

		return i
	})
}
