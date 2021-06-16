package main

import "fmt"

type shape interface {
	area() float64
	circumference() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumference() float64 {
	return s.length * 4
}

func printShape(s shape) {
	fmt.Printf("area of %T is %0.2f \n", s, s.area())
	fmt.Printf("circum of %T is %0.2f \n", s, s.circumference())
}

func main() {
	shapes := []shape{
		square{length: 10.5},
	}

	for _, v := range shapes {
		printShape(v)
	}
}
