package day2

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	expected := 150

	result, err := SolvePart1(input)

	if err != nil {
		t.Error(err)
	}

	if expected != result {
		t.Errorf("Expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	expected := 900

	result, err := SolvePart2(input)

	if err != nil {
		t.Error(err)
	}

	if expected != result {
		t.Errorf("Expected %d actual %d", expected, result)
	}
}
