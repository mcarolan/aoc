package util

func Identity[T any](value T) T {
	return value
}

type Direction int64

const (
	North Direction = iota
	East
	South
	West
)

func (direction Direction) Clockwise() Direction {
	return (direction + 1) % 4
}

func (direction Direction) AntiClockwise() Direction {
	return (direction + 3) % 4
}

