package day12

import (
	"aoc/util"
	"fmt"
)

func floodFill(grid *util.Grid[rune], start util.Cell[rune]) []util.RowCol {
	q := make([]util.Cell[rune], 0)
	q = append(q, start)

	result := make([]util.RowCol, 0)
	visited := make(map[util.RowCol]bool)

	for len(q) > 0 {
		next := q[0]
		q = q[1:]

		if visited[next.Pos] {
			continue
		}
		visited[next.Pos] = true

		result = append(result, next.Pos)

		neighbours := grid.Neighbours(next.Pos, false)

		for _, neighbour := range neighbours {
			if neighbour.Value == start.Value {
				q = append(q, neighbour)
			}
		}
	}

	return result
}

func area(region []util.RowCol) int {
	return len(region)
}

func perimeter(region []util.RowCol) int {
	regionMap := make(map[util.RowCol]bool)
	for _, p := range region {
		regionMap[p] = true
	}

	result := 0

	for _, p := range region {
		above := util.RowCol{Row: p.Row - 1, Col: p.Col}
		below := util.RowCol{Row: p.Row + 1, Col: p.Col}
		left := util.RowCol{Row: p.Row, Col: p.Col - 1}
		right := util.RowCol{Row: p.Row, Col: p.Col + 1}

		if !regionMap[above] {
			result++
		}
		if !regionMap[below] {
			result++
		}
		if !regionMap[left] {
			result++
		}
		if !regionMap[right] {
			result++
		}
	}

	return result
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

func above(p util.RowCol) util.RowCol {
	return util.RowCol{Row: p.Row - 1, Col: p.Col}
}

func below(p util.RowCol) util.RowCol {
	return util.RowCol{Row: p.Row + 1, Col: p.Col}
}

func left(p util.RowCol) util.RowCol {
	return util.RowCol{Row: p.Row, Col: p.Col - 1}
}

func right(p util.RowCol) util.RowCol {
	return util.RowCol{Row: p.Row, Col: p.Col + 1}
}

func direction(p util.RowCol, d Direction) *util.RowCol {
	var res util.RowCol
	switch d {
	case UP:
		res = above(p)
	case DOWN:
		res = below(p)
	case LEFT:
		res = left(p)
	case RIGHT:
		res = right(p)
	default:
		return nil
	}

	return &res
}

func directionToString(dir Direction) string {
	switch dir {
	case UP:
		return "up"
	case DOWN:
		return "down"
	case LEFT:
		return "left"
	case RIGHT:
		return "right"
	}

	return "unknown"
}

func clockwise(dir Direction) Direction {
	switch dir {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}
	return dir
}

func antiClockwise(dir Direction) Direction {
	switch dir {
	case UP:
		return LEFT
	case RIGHT:
		return UP
	case DOWN:
		return RIGHT
	case LEFT:
		return DOWN
	}
	return dir
}

func sides(start util.RowCol, region []util.RowCol) int {
	result := 0
	regionMap := make(map[util.RowCol]bool)
	for _, p := range region {
		regionMap[p] = true
	}

	edgeCells := make(map[util.RowCol]bool)

	for _, p := range region {
		if !regionMap[above(p)] || !regionMap[below(p)] || !regionMap[left(p)] || !regionMap[right(p)] {
			edgeCells[p] = true
		}
	}

	currentDirection := RIGHT
	at := start

	visited := make(map[util.RowCol]map[Direction]bool)

	for at != start || len(visited) == 0 {
		if visited[at] == nil {
			visited[at] = make(map[Direction]bool)
		}
		if visited[at][currentDirection] {
			currentDirection = clockwise(currentDirection)
			continue
		}
		visited[at][currentDirection] = true

		next := direction(at, currentDirection)
		nextAnti := direction(at, antiClockwise(currentDirection))

		if nextAnti != nil && edgeCells[*nextAnti] {
			result++
			at = *nextAnti
			currentDirection = antiClockwise(currentDirection)
		} else if next != nil && edgeCells[*next] {
			at = *next
		} else {
			result++
			currentDirection = clockwise(currentDirection)
		}
	}

	return result - 1
}

func SolvePart1(input string) int {
	grid := util.ParseGrid(util.Lines(input), util.Identity)

	regions := make(map[util.RowCol][]util.RowCol)
	alreadyInRegion := make(map[util.RowCol]bool)

	for cell := range grid.Iterator() {
		if alreadyInRegion[cell.Pos] {
			continue
		}

		region := floodFill(&grid, cell)
		regions[cell.Pos] = region

		for _, c := range region {
			alreadyInRegion[c] = true
		}
	}

	result := 0

	for _, region := range regions {
		result += area(region) * perimeter(region)
	}

	return result
}

func SolvePart2(input string) int {
	grid := util.ParseGrid(util.Lines(input), util.Identity)

	regions := make(map[util.RowCol][]util.RowCol)
	alreadyInRegion := make(map[util.RowCol]bool)

	for cell := range grid.Iterator() {
		if alreadyInRegion[cell.Pos] {
			continue
		}

		region := floodFill(&grid, cell)
		regions[cell.Pos] = region

		for _, c := range region {
			alreadyInRegion[c] = true
		}
	}

	result := 0

	for start, region := range regions {
		value, _ := grid.At(start)
		fmt.Printf("checking %c starting at %v\n", value, start)
		sideValue := sides(start, region)
		fmt.Printf("%c has %d sides", value, sideValue)
		fmt.Println()
		result += area(region) * sideValue
	}

	return result
}
