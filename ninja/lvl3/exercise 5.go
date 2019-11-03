package main

import "fmt"

func main() {
	i := 10
	for {
		if i >100 {
			break
		}
		fmt.Printf("When %v is divided by 4, the remainder is %v\n", i, i%4)
			i++
	}
}