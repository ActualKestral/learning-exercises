package main

import (
	"fmt"
)

//This is deleting from a slice using append and slicing. There is no straight delete/remove from a slice.
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x[:3], x[6:]...) // [42, 43, 44, 48, 49, 50, 51]
	fmt.Println(x)
}
