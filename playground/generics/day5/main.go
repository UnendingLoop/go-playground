package main

import "sort"

// SliceFn - structure for storing slice of comparable type, and func for comparing elements of the slice
type SliceFn[T comparable] struct {
	S       []T
	MinS    T
	MaxS    T
	Compare func(a, b T) bool // true if a < b
}

func (SF *SliceFn[T]) Map(wtd func(T) T) {
	for i, v := range SF.S {
		SF.S[i] = wtd(v)
	}
}

func (SF *SliceFn[T]) Filter(cond func(T) bool) {
	filteredS := make([]T, 0, len(SF.S)/3)
	for _, v := range SF.S {
		if cond(v) {
			filteredS = append(filteredS, v)
		}
	}
	SF.S = filteredS
	SF.Max() //updating max item after applying filter
	SF.Min() //updating min item after applying filter
}

func (SF *SliceFn[T]) Sort() {
	sort.Slice(SF.S, func(i, j int) bool {
		return SF.Compare(SF.S[i], SF.S[j])
	})
}

func (SF *SliceFn[T]) Min() {
	SF.MinS = SF.S[0]
	for _, v := range SF.S[1:] {
		if SF.Compare(v, SF.MinS) {
			SF.MinS = v
		}
	}
}

func (SF *SliceFn[T]) Max() {
	SF.MaxS = SF.S[0]
	for _, v := range SF.S[1:] {
		if SF.Compare(SF.MaxS, v) {
			SF.MaxS = v
		}
	}
}

func (SF SliceFn[T]) Reduce(init T, wtd func(acc, val T) T) T {
	for _, v := range SF.S {
		init = wtd(init, v)
	}
	return init
}

func main() {
	exampleINT := SliceFn[int]{
		S: []int{10, 2, 5, 3, 4, 6, 7, 9, 8},
		Compare: func(a, b int) bool {
			return a < b
		},
	}
	exampleINTreduce := exampleINT.Reduce(0, func(acc, val int) int {
		return acc + val
	})
	exampleSTRING := SliceFn[string]{
		S: []string{"aaaaa", "bbbb", "ccc", "dd", "e"},
		Compare: func(a, b string) bool {
			return len(a) < len(b)
		},
	}
}
