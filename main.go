package main

import (
	"aoc/year2024/day3"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("year2024/inputs/day3.txt")

	if err != nil {
		fmt.Printf("Failed to read input: %s", err)
		os.Exit(1)
	}

	result := day3.SolvePart2(string(input))

	if err != nil {
		fmt.Printf("Failed to solve: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("the result is %d\n", result)
}
