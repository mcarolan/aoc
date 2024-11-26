package day1

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `199
200
208
210
200
207
240
269
260
263`

	expected := 7
	result, err := SolvePart1(input)

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `199
200
208
210
200
207
240
269
260
263`

	expected := 5
	result, err := SolvePart2(input)

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("Expected %d actual %d", expected, result)
	}
}
