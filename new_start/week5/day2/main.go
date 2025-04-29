package main

import "fmt"

func insertion(input [][]int, size int) map[int][]int {
	output := make(map[int][]int, size)
	for _, j := range input {
		output[j[0]] = append(output[j[0]], j[1])
		output[j[1]] = append(output[j[1]], j[0])
	}
	return output
}

func print(graph map[int][]int, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("%d: %v\n", i, graph[i])
	}
}

func main() {
	n := 5
	array := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{2, 3},
		{4, 0},
	}
	graph := insertion(array, n)
	print(graph, n)
}
