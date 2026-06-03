package main

import "fmt"

func main() {
	myMap := map[string]string{
		"name":    "mobarak",
		"success": "ok",
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	for _, value := range myMap {
		fmt.Println(value)
	}

	myArr := [3]string{
		"yellow",
		"red",
		"white",
	}

	for i, value := range myArr {
		fmt.Println(i+1, value)
	}
}
