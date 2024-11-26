package day3

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	result := SolvePart1(input)

	expected := int64(198)

	if expected != result {
		t.Errorf("actual %d != expected %d", result, expected)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	result, err := SolvePart2(input)

	if err != nil {
		t.Error(err)
	}

	expected := int64(230)

	if expected != result {
		t.Errorf("actual %d != expected %d", result, expected)
	}
}
