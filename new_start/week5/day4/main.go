package main //graphы

import (
	"fmt"
	"slices"
)

// Graph - structure which provides methods for manipulations with graph
type Graph struct {
	nodes map[int][]int
}

// AddEdge - method for adding new node to graph
func (numbers *Graph) AddEdge(from, to int) {
	if connections, exists := numbers.nodes[from]; !exists {
		numbers.nodes[from] = []int{to}
		fmt.Printf("New node '%d' and its connection to '%d' is added to graph.\n", from, to)
	} else if !slices.Contains(connections, to) {
		numbers.nodes[from] = append(numbers.nodes[from], to)
		fmt.Printf("Node '%d' already exists, added connection to '%d'\n", from, to)
	} else {
		fmt.Printf("Node '%d' already exists! Nothing to update.\n", from)
	}
	if _, exists := numbers.nodes[to]; !exists {
		numbers.nodes[to] = []int{}
		fmt.Printf("Also created node '%d'.\n", to)
	}
}

// RemoveEdge - method for removing a node
func (numbers *Graph) RemoveEdge(toDelete int) {
	if _, exists := numbers.nodes[toDelete]; !exists {
		fmt.Printf("Node '%d doesn't exist! Nothing to delete.'\n", toDelete)
	} else {
		delete(numbers.nodes, toDelete)
		for node, connections := range numbers.nodes {
			if index := slices.Index(connections, toDelete); index != -1 {
				numbers.nodes[node] = slices.Delete(connections, index, index+1)
			}
		}
		fmt.Printf("Successfully removed node '%d'and its connections.'\n", toDelete)

	}
}

// GetFullInfo - method for printing all nodes with their connections
func (numbers Graph) GetFullInfo() {
	fmt.Printf("List of all nodes in graph:\n")
	for node, connections := range numbers.nodes {
		fmt.Printf("%d: %v\n", node, connections)
	}
}

// BFS - method for Breadth-First Search
func (numbers Graph) BFS(start int) {
	visited := map[int]bool{}
	queue := []int{start}
	fmt.Printf("Doing a BFS starting from node %d:\n", start)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, i := range numbers.nodes[current] {
			if !visited[i] {
				visited[i] = true
				queue = append(queue, i)
			}
		}
		fmt.Printf("Visited node: %d => %v\n", current, numbers.nodes[current])
	}
}

// DFS - method for Depth-First Search
func (numbers Graph) DFS(start int) {
	stack := []int{start}
	visited := map[int]bool{}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for i := len(numbers.nodes[current]) - 1; i >= 0; i-- {
			if !visited[numbers.nodes[current][i]] {
				stack = append(stack, numbers.nodes[current][i])
				visited[numbers.nodes[current][i]] = true
			}
		}
		fmt.Printf("DFS - Visited node: %d => %v\n", current, numbers.nodes[current])
	}
}

// DFSSearch - method for Depth-First Search
// Найти путь от А до В без использования рекурсии
func (numbers Graph) DFSSearch(start, target int) {
	if _, exists := numbers.nodes[start]; !exists {
		fmt.Println("Path is not found! Sorry!")
		return
	}
	if _, exists := numbers.nodes[target]; !exists {
		fmt.Println("Path is not found! Sorry!")
		return
	}

	stack := []int{start}
	visited := map[int]bool{}
	parents := map[int]int{}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		if current == target {
			break
		}
		stack = stack[:len(stack)-1]
		for i := len(numbers.nodes[current]) - 1; i >= 0; i-- {
			if !visited[numbers.nodes[current][i]] {
				stack = append(stack, numbers.nodes[current][i])
				visited[numbers.nodes[current][i]] = true
				parents[numbers.nodes[current][i]] = current
			}
		}
	}
	//Recovering the path
	if _, ok := parents[target]; !ok && start != target {
		fmt.Println("No path exists between the nodes.")
		return
	}

	path := make([]int, 0, len(parents))
	for i := target; i != start; i = parents[i] {
		path = append(path, i)
	}
	path = append(path, start)
	slices.Reverse(path)
	fmt.Printf("The DFS non-recursion path is: %v\n", path)

}

// BFSSearch - method for Breadth-First Search
// Найти путь от А до В
func (numbers Graph) BFSSearch(start, target int) {
	visited := map[int]bool{}
	parents := map[int]int{}

	queue := []int{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, i := range numbers.nodes[current] {
			if !visited[i] {
				visited[i] = true
				queue = append(queue, i)
				parents[i] = current
			}
		}
	}
	//Recovering the path
	if _, ok := parents[target]; !ok && start != target {
		fmt.Println("No path exists between the nodes.")
		return
	}

	path := make([]int, 0, len(parents))
	for i := target; i != start; i = parents[i] {
		path = append(path, i)
	}
	path = append(path, start)
	slices.Reverse(path)
	fmt.Printf("The BFS path is: %v\n", path)
}

func main() {
	numbers := Graph{
		nodes: map[int][]int{
			0: {1, 2},
			1: {2, 3},
			2: {3, 4},
			3: {5, 6},
			4: {3, 5},
			5: {0, 1},
			6: {0, 6},
		},
	}
	numbers.GetFullInfo()
	numbers.AddEdge(7, 5)
	numbers.GetFullInfo()
	numbers.RemoveEdge(7)
	numbers.GetFullInfo()
	numbers.BFS(2)
	numbers.DFS(4)
	numbers.DFSSearch(1, 6)
	numbers.BFSSearch(1, 6)

}
