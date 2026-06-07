package main

import (
	"fmt"
)

// Explain in detail what is happening here.
// How would you make it work?
func main() {
	// str := "Привет"
	// str[2] = 'e'

	runes := []rune("Привет")
	runes[2] = 'e'
	str := string(runes)

	fmt.Println(str)
}

/*
Response
cannot assign to str[2] (neither addressable nor a map index expression)

The code does not compile because strings in Go are immutable.
Additionally, string indexing operates on UTF-8 bytes rather than Unicode characters.
To modify a character, convert the string to a []rune, update the desired rune, and convert it back to a string.
*/
