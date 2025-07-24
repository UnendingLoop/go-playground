package main

import "fmt"

// LongestSlideDown -
func LongestSlideDown(input [][]int) int {
	pyramid := make([][]int, len(input))
	for i := range input {
		pyramid[i] = append([]int(nil), input[i]...)
	}

	// Инициализируем путь
	path := make([][]int, len(input))
	for i := range path {
		path[i] = []int{}
	}

	for i := len(pyramid) - 2; i >= 0; i-- {
		for j := range pyramid[i] {
			maxNeighbor, isLeft := max(pyramid[i+1][j], pyramid[i+1][j+1])
			index := j + 1
			if isLeft {
				index = j
			}
			pyramid[i][j] += maxNeighbor
			path[j] = append(path[j], index)
		}
	}
	fmt.Println("All collected paths:", path)

	return pyramid[0][0]
}

func max(a, b int) (int, bool) {
	if a > b {
		return a, true
	}
	return b, false
}

func main() {
	input := [][]int{
		{3},
		{7, 4},
		{2, 4, 6},
		{8, 5, 9, 3},
		{5, 6, 8, 9, 5},
		{5, 3, 2, 6, 1, 0}}
	answer := LongestSlideDown(input)
	fmt.Println("The answer is: ", answer)
	fmt.Println("Initial array: ", input)

}
