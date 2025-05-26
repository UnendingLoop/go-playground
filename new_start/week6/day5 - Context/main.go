package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type urlProcessor struct {
	wg                  sync.WaitGroup
	ctx                 context.Context
	cancel              context.CancelFunc
	urls                chan string
	success             atomic.Int64
	errors              atomic.Int64
	currentWorkersCount atomic.Int64
	minWorkers          int
	maxWorkers          int
}

func (p *urlProcessor) gracefulCtrlC() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		defer close(sig)
		<-sig
		fmt.Println("Interrupt received")
		p.cancel()
	}()
}

func newURLProcessor(timeout time.Duration, bufSize int) *urlProcessor {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	return &urlProcessor{
		ctx:        ctx,
		cancel:     cancel,
		urls:       make(chan string, bufSize),
		minWorkers: bufSize / 20,
		maxWorkers: bufSize / 10,
	}
}

func (p *urlProcessor) startGenerator(n int) {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		defer close(p.urls)
		for i := range n {
			select {
			case <-p.ctx.Done():
				fmt.Printf("Generator: cancel received: %v. Shutting down...\n", p.ctx.Err())
				return
			case p.urls <- fmt.Sprintf("url_%d", i):
			}
		}
		return
	}()
}

func (p *urlProcessor) managerWorkers() {
	p.wg.Add(1)
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		defer p.wg.Done()
		for i := 0; i < p.minWorkers; i++ { //запускаем минимальное кол-во воркеров
			p.startWorker()
		}
		for { //запускаем цикл проверки каналов и запуск воркеров по необходимости
			select {
			case <-p.ctx.Done():
				return
			case <-ticker.C:
				if len(p.urls) > cap(p.urls)/2 && int(p.currentWorkersCount.Load()) < p.maxWorkers {
					p.startWorker()
				}
			}
		}
	}()
}

func (p *urlProcessor) startWorker() {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		p.currentWorkersCount.Add(1)
		current := p.currentWorkersCount.Load()
		for {
			select {
			case <-p.ctx.Done():
				fmt.Printf("Worker #%d couldn't finish job: %v. Shutting down...\n", current, p.ctx.Err())
				return
			default:
				if len(p.urls) > cap(p.urls)/10 && int(p.currentWorkersCount.Load()) > p.minWorkers {
					p.currentWorkersCount.Add(-1)
					fmt.Printf("Job-channel is empty, stopping worker #%d.\n", current)
					return
					//Или более надёжно — учитывать длительное отсутствие работы через таймер (time.After в select).
				}
				job, ok := <-p.urls
				if !ok {
					p.currentWorkersCount.Add(-1)
					fmt.Printf("Job channel is closed, finishing worker #%d. BYE!\n", current)
					return
				}
				num, err := strconv.Atoi(job[4:])
				if err != nil {
					p.errors.Add(1)
					fmt.Printf("Worker #%d found error while StrConv: %v", current, err)
					continue
				}
				if num%5 == 0 { //обеспечиваем 20% ошибок
					p.errors.Add(1)
					continue
				}
				fmt.Printf("Worker #%d Processed URL #%d\n", current, num)
				time.Sleep(50 * time.Millisecond)
				p.success.Add(1)
			}
		}
	}()
}

func main() {
	p := newURLProcessor(8*time.Second, 100)
	defer p.cancel()
	p.gracefulCtrlC()
	p.startGenerator(3000)
	//time.Sleep(100 * time.Millisecond) //даем генератору время нагеренить значения в канал, чтобы managerworkers смог запустить воркеры
	p.managerWorkers()
	p.wg.Wait()
	fmt.Printf("Results are:\nSuccessfully processed: %d,\nErrors: %d.\n", p.success.Load(), p.errors.Load())

}
