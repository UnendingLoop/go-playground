package kata

func DblLinear(n int) int {
	queue1 := []int{}
	queue2 := []int{}
	seen := make(map[int]bool)

	counter := 0
	current := 1

	for counter < n {
		child1, child2 := 2*current+1, 3*current+1

		if !seen[child1] {
			queue1 = append(queue1, child1)
			seen[child1] = true
		}
		if !seen[child2] {
			queue2 = append(queue2, child2)
			seen[child2] = true
		}
		current = findMin(&queue1, &queue2)
		counter++
	}
	return current
}

func findMin(queue1, queue2 *[]int) int {
	temp := 0

	switch {
	case len(*queue1) == 0:
		temp := (*queue2)[0]
		*queue2 = (*queue2)[1:]
		return temp
	case len(*queue2) == 0:
		temp := (*queue1)[0]
		*queue1 = (*queue1)[1:]
		return temp
	case (*queue1)[0] < (*queue2)[0]:
		temp = (*queue1)[0]
		(*queue1) = (*queue1)[1:]
	case (*queue1)[0] > (*queue2)[0]:
		temp = (*queue2)[0]
		(*queue2) = (*queue2)[1:]
	case (*queue1)[0] == (*queue2)[0]:
		temp = (*queue2)[0]
		(*queue2) = (*queue2)[1:]
		(*queue1) = (*queue1)[1:]
	}

	return temp
}
