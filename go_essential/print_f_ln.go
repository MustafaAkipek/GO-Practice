package main

import "fmt"

func main() {

	// (Print - Println) - Printf

	/* fmt.Println("Merhaba")
	fmt.Print("Merhaba")
	fmt.Println("") // empty line
	fmt.Printf("Merhaba") */

	// name := "Mustafa"

	/* fmt.Print(name)
	fmt.Println("")
	fmt.Println(name)
	fmt.Printf(name) */

	/* fmt.Print("Benim Adim ", name)
	fmt.Println("")
	fmt.Println("Benim Adim", name)
	fmt.Printf("Benim Adim %v %T", name, name) */ // printf --> format method, %v --> value, %T --> type

	/* x := 100
	y := 20
	z := 30

	fmt.Printf("%b %d %o", x, y, z) */ // %b --> binary, d --> decimal, o --> octav

	name, age := "Mustafa", 19

	fmt.Print("Benim Adim ", name, " ve ben ", age, " yasindayim.\n") // \n --> enter new line
	fmt.Println("Benim Adim", name, "ve ben", age, "yasindayim.")
	fmt.Printf("Benim Adim %v ve ben %v yasindayim.", name, age)
}