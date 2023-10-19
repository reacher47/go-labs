package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Введите a: ")
	fmt.Scan(&a)
	fmt.Println("Введите b: ")
	fmt.Scan(&b)
	fmt.Println("a + b =", a+b)
}
