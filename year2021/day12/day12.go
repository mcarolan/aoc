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

func counts(a []string) map[string]int {
	result := make(map[string]int)
	for _, s := range a {
		result[s]++
	}
	return result
}

func mapKey(stringArray []string) string {
	var result strings.Builder

	for _, s := range stringArray {
		result.WriteString(s)
		result.WriteRune(':')
	}

	return result.String()
}

func bfs(adjacencyMatrix map[string][]string, isTargetState func(string) bool, neighbours func([]string) []string) [][]string {
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

		if isTargetState(at) {
			results = append(results, current)
			continue
		}

		neighbours := neighbours(current)

		for _, n := range neighbours {
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
	paths := bfs(adjacencyMatrix, func(at string) bool { return at == "end" }, func(current []string) []string {
		at := current[len(current)-1]
		return util.Filter(adjacencyMatrix[at], func(n string) bool {
			return n == strings.ToUpper(n) || counts(current)[n] == 0
		})
	})
	return len(paths)
}

func SolvePart2(input string) int {
	adjacencyMatrix := parseAdjacencyMatrix(input)

	isValidPath := func(path []string, n string) bool {
		if n == "start" {
			return false
		}

		counts := counts(path)
		counts[n]++

		hasDoubleVisited := false

		for key, count := range counts {
			if key == strings.ToUpper(key) {
				continue
			}

			if count > 2 {
				return false
			}

			if count > 1 {
				if hasDoubleVisited {
					return false
				} else {
					hasDoubleVisited = true
				}
			}
		}

		return true
	}

	paths := bfs(adjacencyMatrix, func(at string) bool { return at == "end" }, func(current []string) []string {
		at := current[len(current)-1]
		return util.Filter(adjacencyMatrix[at], func(n string) bool {
			valid := isValidPath(current, n)
			return valid
		})
	})

	return len(paths)
}
