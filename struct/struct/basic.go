package main

import "fmt"

type Product struct {
	name     string
	price    float64
	discount float64
}

func main() {
	apple := Product{"Apple", 16.22, 0.11}
	fmt.Println(apple)
	fmt.Println(apple.name)
	fmt.Println(apple.discount)
	fmt.Println(apple.price)
}
