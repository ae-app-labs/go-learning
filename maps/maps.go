package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// Looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		9678564: "john",
		4534534: "cathy",
		7345234: "mark",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[4534534])

	phonebook[4534534] = "catherine"
	fmt.Println(phonebook[4534534])
}
