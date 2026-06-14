// Using time.Sleep is not allowed. It would not be considered a valid interview solution.

// 1. Sometimes zeros are returned. What is the problem? Fix it.
// 2. If bank_network_call takes 5 seconds to complete, how long will balance() take to finish?
func balance() int {
	x := make(map[int]int, 1)
	var m sync.Mutex

	// call bank
	// for i := 0; i < 5; i++ {
	// 	i := i
	// 	go func() {
	// 		m.Lock()
	// 		b := bank_network_call(i)

	// 		x[i] = b
	// 		m.Unlock()
	// 	}()
	// }

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		i := i

		wg.Add(1)
		go func() {
			defer wg.Done()

			b := bank_network_call(i)

			m.Lock()
			x[i] = b
			m.Unlock()
		}()
	}

	wg.Wait()

	// Compute the sum of all values in the map and return it.
	return sumOfMap
}

/*
- The bug is that balance() does not wait for the goroutines to finish, so the map may be incomplete when the sum is calculated.
- Fix it by using a sync.WaitGroup.
- The mutex should protect only the map write, not the network call.
- Current code: ~25 seconds (5 requests × 5 seconds, serialized by the mutex).
- After moving bank_network_call() outside the critical section: ~5 seconds (requests run concurrently).
*/