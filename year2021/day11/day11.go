package day11

import (
	"aoc/util"
	"fmt"
)

const MAX_ENERGY = 9

func step(grid *util.Grid[int]) int {
	toFlash := make([]util.RowCol, 0)
	flashed := make(map[util.RowCol]bool)

	for cell := range grid.Iterator() {
		grid.Set(cell.Pos, cell.Value+1)
		if grid.At(cell.Pos) > MAX_ENERGY {
			toFlash = append(toFlash, cell.Pos)
		}
	}

	for len(toFlash) > 0 {
		next := toFlash[0]
		toFlash = toFlash[1:]

		_, hasAlreadyFlashed := flashed[next]
		if hasAlreadyFlashed {
			continue
		}
		flashed[next] = true

		neighbours := grid.Neighbours(next, true)
		for _, cell := range neighbours {
			grid.Set(cell.Pos, cell.Value+1)
			if grid.At(cell.Pos) > MAX_ENERGY {
				toFlash = append(toFlash, cell.Pos)
			}
		}
	}

	for rowCol := range flashed {
		grid.Set(rowCol, 0)
	}

	return len(flashed)
}

func PrintGrid(grid *util.Grid[int]) {
	for row := range grid.Rows {
		for col := range grid.ColsPerRow {
			value := grid.At(util.RowCol{Row: row, Col: col})
			if value == 0 {
				fmt.Print("\033[31m")
				fmt.Print("0")
				fmt.Print("\033[0m")
			} else {
				fmt.Printf("%d", value)
			}
		}
		fmt.Print("\n")
	}
}

func SolvePart1(input string) int {
	grid := util.ParseGrid[int](util.Lines(input), util.RuneToInt)
	totalFlashes := 0

	for range 100 {
		totalFlashes += step(&grid)
	}
	return totalFlashes
}

func SolvePart2(input string) int {
	grid := util.ParseGrid[int](util.Lines(input), util.RuneToInt)

	counter := 0
	for {
		counter++
		if step(&grid) == 100 {
			return counter
		}
	}
}
