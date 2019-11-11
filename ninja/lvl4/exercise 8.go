package main

import "fmt"

//create a map of type string and value of type []string three records(keys)
func main() {
	M := map[string][]string{
		"Bond James":      []string{"Gadgets", "Martinies", "Women"},
		"Moneypenny Miss": []string{"Bond James", "Literature", "Compter Science"},
		"Dr Evil":         []string{"Being evil", "Ice cream", "Sunsets"},
	}
	fmt.Println(M)

	for k, v := range M {
		fmt.Println("record: ", k)
		for j, f := range v {
			fmt.Printf("\t index position: %v, value: %s \n", j, f)
		}
	}
}
