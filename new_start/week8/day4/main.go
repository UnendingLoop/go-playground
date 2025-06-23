package main

import "fmt"

func parenthesisGen(n int) (answer []string) {
	var backtrack func(current string, open, close int)
	backtrack = func(current string, open, close int) {
		if open == n && close == n {
			answer = append(answer, current)
			return
		}
		if open < n {
			backtrack(current+"(", open+1, close)

		}
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	backtrack("", 0, 0)

	return
}
func main() {
	/*balancedParens 0 -> [""]
	balancedParens 1 -> ["()"]
	balancedParens 2 -> ["()()","(())"]
	balancedParens 3 -> ["()()()","(())()","()(())","(()())","((()))"]
	*/
	pairsCount := 0
	fmt.Println("How many parenthesis pairs should be generated?")
	fmt.Scanln(&pairsCount)
	fmt.Println(parenthesisGen(pairsCount))
}
