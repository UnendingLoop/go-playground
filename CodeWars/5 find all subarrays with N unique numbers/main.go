package main

import "fmt"

// Дан массив целых чисел input []int и число k.
// Найди все непрерывные подмассивы, в которых ровно k уникальных чисел.
func uniqueLimitSusbseqsFinder(input []int, k int) (answer [][]int) {
	len := len(input)
	for i, j, counter, store := 0, 0, 0, make(map[int]bool); i < len; i, counter, store = i+1, 0, map[int]bool{} {
		for j = i; j < len; j++ {
			if !store[input[j]] {
				store[input[j]] = true
				counter++
			}
			if counter > k {
				break
			}
			if counter == k {
				answer = append(answer, input[i:j+1])
			}
		}
	}
	return
}

func main() {
	input := []int{1, 1, 2, 3, 3, 4, 5, 6, 4, 3, 2, 2, 1, 1}
	result := uniqueLimitSusbseqsFinder(input, 2)
	fmt.Println("The result is:")
	for _, v := range result {
		fmt.Println(v)
	}
}
