// 1. Merge N channels.
// 2. If any input channel is closed,
//    all the remaining channels must also be closed.
func case3(channels ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	var once sync.Once

	stop := func() {
		once.Do(func() {
			for _, ch := range channels {
				close(ch)
			}

			close(out)
		})
	}

	for _, ch := range channels {
		wg.Add(1)

		go func(c chan int) {
			defer wg.Done()

			for {
				v, ok := <-c
				if !ok {
					stop()

					return
				}

				select {
				case out <- v:
				default:
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		stop()
	}()

	return out
}