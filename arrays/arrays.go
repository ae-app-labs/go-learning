package main

import "fmt"

func main() {
	// arrays have a fixed size
	var ages [3]int = [3]int{20, 25, 30}
	//var ages = [3] int {21, 31, 41}

	fmt.Println(ages, len(ages))

	names := [4]string{"one", "two", "three", "four"}
	names[1] = "Drei"
	fmt.Println(names, len(names))

	// Slices
	var scores = []int{100, 50, 60}
	scores[2] = 80
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// Slice Ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "Cooper")
	fmt.Println(rangeOne)
}
