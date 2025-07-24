package main

import "fmt"

// Input:  []int{1, 1, 2, 2, 2, 3, 4, 4}
// Output: ["1x2", "2x3", "3x1", "4x2"]

func groupAndCount(input []int) []string {
	answer := []string{}
	l := len(input) - 1

	for i, j := 0, 0; i <= l; i = j {
		for j = i; j < l && input[i] == input[j+1]; {
			j++
		}
		answer = append(answer, fmt.Sprint(input[i], "x", j-i+1))
		j++
	}

	return answer
}
func main() {
	input := []int{
		1,
		3,
		4, 4,
		5, 5, 5,
		6, 6, 6,
		7, 7, 7, 7,
		8, 8,
		9}
	answer := groupAndCount(input)
	fmt.Println("The result is:", answer)
}
