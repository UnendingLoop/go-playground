package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type intHeap []uint

func (m intHeap) Len() int {
	return len(m)
}
func (m intHeap) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m *intHeap) Pop() any {
	old := *m
	n := len(old) - 1
	*m = old[0:n]
	return old[n]
}
func (m *intHeap) Push(new any) {
	*m = append(*m, new.(uint))
}
func (m intHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Hammer -
func Hammer(n int) uint {
	h := &intHeap{1}
	y := &intHeap{}
	heap.Init(h)
	heap.Init(y)
	seen := make(map[uint]bool, n)
	seen[1] = true
	mults := []uint{2, 3, 5}

	for i := 1; i <= n; i++ {
		current := heap.Pop(h).(uint)
		heap.Push(y, current)
		fmt.Printf("Current i=%d, value=%d\n", i, current)
		if i == n {
			z := *y
			sort.Slice(z, func(i, j int) bool {
				return z[i] > z[j]
			})
			return z[0]
		}
		for _, v := range mults {
			new := v * current
			if !seen[new] {
				heap.Push(h, new)
				seen[new] = true
			}
		}
	}
	return 0
}
func main() {
	n := 13282
	answer := Hammer(n)
	fmt.Printf("The %dth Hammer number is %d\n", n, answer)

}
