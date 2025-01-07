package day15

import (
	"aoc/util"
	"fmt"
	"strings"
)

func translate(position util.RowCol, instruction rune) util.RowCol {
	switch instruction {
	case '<':
		return util.RowCol{Row: position.Row, Col: position.Col - 1}
	case '>':
		return util.RowCol{Row: position.Row, Col: position.Col + 1}
	case '^':
		return util.RowCol{Row: position.Row - 1, Col: position.Col}
	case 'v':
		return util.RowCol{Row: position.Row + 1, Col: position.Col}
	}

	panic(fmt.Sprintf("unknown instruction %c", instruction))
}

func swap(grid util.Grid[rune], a util.RowCol, b util.RowCol) {
	aValue, _ := grid.At(a)
	bValue, _ := grid.At(b)
	grid.Set(b, aValue)
	grid.Set(a, bValue)
}

func findFirstDifferent(grid util.Grid[rune], position util.RowCol, instruction rune) util.Cell[rune] {
	value, _ := grid.At(position)
	position = translate(position, instruction)
	var valueFound rune

	for {
		nextValue, isValid := grid.At(position)
		if !isValid || nextValue != value {
			valueFound = nextValue
			break
		}
		position = translate(position, instruction)
	}

	return util.Cell[rune]{Pos: position, Value: valueFound}
}

func step(grid util.Grid[rune], currentPosition util.RowCol, instruction rune) util.RowCol {
	nextPosition := translate(currentPosition, instruction)
	value, isValid := grid.At(nextPosition)

	if !isValid || value == '#' {
		return currentPosition
	}

	if value == '.' {
		swap(grid, nextPosition, currentPosition)
		return nextPosition
	} else if value == 'O' {
		firstDifferent := findFirstDifferent(grid, nextPosition, instruction)

		if firstDifferent.Value == '.' {
			swap(grid, nextPosition, firstDifferent.Pos)
			swap(grid, nextPosition, currentPosition)
			return nextPosition
		} else {
			return currentPosition
		}
	} else {
		panic(fmt.Sprintf("unknown cell value %c", value))
	}
}

func SolvePart1(input string) int {
	var parts = strings.Split(input, "\n\n")
	gridRunes := parts[0]
	instructions := parts[1]

	grid := util.ParseGrid[rune](util.Lines(gridRunes), util.Identity)

	var robotPosition *util.RowCol

	for cell := range grid.Iterator() {
		if cell.Value == '@' {
			robotPosition = &cell.Pos
			break
		}
	}

	if robotPosition == nil {
		panic("could not find robot!")
	}

	for _, instruction := range instructions {
		if instruction == '\n' {
			continue
		}

		nextRobotPosition := step(grid, *robotPosition, instruction)
		robotPosition = &nextRobotPosition
	}

	result := 0

	for cell := range grid.Iterator() {
		if cell.Value != 'O' {
			continue
		}

		result += (100 * cell.Pos.Row) + cell.Pos.Col
	}

	return result
}

func SolvePart2(input string) int {
	return 42
}
