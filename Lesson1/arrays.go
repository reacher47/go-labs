package main

import (
	"fmt"
)

func main() {
	array1 := [3]int{1, 3, 7}
	array2 := [2]int{}

	//определение размера при объявлении
	array3 := [...]int{1, 3, 6, 2, 5}

	// не принято объявлять по кодстайлу
	var array4 [2]int = [2]int{4, 7}
	var array5 = [3]int{2, 3, 6}

	fmt.Println(array1, array2, array3, array4, array5)

	// Слайсы (надмножества массивов, могут расширяться)

	var buf0 []int
	buf1 := []int{}
	buf2 := []int{42}
	buf3 := make([]int, 0)
	buf4 := make([]int, 5)     // len=5, cap=5
	buf5 := make([]int, 5, 10) // len=5, cap=10

	fmt.Println(buf0, buf1, buf2, buf3, buf4, buf5)

	// получение среза
	buf := []int{1, 2, 3, 4, 5}
	sl1 := buf[1:4]
	sl2 := buf[:2]
	sl3 := buf[2:]

	fmt.Println(sl1, sl2, sl3)

}
