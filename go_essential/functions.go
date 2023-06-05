package main

import "fmt"

func main() {

	fmt.Println(sum(5, 10))

	hello("Mustafa", 19)
}

func sum(x int, y int) int {
	return x + y
}

func hello(name string, age float32) {
	fmt.Println("Benim Adim", name, age, "yasindayim.")
}