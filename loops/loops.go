package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	names := []string{"david", "amber", "maria", "john"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}
}
