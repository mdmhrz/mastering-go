package main

import "fmt"

func main() {
	// var numbers [6]int // array initialize without valuse
	// numbers[0] = 5
	// fmt.Println(len(number))

	numbers := [6]int{5, 6, 7, 8, 9, 1} // array initialize with numbers
	
	var total int

	for i := 0; i < len(numbers); i++ {
		// fmt.Println(numbers[i])
		total = total + numbers[i]
	}
	fmt.Printf("total fo array %d & length of array is %d\n", total, len(numbers))
}
