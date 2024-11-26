package day7

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `16,1,2,0,4,2,7,1,2,14`

	result, err := SolvePart1(input)

	if err != nil {
		t.Error(err)
	}

	expected := 37

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `16,1,2,0,4,2,7,1,2,14`

	result, err := SolvePart2(input)

	if err != nil {
		t.Error(err)
	}

	expected := 168

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}
