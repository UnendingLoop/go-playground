package main

import (
	"fmt"
	"math"
)

// Dijkstra - algorithm for finding shortest path from A to B
func Dijkstra(graph map[int]map[int]int, start, end int) {
	if func() bool {
		_, exist1 := graph[start]
		_, exist2 := graph[end]
		return !(exist1 || exist2)
	}() {
		fmt.Println("A or B is not in the graph.")
		return
	}

	visited := map[int]bool{}
	parents := map[int]int{}
	distance := map[int]int{}
	for i := range graph {
		distance[i] = math.MaxInt
	}
	distance[start] = 0

	for current := start; current != end; current = minDistNode(distance, visited) {
		if current == -1 {
			fmt.Println("Path not found.")
			return
		}

		visited[current] = true
		for neighbor, dist := range graph[current] {
			if distance[current]+dist < distance[neighbor] {
				distance[neighbor] = distance[current] + dist
				parents[neighbor] = current
			}
		}

	}

	//Recoverying the path
	path := []int{}
	for node := end; node != start; node = parents[node] {
		_, ok := parents[node]
		if !ok {
			fmt.Println("Path not found.")
			return
		}
		path = append([]int{node}, path...)
	}
	path = append([]int{start}, path...)

	fmt.Printf("The path from %d to %d is: %v, total distance is %d.\n", start, end, path, distance[end])

}
func minDistNode(distance map[int]int, visited map[int]bool) int {
	minNode := -1
	minDist := math.MaxInt

	for node, dist := range distance {
		if !visited[node] && minDist > dist {
			minDist = dist
			minNode = node
		}
	}
	return minNode
}

func main() {
	graph := map[int]map[int]int{
		0: {1: 4, 2: 2},
		1: {2: 3, 3: 2},
		2: {1: 1, 3: 5},
		3: {},
	}
	Dijkstra(graph, 0, 3)
}
