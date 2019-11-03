package main

import "fmt"
var i = 0
func main ()  {
		for {
			i++
			if i>12 {
				break				
			} else if i == 13 {
				fmt.Println("13 is also working.")
			}
			if i == 2 {
				fmt.Println("hey guys its working")				
			}
			switch i {
			case (2):
				fmt.Println("Hello Ms Moneypenny")
			}
	}
	switch (i >= 3) {
	case true:
		fmt.Println("Hello Q. Is my switch working yet?")
	}
	
}