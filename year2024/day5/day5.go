package day5

import (
	"aoc/util"
	"slices"
	"strings"
)

func SolvePart1(input string) int {
	inputParts := strings.Split(input, "\n\n")
	rules := util.Map(util.Lines(inputParts[0]), util.AllNumbers)
	updates := util.Map(util.Lines(inputParts[1]), util.AllNumbers)

	result := 0

	for _, update := range updates {
		positions := make(map[int]int)
		for i, entry := range update {
			positions[entry] = i
		}

		validUpdate := true

		for _, rule := range rules {
			a := rule[0]
			b := rule[1]

			aIndex, hasA := positions[a]
			if !hasA {
				continue
			}

			bIndex, hasB := positions[b]
			if !hasB {
				continue
			}

			if aIndex > bIndex {
				validUpdate = false
				break
			}
		}

		if validUpdate {
			result += update[len(update)/2]
		}
	}

	return result
}

func SolvePart2(input string) int {
	inputParts := strings.Split(input, "\n\n")
	rules := util.Map(util.Lines(inputParts[0]), util.AllNumbers)
	updates := util.Map(util.Lines(inputParts[1]), util.AllNumbers)

	result := 0

	invalidUpdates := make(map[int]*[]int)

	for updateIndex, update := range updates {
		positions := make(map[int]int)
		for i, entry := range update {
			positions[entry] = i
		}

		applicableRules := make([]int, 0)

		for ruleIndex, rule := range rules {
			a := rule[0]
			b := rule[1]

			aIndex, hasA := positions[a]
			if !hasA {
				continue
			}

			bIndex, hasB := positions[b]
			if !hasB {
				continue
			}

			applicableRules = append(applicableRules, ruleIndex)

			if aIndex > bIndex {
				invalidUpdates[updateIndex] = &applicableRules
			}
		}
	}

	for updateIndex, applicableRules := range invalidUpdates {
		update := updates[updateIndex]
		slices.SortFunc(update, func(a int, b int) int {
			for _, ruleIndex := range *applicableRules {
				rule := rules[ruleIndex]

				if (rule[0] == a && rule[1] == b) {
					return -1
				} else if (rule[0] == b && rule[1] == a) {
					return 1
				}
			}

			return 0
		})
		result += update[len(update)/2]
	}

	return result
}
