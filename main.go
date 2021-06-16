package main

import "fmt"

func main() {

	fmt.Println("Hello Universe")

	// Strings
	var fName string = "Marco"
	var lName = "Polo"
	var blank string

	fmt.Println(fName, lName, blank)

	fName = "Nicolo"
	blank = "Not Blank"

	fmt.Println(fName, lName, blank)

	// Only inside a function in this manner
	another := "Machiavelli"
	fmt.Println(another)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	//var numOne int8 = 25
	//var numTwo int8 = -128
	//var numThree uint8 = 255

}
