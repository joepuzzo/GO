package main

import "fmt"

func main() {
	var n [3]int

	//Make( array, length, capacity)
	var s []int = make([]int, 3, 5)

	s[0] = 1
	s[1] = 2
	s[2] = 10

	/*----Short hand----*/
	slice1 := []string{
		"Hello",
		"World!",
		"My name",
		"is Joe",
	}

	slice1 = slice1[1:2]

	slice1 = append(slice1, "Bill")

	//Double copy the slice
	slice1 = append(slice1, slice1...)

	//Delete middlestuff
	slice1 = append(slice1[:1], slice[:2]...)
}
