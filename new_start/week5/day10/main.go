package kata

import "container/heap"

type MinHeap []int
type Test int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func DblLinear(n int) int {
	h := &MinHeap{1} // Начинаем с первого элемента (1)
	heap.Init(h)     // Инициализируем кучу

	seen := make(map[int]bool)
	seen[1] = true

	var current int
	for i := 0; i < n; i++ {
		current = heap.Pop(h).(int) // Извлекаем минимальный элемент

		// Добавляем в кучу новые элементы, если они еще не были добавлены
		child1, child2 := 2*current+1, 3*current+1
		if !seen[child1] {
			heap.Push(h, child1)
			seen[child1] = true
		}
		if !seen[child2] {
			heap.Push(h, child2)
			seen[child2] = true
		}
	}

	return current
}
