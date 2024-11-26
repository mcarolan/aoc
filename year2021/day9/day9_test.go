package day9

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`
	expected := 15

	result := SolvePart1(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`
	expected := 1134

	result := SolvePart2(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}
