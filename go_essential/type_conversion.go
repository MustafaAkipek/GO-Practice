package main

import "fmt"

func main() {

	// Part 1 (cannot sum int + float)
	
	/*x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(x + int(y))

	fmt.Printf("%v %T\n", y, y)*/

	// Part 2 (note: float32 --> float64)

	/*
	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(float64(x) + y)
	*/

	// Part 3 (cannot change an string into an int)

	/*
	x := 10
	y := "10"

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(x + int(y))
	*/

	// Part 4 

	num1 := 106
	str1 := string(rune(num1))

	fmt.Printf("%v %T\n", num1, num1)
	fmt.Println()
	fmt.Printf("%v %T\n", str1, str1)
}