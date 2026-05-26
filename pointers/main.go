package main

import "fmt"

func main() {
	a := 42
	p := &a // pinting the memory address of the variable by & sign is called pointer

	a = 100
	*p = 500
	a = 700

	fmt.Println(a)
	fmt.Println(*p) // de-reference the pointer value
}
