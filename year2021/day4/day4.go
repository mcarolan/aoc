package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

var drawRegex = regexp.MustCompile(`([0-9]+),?`)

const BOARD_SIZE = 5

func parseDraws(line string) ([]int, error) {
	return extractNInts(line, drawRegex, -1)
}

var entryregex = regexp.MustCompile(`\s+([0-9]+)`)

func parseBoard(lines string) ([]int, error) {
	return extractNInts(lines, entryregex, BOARD_SIZE*BOARD_SIZE)
}

func parseBoards(lines []string) ([][]int, error) {
	var result [][]int
	offset := 0

	for {
		board, err := parseBoard(strings.Join(lines[offset:], "\n"))
		if err != nil {
			return nil, err
		}
		if board == nil {
			break
		}

		result = append(result, board)
		offset += BOARD_SIZE + 1
	}

	return result, nil
}

type board struct {
	entries    []int
	rowsMarked []int
	colsMarked []int
}

const MARKED = -1

func mark(n int, b *board) bool {
	for i, entry := range b.entries {
		if entry == n {
			b.entries[i] = MARKED
			col := i % BOARD_SIZE
			row := i / BOARD_SIZE
			b.rowsMarked[row] += 1
			b.colsMarked[col] += 1
			return b.rowsMarked[row] == BOARD_SIZE || b.colsMarked[col] == BOARD_SIZE
		}
	}

	return false
}

func parseInput(input string) ([]int, []board, error) {
	lines := strings.Split(input, "\n")

	draws, err := parseDraws(lines[0])
	if err != nil {
		return nil, nil, err
	}

	rawBoards, err := parseBoards(lines[1:])

	if err != nil {
		return nil, nil, err
	}

	var boards []board
	for _, b := range rawBoards {
		boards = append(boards, board{entries: b, rowsMarked: make([]int, BOARD_SIZE), colsMarked: make([]int, BOARD_SIZE)})
	}
	return draws, boards, nil
}

func SolvePart1(input string) (int, error) {
	draws, boards, err := parseInput(input)

	if err != nil {
		return 0, err
	}

	for _, call := range draws {
		for _, board := range boards {
			if mark(call, &board) {
				var sumUnmarked int

				for _, e := range board.entries {
					if e != MARKED {
						sumUnmarked += e
					}
				}

				return sumUnmarked * call, nil
			}
		}
	}

	return 0, fmt.Errorf("no winning board found")
}

func SolvePart2(input string) (int, error) {
	draws, boards, err := parseInput(input)

	if err != nil {
		return 0, err
	}

	var completeBoards = make(map[int]bool)
	lastWinningScore := 0

	for _, call := range draws {
		for i, board := range boards {
			if completeBoards[i] {
				continue
			}
			if mark(call, &board) {
				completeBoards[i] = true
				var sumUnmarked int

				for _, e := range board.entries {
					if e != MARKED {
						sumUnmarked += e
					}
				}

				lastWinningScore = sumUnmarked * call
			}
		}
	}

	return lastWinningScore, nil
}
