package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrErrorsNumGoroutines = errors.New("wrong num of goroutines")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, routines, maxErrors int) error {
	if routines <= 0 {
		return ErrErrorsNumGoroutines
	}
	if maxErrors <= 0 {
		maxErrors = len(tasks) + 1
	}
	ch := make(chan Task)
	wg := sync.WaitGroup{}
	wg.Add(routines)
	var errorsCount int32
	for i := 0; i < routines; i++ {
		go func() {
			defer wg.Done()
			for task := range ch {
				err := task()
				if err != nil {
					atomic.AddInt32(&errorsCount, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errorsCount) >= int32(maxErrors) {
			break
		}
		ch <- task
	}

	close(ch)
	wg.Wait()

	if errorsCount >= int32(maxErrors) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
