package main

import (
	"fmt"
	"sync"
)

// time.Sleep is not allowed. It would not be considered a valid interview solution.
// Explain in detail what happens.
// How can this be fixed while keeping the channel unbuffered?
func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(v int) {
			defer wg.Done()
			ch <- v * v
		}(i)
	}
	// wg.Wait()
	go func() {
		wg.Wait()
		close(ch)
	}()

	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)
}

/*
The program deadlocks.
All three goroutines try to send values to an unbuffered channel:
ch <- v * v

Since there is no receiver yet, they block before reaching wg.Done().
As a result, wg.Wait() never returns, and the main goroutine never reaches the loop that reads from the channel.
*/
