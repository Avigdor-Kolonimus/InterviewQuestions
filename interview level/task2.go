package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func changeName(person *Person) {
	person = &Person{
		Name: "Alice",
	}
}

// What will this program print? How can it be fixed?
func main() {
	person := &Person{
		Name: "Bob",
	}

	fmt.Println(person.Name)
	changeName(person)
	fmt.Println(person.Name)
}

/*
Response
Bob
Bob

person is a pointer, but Go passes function arguments by value.
Inside changeName, only a copy of the pointer is modified, while the original pointer in main still points to the "Bob" object.
Therefore, the output is Bob and Bob.
*/

// Answer:

func changeName(person **Person) {
	*person = &Person{
		Name: "Alice",
	}
}

func main() {
	person := &Person{
		Name: "Bob",
	}

	fmt.Println(person.Name)
	changeName(&person)
	fmt.Println(person.Name)
}

/* Alternatively
func changeName(person *Person) {
	person.Name = "Alice"
}
*/

/*
Response
Bob
Alice
*/
