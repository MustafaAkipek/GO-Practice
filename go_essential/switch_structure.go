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

	// fallthrough

	switch x := 25; {
	case x < 20:
		fmt.Printf("%d kucuktur 20\n", x)
		fallthrough
	
	case x < 50:
		fmt.Printf("%d kucuktur 50\n", x)
		fallthrough

	case x < 100:
		fmt.Printf("%d kucuktur 100\n", x)
		fallthrough

	case x < 200:
		fmt.Printf("%d kucuktur 200\n", x)
	}
}