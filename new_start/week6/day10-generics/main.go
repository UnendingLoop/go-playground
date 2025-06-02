package stack

import "fmt"

type stack[T any] struct {
	queue []T
}

// pop - get the value and remove it from queue
func (S *stack[T]) pop() (T, bool) {
	if len(S.queue) == 0 {
		var zero T
		return zero, false
	}
	last := S.queue[len(S.queue)-1]
	S.queue = S.queue[:(len(S.queue) - 1)]
	return last, true
}

// push - add the value to the queue
func (S *stack[T]) push(new T) {
	S.queue = append(S.queue, new)
}
func (S *stack[T]) reverse() {
	for i, j := 0, S.len()-1; i < j; i, j = i+1, j-1 {
		S.queue[i], S.queue[j] = S.queue[j], S.queue[i]
	}
}

// peek - get the next value, no deletion
func (S *stack[T]) peek() (T, bool) {
	if len(S.queue) == 0 {
		var zero T
		return zero, false
	}
	return S.queue[len(S.queue)-1], true
}

// len - get the length of the queue
func (S *stack[T]) len() int {
	return len(S.queue)
}

// isEmpty - know if the queue is empty
func (S *stack[T]) isEmpty() bool {
	return len(S.queue) == 0
}

// is existing array - for pop and peek
func must[T any](item T, flag bool) T {
	if !flag {
		panic("empty queue")
	}
	return item
}

func main() {
	stack := stack[int]{}
	stack.push(1)
	stack.push(4)
	stack.push(8)
	stack.push(16)
	stack.push(32)
	fmt.Println("The next item in queue:", must(stack.pop()))
	fmt.Println("Can it pop again? Read without deletion:", must(stack.peek()))
	fmt.Println("Current size of a queue:", stack.len())
	fmt.Println("Is the queue empty:", stack.isEmpty())
	fmt.Printf("Current queue: %d\n", stack.queue)
	stack.reverse()
	fmt.Printf("Reversed queue: %d\n", stack.queue)

}
