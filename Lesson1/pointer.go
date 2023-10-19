package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a

	*b += 1
	println(*b)
}

func test() *int {
	a := 3
	b := &a
	fmt.Println(a, &b)
	return b
}
