package main

import (
	"fmt"
	"math"
	"time"
)

func findCompositeSquares(init int) (answer []int) {
	callcounter := 0
	found := false
	var helper func(n, remaining int, path []int)
	helper = func(n, remaining int, path []int) {
		callcounter++
		if found || (len(path) > 0 && n >= path[len(path)-1]) {
			return
		}
		path = append(path, n)
		remaining -= n * n

		if remaining == 0 {
			found = true
			for k := len(path) - 1; k >= 1; k-- {
				answer = append(answer, path[k])
			}
			return
		}

		candidate := int(math.Sqrt(float64(remaining)))

		for i := candidate; i > 0; i-- {
			if i*i > remaining {
				continue
			}
			helper(i, remaining, path)
		}
	}

	helper(init, (init)*(init)*2, []int{})
	fmt.Println("Calls using sqrt:", callcounter)
	return
}
func findCompositeSquaresss(init int) (answer []int) {
	callcounter := 0
	var helper func(n, remaining int) []int
	helper = func(s, i int) []int {
		callcounter++
		if s < 0 {
			return nil
		}
		if s == 0 {
			return []int{}
		}
		var j int = i - 1
		for ; j > 0; j-- {
			var sub = helper(s-j*j, j)
			if sub != nil {
				return append(sub, []int{j}...)
			}
		}
		return nil
	}

	answer = helper((init)*(init), init)
	fmt.Println("Calls using bruteforce:", callcounter)
	return
}

func main() {
	k := 409986
	timer1 := time.Now()
	fmt.Println("The sqrt-method answer is:", findCompositeSquares(k))
	timer1end := time.Since(timer1)
	fmt.Println("Time elapsed:", timer1end)
	timer2 := time.Now()
	fmt.Println("The bruteforce-method answer is:", findCompositeSquaresss(k))
	timer2end := time.Since(timer2)
	fmt.Println("Time elapsed:", timer2end)

}
