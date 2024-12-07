package util

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
		for rowIndex, row := range grid.cells {
			for colIndex, value := range row {
				ch <- Cell[T]{Pos: RowCol{Row: rowIndex, Col: colIndex}, Value: value}
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
	var zero T
	if rowCol.Row >= len(grid.cells) {
		return zero, false
	}

	if rowCol.Col >= len(grid.cells[rowCol.Row]) {
		return zero, false
	}
	return grid.cells[rowCol.Row][rowCol.Col], true
}

func (grid *Grid[T]) Set(rowCol RowCol, value T) {
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
