package main

import (
	"fmt"
	"reflect"
)

// Add code that prints the type of the variable `whoami`.
// func printType(whoami interface{}) {
// }
func printType(whoami interface{}) {
	// optional 1
	fmt.Printf("Type of whoami: %T\n", whoami)

	// optional 2
	fmt.Printf("Type of whoami: %v\n", reflect.TypeOf(whoami))

	// optional 3
	switch v := whoami.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	default:
		fmt.Printf("unknown type: %T\n", v)
	}
}

func main() {
	printType(42)
	printType("im string")
	printType(true)
}
