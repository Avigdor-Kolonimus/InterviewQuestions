package main

import (
	"fmt"
)

// 1
// What is happening?
// How can it be made to work?
func main() {
	// var m map[string]int
	// panic: assignment to entry in nil map
	m := make(map[string]int)

	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
Response
from 1
the 2
best 1
language 1
in 1
hello 1
world 2
*/
