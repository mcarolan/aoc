package day1

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 11

	result := SolvePart1(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 31

	result := SolvePart2(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}
