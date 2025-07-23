package main

import (
	"fmt"
	"sync"
)

func main() {
	token := make(chan int, 1)
	wg := sync.WaitGroup{}
	const workers = 3
	const cycles = 3
	for i := range workers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for j := 0; j < cycles; j++ {
				curr := <-token
				if curr != n {
					token <- curr
					j--
					continue
				}
				fmt.Printf("Worker #%d doing work #%d...\n", i, n)
				token <- (curr + 1) % workers
			}
		}(i)
	}
	token <- 0
	wg.Wait()
	close(token)

}
