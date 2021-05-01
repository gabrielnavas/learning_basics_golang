package main

import "fmt"

func main() {
	//slice não é array, ele define um pedaço de um array
	// myHouses tem referência para posicao [1:3] (3 não incluso)
	houses := []int{55, 77, 100, 22, 88}
	myHouses := houses[1:3] // [77, 100]
	myHouses[0] = 99
	myHouses[1] = 150
	fmt.Println(houses)   // [55, 99, 150, 22, 88]
	fmt.Println(myHouses) // [99, 150]
}
