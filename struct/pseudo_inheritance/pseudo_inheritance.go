package main

import "fmt"

type Car struct {
	Name     string
	MaxSpeed float64
}

type SuperCar struct {
	Car
	TurboOn bool
}

func main() {
	ferrari := SuperCar{
		Car: Car{
			Name:     "Ferrari",
			MaxSpeed: 280,
		},
		TurboOn: true,
	}
	fmt.Println(ferrari) // {{Ferrari 280} true}

	// other example
	lamborghini := SuperCar{}
	lamborghini.Name = "lamborghini"
	lamborghini.MaxSpeed = 300
	lamborghini.TurboOn = false
	fmt.Println(lamborghini) // {{lamborghini 300} false}

}
