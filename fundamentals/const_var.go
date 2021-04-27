package main

import (
	"fmt"
	m "math"
)

func main() {

	// implicit type
	const PI float64 = 3.1415

	// type (float64) inferred by the compiler
	var raio = 3.2

	// reduced way to create a variable
	// "m" is alias to "math"
	area := PI * m.Pow(raio, 2)

	// simple print
	fmt.Println("Area is: ", area)

	// block of the consts
	const (
		name = "Gabriel"
		age  = 22
	)

	// block of the variables
	var (
		country = "Brasil"
		city    = "São Paulo"
	)

	// simple print
	fmt.Println("Country", country)
	fmt.Println("City", city)

	// multi variables in line
	var isTrue, isFalse bool = true, false
	fmt.Println(isTrue, isFalse)

	// multi consts in line
	carName, carYear := "Fusca", 1986
	fmt.Println("My", carName, "is", carYear, "year old")
}
