package main

import "fmt"

// Used composite literal to create a slice of type int.

func main() {
	x := []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
