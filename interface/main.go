package main

import "fmt"

type Animal interface {
	speak()
}

type Dog struct {
}

func (d *Dog) speak() {
	fmt.Println("Woof")
}

func makeSound(a Animal) {
	a.speak()
}

func main() {

	dexter := Dog{}

	fmt.Println(dexter)

}
