package main

import "fmt"

type user struct {
	name       string
	age        int
	isLoggedIn bool
}

func main() {
	myMap := make(map[string]int)

	myMap["userScore"] = 8
	myMap["money"] = 765

	fmt.Println(myMap)
	fmt.Println(myMap["money"])

	userData := map[string]string{
		"name":    "Mobarak",
		"success": "ok",
	}

	delete(userData, "name")

	fmt.Println(userData["name"])

	newUserData := map[string]user{
		"data": {
			name:       "mobarak",
			age:        30,
			isLoggedIn: true,
		},
	}

	fmt.Println(newUserData)

}
