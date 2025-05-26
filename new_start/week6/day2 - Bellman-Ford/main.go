package main

import (
	"fmt"
	"math"
)

func bellmanford(start int, graph map[int]map[int]int) {
	//adding "invisible" nodes to graph
	for _, neighbours := range graph {
		for neighbour := range neighbours {
			if _, ok := graph[neighbour]; !ok {
				graph[neighbour] = map[int]int{}
			}
		}
	}
	//preparing auxiliary tables for main cycle, set distances as infinity
	distance := make(map[int]int, len(graph))
	parents := make(map[int]int, len(graph))
	for i := range graph {
		distance[i] = math.MaxInt
	}
	distance[start] = 0

	//the main cycle
	for i := 0; i < len(graph)-1; i++ {
		for current, neighbours := range graph {
			for neighbour, weight := range neighbours {
				if distance[current] != math.MaxInt && distance[current]+weight < distance[neighbour] {
					distance[neighbour] = distance[current] + weight
					parents[neighbour] = current
				}
			}
		}
	}
	//checking for negative cycles
	for current, neighbours := range graph {
		for neighbour, weight := range neighbours {
			if distance[current] != math.MaxInt && distance[current]+weight < distance[neighbour] {
				fmt.Println("Negative cycle detected! Cannot find paths!")
				return
			}
		}
	}
	//recovering paths
	for node := range graph {
		if distance[node] == math.MaxInt {
			fmt.Printf("No path from %d to %d\n", start, node)
			continue
		}

		if node == start {
			continue
		}
		path := []int{}
		for current := node; current != start; current = parents[current] {
			_, ok := parents[current]
			if !ok {
				fmt.Printf("No path from %d to %d.\n", start, node)
				break
			}
			path = append([]int{current}, path...)
		}
		path = append([]int{start}, path...)
		fmt.Printf("Shortest path from %d to %d is: %v, total weight is %d.\n", start, node, path, distance[node])

	}
}
func main() {
	graph := map[int]map[int]int{
		0: {1: -1, 2: 4},
		1: {2: 3, 3: 2, 4: 2},
		2: {5: 7},
		3: {2: 5},
		4: {3: -3},
	}
	bellmanford(1, graph)
}
