package day11

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `125 17`
	expected := 55312

	result := SolvePart1(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := ``
	expected := 5

	result := SolvePart2(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}
