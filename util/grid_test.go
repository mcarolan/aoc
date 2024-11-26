package util

import (
	"reflect"
	"testing"
)

func TestParseGridInt(t *testing.T) {
	numbers := `123
456
789`
	grid := ParseGrid(Lines(numbers), RuneToInt)

	if grid.Rows != 3 {
		t.Errorf("expected 3 rows actually got %d", grid.Rows)
	}

	if grid.ColsPerRow != 3 {
		t.Errorf("expected rows to have 3 cols, had %d", grid.ColsPerRow)
	}

	if grid.At(RowCol{Row: 2, Col: 1}) != 8 {
		t.Errorf("grid[2][1] to be 8, actually %d", grid.At(RowCol{Row: 2, Col: 1}))
	}
}

func TestNeighboursNoDiagonal(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	grid := ParseGrid(Lines(input), RuneToInt)
	neighbours := grid.Neighbours(RowCol{Row: 2, Col: 2}, false)
	expected := []int{8, 6, 6, 8}

	if !reflect.DeepEqual(neighbours, expected) {
		t.Errorf("expected %v actual %v", expected, neighbours)
	}
}

func TestNeighboursWithDiagonal(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	grid := ParseGrid(Lines(input), RuneToInt)
	neighbours := grid.Neighbours(RowCol{Row: 2, Col: 2}, true)
	expected := []int{8, 7, 6, 7, 6, 7, 8, 9}

	if !reflect.DeepEqual(neighbours, expected) {
		t.Errorf("expected %v actual %v", expected, neighbours)
	}
}

func TestNeighboursEdge(t *testing.T) {

	numbers := `123
456
789`
	grid := ParseGrid(Lines(numbers), RuneToInt)
	neighbours := grid.Neighbours(RowCol{Row: 0, Col: 2}, true)
	expected := []int{6, 5, 2}

	if !reflect.DeepEqual(neighbours, expected) {
		t.Errorf("expected %v actual %v", expected, neighbours)
	}
}
