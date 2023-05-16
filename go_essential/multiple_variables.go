package main

import "fmt"

func main() {

	/* var(
		name string = "Mustafa"
		age int = 300
		isMarried bool = true

		weight float32 = 72.5
		height int = 172
	) */

	/* var name, age, isMarried, weight, height = "Mustafa", 40, true, 72.5, 172 */

	name, age, isMarried, weight, height := "Mustafa", 40, true, 72.5, 172

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)
	fmt.Println(height)
}
