package main

import "fmt"

func main() {

	/* const x = 3
	const y = 5.6
	const z = true
	const t = "test"

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t) */

	const x = 3
	const y = 5.6

	fmt.Printf("%T, %v\n", x, x) // if we assign type of variable(line 17), get error(line 28)
	fmt.Printf("%T, %v\n", y, y) // if we assign type of variable(line 18), get error(line 28)
	fmt.Printf("%T, %v\n", x + y, x + y)
}