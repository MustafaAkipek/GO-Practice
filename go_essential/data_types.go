package main

import "fmt"

func main() {

	var name = "Mustafa"
	var age uint16 = 400
	var isMarried bool = true
	var weight float32 = 72.5

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)

	fmt.Printf("%T\n", name) // "\n" new line
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
	fmt.Printf("%T\n", weight)
}