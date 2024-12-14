package day10

import (
	"aoc/util"
	"fmt"
	"strings"
)

func countTrailheads(grid *util.Grid[int], startPosition util.RowCol) int {
	counter := 0
	q := make([]util.RowCol, 0)
	q = append(q, startPosition)
	visited := make(map[util.RowCol]bool)

	for len(q) > 0 {
		at := q[0]
		q = q[1:]

		_, alreadyVisited := visited[at]

		if alreadyVisited {
			continue
		}
		visited[at] = true

		value, _ := grid.At(at)

		if value == 9 {
			counter++
			continue
		}

		neighbours := grid.Neighbours(at, false)

		for _, n := range neighbours {
			if n.Value == value+1 {
				q = append(q, n.Pos)
			}
		}
	}

	return counter
}

func SolvePart1(input string) int {
	grid := util.ParseGrid(util.Lines(input), util.RuneToInt)

	result := 0

	for cell := range grid.Iterator() {
		if cell.Value != 0 {
			continue
		}
		result += countTrailheads(&grid, cell.Pos)
	}

	return result
}

func keyForPath(path []util.RowCol) string {
	var sb strings.Builder
	for _, rc := range path {
		fmt.Fprintf(&sb, "%v;", rc)
	}
	return sb.String()
}

func countTrailheadPaths(grid *util.Grid[int], startPosition util.RowCol) int {
	counter := 0
	q := make([]util.RowCol, 0)
	q = append(q, startPosition)

	for len(q) > 0 {
		at := q[0]
		q = q[1:]

		value, _ := grid.At(at)

		if value == 9 {
			counter++
			continue
		}

		neighbours := grid.Neighbours(at, false)

		for _, n := range neighbours {
			if n.Value == value+1 {
				q = append(q, n.Pos)
			}
		}
	}

	return counter
}

func SolvePart2(input string) int {
	grid := util.ParseGrid(util.Lines(input), util.RuneToInt)

	result := 0

	for cell := range grid.Iterator() {
		if cell.Value != 0 {
			continue
		}
		result += countTrailheadPaths(&grid, cell.Pos)
	}

	return result
}
