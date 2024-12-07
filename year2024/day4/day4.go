package day4

import "aoc/util"

var SEARCH_PATTERNS = [...][4]util.RowCol{
	{util.RowCol{Row: 0, Col: 0}, util.RowCol{Row: 0, Col: 1}, util.RowCol{Row: 0, Col: 2}, util.RowCol{Row: 0, Col: 3}},    //right
	{util.RowCol{Row: 0, Col: 0}, util.RowCol{Row: 0, Col: -1}, util.RowCol{Row: 0, Col: -2}, util.RowCol{Row: 0, Col: -3}}, //left
	{util.RowCol{Row: 0, Col: 0}, util.RowCol{Row: -1, Col: 0}, util.RowCol{Row: -2, Col: 0}, util.RowCol{Row: -3, Col: 0}}, //above
	{util.RowCol{Row: 0, Col: 0}, util.RowCol{Row: 1, Col: 0}, util.RowCol{Row: 2, Col: 0}, util.RowCol{Row: 3, Col: 0}},    //below

	{util.RowCol{Row: 0, Col: 0}, util.RowCol{Row: -1, Col: 1}, util.RowCol{Row: -2, Col: 2}, util.RowCol{Row: -3, Col: 3}}, // diag ne
	{util.RowCol{Row: 0, Col: 0}, util.RowCol{Row: 1, Col: 1}, util.RowCol{Row: 2, Col: 2}, util.RowCol{Row: 3, Col: 3}},    // diag se

	{util.RowCol{Row: 0, Col: 0}, util.RowCol{Row: 1, Col: -1}, util.RowCol{Row: 2, Col: -2}, util.RowCol{Row: 3, Col: -3}},    //diag nw
	{util.RowCol{Row: 0, Col: 0}, util.RowCol{Row: -1, Col: -1}, util.RowCol{Row: -2, Col: -2}, util.RowCol{Row: -3, Col: -3}}, //diag sw
}

var TARGET_WORD = [...]rune{'X', 'M', 'A', 'S'}

func numberOfMatchesPart1(origin util.RowCol, grid *util.Grid[rune]) int {
	counter := 0

	for _, offset := range SEARCH_PATTERNS {
		match := true

		for i, letterOffset := range offset {
			pos := util.RowCol{Row: origin.Row + letterOffset.Row, Col: origin.Col + letterOffset.Col}
			char, found := grid.At(pos)

			if !found || TARGET_WORD[i] != char {
				match = false
				break
			}
		}

		if match {
			counter++
		}
	}

	return counter
}

func isMatchPart2(origin util.RowCol, grid *util.Grid[rune]) bool {
	ne, found := grid.At(util.RowCol{Row: origin.Row - 1, Col: origin.Col + 1})
	if !found {
		return false
	}

	se, found := grid.At(util.RowCol{Row: origin.Row + 1, Col: origin.Col + 1})
	if !found {
		return false
	}
	nw, found := grid.At(util.RowCol{Row: origin.Row - 1, Col: origin.Col - 1})
	if !found {
		return false
	}

	sw, found := grid.At(util.RowCol{Row: origin.Row + 1, Col: origin.Col - 1})
	if !found {
		return false
	}

	return ((ne == 'S' && sw == 'M') || (ne == 'M' && sw == 'S')) && ((nw == 'S' && se == 'M') || (nw == 'M' && se == 'S'))
}

func SolvePart1(input string) int {
	grid := util.ParseGrid(util.Lines(input), func(r rune) rune { return r })
	counter := 0

	for cell := range grid.Iterator() {
		if cell.Value == 'X' {
			counter += numberOfMatchesPart1(cell.Pos, &grid)
		}
	}

	return counter
}

func SolvePart2(input string) int {
	grid := util.ParseGrid(util.Lines(input), func(r rune) rune { return r })
	counter := 0

	for cell := range grid.Iterator() {
		if cell.Value == 'A' && isMatchPart2(cell.Pos, &grid) {
			counter++
		}
	}

	return counter
}
