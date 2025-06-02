package set

import (
	"sync"
	"unsafe"
)

// Set - structure for storing
type Set[T comparable] struct {
	Items map[T]struct{}
	sync.RWMutex
}

// add
func (S *Set[T]) Add(new T) {
	S.Lock()
	defer S.Unlock()

	S.Items[new] = struct{}{}
}

// remove
func (S *Set[T]) Remove(candidate T) {
	S.Lock()
	defer S.Unlock()

	delete(S.Items, candidate)
}

// contains
func (S *Set[T]) Contains(candidate T) bool {
	S.RLock()
	defer S.RUnlock()

	_, ok := S.Items[candidate]
	return ok
}

// len
func (S *Set[T]) Len() int {
	S.RLock()
	defer S.RUnlock()

	return len(S.Items)
}

// isEmpty
func (S *Set[T]) IsEmpty() bool {
	S.RLock()
	defer S.RUnlock()

	return len(S.Items) == 0
}

// clear
func (S *Set[T]) Clear() {
	S.Lock()
	defer S.Unlock()

	S.Items = make(map[T]struct{})
}

// toSlice
func (S *Set[T]) ToSlice() []T {
	S.RLock()
	defer S.RUnlock()
	response := []T{}
	for i := range S.Items {
		response = append(response, i)
	}
	return response
}

// union
func (S *Set[T]) Union(C *Set[T]) *Set[T] {
	close := lock2sets(S, C)
	defer close()

	R := Set[T]{Items: make(map[T]struct{})}

	for s := range S.Items {
		R.Items[s] = struct{}{}
	}
	for c := range C.Items {
		R.Items[c] = struct{}{}
	}
	return &R
}

// intersection
func (S *Set[T]) Intersection(C *Set[T]) *Set[T] {
	close := lock2sets(S, C)
	defer close()

	R := Set[T]{Items: make(map[T]struct{})}
	for i := range S.Items {
		if _, ok := C.Items[i]; ok {
			R.Items[i] = struct{}{}
		}
	}
	return &R
}

// difference
func (S *Set[T]) Difference(C *Set[T]) *Set[T] {
	close := lock2sets(S, C)
	defer close()

	R := Set[T]{Items: make(map[T]struct{})}
	for i := range S.Items {
		if _, ok := C.Items[i]; !ok {
			R.Items[i] = struct{}{}
		}
	}
	return &R
}

// isSubset
func (S *Set[T]) IsSubset(C *Set[T]) bool {
	close := lock2sets(S, C)
	defer close()

	for i := range S.Items {
		if _, ok := C.Items[i]; !ok {
			return false
		}
	}
	return true
}

// aux lock2sets function - to avoid deadlocks
func lock2sets[T comparable](A, B *Set[T]) func() {
	addrA := uintptr(unsafe.Pointer(A))
	addrB := uintptr(unsafe.Pointer(B))
	if addrA < addrB {
		A.RLock()
		B.RLock()
		return func() {
			B.RUnlock()
			A.RUnlock()
		}
	}
	B.RLock()
	A.RLock()
	return func() {
		A.RUnlock()
		B.RUnlock()
	}
}
