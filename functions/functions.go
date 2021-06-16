package main

import (
	"fmt"
	"math"
)

func greet(n string) {
	fmt.Printf("Hello %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Good bye %v \n", n)
}

func main() {

	names := []string{"cloud", "tifa", "barret"}
	cycleNames(names, greet)
	cycleNames(names, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(16.8)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %.3f and circle 2 is %0.3f", a1, a2)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}
