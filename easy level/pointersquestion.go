package main

import (
	"fmt"
)

// original
// func changePointer(p *int) {
// 	v := 3
// 	p = &v
// }

func changePointer(p *int) {
	*p = 3
}

// alternative
// func changePointer(p **int) {
//     v := 3
//     *p = &v
// }

// What will this code output, and why?
// How can you change v to 3 through the function without adding a return value?
func main() {
	v := 5
	p := &v
	fmt.Println(*p)

	changePointer(p)
	// alternative
	// changePointer(&p)

	fmt.Println(*p)
}

/*
Response
5
5

Go always passes arguments by value.
In changePointer, the parameter p is a copy of the original pointer.
Reassigning p changes only that local copy, so v remains 5.
To modify the original value, dereference the pointer (*p = 3).
If you need to change the pointer itself, pass **int and modify *p.
*/
