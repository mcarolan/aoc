package day7

import (
	"aoc/util"
	"fmt"
	"strconv"
)

func concat(a int, b int) int {
	res, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	if err != nil {
		panic(err)
	}

	return res
}

func hasSolutionPart1(target int, current int, remaining []int) bool {
	if len(remaining) == 0 {
		return current == target
	} else {
		next := remaining[0]
		nextRemaining := remaining[1:]

		return hasSolutionPart1(target, current*next, nextRemaining) || hasSolutionPart1(target, current+next, nextRemaining)
	}
}

func hasSolutionPart2(target int, current int, remaining []int) bool {
	if len(remaining) == 0 {
		return current == target
	} else {
		next := remaining[0]
		nextRemaining := remaining[1:]

		return hasSolutionPart2(target, current*next, nextRemaining) || hasSolutionPart2(target, current+next, nextRemaining) || hasSolutionPart2(target, concat(current, next), nextRemaining)
	}
}

func run(input string, f func(int, int, []int) bool) int {
	equations := util.Map(util.Lines(input), util.AllNumbers)

	result := 0

	for _, equation := range equations {
		target := equation[0]
		operands := equation[1:]
		current := operands[0]
		remaining := operands[1:]

		if f(target, current, remaining) {
			result += target
		}
	}

	return result
}

func SolvePart1(input string) int {
	return run(input, hasSolutionPart1)
}

func SolvePart2(input string) int {
	return run(input, hasSolutionPart2)
}
