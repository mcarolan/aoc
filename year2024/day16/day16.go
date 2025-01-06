package day16

import (
	"aoc/util"
	"container/heap"
	"math"
)

type node struct {
	position  util.RowCol
	direction util.Direction
}

type neighbour struct {
	node   node
	weight int
}

type item struct {
	node     node
	priority int
	index    int
}

type PriorityQueue []*item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)

	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func neighbours(grid util.Grid[rune], n node) []neighbour {
	result := []neighbour{
		{node: node{position: n.position, direction: n.direction.AntiClockwise()}, weight: 1000},
		{node: node{position: n.position, direction: n.direction.Clockwise()}, weight: 1000},
	}

	neighbourPos := n.position.Neighbour(n.direction)
	value, isValid := grid.At(neighbourPos)

	if isValid && value == '.' || value == 'E' {
		result = append(result, neighbour{node: node{position: neighbourPos, direction: n.direction}, weight: 1})
	}

	return result
}

func dijkstra(grid util.Grid[rune], start node) map[node]int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	distances := make(map[node]int)

	for cell := range grid.Iterator() {
		distances[node{position: cell.Pos, direction: util.North}] = math.MaxInt
		distances[node{position: cell.Pos, direction: util.East}] = math.MaxInt
		distances[node{position: cell.Pos, direction: util.South}] = math.MaxInt
		distances[node{position: cell.Pos, direction: util.West}] = math.MaxInt
	}
	distances[start] = 0

	heap.Push(&pq, &item{
		node:     start,
		priority: 0,
	})

	visited := make(map[node]bool)

	for pq.Len() > 0 {
		i := heap.Pop(&pq).(*item)
		current := i.node

		if visited[current] {
			continue
		}

		visited[current] = true

		for _, neighbour := range neighbours(grid, current) {
			if visited[neighbour.node] {
				continue
			}

			alt := distances[current] + neighbour.weight

			if alt < distances[neighbour.node] {
				distances[neighbour.node] = alt
				heap.Push(&pq, &item{
					node:     neighbour.node,
					priority: alt,
				})
			}
		}
	}

	return distances
}

func SolvePart1(input string) int {
	maze := util.ParseGrid[rune](util.Lines(input), util.Identity)

	var start, end *util.RowCol

	for cell := range maze.Iterator() {
		if cell.Value == 'S' {
			start = &cell.Pos
			if end != nil {
				break
			}
		} else if cell.Value == 'E' {
			end = &cell.Pos
			if start != nil {
				break
			}
		}
	}

	distances := dijkstra(maze, node{position: *start, direction: util.East})

	result := math.MaxInt
	for n, distance := range distances {
		if n.position == *end && distance < result {
			result = distance
		}
	}
	return result
}

func SolvePart2(input string) int {
	return 42
}
