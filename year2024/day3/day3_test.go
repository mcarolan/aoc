package day3

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^
do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	expected := 161

	result := SolvePart1(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	expected := 48

	result := SolvePart2(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}
