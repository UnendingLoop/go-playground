package main

import "fmt"

type SliceFn[T any] struct {
	S       []T
	Compare func(T, T) bool
}

type Person struct {
	Name string
	Age  int
}

func (s SliceFn[T]) Max() T {
	if len(s.S) == 0 {
		panic("slice is empty")
	}
	max := s.S[0]
	for _, v := range s.S[1:] {
		if s.Compare(v, max) {
			max = v
		}
	}
	return max
}

func main() {
	people := SliceFn[Person]{
		S: []Person{
			{"Ivan", 30},
			{"Olga", 25},
			{"Petr", 40},
		},
		Compare: func(a, b Person) bool {
			return a.Age > b.Age // сравнение по возрасту
		},
	}

	fmt.Println("Oldest person:", people.Max())
}
