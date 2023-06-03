package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {

		if i%5 == 0 {
			continue
		}

		fmt.Println(i)
	}

	for j := 0; j <= 10; j++ {

		if j == 3 {
			break
		}

		fmt.Println(j)
	}
}
