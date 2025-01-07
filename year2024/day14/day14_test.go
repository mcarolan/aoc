package day14

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`
	expected := 12

	result := SolvePart1(input, 11, 7)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSingleRobot(t *testing.T) {
	input := `p=2,4 v=2,-3`
	expected := 12

	result := SolvePart1(input, 11, 7)

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
