package day17

import (
	"aoc/util"
	"fmt"
	"math"
	"strings"
)

type machine struct {
	registerA int
	registerB int
	registerC int

	instructionPointer int
}

type operand int

const (
	literal0 operand = iota
	literal1
	literal2
	literal3
	registerA
	registerB
	registerC
	reserved
)

type instruction int

const (
	adv instruction = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func comboOperandValue(operand operand, machine machine) int {
	switch operand {
	case literal0, literal1, literal2, literal3:
		return int(operand)
	case registerA:
		return machine.registerA
	case registerB:
		return machine.registerB
	case registerC:
		return machine.registerC
	default:
		panic(fmt.Sprintf("invalid combo operand %d", operand))
	}
}

type stepResult struct {
	halted bool
	output *int
}

func step(program []int, machine *machine) stepResult {
	if machine.instructionPointer >= len(program) {
		return stepResult{halted: true, output: nil}
	}

	instruction := instruction(program[machine.instructionPointer])
	operand := operand(program[machine.instructionPointer+1])

	nextInstruction := machine.instructionPointer + 2
	var output *int

	switch instruction {
	case adv:
		numerator := float64(machine.registerA)
		denominator := math.Pow(2, float64(comboOperandValue(operand, *machine)))
		machine.registerA = int(math.Trunc(numerator / denominator))
	case bxl:
		machine.registerB = machine.registerB | int(operand)
	case bst:
		machine.registerB = comboOperandValue(operand, *machine) % 8
	case jnz:
		if machine.registerA != 0 {
			nextInstruction = int(operand)
		}
	case bxc:
		machine.registerB = machine.registerB ^ machine.registerC
	case out:
		result := comboOperandValue(operand, *machine) % 8
		output = &result
	case bdv:
		numerator := float64(machine.registerA)
		denominator := math.Pow(2, float64(comboOperandValue(operand, *machine)))
		machine.registerB = int(math.Trunc(numerator / denominator))
	case cdv:
		numerator := float64(machine.registerA)
		denominator := math.Pow(2, float64(comboOperandValue(operand, *machine)))
		machine.registerC = int(math.Trunc(numerator / denominator))
	default:
		panic(fmt.Sprintf("invalid instruction %d", instruction))
	}

	machine.instructionPointer = nextInstruction
	return stepResult{halted: false, output: output}
}

func SolvePart1(input string) string {
	parts := strings.Split(input, "\n\n")
	initialRegisterValues := util.AllNumbers(parts[0])
	program := util.AllNumbers(parts[1])

	machine := machine{registerA: initialRegisterValues[0], registerB: initialRegisterValues[1], registerC: initialRegisterValues[2], instructionPointer: 0}

	output := make([]string, 0)
	for {
		stepResult := step(program, &machine)

		if stepResult.halted {
			break
		}

		if stepResult.output != nil {
			output = append(output, fmt.Sprintf("%d", *stepResult.output))
		}
	}

	return strings.Join(output, ",")
}

func SolvePart2(input string) int {
	return 42
}
