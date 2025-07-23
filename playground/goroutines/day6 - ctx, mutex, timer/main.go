package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}

	channel := make(chan int, 5)

	for i := range 5 { // запуск воркеров
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Worker #%d received timeout: %v\nShutting down...\n", n, ctx.Err())
					return
				case j, ok := <-channel:
					if !ok {
						fmt.Printf("Worker #%d: channel is closed! Shutting down...\n", n)
						return
					}
					fmt.Printf("Worker #%d doing job #%d\n", n, j)
					time.Sleep(time.Duration(r.Intn(900)+100) * time.Millisecond)
				}
			}
		}(i)
	}
	wg.Add(1)
	go func() { //Запуск генератора
		defer wg.Done()
		defer close(channel)
		for i := range 20 {
			select {
			case <-ctx.Done():
				fmt.Printf("Generator received timeout: %v\nShutting down...\n", ctx.Err())
				return
			case channel <- i:
				fmt.Printf("Generator sent job: %v\n", i)
				time.Sleep(150 * time.Millisecond)
			}
		}
		fmt.Printf("Generator finished sending jobs.BYE!\n")
	}()

	wg.Wait()
	fmt.Printf("All goroutines are stopped. BYE!\n")

}
