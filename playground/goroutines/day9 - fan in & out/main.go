package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type pipeline struct {
	queue     chan int
	result    chan result
	ctx       context.Context
	cancel    context.CancelFunc
	workersWG sync.WaitGroup
	servWG    sync.WaitGroup
}

type result struct {
	input    int
	workerID int
	output   int
}

func (job *pipeline) ctrlC() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		fmt.Println("\n!!! Received interrupt from OS. Starting shutdown sequence...")
		job.cancel()
	}()
}

func (job *pipeline) generetor(queueSize int) {
	job.servWG.Add(1)
	go func() {
		defer job.servWG.Done()
		defer close(job.queue)

		for i := range queueSize {
			select {
			case <-job.ctx.Done():
				fmt.Printf("Generator received shutdown signal: %v. Finishing work...\n", job.ctx.Err())
				return
			case job.queue <- i:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

func (job *pipeline) newWorker(workersCount int) { //завершается только по закрытии канала queue, чтобы обработать остаток очереди
	for i := range workersCount {
		job.workersWG.Add(1)
		go func(n int) {
			defer job.workersWG.Done()
			for {

				j, ok := <-job.queue
				if !ok {
					fmt.Printf("Worker #%d finished working...\n", n)
					return
				}
				fmt.Printf("Worker #%d received job #%d. Working...\n", n, j)
				job.result <- result{
					input:    j,
					workerID: n,
					output:   j * j,
				}
				time.Sleep(100 * time.Millisecond) //имитация обработки данных
			}
		}(i)
	}
}

func (job *pipeline) consumer() { //завершается только по закрытии канала result, чтобы обработать остаток очереди
	job.servWG.Add(1)
	go func() {
		defer job.servWG.Done()
		results := []int{}
		for {
			r, ok := <-job.result
			if !ok {
				fmt.Printf("Consumer finished working. Printing results:\n %v\n", results)
				return
			}
			results = append(results, r.output)
		}
	}()
}

func (job *pipeline) provisor() {
	job.servWG.Add(1)
	go func() {
		defer job.servWG.Done()
		job.workersWG.Wait()
		fmt.Println("All workers are down. Closing the result channel...")
		close(job.result)
	}()
}

func (job *pipeline) run() {
	job.ctrlC()
	job.generetor(100)
	job.newWorker(3)
	job.consumer()
	job.provisor()
	job.servWG.Wait()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	job := pipeline{
		queue:     make(chan int, 5),
		ctx:       ctx,
		result:    make(chan result, 10),
		cancel:    cancel,
		servWG:    sync.WaitGroup{},
		workersWG: sync.WaitGroup{},
	}
	job.run()
	fmt.Println("All is done! Bye")
}
