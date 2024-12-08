package day8

import (
	"aoc/util"
)

func SolvePart1(input string) int {
	grid := util.ParseGrid(util.Lines(input), func(r rune) rune { return r })
	nodePosition := make(map[rune][]util.RowCol)

	for cell := range grid.Iterator() {
		if cell.Value == '.' {
			continue
		}

		posList, hasPosList := nodePosition[cell.Value]
		if !hasPosList {
			posList = make([]util.RowCol, 0)
		}
		nodePosition[cell.Value] = append(posList, cell.Pos)
	}

	antiNodes := make(map[util.RowCol]bool)

	for node, positions := range nodePosition {
		for pair := range util.Pairs(positions) {
			a := pair[0]
			b := pair[1]

			deltaRow := b.Row - a.Row
			deltaCol := b.Col - a.Col

			antiNode1Pos := util.RowCol{Row: a.Row - deltaRow, Col: a.Col - deltaCol}
			antiNode2Pos := util.RowCol{Row: b.Row + deltaRow, Col: b.Col + deltaCol}

			valueAt, inBounds := grid.At(antiNode1Pos)
			if inBounds && valueAt != node {
				antiNodes[antiNode1Pos] = true
			}

			valueAt, inBounds = grid.At(antiNode2Pos)
			if inBounds && valueAt != node {
				antiNodes[antiNode2Pos] = true
			}
		}
	}

	return len(antiNodes)
}

func SolvePart2(input string) int {
	grid := util.ParseGrid(util.Lines(input), func(r rune) rune { return r })
	nodePosition := make(map[rune][]util.RowCol)

	for cell := range grid.Iterator() {
		if cell.Value == '.' {
			continue
		}

		posList, hasPosList := nodePosition[cell.Value]
		if !hasPosList {
			posList = make([]util.RowCol, 0)
		}
		nodePosition[cell.Value] = append(posList, cell.Pos)
	}

	antiNodes := make(map[util.RowCol]bool)

	for _, positions := range nodePosition {
		for pair := range util.Pairs(positions) {
			a := pair[0]
			b := pair[1]

			deltaRow := b.Row - a.Row
			deltaCol := b.Col - a.Col

			for i := 0; ; i++ {
				antiNode1Pos := util.RowCol{Row: a.Row - deltaRow*i, Col: a.Col - deltaCol*i}
				antiNode2Pos := util.RowCol{Row: b.Row + deltaRow*i, Col: b.Col + deltaCol*i}

				_, exists1 := grid.At(antiNode1Pos)
				_, exists2 := grid.At(antiNode2Pos)

				if !exists1 && !exists2 {
					break
				}

				if exists1 {
					antiNodes[antiNode1Pos] = true
				}
				if exists2 {
					antiNodes[antiNode2Pos] = true
				}
			}
		}
	}

	return len(antiNodes)
}
