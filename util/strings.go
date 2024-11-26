package util

import (
	"fmt"
	"strings"
)

func RuneToInt(r rune) int {
	if r < '0' || r > '9' {
		panic(fmt.Sprintf(`invalid rune for digit %c`, r))
	}
	return int(r - '0')
}

func Lines(s string) []string {
	return strings.Split(s, "\n")
}
