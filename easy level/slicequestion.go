package main

import (
	"fmt"
)

// 1
// What is the output of this program?
func main() {
	a := []string{"a", "b", "c"}
	b := a[1:2]
	b[0] = "q"

	fmt.Println(a)
}

/*
Response
[a q c]

The slice b does not copy the data.
It references the same underlying array as a.
*/

// 2
// What will this code output, and why?
func mod(a []int) {
	a = append(a, 125)
	for i := range a {
		a[i] = 5
	}

	fmt.Println(a)
}

func main() {
	sl := []int{1, 2, 3, 4, 5}

	mod(sl)

	fmt.Println(sl)
}

/*
Response
[5 5 5 5 5 5]
[1 2 3 4 5]

The slice (sl) has:
len = 5
cap = 5

When executing:
a = append(a, 125)

there is no remaining capacity in the backing array.
As a result, append allocates a new array and copies the elements into it.
After that, a points to the new backing array, while sl still points to the original one.
*/

// 3
// What will this code output, and why?
func mod(a []int) {
	for i := range a {
		a[i] = 5
	}

	fmt.Println(a)
}

func main() {
	sl := make([]int, 4, 8)

	sl[0] = 1
	sl[1] = 2
	sl[2] = 3
	sl[3] = 5

	mod(sl)

	fmt.Println(sl)
}

/*
Response
[5 5 5 5]
[5 5 5 5]

modifies the shared underlying array.
Since both slices reference the same array, the changes are visible both inside and outside the function.
*/
