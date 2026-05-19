package main

import "fmt"

func coffeeOrder(name string, CoffeeName string, amount int) string {

	formattedReturn := fmt.Sprintf(
		"Order for %s: %s costs %d taka ",
		name,
		CoffeeName,
		amount,
	)

	return formattedReturn
}

func main() {

	firstOder := coffeeOrder("razu", "Black Coffee", 150)
	fmt.Println(firstOder)

}
