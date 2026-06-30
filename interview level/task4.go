package main

import (
	"fmt"
)

type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil
}

func B() I {
	var ret *impl
	return ret
}

// What will this program print?
func main() {
	a := A()
	b := B()
	fmt.Println(a == b)
}

/*
Answer
false

A() returns a nil interface (type = nil, value = nil).
B() returns an interface containing a typed nil pointer (type = *impl, value = nil).
Although both contain nil values, their dynamic types are different, so the interfaces are not equal.
a: (nil, nil)
b: (*impl, nil)
Therefore, a == b evaluates to false.
*/
