package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 1
// What will this code output, and why?
// How would you fix it?
func main() {
	// var counter int
	var counter int32
	var wg sync.WaitGroup

	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		counter++
	// 	}()
	// }

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()

	}

	wg.Wait()
	fmt.Println(counter)
}

/*
The output is non-deterministic
*/

// 2
// What will this code output, and why?
// How would you fix it while keeping the channel unbuffered?
func main() {
	ch := make(chan int)
	// ch <- 1
	// go func() {
	// 	fmt.Println(<-ch)
	// }()
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()

	ch <- 1

	wg.Wait()
}

/*
fatal error: all goroutines are asleep - deadlock!
ch <- 1
*/

// 3
// Explain in detail what is happening here.
// (Hint: why does it panic?)
// How would you make it work?
// Using time.Sleep is not allowed. It would not be considered a valid answer in an interview.
func main() {
	x := make(map[int]int, 1)
	var mu sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		x[1] = 2
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		x[3] = 7
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		x[123] = 10
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		x[1] = 2
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		x[34] = 7
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		x[1432] = 10
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		x[1] = 2
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		x[100] = 7
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		x[34] = 10
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu.Lock()
		x[1] = 2
		mu.Unlock()
	}()

	time.Sleep(100 * time.Millisecond) // Block for 100 milliseconds.
	wg.Wait()

	mu.RLock()
	fmt.Println("x[1] =", x[1])
	mu.RUnlock()
}

/*
The program launches multiple goroutines that concurrently write to the same map.
Go maps are not thread-safe, so concurrent writes can corrupt the map's internal state.
The runtime detects this and usually panics with fatal error: concurrent map writes.
*/
