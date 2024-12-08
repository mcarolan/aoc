package day6

import (
	"aoc/util"
)

type Facing int

const (
	North Facing = iota
	East
	South
	West
)

func findStartRowCol(grid *util.Grid[rune]) *util.RowCol {
	for cell := range grid.Iterator() {
		if cell.Value == '^' {
			return &cell.Pos
		}
	}

	return nil
}

func nextPosition(pos util.RowCol, facing Facing) util.RowCol {
	switch facing {
	case North:
		return util.RowCol{Row: pos.Row - 1, Col: pos.Col}
	case East:
		return util.RowCol{Row: pos.Row, Col: pos.Col + 1}
	case South:
		return util.RowCol{Row: pos.Row + 1, Col: pos.Col}
	case West:
		return util.RowCol{Row: pos.Row, Col: pos.Col - 1}
	}

	panic("bad facing value")
}

func clockwise(facing Facing) Facing {
	return (facing + 1) % 4
}

func SolvePart1(input string) int {
	grid := util.ParseGrid[rune](util.Lines(input), func(r rune) rune { return r })
	guardPos := findStartRowCol(&grid)
	facing := North

	if guardPos == nil {
		panic("could not find the guard's starting position!")
	}

	visited := make(map[util.RowCol]bool)

	for {
		visited[*guardPos] = true
		nextPos := nextPosition(*guardPos, facing)
		next, nextExists := grid.At(nextPos)

		if !nextExists {
			break
		}

		if next == '#' {
			facing = clockwise(facing)
		} else {
			guardPos = &nextPos
		}
	}

	return len(visited)
}

func isLoop(grid *util.Grid[rune], guardPos util.RowCol, obstaclePos util.RowCol) bool {
	facing := North
	visited := make(map[util.RowCol]map[Facing]bool)

	for {
		_, currentPosExists := grid.At(guardPos)

		if !currentPosExists {
			return false
		}

		visitedFacing, hasVisitedCell := visited[guardPos]

		if hasVisitedCell && visitedFacing[facing] {
			return true
		}

		if !hasVisitedCell {
			visited[guardPos] = make(map[Facing]bool)
		}

		visited[guardPos][facing] = true

		nextPos := nextPosition(guardPos, facing)
		next, nextExists := grid.At(nextPos)

		if !nextExists {
			return false
		}

		if nextPos == obstaclePos || next == '#' {
			facing = clockwise(facing)
		} else {
			guardPos = nextPos
		}
	}
}

func SolvePart2(input string) int {
	grid := util.ParseGrid[rune](util.Lines(input), func(r rune) rune { return r })
	guardPos := findStartRowCol(&grid)
	if guardPos == nil {
		panic("could not find the guard's starting position!")
	}

	counter := 0

	for cell := range grid.Iterator() {
		if cell.Value != '.' {
			continue
		}

		if isLoop(&grid, *guardPos, cell.Pos) {
			counter++
		}
	}

	return counter
}
