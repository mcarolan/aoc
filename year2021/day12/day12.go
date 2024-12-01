package day12

import (
	"aoc/util"
	"strings"
)

func parseAdjacencyMatrix(input string) map[string][]string {
	result := make(map[string][]string)

	for _, line := range util.Lines(input) {
		nodes := strings.Split(line, "-")
		a := nodes[0]
		b := nodes[1]

		listA, hasList := result[a]
		if !hasList {
			listA = make([]string, 0)
		}

		listB, hasList := result[b]

		if !hasList {
			listB = make([]string, 0)
		}

		result[a] = append(listA, b)
		result[b] = append(listB, a)
	}

	return result
}

func contains(a []string, value string) bool {
	for _, s := range a {
		if s == value {
			return true
		}
	}
	return false
}

func mapKey(stringArray []string) string {
	var result strings.Builder

	for _, s := range stringArray {
		result.WriteString(s)
		result.WriteRune(':')
	}

	return result.String()
}

func bfs(adjacencyMatrix map[string][]string) [][]string {
	q := util.NewStack[[]string]()
	q.Push([]string{"start"})

	visited := make(map[string]bool)

	results := make([][]string, 0)

	for !q.IsEmpty() {
		current := q.Pop()
		currentKey := mapKey(current)

		if visited[currentKey] {
			continue
		}
		visited[currentKey] = true

		at := current[len(current)-1]

		if at == "end" {
			results = append(results, current)
			continue
		}

		neighbours := adjacencyMatrix[at]

		for _, n := range neighbours {
			if n == strings.ToLower(n) && contains(current, n) {
				continue
			}

			next := make([]string, len(current)+1)
			copy(next, current)
			next[len(next)-1] = n

			q.Push(next)
		}
	}
	return results
}

func SolvePart1(input string) int {
	adjacencyMatrix := parseAdjacencyMatrix(input)
	paths := bfs(adjacencyMatrix)
	return len(paths)
}

func SolvePart2(input string) int {
	return 42
}
