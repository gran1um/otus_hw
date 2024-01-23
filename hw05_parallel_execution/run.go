package hw05parallelexecution

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

type Task func() error

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

func Run(tasks []Task, n int, m int) error {
	if n <= 0 {
		return errors.New("number of goroutines must be positive")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	tasksChan := make(chan Task, len(tasks))
	var errorsCount int32

	startWorkers(ctx, n, &wg, tasksChan, &errorsCount, m)

	distributeTasks(ctx, tasksChan, tasks, &errorsCount, m)

	close(tasksChan)
	wg.Wait()

	return checkErrors(&errorsCount, m)
}

func startWorkers(ctx context.Context, n int, wg *sync.WaitGroup, tasksChan <-chan Task, errorsCount *int32, m int) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(ctx, wg, tasksChan, errorsCount, m)
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, tasksChan <-chan Task, errorsCount *int32, m int) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-tasksChan:
			if !ok {
				return
			}
			if err := task(); err != nil && m > 0 && atomic.AddInt32(errorsCount, 1) >= int32(m) {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func distributeTasks(ctx context.Context, tasksChan chan<- Task, tasks []Task, errorsCount *int32, m int) {
	for _, task := range tasks {
		if m > 0 && atomic.LoadInt32(errorsCount) >= int32(m) {
			break
		}
		select {
		case tasksChan <- task:
		case <-ctx.Done():
			break
		}
	}
}

func checkErrors(errorsCount *int32, m int) error {
	if (m <= 0 && atomic.LoadInt32(errorsCount) > 0) || (m > 0 && atomic.LoadInt32(errorsCount) >= int32(m)) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
