package main

import "fmt"

var sum = func(x, y int) int {
	return x + y
}

func main() {
	multiply := func(x, y int) int {
		return x * y
	}

	n := sum(1, 2)
	prodN := multiply(2, 3)

	fmt.Printf("1+2 = %v\n", n)
	fmt.Printf("2x3 = %v\n", prodN)
}
