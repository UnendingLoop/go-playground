package main

import "fmt"

func snailClockwiseSymmetric(snailMap [][]int) []int {
	length := len(snailMap)
	if length == 0 || len(snailMap[0]) == 0 {
		return []int{}
	}

	top, bottom := 0, length-1
	left, right := 0, length-1
	result := make([]int, 0)

	for bottom >= top && left <= right {
		//top: L2R
		for i := left; i <= right; i++ {
			result = append(result, snailMap[top][i])
		}
		top++

		//right: T2B
		for i := top; i <= bottom; i++ {
			result = append(result, snailMap[i][right])
		}
		right--

		//bottom: R2L
		if bottom >= top {
			for i := right; i >= left; i-- {
				result = append(result, snailMap[bottom][i])
			}
			bottom--
		}
		//left: B2T
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, snailMap[i][left])
			}
			left++
		}
	}

	return result
}
func snailCounterClockwiseAsymetric(snailMap [][]int) []int {
	lengthX := len(snailMap[0])
	lengthY := len(snailMap)
	if lengthX == 0 {
		return []int{}
	}

	top, bottom := 0, lengthY-1
	left, right := 0, lengthX-1
	result := make([]int, 0, lengthX*lengthY)

	for bottom >= top && left <= right {
		//left: T2B
		for i := top; i <= bottom; i++ {
			result = append(result, snailMap[i][left])
		}
		left++

		//bottom: L2R
		if top <= bottom {
			for i := left; i <= right; i++ {
				result = append(result, snailMap[bottom][i])
			}
			bottom--
		}

		//right: B2T
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, snailMap[i][right])
			}
			right--
		}

		//top: R2L
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, snailMap[top][i])
			}
			top++
		}

	}

	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	answer := snailClockwiseSymmetric(matrix)
	fmt.Println(answer)
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	answer = snailCounterClockwiseAsymetric(matrix)
	fmt.Println(answer)

}
