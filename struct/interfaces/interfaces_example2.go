package main

type Car interface {
	OnPower()
}

type Ferrari struct {
	Name string
	IsOn bool
}

func (p *Ferrari) OnPower() {
	p.IsOn = true
}

func main() {

	// type Ferrari
	person := Ferrari{
		Name: "Apple",
		IsOn: false,
	}

	// literal Ferriri need to be address
	var product Car = &Ferrari{
		Name: "Apple",
		IsOn: false,
	}

	person.OnPower()
	product.OnPower()

}
