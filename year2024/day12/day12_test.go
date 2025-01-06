package day12

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	expected := 1930

	result := SolvePart1(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`
	expected := 236

	result := SolvePart2(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2b(t *testing.T) {
	input := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`
	expected := 368

	result := SolvePart2(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2c(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	expected := 1206

	result := SolvePart2(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}

func TestSolvePart2d(t *testing.T) {
	input := `RRRRIICCFF
			  RRRRIICCCF
			  VVRRRCCFFF
			  VVRCCCJFFF
			  VVVVCJJCFE
			  VVIVCCJJEE
			  VVIIICJJEE
			  MIIIIIJJEE
			  MIIISIJEEE
			  MMMISSJEEE`
	expected := 1206

	result := SolvePart2(input)

	if result != expected {
		t.Errorf("expected %d actual %d", expected, result)
	}
}
