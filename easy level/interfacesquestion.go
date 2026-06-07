package main

import (
	"fmt"
)

// What will the code output and why?
type impl struct{}

func (*impl) C() {}

type I interface {
	C()
}

func A() I {
	return nil
}

func B() I {
	var ret *impl

	return ret
}

func main() {
	a := A()
	b := B()

	fmt.Println(a == b)
}

/*
Response
false

a = (nil, nil)
b = (*impl, nil)
(type1 == type2) && (value1 == value2)
(nil, nil) != (*impl, nil)
*/