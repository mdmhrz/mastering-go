package main

import "fmt"

func main() {
	// Basic for loop
	// count := 11
	// for i := 1; i < count; i++ {
	// 	fmt.Println(i)
	// }

	// Javascript while loop style loop
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// Break, continue
	// count := 11
	// for i := 1; i < count; i++ {
	// 	fmt.Println(i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	count := 10
	for i := 1; i <= count; i++ {

		if i%2 != 0 {
			continue
		}

		fmt.Println(i)
	}

}
