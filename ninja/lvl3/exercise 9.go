package main

import "fmt"

func main() {
	favsport := "icehockey"

	switch favsport {
	case "surfing": 
		fmt.Println("Take me to hawaii")
	case "scubadiving":
		fmt.Println("I love been under the water")
	case "icehockey": 
		fmt.Println("This is the best sport")
	}
}