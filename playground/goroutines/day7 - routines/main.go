package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

var ctx, cancel = context.WithCancel(context.Background())
var jobQueue = make(chan int, 5)
var wg = sync.WaitGroup{}
var syncOnce = sync.Once{}

func timer() {
	deadline := time.NewTimer(5 * time.Second)
	go func() {
		<-deadline.C
		fmt.Println("TIMER: Deadline reached...")
		syncOnce.Do(func() { cancel() })
	}()
}
func ctrlC() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		defer close(sig)
		<-sig
		fmt.Println("Ctrl+C inerruption received...")
		syncOnce.Do(func() { cancel() })
	}()
}

func generator() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(jobQueue)
		for i := range 40 {
			select {
			case <-ctx.Done():
				fmt.Printf("Generator received shutdown signal: %v...\n", ctx.Err())
				return
			case jobQueue <- i:
				fmt.Printf("Generator sent job #%d...\n", i)
				time.Sleep(50 * time.Millisecond)
			}
		}
		fmt.Printf("Generator finished sending all jobs. Closing channel and shutting down...\n")
		syncOnce.Do(func() { cancel() })
	}()
}
func newWorker() {
	for i := range 10 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for {
				select {
				case job, ok := <-jobQueue:
					if !ok {
						fmt.Printf("Worker #%d: channel is closed! Shutting down...\n", n)
						return
					}
					fmt.Printf("Worker #%d doing job #%d...\n", n, job)
					time.Sleep(100 * time.Millisecond)
				case <-ctx.Done():
					fmt.Printf("Worker #%d received shutdown signal: %v...\n", n, ctx.Err())
					return
				}
			}
		}(i)
	}
}

func main() {
	timer()
	ctrlC()
	generator()
	newWorker()
	wg.Wait()
	fmt.Printf("All goroutines completed! Shutting down...\n")

}
