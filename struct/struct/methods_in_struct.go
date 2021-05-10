package main

import "fmt"

type Product struct {
	name     string
	price    float64
	discount float64
}

func (p Product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	apple := Product{
		name:     "Apple",
		price:    100,
		discount: 0.10,
	}
	fmt.Println(apple.priceWithDiscount())
}
