package main

import "fmt"

func main() {

	fmt.Print("Hello, ")
	fmt.Print("Universe!\n")
	fmt.Print("new line\n")

	fmt.Println("hello!")
	fmt.Println("goodbye")

	name := "John Doe"
	age := 26
	fmt.Printf("my name is %q and my age is %v \n", name, age)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %0.1f points! \n", 225.55)

	// Sprintf
	var str = fmt.Sprintf("my name is %v and age is %v", name, age)
	fmt.Println(str)
}
