package main

import (
	"fmt"
)

// Using time.Sleep is not allowed. It would not be considered a valid interview answer.
// 1. How will this code behave?
// 2. How can we make it print only from the first channel (ch)?
// 2.1 You are not allowed to remove anything from the existing code. You may only add new code.
func main() {
	ch := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)

	go func() {
		ch <- true
	}()

	// go func() {
	// 	ch2 <- true
	// }()
	// go func() {
	// 	ch3 <- true
	// }()

	start := make(chan struct{})
	go func() {
		<-start
		ch2 <- true
	}()

	go func() {
		<-start
		ch3 <- true
	}()

	select {
	case <-ch:
		fmt.Printf("val from ch")
	case <-ch2:
		fmt.Printf("val from ch2")
	case <-ch3:
		fmt.Printf("val from ch3")
	}

	close(start)
}

/*
Response
When multiple cases are ready simultaneously, Go chooses one of them pseudo-randomly.
*/
