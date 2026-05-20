package main

import "fmt"

func main() {
	numbers := [5]int{6, 7, 5, 1, 6}

	// slicedNumber := numbers[1:3] // declared start and end
	// slicedNumber := numbers[:3] // declared end, that means from begining start and end at declared index
	slicedNumber := numbers[2:] // declared start, that means start from declared index till end.

	fmt.Println(slicedNumber)

	fmt.Println("the length of the slice is", len(slicedNumber))   // len means actual data lenth of the slice
	fmt.Println("the capacity of the slice is", cap(slicedNumber)) // cap means capacity of the slice that means starting point is cound and till is count
}
