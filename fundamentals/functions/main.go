package main

import "fmt"

// to execute functions
// go run *.go
func main() {
	sum := sumTwoValues(1, 2)
	product := productTwoValues(2, 4)
	fmt.Printf("sum is: %d\n", sum)         // 3
	fmt.Printf("product is: %d\n", product) // 8
}
