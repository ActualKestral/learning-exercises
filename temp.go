package main

import "fmt"

func main ()  {
	i := 0
		for {
			i++
			if i>12 {
				break				
			}
			if i == 2 {
				fmt.Println("hey guys its working")				
			}
			fmt.Println(i)
	}
}