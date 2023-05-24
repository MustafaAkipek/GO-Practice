package main

import "fmt"

func main() {

	x, y := 15, 10
	z := 5.0 / 2 // float / int --> float

	fmt.Printf("%T, %v\n", x, x)             // int, 15
	fmt.Printf("%T, %v\n", y, y)             // int, 10
	fmt.Printf("%T, %v\n", (x + y), (x + y)) // int, 25
	fmt.Printf("%T, %v\n", (x - y), (x - y)) // int, 5
	fmt.Printf("%T, %v\n", (x * y), (x * y)) // int, 150
	fmt.Printf("%T, %v\n", (x / y), (x / y)) // int, 1 (int / int --> int)
	fmt.Printf("%T, %v\n", (x / y), (x / y)) // int, 1 (int / int --> int)
	fmt.Printf("%T, %v\n", (x % y), (x % y)) // int, 5
	fmt.Printf("%T, %v\n", z, z)             // float64, 2.5

	// note: go(lang) has only postfix (w++ or w--)

	w := 10
	fmt.Println(w)

	w++ // increment ++
	fmt.Println(w)

	w-- // decrement --
	fmt.Println(w)
}
