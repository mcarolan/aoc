package day6

import "testing"

func TestSolvePart1(t *testing.T) {
	input := "3,4,3,1,2"

	expected := 5934

	result, err := SolvePart1(input)

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := "3,4,3,1,2"

	expected := 26984457539

	result, err := SolvePart2(input)

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}
