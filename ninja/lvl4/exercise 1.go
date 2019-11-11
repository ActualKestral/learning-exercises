package main

import (
	"fmt"
)

// A composite literal making an array then format printing used to range over the array
func main() {
	x := [5]int{42, 43, 44, 45, 46}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
