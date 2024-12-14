package day10

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	expected := 36

	result := SolvePart1(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`
	expected := 81

	result := SolvePart2(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}
