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

		list, hasList := result[a]

		if !hasList {
			list = make([]string, 0)
		}

		result[a] = append(list, b)
	}

	return result
}

func bfs(at string, adjacencyMatrix map[string][]string, visited map[string]int) [][]string {
	if at == "end" {
		return make([][]string, 0)
	}
	
	visitCount := visited[at]

	if strings.ToLower(at) == at && visitCount > 1 {
		return nil
	}

	visited[at] = visitCount + 1

	neighbours := adjacencyMatrix[at]
}

func SolvePart1(input string) int {
	return 42
}

func SolvePart2(input string) int {
	return 42
}
