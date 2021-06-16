package main

import (
	"fmt"
	"strings"
)

func main() {

	n1, n2 := getInitials("john doe")
	fmt.Println(n1, n2)

	o1, o2 := getInitials("cloud strife")
	fmt.Println(o1, o2)
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	a := strings.Split(s, " ")

	var initials []string
	for _, v := range a {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	} else {
		return initials[0], "_"
	}
}
