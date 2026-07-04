package main

import (
	"fmt"
)

// Fix the function so that it works.
// You cannot change the function signature.
func printNumber(ptrToNumber interface{}) {
	// if ptrToNumber != nil {
	// 	fmt.Println(*ptrToNumber.(*int))
	// } else {
	// 	fmt.Println("nil")
	// }

	if ptr, ok := ptrToNumber.(*int); ok && ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("nil")
	}
}

/*
The problem is that an interface containing a typed nil pointer is not equal to nil.
Therefore, the check ptrToNumber != nil is true even when pv is nil, and the type assertion succeeds, but dereferencing the pointer causes a panic.
*/
func main() {
	v := 10
	printNumber(&v)
	var pv *int

	printNumber(pv)
	pv = &v
	printNumber(pv)
}
