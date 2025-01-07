package day14

import (
	"aoc/util"
	"fmt"
)

type RobotConfiguration struct {
	positionX int
	positionY int
	velocityX int
	velocityY int
}

const GRID_WIDTH = 101
const GRID_HEIGHT = 103
const ITERATIONS = 100

func wrap(i int, n int) int {
	if i%n < 0 {
		return n + (i % n)
	} else {
		return i % n
	}
}

func step(config RobotConfiguration, grid util.Grid[int], iteration int) {
	lastPosition := util.RowCol{Row: wrap((config.positionY + (config.velocityY * (iteration - 1))), grid.Rows), Col: wrap(config.positionX+(config.velocityX*(iteration-1)), grid.ColsPerRow)}
	lastPositionValue, _ := grid.At(lastPosition)
	grid.Set(lastPosition, lastPositionValue-1)

	nextPosition := util.RowCol{Row: wrap(config.positionY+(config.velocityY*iteration), grid.Rows), Col: wrap(config.positionX+(config.velocityX*iteration), grid.ColsPerRow)}
	nextPositionValue, _ := grid.At(nextPosition)
	grid.Set(nextPosition, nextPositionValue+1)
}

func print(grid util.Grid[int], _ int, _ int) {
	for y := range grid.Rows {
		for x := range grid.ColsPerRow {
			// if y == gridHeight/2 {
			// 	fmt.Printf("-")
			// 	continue
			// }
			// if x == gridWidth/2 {
			// 	fmt.Printf("|")
			// 	continue
			// }
			position := util.RowCol{Row: y, Col: x}
			value, _ := grid.At(position)

			if value == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", value)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func SolvePart1(input string, gridWidth int, gridHeight int) int {
	rawInput := util.AllNumbers(input)
	configurations := make([]RobotConfiguration, 0)

	for i := 0; i < len(rawInput); i += 4 {
		configuration := RobotConfiguration{
			positionX: rawInput[i],
			positionY: rawInput[i+1],
			velocityX: rawInput[i+2],
			velocityY: rawInput[i+3],
		}

		configurations = append(configurations, configuration)
	}

	grid := util.NewGrid[int](gridWidth, gridHeight)

	for _, config := range configurations {
		position := util.RowCol{Row: config.positionY, Col: config.positionX}
		value, _ := grid.At(position)
		grid.Set(position, value+1)
	}

	for i := range ITERATIONS {
		for _, config := range configurations {
			step(config, grid, i+1)
		}

		if i < 5 {

			print(grid, gridWidth, gridHeight)
		}
	}

	var tl, tr, bl, br int

	for x := range gridWidth / 2 {
		for y := range gridHeight / 2 {
			value, _ := grid.At(util.RowCol{Row: y, Col: x})
			tl += value
		}
	}

	for x := (gridWidth / 2) + 1; x < gridWidth; x++ {
		for y := range gridHeight / 2 {
			value, _ := grid.At(util.RowCol{Row: y, Col: x})
			tr += value
		}
	}

	for x := range gridWidth / 2 {
		for y := (gridHeight / 2) + 1; y < gridHeight; y++ {
			value, _ := grid.At(util.RowCol{Row: y, Col: x})
			bl += value
		}
	}

	for x := (gridWidth / 2) + 1; x < gridWidth; x++ {
		for y := (gridHeight / 2) + 1; y < gridHeight; y++ {
			value, _ := grid.At(util.RowCol{Row: y, Col: x})
			br += value
		}
	}

	return tl * tr * bl * br
}

func SolvePart2(input string) int {
	return 42
}
