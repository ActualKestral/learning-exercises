package main

import "fmt"

func main ()  {
	for i := 33; i <= 126; i++ {
		fmt.Printf("%d\n", i)
		for n := 0; n <= 2; n++ {
			fmt.Printf("\t %#U\n", i)		
		}
	}
}