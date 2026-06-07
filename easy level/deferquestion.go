package main

import (
	"fmt"
)

// What is the output of this code, and why?
func main() {
	fmt.Println("start")
	
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("end")
}
/*
Response
start
end
3
2
1

In Go, the defer statement postpones the execution of a function until 
the surrounding function returns (main in this case).
*/