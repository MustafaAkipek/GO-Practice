package main

import "fmt"

func main() {

	grade := 5

	switch grade {
	case 5:
		fmt.Println("Pekiyi")
	case 4:
		fmt.Println("Iyi")
	case 3:
		fmt.Println("Orta")
	case 2:
		fmt.Println("Gecer")
	case 1:
		fmt.Println("Basarisiz")
	default:
		fmt.Println("Gecersiz not")
	}
}