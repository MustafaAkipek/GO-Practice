package main

import "fmt"

func main() {

	x := 10

	if x := -5; x < 0 {
		fmt.Println(x, "negatif sayidir")
	} else if x > 0 {
		fmt.Println(x, "pozitif sayidir")
	} else {
		fmt.Println(x, "notr sayidir")
	}

	fmt.Println(x)
}