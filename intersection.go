package main

import (
	"fmt"
)

func intersection(a, b []int) []int {
	var result []int
	counter := make(map[int]int)

	for _, elem := range a {
		counter[elem]++
	}

	for _, elem := range b {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem] -= 1
			result = append(result, elem)
		}
	}

	return result
}

func main() {
	// Case 1
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}
	// [2, 23]
	fmt.Printf("%v\n", intersection(a, b))

	// Case 2
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	// [1, 1, 1]
	fmt.Printf("%v\n", intersection(a, b))
}
