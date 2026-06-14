package main

import (
	"fmt"
	"sync"
	"time"
)

// Explain in detail what happens.
// How long will this program run?
const (
	numRequests = 10_000
)

var (
	count int
	m     sync.Mutex
)

func networkRequest() {
	time.Sleep(time.Millisecond) // Simulates a network request.

	m.Lock()
	count++
	m.Unlock()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			networkRequest()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}

/*
The program launches 10,000 goroutines.

Each goroutine:

Sleeps for approximately 1 ms (simulating I/O).
Acquires the mutex.
Increments the shared counter.
Releases the mutex.
Calls wg.Done().

WaitGroup ensures that main() waits until all goroutines finish.

The mutex protects count from a data race, so the final output is guaranteed to be:

10000
*/
