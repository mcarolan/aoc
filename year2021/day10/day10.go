package day10

import (
	"aoc/util"
	"slices"
)

func firstInvalidCharacter(input string) (rune, bool) {
	stack := util.NewStack[rune]()

	for _, r := range input {
		switch r {
		case '(':
			stack.Push(')')
		case '[':
			stack.Push(']')
		case '{':
			stack.Push('}')
		case '<':
			stack.Push('>')
		default:
			found := stack.Pop()

			if found != r {
				return r, true
			}
		}
	}

	return rune(0), false
}

func SolvePart1(input string) int {
	lines := util.Lines(input)

	score := 0

	for _, line := range lines {
		invalidCharacter, hasInvalidCharacter := firstInvalidCharacter(line)

		if !hasInvalidCharacter {
			continue
		}

		switch invalidCharacter {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		}
	}
	return score
}

func findAutoCompleteScore(input string) (int, bool) {
	stack := util.NewStack[rune]()

	for _, r := range input {
		switch r {
		case '(':
			stack.Push(')')
		case '[':
			stack.Push(']')
		case '{':
			stack.Push('}')
		case '<':
			stack.Push('>')
		default:
			found := stack.Pop()

			if found != r {
				return 0, false // corrupt
			}
		}
	}

	score := 0

	for !stack.IsEmpty() {
		score *= 5

		switch stack.Pop() {
		case ')':
			score += 1
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}

	return score, true
}

func SolvePart2(input string) int {
	lines := util.Lines(input)

	scores := make([]int, 0)

	for _, line := range lines {
		autoCompleteScore, isIncomplete := findAutoCompleteScore(line)

		if !isIncomplete {
			continue
		}
		scores = append(scores, autoCompleteScore)
	}

	slices.Sort(scores)
	return scores[len(scores)/2]
}
