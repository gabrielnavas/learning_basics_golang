package main

import "fmt"

type Product struct {
	ID     int
	Name   string
	Amount int
	Price  float64
}

type Order struct {
	OrderID int
	Orders  []Product
}

func (o Order) getSumPrices() (sum float64) {
	for _, order := range o.Orders {
		sum += order.Price * float64(order.Amount)
	}
	return sum
}

func main() {
	order := Order{
		OrderID: 1,
		Orders: []Product{
			{
				ID:     1,
				Name:   "Apple",
				Amount: 55,
				Price:  1.55,
			}, {
				ID:     2,
				Name:   "Banana",
				Amount: 21,
				Price:  0.75,
			}, {
				ID:     3,
				Name:   "Strawberry",
				Amount: 10,
				Price:  1.05,
			},
		},
	}
	sum := order.getSumPrices()
	fmt.Printf("sum of the prices: %.2f\n", sum)
}
