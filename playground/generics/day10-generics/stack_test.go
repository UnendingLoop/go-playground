package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	var s stack[int]
	if !s.isEmpty() {
		t.Error("Queue must be initially empty!")
	}
	s.push(1)
	s.push(2)
	s.push(4)
	s.push(8)
	s.push(16)
	if s.len() != 5 {
		t.Error("The queue length must be 5!")
	}
	if must(s.peek()) != 16 || s.len() != 5 {
		t.Errorf("Peeking must not remove items from queue!")
	}
	if must(s.pop()) != 16 && s.len() != 4 {
		t.Error("Popping must remove 1 item from queue!")
	}
	s.reverse()
	flag := true
	for i, j := 0, 8; i < (s.len() - 1); i, j = i+1, j/2 {
		if s.queue[i] != j {
			flag = false
		}
		fmt.Printf("Array item: %d, expected value: %d\n", s.queue[i], j)
	}
	if !flag {
		t.Error("Reverse doesn't work properly!")
	}

}
