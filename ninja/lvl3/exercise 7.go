package main

import "fmt"

func main() {
	x := 55
	if x == 99 {
		fmt.Println("The great one.")
	} else if x <99 { 
		fmt.Println("This player is not the great one.")
	} else {
		fmt.Println("Unallowed player number")
	}
}