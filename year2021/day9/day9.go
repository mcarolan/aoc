package day9

import (
	"aoc/util"
	"slices"
)

func SolvePart1(input string) int {
	grid := util.ParseGrid(util.Lines(input), util.RuneToInt)

	riskIndexes := make([]int, 0)

	for cell := range grid.Iterator() {
		neighbours := grid.Neighbours(cell.Pos, false)

		if util.ForAll(neighbours, func(neighbour util.Cell[int]) bool { return neighbour.Value > cell.Value }) {
			riskIndexes = append(riskIndexes, cell.Value+1)
		}
	}

	return util.Sum(riskIndexes)
}

func SolvePart2(input string) int {
	grid := util.ParseGrid(util.Lines(input), util.RuneToInt)

	lowPoints := make([]util.RowCol, 0)

	for cell := range grid.Iterator() {
		neighbours := grid.Neighbours(cell.Pos, false)

		if util.ForAll(neighbours, func(neighbour util.Cell[int]) bool { return neighbour.Value > cell.Value }) {
			lowPoints = append(lowPoints, cell.Pos)
		}
	}

	basinSizes := make([]int, 0)
	for _, lowPoint := range lowPoints {
		basin := util.BFS(lowPoint, func(pos util.RowCol) []util.RowCol {
			neighbours := grid.Neighbours(pos, false)
			return util.Map(util.Filter(neighbours, func(cell util.Cell[int]) bool {
				return cell.Value != 9 && grid.At(pos) < cell.Value
			}), func(cell util.Cell[int]) util.RowCol {
				return cell.Pos
			})
		})

		basinSizes = append(basinSizes, len(basin))
	}

	slices.SortFunc(basinSizes, func(a, b int) int { return b - a })

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}
