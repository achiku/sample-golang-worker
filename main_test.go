package main

import (
	"testing"
	"time"
)

func TestNewWorkerPool(t *testing.T) {
	numWorker, interval := 4, time.Duration(5*time.Second)
	wp := NewWorkerPool(numWorker, interval)
	wp.Start()
}
