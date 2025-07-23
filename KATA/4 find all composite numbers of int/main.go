package main

import (
	"fmt"
)

// partition считает количество разбиений числа n,
// где каждое слагаемое ≤ max
func partition(n, max int) int {
	if n == 0 {
		return 1 // один способ — пустая сумма
	}
	if n < 0 || max == 0 {
		return 0 // нет способов
	}

	// два случая:
	// 1. не использовать max → уменьшаем max
	// 2. использовать хотя бы один max → уменьшаем n
	return partition(n, max-1) + partition(n-max, max)
}

// ExpSum возвращает количество разбиений числа n
func ExpSum(n int) int {
	return partition(n, n)
}

func main() {
	for i := range 10 {
		fmt.Printf("ExpSum(%d) = %d\n", i, ExpSum(i))
	}
}
