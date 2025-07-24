package main

import "fmt"

// Найди все непрерывные подпоследовательности (подмассивы), сумма которых строго меньше N.
func findSubseqBelowLimit(input []int, limit int) (answer [][]int) {
	lenInput := len(input)

	for i, j := 0, 0; i < lenInput; i++ {
		summ := 0
		for j = i; j < lenInput && summ+input[j] < limit; j++ {
			summ += input[j]
			answer = append(answer, input[i:j+1])
		}
	}
	return
}

func main() {
	input := []int{1, 2, 1, 1, 2, 4, 3, 2, 1, 1}
	result := findSubseqBelowLimit(input, 4)
	fmt.Println("The result is:")
	for _, v := range result {
		fmt.Println(v)
	}
}
