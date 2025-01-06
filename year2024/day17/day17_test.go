package day17

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`
	expected := "4,6,3,5,6,3,5,2,1,0"

	result := SolvePart1(input)

	if result != expected {
		t.Errorf("expected %s actual %s", expected, result)
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
