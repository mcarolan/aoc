package util

import "fmt"

type RowCol struct {
	Row int
	Col int
}

type Grid[T any] struct {
	cells      map[int]map[int]T
	ColsPerRow int
	Rows       int
}

type Cell[T any] struct {
	Pos   RowCol
	Value T
}

func (grid Grid[T]) Iterator() <-chan Cell[T] {
	ch := make(chan Cell[T])

	go func() {
		for rowIndex := range len(grid.cells) {
			row := grid.cells[rowIndex]
			for colIndex := range len(row) {
				cell := row[colIndex]
				ch <- Cell[T]{Pos: RowCol{Row: rowIndex, Col: colIndex}, Value: cell}
			}
		}
		close(ch)
	}()

	return ch
}

func ParseGrid[T any](lines []string, cellParser func(rune) T) Grid[T] {
	result := make(map[int]map[int]T)
	for rowIndex, row := range lines {
		result[rowIndex] = make(map[int]T)
		for colIndex, cellRune := range row {
			cell := cellParser(cellRune)
			result[rowIndex][colIndex] = cell
		}
	}
	return Grid[T]{
		cells:      result,
		ColsPerRow: len(result[0]),
		Rows:       len(result),
	}
}

func NewGrid[T any](colsPerRow int, rows int) Grid[T] {
	return Grid[T]{
		cells:      make(map[int]map[int]T),
		ColsPerRow: colsPerRow,
		Rows:       rows,
	}
}

func (grid Grid[rune]) Print() {
	for y := range grid.Rows {
		for x := range grid.ColsPerRow {
			pos := RowCol{Row: y, Col: x}
			value, _ := grid.At(pos)
			fmt.Printf("%c", value)
		}
		fmt.Println()
	}
}

func (grid Grid[T]) Neighbours(rowCol RowCol, includeDiagonal bool) []Cell[T] {
	var indexes []RowCol

	if includeDiagonal {
		indexes = []RowCol{{Row: -1, Col: 0}, {Row: -1, Col: 1}, {Row: 0, Col: 1}, {Row: 1, Col: 1}, {Row: 1, Col: 0}, {Row: 1, Col: -1}, {Row: 0, Col: -1}, {Row: -1, Col: -1}}
	} else {
		indexes = []RowCol{{Row: -1, Col: 0}, {Row: 0, Col: 1}, {Row: 1, Col: 0}, {Row: 0, Col: -1}}
	}

	result := make([]Cell[T], 0)

	for _, i := range indexes {
		neighbourRowIndex := rowCol.Row + i.Row

		if row, exists := grid.cells[neighbourRowIndex]; exists {
			neighbourColIndex := rowCol.Col + i.Col
			cellValue, hasCell := row[neighbourColIndex]
			if hasCell {
				result = append(result, Cell[T]{Pos: RowCol{Row: neighbourRowIndex, Col: neighbourColIndex}, Value: cellValue})
			}
		}

	}

	return result
}

func (grid *Grid[T]) At(rowCol RowCol) (T, bool) {
	if !grid.InBounds(rowCol) {
		var zero T
		return zero, false
	}
	return grid.cells[rowCol.Row][rowCol.Col], true
}

func (grid *Grid[T]) InBounds(rowCol RowCol) bool {
	return rowCol.Row < grid.Rows && rowCol.Row >= 0 && rowCol.Col < grid.ColsPerRow && rowCol.Col >= 0
}

func (grid *Grid[T]) Set(rowCol RowCol, value T) {
	if grid.cells[rowCol.Row] == nil {
		grid.cells[rowCol.Row] = make(map[int]T, grid.ColsPerRow)
	}
	grid.cells[rowCol.Row][rowCol.Col] = value
}

func BFS(start RowCol, neighbourFunc func(RowCol) []RowCol) []RowCol {
	queue := []RowCol{start}
	visited := make(map[RowCol]bool, 0)

	result := make([]RowCol, 0)
	for {
		if len(queue) == 0 {
			break
		}

		at := queue[0]
		queue = queue[1:]

		_, alreadyVisited := visited[at]
		if alreadyVisited {
			continue
		}
		visited[at] = true

		neighbours := neighbourFunc(at)

		result = append(result, at)
		queue = append(queue, neighbours...)
	}
	return result
}

func Manhattan(a RowCol, b RowCol) int {
	return Abs(a.Row-b.Row) + Abs(a.Col-b.Col)
}

func (pos RowCol) Neighbour(direction Direction) RowCol {
	switch direction {
	case North:
		return RowCol{Row: pos.Row - 1, Col: pos.Col}
	case East:
		return RowCol{Row: pos.Row, Col: pos.Col + 1}
	case South:
		return RowCol{Row: pos.Row + 1, Col: pos.Col}
	case West:
		return RowCol{Row: pos.Row, Col: pos.Col - 1}
	}

	panic(fmt.Sprintf("Invalid direction %d", direction))
}
