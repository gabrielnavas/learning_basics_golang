package main

import "fmt"

type Sporty interface {
	OnTurbo()
}

type Luxurious interface {
	MakeGoal()
}

type SportyLuxurious interface {
	Sporty
	Luxurious
}

type Bmw7 struct{}

func (b Bmw7) OnTurbo() {
	fmt.Println("Turn on Turbo")
}

func (b Bmw7) MakeGoal() {
	fmt.Println("Making Goals")
}

func main() {
	var b SportyLuxurious = Bmw7{}
	b.MakeGoal()
	b.OnTurbo()
}
