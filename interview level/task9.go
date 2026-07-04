package main

import (
	"sync"
)

// 1. Sometimes zeros are returned. What's the problem? Fix it.
// 2. If bank_network_call takes 5 seconds to complete,
//    how long will balance() take? How would you fix the problem?
// 3. Assume bank_network_call also returns an error.
//    If at least one call fails, balance() should return an error.
/*
func balance() int {
	x := make(map[int]int, 1)
	var m sync.Mutex

	// Call the bank service.
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			m.Lock()
			b := bank_network_call(i)
			x[i] = b
			m.Unlock()
		}()
	}

	// Calculate the sum of all values in the map and return it.
	return sumOfMap
}
*/
func balance() (int, error) {
	x := make(map[int]int)
	var m sync.Mutex
	var wg sync.WaitGroup

	errCh := make(chan error, 5)

	// Call the bank service.
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)

		go func() {
			defer wg.Done()

			b, err := bank_network_call(i)
			if err != nil {
				errCh <- err
				return
			}

			m.Lock()
			x[i] = b
			m.Unlock()
		}()
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return 0, err
		}
	}

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum, nil
}

/*
Answer
1) Why do zeros sometimes appear?
	balance() returns before the goroutines finish. Use a sync.WaitGroup to wait for all of them.
2) How long does it take?
	About 25 seconds, because the mutex is held during the 5-second bank_network_call, forcing the calls to execute one after another.
	Move the network call outside the critical section, leaving only the map write protected by the mutex.
	Then the function takes about 5 seconds.
3) How to handle errors?
	Make bank_network_call return (int, error), collect errors (e.g., via a channel or errgroup.Group), and if any call fails, return an error from balance().
*/
