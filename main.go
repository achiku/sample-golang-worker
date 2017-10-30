package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Worker worker
type Worker struct {
	Interval time.Duration

	id  int
	mux sync.Mutex
	ch  chan struct{}
}

// Work does the job
func (w *Worker) Work() bool {
	return false
}

// Start starts a worker
func (w *Worker) Start() {
	go func() {
		select {
		case <-w.ch:
			log.Printf("id=%d finished job", w.id)
		case <-time.After(w.Interval):
			for {
				if didWork := w.Work(); !didWork {
					break // didn't do any work, go back to sleep
				}
			}
		}
	}()
}

// WorkerPool worker pool
type WorkerPool struct {
	workers []*Worker
	mux     sync.Mutex
}

// Start start workers in a pool
func (wp *WorkerPool) Start() {
	wp.mux.Lock()
	defer wp.mux.Unlock()

	for i, w := range wp.workers {
		log.Printf("start worker %d: %+v", i, w)
	}
}

// NewWorkerPool creates new worker pool
func NewWorkerPool(numWorker int, interval time.Duration) *WorkerPool {
	var workers []*Worker
	for i := 0; i < numWorker-1; i++ {
		w := &Worker{
			Interval: interval,
			id:       i,
			ch:       make(chan struct{}, 1),
		}
		workers = append(workers, w)
	}
	wp := &WorkerPool{
		workers: workers,
	}
	return wp
}

func main() {
	fmt.Println("vim-go")
}
