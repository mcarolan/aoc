package day5

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	expected := 5

	result, err := SolvePart1(input)

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	expected := 12

	result, err := SolvePart2(input)

	if err != nil {
		t.Error(err)
	}

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}
