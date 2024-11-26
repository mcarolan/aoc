package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type InstructionType int

const (
	MoveForward InstructionType = iota
	MoveDown
	MoveUp
)

type instruction struct {
	instructionType InstructionType
	value           int
}

var instructionRegex = regexp.MustCompile(`(forward|down|up) ([0-9]+)\s*`)

var instructionMap = map[string]InstructionType{
	"forward": MoveForward,
	"down":    MoveDown,
	"up":      MoveUp,
}

func lineToInstruction(line string) (*instruction, error) {
	matches := instructionRegex.FindStringSubmatch(line)

	if len(matches) != 3 {
		return nil, fmt.Errorf("did not match instruction regex. %d matches", len(matches))
	}

	value, err := strconv.Atoi(matches[2])

	if err != nil {
		return nil, err
	}

	instructionType, found := instructionMap[matches[1]]

	if !found {
		return nil, fmt.Errorf("invalid instruction %s", matches[1])
	}

	return &instruction{instructionType, value}, nil
}

func extractInstructions(input string) ([]instruction, error) {
	lines := strings.Split(input, "\n")
	var instructions []instruction

	for i, line := range lines {
		instruction, err := lineToInstruction(line)

		if err != nil {
			return nil, fmt.Errorf("error on input line %d: %s", i, err)
		}

		instructions = append(instructions, *instruction)
	}

	return instructions, nil
}

func SolvePart1(input string) (int, error) {
	var depth int
	var horizontal int

	instructions, err := extractInstructions(input)

	if err != nil {
		return 0, err
	}

	for _, instruction := range instructions {
		switch instruction.instructionType {
		case MoveForward:
			horizontal += instruction.value
		case MoveUp:
			depth -= instruction.value
		case MoveDown:
			depth += instruction.value
		}
	}

	return depth * horizontal, nil
}

func SolvePart2(input string) (int, error) {
	var depth int
	var horizontal int
	var aim int

	instructions, err := extractInstructions(input)

	if err != nil {
		return 0, err
	}

	for _, instruction := range instructions {
		switch instruction.instructionType {
		case MoveForward:
			horizontal += instruction.value
			depth += aim * instruction.value
		case MoveUp:
			aim += -instruction.value
		case MoveDown:
			aim += instruction.value
		}
	}

	return depth * horizontal, nil
}
