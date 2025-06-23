// program for making all combinations of input letters without repeat
package main

import (
	"fmt"
	"sort"
)

var uniqueResults = make(map[string]bool)

func combinator(input []rune, answer string) {
	for i, v := range input {
		forward := append([]rune{}, input[:i]...)
		forward = append(forward, input[i+1:]...)
		combinator(forward, answer+string(v))
	}
	if len(answer) > 0 {
		uniqueResults[answer] = true
	} else {
		return
	}
}
func sorter() []string {
	result := []string{}
	for i := range uniqueResults {
		result = append(result, i)
	}
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) == len(result[j]) {
			return result[i] < result[j]
		}
		return len(result[i]) < len(result[j])
	})

	return result
}

func main() {
	letters := []rune{'a', 'b', 'c', 'd'}
	fmt.Println("Запускаю комбинатор...")
	combinator(letters, "")
	fmt.Println("Сортирую по алфавиту...")
	sortedResult := sorter()
	fmt.Println("Печатаю результат:")
	for _, i := range sortedResult {
		fmt.Println(i)
	}

}
