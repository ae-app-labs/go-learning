package main

import (
	"fmt"
)

func main() {

	age := 30
	fmt.Println(age == 45)

	names := []string{"john", "david", "washington", "doe", "jane"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}

		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break
		}

		fmt.Printf("The vaue at pos %v is %v \n", index, value)
	}
}
