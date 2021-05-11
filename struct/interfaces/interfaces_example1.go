package main

import (
	"fmt"
)

type Printable interface {
	toString() string
}

type Person struct {
	FirstName string
	LastName  string
}

type Product struct {
	Name  string
	Price float64
}

func (p Person) toString() string {
	return p.FirstName + " " + p.LastName
}

func (p Product) toString() string {
	return p.Name + ": " + fmt.Sprint(p.Price)
}

func print(o Printable) {
	fmt.Println(o.toString())
}

func main() {
	var person Printable = Person{
		FirstName: "Gabriel",
		LastName:  "Silva",
	}

	var product Printable = Product{
		Name:  "Apple",
		Price: 1.21,
	}

	print(person)
	print(product)
}
