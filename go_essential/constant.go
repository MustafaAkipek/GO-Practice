package main

import (
	"fmt"
	"math"
)

func main() {

	var r float64 = 3.0 // the value of the variable can be changed

	const pi float64 = 3.14 // the value of the variable cannot be changed

	areaOfCircle := pi * (math.Pow(r, 2))

	fmt.Println(areaOfCircle)
}