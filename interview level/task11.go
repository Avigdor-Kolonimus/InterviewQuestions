// Write a function that merges N channels.
// All values from the input channels should be forwarded
// to a single output channel.
func joinChannels(chs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	for _, ch := range chs {
		wg.Add(1)

		go func(c <-chan int) {
			defer wg.Done()

			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}