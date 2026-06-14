package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	timeOut = time.Second
)

func init() {
	rand.NewSource(time.Now().UnixNano()) // used to generate random numbers; you don't need to know more for this task
}

// Using time.Sleep is not allowed. It would not be considered a valid interview solution.
// There is a function unpredictableFunc that takes an unpredictable amount of time
// and returns a number.
// Its implementation cannot be modified (assume it performs a network request).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000) // random number
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

// You need to implement a wrapper function predictableFunc with a fixed timeout (for example, 1 second).
// If the long-running function completes within the timeout, return its result.
// Otherwise, return an error. In that case, we don't care about the eventual result.

// func predictableFunc() int64 {
// }

func predictableFunc() (int64, error) {
	resultCh := make(chan int64, 1)

	go func() {
		resultCh <- unpredictableFunc()
	}()

	select {
	case result := <-resultCh:
		return result, nil

	case <-time.After(timeOut):
		return 0, fmt.Errorf("timeout exceeded")
	}
}

func main() {
	fmt.Println("started")

	// fmt.Println(predictableFunc())
	result, err := predictableFunc()
	if err != nil {
		fmt.Println("error:", err)

		return
	}

	fmt.Println(result)
}
