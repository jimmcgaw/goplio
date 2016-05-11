package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2 // e.g. x = 2
	fmt.Println(x)
}
