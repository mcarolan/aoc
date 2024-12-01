package day12

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
	expected := 10

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
