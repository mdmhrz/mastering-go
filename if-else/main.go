package main

import "fmt"

func main() {
	marks := 50

	if marks >= 80 {
		fmt.Println("You've got GPA-5.00")
	} else if marks >= 60 {
		fmt.Println("You've got GAP-4.00")
	} else {
		fmt.Println("You are an average student")
	}
}
