package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) setFirstName(firstName string) {
	p.FirstName = firstName
}

func (p *Person) setLastName(lastName string) {
	p.LastName = lastName
}

func main() {
	gabriel := Person{
		FirstName: "Gabriel",
		LastName:  "Silva",
	}

	gabriel.setFirstName("John")
	fmt.Println(gabriel) // {Gabriel Silva}

	gabriel.setLastName("Nozes") // change the last name by pointer reference
	fmt.Println(gabriel)         // {Gabriel Nozes}
}
