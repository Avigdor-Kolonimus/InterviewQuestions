package main

import (
	"fmt"
)

// What will this program print?
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic caught:", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Something went wrong!")
	fmt.Println("After panic")
}

/*
Response
Before panic
Panic caught: Something went wrong!

panic() stops normal execution and starts unwinding the stack.
The deferred function runs, recover() catches the panic, and the program exits normally.
The line fmt.Println("After panic") is never executed.
*/
