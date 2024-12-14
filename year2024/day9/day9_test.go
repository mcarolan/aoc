package day9

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `2333133121414131402`
	expected := 1928

	result := SolvePart1(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `2333133121414131402`
	expected := 2858

	result := SolvePart2(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}
