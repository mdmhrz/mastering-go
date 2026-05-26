package main

import "fmt"

// type metaInfo struct {
// 	phone   int
// 	address string
// }

// type user struct {
// 	name  string
// 	email string
// 	metaInfo
// }

type user struct {
	name string
	age  int
	role string
}

func main() {
	// john := user{"john", "john@gmailcom"}
	// john := user{
	// 	name:  "john",
	// 	email: "john@gmailcom",
	// 	metaInfo: metaInfo{
	// 		phone:   0123,
	// 		address: "165",
	// 	},
	// }
	// john.email = "johnerbaccha@gmail.com"

	// fmt.Printf("%+v", john)

	newUser := func(name string, age int, role string) user {
		return user{
			name: name,
			age:  age,
			role: role,
		}
	}

	fmt.Print(newUser("jamal", 34, "student"))
}
