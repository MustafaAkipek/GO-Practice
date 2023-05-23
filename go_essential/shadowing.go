package main

import "fmt"

func main() {

	x := 5

	if true{
		x := 10
		x++
		fmt.Println(x) // 11
	}

	fmt.Println(x) // 5

	if true{
		x = 10
		x++
		fmt.Println(x) // 11
	}

	fmt.Println(x) // 11
}