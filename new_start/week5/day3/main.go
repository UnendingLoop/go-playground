package main

import "fmt"

// var path []int
var visited map[int]bool

func dfs(graph map[int][]int, node, target int, path []int) {
	if visited[node] {
		return
	}
	path = append(path, node)
	visited[node] = true                 // Помечаем вершину как посещенную
	fmt.Println("Visiting node: ", node) // Выводим вершину

	fmt.Println("Path to target is: ", path)

	for _, neighbour := range graph[node] {
		dfs(graph, neighbour, target, path)
	}
}

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0},
		3: {1},
		4: {1},
		5: {6, 7},
		6: {5},
		7: {5},
		8: {},
	}
	visited = make(map[int]bool)
	counter := 0
	path := make([]int, 0)
	target := 4
	for node := range len(graph) {
		if !visited[node] {
			dfs(graph, node, target, path) // Начнем поиск в глубину с вершины 0
			counter++
		}
		path = make([]int, 0)
	}
	fmt.Println("Number of islands: ", counter)

}
