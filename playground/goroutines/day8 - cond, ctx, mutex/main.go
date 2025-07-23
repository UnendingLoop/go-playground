package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

var cancel context.CancelFunc
var onceDo sync.Once

type jobQueue struct {
	queue []int
	cond  *sync.Cond
	wg    sync.WaitGroup
	ctx   context.Context
}

func (jQ *jobQueue) ctrlC() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		defer close(sig)
		<-sig
		fmt.Println("\n!!! Received Ctrl+C interruption. Starting shutdown sequence...")
		onceDo.Do(func() {
			cancel()
			jQ.cond.Broadcast() //будим все горутины чтобы они завершились
		})
	}()
}

func (jQ *jobQueue) generator(queueSize int) {
	jQ.wg.Add(1)
	go func() {
		defer jQ.wg.Done()
		for i := range 20 {
			select {
			case <-jQ.ctx.Done():
				fmt.Printf("Generator: received a shutdown signal: %v. Shutting down...\n", jQ.ctx.Err())
				jQ.cond.Broadcast() //будим все горутины чтобы они завершились
				return
			default:
				jQ.cond.L.Lock()
				jQ.queue = append(jQ.queue, i)
				jQ.cond.L.Unlock()
				jQ.cond.Signal()
				fmt.Printf("Generator: sent job #%d...\n", i)
				time.Sleep(300 * time.Millisecond)
			}
		}
		fmt.Println("Generator finished sending jobs! Starting shutdown sequence...")
		onceDo.Do(func() {
			cancel()
			jQ.cond.Broadcast() //будим все горутины чтобы они завершились
		})
	}()
}

func (jQ *jobQueue) newWorker() { //даже при получении сигнала завершения, горутина дообработает задачи в канале
	for i := range 3 {
		jQ.wg.Add(1)
		go func(n int) {
			defer jQ.wg.Done()
			for {
				jQ.cond.L.Lock()
				for len(jQ.queue) == 0 {
					if jQ.ctx.Err() != nil {
						jQ.cond.L.Unlock()
						fmt.Printf("Worker #%d: no more jobs, ctx done. Exiting.\n", n)
						return
					}
					fmt.Printf("Worker #%d falling asleep...\n", n)
					jQ.cond.Wait()
				}
				job := jQ.queue[0]
				jQ.queue = jQ.queue[1:]
				jQ.cond.L.Unlock()
				fmt.Printf("Worker #%d doing job #%d...\n", n, job)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
}

func main() {
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)
	ctx, temp := context.WithTimeout(context.Background(), 3*time.Second)
	cancel = temp
	jQ := jobQueue{
		queue: []int{},
		cond:  cond,
		wg:    sync.WaitGroup{},
		ctx:   ctx,
	}
	jQ.ctrlC() // запускаем обработчик прерываний
	jQ.generator()
	jQ.newWorker()
	jQ.wg.Wait()
	fmt.Println("All goroutines are finished!")
}
