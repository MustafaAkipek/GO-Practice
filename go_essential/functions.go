package main

import "fmt"

func main() {

	fmt.Println(sum(5, 10))

	hello("Mustafa", 19)

	bolum , kalan := bolme(104, 5)
	fmt.Println(bolum, kalan)
}

func sum(x int, y int) int {
	return x + y
}

func hello(name string, age float32) {
	fmt.Println("Benim Adim", name, age, "yasindayim.")
}

func bolme(bolunen, bolen int) (bolum, kalan int) {
	bolum = bolunen / bolen
	kalan = bolunen % bolen
	
	return bolum, kalan
}