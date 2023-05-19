package main

import "fmt"

var packVar = "Package Scope"
var funcVar = "Func(Package) Scope"

func main() {

	var funcVar = "Func Scope"

	fmt.Println(funcVar)
	fmt.Println(packVar)

	myFunc()
}

func myFunc() {
	fmt.Println(funcVar)
}