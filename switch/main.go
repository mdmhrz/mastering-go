package main

import "fmt"

func main() {
	day := "sat"

	switch day {
	case "sat":
		fmt.Println("This is offday")
	case "sun":
		fmt.Println("This is working day")
	}
}
