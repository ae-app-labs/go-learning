package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput(reader, "Create a new Bill name: ")

	myBill := newBill(name)
	fmt.Println("Created the bill -", name)
	return myBill
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput(reader, "Choose option (a - add item, s - save bill, t - add tip): ")

	switch opt {
	case "a":
		name, _ := getInput(reader, "Enter item name: ")
		price, _ := getInput(reader, "Enter item price: ")

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput(reader, "Enter tip: ")
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added - ", t)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("Billed saved - ", b.name)
	default:
		fmt.Println("Not a valid option...")
		promptOptions(b)
	}
}

func main() {

	mybill := createBill()
	promptOptions(mybill)

	fmt.Println(mybill.format())
}
