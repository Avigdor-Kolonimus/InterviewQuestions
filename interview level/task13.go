package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// There is a function `unpredictableFunc` that takes an unpredictable
// amount of time to complete and returns a number.
// Its implementation cannot be changed (assume it performs a network request).
//
// Implement a wrapper function `predictableFunc` that runs with a fixed timeout
// (e.g., 1 second).
//
// If the long-running function completes within the timeout,
// return its result.
//
// Otherwise, return an error. In that case, the actual result is not important.
//
// Additionally, measure and log how long the function execution took.
//
// The wrapper function's signature may be changed.

func init() {
	rand.Seed(time.Now().UnixNano())
}

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

// func predictableFunc() int64 {

// }

// func main() {
// 	fmt.Println("started")
// 	fmt.Println(predictableFunc())
// }

func predictableFunc(timeout time.Duration) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	start := time.Now()

	resultCh := make(chan int64, 1)

	go func() {
		resultCh <- unpredictableFunc()
	}()

	select {
	case res := <-resultCh:
		fmt.Println("execution time:", time.Since(start))
		return res, nil

	case <-ctx.Done():
		fmt.Println("execution time:", time.Since(start))
		return 0, ctx.Err()
	}
}

func main() {
	fmt.Println("started")

	res, err := predictableFunc(time.Second)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(res)
}
