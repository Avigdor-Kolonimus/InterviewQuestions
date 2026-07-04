package main

import (
	"fmt"
)

const (
	maxLenght = 3
)

var letters = []rune{'a', 'b', 'z', 'f'}

func main() {
	base := len(letters)
	for length := 1; length <= maxLenght; length++ {
		total := 1
		for i := 0; i < length; i++ {
			total *= base
		}

		for num := 0; num < total; num++ {
			x := num
			word := make([]rune, length)

			for pos := length - 1; pos >= 0; pos-- {
				word[pos] = letters[x%base]
				x /= base
			}

			fmt.Println(string(word))
		}
	}
}
