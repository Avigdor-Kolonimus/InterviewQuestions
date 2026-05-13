package main

import (
	"context"
	"fmt"
	"sync"
)

const (
	numWorkers = 3
	numJobs    = 5
)

type Result struct {
	Value int
	Err   error
}

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- Result, f func(int) (int, error), wg *sync.WaitGroup) {
	defer wg.Done()

	defer func() {
		if r := recover(); r != nil {
			results <- Result{
				Err: fmt.Errorf("panic: %v", r),
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d canceled\n", id)
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

			func() {
				defer func() {
					if r := recover(); r != nil {
						results <- Result{
							Err: fmt.Errorf("worker %d panic: %v", id, r),
						}
					}
				}()

				fmt.Printf("worker %d processed %d -> ?\n", id, job)

				value, err := f(job)

				results <- Result{
					Value: value,
					Err:   err,
				}
			}()
		}
	}
}

func process(n int) (int, error) {
	if n == 4 {
		panic("boom")
	}

	if n == 3 {
		return 0, fmt.Errorf("cannot process %d", n)
	}

	return n * 10, nil
}

func main() {
	jobs := make(chan int)
	results := make(chan Result)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)

		go worker(ctx, i, jobs, results, process, &wg)
	}

	// producer
	go func() {
		defer close(jobs)

		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
	}()

	// close results after all workers done
	go func() {
		wg.Wait()
		close(results)
	}()

	// consume results
	for result := range results {
		fmt.Println("Value: ", result.Value, " Error: ", result.Err)
	}
}
