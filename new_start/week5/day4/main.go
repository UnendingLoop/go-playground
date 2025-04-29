package main

import "fmt"

var visited map[int]bool

func bfs(graph map[int][]int, queue []int, node, target int, path []int) {
	if !visited[node] {
		queue = queue[1:] //убираем из очереди первый элемент
		visited[node] = true
		path = append(path, node)
		if node == target {
			fmt.Println("Path to target is: ", path)
		}
		queue = append(queue, graph[node]...) //добавляем в очередь детей текущей ноды
		fmt.Println("Current node: ", node)
		if len(queue) > 0 {
			bfs(graph, queue, queue[0], target, path)
		}
	}
}
func main() {
	visited = make(map[int]bool)
	graph := map[int][]int{
		0: {1, 2},
		1: {3, 4},
		5: {6, 7},
		8: {},
	}
	target := 7
	path := make([]int, 0)
	for node := range graph {
		bfs(graph, []int{node}, node, target, path)

	}
}
